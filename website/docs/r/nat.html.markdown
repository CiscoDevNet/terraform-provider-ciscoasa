---
layout: "ciscoasa"
page_title: "Cisco ASA: ciscoasa_nat"
sidebar_current: "docs-ciscoasa-resource-nat"
description: |-
  Provides a Cisco ASA NAT resource.
---

# ciscoasa_nat

Provides a Cisco ASA Failover Setup resource.

## Example Usage

```hcl
resource "ciscoasa_nat" "auto_test" {
  section                   = "auto"
  description               = "static auto test"
  mode                      = "static"
  original_interface_name   = ciscoasa_interface_physical.inside.name
  translated_interface_name = ciscoasa_interface_physical.outside.name
  original_source_kind      = "objectRef#NetworkObj"
  original_source_value     = ciscoasa_network_object.inside_host.name
  translated_source_kind    = "objectRef#NetworkObj"
  translated_source_value   = ciscoasa_network_object.inside_host_translated.name
}

resource "ciscoasa_nat" "after_test1" {
  active                     = false
  section                    = "after"
  description                = "dynamic after test 1"
  mode                       = "dynamic"
  position                   = 1
  original_interface_name    = ciscoasa_interface_physical.inside.name
  translated_interface_name  = ciscoasa_interface_physical.outside.name
  original_source_kind       = "objectRef#NetworkObj"
  original_source_value      = ciscoasa_network_object.inside_host.name
  original_destination_kind  = "objectRef#NetworkObj"
  original_destination_value = ciscoasa_network_object.outside_pool.name
  original_service_kind      = "objectRef#TcpUdpServiceObj"
  original_service_value     = ciscoasa_network_service.tcp_port.name
  translated_source_kind     = "objectRef#NetworkObj"
  translated_source_value    = ciscoasa_network_object.host_101.name
  translated_service_kind    = "objectRef#TcpUdpServiceObj"
  translated_service_value   = ciscoasa_network_service.tcp_dns.name
}
```

## Argument Reference

The following arguments are supported:

* `section` - (Required) Section of the NAT rules ("auto", "before", "after")
* `active` - (Optional) Indicates if the rule is enabled
* `block_allocation` - (Optional) 
* `description` - (Optional) Description, remarks for the rule
* `extended` - (Optional) Enables extension of PAT uniqueness to per destination instead of per interface for PAT pool
* `flat` - (Optional) Enables use of entire 1024 to 65535 port range
* `include_reserve` - (Optional) If true, it uses all possible port numbers including ports 1 to 1023
* `dns` - (Optional) If true, translates DNS replies that match this rule, applicable only when translated service is not entered
* `interface_pat` - (Optional) Set this value to true, to enable Fall through to Interface PAT. Not applicable in transparent mode
* `net_to_net` - (Optional) If true, one-to-one address translation is enabled
* `no_proxy_arp` - (Optional) If true, enables proxy ARP for incoming packets to the mapped IP address
* `pat_pool` - (Optional) Enable PAT Pool for this rule. Applicable only for dynamic mode
* `round_robin` - (Optional) If true, it enables round robin addresses for a PAT pool. Applicable only for dynamic mode
* `route_lookup` - (Optional) For identity NAT in routed mode, determines the egress interface using a route lookup instead of using the interface specified in the NAT command
* `unidirectional` - (Optional) If true, makes the translation unidirectional, default is bidirectional
* `mode` - (Required) Source NAT type. Allowed values = 'static', 'dynamic'
* `position` - (Optional) Position/line number of the rule ( value > 0). Applicable only if more than one rule in section
* `use_destination_interface_ipv6` - (Optional) If true, IPV6 address of the mapped interface will be used for translated destination
* `use_source_interface_ipv6` - (Optional) If true, IPV6 address of the real interface will be used for translated source
* `original_destination` - (Optional) Original destination network object reference
* `original_interface_name` - (Optional) Real interface object reference
* `original_service` - (Optional) Original service object reference
* `original_source` - (Optional) Original source network object reference
* `translate_destination` - (Optional) Mapped/translated destination network object reference
* `translated_interface_name` - (Optional) Mapped/translated interface object reference
* `translated_service` - (Optional) Mapped/translated service object reference
* `translated_source` - (Optional) Mapped/translated source network object reference
* `translated_source_pat_pool` - (Optional) PAT Pool network address for mapped source