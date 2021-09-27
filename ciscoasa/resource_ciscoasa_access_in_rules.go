package ciscoasa

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCiscoASAAccessInRules() *schema.Resource {
	return &schema.Resource{
		Create:        resourceCiscoASAAccessInRulesCreate,
		Read:          resourceCiscoASAAccessInRulesRead,
		Update:        resourceCiscoASAAccessInRulesUpdate,
		Delete:        resourceCiscoASAAccessInRulesDelete,
		CustomizeDiff: resourceCiscoASAAccessInRulesDiff,

		Schema: map[string]*schema.Schema{
			"interface": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"managed": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"rule": {
				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"source": {
							Type:     schema.TypeString,
							Required: true,
						},

						"source_service": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"destination": {
							Type:     schema.TypeString,
							Required: true,
						},

						"destination_service": {
							Type:     schema.TypeString,
							Required: true,
						},

						"active": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						"permit": {
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						"time_range": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"id": {
							Type:     schema.TypeString,
							Computed: true,
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

func resourceCiscoASAAccessInRulesDiff(_ context.Context, d *schema.ResourceDiff, meta interface{}) error {
	o, n := d.GetChange("rule")
	if len(o.([]interface{})) != len(n.([]interface{})) {
		d.ForceNew("rule")
		return nil
	}

	return nil
}

func resourceCiscoASAAccessInRulesCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	// We need to set this upfront in order to be able to save a partial state
	d.SetId(d.Get("interface").(string))

	// Create all rules that are configured
	if nrs := d.Get("rule").([]interface{}); len(nrs) > 0 {
		// Create an empty rule list to hold all newly created rules
		rules := resourceCiscoASAAccessInRules().Schema["rule"].ZeroValue().([]interface{})

		for _, rule := range nrs {
			rule := rule.(map[string]interface{})

			id, err := ca.Access.CreateAccessInRule(
				d.Id(),
				cidrToAddress(rule["source"].(string)),
				rule["source_service"].(string),
				cidrToAddress(rule["destination"].(string)),
				rule["destination_service"].(string),
				rule["time_range"].(string),
				rule["active"].(bool),
				rule["permit"].(bool),
			)

			if err != nil {
				return fmt.Errorf("Error creating ACE on interface %s: %v", d.Id(), err)
			}

			rule["id"] = id
			rules = append(rules, rule)

			// We need to update this first to preserve the correct state
			d.Set("rule", rules)
		}
	}

	return resourceCiscoASAAccessInRulesRead(d, meta)
}

func createCiscoASAAccessInRulesRules(meta interface{}, iface string, rules []interface{}, nrs []interface{}) ([]interface{}, error) {
	ca := meta.(*ciscoasa.Client)

	for _, rule := range nrs {
		rule := rule.(map[string]interface{})

		id, err := ca.Access.CreateAccessInRule(
			iface,
			cidrToAddress(rule["source"].(string)),
			rule["source_service"].(string),
			cidrToAddress(rule["destination"].(string)),
			rule["destination_service"].(string),
			rule["time_range"].(string),
			rule["active"].(bool),
			rule["permit"].(bool),
		)
		if err != nil {
			return rules, fmt.Errorf("Error creating ACE on interface %s: %v", iface, err)
		}

		rule["id"] = id
		rules = append(rules, rule)
	}

	return rules, nil
}

func resourceCiscoASAAccessInRulesRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	// Get all the rules from the running environment
	l, err := ca.Access.ListAccessInRules(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf("[DEBUG] Rule %s no longer exists", d.Id())
			d.SetId("")
			return err
		}

		return fmt.Errorf("Error reading interface %s rules: %v", d.Id(), err)
	}

	// Make a map of all the rules so we can easily find a rule
	ruleMap := make(map[string]*ciscoasa.ExtendedACEObject, l.RangeInfo.Total)
	for _, r := range l.Items {
		ruleMap[r.ObjectID] = r
	}

	// Create an empty list to hold all rules
	rules := resourceCiscoASAAccessInRules().Schema["rule"].ZeroValue().([]interface{})

	// Read all rules that are configured
	if rs := d.Get("rule").([]interface{}); len(rs) > 0 {
		for _, rule := range rs {
			rule := rule.(map[string]interface{})
			id := rule["id"].(string)

			// Get the rule
			r, ok := ruleMap[id]
			if !ok {
				continue
			}

			// Delete the known rule so only unknown rules remain in the ruleMap
			delete(ruleMap, id)

			rule["source"] = addressToCIDR(r.SrcAddress.String())
			rule["destination"] = addressToCIDR(r.DstAddress.String())
			rule["destination_service"] = r.DstService.String()
			rule["active"] = r.Active
			rule["permit"] = r.Permit
			rules = append(rules, rule)
		}
	}

	// If this is a managed firewall, add all unknown rules into dummy rules
	managed := d.Get("managed").(bool)
	if managed && len(ruleMap) > 0 {
		for _, r := range ruleMap {
			rule := make(map[string]interface{})

			rule["source"] = addressToCIDR(r.SrcAddress.String())
			rule["destination"] = addressToCIDR(r.DstAddress.String())
			rule["destination_service"] = r.DstService.String()
			rule["active"] = r.Active
			rule["permit"] = r.Permit
			rule["id"] = r.ObjectID
			rules = append(rules, rule)
		}
	}

	if len(rules) > 0 {
		d.Set("rule", rules)
	} else if !managed {
		d.SetId("")
	}

	return nil
}

func resourceCiscoASAAccessInRulesUpdate(d *schema.ResourceData, meta interface{}) error {
	// Check if the rule set as a whole has changed
	if d.HasChange("rule") {
		o, n := d.GetChange("rule")
		// Create a list with old rules to track successfully updated rules
		rules := o.([]interface{})
		ca := meta.(*ciscoasa.Client)

		for i, rule := range n.([]interface{}) {
			rule := rule.(map[string]interface{})
			id, err := ca.Access.UpdateAccessInRule(
				d.Id(),
				rule["id"].(string),
				cidrToAddress(rule["source"].(string)),
				rule["source_service"].(string),
				cidrToAddress(rule["destination"].(string)),
				rule["destination_service"].(string),
				rule["time_range"].(string),
				rule["active"].(bool),
				rule["permit"].(bool),
			)
			if err != nil {
				return fmt.Errorf("Error updating ACE on interface %s: %v", d.Id(), err)
			}
			rule["id"] = id
			rules[i] = rule
			d.Set("rule", rules)
		}

		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCiscoASAAccessInRulesRead(d, meta)
}

func resourceCiscoASAAccessInRulesDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	// Delete all rules
	if ors := d.Get("rule").([]interface{}); len(ors) > 0 {
		// Each rule that is successfully deleted will be removed
		// from this list, leaving only rules that
		// could not be deleted properly and should be saved in the state.
		for i := 0; i < len(ors); i++ {
			rule := ors[i].(map[string]interface{})

			err := ca.Access.DeleteAccessInRule(d.Id(), rule["id"].(string))
			if err != nil {
				if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
					log.Printf(
						"[DEBUG] ACE %s from interface %s no longer exists", rule["id"].(string), d.Id())
					continue
				}

				// We need to update this first to preserve the correct state
				d.Set("rule", ors)

				return fmt.Errorf("Error deleting ACE %s from interface %s: %v", rule["id"].(string), d.Id(), err)
			}
			ors = append(ors[:i], ors[i+1:]...)
			i--
		}
	}

	return nil
}
