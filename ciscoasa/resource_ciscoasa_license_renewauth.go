package ciscoasa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func resourceCiscoASALicenseRenewAuth() *schema.Resource {
	return &schema.Resource{
		Create:        resourceCiscoASALicenseRenewAuthCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,
	}
}

func resourceCiscoASALicenseRenewAuthCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Licensing.RenewAuthLicense()
	if err != nil {
		return fmt.Errorf(
			"Error creating LicenseRenewAuth: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}
