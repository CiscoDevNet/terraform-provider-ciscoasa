---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_license_config"
sidebar_current: "docs-ciscoasa-resource-license-config"
description: |-
  Provides a Cisco ASA Smart License Config resource.
---

# ciscoasa_license_config

Provides a Cisco ASA Smart License Config resource.

## Example Usage

```hcl
resource "ciscoasa_license_config" "test" {
  throughput         = "2G"
}
```

## Argument Reference

The following arguments are supported:

* `license_server_url` - (Optional) The URL of the Licensing Authority. Unless directed by Cisco TAC, you should not change the License Authority URL.
* `transport_url` - (Optional)
* `privacy_host_name` - (Optional)
* `privacy_version` - (Optional)
* `throughput` - (Required) Throughtput level ("100M", "1G", "2G", "10G", "20G").
