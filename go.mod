module github.com/terraform-providers/terraform-provider-ciscoasa

go 1.16

require (
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hil v0.0.0-20190212132231-97b3a9cdfa93 // indirect
	github.com/hashicorp/terraform v0.12.8 // indirect
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.7.0
	github.com/hashicorp/yamux v0.0.0-20181012175058-2f1d1f20f75d // indirect
	github.com/ulikunitz/xz v0.5.8 // indirect
	github.com/vmihailenco/msgpack v4.0.4+incompatible // indirect
	github.com/CiscoDevNet/go-ciscoasa v0.2.3
)

replace github.com/CiscoDevNet/go-ciscoasa v0.2.3 => github.com/id27182/go-ciscoasa v0.2.5
