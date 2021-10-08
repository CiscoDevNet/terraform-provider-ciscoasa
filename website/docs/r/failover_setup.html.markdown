---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_failover_setup"
sidebar_current: "docs-ciscoasa-resource-failover-setup"
description: |-
  Provides a Cisco ASA Failover Setup resource.
---

# ciscoasa_failover_setup

Provides a Cisco ASA Failover Setup resource.

## Example Usage

```hcl
resource "ciscoasa_failover_setup" "test" {
  enable                           = true
  lan_interface_hw_id              = "TenGigabitEthernet0/0"
  lan_failover_name                = "test-fo"
  lan_primary_ip                   = "192.168.20.11"
  lan_secondary_ip                 = "192.168.20.12"
  lan_net_mask                     = "255.255.255.0"
  lan_preferred_role               = "secondary"
  state_interface_hw_id            = "TenGigabitEthernet0/0"
  state_failover_name              = "test-fo"
  state_primary_ip                 = "192.168.20.11"
  state_secondary_ip               = "192.168.20.12"
  state_net_mask                   = "255.255.255.0"
  failed_interfaces_threshold      = "10"
  failed_interfaces_threshold_unit = "Percentage"
}
```

## Argument Reference

The following arguments are supported:

* `enable` - (Optional) If failover is enabled.
* `shared_key` - (Optional) Failover key.
* `ipsec_key` - (Optional) IP Secret key.
* `hex_key` - (Optional) If using 32 hexadecimal character key.
* `state_failover_name` - (Optional) 
* `http_replication` - (Optional) Is HTTP replication enabled.
* `replication_rate` - (Optional) Replication Rate(Connections per second).
* `failed_interfaces_threshold_unit` - (Optional) Unit of failed Interfaces that triggers failover: Number or Percentage.
* `failed_interfaces_threshold` - (Optional) The value of failed Interfaces that triggers failover.
* `unit_poll_time` - (Optional) Unit Failover for Failover Poll Times.
* `unit_poll_time_unit` - (Optional) Unit of Unit Failover for Failover Poll Times: seconds or milliseconds.
* `unit_hold_time` - (Optional) Unit Hold Time for Failover Poll Times.
* `unit_hold_time_unit` - (Optional) Unit of Unit Hold Time for Failover Poll Times: seconds or milliseconds.
* `monitored_poll_time` - (Optional) Monitored Interfaces for Failover Poll Times.
* `monitored_poll_time_unit` - (Optional) Unit of Monitored Interfaces for Failover Poll Times: seconds or milliseconds.
* `interface_hold_time` - (Optional) Interface hold time in seconds. Range 5-75 and at least 5 times interface poll time.
* `lan_interface_hw_id` - (Required) Interface for LAN failover.
* `lan_failover_name` - (Required) Logical name for LAN failover interface.
* `lan_primary_ip` - (Required) Active IP for LAN failover.
* `lan_secondary_ip` - (Required) Standby IP for LAN failover.
* `lan_net_mask` - (Required) Subnet Mask for LAN failover.
* `lan_preferred_role` - (Required) Lan failover interface preferred role.
* `state_interface_hw_id` - (Optional) Interface for State Failover.
* `state_failover_name` - (Optional) Logical name for State Failover.
* `state_primary_ip` - (Optional) Active IP for State Failover.
* `state_secondary_ip` - (Optional) Standby IP for State Failover.
* `state_net_mask` - (Optional) Subnet Mask for State Failover.
