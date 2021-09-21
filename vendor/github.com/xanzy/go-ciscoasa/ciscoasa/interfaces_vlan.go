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

// VlanInterfaceCollection represents a collection of vlan interfaces.
type VlanInterfaceCollection struct {
	RangeInfo RangeInfo        `json:"rangeInfo"`
	Items     []*VlanInterface `json:"items"`
	Kind      string           `json:"kind"`
	SelfLink  string           `json:"selfLink"`
}

// VlanInterface represents an vlan interface.
type VlanInterface struct {
	HardwareID        string     `json:"hardwareID"`
	InterfaceDesc     string     `json:"interfaceDesc"`
	ForwardTrafficCX  bool       `json:"forwardTrafficCX"`
	ForwardTrafficSFR bool       `json:"forwardTrafficSFR"`
	ActiveMacAddress  string     `json:"activeMacAddress"`
	StandByMacAddress string     `json:"standByMacAddress"`
	ManagementOnly    bool       `json:"managementOnly"`
	Mtu               int        `json:"mtu"`
	Name              string     `json:"name"`
	SecurityLevel     int        `json:"securityLevel"`
	Shutdown          bool       `json:"shutdown"`
	VlanID            int        `json:"vlanID"`
	IPAddress         *IPAddress `json:"ipAddress,omitempty"`
	Ipv6Info          *IPv6Info  `json:"ipv6Info,omitempty"`
	Kind              string     `json:"kind"`
	ObjectID          string     `json:"objectId,omitempty"`
	SelfLink          string     `json:"selfLink,omitempty"`
}

// ListVlanInterfaces returns a collection of interfaces.
func (s *interfaceService) ListVlanInterfaces() (*VlanInterfaceCollection, error) {
	result := &VlanInterfaceCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/interfaces/vlan?limit=%d&offset=%d", s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		e := &VlanInterfaceCollection{}
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

// CreateVlanInterface creates a vlan interface.
func (s *interfaceService) CreateVlanInterface(
	activeMacAddress string,
	forwardTrafficCX bool,
	forwardTrafficSFR bool,
	hardwareID string,
	interfaceDesc string,
	ipAddress *IPAddress,
	ipv6Info *IPv6Info,
	kind string,
	managementOnly bool,
	mtu int,
	name string,
	securityLevel int,
	shutdown bool,
	standByMacAddress string,
	vlanId int,
) (string, error) {
	u := "/api/interfaces/vlan/"

	r := &VlanInterface{
		ActiveMacAddress:  activeMacAddress,
		ForwardTrafficCX:  forwardTrafficCX,
		ForwardTrafficSFR: forwardTrafficSFR,
		HardwareID:        hardwareID,
		InterfaceDesc:     interfaceDesc,
		IPAddress:         ipAddress,
		Ipv6Info:          ipv6Info,
		Kind:              kind,
		ManagementOnly:    managementOnly,
		Mtu:               mtu,
		Name:              name,
		SecurityLevel:     securityLevel,
		Shutdown:          shutdown,
		StandByMacAddress: standByMacAddress,
		VlanID:            vlanId,
	}

	req, err := s.newRequest("POST", u, r)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// UpdateVlanInterface updates a vlan interface.
func (s *interfaceService) UpdateVlanInterface(
	activeMacAddress string,
	forwardTrafficCX bool,
	forwardTrafficSFR bool,
	hardwareID string,
	interfaceDesc string,
	ipAddress *IPAddress,
	ipv6Info *IPv6Info,
	kind string,
	managementOnly bool,
	mtu int,
	name string,
	objectID string,
	securityLevel int,
	shutdown bool,
	standByMacAddress string,
	vlanId int,
) error {
	u := fmt.Sprintf("/api/interfaces/vlan/%s", objectID)

	r := &VlanInterface{
		ActiveMacAddress:  activeMacAddress,
		ForwardTrafficCX:  forwardTrafficCX,
		ForwardTrafficSFR: forwardTrafficSFR,
		HardwareID:        hardwareID,
		InterfaceDesc:     interfaceDesc,
		IPAddress:         ipAddress,
		Ipv6Info:          ipv6Info,
		Kind:              kind,
		ManagementOnly:    managementOnly,
		Mtu:               mtu,
		Name:              name,
		ObjectID:          objectID,
		SecurityLevel:     securityLevel,
		Shutdown:          shutdown,
		StandByMacAddress: standByMacAddress,
		VlanID:            vlanId,
	}

	req, err := s.newRequest("PUT", u, r)
	if err != nil {
		return err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return err
	}

	err = checkResponse(resp)

	return err
}

// GetVlanInterface retrieves a vlan interface.
func (s *interfaceService) GetVlanInterface(objectID string) (*VlanInterface, error) {
	u := fmt.Sprintf("/api/interfaces/vlan/%s", objectID)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &VlanInterface{}
	_, err = s.do(req, r)

	return r, err
}

// DeleteVlanInterface deletes a vlan interface.
func (s *interfaceService) DeleteVlanInterface(objectID string) error {
	u := fmt.Sprintf("/api/interfaces/vlan/%s", objectID)

	req, err := s.newRequest("DELETE", u, nil)
	if err != nil {
		return err
	}

	_, err = s.do(req, nil)

	return err
}
