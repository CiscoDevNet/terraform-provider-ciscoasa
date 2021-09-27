package ciscoasa

import (
	"fmt"
	"reflect"
	"strings"

	"net"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/xanzy/go-ciscoasa/ciscoasa"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_API_URL", nil),
			},

			"username": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_USERNAME", nil),
			},

			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_PASSWORD", nil),
			},

			"ssl_no_verify": {
				Type:        schema.TypeBool,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("CISCOASA_SSLNOVERIFY", false),
			},
		},

		ResourcesMap: map[string]*schema.Resource{
			"ciscoasa_access_in_rules":           resourceCiscoASAAccessInRules(),
			"ciscoasa_access_out_rules":          resourceCiscoASAAccessOutRules(),
			"ciscoasa_interface_physical":        resourceCiscoASAPhysicalInterface(),
			"ciscoasa_interface_vlan":            resourceCiscoASAVlanInterface(),
			"ciscoasa_acl":                       resourceCiscoASAACL(),
			"ciscoasa_network_object":            resourceCiscoASANetworkObject(),
			"ciscoasa_network_object_group":      resourceCiscoASANetworkObjectGroup(),
			"ciscoasa_network_service":           resourceCiscoASANetworkService(),
			"ciscoasa_network_service_group":     resourceCiscoASANetworkServiceGroup(),
			"ciscoasa_static_route":              resourceCiscoASAStaticRoute(),
			"ciscoasa_timerange":                 resourceCiscoASATimeRange(),
			"ciscoasa_ntp_server":                resourceCiscoASANtpServer(),
			"ciscoasa_dhcp_server":               resourceCiscoASADhcpServer(),
			"ciscoasa_dhcp_relay_globalsettings": resourceCiscoASADhcpRelayGlobalsettings(),
			"ciscoasa_dhcp_relay_local":          resourceCiscoASADhcpRelayLocal(),
			"ciscoasa_backup":                    resourceCiscoASABackup(),
			"ciscoasa_nat":                       resourceCiscoASANat(),
			"ciscoasa_failover_interface":        resourceCiscoASAFailoverInterface(),
			"ciscoasa_failover_setup":            resourceCiscoASAFailoverSetup(),
			"ciscoasa_license_config":            resourceCiscoASALicenseConfig(),
			"ciscoasa_license_register":          resourceCiscoASALicenseRegister(),
			"ciscoasa_license_renewauth":         resourceCiscoASALicenseRenewAuth(),
			"ciscoasa_license_renewid":           resourceCiscoASALicenseRenewId(),
		},

		DataSourcesMap: map[string]*schema.Resource{
			"ciscoasa_interfaces_physical": dataSourceCiscoASAPhysicalInterfaces(),
			"ciscoasa_interface_physical":  dataSourceCiscoASAPhysicalInterface(),
			"ciscoasa_interface_vlan":      dataSourceCiscoASAVlanInterface(),
			"ciscoasa_interfaces_vlan":     dataSourceCiscoASAVlanInterfaces(),
			"ciscoasa_network_object":      dataSourceCiscoASANetworkObject(),
			"ciscoasa_network_objects":     dataSourceCiscoASANetworkObjects(),
			"ciscoasa_network_service":     dataSourceCiscoASANetworkService(),
			"ciscoasa_network_services":    dataSourceCiscoASANetworkServices(),
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

func expandIPAddress(d *schema.ResourceData) (*ciscoasa.IPAddress, error) {
	ipAddressSet := d.Get("ip_address").(*schema.Set).List()

	result := &ciscoasa.IPAddress{}

	err := fmt.Errorf("Error in 'ip_address' parameters")

	attrs, ok := ipAddressSet[0].(map[string]interface{})
	if ok {
		if reflect.ValueOf(attrs["static"]).Len() > 0 {
			staticAttrs := attrs["static"].([]interface{})[0].(map[string]interface{})
			result.Kind = "StaticIP"
			ip := &ciscoasa.Address{}
			ip.Kind = "IPv4Address"
			ip.Value = staticAttrs["ip"].(string)
			result.IP = ip
			netMask := &ciscoasa.Address{}
			netMask.Kind = "IPv4NetMask"
			netMask.Value = staticAttrs["net_mask"].(string)
			result.NetMask = netMask
		}
		if reflect.ValueOf(attrs["dhcp"]).Len() > 0 {
			dhcpAttrs := attrs["dhcp"].([]interface{})[0].(map[string]interface{})
			result.Kind = "DHCP"
			result.DhcpOptionUsingMac = dhcpAttrs["dhcp_option_using_mac"].(bool)
			result.DhcpBroadcast = dhcpAttrs["dhcp_broadcast"].(bool)
			dhcpClient := &ciscoasa.DhcpClient{}
			dhcpClientAttrs := dhcpAttrs["dhcp_client"].([]interface{})[0].(map[string]interface{})
			dhcpClient.SetDefaultRoute = dhcpClientAttrs["set_default_route"].(bool)
			dhcpClient.Metric = dhcpClientAttrs["metric"].(int)
			dhcpClient.PrimaryTrackId = dhcpClientAttrs["primary_track_id"].(int)
			dhcpClient.TrackingEnabled = dhcpClientAttrs["tracking_enabled"].(bool)
			if dhcpClient.TrackingEnabled {
				slaTracking := &ciscoasa.SlaTracking{}
				slaTrackingAttrs := dhcpClientAttrs["sla_tracking_settings"].([]interface{})[0].(map[string]interface{})
				slaTracking.SlaId = slaTrackingAttrs["sla_id"].(int)
				slaTracking.TrackedIP = slaTrackingAttrs["tracked_ip"].(string)
				slaTracking.FrequencyInSeconds = slaTrackingAttrs["frequency_in_seconds"].(int)
				slaTracking.DataSizeInBytes = slaTrackingAttrs["data_size_in_bytes"].(int)
				slaTracking.ThresholdInMilliseconds = slaTrackingAttrs["threshold_in_milliseconds"].(int)
				slaTracking.ToS = slaTrackingAttrs["tos"].(int)
				slaTracking.TimeoutInMilliseconds = slaTrackingAttrs["timeout_in_milliseconds"].(int)
				slaTracking.NumPackets = slaTrackingAttrs["num_packets"].(int)
				dhcpClient.SlaTrackingSettings = slaTracking
			}
			result.DhcpClient = dhcpClient
		}
	} else {
		return nil, err
	}

	return result, nil
}

func expandIPv6Info(d *schema.ResourceData) (*ciscoasa.IPv6Info, error) {
	ipV6InfoSet := d.Get("ipv6_info").(*schema.Set).List()

	result := &ciscoasa.IPv6Info{}

	err := fmt.Errorf("Error in 'ipv6_info' parameters")

	attrs, ok := ipV6InfoSet[0].(map[string]interface{})
	if ok {
		result.Enabled = attrs["enabled"].(bool)
		result.AutoConfig = attrs["auto_config"].(bool)
		result.EnforceEUI64 = attrs["enforce_eui64"].(bool)
		result.ManagedAddressConfig = attrs["managed_address_config"].(bool)
		result.NsInterval = attrs["ns_interval"].(int)
		result.DadAttempts = attrs["dad_attempts"].(int)
		if reflect.ValueOf(attrs["n_discovery_prefix_list"]).Len() > 0 {
			v := make([]*ciscoasa.NDiscoveryPrefix, 0)
			nDiscoveryPrefix := &ciscoasa.NDiscoveryPrefix{}
			nDiscoveryPrefixAttrs := attrs["n_discovery_prefix_list"].([]interface{})[0].(map[string]interface{})
			nDiscoveryPrefix.Kind = "IPV6Prefix"
			nDiscoveryPrefix.OffLink = nDiscoveryPrefixAttrs["off_link"].(bool)
			nDiscoveryPrefix.NoAdvertise = nDiscoveryPrefixAttrs["no_advertise"].(bool)
			nDiscoveryPrefix.PreferredLifetime = nDiscoveryPrefixAttrs["preferred_lifetime"].(int)
			nDiscoveryPrefix.ValidLifetime = nDiscoveryPrefixAttrs["valid_lifetime"].(int)
			nDiscoveryPrefix.HasDuration = nDiscoveryPrefixAttrs["has_duration"].(bool)
			nDiscoveryPrefix.DefaultPrefix = nDiscoveryPrefixAttrs["default_prefix"].(bool)
			v = append(v, nDiscoveryPrefix)
			result.NDiscoveryPrefixList = v
		}
		result.OtherStatefulConfig = attrs["other_stateful_config"].(bool)
		result.RouterAdvertInterval = attrs["router_advert_interval"].(int)
		result.RouterAdvertIntervalUnit = attrs["router_advert_interval_unit"].(string)
		result.RouterAdvertLifetime = attrs["router_advert_lifetime"].(int)
		result.SuppressRouterAdvert = attrs["suppress_router_advert"].(bool)
		result.ReachableTime = attrs["reachable_time"].(int)
		if reflect.ValueOf(attrs["link_local_address"]).Len() > 0 {
			llAddress := &ciscoasa.Ipv6Address{}
			llAddressAttrs := attrs["link_local_address"].([]interface{})[0].(map[string]interface{})
			llAddress.Kind = "object#Ipv6InterfaceAddress"
			// llAddress.IsEUI64 = llAddressAttrs["is_eui64"].(bool)
			address := &ciscoasa.Address{}
			address.Kind = "IPv6Address"
			address.Value = llAddressAttrs["address"].(string)
			llAddress.Address = address
			standby := &ciscoasa.Address{}
			standby.Kind = "IPv6Address"
			standby.Value = llAddressAttrs["standby"].(string)
			llAddress.Standby = standby
			result.LinkLocalAddress = llAddress
		}
		if reflect.ValueOf(attrs["ipv6_addresses"]).Len() > 0 {
			v := make([]*ciscoasa.Ipv6Address, 0)
			ipv6Address := &ciscoasa.Ipv6Address{}
			ipv6AddressAttrs := attrs["ipv6_addresses"].([]interface{})[0].(map[string]interface{})
			ipv6Address.Kind = "object#Ipv6InterfaceAddress"
			// ipv6Address.IsEUI64 = ipv6AddressAttrs["is_eui64"].(bool)
			address := &ciscoasa.Address{}
			address.Kind = "IPv6Address"
			address.Value = ipv6AddressAttrs["address"].(string)
			ipv6Address.Address = address
			standby := &ciscoasa.Address{}
			standby.Kind = "IPv6Address"
			standby.Value = ipv6AddressAttrs["standby"].(string)
			ipv6Address.Standby = standby
			ipv6Address.PrefixLength = ipv6AddressAttrs["prefix_length"].(int)
			v = append(v, ipv6Address)
			result.Ipv6Addresses = v
		}
		result.Kind = "object#Ipv6InterfaceInfo"
	} else {
		return nil, err
	}

	return result, nil
}

func flattenIPAddress(in *ciscoasa.IPAddress) []map[string]interface{} {
	var out = make([]map[string]interface{}, 0, 0)

	m := make(map[string]interface{})

	if in.Kind == "DHCP" {
		dhcpNested := make([]map[string]interface{}, 0, 0)
		dhcp := make(map[string]interface{})
		dhcp["dhcp_option_using_mac"] = in.DhcpOptionUsingMac
		dhcp["dhcp_broadcast"] = in.DhcpBroadcast
		dhcpClientNested := make([]map[string]interface{}, 0, 0)
		dhcpClient := make(map[string]interface{})
		dhcpClient["set_default_route"] = in.DhcpClient.SetDefaultRoute
		dhcpClient["metric"] = in.DhcpClient.Metric
		dhcpClient["primary_track_id"] = in.DhcpClient.PrimaryTrackId
		dhcpClient["tracking_enabled"] = in.DhcpClient.TrackingEnabled
		if in.DhcpClient.TrackingEnabled {
			slaTrackingNested := make([]map[string]interface{}, 0, 0)
			slaTracking := make(map[string]interface{})
			slaTracking["sla_id"] = in.DhcpClient.SlaTrackingSettings.SlaId
			slaTracking["tracked_ip"] = in.DhcpClient.SlaTrackingSettings.TrackedIP
			slaTracking["frequency_in_seconds"] = in.DhcpClient.SlaTrackingSettings.FrequencyInSeconds
			slaTracking["data_size_in_bytes"] = in.DhcpClient.SlaTrackingSettings.DataSizeInBytes
			slaTracking["threshold_in_milliseconds"] = in.DhcpClient.SlaTrackingSettings.ThresholdInMilliseconds
			slaTracking["tos"] = in.DhcpClient.SlaTrackingSettings.ToS
			slaTracking["timeout_in_milliseconds"] = in.DhcpClient.SlaTrackingSettings.TimeoutInMilliseconds
			slaTracking["num_packets"] = in.DhcpClient.SlaTrackingSettings.NumPackets
			slaTrackingNested = append(slaTrackingNested, slaTracking)
			dhcpClient["sla_tracking_settings"] = slaTrackingNested
		}
		dhcpClientNested = append(dhcpClientNested, dhcpClient)
		dhcp["dhcp_client"] = dhcpClientNested
		dhcpNested = append(dhcpNested, dhcp)
		m["dhcp"] = dhcpNested
	} else {
		staticNested := make([]map[string]interface{}, 0, 0)
		ip := make(map[string]interface{})
		ip["ip"] = in.IP.Value
		ip["net_mask"] = in.NetMask.Value
		staticNested = append(staticNested, ip)
		m["static"] = staticNested
	}

	out = append(out, m)

	return out
}

func flattenIPv6Info(in *ciscoasa.IPv6Info) []map[string]interface{} {

	var out = make([]map[string]interface{}, 0, 0)

	m := make(map[string]interface{})

	m["auto_config"] = in.AutoConfig
	m["dad_attempts"] = in.DadAttempts
	m["enabled"] = in.Enabled
	m["enforce_eui64"] = in.EnforceEUI64
	if in.LinkLocalAddress != nil {
		llNested := make([]map[string]interface{}, 0, 0)
		ll := make(map[string]interface{})
		// ll["is_eui64"] = in.LinkLocalAddress.IsEUI64
		ll["address"] = in.LinkLocalAddress.Address.Value
		ll["standby"] = in.LinkLocalAddress.Standby.Value
		llNested = append(llNested, ll)
		m["link_local_address"] = llNested
	}
	if reflect.ValueOf(in.Ipv6Addresses).Len() > 0 {
		ipv6Nested := make([]map[string]interface{}, 0, 0)
		ipv6 := make(map[string]interface{})
		// ipv6["is_eui64"] = in.Ipv6Addresses[0].IsEUI64
		ipv6["address"] = in.Ipv6Addresses[0].Address.Value
		ipv6["standby"] = in.Ipv6Addresses[0].Standby.Value
		ipv6["prefix_length"] = in.Ipv6Addresses[0].PrefixLength
		ipv6Nested = append(ipv6Nested, ipv6)
		m["ipv6_addresses"] = ipv6Nested
	}
	m["managed_address_config"] = in.ManagedAddressConfig
	if reflect.ValueOf(in.NDiscoveryPrefixList).Len() > 0 {
		ndplNested := make([]map[string]interface{}, 0, 0)
		ndpl := make(map[string]interface{})
		ndpl["off_link"] = in.NDiscoveryPrefixList[0].OffLink
		ndpl["no_advertise"] = in.NDiscoveryPrefixList[0].NoAdvertise
		ndpl["preferred_lifetime"] = in.NDiscoveryPrefixList[0].PreferredLifetime
		ndpl["valid_lifetime"] = in.NDiscoveryPrefixList[0].ValidLifetime
		ndpl["has_duration"] = in.NDiscoveryPrefixList[0].HasDuration
		ndpl["default_prefix"] = in.NDiscoveryPrefixList[0].DefaultPrefix
		ndplNested = append(ndplNested, ndpl)
		m["n_discovery_prefix_list"] = ndplNested
	}
	m["ns_interval"] = in.NsInterval
	m["other_stateful_config"] = in.OtherStatefulConfig
	m["reachable_time"] = in.ReachableTime
	m["router_advert_interval"] = in.RouterAdvertInterval
	m["router_advert_interval_unit"] = in.RouterAdvertIntervalUnit
	m["router_advert_lifetime"] = in.RouterAdvertLifetime
	m["suppress_router_advert"] = in.SuppressRouterAdvert

	out = append(out, m)

	return out
}

func isInterfaceHwId(iface string) bool {
	return strings.Contains(iface, "/")
}

func objectIdFromHwId(iface string) string {
	return strings.Replace(iface, "/", "_API_SLASH_", 1)
}

func hwIdFromObjId(id string) string {
	return strings.Replace(id, "_API_SLASH_", "/", 1)
}
