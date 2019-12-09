package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASANetworkObject() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASANetworkObjectCreate,
		Read:   resourceCiscoASANetworkObjectRead,
		Update: resourceCiscoASANetworkObjectUpdate,
		Delete: resourceCiscoASANetworkObjectDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func resourceCiscoASANetworkObjectCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.CreateNetworkObject(d.Get("name").(string), "", d.Get("value").(string))
	if err != nil {
		return fmt.Errorf("Error creating Network Object %s: %v", d.Get("name").(string), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkObjectRead(d, meta)
}

func resourceCiscoASANetworkObjectRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.GetNetworkObject(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf("[DEBUG] Network Object %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading Network Object %s: %v", d.Id(), err)
	}

	d.Set("name", n.Name)
	d.Set("value", n.Host.Value)

	return nil
}

func resourceCiscoASANetworkObjectUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.UpdateNetworkObject(d.Id(), "", d.Get("value").(string))
	if err != nil {
		return fmt.Errorf("Error updating Network Object %s: %v", d.Id(), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkObjectRead(d, meta)
}

func resourceCiscoASANetworkObjectDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Objects.DeleteNetworkObject(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Network Object %s: %v", d.Id(), err)
	}

	return nil
}
