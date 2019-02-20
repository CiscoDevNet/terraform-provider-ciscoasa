---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_object"
sidebar_current: "docs-ciscoasa-resource-network-object"
description: |-
  Provides a Cisco ASA Network Object resource.
---

# github_issue_label

Provides a Cisco ASA Network Object resource.

## Example Usage

```hcl
resource "ciscoasa_network_object" "ipv4host" {
  name = "ipv4_host"
  value = "192.168.10.5"
}
resource "ciscoasa_network_object" "ipv4range" {
  name = "ipv4_range"
  value = "192.168.10.5-192.168.10.15"
}
resource "ciscoasa_network_object" "ipv4_subnet" {
  name = "ipv4_subnet"
  value = "192.168.10.128/25"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group.
* `value` - (Required) The value representing the object. This can be a single host, a range of hosts (`<ip>-<ip>`), or a CIDR.
