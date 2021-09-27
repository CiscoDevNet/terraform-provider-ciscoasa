package ciscoasa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASALicenseRegister() *schema.Resource {
	return &schema.Resource{
		Create:        resourceCiscoASALicenseRegisterCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,

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
