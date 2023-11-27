// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

type IAMAPIKey struct {
	pulumi.CustomResourceState

	// The IAM API Key to match.
	Key pulumi.StringOutput `pulumi:"key"`
	// ❗ IAM API Key name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ❗ IAM API role ID.
	RoleId pulumi.StringOutput `pulumi:"roleId"`
	// Secret for the IAM API Key.
	Secret   pulumi.StringOutput        `pulumi:"secret"`
	Timeouts IAMAPIKeyTimeoutsPtrOutput `pulumi:"timeouts"`
}

// NewIAMAPIKey registers a new resource with the given unique name, arguments, and options.
func NewIAMAPIKey(ctx *pulumi.Context,
	name string, args *IAMAPIKeyArgs, opts ...pulumi.ResourceOption) (*IAMAPIKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RoleId == nil {
		return nil, errors.New("invalid value for required argument 'RoleId'")
	}
	secrets := pulumi.AdditionalSecretOutputs([]string{
		"secret",
	})
	opts = append(opts, secrets)
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource IAMAPIKey
	err := ctx.RegisterResource("exoscale:index/iAMAPIKey:IAMAPIKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetIAMAPIKey gets an existing IAMAPIKey resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetIAMAPIKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *IAMAPIKeyState, opts ...pulumi.ResourceOption) (*IAMAPIKey, error) {
	var resource IAMAPIKey
	err := ctx.ReadResource("exoscale:index/iAMAPIKey:IAMAPIKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering IAMAPIKey resources.
type iamapikeyState struct {
	// The IAM API Key to match.
	Key *string `pulumi:"key"`
	// ❗ IAM API Key name.
	Name *string `pulumi:"name"`
	// ❗ IAM API role ID.
	RoleId *string `pulumi:"roleId"`
	// Secret for the IAM API Key.
	Secret   *string            `pulumi:"secret"`
	Timeouts *IAMAPIKeyTimeouts `pulumi:"timeouts"`
}

type IAMAPIKeyState struct {
	// The IAM API Key to match.
	Key pulumi.StringPtrInput
	// ❗ IAM API Key name.
	Name pulumi.StringPtrInput
	// ❗ IAM API role ID.
	RoleId pulumi.StringPtrInput
	// Secret for the IAM API Key.
	Secret   pulumi.StringPtrInput
	Timeouts IAMAPIKeyTimeoutsPtrInput
}

func (IAMAPIKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*iamapikeyState)(nil)).Elem()
}

type iamapikeyArgs struct {
	// ❗ IAM API Key name.
	Name *string `pulumi:"name"`
	// ❗ IAM API role ID.
	RoleId   string             `pulumi:"roleId"`
	Timeouts *IAMAPIKeyTimeouts `pulumi:"timeouts"`
}

// The set of arguments for constructing a IAMAPIKey resource.
type IAMAPIKeyArgs struct {
	// ❗ IAM API Key name.
	Name pulumi.StringPtrInput
	// ❗ IAM API role ID.
	RoleId   pulumi.StringInput
	Timeouts IAMAPIKeyTimeoutsPtrInput
}

func (IAMAPIKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*iamapikeyArgs)(nil)).Elem()
}

type IAMAPIKeyInput interface {
	pulumi.Input

	ToIAMAPIKeyOutput() IAMAPIKeyOutput
	ToIAMAPIKeyOutputWithContext(ctx context.Context) IAMAPIKeyOutput
}

func (*IAMAPIKey) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAPIKey)(nil)).Elem()
}

func (i *IAMAPIKey) ToIAMAPIKeyOutput() IAMAPIKeyOutput {
	return i.ToIAMAPIKeyOutputWithContext(context.Background())
}

func (i *IAMAPIKey) ToIAMAPIKeyOutputWithContext(ctx context.Context) IAMAPIKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAPIKeyOutput)
}

// IAMAPIKeyArrayInput is an input type that accepts IAMAPIKeyArray and IAMAPIKeyArrayOutput values.
// You can construct a concrete instance of `IAMAPIKeyArrayInput` via:
//
//	IAMAPIKeyArray{ IAMAPIKeyArgs{...} }
type IAMAPIKeyArrayInput interface {
	pulumi.Input

	ToIAMAPIKeyArrayOutput() IAMAPIKeyArrayOutput
	ToIAMAPIKeyArrayOutputWithContext(context.Context) IAMAPIKeyArrayOutput
}

type IAMAPIKeyArray []IAMAPIKeyInput

func (IAMAPIKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMAPIKey)(nil)).Elem()
}

func (i IAMAPIKeyArray) ToIAMAPIKeyArrayOutput() IAMAPIKeyArrayOutput {
	return i.ToIAMAPIKeyArrayOutputWithContext(context.Background())
}

