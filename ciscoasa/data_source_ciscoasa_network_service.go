package ciscoasa

import (
	"context"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func dataSourceCiscoASANetworkService() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceCiscoASANetworkServiceRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},

			"value": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceCiscoASANetworkServiceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ca := meta.(*ciscoasa.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	objectID := d.Get("name").(string)

	r, err := ca.Objects.GetNetworkService(objectID)
	if err != nil {
		return diag.FromErr(err)
	}

	if err := d.Set("value", r.Value); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(objectID)

	return diags
}
