package ciscoasa

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/CiscoDevNet/go-ciscoasa/ciscoasa"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCiscoASANat() *schema.Resource {
	return &schema.Resource{
		Create: resourceCiscoASANatCreate,
		Read:   resourceCiscoASANatRead,
		Update: resourceCiscoASANatUpdate,
		Delete: resourceCiscoASANatDelete,

		Schema: map[string]*schema.Schema{

			"section": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"after", "auto", "before",
				}, false),
				ForceNew: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"original_interface_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"translated_interface_name": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"original_source_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AnyIPAddress",
					"IPv4Range",
					"IPv6Range",
					"IPv4Network",
					"IPv6Network",
					"IPv4Address",
					"IPv6Address",
					"objectRef#NetworkObj",
					"objectRef#NetworkObjGroup",
				}, false),
				Default: "AnyIPAddress",
			},

			"original_source_value": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "any",
			},

			"original_destination_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AnyIPAddress",
					"IPv4Range",
					"IPv6Range",
					"IPv4Network",
					"IPv6Network",
					"IPv4Address",
					"IPv6Address",
					"objectRef#NetworkObj",
					"objectRef#NetworkObjGroup",
				}, false),
				Default: "AnyIPAddress",
			},

			"original_destination_value": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "any",
			},

			"original_service_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"AnyService",
					"objectRef#ICMP6ServiceObj",
					"objectRef#ICMPServiceObj",
					"objectRef#NetworkProtocolObj",
					"objectRef#TcpUdpServiceObj",
				}, false),
				Default: "AnyService",
			},

			"original_service_value": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "any",
			},

			"translated_source_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"IPv4Range",
					"IPv6Range",
					"IPv4Network",
					"IPv6Network",
					"IPv4Address",
					"IPv6Address",
					"objectRef#NetworkObj",
					"objectRef#NetworkObjGroup",
				}, false),
			},

			"translated_source_value": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"translated_destination_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"IPv4Range",
					"IPv6Range",
					"IPv4Network",
					"IPv6Network",
					"IPv4Address",
					"IPv6Address",
					"objectRef#NetworkObj",
					"objectRef#NetworkObjGroup",
				}, false),
			},

			"translated_destination_value": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"translated_service_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"objectRef#ICMP6ServiceObj",
					"objectRef#ICMPServiceObj",
					"objectRef#NetworkProtocolObj",
					"objectRef#TcpUdpServiceObj",
				}, false),
			},

			"translated_service_value": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"mode": {
				Type:     schema.TypeString,
				Required: true,
			},

			"interface_pat": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"use_source_interface_ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"pat_pool": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"translated_source_pat_pool_kind": {
				Type:     schema.TypeString,
				Optional: true,
				ValidateFunc: validation.StringInSlice([]string{
					"IPv4Range",
					"IPv6Range",
					"IPv4Network",
					"IPv6Network",
					"IPv4Address",
					"IPv6Address",
					"objectRef#NetworkObj",
					"objectRef#NetworkObjGroup",
				}, false),
			},

			"translated_source_pat_pool_value": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"extended": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"flat": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"include_reserve": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"block_allocation": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"round_robin": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"use_destination_interface_ipv6": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"net_to_net": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"dns": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"unidirectional": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"no_proxy_arp": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"route_lookup": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"active": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
				DiffSuppressFunc: func(k, old, new string, d *schema.ResourceData) bool {
					if d.Get("section").(string) == "auto" {
						return true
					}
					return false
				},
			},

			"position": {
				Type:         schema.TypeInt,
				Optional:     true,
				Computed:     true,
				ValidateFunc: validation.IntAtLeast(1),
			},

			"last_updated": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceCiscoASANatCreate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	section := d.Get("section").(string)

	Nat := ciscoasa.Nat{
		Active:                      d.Get("active").(bool),
		BlockAllocation:             d.Get("block_allocation").(bool),
		Description:                 d.Get("description").(string),
		Extended:                    d.Get("extended").(bool),
		Flat:                        d.Get("flat").(bool),
		IncludeReserve:              d.Get("include_reserve").(bool),
		IsDNS:                       d.Get("dns").(bool),
		IsInterfacePAT:              d.Get("interface_pat").(bool),
		IsNetToNet:                  d.Get("net_to_net").(bool),
		IsNoProxyArp:                d.Get("no_proxy_arp").(bool),
		IsPatPool:                   d.Get("pat_pool").(bool),
		IsRoundRobin:                d.Get("round_robin").(bool),
		IsRouteLookup:               d.Get("route_lookup").(bool),
		IsUnidirectional:            d.Get("unidirectional").(bool),
		Mode:                        d.Get("mode").(string),
		Position:                    d.Get("position").(int),
		UseDestinationInterfaceIPv6: d.Get("use_destination_interface_ipv6").(bool),
		UseSourceInterfaceIPv6:      d.Get("use_source_interface_ipv6").(bool),
	}

	Nat.OriginalDestination = objFromKindValue(d, "original_destination")
	if originalInterfaceName, ok := d.GetOk("original_interface_name"); ok {
		Nat.OriginalInterface = objFromInterfaceName(originalInterfaceName)
	}
	Nat.OriginalService = objFromKindValue(d, "original_service")
	Nat.OriginalSource = objFromKindValue(d, "original_source")
	Nat.TranslatedDestination = objFromKindValue(d, "translated_destination")
	if translatedInterfaceName, ok := d.GetOk("translated_interface_name"); ok {
		Nat.TranslatedInterface = objFromInterfaceName(translatedInterfaceName)
	}
	Nat.TranslatedService = objFromKindValue(d, "translated_service")
	Nat.TranslatedSource = objFromKindValue(d, "translated_source")
	Nat.TranslatedSourcePatPool = objFromKindValue(d, "translated_source_pat_pool")

	NatId, err := ca.Nat.CreateNat(section, &Nat)

	if err != nil {
		return fmt.Errorf(
			"Error creating NAT %s: %v", section, err)
	}

	d.SetId(NatId)

	return resourceCiscoASANatRead(d, meta)
}

