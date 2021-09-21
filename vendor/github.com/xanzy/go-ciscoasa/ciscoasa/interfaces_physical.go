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

// PhysicalInterfaceCollection represents a collection of physical interfaces.
type PhysicalInterfaceCollection struct {
	RangeInfo RangeInfo            `json:"rangeInfo"`
	Items     []*PhysicalInterface `json:"items"`
	Kind      string               `json:"kind"`
	SelfLink  string               `json:"selfLink"`
}

// PhysicalInterface represents an interface.
type PhysicalInterface struct {
	HardwareID        string     `json:"hardwareID"`
	InterfaceDesc     string     `json:"interfaceDesc"`
	ChannelGroupID    string     `json:"channelGroupID"`
	ChannelGroupMode  string     `json:"channelGroupMode"`
	Duplex            string     `json:"duplex,omitempty"`
	FlowcontrolOn     bool       `json:"flowcontrolOn"`
	FlowcontrolHigh   int        `json:"flowcontrolHigh"`
	FlowcontrolLow    int        `json:"flowcontrolLow"`
	FlowcontrolPeriod int        `json:"flowcontrolPeriod"`
	ForwardTrafficCX  bool       `json:"forwardTrafficCX"`
	ForwardTrafficSFR bool       `json:"forwardTrafficSFR"`
	LacpPriority      int        `json:"lacpPriority"`
	ActiveMacAddress  string     `json:"activeMacAddress"`
	StandByMacAddress string     `json:"standByMacAddress"`
	ManagementOnly    bool       `json:"managementOnly"`
	Mtu               int        `json:"mtu"`
	Name              string     `json:"name"`
	SecurityLevel     int        `json:"securityLevel"`
	Shutdown          bool       `json:"shutdown"`
	Speed             string     `json:"speed,omitempty"`
	IPAddress         *IPAddress `json:"ipAddress,omitempty"`
	Ipv6Info          *IPv6Info  `json:"ipv6Info,omitempty"`
	Kind              string     `json:"kind"`
	ObjectID          string     `json:"objectId,omitempty"`
	SelfLink          string     `json:"selfLink,omitempty"`
}

// ListPhysicalInterfaces returns a collection of interfaces.
func (s *interfaceService) ListPhysicalInterfaces() (*PhysicalInterfaceCollection, error) {
	result := &PhysicalInterfaceCollection{}
	page := 0

	for {
		offset := page * s.pageLimit
		u := fmt.Sprintf("/api/interfaces/physical?limit=%d&offset=%d", s.pageLimit, offset)

		req, err := s.newRequest("GET", u, nil)
		if err != nil {
			return nil, err
		}

		e := &PhysicalInterfaceCollection{}
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

// UpdatePhysicalInterface updates a physical interface
// as there is no way of creating a physical interface.
func (s *interfaceService) UpdatePhysicalInterface(
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
) (string, error) {
	u := fmt.Sprintf("/api/interfaces/physical/%s", objectID)

	duplex := "auto"
	speed := "auto"
	if kind == "object#TenGigInterface" {
		duplex = ""
		speed = ""
	}

	r := &PhysicalInterface{
		ActiveMacAddress:  activeMacAddress,
		ChannelGroupID:    "",
		ChannelGroupMode:  "active",
		Duplex:            duplex,
		FlowcontrolHigh:   -1,
		FlowcontrolLow:    -1,
		FlowcontrolOn:     false,
		FlowcontrolPeriod: -1,
		ForwardTrafficCX:  forwardTrafficCX,
		ForwardTrafficSFR: forwardTrafficSFR,
		HardwareID:        hardwareID,
		InterfaceDesc:     interfaceDesc,
		IPAddress:         ipAddress,
		Ipv6Info:          ipv6Info,
		Kind:              kind,
		LacpPriority:      -1,
		ManagementOnly:    managementOnly,
		Mtu:               mtu,
		Name:              name,
		ObjectID:          objectID,
		SecurityLevel:     securityLevel,
		Shutdown:          shutdown,
		Speed:             speed,
		StandByMacAddress: standByMacAddress,
	}

	req, err := s.newRequest("PUT", u, r)
	if err != nil {
		return "", err
	}

	resp, err := s.do(req, nil)
	if err != nil {
		return "", err
	}

	return idFromResponse(resp)
}

// GetPhysicalInterface retrieves a physical interface.
func (s *interfaceService) GetPhysicalInterface(objectID string) (*PhysicalInterface, error) {
	u := fmt.Sprintf("/api/interfaces/physical/%s", objectID)

	req, err := s.newRequest("GET", u, nil)
	if err != nil {
		return nil, err
	}

	r := &PhysicalInterface{}
	_, err = s.do(req, r)

	return r, err
}

// DeletePhysicalInterface sets values to defaults
// as there is no way of deletion a physical interface.
func (s *interfaceService) DeletePhysicalInterface(
	hardwareID string,
	kind string,
	objectID string,
) error {
	u := fmt.Sprintf("/api/interfaces/physical/%s", objectID)

	duplex := "auto"
	speed := "auto"
	if kind == "object#TenGigInterface" {
		duplex = ""
		speed = ""
	}

	r := &PhysicalInterface{
		ActiveMacAddress:  "",
		ChannelGroupID:    "",
		ChannelGroupMode:  "active",
		Duplex:            duplex,
		FlowcontrolHigh:   -1,
		FlowcontrolLow:    -1,
		FlowcontrolOn:     false,
		FlowcontrolPeriod: -1,
		ForwardTrafficCX:  false,
		ForwardTrafficSFR: false,
		HardwareID:        hardwareID,
		InterfaceDesc:     "",
		IPAddress:         nil,
		Ipv6Info:          nil,
		Kind:              kind,
		LacpPriority:      -1,
		ManagementOnly:    false,
		Mtu:               1500,
		Name:              "",
		SecurityLevel:     -1,
		Shutdown:          false,
		Speed:             speed,
		StandByMacAddress: "",
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
