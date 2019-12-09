package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASANetworkObjectGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASANetworkObjectGroupCreate,
		Read:   resourceCiscoASANetworkObjectGroupRead,
		Update: resourceCiscoASANetworkObjectGroupUpdate,
		Delete: resourceCiscoASANetworkObjectGroupDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"members": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
		},
	}
}

func resourceCiscoASANetworkObjectGroupCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	members := []string{}
	for _, member := range d.Get("members").(*schema.Set).List() {
		members = append(members, member.(string))
	}

	n, err := ca.Objects.CreateNetworkObjectGroup(d.Get("name").(string), "", members)
	if err != nil {
		return fmt.Errorf("Error creating Network Object Group %s: %v", d.Get("name").(string), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkObjectGroupRead(d, meta)
}

func resourceCiscoASANetworkObjectGroupRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.GetNetworkObjectGroup(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf("[DEBUG] Network Object Group %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading Network Object Group %s: %v", d.Id(), err)
	}

	members := resourceCiscoASANetworkObjectGroup().Schema["members"].ZeroValue().(*schema.Set)
	for _, member := range n.Members {
		members.Add(member.String())
	}

	d.Set("members", members)

	return nil
}

func resourceCiscoASANetworkObjectGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	members := []string{}
	for _, member := range d.Get("members").(*schema.Set).List() {
		members = append(members, member.(string))
	}

	n, err := ca.Objects.UpdateNetworkObjectGroup(d.Id(), "", members)
	if err != nil {
		return fmt.Errorf("Error updating Network Object Group %s: %v", d.Id(), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkObjectGroupRead(d, meta)
}

func resourceCiscoASANetworkObjectGroupDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Objects.DeleteNetworkObjectGroup(d.Id())
	if err != nil {
		if err != nil {
			return fmt.Errorf("Error deleting Network Object Group %s: %v", d.Id(), err)
		}
	}

	return nil
}
