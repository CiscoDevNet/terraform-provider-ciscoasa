package ciscoasa

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCiscoASADhcpRelayLocal() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASADhcpRelayLocalCreate,
		Read:   resourceCiscoASADhcpRelayLocalRead,
		Update: resourceCiscoASADhcpRelayLocalUpdate,
		Delete: resourceCiscoASADhcpRelayLocalDelete,

		Schema: map[string]*schema.Schema{
			"interface": {
				Type:     schema.TypeString,
				Required: true,
			},

			"servers": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},

			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCiscoASADhcpRelayLocalCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	iface := d.Get("interface").(string)
	servers := d.Get("servers").([]interface{})
	s := make([]string, len(servers))

	for i := range s {
		s[i] = servers[i].(string)
	}

	err := ca.Dhcp.CreateDhcpRelayLocal(
		iface,
		s,
	)
	if err != nil {
		return fmt.Errorf(
			"Error creating DHCP Relay interface server %s: %v", iface, err)
	}

	d.SetId(iface)

	return resourceCiscoASADhcpRelayLocalRead(d, meta)
}

func resourceCiscoASADhcpRelayLocalRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.Dhcp.GetDhcpRelayLocal(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf(
				"[DEBUG] DHCP Relay interface server for %s no longer exists", d.Get("interface").(string))
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Error reading DHCP Relay interface server %s: %v", d.Get("interface").(string), err)
	}

	d.Set("interface", r.Interface)
	d.Set("servers", r.Servers)

	return nil
}

func resourceCiscoASADhcpRelayLocalUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	iface := d.Get("interface").(string)
	servers := d.Get("servers").([]interface{})
	s := make([]string, len(servers))

	for i := range s {
		s[i] = servers[i].(string)
	}

	err := ca.Dhcp.UpdateDhcpRelayLocal(
		iface,
		s,
	)
	if err != nil {
		return fmt.Errorf(
			"Error updating DHCP Relay interface server %s: %v", iface, err)
	}

	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceCiscoASADhcpRelayLocalRead(d, meta)
}

func resourceCiscoASADhcpRelayLocalDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Dhcp.DeleteDhcpRelayLocal(d.Id())
	if err != nil {
		return fmt.Errorf(
			"Error deleting DHCP Relay interface server %s: %v", d.Get("interface").(string), err)
	}

	return nil
}
