---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_interface_vlan"
sidebar_current: "docs-ciscoasa-datasource-interface-vlan"
description: |-
  Provides a Cisco ASA VLAN Interface Data Source.
---

# ciscoasa_interface_vlan

Provides a Cisco ASA VLAN Interface Data Source.

## Example Usage

```hcl
data "ciscoasa_interface_vlan" "vlantengig140" {
  hardware_id = "TenGigabitEthernet0/0.140"
}
```

## Argument Reference

The following argument is required:

* `hardware_id` - Vlan interface ID (hardware ID)

## Attributes Exported

The following attributes are exported:

* `active_mac_address` - Active MAC Address
* `forward_traffic_cx` - Enable traffic forwarding to CX module
* `forward_traffic_sfr` - Enable traffic forwarding to FirePOWER module
* `hardware_id` - Physical interface ID (hardware ID)
* `interface_desc` - Interface description
* `ip_address` - IP Address option for the interface. The configuration is detailed below
* `ipv6_info` - IPv6 Information. The configuration is detailed below
* `management_only` - Designate this interface as management-only
* `mtu` - MTU
* `name` - To provide a name for an interface
* `security_level` - The interface security level
* `shutdown` - Disable this interface
* `stand_by_mac_address` - Standby MAC Address
* `vlan_id` - Provide a VLAN ID for the subinterface

### `ip_address` block contains:

* `static` - The configuration is detailed below
* `dhcp` - The configuration is detailed below

### `static` block contains:

* `ip` - IPv4 address
* `net_mask` - NetMask

### `dhcp` block contains:

* `dhcp_option_using_mac`
* `dhcp_broadcast`
* `dhcp_client` - The configuration is detailed below

### `dhcp_client` block contains:

* `set_default_route`
* `metric`
* `primary_track_id`
* `tracking_enabled`
* `sla_tracking_settings` - The configuration is detailed below

### `sla_tracking_settings` block contains:

* `sla_id`
* `tracked_ip`
* `frequency_in_seconds`
* `data_size_in_bytes`
* `threshold_in_milliseconds`
* `tos`
* `timeout_in_milliseconds`
* `num_packets`

### `ipv6_info` block contains:

* `auto_config` - IPv6 Auto Configuration enabled
* `dad_attempts` - Neighbor solicitation message count (DAD attempts)
* `enabled` - IPv6 enabled
* `enforce_eui64` - Enforce Extended Unique Identifier (EUI)
* `link_local_address` - The configuration is detailed below
* `ipv6_addresses` - The configuration is detailed below
* `managed_address_config` - Managed address configuration flag
* `ns_interval` - Neighbor solicitation retransmit interval
* `other_stateful_config` - Other stateful configuration flag
* `reachable_time` - Reachable time
* `router_advert_interval` - Router advertisement interval
* `router_advert_interval_unit` - Router advertisement interval unit
* `router_advert_lifetime` - Router advertisement lifetime
* `suppress_router_advert` - Suppress router advertisement messages on this IPv6 interface

### `link_local_address` block contains:

* `address` - IPv6 address
* `standby` - IPv6 standby address

### `ipv6_addresses` block contains:

* `address` - IPv6 address
* `standby` - IPv6 standby address
* `prefix_length` - IPv6 prefix length

### `n_discovery_prefix_list` block contains:

* `off_link`
* `no_advertise`
* `preferred_lifetime`
* `valid_lifetime`
* `has_duration`
* `default_prefix`
