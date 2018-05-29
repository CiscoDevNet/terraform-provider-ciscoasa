package ciscoasa

import (
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

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

			"managed": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},

			"rule": &schema.Schema{
				Type:     schema.TypeSet,
				Required: true,
				Elem: &schema.Resource{
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

						"permit": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
						},

						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceCiscoASAACLCreate(d *schema.ResourceData, meta interface{}) error {
	// We need to set this upfront in order to be able to save a partial state
	d.SetId(d.Get("name").(string))

	// Create all rules that are configured
	if nrs := d.Get("rule").(*schema.Set); nrs.Len() > 0 {
		// Create an empty rule set to hold all newly created rules
		rules := resourceCiscoASAACL().Schema["rule"].ZeroValue().(*schema.Set)

		err := createCiscoASAACLRules(meta, d.Id(), rules, nrs)

		// We need to update this first to preserve the correct state
		d.Set("rule", rules)

		if err != nil {
			return err
		}
	}

	return resourceCiscoASAACLRead(d, meta)
}

func createCiscoASAACLRules(meta interface{}, acl string, rules *schema.Set, nrs *schema.Set) error {
	ca := meta.(*ciscoasa.Client)

	for _, rule := range nrs.List() {
		rule := rule.(map[string]interface{})

		id, err := ca.Objects.CreateExtendedACLACE(
			acl,
			cidrToAddress(rule["source"].(string)),
			rule["source_service"].(string),
			cidrToAddress(rule["destination"].(string)),
			rule["destination_service"].(string),
			rule["active"].(bool),
			rule["permit"].(bool),
		)
		if err != nil {
			return fmt.Errorf("Error creating ACE on ACL %s: %v", acl, err)
		}

		rule["id"] = id
		rules.Add(rule)
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

	// Create an empty schema.Set to hold all rules
	rules := resourceCiscoASAACL().Schema["rule"].ZeroValue().(*schema.Set)

	// Read all rules that are configured
	if rs := d.Get("rule").(*schema.Set); rs.Len() > 0 {
		for _, rule := range rs.List() {
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
			rules.Add(rule)
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
			rules.Add(rule)
		}
	}

	if rules.Len() > 0 {
		d.Set("rule", rules)
	} else if !managed {
		d.SetId("")
	}

	return nil
}

func resourceCiscoASAACLUpdate(d *schema.ResourceData, meta interface{}) error {
	// Check if the rule set as a whole has changed
	if d.HasChange("rule") {
		o, n := d.GetChange("rule")
		ors := o.(*schema.Set).Difference(n.(*schema.Set))
		nrs := n.(*schema.Set).Difference(o.(*schema.Set))

		// We need to start with a rule set containing all the rules we
		// already have and want to keep. Any rules that are not deleted
		// correctly and any newly created rules, will be added to this
		// set to make sure we end up in a consistent state
		rules := o.(*schema.Set).Intersection(n.(*schema.Set))

		// First loop through all the new rules and create them
		if nrs.Len() > 0 {
			err := createCiscoASAACLRules(meta, d.Id(), rules, nrs)

			// We need to update this first to preserve the correct state
			d.Set("rule", rules)

			if err != nil {
				return err
			}
		}

		// Then loop through all the old rules and remove them
		if ors.Len() > 0 {
			err := deleteCiscoASAACLRules(meta, d.Id(), rules, ors)

			// We need to update this first to preserve the correct state
			d.Set("rule", rules)

			if err != nil {
				return err
			}
		}
	}

	return resourceCiscoASAACLRead(d, meta)
}

func resourceCiscoASAACLDelete(d *schema.ResourceData, meta interface{}) error {
	// Delete all rules
	if ors := d.Get("rule").(*schema.Set); ors.Len() > 0 {
		// Create an additional set with all the existing rules. Each rule that is
		// succesfully deleted will be removed from this set, leaving only rules that
		// could not be deleted properly and should be saved in the state.
		rules := d.Get("rule").(*schema.Set)

		err := deleteCiscoASAACLRules(meta, d.Id(), rules, ors)

		// We need to update this first to preserve the correct state
		d.Set("rule", rules)

		if err != nil {
			return err
		}
	}

	return nil
}

func deleteCiscoASAACLRules(meta interface{}, acl string, rules *schema.Set, ors *schema.Set) error {
	ca := meta.(*ciscoasa.Client)

	for _, rule := range ors.List() {
		rule := rule.(map[string]interface{})

		err := ca.Objects.DeleteExtendedACLACE(acl, rule["id"].(string))
		if err != nil {
			if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
				log.Printf(
					"[DEBUG] ACE %s from ACL %s no longer exists", rule["id"].(string), acl)
				continue
			}

			return fmt.Errorf("Error deleting ACE %s from ACL %s: %v", rule["id"].(string), acl, err)
		}

		rules.Remove(rule)
	}

	return nil
}
