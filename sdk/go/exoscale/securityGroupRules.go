// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumix"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

// !> **WARNING:** This resource is **DEPRECATED** and will be removed in the next major version. Please use the SecurityGroupRule instead (or refer to the ad-hoc migration guide).
type SecurityGroupRules struct {
	pulumi.CustomResourceState

	// A security group rule definition (can be specified multiple times).
	Egresses SecurityGroupRulesEgressArrayOutput `pulumi:"egresses"`
	// A security group rule definition (can be specified multiple times).
	Ingresses SecurityGroupRulesIngressArrayOutput `pulumi:"ingresses"`
	// ❗ The security group (name) the rules apply to (conflicts with `securityGroupId`).
	SecurityGroup pulumi.StringOutput `pulumi:"securityGroup"`
	// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
	SecurityGroupId pulumi.StringOutput `pulumi:"securityGroupId"`
}

// NewSecurityGroupRules registers a new resource with the given unique name, arguments, and options.
func NewSecurityGroupRules(ctx *pulumi.Context,
	name string, args *SecurityGroupRulesArgs, opts ...pulumi.ResourceOption) (*SecurityGroupRules, error) {
	if args == nil {
		args = &SecurityGroupRulesArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SecurityGroupRules
	err := ctx.RegisterResource("exoscale:index/securityGroupRules:SecurityGroupRules", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSecurityGroupRules gets an existing SecurityGroupRules resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSecurityGroupRules(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityGroupRulesState, opts ...pulumi.ResourceOption) (*SecurityGroupRules, error) {
	var resource SecurityGroupRules
	err := ctx.ReadResource("exoscale:index/securityGroupRules:SecurityGroupRules", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SecurityGroupRules resources.
type securityGroupRulesState struct {
	// A security group rule definition (can be specified multiple times).
	Egresses []SecurityGroupRulesEgress `pulumi:"egresses"`
	// A security group rule definition (can be specified multiple times).
	Ingresses []SecurityGroupRulesIngress `pulumi:"ingresses"`
	// ❗ The security group (name) the rules apply to (conflicts with `securityGroupId`).
	SecurityGroup *string `pulumi:"securityGroup"`
	// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
	SecurityGroupId *string `pulumi:"securityGroupId"`
}

type SecurityGroupRulesState struct {
	// A security group rule definition (can be specified multiple times).
	Egresses SecurityGroupRulesEgressArrayInput
	// A security group rule definition (can be specified multiple times).
	Ingresses SecurityGroupRulesIngressArrayInput
	// ❗ The security group (name) the rules apply to (conflicts with `securityGroupId`).
	SecurityGroup pulumi.StringPtrInput
	// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
	SecurityGroupId pulumi.StringPtrInput
}

func (SecurityGroupRulesState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupRulesState)(nil)).Elem()
}

type securityGroupRulesArgs struct {
	// A security group rule definition (can be specified multiple times).
	Egresses []SecurityGroupRulesEgress `pulumi:"egresses"`
	// A security group rule definition (can be specified multiple times).
	Ingresses []SecurityGroupRulesIngress `pulumi:"ingresses"`
	// ❗ The security group (name) the rules apply to (conflicts with `securityGroupId`).
	SecurityGroup *string `pulumi:"securityGroup"`
	// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
	SecurityGroupId *string `pulumi:"securityGroupId"`
}

// The set of arguments for constructing a SecurityGroupRules resource.
type SecurityGroupRulesArgs struct {
	// A security group rule definition (can be specified multiple times).
	Egresses SecurityGroupRulesEgressArrayInput
	// A security group rule definition (can be specified multiple times).
	Ingresses SecurityGroupRulesIngressArrayInput
	// ❗ The security group (name) the rules apply to (conflicts with `securityGroupId`).
	SecurityGroup pulumi.StringPtrInput
	// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
	SecurityGroupId pulumi.StringPtrInput
}

func (SecurityGroupRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityGroupRulesArgs)(nil)).Elem()
}

type SecurityGroupRulesInput interface {
	pulumi.Input

	ToSecurityGroupRulesOutput() SecurityGroupRulesOutput
	ToSecurityGroupRulesOutputWithContext(ctx context.Context) SecurityGroupRulesOutput
}

func (*SecurityGroupRules) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupRules)(nil)).Elem()
}

