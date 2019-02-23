---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_service_group"
sidebar_current: "docs-ciscoasa-resource-network-service-group"
description: |-
  Provides a Cisco ASA Network Service Group resource.
---

# github_issue_label

Provides a Cisco ASA Network Service Group.

## Example Usage

```hcl
resource "ciscoasa_network_service_group" "service_group" {
  name = "service_group"
  
  members = [
    "tcp/80",
    "udp/53",
    "tcp/6001-6500",
    "icmp/0",
  ]
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the group.
* `members` - (Required) The list of the group members.
