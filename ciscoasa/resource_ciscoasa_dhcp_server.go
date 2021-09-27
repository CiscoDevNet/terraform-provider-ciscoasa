package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func resourceCiscoASADhcpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASADhcpServerUpdate,
		Read:   resourceCiscoASADhcpServerRead,
		Update: resourceCiscoASADhcpServerUpdate,
		Delete: resourceCiscoASADhcpServerDelete,

		Schema: map[string]*schema.Schema{
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},

			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"pool_start_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
			},

			"pool_end_ip": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
			},

			"dns_ip_primary": {
				Type:         schema.TypeString,
				Optional:     true,
				Default:      "",
				ValidateFunc: validation.IsIPAddress,
			},

			"dns_ip_secondary": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
			},

			"wins_ip_primary": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
			},

			"wins_ip_secondary": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.IsIPAddress,
			},

			"lease_length": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"ping_timeout": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"domain_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"auto_config_enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"auto_config_interface": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"vpn_override": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"options": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"type": {
							Type:     schema.TypeString,
							Required: true,
							ValidateFunc: validation.StringInSlice(
								[]string{"ascii", "hex", "ip"},
								false,
							),
						},
						"code": {
							Type:         schema.TypeInt,
							Required:     true,
							ValidateFunc: validation.IntBetween(1, 255),
						},
						"value1": {
							Type:     schema.TypeString,
							Required: true,
						},
						"value2": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"ddns_update_dns_client": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"ddns_update_both_records": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"ddns_override_client_settings": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

// There is no real creation of the DHCP server.
// All the changes are made on already existing object,
// which is created with the Interface.
func resourceCiscoASADhcpServerUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	options := d.Get("options").([]interface{})
	dhcpOptions := []*ciscoasa.DhcpServerOptions{}

	for _, option := range options {
		o := option.(map[string]interface{})
		dhcpOption := &ciscoasa.DhcpServerOptions{
			Type:   o["type"].(string),
			Code:   o["code"].(int),
			Value1: o["value1"].(string),
			Value2: o["value2"].(string),
		}
		dhcpOptions = append(dhcpOptions, dhcpOption)
	}

	iface := d.Get("interface").(string)
	dhcpInterface := &ciscoasa.InterfaceRef{
		Kind: "objectRef#Interface",
		Name: iface,
	}

	DhcpServer := ciscoasa.DhcpServer{
		Interface:             dhcpInterface,
		Enabled:               d.Get("enabled").(bool),
		PoolStartIP:           d.Get("pool_start_ip").(string),
		PoolEndIP:             d.Get("pool_end_ip").(string),
		DnsIP1:                d.Get("dns_ip_primary").(string),
		DnsIP2:                d.Get("dns_ip_secondary").(string),
		WinsIP1:               d.Get("wins_ip_primary").(string),
		WinsIP2:               d.Get("wins_ip_secondary").(string),
		LeaseLengthInSec:      d.Get("lease_length").(string),
		PingTimeoutInMilliSec: d.Get("ping_timeout").(string),
		DomainName:            d.Get("domain_name").(string),
		IsAutoConfigEnabled:   d.Get("auto_config_enabled").(bool),
		AutoConfigInterface:   d.Get("auto_config_interface").(string),
		IsVpnOverride:         d.Get("vpn_override").(bool),
		Options:               dhcpOptions,
		Kind:                  "object#DhcpServer",
		ObjectId:              iface,
	}

	DhcpServer.Ddns.UpdateDNSClient = d.Get("ddns_update_dns_client").(bool)
	DhcpServer.Ddns.UpdateBothRecords = d.Get("ddns_update_both_records").(bool)
	DhcpServer.Ddns.OverrideClientSettings = d.Get("ddns_override_client_settings").(bool)

	DhcpServerId, err := ca.Dhcp.UpdateDhcpServer(&DhcpServer)

	if err != nil {
		return fmt.Errorf(
			"Error creating DHCP Server %s: %v", iface, err)
	}

	d.SetId(DhcpServerId)

	return resourceCiscoASADhcpServerRead(d, meta)
}

func resourceCiscoASADhcpServerRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.Dhcp.GetDhcpServer(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf(
				"[DEBUG] DHCP Server for %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Error reading DHCP Server %s: %v", d.Id(), err)
	}

	d.Set("interface", r.Interface.Name)
	d.Set("enabled", r.Enabled)
	d.Set("pool_start_ip", r.PoolStartIP)
	d.Set("pool_end_ip", r.PoolEndIP)
	d.Set("dns_ip_primary", r.DnsIP1)
	d.Set("dns_ip_secondary", r.DnsIP2)
	d.Set("wins_ip_primary", r.WinsIP1)
	d.Set("wins_ip_secondary", r.WinsIP2)
	d.Set("lease_length", r.LeaseLengthInSec)
	d.Set("ping_timeout", r.PingTimeoutInMilliSec)
	d.Set("domain_name", r.DomainName)
	d.Set("auto_config_enabled", r.IsAutoConfigEnabled)
	d.Set("auto_config_interface", r.AutoConfigInterface)
	d.Set("vpn_override", r.IsVpnOverride)
	d.Set("options", flattenDhcpOptions(r.Options))
	d.Set("ddns_update_dns_client", r.Ddns.UpdateDNSClient)
	d.Set("ddns_update_both_records", r.Ddns.UpdateBothRecords)
	d.Set("ddns_override_client_settings", r.Ddns.OverrideClientSettings)

	return nil
}

func resourceCiscoASADhcpServerDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	dhcpOptions := []*ciscoasa.DhcpServerOptions{}

	iface := d.Get("interface").(string)
	dhcpInterface := &ciscoasa.InterfaceRef{
		Kind: "objectRef#Interface",
		Name: iface,
	}

	DhcpServer := ciscoasa.DhcpServer{
		Interface:             dhcpInterface,
		Enabled:               resourceCiscoASADhcpServer().Schema["enabled"].ZeroValue().(bool),
		PoolStartIP:           resourceCiscoASADhcpServer().Schema["pool_start_ip"].ZeroValue().(string),
		PoolEndIP:             resourceCiscoASADhcpServer().Schema["pool_end_ip"].ZeroValue().(string),
		DnsIP1:                resourceCiscoASADhcpServer().Schema["dns_ip_primary"].ZeroValue().(string),
		DnsIP2:                resourceCiscoASADhcpServer().Schema["dns_ip_secondary"].ZeroValue().(string),
		WinsIP1:               resourceCiscoASADhcpServer().Schema["wins_ip_primary"].ZeroValue().(string),
		WinsIP2:               resourceCiscoASADhcpServer().Schema["wins_ip_secondary"].ZeroValue().(string),
		LeaseLengthInSec:      resourceCiscoASADhcpServer().Schema["lease_length"].ZeroValue().(string),
		PingTimeoutInMilliSec: resourceCiscoASADhcpServer().Schema["ping_timeout"].ZeroValue().(string),
		DomainName:            resourceCiscoASADhcpServer().Schema["domain_name"].ZeroValue().(string),
		IsAutoConfigEnabled:   resourceCiscoASADhcpServer().Schema["auto_config_enabled"].ZeroValue().(bool),
		AutoConfigInterface:   resourceCiscoASADhcpServer().Schema["auto_config_interface"].ZeroValue().(string),
		IsVpnOverride:         resourceCiscoASADhcpServer().Schema["vpn_override"].ZeroValue().(bool),
		Options:               dhcpOptions,
		Kind:                  "object#DhcpServer",
		ObjectId:              iface,
	}

	DhcpServer.Ddns.UpdateDNSClient = resourceCiscoASADhcpServer().Schema["ddns_update_dns_client"].ZeroValue().(bool)
	DhcpServer.Ddns.UpdateBothRecords = resourceCiscoASADhcpServer().Schema["ddns_update_both_records"].ZeroValue().(bool)
	DhcpServer.Ddns.OverrideClientSettings = resourceCiscoASADhcpServer().Schema["ddns_override_client_settings"].ZeroValue().(bool)

	_, err := ca.Dhcp.UpdateDhcpServer(&DhcpServer)
	if err != nil {
		return fmt.Errorf(
			"Error deleting DHCP Server %s: %v", iface, err)
	}

	return nil
}

func flattenDhcpOptions(dhcpOptions []*ciscoasa.DhcpServerOptions) []interface{} {
	if dhcpOptions != nil {
		dos := make([]interface{}, len(dhcpOptions), len(dhcpOptions))

		for i, dhcpOption := range dhcpOptions {
			do := make(map[string]interface{})

			do["type"] = dhcpOption.Type
			do["code"] = dhcpOption.Code
			do["value1"] = dhcpOption.Value1
			do["value2"] = dhcpOption.Value2

			dos[i] = do
		}

		return dos
	}

	return make([]interface{}, 0)
}
