---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_dhcp_relay"
sidebar_current: "docs-ciscoasa-resource-dhcp-relay-local"
description: |-
  Provides a Cisco ASA DHCP Relay Interface Server resource.
---

# ciscoasa_dhcp_relay_local

Provides a Cisco ASA DHCP Relay Interface Server resource.

## Example Usage

```hcl
resource "ciscoasa_interface_physical" "tengig_ipv4_dhcp" {
  name           = "tengig"
  hardware_id    = "TenGigabitEthernet0/2"
  interface_desc = "Interface DHCP"
  ip_address {
    dhcp {
      dhcp_option_using_mac = false
      dhcp_broadcast        = true
      dhcp_client {
        set_default_route = true
        metric            = 1
        primary_track_id  = -1
        tracking_enabled  = false
      }
    }
  }
  security_level = 0
}


resource "ciscoasa_dhcp_relay_local" "dhcp_relay_test" {
  interface = ciscoasa_interface_physical.tengig_ipv4_dhcp.name
  servers = [
    "166.177.180.190",
    "20.20.30.26",
    "135.144.153.163"
  ]
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) Name of the interface connected to the DHCP network.
* `servers` - (Required) IP addresses of DHCP servers to be used for DHCP requests that enter that interface.