func (i IAMAPIKeyArray) ToIAMAPIKeyArrayOutputWithContext(ctx context.Context) IAMAPIKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAPIKeyArrayOutput)
}

// IAMAPIKeyMapInput is an input type that accepts IAMAPIKeyMap and IAMAPIKeyMapOutput values.
// You can construct a concrete instance of `IAMAPIKeyMapInput` via:
//
//	IAMAPIKeyMap{ "key": IAMAPIKeyArgs{...} }
type IAMAPIKeyMapInput interface {
	pulumi.Input

	ToIAMAPIKeyMapOutput() IAMAPIKeyMapOutput
	ToIAMAPIKeyMapOutputWithContext(context.Context) IAMAPIKeyMapOutput
}

type IAMAPIKeyMap map[string]IAMAPIKeyInput

func (IAMAPIKeyMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMAPIKey)(nil)).Elem()
}

func (i IAMAPIKeyMap) ToIAMAPIKeyMapOutput() IAMAPIKeyMapOutput {
	return i.ToIAMAPIKeyMapOutputWithContext(context.Background())
}

func (i IAMAPIKeyMap) ToIAMAPIKeyMapOutputWithContext(ctx context.Context) IAMAPIKeyMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IAMAPIKeyMapOutput)
}

type IAMAPIKeyOutput struct{ *pulumi.OutputState }

func (IAMAPIKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IAMAPIKey)(nil)).Elem()
}

func (o IAMAPIKeyOutput) ToIAMAPIKeyOutput() IAMAPIKeyOutput {
	return o
}

func (o IAMAPIKeyOutput) ToIAMAPIKeyOutputWithContext(ctx context.Context) IAMAPIKeyOutput {
	return o
}

// The IAM API Key to match.
func (o IAMAPIKeyOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAPIKey) pulumi.StringOutput { return v.Key }).(pulumi.StringOutput)
}

// ❗ IAM API Key name.
func (o IAMAPIKeyOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAPIKey) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ❗ IAM API role ID.
func (o IAMAPIKeyOutput) RoleId() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAPIKey) pulumi.StringOutput { return v.RoleId }).(pulumi.StringOutput)
}

// Secret for the IAM API Key.
func (o IAMAPIKeyOutput) Secret() pulumi.StringOutput {
	return o.ApplyT(func(v *IAMAPIKey) pulumi.StringOutput { return v.Secret }).(pulumi.StringOutput)
}

func (o IAMAPIKeyOutput) Timeouts() IAMAPIKeyTimeoutsPtrOutput {
	return o.ApplyT(func(v *IAMAPIKey) IAMAPIKeyTimeoutsPtrOutput { return v.Timeouts }).(IAMAPIKeyTimeoutsPtrOutput)
}

type IAMAPIKeyArrayOutput struct{ *pulumi.OutputState }

func (IAMAPIKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*IAMAPIKey)(nil)).Elem()
}

func (o IAMAPIKeyArrayOutput) ToIAMAPIKeyArrayOutput() IAMAPIKeyArrayOutput {
	return o
}

func (o IAMAPIKeyArrayOutput) ToIAMAPIKeyArrayOutputWithContext(ctx context.Context) IAMAPIKeyArrayOutput {
	return o
}

func (o IAMAPIKeyArrayOutput) Index(i pulumi.IntInput) IAMAPIKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *IAMAPIKey {
		return vs[0].([]*IAMAPIKey)[vs[1].(int)]
	}).(IAMAPIKeyOutput)
}

type IAMAPIKeyMapOutput struct{ *pulumi.OutputState }

func (IAMAPIKeyMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*IAMAPIKey)(nil)).Elem()
}

func (o IAMAPIKeyMapOutput) ToIAMAPIKeyMapOutput() IAMAPIKeyMapOutput {
	return o
}

func (o IAMAPIKeyMapOutput) ToIAMAPIKeyMapOutputWithContext(ctx context.Context) IAMAPIKeyMapOutput {
	return o
}

func (o IAMAPIKeyMapOutput) MapIndex(k pulumi.StringInput) IAMAPIKeyOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *IAMAPIKey {
		return vs[0].(map[string]*IAMAPIKey)[vs[1].(string)]
	}).(IAMAPIKeyOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAPIKeyInput)(nil)).Elem(), &IAMAPIKey{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAPIKeyArrayInput)(nil)).Elem(), IAMAPIKeyArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IAMAPIKeyMapInput)(nil)).Elem(), IAMAPIKeyMap{})
	pulumi.RegisterOutputType(IAMAPIKeyOutput{})
	pulumi.RegisterOutputType(IAMAPIKeyArrayOutput{})
	pulumi.RegisterOutputType(IAMAPIKeyMapOutput{})
}
