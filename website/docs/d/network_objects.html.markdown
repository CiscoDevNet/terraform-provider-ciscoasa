---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_objects"
sidebar_current: "docs-ciscoasa-datasource-network-objects"
description: |-
  Provides a Cisco ASA Network Objects Data Source.
---

# ciscoasa_network_objects

Provides a Cisco ASA Network Objects Data Source.

## Example Usage

```hcl
data "ciscoasa_network_objects" "all" {}
```

## Attributes Exported

The following attribute is exported:

* `network_objects` - The configuration is detailed below

### `network_objects` block contains:

* `name` - The name of the object
* `value` - The value representing the object
