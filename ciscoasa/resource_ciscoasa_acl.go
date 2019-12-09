package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

var aclRule = &schema.Resource{
	Schema: map[string]*schema.Schema{
		"source": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},

		"source_service": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
		},

		"destination": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},

		"destination_service": &schema.Schema{
			Type:     schema.TypeString,
			Required: true,
		},

		"active": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},

		"log_interval": &schema.Schema{
			Type:     schema.TypeInt,
			Optional: true,
			Default:  300,
		},

		"log_status": &schema.Schema{
			Type:     schema.TypeString,
			Optional: true,
			Default:  "Default",
			ValidateFunc: validation.StringInSlice(
				[]string{"Default", "Debugging", "Disabled", "Notifications", "Critical",
					"Emergencies", "Warnings", "Errors", "Informational", "Alerts"},
				false,
			),
		},

		"permit": &schema.Schema{
			Type:     schema.TypeBool,
			Optional: true,
			Default:  true,
		},

		"remarks": &schema.Schema{
			Type:     schema.TypeList,
			Optional: true,
			Elem:     &schema.Schema{Type: schema.TypeString},
		},

		"id": &schema.Schema{
			Type:     schema.TypeString,
			Computed: true,
		},
	},
}

func resourceCiscoASAACL() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASAACLCreate,
		Read:   resourceCiscoASAACLRead,
		Update: resourceCiscoASAACLUpdate,
		Delete: resourceCiscoASAACLDelete,

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"rule": &schema.Schema{
				Type:     schema.TypeList,
				Required: true,
				Elem:     aclRule,
			},
		},
	}
}

func resourceCiscoASAACLCreate(d *schema.ResourceData, meta interface{}) error {
	// We need to set this upfront in order to be able to save a partial state
	d.SetId(d.Get("name").(string))

	// Create all rules that are configured
	if nrs, ok := d.Get("rule").([]interface{}); ok && len(nrs) > 0 {
		// Create an empty list to hold all newly created rules
		var rules []interface{}

		err := createCiscoASAACLRules(meta, d.Id(), &rules, nrs)

		// We need to update this first to preserve the correct state
		d.Set("rule", rules)

		if err != nil {
			return err
		}
	}

	return resourceCiscoASAACLRead(d, meta)
}

func createCiscoASAACLRules(meta interface{}, acl string, rules *[]interface{}, nrs []interface{}) error {
	ca := meta.(*ciscoasa.Client)

	for _, rule := range nrs {
		rule := rule.(map[string]interface{})

		options := ciscoasa.CreateExtendedACLACEOptions{
			Src:        cidrToAddress(rule["source"].(string)),
			SrcService: rule["source_service"].(string),
			Dst:        cidrToAddress(rule["destination"].(string)),
			DstService: rule["destination_service"].(string),
			Active:     rule["active"].(bool),
			Permit:     rule["permit"].(bool),
		}

		options.RuleLogging = &ciscoasa.RuleLogging{
			LogInterval: rule["log_interval"].(int),
			LogStatus:   rule["log_status"].(string),
		}

		// When we need to insert a rule at a specific position, we add the
		// position to the rule map and use it here to set the position.
		if position, ok := rule["position"]; ok {
			options.Position = position.(int)
		}

		if remarks, ok := rule["remarks"]; ok {
			for _, remark := range remarks.([]interface{}) {
				options.Remarks = append(options.Remarks, remark.(string))
			}
		}

		id, err := ca.Objects.CreateExtendedACLACE(acl, options)
		if err != nil {
			return fmt.Errorf("Error creating ACE on ACL %s: %v", acl, err)
		}

		rule["id"] = id
		*rules = append(*rules, rule)
	}

	return nil
}

func resourceCiscoASAACLRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	// Get all the rules from the running environment
	l, err := ca.Objects.ListExtendedACLACEs(d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf("[DEBUG] ACL %s no longer exists", d.Id())
			d.SetId("")
			return err
		}

		return fmt.Errorf("Error reading ACL %s rules: %v", d.Id(), err)
	}

	// Make a map of all the rules so we can easily find a rule
	ruleMap := make(map[string]*ciscoasa.ExtendedACEObject, l.RangeInfo.Total)
	for _, r := range l.Items {
		ruleMap[r.ObjectID] = r
	}

	// Create an empty list to hold all rules
	var rules []interface{}

	// Read all rules that are configured
	if rs, ok := d.Get("rule").([]interface{}); ok && len(rs) > 0 {
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

			if v, ok := rule["source_service"].(string); ok && v != "" {
				rule["source_service"] = r.SrcService.String()
			}

			if r.RuleLogging != nil {
				rule["log_interval"] = r.RuleLogging.LogInterval
				rule["log_status"] = r.RuleLogging.LogStatus
			}

			if v, ok := rule["remarks"].([]interface{}); ok && len(v) != 0 {
				var remarks []interface{}
				for _, remark := range r.Remarks {
					remarks = append(remarks, remark)
				}
				rule["remarks"] = remarks
			}

			rules = append(rules, rule)
		}
	}

	if len(ruleMap) > 0 {
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
	} else {
		d.SetId("")
	}

	return nil
}

