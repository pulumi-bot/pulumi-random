// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// The resource `RandomInteger` generates random values from a given range, described by the `min` and `max` attributes of a given resource.
//
// This resource can be used in conjunction with resources that have
// the `createBeforeDestroy` lifecycle flag set, to avoid conflicts with
// unique names during the brief period where both the old and new resources
// exist concurrently.
//
// ## Example Usage
//
// The following example shows how to generate a random priority between 1 and 50000 for
// a `awsAlbListenerRule` resource:
//
// ```go
// package main
//
// import (
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/alb"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		priority, err := random.NewRandomInteger(ctx, "priority", &random.RandomIntegerArgs{
// 			Keepers: pulumi.AnyMap{
// 				"listener_arn": pulumi.Any(_var.Listener_arn),
// 			},
// 			Max: pulumi.Int(50000),
// 			Min: pulumi.Int(1),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = alb.NewListenerRule(ctx, "main", &alb.ListenerRuleArgs{
// 			Actions: alb.ListenerRuleActionArray{
// 				&alb.ListenerRuleActionArgs{
// 					TargetGroupArn: pulumi.Any(_var.Target_group_arn),
// 					Type:           pulumi.String("forward"),
// 				},
// 			},
// 			ListenerArn: pulumi.Any(_var.Listener_arn),
// 			Priority:    priority.Result,
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
//
// The result of the above will set a random priority.
//
// ## Import
//
// Random integers can be imported using the `result`, `min`, and `max`, with an optional `seed`. This can be used to replace a config value with a value interpolated from the random provider without experiencing diffs. Example (values are separated by a `,`)
//
// ```sh
//  $ pulumi import random:index/randomInteger:RandomInteger priority 15390,1,50000
// ```
type RandomInteger struct {
	pulumi.CustomResourceState

	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapOutput `pulumi:"keepers"`
	// The maximum inclusive value of the range.
	Max pulumi.IntOutput `pulumi:"max"`
	// The minimum inclusive value of the range.
	Min pulumi.IntOutput `pulumi:"min"`
	// (int) The random Integer result.
	Result pulumi.IntOutput `pulumi:"result"`
	// A custom seed to always produce the same value.
	Seed pulumi.StringPtrOutput `pulumi:"seed"`
}

// NewRandomInteger registers a new resource with the given unique name, arguments, and options.
func NewRandomInteger(ctx *pulumi.Context,
	name string, args *RandomIntegerArgs, opts ...pulumi.ResourceOption) (*RandomInteger, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Max == nil {
		return nil, errors.New("invalid value for required argument 'Max'")
	}
	if args.Min == nil {
		return nil, errors.New("invalid value for required argument 'Min'")
	}
	var resource RandomInteger
	err := ctx.RegisterResource("random:index/randomInteger:RandomInteger", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomInteger gets an existing RandomInteger resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomInteger(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomIntegerState, opts ...pulumi.ResourceOption) (*RandomInteger, error) {
	var resource RandomInteger
	err := ctx.ReadResource("random:index/randomInteger:RandomInteger", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomInteger resources.
type randomIntegerState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The maximum inclusive value of the range.
	Max *int `pulumi:"max"`
	// The minimum inclusive value of the range.
	Min *int `pulumi:"min"`
	// (int) The random Integer result.
	Result *int `pulumi:"result"`
	// A custom seed to always produce the same value.
	Seed *string `pulumi:"seed"`
}

type RandomIntegerState struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapInput
	// The maximum inclusive value of the range.
	Max pulumi.IntPtrInput
	// The minimum inclusive value of the range.
	Min pulumi.IntPtrInput
	// (int) The random Integer result.
	Result pulumi.IntPtrInput
	// A custom seed to always produce the same value.
	Seed pulumi.StringPtrInput
}

func (RandomIntegerState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomIntegerState)(nil)).Elem()
}

type randomIntegerArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers map[string]interface{} `pulumi:"keepers"`
	// The maximum inclusive value of the range.
	Max int `pulumi:"max"`
	// The minimum inclusive value of the range.
	Min int `pulumi:"min"`
	// A custom seed to always produce the same value.
	Seed *string `pulumi:"seed"`
}

// The set of arguments for constructing a RandomInteger resource.
type RandomIntegerArgs struct {
	// Arbitrary map of values that, when changed, will
	// trigger a new id to be generated.
	Keepers pulumi.MapInput
	// The maximum inclusive value of the range.
	Max pulumi.IntInput
	// The minimum inclusive value of the range.
	Min pulumi.IntInput
	// A custom seed to always produce the same value.
	Seed pulumi.StringPtrInput
}

func (RandomIntegerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomIntegerArgs)(nil)).Elem()
}

type RandomIntegerInput interface {
	pulumi.Input

	ToRandomIntegerOutput() RandomIntegerOutput
	ToRandomIntegerOutputWithContext(ctx context.Context) RandomIntegerOutput
}

func (*RandomInteger) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomInteger)(nil))
}

func (i *RandomInteger) ToRandomIntegerOutput() RandomIntegerOutput {
	return i.ToRandomIntegerOutputWithContext(context.Background())
}

func (i *RandomInteger) ToRandomIntegerOutputWithContext(ctx context.Context) RandomIntegerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIntegerOutput)
}

func (i *RandomInteger) ToRandomIntegerPtrOutput() RandomIntegerPtrOutput {
	return i.ToRandomIntegerPtrOutputWithContext(context.Background())
}

func (i *RandomInteger) ToRandomIntegerPtrOutputWithContext(ctx context.Context) RandomIntegerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIntegerPtrOutput)
}

