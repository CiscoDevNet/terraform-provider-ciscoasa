---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_dhcp_server"
sidebar_current: "docs-ciscoasa-resource-dhcp-server"
description: |-
  Provides a Cisco ASA DHCP Server resource.
---

# ciscoasa_dhcp_server

Provides a Cisco ASA DHCP Server resource.

## Example Usage

```hcl
resource "ciscoasa_interface_physical" "ipv4_interface" {
  name           = "test"
  hardware_id    = "TenGigabitEthernet0/1"
  interface_desc = "test descr"
  ip_address {
    static {
      ip       = "8.8.8.1"
      net_mask = "255.255.255.0"
    }
  }
  security_level = 0
}

resource "ciscoasa_dhcp_server" "dhcp_test" {
  interface             = ciscoasa_interface_physical.ipv4_interface.name
  enabled               = true
  pool_start_ip         = "8.8.8.4"
  pool_end_ip           = "8.8.8.20"
  dns_ip_primary        = "3.3.3.3"
  dns_ip_secondary      = "5.5.5.5"
  wins_ip_primary       = "4.4.4.4"
  wins_ip_secondary     = "6.6.6.6"
  lease_length          = "305"
  ping_timeout          = "40"
  domain_name           = "testing1"
  auto_config_enabled   = true
  auto_config_interface = false
  vpn_override          = false
  options {
    type   = "hex"
    code   = 2
    value1 = "c52f"
  }
  options {
    type   = "ascii"
    code   = 4
    value1 = "1261"
  }
  options {
    type   = "ip"
    code   = 13
    value1 = "1.1.1.2"
    value2 = "1.1.2.1"
  }
  ddns_update_dns_client        = true
  ddns_update_both_records      = true
  ddns_override_client_settings = true
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) Name of the interface for DHCP server.
* `enabled` - (Required) Enable the DHCP server.
* `pool_start_ip` - (Required) DHCP address pool start IP addresses.
* `pool_end_ip` - (Required) DHCP address pool end IP addresses.
* `dns_ip_primary` - (Optional) IP addresses of the primary DNS server.
* `dns_ip_secondary` - (Optional) IP addresses of the secondary DNS server.
* `wins_ip_primary` - (Optional) IP addresses of the primary WINS server.
* `wins_ip_secondary` - (Optional) IP addresses of the secondary WINS server.
* `lease_length` - (Optional) Duration of time that the DHCP server configured on the interface allows DHCP clients to use an assigned IP address.
* `ping_timeout` - (Optional) Time in milliseconds that the ASA will wait for an ICMP ping response on the interface.
* `domain_name` - (Optional) DNS domain name.
* `auto_config_enabled` - (Optional) Enable DHCP auto configuration only if the ASA is acting as a DHCP client on a specified interface (usually outside).
* `auto_config_interface` - (Optional) Name of the interface on a DHCP client that provides DNS, WINS, and domain name information for automatic configuration.
* `vpn_override` - (Optional) When auto-configuration is enabled, override the interface DHCP or PPPoE client WINS parameter with the VPN client parameter.
* `options` - (Optional) One or more `options` element as defined below.
* `ddns_update_dns_client` - (Optional) Configure custom DNS client update actions, in addition to the default action of updating the client PTR resource records.
* `ddns_update_both_records` - (Optional) DHCP server should update both the PTR and the A resource records.
* `ddns_override_client_settings` - (Optional) DHCP server actions should override any update actions requested by the DHCP client.

### `options` supports the following:

* `type` - (Required) Option type (one of "ascii", "hex" and "ip").
* `code` - (Required) Option code. All DHCP options (1 through 255) are supported except for 1, 12, 50-54, 58-59, 61, 67, and 82.
* `value1` - (Required) Option value.
* `value2` - (Optional) Secondary option value (applicable to the "ip" type only).
