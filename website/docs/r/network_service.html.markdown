---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_network_service"
sidebar_current: "docs-ciscoasa-resource-network-service"
description: |-
  Provides a Cisco ASA Network Service Object resource.
---

# ciscoasa_network_service

Provides a Cisco ASA Network Service Object resource.

## Example Usage

```hcl
resource "ciscoasa_network_service" "tcp-port" {
  name  = "tcp-port"
  value = "tcp/800"
}
resource "ciscoasa_network_service" "tcp-with-source" {
  name  = "tcp-with-source"
  value = "tcp/https,source=7000-8000"
}
resource "ciscoasa_network_service" "icmp4-with-type-code" {
  name  = "icmp4-with-type-code"
  value = "icmp/50/60"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) The name of the service.
* `value` - (Required) The value representing the object - complete network service data in text format (protocol, port information etc.).