type RandomIntegerPtrInput interface {
	pulumi.Input

	ToRandomIntegerPtrOutput() RandomIntegerPtrOutput
	ToRandomIntegerPtrOutputWithContext(ctx context.Context) RandomIntegerPtrOutput
}

type randomIntegerPtrType RandomIntegerArgs

func (*randomIntegerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomInteger)(nil))
}

func (i *randomIntegerPtrType) ToRandomIntegerPtrOutput() RandomIntegerPtrOutput {
	return i.ToRandomIntegerPtrOutputWithContext(context.Background())
}

func (i *randomIntegerPtrType) ToRandomIntegerPtrOutputWithContext(ctx context.Context) RandomIntegerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIntegerOutput).ToRandomIntegerPtrOutput()
}

// RandomIntegerArrayInput is an input type that accepts RandomIntegerArray and RandomIntegerArrayOutput values.
// You can construct a concrete instance of `RandomIntegerArrayInput` via:
//
//          RandomIntegerArray{ RandomIntegerArgs{...} }
type RandomIntegerArrayInput interface {
	pulumi.Input

	ToRandomIntegerArrayOutput() RandomIntegerArrayOutput
	ToRandomIntegerArrayOutputWithContext(context.Context) RandomIntegerArrayOutput
}

type RandomIntegerArray []RandomIntegerInput

func (RandomIntegerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RandomInteger)(nil))
}

func (i RandomIntegerArray) ToRandomIntegerArrayOutput() RandomIntegerArrayOutput {
	return i.ToRandomIntegerArrayOutputWithContext(context.Background())
}

func (i RandomIntegerArray) ToRandomIntegerArrayOutputWithContext(ctx context.Context) RandomIntegerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIntegerArrayOutput)
}

// RandomIntegerMapInput is an input type that accepts RandomIntegerMap and RandomIntegerMapOutput values.
// You can construct a concrete instance of `RandomIntegerMapInput` via:
//
//          RandomIntegerMap{ "key": RandomIntegerArgs{...} }
type RandomIntegerMapInput interface {
	pulumi.Input

	ToRandomIntegerMapOutput() RandomIntegerMapOutput
	ToRandomIntegerMapOutputWithContext(context.Context) RandomIntegerMapOutput
}

type RandomIntegerMap map[string]RandomIntegerInput

func (RandomIntegerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RandomInteger)(nil))
}

func (i RandomIntegerMap) ToRandomIntegerMapOutput() RandomIntegerMapOutput {
	return i.ToRandomIntegerMapOutputWithContext(context.Background())
}

func (i RandomIntegerMap) ToRandomIntegerMapOutputWithContext(ctx context.Context) RandomIntegerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomIntegerMapOutput)
}

type RandomIntegerOutput struct {
	*pulumi.OutputState
}

func (RandomIntegerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomInteger)(nil))
}

func (o RandomIntegerOutput) ToRandomIntegerOutput() RandomIntegerOutput {
	return o
}

func (o RandomIntegerOutput) ToRandomIntegerOutputWithContext(ctx context.Context) RandomIntegerOutput {
	return o
}

func (o RandomIntegerOutput) ToRandomIntegerPtrOutput() RandomIntegerPtrOutput {
	return o.ToRandomIntegerPtrOutputWithContext(context.Background())
}

func (o RandomIntegerOutput) ToRandomIntegerPtrOutputWithContext(ctx context.Context) RandomIntegerPtrOutput {
	return o.ApplyT(func(v RandomInteger) *RandomInteger {
		return &v
	}).(RandomIntegerPtrOutput)
}

type RandomIntegerPtrOutput struct {
	*pulumi.OutputState
}

func (RandomIntegerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomInteger)(nil))
}

func (o RandomIntegerPtrOutput) ToRandomIntegerPtrOutput() RandomIntegerPtrOutput {
	return o
}

func (o RandomIntegerPtrOutput) ToRandomIntegerPtrOutputWithContext(ctx context.Context) RandomIntegerPtrOutput {
	return o
}

type RandomIntegerArrayOutput struct{ *pulumi.OutputState }

func (RandomIntegerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RandomInteger)(nil))
}

func (o RandomIntegerArrayOutput) ToRandomIntegerArrayOutput() RandomIntegerArrayOutput {
	return o
}

func (o RandomIntegerArrayOutput) ToRandomIntegerArrayOutputWithContext(ctx context.Context) RandomIntegerArrayOutput {
	return o
}

func (o RandomIntegerArrayOutput) Index(i pulumi.IntInput) RandomIntegerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RandomInteger {
		return vs[0].([]RandomInteger)[vs[1].(int)]
	}).(RandomIntegerOutput)
}

type RandomIntegerMapOutput struct{ *pulumi.OutputState }

func (RandomIntegerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RandomInteger)(nil))
}

func (o RandomIntegerMapOutput) ToRandomIntegerMapOutput() RandomIntegerMapOutput {
	return o
}

func (o RandomIntegerMapOutput) ToRandomIntegerMapOutputWithContext(ctx context.Context) RandomIntegerMapOutput {
	return o
}

func (o RandomIntegerMapOutput) MapIndex(k pulumi.StringInput) RandomIntegerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RandomInteger {
		return vs[0].(map[string]RandomInteger)[vs[1].(string)]
	}).(RandomIntegerOutput)
}

func init() {
	pulumi.RegisterOutputType(RandomIntegerOutput{})
	pulumi.RegisterOutputType(RandomIntegerPtrOutput{})
	pulumi.RegisterOutputType(RandomIntegerArrayOutput{})
	pulumi.RegisterOutputType(RandomIntegerMapOutput{})
}
