// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

func LookupNLB(ctx *pulumi.Context, args *LookupNLBArgs, opts ...pulumi.InvokeOption) (*LookupNLBResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupNLBResult
	err := ctx.Invoke("exoscale:index/getNLB:getNLB", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getNLB.
type LookupNLBArgs struct {
	// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The NLB name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getNLB.
type LookupNLBResult struct {
	// The NLB creation date.
	CreatedAt string `pulumi:"createdAt"`
	// The Network Load Balancers (NLB) description.
	Description string `pulumi:"description"`
	// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The NLB public IPv4 address.
	IpAddress string `pulumi:"ipAddress"`
	// The NLB name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The current NLB state.
	State string `pulumi:"state"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func LookupNLBOutput(ctx *pulumi.Context, args LookupNLBOutputArgs, opts ...pulumi.InvokeOption) LookupNLBResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNLBResult, error) {
			args := v.(LookupNLBArgs)
			r, err := LookupNLB(ctx, &args, opts...)
			var s LookupNLBResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNLBResultOutput)
}

// A collection of arguments for invoking getNLB.
type LookupNLBOutputArgs struct {
	// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The NLB name to match (conflicts with `id`).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupNLBOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNLBArgs)(nil)).Elem()
}

// A collection of values returned by getNLB.
type LookupNLBResultOutput struct{ *pulumi.OutputState }

func (LookupNLBResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNLBResult)(nil)).Elem()
}

func (o LookupNLBResultOutput) ToLookupNLBResultOutput() LookupNLBResultOutput {
	return o
}

func (o LookupNLBResultOutput) ToLookupNLBResultOutputWithContext(ctx context.Context) LookupNLBResultOutput {
	return o
}

// The NLB creation date.
func (o LookupNLBResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNLBResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// The Network Load Balancers (NLB) description.
func (o LookupNLBResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNLBResult) string { return v.Description }).(pulumi.StringOutput)
}

// The Network Load Balancers (NLB) ID to match (conflicts with `name`).
func (o LookupNLBResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNLBResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The NLB public IPv4 address.
func (o LookupNLBResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNLBResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// The NLB name to match (conflicts with `id`).
func (o LookupNLBResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNLBResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The current NLB state.
func (o LookupNLBResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNLBResult) string { return v.State }).(pulumi.StringOutput)
}

// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o LookupNLBResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNLBResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNLBResultOutput{})
}
