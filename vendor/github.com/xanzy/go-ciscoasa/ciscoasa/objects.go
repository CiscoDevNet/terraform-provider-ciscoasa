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
	"errors"
	"fmt"
	"regexp"
	"strings"
)

type objectsService struct {
	*Client
}
type Periodic struct {
	Frequency   string `json:"frequency"`
	StartHour   int    `json:"startHour"`
	StartMinute int    `json:"startMinute"`
	EndHour     int    `json:"endHour"`
	EndMinute   int    `json:"endMinute"`
}

func (s *objectsService) objectFromAddress(address string) (*AddressObject, error) {
	if address == "" {
		return nil, errors.New("an address cannot be an empty string")
	}

	o := &AddressObject{}

	// Test if the address is referencing a network object.
	if n, err := s.GetNetworkObject(address); err == nil {
		o.Kind = strings.Replace(n.Kind, "object#", "objectRef#", 1)
		o.ObjectID = address
		return o, nil
	}

	// Test if the address is referencing a network object group.
	if n, err := s.GetNetworkObjectGroup(address); err == nil {
		o.Kind = strings.Replace(n.Kind, "object#", "objectRef#", 1)
		o.ObjectID = address
		return o, nil
	}

	// Test if the address kind can be inferred from it's value.
	if k, err := kindFromValue(address); err == nil {
		if strings.HasSuffix(o.Kind, "Range") {
			return nil, errors.New("an address cannot be an IP range")
		}
		o.Kind = k
		o.Value = address
		return o, nil
	}

	return nil, fmt.Errorf("failed to create object from address %q", address)
}

var regexpPorts = regexp.MustCompile("^[0-9]+(-[0-9]+)?")

func (s *objectsService) objectFromService(service string) (*ServiceObject, error) {
	if service == "" {
		return nil, nil
	}

	o := &ServiceObject{}
	parts := strings.SplitN(service, "/", 2)

	if len(parts) == 1 {
		// Test if the service is referencing a network service.
		if n, err := s.GetNetworkService(parts[0]); err == nil {
			o.Kind = strings.Replace(n.Kind, "object#", "objectRef#", 1)
			o.ObjectID = n.ObjectID
			return o, nil
		}

		// Test is the service is referencing a network service group.
		if n, err := s.GetNetworkServiceGroup(parts[0]); err == nil {
			o.Kind = strings.Replace(n.Kind, "object#", "objectRef#", 1)
			o.ObjectID = n.ObjectID
			return o, nil
		}
	}

	p, err := protocolFromValue(parts[0])
	if err != nil {
		return nil, err
	}

	if len(parts) == 1 || parts[1] == "" {
		o.Kind = "NetworkProtocol"
		o.Value = p.Prefix
		return o, nil
	}

	o.Kind = p.Kind
	o.Value = fmt.Sprintf("%s/%s", p.Prefix, parts[1])

	return o, nil
}

var regexpService = regexp.MustCompile(`^<?(tcp|udp|tcpudp|icmp6?)?(/)?[a-zA-Z0-9_]*(,)?.*`)

func (s *objectsService) kindFromService(service string) (*ServiceObject, error) {

	o := &ServiceObject{}

	submatches := regexpService.FindStringSubmatch(service)

	if protocol := submatches[1]; protocol != "" {
		p, err := protocolFromValue(protocol)
		if err != nil {
			return nil, err
		}
		o.Kind = fmt.Sprintf("object#%sObj", p.Kind)
	} else {
		o.Kind = "object#NetworkProtocolObj"
	}

	protoDelim, srcDelim := submatches[2], submatches[3]
	if protoDelim == "" && srcDelim == "," {
		parts := strings.SplitN(service, ",", 2)
		o.Value = fmt.Sprintf("%s/1-65535,%s", parts[0], parts[1])
	} else {
		o.Value = service
	}

	return o, nil
}
