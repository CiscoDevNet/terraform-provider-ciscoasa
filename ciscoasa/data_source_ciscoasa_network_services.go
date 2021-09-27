package ciscoasa

import (
	"context"
	"strconv"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func dataSourceCiscoASANetworkServices() *schema.Resource {

	return &schema.Resource{
		ReadContext: dataSourceCiscoASANetworkServicesRead,

		Schema: map[string]*schema.Schema{
			"network_services": {
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

func dataSourceCiscoASANetworkServicesRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	ca := meta.(*ciscoasa.Client)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics

	l, err := ca.Objects.ListNetworkServices()
	if err != nil {
		return diag.FromErr(err)
	}

	// Create an empty list to hold all NetworkServices
	networkServices := make([]interface{}, len(l.Items), len(l.Items))

	for i, r := range l.Items {
		networkServiceMap := make(map[string]interface{})
		networkServiceMap["name"] = r.ObjectID
		networkServiceMap["value"] = r.Value

		networkServices[i] = networkServiceMap
	}

	if err := d.Set("network_services", networkServices); err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	return diags
}
