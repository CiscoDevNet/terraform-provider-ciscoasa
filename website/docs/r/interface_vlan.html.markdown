---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_interface_vlan"
sidebar_current: "docs-ciscoasa-resource-interface-vlan"
description: |-
  Provides a Cisco ASA Interface Vlan resource.
---

# ciscoasa_interface_vlan

Provides a Cisco ASA interface vlan resource.

## Example Usage

```hcl
resource "ciscoasa_interface_vlan" "tengig_no_config" {
  name           = "vlantengig140"
  hardware_id    = "TenGigabitEthernet0/0.140"
  interface_desc = "VlanTenGig Zero Conf"
  vlan_id        = 141
  security_level = 0
}

resource "ciscoasa_interface_vlan" "tengig_ipv4_static" {
  name           = "vlantengig150"
  hardware_id    = "TenGigabitEthernet0/1.150"
  interface_desc = "VlanTenGig Static"
  ip_address {
    static {
      ip       = "10.0.2.6"
      net_mask = "255.255.255.0"
    }
  }
  security_level = 0
}

resource "ciscoasa_interface_vlan" "tengig_ipv4_dhcp_sla" {
  name           = "vlantengig160"
  hardware_id    = "TenGigabitEthernet0/1.160"
  interface_desc = "VlanTenGig DHCP Sla tracking"

  ip_address {
    dhcp {
      dhcp_option_using_mac = false
      dhcp_broadcast        = true
      dhcp_client {
        set_default_route = false
        metric            = 1
        primary_track_id  = 6
        tracking_enabled  = true
        sla_tracking_settings {
          sla_id                    = 8
          tracked_ip                = "10.0.2.6"
          frequency_in_seconds      = 61
          data_size_in_bytes        = 30
          threshold_in_milliseconds = 5001
          tos                       = 2
          timeout_in_milliseconds   = 5002
          num_packets               = 2
        }
      }
    }
  }
  security_level = 0
}

resource "ciscoasa_interface_vlan" "tengig_ipv6" {
  name           = "vlantengig170"
  hardware_id    = "TenGigabitEthernet0/2.170"
  interface_desc = "VlanTenGig Ipv6 with standby address"
  ipv6_info {
    auto_config   = false
    dad_attempts  = 1
    enabled       = true
    enforce_eui64 = false
    link_local_address {
      address = "fe80::20e:cff:fe3b:883c"
      standby = "fe80::20e:cff"
    }
    ipv6_addresses {
      address       = "2001:db8:a0b:12f0::47"
      standby       = "2001:db8:a0b:12f0::46"
      prefix_length = 64
    }
    managed_address_config = true
    n_discovery_prefix_list {
      off_link           = false
      no_advertise       = true
      preferred_lifetime = 604800
      valid_lifetime     = 2592000
      has_duration       = true
      default_prefix     = true
    }
    ns_interval                 = 1000
    other_stateful_config       = true
    reachable_time              = 0
    router_advert_interval      = 250
    router_advert_interval_unit = "sec"
    router_advert_lifetime      = 1800
    suppress_router_advert      = true
  }
  security_level = 0
}
```

## Argument Reference

The following arguments are supported:

* `active_mac_address` - (Optional) Active MAC Address
* `forward_traffic_cx` - (Optional) Enable traffic forwarding to CX module
* `forward_traffic_sfr` - (Optional) Enable traffic forwarding to FirePOWER module
* `hardware_id` - (Required) Vlan interface ID (hardware ID)
* `interface_desc` - (Optional) Interface description; up to 240 characters
* `ip_address` - (Optional) One `ip_address` element as defined below
* `ipv6_info` - (Optional) One `ipv6_info` element as defined below
* `management_only` - (Optional) Designate this interface as management-only
* `mtu` - (Optional) Specify MTU; 300 to 65535 bytes
* `name` - (Optional) To provide a name for an interface
* `security_level` - (Required) The interface security level: enter a value between 0 (lowest) and 100 (highest)
* `shutdown` - (Optional) Disable this interface
* `stand_by_mac_address` - (Optional) Standby MAC Address
* `vlan_id` - (Optional) Provide a VLAN ID for the subinterface

### `ip_address` supports the following:

* `static` - (Optional) One `static` element as defined below
* `dhcp` - (Optional) One `dhcp` element as defined below

### `static` supports the following:

* `ip` - (Required) IPv4 address
* `net_mask` - (Required) NetMask

### `dhcp` supports the following:

* `dhcp_option_using_mac` - (Required)
* `dhcp_broadcast` - (Required)
* `dhcp_client` - (Optional) One `dhcp_client` element as defined below

### `dhcp_client` supports the following:

* `set_default_route` - (Required)
* `metric` - (Required)
* `primary_track_id` - (Required)
* `tracking_enabled` - (Required)
* `sla_tracking_settings` - (Optional) One `sla_tracking_settings` element as defined below if `tracking_enabled` is set to `true`

### `sla_tracking_settings` supports the following:

* `sla_id` - (Required)
* `tracked_ip` - (Required)
* `frequency_in_seconds` - (Required)
* `data_size_in_bytes` - (Required)
* `threshold_in_milliseconds` - (Required)
* `tos` - (Required)
* `timeout_in_milliseconds` - (Required)
* `num_packets` - (Required)

### `ipv6_info` supports the following:

* `auto_config` - (Required) IPv6 Auto Configuration enabled
* `dad_attempts` - (Required) Neighbor solicitation message count (DAD attempts). Range is 0 to 600. Default is 1
* `enabled` - (Required) IPv6 enabled
* `enforce_eui64` - (Required) Enforce Extended Unique Identifier (EUI), as per RFC2373
* `link_local_address` - (Optional) One `link_local_address` element as defined below
* `ipv6_addresses` - (Optional) One `ipv6_addresses` element as defined below
* `managed_address_config` - (Required) Managed address configuration flag
* `ipv6_addresses` - (Optional) One `ipv6_addresses` element as defined below
* `ns_interval` - (Required) Neighbor solicitation retransmit interval. Range is from 1000 msec to 3600000 msec. Default is 1000
* `other_stateful_config` - (Required) Other stateful configuration flag
* `reachable_time` - (Required) Reachable time. Range is from 0 sec to 3600000 msec. Default is 0
* `router_advert_interval` - (Required) Router advertisement interval. Range is from 3 sec to 1800 sec (or from 1000 msec to 3600000 msec). Default is 200 sec
* `router_advert_interval_unit` - (Required) Router advertisement interval unit, 'sec' or 'msec'. Default is 'sec'
* `router_advert_lifetime` - (Required) Router advertisement lifetime. Range is 0 seconds to 9000 seconds. Default is 1800 seconds
* `suppress_router_advert` - (Required) Suppress router advertisement messages on this IPv6 interface

### `link_local_address` supports the following:

* `address` - (Required) IPv6 address
* `standby` - (Required) IPv6 standby address

### `ipv6_addresses` supports the following:

* `address` - (Required) IPv6 address
* `standby` - (Required) IPv6 standby address
* `prefix_length` - (Required) IPv6 prefix length

### `n_discovery_prefix_list` supports the following:

* `off_link` - (Required)
* `no_advertise` - (Required)
* `preferred_lifetime` - (Required)
* `valid_lifetime` - (Required)
* `has_duration` - (Required)
* `default_prefix` - (Required)