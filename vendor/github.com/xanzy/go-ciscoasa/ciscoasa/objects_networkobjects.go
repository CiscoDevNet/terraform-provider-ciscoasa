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

// NetworkObjectCollection represents a collection of network objects.
type NetworkObjectCollection struct {
	RangeInfo RangeInfo        `json:"rangeInfo"`
	Items     []*NetworkObject `json:"items"`
	Kind      string           `json:"kind"`
	SelfLink  string           `json:"selfLink"`
}

// NetworkObject represents a network object.
type NetworkObject struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Host        struct {
		Kind  string `json:"kind"`
		Value string `json:"value"`
	} `json:"host"`
	Kind     string `json:"kind"`
	ObjectID string `json:"objectId,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// ListNetworkObjects returns a collection of network objects.
func (s *objectsService) ListNetworkObjects() (*NetworkObjectCollection, error) {
	result := &NetworkObjectCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/objects/networkobjects?limit=%d&offset=%d", s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		n := &NetworkObjectCollection{}
		_, err = s.do(req, n)
		if err != nil {
			return nil, err
		}

		result.RangeInfo = n.RangeInfo
		result.Items = append(result.Items, n.Items...)
		result.Kind = n.Kind
		result.SelfLink = n.SelfLink

		if n.RangeInfo.Offset+n.RangeInfo.Limit == n.RangeInfo.Total {
			break
		}
		page++
	}

	return result, nil
}

// CreateNetworkObject creates a new network object.
func (s *objectsService) CreateNetworkObject(name, description, value string) (*NetworkObject, error) {
	u := "/api/objects/networkobjects"

	n := &NetworkObject{
		Name:        name,
		Description: description,
		Kind:        "object#NetworkObj",
	}

	kind, err := kindFromValue(value)
	if err != nil {
		return nil, err
	}

	n.Host.Kind = kind
	n.Host.Value = value

	req, err := s.newRequest("POST", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkObject(name)
}

// GetNetworkObject retrieves a network object.
func (s *objectsService) GetNetworkObject(name string) (*NetworkObject, error) {
	u := fmt.Sprintf("/api/objects/networkobjects/%s", name)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	n := &NetworkObject{}
	_, err = s.do(req, n)

	return n, err
}

// UpdateNetworkObject updates a network object.
func (s *objectsService) UpdateNetworkObject(name, description, value string) (*NetworkObject, error) {
	u := fmt.Sprintf("/api/objects/networkobjects/%s", name)

	n := &NetworkObject{
		Name:        name,
		Description: description,
		Kind:        "object#NetworkObj",
	}

	kind, err := kindFromValue(value)
	if err != nil {
		return nil, err
	}

	n.Host.Kind = kind
	n.Host.Value = value

	req, err := s.newRequest("PUT", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkObject(name)
}

// DeleteNetworkObject deletes a network object.
func (s *objectsService) DeleteNetworkObject(name string) error {
	u := fmt.Sprintf("/api/objects/networkobjects/%s", name)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
