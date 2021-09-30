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

type InterfaceRef struct {
	Kind     string `json:"kind"`
	RefLink  string `json:"refLink,omitempty"`
	ObjectId string `json:"objectId,omitempty"`
	Name     string `json:"name,omitempty"`
}

// SlaTracking represents an SlaTracking Settings.
type SlaTracking struct {
	SlaId                   int    `json:"slaId"`
	TrackedIP               string `json:"trackedIP"`
	FrequencyInSeconds      int    `json:"frequencyInSeconds"`
	DataSizeInBytes         int    `json:"dataSizeInBytes"`
	ThresholdInMilliseconds int    `json:"thresholdInMilliseconds"`
	ToS                     int    `json:"ToS"`
	TimeoutInMilliseconds   int    `json:"timeoutInMilliseconds"`
	NumPackets              int    `json:"numPackets"`
}

// DhcpClient represents an DHCP Settings.
type DhcpClient struct {
	SetDefaultRoute     bool         `json:"setDefaultRoute"`
	Metric              int          `json:"metric"`
	PrimaryTrackId      int          `json:"primaryTrackId"`
	TrackingEnabled     bool         `json:"trackingEnabled"`
	SlaTrackingSettings *SlaTracking `json:"slaTrackingSettings"`
}

// Address represents a static IPv4/IPv6 address settings.
type Address struct {
	Kind  string `json:"kind"`
	Value string `json:"value"`
}

// IPAddress represents an IP address settings.
type IPAddress struct {
	IP                 *Address    `json:"ip,omitempty"`
	NetMask            *Address    `json:"netMask,omitempty"`
	Kind               string      `json:"kind"`
	DhcpOptionUsingMac bool        `json:"dhcpOptionUsingMac,omitempty"`
	DhcpBroadcast      bool        `json:"dhcpBroadcast,omitempty"`
	DhcpClient         *DhcpClient `json:"dhcpClient,omitempty"`
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

// nDiscoveryPrefix represents an nDiscoveryPrefix list.
type NDiscoveryPrefix struct {
	OffLink           bool   `json:"offLink"`
	NoAdvertise       bool   `json:"noAdvertise"`
	PreferredLifetime int    `json:"preferredLifetime"`
	ValidLifetime     int    `json:"validLifetime"`
	HasDuration       bool   `json:"hasDuration"`
	DefaultPrefix     bool   `json:"defaultPrefix"`
	Kind              string `json:"kind"`
}

// Ipv6Address represents an Ipv6Address.
type Ipv6Address struct {
	PrefixLength int      `json:"prefixLength,omitempty"`
	Standby      *Address `json:"standby,omitempty"`
	Address      *Address `json:"address,omitempty"`
	// IsEUI64      bool     `json:"isEUI64"`
	Kind string `json:"kind"`
}

// IPv6Info represents an IPv6 address.
type IPv6Info struct {
	Enabled                  bool                `json:"enabled"`
	AutoConfig               bool                `json:"autoConfig"`
	EnforceEUI64             bool                `json:"enforceEUI64"`
	ManagedAddressConfig     bool                `json:"managedAddressConfig"`
	NsInterval               int                 `json:"nsInterval"`
	DadAttempts              int                 `json:"dadAttempts"`
	NDiscoveryPrefixList     []*NDiscoveryPrefix `json:"nDiscoveryPrefixList,omitempty"`
	OtherStatefulConfig      bool                `json:"otherStatefulConfig"`
	RouterAdvertInterval     int                 `json:"routerAdvertInterval"`
	RouterAdvertIntervalUnit string              `json:"routerAdvertIntervalUnit"`
	RouterAdvertLifetime     int                 `json:"routerAdvertLifetime"`
	SuppressRouterAdvert     bool                `json:"suppressRouterAdvert"`
	ReachableTime            int                 `json:"reachableTime"`
	LinkLocalAddress         *Ipv6Address        `json:"linkLocalAddress,omitempty"`
	Ipv6Addresses            []*Ipv6Address      `json:"ipv6Addresses,omitempty"`
	Kind                     string              `json:"kind"`
}

// UnmarshalJSON implements json.Unmarshaler interface.
func (ip *IPv6Info) UnmarshalJSON(b []byte) error {
	type alias IPv6Info
	if err := json.Unmarshal(b, (*alias)(ip)); err != nil {
		ip = nil
	}
	return nil
}
