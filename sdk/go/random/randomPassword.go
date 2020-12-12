// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package random

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

// Identical to the `RandomString` resource with the exception that the
// result is treated as sensitive and, thus, _not_ displayed in console output.
//
// > **Note:** All attributes including the generated password will be stored in
// the raw state as plain-text.
//
// This resource *does* use a cryptographic random number generator.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-aws/sdk/v3/go/aws/rds"
// 	"github.com/pulumi/pulumi-random/sdk/v2/go/random"
// 	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		password, err := random.NewRandomPassword(ctx, "password", &random.RandomPasswordArgs{
// 			Length:          pulumi.Int(16),
// 			Special:         pulumi.Bool(true),
// 			OverrideSpecial: pulumi.String(fmt.Sprintf("%v%v%v", "_", "%", "@")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = rds.NewInstance(ctx, "example", &rds.InstanceArgs{
// 			InstanceClass:    pulumi.String("db.t3.micro"),
// 			AllocatedStorage: pulumi.Int(64),
// 			Engine:           pulumi.String("mysql"),
// 			Username:         pulumi.String("someone"),
// 			Password:         password.Result,
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
// Random Password can be imported by specifying the value of the string
//
// ```sh
//  $ pulumi import random:index/randomPassword:RandomPassword password securepassword
// ```
type RandomPassword struct {
	pulumi.CustomResourceState

	Keepers         pulumi.MapOutput       `pulumi:"keepers"`
	Length          pulumi.IntOutput       `pulumi:"length"`
	Lower           pulumi.BoolPtrOutput   `pulumi:"lower"`
	MinLower        pulumi.IntPtrOutput    `pulumi:"minLower"`
	MinNumeric      pulumi.IntPtrOutput    `pulumi:"minNumeric"`
	MinSpecial      pulumi.IntPtrOutput    `pulumi:"minSpecial"`
	MinUpper        pulumi.IntPtrOutput    `pulumi:"minUpper"`
	Number          pulumi.BoolPtrOutput   `pulumi:"number"`
	OverrideSpecial pulumi.StringPtrOutput `pulumi:"overrideSpecial"`
	Result          pulumi.StringOutput    `pulumi:"result"`
	Special         pulumi.BoolPtrOutput   `pulumi:"special"`
	Upper           pulumi.BoolPtrOutput   `pulumi:"upper"`
}

