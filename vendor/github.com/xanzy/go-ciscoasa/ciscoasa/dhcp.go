package ciscoasa

type dhcpService struct {
	*Client
}

// DhcpServerCollection represents a collection of DHCP Relay Interface Servers.
type DhcpServerCollection struct {
	RangeInfo RangeInfo     `json:"rangeInfo"`
	Items     []*DhcpServer `json:"items"`
	Kind      string        `json:"kind"`
	SelfLink  string        `json:"selfLink"`
}

// DhcpServerOptions represents a ntp server options.
type DhcpServerOptions struct {
	Type   string `json:"type"`
	Code   int    `json:"code"`
	Value1 string `json:"value1"`
	Value2 string `json:"value2,omitempty"`
}

// DhcpServer represents a DHCP Relay Interface Server.
type DhcpServer struct {
	Interface             *InterfaceRef        `json:"interface"`
	Enabled               bool                 `json:"enabled"`
	PoolStartIP           string               `json:"poolStartIP"`
	PoolEndIP             string               `json:"poolEndIP"`
	DnsIP1                string               `json:"dnsIP1"`
	DnsIP2                string               `json:"dnsIP2"`
	WinsIP1               string               `json:"winsIP1"`
	WinsIP2               string               `json:"winsIP2"`
	LeaseLengthInSec      string               `json:"leaseLengthInSec,omitempty"`
	PingTimeoutInMilliSec string               `json:"pingTimeoutInMilliSec,omitempty"`
	DomainName            string               `json:"domainName,omitempty"`
	IsAutoConfigEnabled   bool                 `json:"isAutoConfigEnabled"`
	AutoConfigInterface   string               `json:"autoConfigInterface,omitempty"`
	IsVpnOverride         bool                 `json:"isVpnOverride,omitempty"`
	Options               []*DhcpServerOptions `json:"options,omitempty"`
	Ddns                  struct {
		UpdateDNSClient        bool `json:"updateDNSClient"`
		UpdateBothRecords      bool `json:"updateBothRecords,omitempty"`
		OverrideClientSettings bool `json:"overrideClientSettings,omitempty"`
	} `json:"ddns"`
	Kind     string `json:"kind"`
	ObjectId string `json:"objectId,omitempty"`
	SelfLink string `json:"selfLink,omitempty"`
}

// DhcpRelayLocalCollection represents a collection of DHCP Relay Interface Servers.
type DhcpRelayLocalCollection struct {
	RangeInfo RangeInfo         `json:"rangeInfo"`
	Items     []*DhcpRelayLocal `json:"items"`
	Kind      string            `json:"kind"`
	SelfLink  string            `json:"selfLink"`
}

// DhcpRelayLocal represents a DHCP Relay Interface Server.
type DhcpRelayLocal struct {
	Servers   []string `json:"servers"`
	Interface string   `json:"interface"`
	Kind      string   `json:"kind"`
}

// DhcpRelayGS represents a DHCP Relay Global Settings object.
type DhcpRelayGS struct {
	Ipv4Timeout            int    `json:"ipv4Timeout"`
	Ipv6Timeout            int    `json:"ipv6Timeout"`
	TrustedOnAllInterfaces bool   `json:"trustedOnAllInterfaces"`
	Kind                   string `json:"kind"`
}
