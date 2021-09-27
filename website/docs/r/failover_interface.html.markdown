---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_failover_interface"
sidebar_current: "docs-ciscoasa-resource-failover-interface"
description: |-
  Provides a Cisco ASA Failover Interface resource.
---

# ciscoasa_failover_interface

Provides a Cisco ASA Failover Interface resource.

## Example Usage

```hcl
resource "ciscoasa_failover_interface" "inside" {
  hardware_id = ciscoasa_interface_physical.inside.hardware_id
  standby_ip  = ciscoasa_network_object.inside_host.value
  monitored   = false
}
```

## Argument Reference

The following arguments are supported:

* `hardware_id` - (Required) Failover Interface Hardware ID.
* `standby_ip` - (Required) Standby IP Address for the failover Interface.
* `monitored` - (Optional) If it is monitored.