// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The resource `RandomShuffle` generates a random permutation of a list
// of strings given as an argument.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v2/go/aws/elb"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		az, err := random.NewRandomShuffle(ctx, "az", &random.RandomShuffleArgs{
// 			Inputs: pulumi.StringArray{
// 				pulumi.String("us-west-1a"),
// 				pulumi.String("us-west-1c"),
// 				pulumi.String("us-west-1d"),
// 				pulumi.String("us-west-1e"),
// 			},
// 			ResultCount: pulumi.Int(2),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = elb.NewLoadBalancer(ctx, "example", &elb.LoadBalancerArgs{
// 			AvailabilityZones: az.Results,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type RandomShuffle struct {
	pulumi.CustomResourceState

	// The list of strings to shuffle.
	Inputs pulumi.StringArrayOutput `pulumi:"inputs"`
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The number of results to return. Defaults to
	// the number of items in the `input` list. If fewer items are requested,
	// some elements will be excluded from the result. If more items are requested,
	// items will be repeated in the result but not more frequently than the number
	// of items in the input list.
	ResultCount pulumi.IntPtrOutput `pulumi:"resultCount"`
	// Random permutation of the list of strings given in `input`.
	Results pulumi.StringArrayOutput `pulumi:"results"`
	// Arbitrary string with which to seed the random number
	// generator, in order to produce less-volatile permutations of the list.
	// **Important:** Even with an identical seed, it is not guaranteed that the
	// same permutation will be produced across different versions of the provider.
	// This argument causes the result to be *less volatile*, but not fixed for
	// all time.
	Seed pulumi.StringPtrOutput `pulumi:"seed"`
}

// NewRandomShuffle registers a new resource with the given unique name, arguments, and options.
func NewRandomShuffle(ctx *pulumi.Context,
	name string, args *RandomShuffleArgs, opts ...pulumi.ResourceOption) (*RandomShuffle, error) {
	if args == nil || args.Inputs == nil {
		return nil, errors.New("missing required argument 'Inputs'")
	}
	if args == nil {
		args = &RandomShuffleArgs{}
	}
	var resource RandomShuffle
	err := ctx.RegisterResource("random:index/randomShuffle:RandomShuffle", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomShuffle gets an existing RandomShuffle resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomShuffle(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomShuffleState, opts ...pulumi.ResourceOption) (*RandomShuffle, error) {
	var resource RandomShuffle
	err := ctx.ReadResource("random:index/randomShuffle:RandomShuffle", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomShuffle resources.
type randomShuffleState struct {
	// The list of strings to shuffle.
	Inputs []string `pulumi:"inputs"`
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The number of results to return. Defaults to
	// the number of items in the `input` list. If fewer items are requested,
	// some elements will be excluded from the result. If more items are requested,
	// items will be repeated in the result but not more frequently than the number
	// of items in the input list.
	ResultCount *int `pulumi:"resultCount"`
	// Random permutation of the list of strings given in `input`.
	Results []string `pulumi:"results"`
	// Arbitrary string with which to seed the random number
	// generator, in order to produce less-volatile permutations of the list.
	// **Important:** Even with an identical seed, it is not guaranteed that the
	// same permutation will be produced across different versions of the provider.
	// This argument causes the result to be *less volatile*, but not fixed for
	// all time.
	Seed *string `pulumi:"seed"`
}

type RandomShuffleState struct {
	// The list of strings to shuffle.
	Inputs pulumi.StringArrayInput
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapInput
	// The number of results to return. Defaults to
	// the number of items in the `input` list. If fewer items are requested,
	// some elements will be excluded from the result. If more items are requested,
	// items will be repeated in the result but not more frequently than the number
	// of items in the input list.
	ResultCount pulumi.IntPtrInput
	// Random permutation of the list of strings given in `input`.
	Results pulumi.StringArrayInput
	// Arbitrary string with which to seed the random number
	// generator, in order to produce less-volatile permutations of the list.
	// **Important:** Even with an identical seed, it is not guaranteed that the
	// same permutation will be produced across different versions of the provider.
	// This argument causes the result to be *less volatile*, but not fixed for
	// all time.
	Seed pulumi.StringPtrInput
}

func (RandomShuffleState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomShuffleState)(nil)).Elem()
}

type randomShuffleArgs struct {
	// The list of strings to shuffle.
	Inputs []string `pulumi:"inputs"`
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The number of results to return. Defaults to
	// the number of items in the `input` list. If fewer items are requested,
	// some elements will be excluded from the result. If more items are requested,
	// items will be repeated in the result but not more frequently than the number
	// of items in the input list.
	ResultCount *int `pulumi:"resultCount"`
	// Arbitrary string with which to seed the random number
	// generator, in order to produce less-volatile permutations of the list.
	// **Important:** Even with an identical seed, it is not guaranteed that the
	// same permutation will be produced across different versions of the provider.
	// This argument causes the result to be *less volatile*, but not fixed for
	// all time.
	Seed *string `pulumi:"seed"`
}

// The set of arguments for constructing a RandomShuffle resource.
type RandomShuffleArgs struct {
	// The list of strings to shuffle.
	Inputs pulumi.StringArrayInput
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapInput
	// The number of results to return. Defaults to
	// the number of items in the `input` list. If fewer items are requested,
	// some elements will be excluded from the result. If more items are requested,
	// items will be repeated in the result but not more frequently than the number
	// of items in the input list.
	ResultCount pulumi.IntPtrInput
	// Arbitrary string with which to seed the random number
	// generator, in order to produce less-volatile permutations of the list.
	// **Important:** Even with an identical seed, it is not guaranteed that the
	// same permutation will be produced across different versions of the provider.
	// This argument causes the result to be *less volatile*, but not fixed for
	// all time.
	Seed pulumi.StringPtrInput
}

func (RandomShuffleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomShuffleArgs)(nil)).Elem()
}
