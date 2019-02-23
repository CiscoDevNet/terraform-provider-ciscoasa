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

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
provider "ciscoasa" {
  api_url       = "https://10.0.0.5"
  username      = "admin"
  password      = # your password here
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
