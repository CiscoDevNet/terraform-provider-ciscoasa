---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_service"
sidebar_current: "docs-ciscoasa-datasource-network-service"
description: |-
  Provides a Cisco ASA Network Service Data Source.
---

# ciscoasa_network_service

Provides a Cisco ASA Network Service Data Source.

## Example Usage

```hcl
data "ciscoasa_network_service" "tcp-port" {
  name = "tcp-port"
}
```

## Argument Reference

The following argument is required:

* `name` - The name of the service

## Attributes Exported

The following attributes are exported:

* `name` - The name of the service
* `value` - The value representing the service
