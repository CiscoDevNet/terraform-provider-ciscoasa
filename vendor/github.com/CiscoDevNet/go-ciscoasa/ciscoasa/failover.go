package ciscoasa

type failoverService struct {
	*Client
}

// FailoverInterfacesCollection represents a collection of Failover Interfaces.
type FailoverInterfacesCollection struct {
	RangeInfo RangeInfo            `json:"rangeInfo"`
	Items     []*FailoverInterface `json:"items"`
	Kind      string               `json:"kind"`
	SelfLink  string               `json:"selfLink"`
}

// FailoverInterface represents a Failover Interface.
type FailoverInterface struct {
	InterfaceName    string   `json:"interfaceName"`
	Name             string   `json:"name"`
	ActiveIPAddress  Address  `json:"activeIPAddress"`
	SubnetMask       Address  `json:"subnetMask"`
	StandbyIPAddress *Address `json:"standbyIPAddress,omitempty"`
	IsMonitored      bool     `json:"isMonitored"`
	Kind             string   `json:"kind"`
	ObjectId         string   `json:"objectId,omitempty"`
	SelfLink         string   `json:"selfLink,omitempty"`
}

// FailoverSetup represents a Failover Setup.
type FailoverSetup struct {
	EnableFOCheck                      bool          `json:"enableFOCheck"`
	SecretKey                          string        `json:"secretKey"`
	IpSecKey                           string        `json:"ipSecKey"`
	HexKey                             bool          `json:"hexKey"`
	LanFoInterface                     *InterfaceRef `json:"lanFoInterface,omitempty"`
	LanActiveIP                        *Address      `json:"lanActiveIP,omitempty"`
	LanSubnet                          *Address      `json:"lanSubnet,omitempty"`
	LanStandby                         *Address      `json:"lanStandby,omitempty"`
	LanIFCName                         string        `json:"lanIFCName,omitempty"`
	IsLANInterfacePreferredPrimary     bool          `json:"isLANInterfacePreferredPrimary"`
	IsLANInterfacePreferredSecondary   bool          `json:"isLANInterfacePreferredSecondary"`
	StateFoInterface                   *InterfaceRef `json:"stateFoInterface,omitempty"`
	StateActiveIP                      *Address      `json:"stateActiveIP,omitempty"`
	StateSubnet                        *Address      `json:"stateSubnet,omitempty"`
	StateStandbyIP                     *Address      `json:"stateStandbyIP,omitempty"`
	StateIFCName                       string        `json:"stateIFCName,omitempty"`
	HttpReplicate                      bool          `json:"httpReplicate"`
	ReplicateRate                      int           `json:"replicateRate"`
	FailedInterfacesUnit               string        `json:"failedInterfacesUnit"`
	FailedInterfacesTriggeringFailover string        `json:"failedInterfacesTriggeringFailover"`
	UnitPollTime                       string        `json:"unitPollTime"`
	UnitPollTimeUnit                   string        `json:"unitPollTimeUnit"`
	UnitHoldTime                       string        `json:"unitHoldTime"`
	UnitHoldTimeUnit                   string        `json:"unitHoldTimeUnit"`
	MonitoredPollTime                  string        `json:"monitoredPollTime"`
	MonitoredPollTimeUnit              string        `json:"monitoredPollTimeUnit"`
	InterfaceHoldTime                  string        `json:"interfaceHoldTime"`
	Kind                               string        `json:"kind,omitempty"`
	ObjectId                           string        `json:"objectId,omitempty"`
	SelfLink                           string        `json:"selfLink,omitempty"`
}