func (i *SecurityGroupRules) ToSecurityGroupRulesOutput() SecurityGroupRulesOutput {
	return i.ToSecurityGroupRulesOutputWithContext(context.Background())
}

func (i *SecurityGroupRules) ToSecurityGroupRulesOutputWithContext(ctx context.Context) SecurityGroupRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRulesOutput)
}

func (i *SecurityGroupRules) ToOutput(ctx context.Context) pulumix.Output[*SecurityGroupRules] {
	return pulumix.Output[*SecurityGroupRules]{
		OutputState: i.ToSecurityGroupRulesOutputWithContext(ctx).OutputState,
	}
}

// SecurityGroupRulesArrayInput is an input type that accepts SecurityGroupRulesArray and SecurityGroupRulesArrayOutput values.
// You can construct a concrete instance of `SecurityGroupRulesArrayInput` via:
//
//	SecurityGroupRulesArray{ SecurityGroupRulesArgs{...} }
type SecurityGroupRulesArrayInput interface {
	pulumi.Input

	ToSecurityGroupRulesArrayOutput() SecurityGroupRulesArrayOutput
	ToSecurityGroupRulesArrayOutputWithContext(context.Context) SecurityGroupRulesArrayOutput
}

type SecurityGroupRulesArray []SecurityGroupRulesInput

func (SecurityGroupRulesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroupRules)(nil)).Elem()
}

func (i SecurityGroupRulesArray) ToSecurityGroupRulesArrayOutput() SecurityGroupRulesArrayOutput {
	return i.ToSecurityGroupRulesArrayOutputWithContext(context.Background())
}

func (i SecurityGroupRulesArray) ToSecurityGroupRulesArrayOutputWithContext(ctx context.Context) SecurityGroupRulesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRulesArrayOutput)
}

func (i SecurityGroupRulesArray) ToOutput(ctx context.Context) pulumix.Output[[]*SecurityGroupRules] {
	return pulumix.Output[[]*SecurityGroupRules]{
		OutputState: i.ToSecurityGroupRulesArrayOutputWithContext(ctx).OutputState,
	}
}

// SecurityGroupRulesMapInput is an input type that accepts SecurityGroupRulesMap and SecurityGroupRulesMapOutput values.
// You can construct a concrete instance of `SecurityGroupRulesMapInput` via:
//
//	SecurityGroupRulesMap{ "key": SecurityGroupRulesArgs{...} }
type SecurityGroupRulesMapInput interface {
	pulumi.Input

	ToSecurityGroupRulesMapOutput() SecurityGroupRulesMapOutput
	ToSecurityGroupRulesMapOutputWithContext(context.Context) SecurityGroupRulesMapOutput
}

type SecurityGroupRulesMap map[string]SecurityGroupRulesInput

func (SecurityGroupRulesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroupRules)(nil)).Elem()
}

func (i SecurityGroupRulesMap) ToSecurityGroupRulesMapOutput() SecurityGroupRulesMapOutput {
	return i.ToSecurityGroupRulesMapOutputWithContext(context.Background())
}

func (i SecurityGroupRulesMap) ToSecurityGroupRulesMapOutputWithContext(ctx context.Context) SecurityGroupRulesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityGroupRulesMapOutput)
}

func (i SecurityGroupRulesMap) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecurityGroupRules] {
	return pulumix.Output[map[string]*SecurityGroupRules]{
		OutputState: i.ToSecurityGroupRulesMapOutputWithContext(ctx).OutputState,
	}
}

type SecurityGroupRulesOutput struct{ *pulumi.OutputState }

func (SecurityGroupRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityGroupRules)(nil)).Elem()
}

func (o SecurityGroupRulesOutput) ToSecurityGroupRulesOutput() SecurityGroupRulesOutput {
	return o
}

func (o SecurityGroupRulesOutput) ToSecurityGroupRulesOutputWithContext(ctx context.Context) SecurityGroupRulesOutput {
	return o
}

