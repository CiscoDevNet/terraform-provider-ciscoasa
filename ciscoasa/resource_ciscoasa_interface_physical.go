package ciscoasa

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func resourceCiscoASAPhysicalInterface() *schema.Resource {

	return &schema.Resource{
		Create: resourceCiscoASAPhysicalInterfaceCreate,
		Read:   resourceCiscoASAPhysicalInterfaceRead,
		Update: resourceCiscoASAPhysicalInterfaceUpdate,
		Delete: resourceCiscoASAPhysicalInterfaceDelete,

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

			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCiscoASAPhysicalInterfaceCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	var err error

	objectID := strings.Replace(d.Get("hardware_id").(string), "/", "_API_SLASH_", 1)

	hardwareID := d.Get("hardware_id").(string)

	kind, err := physicalKindFromHwId(hardwareID)
	if err != nil {
		return err
	}

	physicalInterface := ciscoasa.PhysicalInterface{
		ActiveMacAddress:  d.Get("active_mac_address").(string),
		ForwardTrafficCX:  d.Get("forward_traffic_cx").(bool),
		ForwardTrafficSFR: d.Get("forward_traffic_sfr").(bool),
		HardwareID:        hardwareID,
		InterfaceDesc:     d.Get("interface_desc").(string),
		Kind:              kind,
		ManagementOnly:    d.Get("management_only").(bool),
		Mtu:               d.Get("mtu").(int),
		Name:              d.Get("name").(string),
		ObjectID:          objectID,
		SecurityLevel:     d.Get("security_level").(int),
		Shutdown:          d.Get("shutdown").(bool),
		StandByMacAddress: d.Get("stand_by_mac_address").(string),
	}

	if _, ok := d.GetOk("ip_address"); ok {
		physicalInterface.IPAddress, err = expandIPAddress(d)
		if err != nil {
			return err
		}
	}

	if _, ok := d.GetOk("ipv6_info"); ok {
		physicalInterface.Ipv6Info, err = expandIPv6Info(d)
		if err != nil {
			return err
		}
	}

	physicalInterfaceId, err := ca.Interfaces.UpdatePhysicalInterface(
		physicalInterface.ActiveMacAddress,
		physicalInterface.ForwardTrafficCX,
		physicalInterface.ForwardTrafficSFR,
		physicalInterface.HardwareID,
		physicalInterface.InterfaceDesc,
		physicalInterface.IPAddress,
		physicalInterface.Ipv6Info,
		physicalInterface.Kind,
		physicalInterface.ManagementOnly,
		physicalInterface.Mtu,
		physicalInterface.Name,
		physicalInterface.ObjectID,
		physicalInterface.SecurityLevel,
		physicalInterface.Shutdown,
		physicalInterface.StandByMacAddress,
	)

	if err != nil {
		return fmt.Errorf(
			"Error creating Physical Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
	}

	d.SetId(physicalInterfaceId)

	return resourceCiscoASAPhysicalInterfaceRead(d, meta)
}

func resourceCiscoASAPhysicalInterfaceRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.Interfaces.GetPhysicalInterface(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf(
				"[DEBUG] Physical Interface for %s->%s not found", d.Get("name").(string), d.Get("hardware_id").(string))
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Error reading Physical Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
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

	return nil
}

func resourceCiscoASAPhysicalInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	var err error

	physicalInterfaceId := d.Id()

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
	) {
		physicalInterface := ciscoasa.PhysicalInterface{
			ActiveMacAddress:  d.Get("active_mac_address").(string),
			ForwardTrafficCX:  d.Get("forward_traffic_cx").(bool),
			ForwardTrafficSFR: d.Get("forward_traffic_sfr").(bool),
			HardwareID:        d.Get("hardware_id").(string),
			InterfaceDesc:     d.Get("interface_desc").(string),
			ManagementOnly:    d.Get("management_only").(bool),
			Mtu:               d.Get("mtu").(int),
			Name:              d.Get("name").(string),
			SecurityLevel:     d.Get("security_level").(int),
			Shutdown:          d.Get("shutdown").(bool),
			StandByMacAddress: d.Get("stand_by_mac_address").(string),
		}

		if _, ok := d.GetOk("ip_address"); ok {
			physicalInterface.IPAddress, err = expandIPAddress(d)
			if err != nil {
				return err
			}
		}

		if _, ok := d.GetOk("ipv6_info"); ok {
			physicalInterface.Ipv6Info, err = expandIPv6Info(d)
			if err != nil {
				return err
			}
		}

		kind, _ := physicalKindFromHwId(physicalInterfaceId)

		_, err := ca.Interfaces.UpdatePhysicalInterface(
			physicalInterface.ActiveMacAddress,
			physicalInterface.ForwardTrafficCX,
			physicalInterface.ForwardTrafficSFR,
			physicalInterface.HardwareID,
			physicalInterface.InterfaceDesc,
			physicalInterface.IPAddress,
			physicalInterface.Ipv6Info,
			kind,
			physicalInterface.ManagementOnly,
			physicalInterface.Mtu,
			physicalInterface.Name,
			physicalInterfaceId,
			physicalInterface.SecurityLevel,
			physicalInterface.Shutdown,
			physicalInterface.StandByMacAddress,
		)
		if err != nil {
			return fmt.Errorf(
				"Error updating Physical Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCiscoASAPhysicalInterfaceRead(d, meta)
}

func resourceCiscoASAPhysicalInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	physicalInterfaceId := d.Id()

	kind, _ := physicalKindFromHwId(physicalInterfaceId)

	err := ca.Interfaces.DeletePhysicalInterface(
		d.Get("hardware_id").(string),
		kind,
		physicalInterfaceId,
	)

	if err != nil {
		return fmt.Errorf(
			"Error deleting Physical Interface %s->%s: %v", d.Get("name").(string), d.Get("hardware_id").(string), err)
	}

	return nil
}

// Get kind from simple parsing
func physicalKindFromHwId(hardwareID string) (string, error) {
	if hardwareID == "" {
		return "", errors.New("a 'hardwareID' cannot be an empty string")
	}
	re := regexp.MustCompile(`(Gigabit|Management|TenGigabit)`)
	match := re.FindStringSubmatch(hardwareID)

	if match != nil {
		if match[0] == "Gigabit" {
			return "object#GigabitInterface", nil
		} else if match[0] == "TenGigabit" {
			return "object#TenGigInterface", nil
		} else if match[0] == "Management" {
			return "object#MgmtInterface", nil
		}
	}

	return "", fmt.Errorf("failed to get kind from 'hardwareID': %q", hardwareID)
}
