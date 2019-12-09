package ciscoasa

import (
	"strings"

	"net"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// Provider returns a terraform.ResourceProvider.
func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_API_URL", nil),
			},

			"username": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_USERNAME", nil),
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_PASSWORD", nil),
			},

			"ssl_no_verify": &schema.Schema{
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_SSLNOVERIFY", false),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ciscoasa_access_in_rules":       resourceCiscoASAAccessInRules(),
			"ciscoasa_access_out_rules":      resourceCiscoASAAccessOutRules(),
			"ciscoasa_acl":                   resourceCiscoASAACL(),
			"ciscoasa_network_object":        resourceCiscoASANetworkObject(),
			"ciscoasa_network_object_group":  resourceCiscoASANetworkObjectGroup(),
			"ciscoasa_network_service_group": resourceCiscoASANetworkServiceGroup(),
			"ciscoasa_static_route":          resourceCiscoASAStaticRoute(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := Config{
		APIURL:      d.Get("api_url").(string),
		Username:    d.Get("username").(string),
		Password:    d.Get("password").(string),
		SSLNoVerify: d.Get("ssl_no_verify").(bool),
	}

	return config.NewClient()
}

// cidrToAddress handles reserved cidr notations which within
// acl's are not allowed and need to be provided as reserved key words
func cidrToAddress(s string) string {
	switch strings.ToLower(s) {
	case "0.0.0.0/0":
		return "any4"
	case "::/0":
		return "any6"
	}

	return trimNetworkPrefix(s)
}

func addressToCIDR(s string) string {
	switch strings.ToLower(s) {
	case "any":
		return "0.0.0.0/0"
	case "any4":
		return "0.0.0.0/0"
	case "any6":
		return "::/0"
	}

	return addNetworkPrefix(s)
}

func trimNetworkPrefix(s string) string {
	addr, _, err := net.ParseCIDR(s)
	if err == nil {
		if addr.To4() != nil {
			return strings.TrimSuffix(s, "/32")
		}
		if addr.To16() != nil {
			return strings.TrimSuffix(s, "/128")
		}
	}

	return s
}

func addNetworkPrefix(s string) string {
	if !strings.Contains(s, "/") {
		addr, _, err := net.ParseCIDR(s + "/32")
		if err == nil {
			if addr.To4() != nil {
				return s + "/32"
			}
			if addr.To16() != nil {
				return s + "/128"
			}
		}
	}

	return s
}
