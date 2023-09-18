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

func LookupAntiAffinityGroup(ctx *pulumi.Context, args *LookupAntiAffinityGroupArgs, opts ...pulumi.InvokeOption) (*LookupAntiAffinityGroupResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupAntiAffinityGroupResult
	err := ctx.Invoke("exoscale:index/getAntiAffinityGroup:getAntiAffinityGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getAntiAffinityGroup.
type LookupAntiAffinityGroupArgs struct {
	// The anti-affinity group ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The group name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
}

// A collection of values returned by getAntiAffinityGroup.
type LookupAntiAffinityGroupResult struct {
	// The anti-affinity group ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The list of attached exoscale*compute*instance (IDs).
	Instances []string `pulumi:"instances"`
	// The group name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
}

func LookupAntiAffinityGroupOutput(ctx *pulumi.Context, args LookupAntiAffinityGroupOutputArgs, opts ...pulumi.InvokeOption) LookupAntiAffinityGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAntiAffinityGroupResult, error) {
			args := v.(LookupAntiAffinityGroupArgs)
			r, err := LookupAntiAffinityGroup(ctx, &args, opts...)
			var s LookupAntiAffinityGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAntiAffinityGroupResultOutput)
}

// A collection of arguments for invoking getAntiAffinityGroup.
type LookupAntiAffinityGroupOutputArgs struct {
	// The anti-affinity group ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The group name to match (conflicts with `id`).
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (LookupAntiAffinityGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAntiAffinityGroupArgs)(nil)).Elem()
}

// A collection of values returned by getAntiAffinityGroup.
type LookupAntiAffinityGroupResultOutput struct{ *pulumi.OutputState }

func (LookupAntiAffinityGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAntiAffinityGroupResult)(nil)).Elem()
}

func (o LookupAntiAffinityGroupResultOutput) ToLookupAntiAffinityGroupResultOutput() LookupAntiAffinityGroupResultOutput {
	return o
}

func (o LookupAntiAffinityGroupResultOutput) ToLookupAntiAffinityGroupResultOutputWithContext(ctx context.Context) LookupAntiAffinityGroupResultOutput {
	return o
}

func (o LookupAntiAffinityGroupResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupAntiAffinityGroupResult] {
	return pulumix.Output[LookupAntiAffinityGroupResult]{
		OutputState: o.OutputState,
	}
}

// The anti-affinity group ID to match (conflicts with `name`).
func (o LookupAntiAffinityGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAntiAffinityGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The list of attached exoscale*compute*instance (IDs).
func (o LookupAntiAffinityGroupResultOutput) Instances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupAntiAffinityGroupResult) []string { return v.Instances }).(pulumi.StringArrayOutput)
}

// The group name to match (conflicts with `id`).
func (o LookupAntiAffinityGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAntiAffinityGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAntiAffinityGroupResultOutput{})
}
