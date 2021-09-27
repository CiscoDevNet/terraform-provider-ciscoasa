package ciscoasa

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func dataSourceCiscoASAPhysicalInterfaces() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceCiscoASAPhysicalInterfacesRead,

		Schema: map[string]*schema.Schema{
			"interfaces_physical": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"active_mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"forward_traffic_cx": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"forward_traffic_sfr": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"hardware_id": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"interface_desc": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"ip_address": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"static": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"ip": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"net_mask": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},

									"dhcp": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"dhcp_option_using_mac": {
													Type:     schema.TypeBool,
													Computed: true,
												},

												"dhcp_broadcast": {
													Type:     schema.TypeBool,
													Computed: true,
												},

												"dhcp_client": {
													Type:     schema.TypeList,
													Computed: true,
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"set_default_route": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"metric": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"primary_track_id": {
																Type:     schema.TypeInt,
																Computed: true,
															},
															"tracking_enabled": {
																Type:     schema.TypeBool,
																Computed: true,
															},
															"sla_tracking_settings": {
																Type:     schema.TypeList,
																Computed: true,
																Elem: &schema.Resource{
																	Schema: map[string]*schema.Schema{
																		"sla_id": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"tracked_ip": {
																			Type:     schema.TypeString,
																			Computed: true,
																		},
																		"frequency_in_seconds": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"data_size_in_bytes": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"threshold_in_milliseconds": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"tos": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"timeout_in_milliseconds": {
																			Type:     schema.TypeInt,
																			Computed: true,
																		},
																		"num_packets": {
																			Type:     schema.TypeInt,
																			Computed: true,
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
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"auto_config": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"dad_attempts": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"enabled": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"enforce_eui64": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"link_local_address": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"standby": {
													Type:     schema.TypeString,
													Computed: true,
												},
											},
										},
									},
									"ipv6_addresses": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"address": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"standby": {
													Type:     schema.TypeString,
													Computed: true,
												},
												"prefix_length": {
													Type:     schema.TypeInt,
													Computed: true,
												},
											},
										},
									},

									"managed_address_config": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"n_discovery_prefix_list": {
										Type:     schema.TypeList,
										Computed: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"off_link": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"no_advertise": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"preferred_lifetime": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"valid_lifetime": {
													Type:     schema.TypeInt,
													Computed: true,
												},
												"has_duration": {
													Type:     schema.TypeBool,
													Computed: true,
												},
												"default_prefix": {
													Type:     schema.TypeBool,
													Computed: true,
												},
											},
										},
									},
									"ns_interval": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"other_stateful_config": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"reachable_time": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"router_advert_interval": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"router_advert_interval_unit": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"router_advert_lifetime": {
										Type:     schema.TypeInt,
										Computed: true,
									},
									"suppress_router_advert": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},

						"management_only": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"mtu": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"security_level": {
							Type:     schema.TypeInt,
							Computed: true,
						},

						"shutdown": {
							Type:     schema.TypeBool,
							Computed: true,
						},

						"stand_by_mac_address": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCiscoASAPhysicalInterfacesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ca := meta.(*ciscoasa.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	l, err := ca.Interfaces.ListPhysicalInterfaces()
	if err != nil {
		return diag.FromErr(err)
	}

	// Create an empty list to hold all interfaces
	interfaces := make([]interface{}, len(l.Items), len(l.Items))

	for i, r := range l.Items {
		interfaceMap := make(map[string]interface{})
		interfaceMap["id"] = r.ObjectID
		interfaceMap["active_mac_address"] = r.ActiveMacAddress
		interfaceMap["forward_traffic_cx"] = r.ForwardTrafficCX
		interfaceMap["forward_traffic_sfr"] = r.ForwardTrafficSFR
		interfaceMap["hardware_id"] = r.HardwareID
		interfaceMap["interface_desc"] = r.InterfaceDesc
		interfaceMap["management_only"] = r.ManagementOnly
		interfaceMap["mtu"] = r.Mtu
		interfaceMap["name"] = r.Name
		if r.IPAddress.Kind != "" {
			interfaceMap["ip_address"] = flattenIPAddress(r.IPAddress)
		}
		interfaceMap["ipv6_info"] = flattenIPv6Info(r.Ipv6Info)
		interfaceMap["security_level"] = r.SecurityLevel
		interfaceMap["shutdown"] = r.Shutdown
		interfaceMap["stand_by_mac_address"] = r.StandByMacAddress

		interfaces[i] = interfaceMap
	}

	if err := d.Set("interfaces_physical", interfaces); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
