---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_static_route"
sidebar_current: "docs-ciscoasa-resource-static-route"
description: |-
  Provides a Cisco ASA Static Route resource.
---

# github_issue_label

Provides a Cisco ASA static route resource.

## Example Usage

```hcl
resource "ciscoasa_static_route" "ipv4_static_route" {
  "interface" = "inside"
  "network" = "10.254.0.0/16"
  "gateway" = "192.168.10.20"
}

resource "ciscoasa_static_route" "ipv6_static_route" {
  "interface" = "inside"
  "network" = "fd01:1337::/64"
  "gateway" = "fd01:1338::1"
}
```

## Argument Reference

The following arguments are supported:

* `interface` - (Required) The name of the interface.
* `network` - (Required)
* `gateway` - (Required)
* `metric` - (Optional) Default `1`.
* `tracked` - (Optional) Default `false`.
* `tunneled` - (Optional) Default `false`.
