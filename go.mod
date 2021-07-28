module github.com/terraform-providers/terraform-provider-ciscoasa

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/xanzy/go-ciscoasa //This one is waiting for https://github.com/svanharmelen/go-ciscoasa/pull/12 to be merged
)