// NewRandomPassword registers a new resource with the given unique name, arguments, and options.
func NewRandomPassword(ctx *pulumi.Context,
	name string, args *RandomPasswordArgs, opts ...pulumi.ResourceOption) (*RandomPassword, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Length == nil {
		return nil, errors.New("invalid value for required argument 'Length'")
	}
	var resource RandomPassword
	err := ctx.RegisterResource("random:index/randomPassword:RandomPassword", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRandomPassword gets an existing RandomPassword resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRandomPassword(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RandomPasswordState, opts ...pulumi.ResourceOption) (*RandomPassword, error) {
	var resource RandomPassword
	err := ctx.ReadResource("random:index/randomPassword:RandomPassword", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering RandomPassword resources.
type randomPasswordState struct {
	Keepers         map[string]interface{} `pulumi:"keepers"`
	Length          *int                   `pulumi:"length"`
	Lower           *bool                  `pulumi:"lower"`
	MinLower        *int                   `pulumi:"minLower"`
	MinNumeric      *int                   `pulumi:"minNumeric"`
	MinSpecial      *int                   `pulumi:"minSpecial"`
	MinUpper        *int                   `pulumi:"minUpper"`
	Number          *bool                  `pulumi:"number"`
	OverrideSpecial *string                `pulumi:"overrideSpecial"`
	Result          *string                `pulumi:"result"`
	Special         *bool                  `pulumi:"special"`
	Upper           *bool                  `pulumi:"upper"`
}

type RandomPasswordState struct {
	Keepers         pulumi.MapInput
	Length          pulumi.IntPtrInput
	Lower           pulumi.BoolPtrInput
	MinLower        pulumi.IntPtrInput
	MinNumeric      pulumi.IntPtrInput
	MinSpecial      pulumi.IntPtrInput
	MinUpper        pulumi.IntPtrInput
	Number          pulumi.BoolPtrInput
	OverrideSpecial pulumi.StringPtrInput
	Result          pulumi.StringPtrInput
	Special         pulumi.BoolPtrInput
	Upper           pulumi.BoolPtrInput
}

func (RandomPasswordState) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPasswordState)(nil)).Elem()
}

type randomPasswordArgs struct {
	Keepers         map[string]interface{} `pulumi:"keepers"`
	Length          int                    `pulumi:"length"`
	Lower           *bool                  `pulumi:"lower"`
	MinLower        *int                   `pulumi:"minLower"`
	MinNumeric      *int                   `pulumi:"minNumeric"`
	MinSpecial      *int                   `pulumi:"minSpecial"`
	MinUpper        *int                   `pulumi:"minUpper"`
	Number          *bool                  `pulumi:"number"`
	OverrideSpecial *string                `pulumi:"overrideSpecial"`
	Special         *bool                  `pulumi:"special"`
	Upper           *bool                  `pulumi:"upper"`
}

// The set of arguments for constructing a RandomPassword resource.
type RandomPasswordArgs struct {
	Keepers         pulumi.MapInput
	Length          pulumi.IntInput
	Lower           pulumi.BoolPtrInput
	MinLower        pulumi.IntPtrInput
	MinNumeric      pulumi.IntPtrInput
	MinSpecial      pulumi.IntPtrInput
	MinUpper        pulumi.IntPtrInput
	Number          pulumi.BoolPtrInput
	OverrideSpecial pulumi.StringPtrInput
	Special         pulumi.BoolPtrInput
	Upper           pulumi.BoolPtrInput
}

func (RandomPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*randomPasswordArgs)(nil)).Elem()
}

type RandomPasswordInput interface {
	pulumi.Input

	ToRandomPasswordOutput() RandomPasswordOutput
	ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput
}

func (*RandomPassword) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomPassword)(nil))
}

func (i *RandomPassword) ToRandomPasswordOutput() RandomPasswordOutput {
	return i.ToRandomPasswordOutputWithContext(context.Background())
}

func (i *RandomPassword) ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordOutput)
}

func (i *RandomPassword) ToRandomPasswordPtrOutput() RandomPasswordPtrOutput {
	return i.ToRandomPasswordPtrOutputWithContext(context.Background())
}

func (i *RandomPassword) ToRandomPasswordPtrOutputWithContext(ctx context.Context) RandomPasswordPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RandomPasswordPtrOutput)
}

type RandomPasswordPtrInput interface {
	pulumi.Input

	ToRandomPasswordPtrOutput() RandomPasswordPtrOutput
	ToRandomPasswordPtrOutputWithContext(ctx context.Context) RandomPasswordPtrOutput
}

type RandomPasswordOutput struct {
	*pulumi.OutputState
}

func (RandomPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RandomPassword)(nil))
}

func (o RandomPasswordOutput) ToRandomPasswordOutput() RandomPasswordOutput {
	return o
}

func (o RandomPasswordOutput) ToRandomPasswordOutputWithContext(ctx context.Context) RandomPasswordOutput {
	return o
}

type RandomPasswordPtrOutput struct {
	*pulumi.OutputState
}

func (RandomPasswordPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RandomPassword)(nil))
}

func (o RandomPasswordPtrOutput) ToRandomPasswordPtrOutput() RandomPasswordPtrOutput {
	return o
}

func (o RandomPasswordPtrOutput) ToRandomPasswordPtrOutputWithContext(ctx context.Context) RandomPasswordPtrOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(RandomPasswordOutput{})
	pulumi.RegisterOutputType(RandomPasswordPtrOutput{})
}