func resourceCiscoASANatRead(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	section := d.Get("section").(string)
	r, err := ca.Nat.GetNat(section, d.Id())
	if err != nil {
		if strings.Contains(err.Error(), "RESOURCE-NOT-FOUND") {
			log.Printf(
				"[DEBUG] NAT for %s->%s no longer exists", section, d.Id())
			d.SetId("")
			return nil
		}

		return fmt.Errorf(
			"Error reading NAT %s->%s: %v", section, d.Id(), err)
	}
	if r.Description != "" {
		d.Set("description", r.Description)
	}
	if r.OriginalInterface != nil {
		d.Set("original_interface_name", r.OriginalInterface.Name)
	}
	if r.TranslatedInterface != nil {
		d.Set("translated_interface_name", r.TranslatedInterface.Name)
	}
	if r.OriginalSource != nil {
		d.Set("original_source_kind", r.OriginalSource.Kind)
		d.Set("original_source_value", valueFromObj(r.OriginalSource))
	}
	if r.OriginalDestination != nil {
		d.Set("original_destination_kind", r.OriginalDestination.Kind)
		d.Set("original_destination_value", valueFromObj(r.OriginalDestination))
	}
	if r.OriginalService != nil {
		d.Set("original_service_kind", r.OriginalService.Kind)
		d.Set("original_service_value", valueFromObj(r.OriginalService))
	}
	if r.TranslatedSource != nil {
		d.Set("translated_source_kind", r.TranslatedSource.Kind)
		d.Set("translated_source_value", valueFromObj(r.TranslatedSource))
	}
	if r.TranslatedDestination != nil {
		d.Set("translated_destination_kind", r.TranslatedDestination.Kind)
		d.Set("translated_destination_value", valueFromObj(r.TranslatedDestination))
	}
	if r.TranslatedService != nil {
		d.Set("translated_service_kind", r.TranslatedService.Kind)
		d.Set("translated_service_value", valueFromObj(r.TranslatedService))
	}
	d.Set("mode", r.Mode)
	d.Set("interface_pat", r.IsInterfacePAT)
	d.Set("use_source_interface_ipv6", r.UseSourceInterfaceIPv6)
	d.Set("pat_pool", r.IsPatPool)
	if r.TranslatedSourcePatPool != nil {
		d.Set("translated_source_pat_pool_kind", r.TranslatedSourcePatPool.Kind)
		d.Set("translated_source_pat_pool_value", valueFromObj(r.TranslatedSourcePatPool))
	}
	d.Set("extended", r.Extended)
	d.Set("flat", r.Flat)
	d.Set("include_reserve", r.IncludeReserve)
	d.Set("block_allocation", r.BlockAllocation)
	d.Set("round_robin", r.IsRoundRobin)
	d.Set("use_destination_interface_ipv6", r.UseDestinationInterfaceIPv6)
	d.Set("net_to_net", r.IsNetToNet)
	d.Set("dns", r.IsDNS)
	d.Set("unidirectional", r.IsUnidirectional)
	d.Set("no_proxy_arp", r.IsNoProxyArp)
	d.Set("route_lookup", r.IsRouteLookup)
	if section != "auto" {
		d.Set("active", r.Active)
	}
	d.Set("position", r.Position)

	return nil
}

