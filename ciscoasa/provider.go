package ciscoasa

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
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
