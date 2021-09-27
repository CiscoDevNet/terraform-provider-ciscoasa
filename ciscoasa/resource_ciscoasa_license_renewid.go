package ciscoasa

import (
	"fmt"
	"strconv"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCiscoASALicenseRenewId() *schema.Resource {
	return &schema.Resource{
		Create:        resourceCiscoASALicenseRenewIdCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,
	}
}

func resourceCiscoASALicenseRenewIdCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Licensing.RenewIdLicense()
	if err != nil {
		return fmt.Errorf(
			"Error creating LicenseRenewId: %v", err)
	}

	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return nil
}
