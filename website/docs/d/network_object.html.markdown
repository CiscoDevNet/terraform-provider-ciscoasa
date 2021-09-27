---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_object"
sidebar_current: "docs-ciscoasa-datasource-network-object"
description: |-
  Provides a Cisco ASA Network Object Data Source.
---

# ciscoasa_network_object

Provides a Cisco ASA Network Object Data Source.

## Example Usage

```hcl
data "ciscoasa_network_object" "obj_ipv4subnet" {
  name = "test_ipv4subnet"
}
```

## Argument Reference

The following argument is required:

* `name` - The name of the object

## Attributes Exported

The following attributes are exported:

* `name` - The name of the object
* `value` - The value representing the object
