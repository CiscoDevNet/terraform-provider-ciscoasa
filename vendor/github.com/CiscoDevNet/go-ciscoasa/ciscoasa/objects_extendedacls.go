//
// Copyright 2017, Rutger te Nijenhuis & Sander van Harmelen
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package ciscoasa

import "fmt"

// ExtendedACLObjectCollection represents a collection of access control list objects.
type ExtendedACLObjectCollection struct {
	RangeInfo RangeInfo            `json:"rangeInfo"`
	Items     []*ExtendedACLObject `json:"items"`
	Kind      string               `json:"kind"`
	SelfLink  string               `json:"selfLink"`
}

// ExtendedACLObject represents an access control list object.
type ExtendedACLObject struct {
	Name     string `json:"name,omitempty"`
	Kind     string `json:"kind,omitempty"`
	ObjectID string `json:"objectId,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// ExtendedACEObjectCollection represents a collection of access control element objects.
type ExtendedACEObjectCollection struct {
	RangeInfo RangeInfo            `json:"rangeInfo"`
	Items     []*ExtendedACEObject `json:"items"`
	Kind      string               `json:"kind"`
	SelfLink  string               `json:"selfLink"`
}

// ExtendedACEObject represents an access control element object
type ExtendedACEObject struct {
	SrcAddress   *AddressObject `json:"sourceAddress,omitempty"`
	SrcService   *ServiceObject `json:"sourceService,omitempty"`
	DstAddress   *AddressObject `json:"destinationAddress,omitempty"`
	DstService   *ServiceObject `json:"destinationService,omitempty"`
	RuleLogging  *RuleLogging   `json:"ruleLogging,omitempty"`
	TimeRange    *TimeRange     `json:"timeRange,omitempty"`
	Position     int            `json:"position,omitempty"`
	Permit       bool           `json:"permit,omitempty"`
	Active       bool           `json:"active"`
	IsAccessRule bool           `json:"isAccessRule"`
	Remarks      []string       `json:"remarks,omitempty"`
	Kind         string         `json:"kind,omitempty"`
	ObjectID     string         `json:"objectId,omitempty"`
	SelfLink     string         `json:"selfLink,omitempty"`
}

// RuleLogging represents the rule logging settings
type RuleLogging struct {
	LogStatus   string `json:"logStatus,omitempty"`
	LogInterval int    `json:"logInterval,omitempty"`
}

// ListExtendedACLs returns a collection of access control list objects.
func (s *objectsService) ListExtendedACLs() (*ExtendedACLObjectCollection, error) {
	result := &ExtendedACLObjectCollection{}
	page := 0

	for {
		offset := page * s.pageLimit

		u := fmt.Sprintf("/api/objects/extendedacls?limit=%d&offset=%d", s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		e := &ExtendedACLObjectCollection{}
		_, err = s.do(req, e)
		if err != nil {
			return nil, err
		}

		result.RangeInfo = e.RangeInfo
		result.Items = append(result.Items, e.Items...)
		result.Kind = e.Kind
		result.SelfLink = e.SelfLink

		if e.RangeInfo.Offset+e.RangeInfo.Limit == e.RangeInfo.Total {
			break
		}
		page++
	}

	return result, nil
}

// ListExtendedACLACEs returns a collection of access control element objects.
func (s *objectsService) ListExtendedACLACEs(aclName string) (*ExtendedACEObjectCollection, error) {
	result := &ExtendedACEObjectCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/objects/extendedacls/%s/aces?limit=%d&offset=%d", aclName, s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		e := &ExtendedACEObjectCollection{}
		_, err = s.do(req, e)
		if err != nil {
			return nil, err
		}

		result.RangeInfo = e.RangeInfo
		result.Items = append(result.Items, e.Items...)
		result.Kind = e.Kind
		result.SelfLink = e.SelfLink

		if e.RangeInfo.Offset+e.RangeInfo.Limit == e.RangeInfo.Total {
			break
		}
		page++
	}

	return result, nil
}

type CreateExtendedACLACEOptions struct {
	Active      bool
	Dst         string
	DstService  string
	Permit      bool
	Remarks     []string
	Position    int
	RuleLogging *RuleLogging
	Src         string
	SrcService  string
}

// CreateExtendedACLACE creates an access control element.
func (s *objectsService) CreateExtendedACLACE(aclName string, options CreateExtendedACLACEOptions) (string, error) {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces", aclName)

	e := &ExtendedACEObject{
		Active:      options.Active,
		Permit:      options.Permit,
		Remarks:     options.Remarks,
		Position:    options.Position,
		RuleLogging: options.RuleLogging,
		Kind:        "object#ExtendedACE",
	}

	var err error
	if e.DstAddress, err = s.Objects.objectFromAddress(options.Dst); err != nil {
		return "", err
	}
	if e.DstService, err = s.Objects.objectFromService(options.DstService); err != nil {
		return "", err
	}
	if e.SrcAddress, err = s.Objects.objectFromAddress(options.Src); err != nil {
		return "", err
	}
	if e.SrcService, err = s.Objects.objectFromService(options.SrcService); err != nil {
		return "", err
	}

	req, err := s.newRequest("POST", u, e)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// GetExtendedACLACE retrieves an access control element.
func (s *objectsService) GetExtendedACLACE(aclName string, aceID string) (*ExtendedACEObject, error) {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces/%s", aclName, aceID)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	e := &ExtendedACEObject{}
	_, err = s.do(req, e)

	return e, err
}

type UpdateExtendedACLACEOptions struct {
	Active      bool
	Dst         string
	DstService  string
	Permit      bool
	Remarks     []string
	Position    int
	RuleLogging *RuleLogging
	Src         string
	SrcService  string
}

// UpdateExtendedACLACE updates an access control element.
func (s *objectsService) UpdateExtendedACLACE(aclName, aceID string, options UpdateExtendedACLACEOptions) (string, error) {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces/%s", aclName, aceID)

	e := &ExtendedACEObject{
		Active:      options.Active,
		Permit:      options.Permit,
		Remarks:     options.Remarks,
		Position:    options.Position,
		RuleLogging: options.RuleLogging,
		Kind:        "object#ExtendedACE",
	}

	var err error
	if e.DstAddress, err = s.Objects.objectFromAddress(options.Dst); err != nil {
		return "", err
	}
	if e.DstService, err = s.Objects.objectFromService(options.DstService); err != nil {
		return "", err
	}
	if e.SrcAddress, err = s.Objects.objectFromAddress(options.Src); err != nil {
		return "", err
	}
	if e.SrcService, err = s.Objects.objectFromService(options.SrcService); err != nil {
		return "", err
	}

	req, err := s.newRequest("POST", u, e)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// DeleteExtendedACLACE deletes an access control element.
func (s *objectsService) DeleteExtendedACLACE(aclName string, aceID string) error {
	u := fmt.Sprintf("/api/objects/extendedacls/%s/aces/%s", aclName, aceID)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
