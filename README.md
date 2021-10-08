Terraform Provider
==================

> **Note:** this Terraform provider is now publically available on the [Terraform Registry](https://registry.terraform.io/providers/CiscoDevNet/ciscoasa/latest).

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 1.0.x
-	[Go](https://golang.org/doc/install) 1.9 (to build the provider plugin)

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/CiscoDevNet/terraform-provider-ciscoasa`

```sh
$ mkdir -p $GOPATH/src/github.com/terraform-providers; cd $GOPATH/src/github.com/terraform-providers
$ git clone git@github.com:CiscoDevNet/terraform-provider-ciscoasa
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/CiscoDevNet/terraform-provider-ciscoasa
$ make build
```

Using the provider
----------------------
If you're building the provider, follow the instructions to
[install it as a plugin.](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin)
After placing it into your plugins directory,  run `terraform init` to initialize it.

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed
on your machine (version 1.9+ is *required*). You'll also need to correctly setup a
[GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary
in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-ciscoasa
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the Acceptance tests:
1. Prepare environment with Cisco ASA instance and Rest API installed and configured (check `testinfra` folder for Terraform configs example) 
2. Check/export necessary Environment Variables
3. Run `make testinfra-testacc`.

The list of current available Environment Variables:

|EnvVar|Default value|Usage|
|---|---|---|
|CISCOASA_SSLNOVERIFY|`true`|Skip SSL verification on connection to Cisco ASA API|
|CISCOASA_OBJECT_PREFIX|`acc`|Prefix for objects created by tests|
|CISCOASA_INTERFACE_NAME|`inside`|Named interface which will be used in tests|
|CISCOASA_INTERFACE_HW_ID_BASE|`TenGigabitEthernet0`|Base part of Interface Hardware ID|
|CISCOASA_INTERFACE_HW_IDS|`1,2`|Interfaces Hardware ID indexes|
|CISCOASA_USERNAME|`$(cd testinfra; terraform output asav_username)`|Username for Cisco ASA API|
|CISCOASA_PASSWORD|`$(cd testinfra; terraform output asav_password)`|Password for Cisco ASA API|
|CISCOASA_API_URL|`https://$(cd testinfra; terraform output asav_public_ip)`|URL for Cisco ASA API|

```sh
$ make testinfra-testacc
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```
