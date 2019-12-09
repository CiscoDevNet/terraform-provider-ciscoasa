package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASANetworkServiceGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASANetworkServiceGroupCreate,
		Read:   resourceCiscoASANetworkServiceGroupRead,
		Update: resourceCiscoASANetworkServiceGroupUpdate,
		Delete: resourceCiscoASANetworkServiceGroupDelete,

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

func resourceCiscoASANetworkServiceGroupCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	members := []string{}
	for _, member := range d.Get("members").(*schema.Set).List() {
		members = append(members, member.(string))
	}

	n, err := ca.Objects.CreateNetworkServiceGroup(d.Get("name").(string), "", members)
	if err != nil {
		return fmt.Errorf("Error creating Network Service Group %s: %v", d.Get("name").(string), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkServiceGroupRead(d, meta)
}

func resourceCiscoASANetworkServiceGroupRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.GetNetworkServiceGroup(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf("[DEBUG] Network Service Group %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading Network Service Group %s: %v", d.Id(), err)
	}

	members := resourceCiscoASANetworkServiceGroup().Schema["members"].ZeroValue().(*schema.Set)
	for _, member := range n.Members {
		members.Add(member.String())
	}

	d.Set("members", members)

	return nil
}

func resourceCiscoASANetworkServiceGroupUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	members := []string{}
	for _, member := range d.Get("members").(*schema.Set).List() {
		members = append(members, member.(string))
	}

	n, err := ca.Objects.UpdateNetworkServiceGroup(d.Id(), "", members)
	if err != nil {
		return fmt.Errorf("Error updating Network Service Group %s: %v", d.Get("name").(string), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASANetworkServiceGroupRead(d, meta)
}

func resourceCiscoASANetworkServiceGroupDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Objects.DeleteNetworkServiceGroup(d.Id())
	if err != nil {
		if err != nil {
			return fmt.Errorf("Error deleting Network Service Group %s: %v", d.Id(), err)
		}
	}

	return nil
}