func resourceCiscoASANatUpdate(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	section := d.Get("section").(string)

	if d.HasChanges(
		"active",
		"block_allocation",
		"description",
		"extended",
		"flat",
		"include_reserve",
		"dns",
		"interface_pat",
		"net_to_net",
		"no_proxy_arp",
		"pat_pool",
		"round_robin",
		"route_lookup",
		"unidirectional",
		"mode",
		"position",
		"use_destination_interface_ipv6",
		"use_source_interface_ipv6",
		"original_destination",
		"original_interface_name",
		"original_service",
		"original_source",
		"translated_destination",
		"translated_interface_name",
		"translated_service",
		"translated_source",
		"translated_source_pat_pool",
	) {
		Nat := ciscoasa.Nat{
			Active:                      d.Get("active").(bool),
			BlockAllocation:             d.Get("block_allocation").(bool),
			Description:                 d.Get("description").(string),
			Extended:                    d.Get("extended").(bool),
			Flat:                        d.Get("flat").(bool),
			IncludeReserve:              d.Get("include_reserve").(bool),
			IsDNS:                       d.Get("dns").(bool),
			IsInterfacePAT:              d.Get("interface_pat").(bool),
			IsNetToNet:                  d.Get("net_to_net").(bool),
			IsNoProxyArp:                d.Get("no_proxy_arp").(bool),
			IsPatPool:                   d.Get("pat_pool").(bool),
			IsRoundRobin:                d.Get("round_robin").(bool),
			IsRouteLookup:               d.Get("route_lookup").(bool),
			IsUnidirectional:            d.Get("unidirectional").(bool),
			Mode:                        d.Get("mode").(string),
			Position:                    d.Get("position").(int),
			UseDestinationInterfaceIPv6: d.Get("use_destination_interface_ipv6").(bool),
			UseSourceInterfaceIPv6:      d.Get("use_source_interface_ipv6").(bool),
		}

		Nat.OriginalDestination = objFromKindValue(d, "original_destination")
		if originalInterfaceName, ok := d.GetOk("original_interface_name"); ok {
			Nat.OriginalInterface = objFromInterfaceName(originalInterfaceName)
		}
		Nat.OriginalService = objFromKindValue(d, "original_service")
		Nat.OriginalSource = objFromKindValue(d, "original_source")
		Nat.TranslatedDestination = objFromKindValue(d, "translated_destination")
		if translatedInterfaceName, ok := d.GetOk("translated_interface_name"); ok {
			Nat.TranslatedInterface = objFromInterfaceName(translatedInterfaceName)
		}
		Nat.TranslatedService = objFromKindValue(d, "translated_service")
		Nat.TranslatedSource = objFromKindValue(d, "translated_source")
		Nat.TranslatedSourcePatPool = objFromKindValue(d, "translated_source_pat_pool")

		NatId, err := ca.Nat.UpdateNat(section, d.Id(), &Nat)

		if err != nil {
			return fmt.Errorf(
				"Error updating NAT %s: %v", section, err)
		}

		d.SetId(NatId)
		d.Set("last_updated", time.Now().Format(time.RFC850))
	}

	return resourceCiscoASANatRead(d, meta)
}

func resourceCiscoASANatDelete(d *schema.ResourceData, meta interface{}) error {
	ca := meta.(*ciscoasa.Client)

	section := d.Get("section").(string)
	err := ca.Nat.DeleteNat(section, d.Id())
	if err != nil {
		return fmt.Errorf(
			"Error deleting NAT %s->%s: %v", section, d.Id(), err)
	}

	return nil
}

func isObjectRef(kind string) bool {
	return strings.Contains(kind, "objectRef#")
}

func objFromKindValue(d *schema.ResourceData, name string) *ciscoasa.TranslatedOriginalObj {
	if kind, ok := d.GetOk(fmt.Sprintf("%s_kind", name)); ok {
		Obj := &ciscoasa.TranslatedOriginalObj{}
		Obj.Kind = kind.(string)
		if value, ok := d.GetOk(fmt.Sprintf("%s_value", name)); ok {
			if isObjectRef(kind.(string)) {
				Obj.ObjectId = value.(string)
			} else {
				Obj.Value = value.(string)
			}
		}
		return Obj
	}
	return nil
}

func objFromInterfaceName(name interface{}) *ciscoasa.TranslatedOriginalInterface {
	Obj := &ciscoasa.TranslatedOriginalInterface{}

	Obj.Kind = "objectRef#Interface"
	Obj.Name = name.(string)

	return Obj
}

func valueFromObj(obj *ciscoasa.TranslatedOriginalObj) string {
	if isObjectRef(obj.Kind) {
		return obj.ObjectId
	}

	return obj.Value
}
