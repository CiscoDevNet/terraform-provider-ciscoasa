package ciscoasa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCiscoASALicenseConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASALicenseConfigUpdate,
		Read:   resourceCiscoASALicenseConfigRead,
		Update: resourceCiscoASALicenseConfigUpdate,
		Delete: schema.Noop,

		Schema: map[string]*schema.Schema{
			"throughput": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"100M",
					"1G",
					"2G",
				}, false),
			},

			"license_server_url": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "https://tools.cisco.com/its/service/oddce/services/DDCEService",
			},

			"transport_url": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},

			"privacy_host_name": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"privacy_version": {
				Type:     schema.TypeBool,
				Optional: true,
			},
		},
	}
}

func resourceCiscoASALicenseConfigUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Licensing.UpdateLicenseConfig(&ciscoasa.SmartLicenseConfig{
		LicenseServerURL: d.Get("license_server_url").(string),
		TransportURL:     d.Get("transport_url").(bool),
		PrivacyHostName:  d.Get("privacy_host_name").(bool),
		PrivacyVersion:   d.Get("privacy_version").(bool),
		Throughput:       d.Get("throughput").(string),
		FeatureTier:      "standard",
	})

	if err != nil {
		return fmt.Errorf(
			"Error creating License Config: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return resourceCiscoASALicenseConfigRead(d, meta)
}

func resourceCiscoASALicenseConfigRead(d *schema.ResourceData, meta interface{}) error {

	ca := meta.(*ciscoasa.Client)

	r, err := ca.Licensing.GetLicenseConfig()
	if err != nil {
		return fmt.Errorf("Error reading License Config: %v", err)
	}

	d.Set("license_server_url", r.LicenseServerURL)
	d.Set("transport_url", r.TransportURL)
	d.Set("privacy_host_name", r.PrivacyHostName)
	d.Set("privacy_version", r.PrivacyVersion)
	d.Set("throughput", r.Throughput)

	return nil
}
