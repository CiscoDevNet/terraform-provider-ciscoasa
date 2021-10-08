---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_license_register"
sidebar_current: "docs-ciscoasa-resource-license-register"
description: |-
  Provides a Cisco ASA Smart License Registration resource.
---

# ciscoasa_license_register

Provides a Cisco ASA Smart License Registration resource.
terraform destroy command will deregister the Cisco ASA's smart license

## Example Usage

```hcl
resource "ciscoasa_license_register" "test" {
  id_token = "<registration token>"
}
```

## Argument Reference

The following arguments are supported:

* `id_token` - (Required) Token ID of the virtual account to which this ASAv will be assigned.
* `force` - (Optional) Use the force keyword to register an ASAv that is already registered, but that might be out of sync with the License Authority.
