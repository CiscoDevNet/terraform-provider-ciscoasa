---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_license_register"
sidebar_current: "docs-ciscoasa-resource-license-register"
description: |-
  Provides a Cisco ASA Register the Smart License resource.
---

# ciscoasa_license_register

Provides a Cisco ASA Register the Smart License resource.

## Example Usage

```hcl
resource "ciscoasa_license_register" "test" {
  id_token = "ZDBmOTJjOWItMTk5NS00ODNhLThmZWUtMDQ0NjNkMjM4YzlmLTE0NDUwOTQw%0AMzUyNzN8T1ArdTVHaHpWeWcwaHRwMzhMaWRtaW9FTWxoNHRXc3RTOGt3Tk1V%0AZ0JLMD0%3D%0A"
}
```

## Argument Reference

The following arguments are supported:

* `id_token` - (Required) Token ID of the virtual account to which this ASAv will be assigned.
* `force` - (Optional) Use the force keyword to register an ASAv that is already registered, but that might be out of sync with the License Authority.