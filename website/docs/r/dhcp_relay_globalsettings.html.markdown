---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_dhcp_relay_globalsettings"
sidebar_current: "docs-ciscoasa-resource-dhcp-relay-globalsettings"
description: |-
  Provides a Cisco ASA DHCP Relay Global Settings resource.
---

# ciscoasa_dhcp_relay_globalsettings

Provides a Cisco ASA DHCP Relay Global Settings resource.

## Example Usage

```hcl
resource "ciscoasa_dhcp_relay_globalsettings" "test" {
  ipv4_timeout              = 90
  ipv6_timeout              = 90
  trusted_on_all_interfaces = true
}
```

## Argument Reference

The following arguments are supported:

* `ipv4_timeout` - (Optional) Time in seconds (1 to 3600) allowed for IPv4 DHCP relay address handling. Default `60`.
* `ipv6_timeout` - (Optional) Time in seconds (1 to 3600) allowed for IPv6 DHCP relay address handling.  Default `60`.
* `trusted_on_all_interfaces` - (Optional) Configures all interfaces as trusted. Default `false`.
