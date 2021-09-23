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

// NetworkServiceCollection represents a collection of network services.
type NetworkServiceCollection struct {
	RangeInfo RangeInfo         `json:"rangeInfo"`
	Items     []*NetworkService `json:"items"`
	Kind      string            `json:"kind"`
	SelfLink  string            `json:"selfLink"`
}

// NetworkService represents a network service.
type NetworkService struct {
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	Value       string `json:"value"`
	Kind        string `json:"kind"`
	ObjectID    string `json:"objectId,omitempty"`
	SelfLink    string `json:"selfLink,omitempty"`
}

// ListNetworkServices returns a collection of network services.
func (s *objectsService) ListNetworkServices() (*NetworkServiceCollection, error) {
	result := &NetworkServiceCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/objects/networkservices?limit=%d&offset=%d", s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		n := &NetworkServiceCollection{}
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

// CreateNetworkService creates a new network service.
func (s *objectsService) CreateNetworkService(name, description, service string) (*NetworkService, error) {
	u := "/api/objects/networkobjectgroups"

	o, err := s.objectFromService(service)
	if err != nil {
		return nil, err
	}

	n := &NetworkService{
		Name:        name,
		Description: description,
		Value:       o.Value,
		Kind:        o.Kind,
		ObjectID:    o.ObjectID,
	}

	req, err := s.newRequest("POST", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkService(name)
}

// GetNetworkService retrieves a network service.
func (s *objectsService) GetNetworkService(name string) (*NetworkService, error) {
	u := fmt.Sprintf("/api/objects/networkservices/%s", name)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	n := &NetworkService{}
	_, err = s.do(req, n)

	return n, err
}

// UpdateNetworkService updates a network service.
func (s *objectsService) UpdateNetworkService(name, description, service string) (*NetworkService, error) {
	u := fmt.Sprintf("/api/objects/networkservices/%s", name)

	o, err := s.objectFromService(service)
	if err != nil {
		return nil, err
	}

	n := &NetworkService{
		Name:        name,
		Description: description,
		Value:       o.Value,
		Kind:        o.Kind,
		ObjectID:    o.ObjectID,
	}

	req, err := s.newRequest("PUT", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkService(name)
}

// DeleteNetworkService deletes a network server.
func (s *objectsService) DeleteNetworkService(name string) error {
	u := fmt.Sprintf("/api/objects/networkservices/%s", name)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
