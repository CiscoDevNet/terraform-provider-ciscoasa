package ciscoasa

import (
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASAFailoverInterface() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASAFailoverInterfaceUpdate,
		Read:   resourceCiscoASAFailoverInterfaceRead,
		Update: resourceCiscoASAFailoverInterfaceUpdate,
		Delete: resourceCiscoASAFailoverInterfaceDelete,

		Schema: map[string]*schema.Schema{
			"hardware_id": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"standby_ip": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.IsIPAddress,
			},

			"monitored": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
		},
	}
}

// There is no real creation of the Failover Interface.
// All the changes are made on already existing object,
// which is created with the Interface.
func resourceCiscoASAFailoverInterfaceUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	FailoverInterfaceId := objectIdFromHwId(d.Get("hardware_id").(string))

	FailoverInterface, err := getFailoverInterface(FailoverInterfaceId, meta)
	if err != nil {
		return err
	}

	StandbyIPAddress := &ciscoasa.Address{
		Kind:  "IPv4Address",
		Value: d.Get("standby_ip").(string),
	}

	FailoverInterface.StandbyIPAddress = StandbyIPAddress
	FailoverInterface.IsMonitored = d.Get("monitored").(bool)

	err = ca.Failover.UpdateFailoverInterface(FailoverInterface)

	if err != nil {
		return fmt.Errorf(
			"Error creating Failover Interface %s: %v", FailoverInterface.InterfaceName, err)
	}

	d.SetId(FailoverInterfaceId)

	return resourceCiscoASAFailoverInterfaceRead(d, meta)
}

func resourceCiscoASAFailoverInterfaceRead(d *schema.ResourceData, meta interface{}) error {

	r, err := getFailoverInterface(d.Id(), meta)
	if err != nil {
		return err
	}

	d.Set("hardware_id", r.InterfaceName)
	d.Set("standby_ip", r.StandbyIPAddress.Value)
	d.Set("monitored", r.IsMonitored)

	return nil
}

func resourceCiscoASAFailoverInterfaceDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	r, err := getFailoverInterface(d.Id(), meta)
	if err != nil {
		return err
	}

	r.StandbyIPAddress = nil
	r.IsMonitored = true

	err = ca.Failover.UpdateFailoverInterface(r)
	if err != nil {
		return fmt.Errorf(
			"Error deleting Failover Interface %s: %v", r.InterfaceName, err)
	}

	return nil
}

func getFailoverInterface(id string, meta interface{}) (*ciscoasa.FailoverInterface, error) {
	ca := meta.(*ciscoasa.Client)
	r, err := ca.Failover.GetFailoverInterface(id)
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			return nil, fmt.Errorf("Failover Interface for %s no longer exists", id)
		}
		return nil, fmt.Errorf("Error reading Failover Interface %s: %v", id, err)
	}
	return r, nil
}
