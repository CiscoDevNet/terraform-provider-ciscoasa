---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_acl"
sidebar_current: "docs-ciscoasa-resource-acl"
description: |-
  Provides a Cisco ASA ACL resource.
---

# ciscoasa_acl

Provides a Cisco ASA ACL resource.

## Example Usage

```hcl
resource "ciscoasa_acl" "foo" {
  name = "aclname"
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

* `name` - (Required) The name of the ACL.
* `rule` - (Required) One or more `rule` elements as defined below.

### `rule` supports the following:

* `destination_service` - (Required)
* `destination` - (Required)
* `source` - (Required)
* `active` - (Optional) Default `true`.
* `log_interval` - (Optional) Default `300`.
* `log_status` - (Optional) Must be one of `Default`, `Debugging`, `Disabled`, `Notifications`, `Critical`, `Emergencies`, `Warnings`, `Errors`, `Informational`, `Alerts`. Default `Default`.
* `permit` - (Optional) Default `true`.
* `remarks` - (Optional)
* `source_service` - (Optional)
* `id` - (Computed)
