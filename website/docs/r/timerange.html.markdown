---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_access_timerange"
sidebar_current: "docs-ciscoasa-resource-timerange"
description: |-
  Provides Cisco ASA Time Range Resource.
---

# ciscoasa_timerange

Provides Cisco ASA time range resource.

## Example Usage

```hcl
resource "ciscoasa_timerange" "tr" {
  name = "tr"
  value {
    start = "now"
    end   = "03:47 May 14 2025"
    periodic {
      frequency    = "Wednesday to Thursday"
      start_hour   = 4
      start_minute = 3
      end_hour     = 23
      end_minute   = 59
    }
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required)
* `value` - (Required) One `value` element as defined below.

### `value` supports the following:

* `start` - (Required) Time at which this one-time time-range starts.
* `end` - (Required) Time at which this one-time time-range ends.
* `periodic` - (Optional) One or more `periodic` element as defined below.

### `periodic` supports the following:

* `frequency` - (Required)
* `start_hour` - (Required)
* `start_minute` - (Required)
* `end_hour` - (Required)
* `end_minute` - (Required)
