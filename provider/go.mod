module github.com/pulumi/pulumi-random/provider/v2

go 1.14

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.4.1-0.20200608011815-6feeb51f2d39
	github.com/pulumi/pulumi/sdk/v2 v2.5.1-0.20200630091945-bb358c4d2173
	github.com/terraform-providers/terraform-provider-random v0.0.0-20190925210718-83518d96ae4f
)

replace github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible

replace github.com/pulumi/pulumi-terraform-bridge/v2 => ../../pulumi-terraform-bridge
