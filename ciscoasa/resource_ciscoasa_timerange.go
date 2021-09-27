package ciscoasa

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func resourceCiscoASATimeRange() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASATimeRangeCreate,
		Read:   resourceCiscoASATimeRangeRead,
		Update: resourceCiscoASATimeRangeUpdate,
		Delete: resourceCiscoASATimeRangeDelete,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"value": {
				Type:     schema.TypeList,
				Required: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"start": {
							Type:     schema.TypeString,
							Required: true,
						},

						"end": {
							Type:     schema.TypeString,
							Required: true,
						},

						"periodic": {
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"frequency": {
										Type:     schema.TypeString,
										Required: true,
									},

									"start_hour": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(0, 23),
									},

									"start_minute": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(0, 59),
									},

									"end_hour": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(0, 23),
									},

									"end_minute": {
										Type:         schema.TypeInt,
										Required:     true,
										ValidateFunc: validation.IntBetween(0, 59),
									},
								},
							},
						},
					},
				},
			},
			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCiscoASATimeRangeCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	name := d.Get("name").(string)
	value := d.Get("value").([]interface{})
	valueAttrs := value[0].(map[string]interface{})
	start := valueAttrs["start"].(string)
	end := valueAttrs["end"].(string)
	periodic := make([]*ciscoasa.Periodic, 0)
	periodicAttrs := valueAttrs["periodic"].(*schema.Set)

	if periodicAttrs.Len() > 0 {
		periodic = expandPeriodic(periodicAttrs)
	}

	n, err := ca.Objects.CreateTimeRange(name, start, end, periodic)
	if err != nil {
		return fmt.Errorf("Error creating Time Range %s: %v", d.Get("name").(string), err)
	}

	d.SetId(n.Name)

	return resourceCiscoASATimeRangeRead(d, meta)
}

func resourceCiscoASATimeRangeRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	n, err := ca.Objects.GetTimeRange(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf("[DEBUG] Time Range %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf("Error reading Time Range %s: %v", d.Id(), err)
	}

	d.Set("name", n.Name)
	d.Set("value", flattenTRValue(n.Value))

	return nil
}

func resourceCiscoASATimeRangeUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	if d.HasChange("value") {
		value := d.Get("value").([]interface{})
		valueAttrs := value[0].(map[string]interface{})
		start := valueAttrs["start"].(string)
		end := valueAttrs["end"].(string)
		periodic := make([]*ciscoasa.Periodic, 0)
		periodicAttrs := valueAttrs["periodic"].(*schema.Set)

		if periodicAttrs.Len() > 0 {
			periodic = expandPeriodic(periodicAttrs)
		}

		_, err := ca.Objects.UpdateTimeRange(d.Id(), start, end, periodic)
		if err != nil {
			return fmt.Errorf("Error updating Time Range %s: %v", d.Id(), err)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCiscoASATimeRangeRead(d, meta)
}

func resourceCiscoASATimeRangeDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	err := ca.Objects.DeleteTimeRange(d.Id())
	if err != nil {
		return fmt.Errorf("Error deleting Time Range %s: %v", d.Id(), err)
	}

	return nil
}

func flattenTRValue(in *ciscoasa.TRValue) []interface{} {
	var out = make([]interface{}, 1, 1)
	m := make(map[string]interface{})
	m["start"] = in.Start
	m["end"] = in.End
	if len(in.Periodic) > 0 {
		m["periodic"] = flattenPeriodic(in.Periodic)
	}
	out[0] = m
	return out
}

func flattenPeriodic(in []*ciscoasa.Periodic) []map[string]interface{} {

	var out = make([]map[string]interface{}, len(in), len(in))
	for i, v := range in {
		m := make(map[string]interface{})
		m["frequency"] = v.Frequency
		m["start_hour"] = v.StartHour
		m["start_minute"] = v.StartMinute
		m["end_hour"] = v.EndHour
		m["end_minute"] = v.EndMinute
		out[i] = m
	}
	return out
}

func expandPeriodic(periodic *schema.Set) []*ciscoasa.Periodic {

	result := make([]*ciscoasa.Periodic, periodic.Len(), periodic.Len())

	for i, p := range periodic.List() {
		out := &ciscoasa.Periodic{}
		v := p.(map[string]interface{})
		out.Frequency = v["frequency"].(string)
		out.StartHour = v["start_hour"].(int)
		out.StartMinute = v["start_minute"].(int)
		out.EndHour = v["end_hour"].(int)
		out.EndMinute = v["end_minute"].(int)
		result[i] = out
	}

	return result
}
