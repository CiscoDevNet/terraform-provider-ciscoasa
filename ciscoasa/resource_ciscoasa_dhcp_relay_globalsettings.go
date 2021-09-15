package ciscoasa

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASADhcpRelayGlobalsettings() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASADhcpRelayGlobalsettingsUpdate,
		Read:   resourceCiscoASADhcpRelayGlobalsettingsRead,
		Update: resourceCiscoASADhcpRelayGlobalsettingsUpdate,
		Delete: resourceCiscoASADhcpRelayGlobalsettingsDelete,

		Schema: map[string]*schema.Schema{
			"ipv4_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      60,
				ValidateFunc: validation.IntBetween(1, 3600),
			},

			"ipv6_timeout": {
				Type:         schema.TypeInt,
				Optional:     true,
				Default:      60,
				ValidateFunc: validation.IntBetween(1, 3600),
			},

			"trusted_on_all_interfaces": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceCiscoASADhcpRelayGlobalsettingsUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Dhcp.UpdateDhcpRelayGlobalsettings(
		&ciscoasa.DhcpRelayGS{
			Ipv4Timeout:            d.Get("ipv4_timeout").(int),
			Ipv6Timeout:            d.Get("ipv6_timeout").(int),
			TrustedOnAllInterfaces: d.Get("trusted_on_all_interfaces").(bool),
			Kind:                   "object#DHCPRelayGlobalSettings",
		},
	)

	if err != nil {
		return fmt.Errorf(
			"Error creating DHCP Relay Global Settings: %v", err)
	}

	d.SetId("globalsettings")

	return resourceCiscoASADhcpRelayGlobalsettingsRead(d, meta)
}

func resourceCiscoASADhcpRelayGlobalsettingsRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.Dhcp.GetDhcpRelayGlobalsettings()
	if err != nil {
		return fmt.Errorf(
			"Error reading DHCP Relay Global Settings: %v", err)
	}

	d.Set("ipv4_timeout", r.Ipv4Timeout)
	d.Set("ipv6_timeout", r.Ipv6Timeout)
	d.Set("trusted_on_all_interfaces", r.TrustedOnAllInterfaces)

	return nil
}

func resourceCiscoASADhcpRelayGlobalsettingsDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	ipv4Timeout, _ := resourceCiscoASADhcpRelayGlobalsettings().Schema["ipv4_timeout"].DefaultValue()
	ipv6Timeout, _ := resourceCiscoASADhcpRelayGlobalsettings().Schema["ipv6_timeout"].DefaultValue()
	truested, _ := resourceCiscoASADhcpRelayGlobalsettings().Schema["trusted_on_all_interfaces"].DefaultValue()

	err := ca.Dhcp.UpdateDhcpRelayGlobalsettings(
		&ciscoasa.DhcpRelayGS{
			Ipv4Timeout:            ipv4Timeout.(int),
			Ipv6Timeout:            ipv6Timeout.(int),
			TrustedOnAllInterfaces: truested.(bool),
			Kind:                   "object#DHCPRelayGlobalSettings",
		},
	)

	if err != nil {
		return fmt.Errorf(
			"Error deleting DHCP Relay Global Settings: %v", err)
	}

	return nil
}
