package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCiscoASANetworkService() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASANetworkServiceCreate,
		Read:   resourceCiscoASANetworkServiceRead,
		Update: resourceCiscoASANetworkServiceUpdate,
		Delete: resourceCiscoASANetworkServiceDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": {
				Type:     schema.TypeString,
				Required: true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if strings.Replace(old, "/1-65535,", ",", 1) == strings.Replace(new, "/1-65535,", ",", 1) {
						return true
					}
					return false
				},
			},
		},
	}
}

func resourceCiscoASANetworkServiceCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.CreateNetworkService(d.Get("name").(string), "", d.Get("value").(string))
	if err != nil {
		return fmt.Errorf("Error creating Network Service %s: %v", d.Get("name").(string), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkServiceRead(d, meta)
}

func resourceCiscoASANetworkServiceRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.GetNetworkService(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf("[DEBUG] Network Service %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading Network Service %s: %v", d.Id(), err)
	}

	d.Set("name", n.Name)
	d.Set("value", n.Value)

	return nil
}

func resourceCiscoASANetworkServiceUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.UpdateNetworkService(d.Id(), "", d.Get("value").(string))
	if err != nil {
		return fmt.Errorf("Error updating Network Service %s: %v", d.Id(), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkServiceRead(d, meta)
}

func resourceCiscoASANetworkServiceDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Objects.DeleteNetworkService(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Network Service %s: %v", d.Id(), err)
	}

	return nil
}
