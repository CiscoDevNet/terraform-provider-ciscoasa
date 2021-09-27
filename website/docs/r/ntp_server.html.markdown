---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_ntp_server"
sidebar_current: "docs-ciscoasa-resource-ntp-server"
description: |-
  Provides a Cisco ASA NTP Server resource.
---

# ciscoasa_ntp_server

Provides a Cisco ASA NTP Server resource.

## Example Usage

```hcl
resource "ciscoasa_ntp_server" "ntp_test" {
  ip_address = "2.2.2.2"
  interface  = "inside"
  key_number = "3"
  key_value  = "test3"
  preferred  = true
}
```

## Argument Reference

The following arguments are supported:

* `ip_address` - (Required) IP address of the server.
* `interface` - (Optional) Outgoing interface for NTP packets.
* `key_number` - (Required) Key ID for this authentication key.
* `key_value` - (Required) Key value for this authentication key.
* `key_trusted` - (Optional) Sets this authentication key as a trusted key, which is required for authentication to succeed.
* `preferred` - (Optional) Sets this server as a preferred server.
