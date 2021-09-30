---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_write_memory"
sidebar_current: "docs-ciscoasa-resource-write-memory"
description: |-
  Provides a resource, which allows save in-memory changes for Cisco ASA.
---

# ciscoasa_write_memory

Provides a resource, which allows save in-memory changes for Cisco ASA. 

Write memory ASA operation is triggerred on resource creation. The ```triggers``` argument allows specifying an arbitrary set of values that, when changed, will cause the resource to be replaced.
It's recommended to use ```triggers``` argument in conjunction with ```jsonencode``` terraform function to trigger resource recreation (and write memory ASA operation) each time other ciscoasa resources are updated.   

## Example Usage

```hcl
resource "ciscoasa_static_route" "ipv4_static_route" {
  interface = "inside"
  network   = "10.254.0.0/16"
  gateway   = "192.168.10.20"
}

resource "ciscoasa_write_memory" "write_memory" { 
    triggers = {
        ciscoasa_static_route.ipv4_static_route = jsonencode(ciscoasa_static_route.ipv4_static_route)
    }
}
```

## Argument Reference

The following arguments are supported:

* `triggers` - (Optional) A map of arbitrary strings that, when changed, will force the null resource to be replaced, re-running any associated provisioners.
