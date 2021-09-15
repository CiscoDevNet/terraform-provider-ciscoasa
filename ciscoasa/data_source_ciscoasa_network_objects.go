package ciscoasa

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func dataSourceCiscoASANetworkObjects() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceCiscoASANetworkObjectsRead,

		Schema: map[string]*schema.Schema{
			"network_objects": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},

						"value": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceCiscoASANetworkObjectsRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ca := meta.(*ciscoasa.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	l, err := ca.Objects.ListNetworkObjects()
	if err != nil {
		return diag.FromErr(err)
	}

	// Create an empty list to hold all networkObjects
	networkObjects := make([]interface{}, len(l.Items), len(l.Items))

	for i, r := range l.Items {
		networkObjectMap := make(map[string]interface{})
		networkObjectMap["name"] = r.ObjectID
		networkObjectMap["value"] = r.Host.Value

		networkObjects[i] = networkObjectMap
	}

	if err := d.Set("network_objects", networkObjects); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
