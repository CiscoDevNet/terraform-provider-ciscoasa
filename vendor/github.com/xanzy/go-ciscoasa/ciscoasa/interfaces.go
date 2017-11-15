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
	"encoding/json"
	"net"
	"strconv"
)

type interfaceService struct {
	*Client
}

// IPAddress represents an IPv4 address.
type IPAddress struct {
	IP struct {
		Kind  string `json:"kind"`
		Value string `json:"value"`
	} `json:"ip"`
	NetMask struct {
		Kind  string `json:"kind"`
		Value string `json:"value"`
	} `json:"netMask"`
	Kind string `json:"kind"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ip *IPAddress) UnmarshalJSON(b []byte) error {
	type alias IPAddress
	if err := json.Unmarshal(b, (*alias)(ip)); err != nil {
		ip = nil
	}
	return nil
}

func (ip *IPAddress) String() string {
	n := net.IPMask(net.ParseIP(ip.NetMask.Value).To4())
	b, _ := n.Size()
	bitsize := strconv.Itoa(b)

	return ip.IP.Value + "/" + bitsize
}

// IPv6Info represents an IPv6 address.
type IPv6Info struct {
	Enabled                  bool     `json:"enabled"`
	AutoConfig               bool     `json:"autoConfig"`
	EnforceEUI64             bool     `json:"enforceEUI64"`
	ManagedAddressConfig     bool     `json:"managedAddressConfig"`
	NsInterval               int      `json:"nsInterval"`
	DadAttempts              int      `json:"dadAttempts"`
	NDiscoveryPrefixList     []string `json:"nDiscoveryPrefixList"`
	OtherStatefulConfig      bool     `json:"otherStatefulConfig"`
	RouterAdvertInterval     int      `json:"routerAdvertInterval"`
	RouterAdvertIntervalUnit string   `json:"routerAdvertIntervalUnit"`
	RouterAdvertLifetime     int      `json:"routerAdvertLifetime"`
	SuppressRouterAdvert     bool     `json:"suppressRouterAdvert"`
	ReachableTime            int      `json:"reachableTime"`
	Ipv6Addresses            []string `json:"ipv6Addresses"`
	Kind                     string   `json:"kind"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ip *IPv6Info) UnmarshalJSON(b []byte) error {
	type alias IPv6Info
	if err := json.Unmarshal(b, (*alias)(ip)); err != nil {
		ip = nil
	}
	return nil
}
