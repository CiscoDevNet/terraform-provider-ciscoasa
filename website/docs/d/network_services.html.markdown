---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_services"
sidebar_current: "docs-ciscoasa-datasource-network-services"
description: |-
  Provides a Cisco ASA Network Services Data Source.
---

# ciscoasa_network_services

Provides a Cisco ASA Network Services Data Source.

## Example Usage

```hcl
data "ciscoasa_network_services" "all" {}
```

## Attributes Exported

The following attribute is exported:

* `network_services` - The configuration is detailed below

### `network_services` block contains:

* `name` - The name of the service
* `value` - The value representing the service