func resourceCiscoASAACLUpdate(d *schema.ResourceData, meta interface{}) error {
	// Check if the rule set as a whole has changed
	if d.HasChange("rule") {
		o, n := d.GetChange("rule")
		ors := o.([]interface{})
		nrs := n.([]interface{})

		// Create three new lists to hold all the old (remove), current (keep)
		// and new (create) rules that need to be deleted, kept and created.
		var remove, keep, create []interface{}

		// Create a temporary set to check which old rules still exist.
		hashResource := schema.HashResource(aclRule)
		nrsSet := schema.NewSet(hashResource, nrs)

		// Filter out all old rules that no longer exist in the new rules while
		// maintaining the correct order of the remaining old rules.
		filtered := ors[:0]
		for _, or := range ors {
			if !nrsSet.Contains(or) {
				remove = append(remove, or)
				continue
			}
			filtered = append(filtered, or)
		}
		ors = filtered

		for nIdx, nr := range nrs {
			nrHash := hashResource(nr)

			// Search for the config rule in the state rules.
			oIdx := -1
			for idx, or := range ors {
				if nrHash == hashResource(or) {
					oIdx = idx
					break
				}
			}

			// If both rules are in the same position, do nothing.
			if nIdx == oIdx+len(create) {
				keep = append(keep, ors[oIdx])
				continue
			}

			// Add the position where to insert this rule to make
			// sure we keep the correct order.
			nr.(map[string]interface{})["position"] = nIdx + 1
			create = append(create, nr)

			if oIdx != -1 {
				remove = append(remove, ors[oIdx])
				ors = append(ors[:oIdx], ors[oIdx+1:]...)
			}
		}

		// First loop through all the old rules and delete them
		if len(remove) > 0 {
			rules := append([]interface{}(nil), remove...)

			err := deleteCiscoASAACLRules(meta, d.Id(), &rules, remove)

			// We need to update this first to preserve the correct state
			d.Set("rule", append(keep, rules...))

			if err != nil {
				return err
			}
		}

		// Then loop through all the new rules and create them
		if len(create) > 0 {
			var rules []interface{}

			err := createCiscoASAACLRules(meta, d.Id(), &rules, create)

			// When we need to insert a rule at a specific position, we add the
			// required position to the rule map. We use that same position to
			// insert the new rule in the keep map so that the order will be the
			// same in the state. And since the aclRule schema doesn't have a
			// position field, we delete the entry from the map before we add it
			// to the keep map and save it to the state
			for _, rule := range rules {
				rule := rule.(map[string]interface{})
				position, ok := rule["position"].(int)
				delete(rule, "position")

				if ok && position > 0 {
					keep = append(keep[:position-1], append([]interface{}{rule}, keep[position-1:]...)...)
				} else {
					keep = append(keep, rule)
				}
			}

			// We need to update this first to preserve the correct state
			d.Set("rule", keep)

			if err != nil {
				return err
			}
		}
	}

	return resourceCiscoASAACLRead(d, meta)
}

func resourceCiscoASAACLDelete(d *schema.ResourceData, meta interface{}) error {
	// Delete all rules
	if ors, ok := d.Get("rule").([]interface{}); ok && len(ors) > 0 {
		// Create an additional list with all the existing rules. Each rule that is
		// succesfully deleted will be removed from this list, leaving only rules that
		// could not be deleted properly and should be saved in the state.
		rules := append([]interface{}(nil), ors...)

		err := deleteCiscoASAACLRules(meta, d.Id(), &rules, ors)

		// We need to update this first to preserve the correct state
		d.Set("rule", rules)

		if err != nil {
			return err
		}
	}

	return nil
}

func deleteCiscoASAACLRules(meta interface{}, acl string, rules *[]interface{}, ors []interface{}) error {
	ca := meta.(*ciscoasa.Client)

	deleted := 0
	for i, rule := range ors {
		rule := rule.(map[string]interface{})

		err := ca.Objects.DeleteExtendedACLACE(acl, rule["id"].(string))
		if err != nil {
			if !strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
				return fmt.Errorf("Error deleting ACE %s from ACL %s: %v", rule["id"].(string), acl, err)
			}
			log.Printf("[DEBUG] ACE %s from ACL %s no longer exists", rule["id"].(string), acl)
		}

		if len(*rules) < (i-deleted)+1 {
			*rules = append((*rules)[:i-deleted], (*rules)[(i-deleted):]...)
		} else {
			*rules = append((*rules)[:i-deleted], (*rules)[(i-deleted)+1:]...)
		}
		deleted++
	}

	return nil
}
