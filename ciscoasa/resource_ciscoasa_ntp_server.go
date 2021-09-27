package ciscoasa

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCiscoASANtpServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASANtpServerCreate,
		Read:   resourceCiscoASANtpServerRead,
		Update: resourceCiscoASANtpServerUpdate,
		Delete: resourceCiscoASANtpServerDelete,

		Schema: map[string]*schema.Schema{
			"ip_address": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},

			"preferred": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"interface": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"key_number": {
				Type:     schema.TypeString,
				Required: true,
			},

			"key_value": {
				Type:     schema.TypeString,
				Required: true,
			},

			"key_trusted": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCiscoASANtpServerCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	ipAddress := d.Get("ip_address").(string)
	iface := d.Get("interface").(string)

	if isInterfaceHwId(iface) {
		iface = objectIdFromHwId(iface)
	}

	err := ca.DeviceSetup.CreateNtpServer(
		ipAddress,
		d.Get("preferred").(bool),
		iface,
		d.Get("key_number").(string),
		d.Get("key_value").(string),
		d.Get("key_trusted").(bool),
	)
	if err != nil {
		return fmt.Errorf(
			"Error creating NTP Server %s->%s: %v", ipAddress, d.Get("key_number").(string), err)
	}

	d.SetId(ipAddress)

	return resourceCiscoASANtpServerRead(d, meta)
}

func resourceCiscoASANtpServerRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.DeviceSetup.GetNtpServer(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf(
				"[DEBUG] NTP Server for %s->%s no longer exists", d.Get("ip_address").(string), d.Get("key_number").(string))
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Error reading NTP Server %s->%s: %v", d.Get("ip_address").(string), d.Get("key_number").(string), err)
	}

	d.Set("ip_address", r.IpAddress)
	d.Set("preferred", r.IsPreferred)
	if r.Interface != nil {
		if d.Get("interface").(string) == hwIdFromObjId(r.Interface.ObjectId) {
			d.Set("interface", hwIdFromObjId(r.Interface.ObjectId))
		} else {
			d.Set("interface", r.Interface.Name)
		}
	}
	d.Set("key_number", r.Key.Number)
	d.Set("key_value", r.Key.Value)
	d.Set("key_trusted", r.Key.IsTrusted)

	return nil
}

func resourceCiscoASANtpServerUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	ipAddress := d.Get("ip_address").(string)
	iface := d.Get("interface").(string)

	if isInterfaceHwId(iface) {
		iface = objectIdFromHwId(iface)
	}

	err := ca.DeviceSetup.UpdateNtpServer(
		d.Id(),
		ipAddress,
		d.Get("preferred").(bool),
		iface,
		d.Get("key_number").(string),
		d.Get("key_value").(string),
		d.Get("key_trusted").(bool),
	)
	if err != nil {
		return fmt.Errorf(
			"Error updating NTP Server %s->%s: %v", ipAddress, d.Get("key_number").(string), err)
	}

	d.SetId(ipAddress)
	d.Set("last_updated", time.Now().Format(time.RFC850))

	return resourceCiscoASANtpServerRead(d, meta)
}

func resourceCiscoASANtpServerDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.DeviceSetup.DeleteNtpServer(d.Id())
	if err != nil {
		return fmt.Errorf(
			"Error deleting NTP Server %s->%s: %v", d.Get("ip_address").(string), d.Get("key_number").(string), err)
	}

	return nil
}
