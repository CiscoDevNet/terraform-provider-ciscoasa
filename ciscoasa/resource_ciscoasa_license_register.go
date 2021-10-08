package ciscoasa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCiscoASALicenseRegister() *schema.Resource {
	return &schema.Resource{
		Create:        resourceCiscoASALicenseRegisterCreate,
		ReadContext:   schema.NoopContext,
		Delete: resourceCiscoASALicenseRegisterDelete,

		Schema: map[string]*schema.Schema{
			"id_token": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"force": {
				Type:     schema.TypeBool,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCiscoASALicenseRegisterCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Licensing.RegisterLicense(
		d.Get("id_token").(string),
		d.Get("force").(bool),
	)
	if err != nil {
		return fmt.Errorf(
			"Error creating LicenseRegister: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}

func resourceCiscoASALicenseRegisterDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Licensing.DeregisterLicense()
	if err != nil {
		return fmt.Errorf(
			"Error deleting LicenseRegister: %v", err)
	}

	d.SetId("")

	return nil
}
