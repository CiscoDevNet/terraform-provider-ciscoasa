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

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// NetworkServiceGroupCollection represents a collection of network service groups.
type NetworkServiceGroupCollection struct {
	RangeInfo RangeInfo              `json:"rangeInfo"`
	Items     []*NetworkServiceGroup `json:"items"`
	Kind      string                 `json:"kind"`
	SelfLink  string                 `json:"selfLink"`
}

// NetworkServiceGroup represents a network service group.
type NetworkServiceGroup struct {
	Name        string           `json:"name"`
	Description string           `json:"description,omitempty"`
	Members     []*ServiceObject `json:"members"`
	Kind        string           `json:"kind"`
	ObjectID    string           `json:"objectId,omitempty"`
	SelfLink    string           `json:"selfLink,omitempty"`
}

// ServiceObject represents an service object.
type ServiceObject struct {
	Value    string `json:"value,omitempty"`
	Kind     string `json:"kind"`
	ObjectID string `json:"objectId,omitempty"`
	Reflink  string `json:"refLink,omitempty"`
}

// String returns the numerical description of the service.
func (o *ServiceObject) String() string {
	if strings.HasPrefix(o.Kind, "objectRef#") {
		return o.ObjectID
	}

	parts := strings.Split(o.Value, "/")

	if len(parts) >= 2 && !regexpPorts.MatchString(parts[1]) {
		switch parts[0] {
		case "icmp":
			if part1, ok := icmpType[parts[1]]; ok {
				parts[1] = part1
			}
		case "tcp":
			if part1, ok := tcpType[parts[1]]; ok {
				parts[1] = part1
			} else {
				part1, err := net.LookupPort(parts[0], parts[1])
				if err == nil {
					parts[1] = strconv.Itoa(part1)
				}
			}
		case "udp":
			if part1, ok := udpType[parts[1]]; ok {
				parts[1] = part1
			} else {
				part1, err := net.LookupPort(parts[0], parts[1])
				if err == nil {
					parts[1] = strconv.Itoa(part1)
				}
			}
		}
	}

	return strings.Join(parts, "/")
}

// ListNetworkServiceGroups returns a collection of network service groups.
func (s *objectsService) ListNetworkServiceGroups() (*NetworkServiceGroupCollection, error) {
	result := &NetworkServiceGroupCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/objects/networkservicegroups?limit=%d&offset=%d", s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		n := &NetworkServiceGroupCollection{}
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

// CreateNetworkServiceGroup creates a new network service group.
func (s *objectsService) CreateNetworkServiceGroup(name, description string, members []string) (*NetworkServiceGroup, error) {
	u := "/api/objects/networkservicegroups"

	n := NetworkServiceGroup{
		Name:        name,
		Description: description,
		Kind:        "object#NetworkServiceGroup",
	}

	for _, member := range members {
		o, err := s.objectFromService(member)
		if err != nil {
			return nil, err
		}

		n.Members = append(n.Members, o)
	}

	req, err := s.newRequest("POST", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkServiceGroup(name)
}

// GetNetworkServiceGroup retrieves a network service group.
func (s *objectsService) GetNetworkServiceGroup(name string) (*NetworkServiceGroup, error) {
	u := fmt.Sprintf("/api/objects/networkservicegroups/%s", name)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	n := &NetworkServiceGroup{}
	_, err = s.do(req, n)

	return n, err
}

// UpdateNetworkServiceGroup updates a network service group.
func (s *objectsService) UpdateNetworkServiceGroup(name, description string, members []string) (*NetworkServiceGroup, error) {
	u := fmt.Sprintf("/api/objects/networkservicegroups/%s", name)

	n := NetworkServiceGroup{
		Name:        name,
		Description: description,
		Kind:        "object#NetworkServiceGroup",
	}

	for _, member := range members {
		o, err := s.objectFromService(member)
		if err != nil {
			return nil, err
		}

		n.Members = append(n.Members, o)
	}

	req, err := s.newRequest("PUT", u, n)
	if err != nil {
		return nil, err
	}

	_, err = s.do(req, nil)
	if err != nil {
		return nil, err
	}

	return s.GetNetworkServiceGroup(name)
}

// DeleteNetworkServiceGroup deletes a network service group.
func (s *objectsService) DeleteNetworkServiceGroup(name string) error {
	u := fmt.Sprintf("/api/objects/networkservicegroups/%s", name)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
