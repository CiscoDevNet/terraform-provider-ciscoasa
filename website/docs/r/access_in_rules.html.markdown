---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_access_in_rules"
sidebar_current: "docs-ciscoasa-resource-access-in-rules"
description: |-
  Provides Cisco ASA Inbound Access Rules.
---

# ciscoasa_access_in_rules

Provides Cisco ASA inbound access rule. Inbound access rules apply to traffic as it enters an interface. Global access rules are always inbound.

## Example Usage

```hcl
resource "ciscoasa_access_in_rules" "foo" {
  interface = "inside"
  rule {
    source              = "192.168.10.5/32"
    destination         = "192.168.15.0/25"
    destination_service = "tcp/443"
  }
  rule {
    source              = "192.168.10.0/24"
		source_service      = "udp"
    destination         = "192.168.15.6/32"
    destination_service = "udp/53"
  }
  rule {
    source              = "192.168.10.0/23"
    destination         = "192.168.12.0/23"
    destination_service = "icmp/0"
  }
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required)
* `rule` - (Required) One or more `rule` elements as defined below.
* `managed` - (Optional) Default `false`.

### `rule` supports the following:

* `destination_service` - (Required)
* `destination` - (Required)
* `source` - (Required)
* `active` - (Optional) Default `true`.
* `permit` - (Optional) Default `true`.
* `source_service` - (Optional)
* `id` - (Computed)
