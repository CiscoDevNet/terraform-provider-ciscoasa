package ciscoasa

import (
	"fmt"
	"math/rand"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
)

func resourceCiscoASAWriteMemory() * schema.Resource {
	return &schema.Resource{
		Description: `The ` + "`null_resource`" + ` resource implements the standard resource lifecycle but takes no further action.
The ` + "`triggers`" + ` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.`,

		Create: resourceCiscoASAWriteMemoryCreate,
		Read:   resourceCiscoASAWriteMemoryRead,
		Delete: resourceCiscoASAWriteMemoryDelete,

		Schema: map[string]*schema.Schema{
			"triggers": {
				Description: "A map of arbitrary strings that, when changed, will force the write_memory resource to be replaced, re-running saving changes.",
				Type:        schema.TypeMap,
				Optional:    true,
				ForceNew:    true,
			},

			"id": {
				Description: "This is set to a random value at create time.",
				Computed:    true,
				Type:        schema.TypeString,
			},
		},
	}
}

func resourceCiscoASAWriteMemoryCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Save.WriteMem()
	if err != nil {
		return fmt.Errorf("Error saving changes: %v", err)
	}


	d.SetId(fmt.Sprintf("%d", rand.Int()))

	return nil
}

func resourceCiscoASAWriteMemoryRead(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func resourceCiscoASAWriteMemoryDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}