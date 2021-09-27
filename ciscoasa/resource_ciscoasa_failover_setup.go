package ciscoasa

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func resourceCiscoASAFailoverSetup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASAFailoverSetupUpdate,
		Read:   resourceCiscoASAFailoverSetupRead,
		Update: resourceCiscoASAFailoverSetupUpdate,
		Delete: resourceCiscoASAFailoverSetupDelete,

		Schema: map[string]*schema.Schema{
			"enable": {
				Type:         schema.TypeBool,
				Optional:     true,
				RequiredWith: []string{"lan_interface_hw_id"},
			},

			"shared_key": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"hex_key": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"ipsec_key": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"lan_interface_hw_id": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"lan_failover_name",
					"lan_primary_ip",
					"lan_secondary_ip",
					"lan_net_mask",
					"lan_preferred_role",
				},
			},

			"lan_failover_name": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"lan_interface_hw_id",
					"lan_primary_ip",
					"lan_secondary_ip",
					"lan_net_mask",
					"lan_preferred_role",
				},
			},

			"lan_primary_ip": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"lan_interface_hw_id",
					"lan_failover_name",
					"lan_secondary_ip",
					"lan_net_mask",
					"lan_preferred_role",
				},
				ValidateFunc: validation.IsIPAddress,
			},

			"lan_secondary_ip": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"lan_interface_hw_id",
					"lan_failover_name",
					"lan_primary_ip",
					"lan_net_mask",
					"lan_preferred_role",
				},
				ValidateFunc: validation.IsIPAddress,
			},

			"lan_net_mask": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"lan_interface_hw_id",
					"lan_failover_name",
					"lan_primary_ip",
					"lan_secondary_ip",
					"lan_preferred_role",
				},
				ValidateFunc: validation.IsIPAddress,
			},

			"lan_preferred_role": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"lan_interface_hw_id",
					"lan_failover_name",
					"lan_primary_ip",
					"lan_secondary_ip",
					"lan_net_mask",
				},
				ValidateFunc: validation.StringInSlice([]string{
					"primary", "secondary",
				}, false),
			},

			"state_interface_hw_id": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"state_failover_name",
					"state_primary_ip",
					"state_secondary_ip",
					"state_net_mask",
				},
			},

			"state_failover_name": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"state_interface_hw_id",
					"state_primary_ip",
					"state_secondary_ip",
					"state_net_mask",
				},
			},

			"state_primary_ip": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"state_interface_hw_id",
					"state_failover_name",
					"state_secondary_ip",
					"state_net_mask",
				},
				ValidateFunc: validation.IsIPAddress,
			},

			"state_secondary_ip": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"state_interface_hw_id",
					"state_failover_name",
					"state_primary_ip",
					"state_net_mask",
				},
				ValidateFunc: validation.IsIPAddress,
			},

			"state_net_mask": {
				Type:     schema.TypeString,
				Optional: true,
				RequiredWith: []string{
					"state_interface_hw_id",
					"state_failover_name",
					"state_primary_ip",
					"state_secondary_ip",
				},
				ValidateFunc: validation.IsIPAddress,
			},

			"http_replication": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"interface_hold_time": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "25",
			},

			"unit_hold_time": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "15",
			},

			"unit_hold_time_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "seconds",
				ValidateFunc: validation.StringInSlice([]string{
					"seconds", "milliseconds",
				}, false),
			},

			"unit_poll_time": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1",
			},

			"unit_poll_time_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "seconds",
				ValidateFunc: validation.StringInSlice([]string{
					"seconds", "milliseconds",
				}, false),
			},

			"monitored_poll_time": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "5",
			},

			"monitored_poll_time_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "seconds",
				ValidateFunc: validation.StringInSlice([]string{
					"seconds", "milliseconds",
				}, false),
			},

			"replication_rate": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  -1,
			},

			"failed_interfaces_threshold": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "1",
			},

			"failed_interfaces_threshold_unit": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Number",
				ValidateFunc: validation.StringInSlice([]string{
					"Number", "Percentage",
				}, false),
			},
		},
	}
}

