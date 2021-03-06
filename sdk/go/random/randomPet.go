// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"reflect"

	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The resource `.RandomPet` generates random pet names that are intended to be
// used as unique identifiers for other resources.
//
// This resource can be used in conjunction with resources that have
// the `createBeforeDestroy` lifecycle flag set, to avoid conflicts with
// unique names during the brief period where both the old and new resources
// exist concurrently.
type RandomPet struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The length (in words) of the pet name.
	Length pulumi.IntPtrOutput `pulumi:"length"`
	// A string to prefix the name with.
	Prefix pulumi.StringPtrOutput `pulumi:"prefix"`
	// The character to separate words in the pet name.
	Separator pulumi.StringPtrOutput `pulumi:"separator"`
}

// NewRandomPet registers a new resource with the given unique name, arguments, and options.
func NewRandomPet(ctx *pulumi.Context,
	name string, args *RandomPetArgs, opts ...pulumi.ResourceOption) (*RandomPet, error) {
	if args == nil {
		args = &RandomPetArgs{}
	}
	var resource RandomPet
	err := ctx.RegisterResource("random:index/randomPet:RandomPet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomPet gets an existing RandomPet resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomPet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomPetState, opts ...pulumi.ResourceOption) (*RandomPet, error) {
	var resource RandomPet
	err := ctx.ReadResource("random:index/randomPet:RandomPet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomPet resources.
type randomPetState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length (in words) of the pet name.
	Length *int `pulumi:"length"`
	// A string to prefix the name with.
	Prefix *string `pulumi:"prefix"`
	// The character to separate words in the pet name.
	Separator *string `pulumi:"separator"`
}

type RandomPetState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers pulumi.MapInput
	// The length (in words) of the pet name.
	Length pulumi.IntPtrInput
	// A string to prefix the name with.
	Prefix pulumi.StringPtrInput
	// The character to separate words in the pet name.
	Separator pulumi.StringPtrInput
}

func (RandomPetState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPetState)(nil)).Elem()
}

type randomPetArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The length (in words) of the pet name.
	Length *int `pulumi:"length"`
	// A string to prefix the name with.
	Prefix *string `pulumi:"prefix"`
	// The character to separate words in the pet name.
	Separator *string `pulumi:"separator"`
}

// The set of arguments for constructing a RandomPet resource.
type RandomPetArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated. See
	// the main provider documentation for more information.
	Keepers pulumi.MapInput
	// The length (in words) of the pet name.
	Length pulumi.IntPtrInput
	// A string to prefix the name with.
	Prefix pulumi.StringPtrInput
	// The character to separate words in the pet name.
	Separator pulumi.StringPtrInput
}

func (RandomPetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPetArgs)(nil)).Elem()
}
