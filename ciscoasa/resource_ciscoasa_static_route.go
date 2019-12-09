package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASAStaticRoute() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASAStaticRouteCreate,
		Read:   resourceCiscoASAStaticRouteRead,
		Update: resourceCiscoASAStaticRouteUpdate,
		Delete: resourceCiscoASAStaticRouteDelete,

		Schema: map[string]*schema.Schema{
			"interface": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"network": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"gateway": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},

			"metric": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Default:  1,
			},

			"tracked": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"tunneled": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceCiscoASAStaticRouteCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	routeID, err := ca.Routing.CreateStaticRoute(
		d.Get("interface").(string),
		cidrToAddress(d.Get("network").(string)),
		d.Get("gateway").(string),
		d.Get("metric").(int),
		d.Get("tracked").(bool),
		d.Get("tunneled").(bool),
	)
	if err != nil {
		return fmt.Errorf(
			"Error creating Static Route %s->%s: %v", d.Get("network").(string), d.Get("gateway").(string), err)
	}

	d.SetId(routeID)

	return resourceCiscoASAStaticRouteRead(d, meta)
}

func resourceCiscoASAStaticRouteRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := ca.Routing.GetStaticRoute(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf(
				"[DEBUG] Static Route for %s->%s no longer exists", d.Get("network").(string), d.Get("gateway").(string))
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Error reading Static Route %s->%s: %v", d.Get("network").(string), d.Get("gateway").(string), err)
	}

	d.Set("interface", r.Interface.Name)
	d.Set("network", addressToCIDR(r.Network.String()))
	d.Set("gateway", r.Gateway.String())
	d.Set("metric", r.DistanceMetric)
	d.Set("tunneled", r.Tunneled)
	d.Set("tracked", r.Tracked)

	return nil
}

func resourceCiscoASAStaticRouteUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	routeID, err := ca.Routing.UpdateStaticRoute(
		d.Id(),
		d.Get("interface").(string),
		cidrToAddress(d.Get("network").(string)),
		d.Get("gateway").(string),
		d.Get("metric").(int),
		d.Get("tracked").(bool),
		d.Get("tunneled").(bool),
	)
	if err != nil {
		return fmt.Errorf(
			"Error updating Static Route %s->%s: %v", d.Get("network").(string), d.Get("gateway").(string), err)
	}

	d.SetId(routeID)

	return resourceCiscoASAStaticRouteRead(d, meta)
}

func resourceCiscoASAStaticRouteDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Routing.DeleteStaticRoute(d.Id())
	if err != nil {
		return fmt.Errorf(
			"Error deleting Static Route %s->%s: %v", d.Get("network").(string), d.Get("gateway").(string), err)
	}

	return nil
}
