---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_backup"
sidebar_current: "docs-ciscoasa-resource-backup"
description: |-
  Provides a Cisco ASA Backup Configuration resource.
---

# ciscoasa_backup

Provides a Cisco ASA Backup Configuration resource.

## Example Usage

```hcl
resource "ciscoasa_backup" "test" {
  passphrase = "123456"
  location   = "disk0:/backup.cfg"
}
```

## Argument Reference

The following arguments are supported:

* `context` - (Optional) Context to backup.
* `location` - (Optional) Path of backup file. Default `""`.
* `passphrase` - (Optional) Passphrase to encrypt certificates and passwords.
