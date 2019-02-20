---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_object_group"
sidebar_current: "docs-ciscoasa-resource-network-object-group"
description: |-
  Provides a Cisco ASA Network Object Group resource.
---

# github_issue_label

Provides a Cisco ASA Network Object Group.

## Example Usage

```hcl
resource "ciscoasa_network_object" "ipv4host" {
  name = "my_object"
  value = "192.168.10.5"
}

resource "ciscoasa_network_object_group" "objgrp_mixed" {
  name = "my_group"
  members = [
    "${ciscoasa_network_object.obj_ipv4host.name}",
    "192.168.10.15",
  	"10.5.10.0/24",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group.
* `members` - (Required) The list of the group members.
