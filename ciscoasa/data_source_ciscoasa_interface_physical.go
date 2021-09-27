package ciscoasa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func dataSourceCiscoASAPhysicalInterface() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceCiscoASAPhysicalInterfaceRead,

		Schema: map[string]*schema.Schema{
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
				Required: true,
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
	}
}

func dataSourceCiscoASAPhysicalInterfaceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ca := meta.(*ciscoasa.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	objectID := objectIdFromHwId(d.Get("hardware_id").(string))

	r, err := ca.Interfaces.GetPhysicalInterface(objectID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("active_mac_address", r.ActiveMacAddress); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("forward_traffic_cx", r.ForwardTrafficCX); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("forward_traffic_sfr", r.ForwardTrafficSFR); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("hardware_id", r.HardwareID); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("interface_desc", r.InterfaceDesc); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("management_only", r.ManagementOnly); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("mtu", r.Mtu); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("name", r.Name); err != nil {
		return diag.FromErr(err)
	}
	if r.IPAddress.Kind == "" {
		d.Set("ip_address", nil)
	} else {
		if err := d.Set("ip_address", flattenIPAddress(r.IPAddress)); err != nil {
			return diag.FromErr(err)
		}
	}
	if err := d.Set("ipv6_info", flattenIPv6Info(r.Ipv6Info)); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("security_level", r.SecurityLevel); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("shutdown", r.Shutdown); err != nil {
		return diag.FromErr(err)
	}
	if err := d.Set("stand_by_mac_address", r.StandByMacAddress); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(objectID)

	return diags
}
