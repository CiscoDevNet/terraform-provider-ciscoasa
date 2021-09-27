package ciscoasa

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func resourceCiscoASABackup() *schema.Resource {
	return &schema.Resource{
		Create:        resourceCiscoASABackupCreate,
		ReadContext:   schema.NoopContext,
		DeleteContext: schema.NoopContext,

		Schema: map[string]*schema.Schema{
			"context": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"location": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "",
				ForceNew: true,
			},

			"passphrase": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
		},
	}
}

func resourceCiscoASABackupCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	context := d.Get("context").(string)

	err := ca.CreateBackup(
		context,
		d.Get("location").(string),
		d.Get("passphrase").(string),
	)
	if err != nil {
		return fmt.Errorf(
			"Error creating Backup: %v", err)
	}

	if context != "" {
		d.SetId(context)
	} else {
		d.SetId("full")
	}

	return nil
}
