---
layout: "ciscoasa"
page_title: "Provider: Cisco ASA"
sidebar_current: "docs-ciscoasa-index"
description: |-
  The Cisco ASA provider is used to interact with Cisco ASA hardware devices or the Cisco ASAv virtual appliance.
  The provider needs to be configured with the proper credentials before it can be used.
---

# Cisco ASA Provider

The Cisco ASA provider is used to interact with Cisco ASA hardware devices or the Cisco ASAv virtual appliance. The provider needs to be configured with the proper credentials before it can be used.

Use the navigation to the left to read about the available Resources and Data Sources.

To learn the basics of Terraform using this provider, follow the hands-on [get started tutorials](https://learn.hashicorp.com/tutorials/terraform/infrastructure-as-code?in=terraform/azure-get-started&_ga=2.39967714.1393978341.1633676960-942932012.1624353028) on HashiCorp's Learn platform.

## Prerequisites

Before using the provider, Cisco ASA Rest API agent should be installed and enabled on the ASA. [This document](https://www.cisco.com/c/en/us/td/docs/security/asa/api/qsg-asa-api.html) can be used for reference.


## Example Usage

```hcl
provider "ciscoasa" {
  api_url       = "<IP Address of Cisco ASA>"
  username      = "<Username>"
  password      = "<Password>"
  ssl_no_verify = false
}
```

## Configuration Reference

The following keys can be used to configure the provider.

* `api_url` - (Required) URL of the API for the ASA Firewall. This is typically not enabled by default, please refer to the [Cisco documentation](https://www.cisco.com/c/en/us/td/docs/security/asa/api/qsg-asa-api.html) for how to enable it.

  This can also be set as the `CISCOASA_API_URL` environment variable.

* `username` - (Required) The username for logging in to the API.

  This can also be set as the `CISCOASA_USERNAME` environment variable.

* `password` - (Required) The password for logging in to the API.

  This can also be set as the `CISCOASA_PASSWORD` environment variable.

* `ssl_no_verify` - (Required) A flag indicating whether or not to verify the TLS certificate.

  This can also be set as the `CISCOASA_SSLNOVERIFY` environment variable.

## Issues

For Cisco ASAv on Microsoft Azure, the Physical and VLAN interface Resources creation and update might get stuck. Users running ASAv on Microsoft Azure, can create the interfaces manually and reference the interfaces as data source for other resources.

## Feature Enhancement and Filing Bugs

Bugs files and feature ehancement requests can be found in the GitHub repo issues section of Ciscoasa Terraform Provider. 
Please avoid "me too" or "+1" comments. Instead, use a thumbs up reaction on enhancement requests. Provider maintainers will often prioritize work based on the number of thumbs on an issue.

Community input is appreciated on outstanding issues! We would also love to hear about new use cases that can be addressed using the Cisco ASA Provider.

If you're interested in working on an issue please leave a comment on that issue
