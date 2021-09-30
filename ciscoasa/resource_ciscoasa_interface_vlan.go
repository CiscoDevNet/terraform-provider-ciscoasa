package ciscoasa

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCiscoASAVlanInterface() *schema.Resource {

	return &schema.Resource{
		Create: resourceCiscoASAVlanInterfaceCreate,
		Read:   resourceCiscoASAVlanInterfaceRead,
		Update: resourceCiscoASAVlanInterfaceUpdate,
		Delete: resourceCiscoASAVlanInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"active_mac_address": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: false,
			},

			"forward_traffic_cx": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},

			"forward_traffic_sfr": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},

			"hardware_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"interface_desc": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: false,
			},

			"ip_address": {
				Type:     schema.TypeSet,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"static": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"ip": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: false,
									},
									"net_mask": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: false,
									},
								},
							},
						},

						"dhcp": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"dhcp_option_using_mac": {
										Type:     schema.TypeBool,
										Required: true,
										ForceNew: false,
									},

									"dhcp_broadcast": {
										Type:     schema.TypeBool,
										Required: true,
										ForceNew: false,
									},

									"dhcp_client": {
										Type:       schema.TypeList,
										Required:   true,
										ConfigMode: schema.SchemaConfigModeAttr,
										MaxItems:   1,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"set_default_route": {
													Type:     schema.TypeBool,
													Required: true,
													ForceNew: false,
												},
												"metric": {
													Type:     schema.TypeInt,
													Required: true,
													ForceNew: false,
												},
												"primary_track_id": {
													Type:     schema.TypeInt,
													Required: true,
													ForceNew: false,
												},
												"tracking_enabled": {
													Type:     schema.TypeBool,
													Required: true,
													ForceNew: false,
												},
												"sla_tracking_settings": {
													Type:       schema.TypeList,
													Computed:   true,
													Optional:   true,
													ConfigMode: schema.SchemaConfigModeAttr,
													MaxItems:   1,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"sla_id": {
																Type:     schema.TypeInt,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
															"tracked_ip": {
																Type:     schema.TypeString,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
															"frequency_in_seconds": {
																Type:     schema.TypeInt,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
															"data_size_in_bytes": {
																Type:     schema.TypeInt,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
															"threshold_in_milliseconds": {
																Type:     schema.TypeInt,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
															"tos": {
																Type:     schema.TypeInt,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
															"timeout_in_milliseconds": {
																Type:     schema.TypeInt,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
															"num_packets": {
																Type:     schema.TypeInt,
																Computed: true,
																Optional: true,
																ForceNew: false,
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},

			"ipv6_info": {
				Type:     schema.TypeSet,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auto_config": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: false,
						},
						"dad_attempts": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1,
							ForceNew: false,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: false,
						},
						"enforce_eui64": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: false,
						},
						"link_local_address": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// "is_eui64": {
									// 	Type:     schema.TypeBool,
									// 	Required: true,
									// 	ForceNew: false,
									// },
									"address": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: false,
									},
									"standby": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: false,
									},
								},
							},
						},
						"ipv6_addresses": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									// "is_eui64": {
									// 	Type:     schema.TypeBool,
									// 	Required: true,
									// 	ForceNew: false,
									// },
									"address": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: false,
									},
									"standby": {
										Type:     schema.TypeString,
										Required: true,
										ForceNew: false,
									},
									"prefix_length": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: false,
									},
								},
							},
						},
						"managed_address_config": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: false,
						},
						"n_discovery_prefix_list": {
							Type:       schema.TypeList,
							Optional:   true,
							ConfigMode: schema.SchemaConfigModeAttr,
							MaxItems:   1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"off_link": {
										Type:     schema.TypeBool,
										Required: true,
										ForceNew: false,
									},
									"no_advertise": {
										Type:     schema.TypeBool,
										Required: true,
										ForceNew: false,
									},
									"preferred_lifetime": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: false,
									},
									"valid_lifetime": {
										Type:     schema.TypeInt,
										Required: true,
										ForceNew: false,
									},
									"has_duration": {
										Type:     schema.TypeBool,
										Required: true,
										ForceNew: false,
									},
									"default_prefix": {
										Type:     schema.TypeBool,
										Required: true,
										ForceNew: false,
									},
								},
							},
						},
						"ns_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1000,
							ForceNew: false,
						},
						"other_stateful_config": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: false,
						},
						"reachable_time": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  0,
							ForceNew: false,
						},
						"router_advert_interval": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  200,
							ForceNew: false,
						},
						"router_advert_interval_unit": {
							Type:     schema.TypeString,
							Optional: true,
							Default:  "sec",
							ValidateFunc: validation.StringInSlice([]string{
								"sec", "msec",
							}, false),
							ForceNew: false,
						},
						"router_advert_lifetime": {
							Type:     schema.TypeInt,
							Optional: true,
							Default:  1800,
							ForceNew: false,
						},
						"suppress_router_advert": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  false,
							ForceNew: false,
						},
					},
				},
			},

			"management_only": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: false,
			},

			"mtu": {
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1500,
				ForceNew: false,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: false,
			},

			"security_level": {
				Type:     schema.TypeInt,
				Required: true,
				ForceNew: false,
			},

			"shutdown": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				ForceNew: false,
			},

			"stand_by_mac_address": {
				Type:     schema.TypeString,
				Default:  "",
				Optional: true,
				ForceNew: false,
			},

			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
				ForceNew: false,
			},

			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCiscoASAVlanInterfaceCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	var err error

	hardwareID := d.Get("hardware_id").(string)

	kind, err := vlanKindFromHwId(hardwareID)
	if err != nil {
		return err
	}

	vlanId, ok := d.GetOk("vlan_id")
	if !ok {
		vlanId, err = vlanIdFromHwId(d)
		if err != nil {
			return err
		}
	}

	VlanInterface := ciscoasa.VlanInterface{
		ActiveMacAddress:  d.Get("active_mac_address").(string),
		ForwardTrafficCX:  d.Get("forward_traffic_cx").(bool),
		ForwardTrafficSFR: d.Get("forward_traffic_sfr").(bool),
		HardwareID:        hardwareID,
		InterfaceDesc:     d.Get("interface_desc").(string),
		Kind:              kind,
		ManagementOnly:    d.Get("management_only").(bool),
		Mtu:               d.Get("mtu").(int),
		Name:              d.Get("name").(string),
		SecurityLevel:     d.Get("security_level").(int),
		Shutdown:          d.Get("shutdown").(bool),
		StandByMacAddress: d.Get("stand_by_mac_address").(string),
		VlanID:            vlanId.(int),
	}

	if _, ok := d.GetOk("ip_address"); ok {
		VlanInterface.IPAddress, err = expandIPAddress(d)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("ipv6_info"); ok {
		VlanInterface.Ipv6Info, err = expandIPv6Info(d)
		if err != nil {
			return err
		}
	}

	VlanInterfaceId, err := ca.Interfaces.CreateVlanInterface(

		VlanInterface.ActiveMacAddress,
		VlanInterface.ForwardTrafficCX,
		VlanInterface.ForwardTrafficSFR,
		VlanInterface.HardwareID,
		VlanInterface.InterfaceDesc,
		VlanInterface.IPAddress,
		VlanInterface.Ipv6Info,
		VlanInterface.Kind,
		VlanInterface.ManagementOnly,
		VlanInterface.Mtu,
		VlanInterface.Name,
		VlanInterface.SecurityLevel,
		VlanInterface.Shutdown,
		VlanInterface.StandByMacAddress,
		VlanInterface.VlanID,
	)

	if err != nil {
		return fmt.Errorf(
			"Error creating Vlan Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
	}

	d.SetId(VlanInterfaceId)

	return resourceCiscoASAVlanInterfaceRead(d, meta)
}

func resourceCiscoASAVlanInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.Interfaces.GetVlanInterface(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf(
				"[DEBUG] Vlan Interface for %s->%s no longer exists", d.Get("name").(string), d.Get("hardware_id").(string))
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Error reading Vlan Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
	}

	d.Set("active_mac_address", r.ActiveMacAddress)
	d.Set("forward_traffic_cx", r.ForwardTrafficCX)
	d.Set("forward_traffic_sfr", r.ForwardTrafficSFR)
	d.Set("hardware_id", r.HardwareID)
	d.Set("interface_desc", r.InterfaceDesc)
	d.Set("management_only", r.ManagementOnly)
	d.Set("mtu", r.Mtu)
	d.Set("name", r.Name)
	if r.IPAddress.Kind == "" {
		d.Set("ip_address", nil)
	} else {
		d.Set("ip_address", flattenIPAddress(r.IPAddress))
	}
	d.Set("ipv6_info", flattenIPv6Info(r.Ipv6Info))
	d.Set("security_level", r.SecurityLevel)
	d.Set("shutdown", r.Shutdown)
	d.Set("stand_by_mac_address", r.StandByMacAddress)
	d.Set("vlan_id", r.VlanID)

	return nil
}

func resourceCiscoASAVlanInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	var err error

	VlanInterfaceId := d.Id()

	if d.HasChanges(
		"active_mac_address",
		"forward_traffic_cx",
		"forward_traffic_sfr",
		"interface_desc",
		"ip_address",
		"ipv6_info",
		"management_only",
		"mtu",
		"name",
		"security_level",
		"shutdown",
		"stand_by_mac_address",
		"vlan_id",
	) {
		vlanId, ok := d.GetOk("vlan_id")
		if !ok {
			vlanId, err = vlanIdFromHwId(d)
			if err != nil {
				return err
			}
		}

		hardwareID := d.Get("hardware_id").(string)

		kind, _ := vlanKindFromHwId(hardwareID)

		VlanInterface := ciscoasa.VlanInterface{
			ActiveMacAddress:  d.Get("active_mac_address").(string),
			ForwardTrafficCX:  d.Get("forward_traffic_cx").(bool),
			ForwardTrafficSFR: d.Get("forward_traffic_sfr").(bool),
			HardwareID:        hardwareID,
			InterfaceDesc:     d.Get("interface_desc").(string),
			Kind:              kind,
			ManagementOnly:    d.Get("management_only").(bool),
			Mtu:               d.Get("mtu").(int),
			Name:              d.Get("name").(string),
			SecurityLevel:     d.Get("security_level").(int),
			Shutdown:          d.Get("shutdown").(bool),
			StandByMacAddress: d.Get("stand_by_mac_address").(string),
			VlanID:            vlanId.(int),
		}

		if _, ok := d.GetOk("ip_address"); ok {
			VlanInterface.IPAddress, err = expandIPAddress(d)
			if err != nil {
				return err
			}
		}

		if _, ok := d.GetOk("ipv6_info"); ok {
			VlanInterface.Ipv6Info, err = expandIPv6Info(d)
			if err != nil {
				return err
			}
		}

		err := ca.Interfaces.UpdateVlanInterface(
			VlanInterface.ActiveMacAddress,
			VlanInterface.ForwardTrafficCX,
			VlanInterface.ForwardTrafficSFR,
			VlanInterface.HardwareID,
			VlanInterface.InterfaceDesc,
			VlanInterface.IPAddress,
			VlanInterface.Ipv6Info,
			VlanInterface.Kind,
			VlanInterface.ManagementOnly,
			VlanInterface.Mtu,
			VlanInterface.Name,
			VlanInterfaceId,
			VlanInterface.SecurityLevel,
			VlanInterface.Shutdown,
			VlanInterface.StandByMacAddress,
			VlanInterface.VlanID,
		)
		if err != nil {
			return fmt.Errorf(
				"Error updating Vlan Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCiscoASAVlanInterfaceRead(d, meta)
}

func resourceCiscoASAVlanInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Interfaces.DeleteVlanInterface(d.Id())

	if err != nil {
		return fmt.Errorf(
			"Error deleting Vlan Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
	}

	return nil
}

// Get kind from simple parsing
func vlanKindFromHwId(hardwareID string) (string, error) {
	if hardwareID == "" {
		return "", errors.New("a 'hardwareID' cannot be an empty string")
	}
	re := regexp.MustCompile(`(Gigabit|Management|TenGigabit|Port-channel|Redundant)`)
	match := re.FindStringSubmatch(hardwareID)

	if match != nil {
		if match[0] == "Gigabit" {
			return "object#VlanGigInterface", nil
		} else if match[0] == "TenGigabit" {
			return "object#VlanTenGigInterface", nil
		} else if match[0] == "Management" {
			return "object#VlanMgmtInterface", nil
		} else if match[0] == "Port-channel" {
			return "object#VlanPCInterface", nil
		} else if match[0] == "Redundant" {
			return "object#VlanRedInterface", nil
		}
	}

	return "", fmt.Errorf("failed to get kind from 'hardwareID': %q", hardwareID)
}

// Get VlanId if not provided in parameters
func vlanIdFromHwId(d *schema.ResourceData) (int, error) {
	if hardwareID, ok := d.GetOk("hardware_id"); ok {
		s := strings.Split(hardwareID.(string), ".")
		var i int
		if _, err := fmt.Sscanf(s[len(s)-1], "%d", &i); err == nil {
			return i, nil
		}
	}
	return 0, errors.New("Cannot get a 'hardwareID'")
}