// There is no real creation of the Failover Setup.
// All the changes are made on already existing object.
func resourceCiscoASAFailoverSetupUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	FailoverSetup := &ciscoasa.FailoverSetup{
		EnableFOCheck:                      d.Get("enable").(bool),
		SecretKey:                          d.Get("shared_key").(string),
		IpSecKey:                           d.Get("ipsec_key").(string),
		HexKey:                             d.Get("hex_key").(bool),
		LanIFCName:                         d.Get("lan_failover_name").(string),
		StateIFCName:                       d.Get("state_failover_name").(string),
		HttpReplicate:                      d.Get("http_replication").(bool),
		ReplicateRate:                      d.Get("replication_rate").(int),
		FailedInterfacesUnit:               d.Get("failed_interfaces_threshold_unit").(string),
		FailedInterfacesTriggeringFailover: d.Get("failed_interfaces_threshold").(string),
		UnitPollTime:                       d.Get("unit_poll_time").(string),
		UnitPollTimeUnit:                   d.Get("unit_poll_time_unit").(string),
		UnitHoldTime:                       d.Get("unit_hold_time").(string),
		UnitHoldTimeUnit:                   d.Get("unit_hold_time_unit").(string),
		MonitoredPollTime:                  d.Get("monitored_poll_time").(string),
		MonitoredPollTimeUnit:              d.Get("monitored_poll_time_unit").(string),
		InterfaceHoldTime:                  d.Get("interface_hold_time").(string),
	}

	if lanIface, ok := d.GetOk("lan_interface_hw_id"); ok {
		FailoverSetup.LanFoInterface = makeInterfaceRefObj(lanIface.(string))
	}

	if lanPrimIp, ok := d.GetOk("lan_primary_ip"); ok {
		FailoverSetup.LanActiveIP = makeIpAdressObj(lanPrimIp.(string))
	}

	if lanSecIp, ok := d.GetOk("lan_secondary_ip"); ok {
		FailoverSetup.LanStandby = makeIpAdressObj(lanSecIp.(string))
	}

	if lanNetMask, ok := d.GetOk("lan_net_mask"); ok {
		FailoverSetup.LanSubnet = makeNetmaskObj(lanNetMask.(string))
	}

	if role, ok := d.GetOk("lan_preferred_role"); ok {
		if role.(string) == "primary" {
			FailoverSetup.IsLANInterfacePreferredPrimary = true
			FailoverSetup.IsLANInterfacePreferredSecondary = false
		} else {
			FailoverSetup.IsLANInterfacePreferredPrimary = false
			FailoverSetup.IsLANInterfacePreferredSecondary = true
		}
	}

	if stateIface, ok := d.GetOk("state_interface_hw_id"); ok {
		FailoverSetup.StateFoInterface = makeInterfaceRefObj(stateIface.(string))
	}

	if statePrimIp, ok := d.GetOk("state_primary_ip"); ok {
		FailoverSetup.StateActiveIP = makeIpAdressObj(statePrimIp.(string))
	}

	if stateSecIp, ok := d.GetOk("state_secondary_ip"); ok {
		FailoverSetup.StateStandbyIP = makeIpAdressObj(stateSecIp.(string))
	}

	if stateNetMask, ok := d.GetOk("state_net_mask"); ok {
		FailoverSetup.StateSubnet = makeNetmaskObj(stateNetMask.(string))
	}

	err := ca.Failover.UpdateFailoverSetup(FailoverSetup)

	if err != nil {
		return fmt.Errorf(
			"Error creating Failover Setup: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return resourceCiscoASAFailoverSetupRead(d, meta)
}

func resourceCiscoASAFailoverSetupRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.Failover.GetFailoverSetup()
	if err != nil {
		return fmt.Errorf(
			"Error reading Failover Setup: %v", err)
	}

	d.Set("enable", r.EnableFOCheck)
	d.Set("shared_key", r.SecretKey)
	d.Set("ipsec_key", r.IpSecKey)
	d.Set("hex_key", r.HexKey)
	d.Set("state_failover_name", r.StateIFCName)
	d.Set("http_replication", r.HttpReplicate)
	d.Set("replication_rate", r.ReplicateRate)
	d.Set("failed_interfaces_threshold_unit", r.FailedInterfacesUnit)
	d.Set("failed_interfaces_threshold", r.FailedInterfacesTriggeringFailover)
	d.Set("unit_poll_time", r.UnitPollTime)
	d.Set("unit_poll_time_unit", r.UnitPollTimeUnit)
	d.Set("unit_hold_time", r.UnitHoldTime)
	d.Set("unit_hold_time_unit", r.UnitHoldTimeUnit)
	d.Set("monitored_poll_time", r.MonitoredPollTime)
	d.Set("monitored_poll_time_unit", r.MonitoredPollTimeUnit)
	d.Set("interface_hold_time", r.InterfaceHoldTime)

	if r.LanFoInterface != nil {
		d.Set("lan_interface_hw_id", hwIdFromObjId(r.LanFoInterface.ObjectId))
		d.Set("lan_failover_name", r.LanIFCName)
		d.Set("lan_primary_ip", r.LanActiveIP.Value)
		d.Set("lan_secondary_ip", r.LanStandby.Value)
		d.Set("lan_net_mask", r.LanSubnet.Value)
	}

	if r.IsLANInterfacePreferredPrimary == true && r.IsLANInterfacePreferredSecondary == false {
		d.Set("lan_preferred_role", "primary")
	} else if r.IsLANInterfacePreferredPrimary == false && r.IsLANInterfacePreferredSecondary == true {
		d.Set("lan_preferred_role", "secondary")
	}

	if r.StateFoInterface != nil {
		if r.StateFoInterface.Name != "" {
			d.Set("state_interface_hw_id", r.StateFoInterface.Name)
		} else {
			d.Set("state_interface_hw_id", hwIdFromObjId(r.StateFoInterface.ObjectId))
		}
		d.Set("state_failover_name", r.LanIFCName)
		d.Set("state_primary_ip", r.LanActiveIP.Value)
		d.Set("state_secondary_ip", r.LanStandby.Value)
		d.Set("state_net_mask", r.LanSubnet.Value)
	}

	return nil
}

func resourceCiscoASAFailoverSetupDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	failoverSetupFlush := &ciscoasa.FailoverSetup{
		EnableFOCheck:                      false,
		SecretKey:                          "",
		IpSecKey:                           "",
		HexKey:                             false,
		LanIFCName:                         "",
		LanFoInterface:                     nil,
		LanActiveIP:                        nil,
		LanStandby:                         nil,
		LanSubnet:                          nil,
		IsLANInterfacePreferredPrimary:     false,
		IsLANInterfacePreferredSecondary:   false,
		StateIFCName:                       "",
		StateFoInterface:                   nil,
		StateActiveIP:                      nil,
		StateStandbyIP:                     nil,
		StateSubnet:                        nil,
		HttpReplicate:                      false,
		ReplicateRate:                      -1,
		FailedInterfacesUnit:               "Number",
		FailedInterfacesTriggeringFailover: "1",
		UnitPollTime:                       "1",
		UnitPollTimeUnit:                   "seconds",
		UnitHoldTime:                       "15",
		UnitHoldTimeUnit:                   "seconds",
		MonitoredPollTime:                  "5",
		MonitoredPollTimeUnit:              "seconds",
		InterfaceHoldTime:                  "25",
	}

	err := ca.Failover.UpdateFailoverSetup(failoverSetupFlush)
	if err != nil {
		// Workaround for disabling FO first
		if strings.Contains(err.Error(), "Interface is in use by failover.") {
			err = ca.Failover.UpdateFailoverSetup(failoverSetupFlush)
			if err != nil {
				return fmt.Errorf(
					"Error deleting Failover Setup: %v", err)
			}
		}
		return fmt.Errorf(
			"Error deleting Failover Setup: %v", err)
	}

	return nil
}

func makeInterfaceRefObj(iface string) *ciscoasa.InterfaceRef {
	i := &ciscoasa.InterfaceRef{
		Kind: "objectRef#Interface",
	}
	if isInterfaceHwId(iface) {
		i.ObjectId = objectIdFromHwId(iface)
	} else {
		i.Name = iface
	}
	return i
}

func makeIpAdressObj(value string) *ciscoasa.Address {
	return &ciscoasa.Address{
		Kind:  "IPv4Address",
		Value: value,
	}
}

func makeNetmaskObj(value string) *ciscoasa.Address {
	return &ciscoasa.Address{
		Kind:  "IPv4NetMask",
		Value: value,
	}
}