func (o SecurityGroupRulesOutput) ToOutput(ctx context.Context) pulumix.Output[*SecurityGroupRules] {
	return pulumix.Output[*SecurityGroupRules]{
		OutputState: o.OutputState,
	}
}

// A security group rule definition (can be specified multiple times).
func (o SecurityGroupRulesOutput) Egresses() SecurityGroupRulesEgressArrayOutput {
	return o.ApplyT(func(v *SecurityGroupRules) SecurityGroupRulesEgressArrayOutput { return v.Egresses }).(SecurityGroupRulesEgressArrayOutput)
}

// A security group rule definition (can be specified multiple times).
func (o SecurityGroupRulesOutput) Ingresses() SecurityGroupRulesIngressArrayOutput {
	return o.ApplyT(func(v *SecurityGroupRules) SecurityGroupRulesIngressArrayOutput { return v.Ingresses }).(SecurityGroupRulesIngressArrayOutput)
}

// ❗ The security group (name) the rules apply to (conflicts with `securityGroupId`).
func (o SecurityGroupRulesOutput) SecurityGroup() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupRules) pulumi.StringOutput { return v.SecurityGroup }).(pulumi.StringOutput)
}

// ❗ The security group (ID) the rules apply to (conficts with `security_group)`.
func (o SecurityGroupRulesOutput) SecurityGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityGroupRules) pulumi.StringOutput { return v.SecurityGroupId }).(pulumi.StringOutput)
}

type SecurityGroupRulesArrayOutput struct{ *pulumi.OutputState }

func (SecurityGroupRulesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SecurityGroupRules)(nil)).Elem()
}

func (o SecurityGroupRulesArrayOutput) ToSecurityGroupRulesArrayOutput() SecurityGroupRulesArrayOutput {
	return o
}

func (o SecurityGroupRulesArrayOutput) ToSecurityGroupRulesArrayOutputWithContext(ctx context.Context) SecurityGroupRulesArrayOutput {
	return o
}

func (o SecurityGroupRulesArrayOutput) ToOutput(ctx context.Context) pulumix.Output[[]*SecurityGroupRules] {
	return pulumix.Output[[]*SecurityGroupRules]{
		OutputState: o.OutputState,
	}
}

func (o SecurityGroupRulesArrayOutput) Index(i pulumi.IntInput) SecurityGroupRulesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SecurityGroupRules {
		return vs[0].([]*SecurityGroupRules)[vs[1].(int)]
	}).(SecurityGroupRulesOutput)
}

type SecurityGroupRulesMapOutput struct{ *pulumi.OutputState }

func (SecurityGroupRulesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SecurityGroupRules)(nil)).Elem()
}

func (o SecurityGroupRulesMapOutput) ToSecurityGroupRulesMapOutput() SecurityGroupRulesMapOutput {
	return o
}

func (o SecurityGroupRulesMapOutput) ToSecurityGroupRulesMapOutputWithContext(ctx context.Context) SecurityGroupRulesMapOutput {
	return o
}

func (o SecurityGroupRulesMapOutput) ToOutput(ctx context.Context) pulumix.Output[map[string]*SecurityGroupRules] {
	return pulumix.Output[map[string]*SecurityGroupRules]{
		OutputState: o.OutputState,
	}
}

func (o SecurityGroupRulesMapOutput) MapIndex(k pulumi.StringInput) SecurityGroupRulesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SecurityGroupRules {
		return vs[0].(map[string]*SecurityGroupRules)[vs[1].(string)]
	}).(SecurityGroupRulesOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupRulesInput)(nil)).Elem(), &SecurityGroupRules{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupRulesArrayInput)(nil)).Elem(), SecurityGroupRulesArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityGroupRulesMapInput)(nil)).Elem(), SecurityGroupRulesMap{})
	pulumi.RegisterOutputType(SecurityGroupRulesOutput{})
	pulumi.RegisterOutputType(SecurityGroupRulesArrayOutput{})
	pulumi.RegisterOutputType(SecurityGroupRulesMapOutput{})
}
