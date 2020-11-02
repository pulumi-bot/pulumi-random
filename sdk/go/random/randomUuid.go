// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The resource `RandomUuid` generates random uuid string that is intended to be
// used as unique identifiers for other resources.
//
// This resource uses the `hashicorp/go-uuid` to generate a UUID-formatted string
// for use with services needed a unique string identifier.
//
// ## Example Usage
//
// The following example shows how to generate a unique name for an Azure Resource Group.
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-azure/sdk/v3/go/azure/core"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		_, err := random.NewRandomUuid(ctx, "testRandomUuid", nil)
// 		if err != nil {
// 			return err
// 		}
// 		_, err = core.NewResourceGroup(ctx, "testResourceGroup", &core.ResourceGroupArgs{
// 			Location: pulumi.String("Central US"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// ## Import
//
// Random UUID's can be imported. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example
type RandomUuid struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will
	// trigger a new uuid to be generated.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The generated uuid presented in string format.
	Result pulumi.StringOutput `pulumi:"result"`
}

// NewRandomUuid registers a new resource with the given unique name, arguments, and options.
func NewRandomUuid(ctx *pulumi.Context,
	name string, args *RandomUuidArgs, opts ...pulumi.ResourceOption) (*RandomUuid, error) {
	if args == nil {
		args = &RandomUuidArgs{}
	}
	var resource RandomUuid
	err := ctx.RegisterResource("random:index/randomUuid:RandomUuid", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomUuid gets an existing RandomUuid resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomUuid(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomUuidState, opts ...pulumi.ResourceOption) (*RandomUuid, error) {
	var resource RandomUuid
	err := ctx.ReadResource("random:index/randomUuid:RandomUuid", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomUuid resources.
type randomUuidState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new uuid to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The generated uuid presented in string format.
	Result *string `pulumi:"result"`
}

type RandomUuidState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new uuid to be generated.
	Keepers pulumi.MapInput
	// The generated uuid presented in string format.
	Result pulumi.StringPtrInput
}

func (RandomUuidState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomUuidState)(nil)).Elem()
}

type randomUuidArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new uuid to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
}

// The set of arguments for constructing a RandomUuid resource.
type RandomUuidArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new uuid to be generated.
	Keepers pulumi.MapInput
}

func (RandomUuidArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomUuidArgs)(nil)).Elem()
}
