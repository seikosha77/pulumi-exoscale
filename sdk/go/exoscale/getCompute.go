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

// !> **WARNING:** This data source is **DEPRECATED** and will be removed in the next major version. Please use ComputeInstance instead.
func LookupCompute(ctx *pulumi.Context, args *LookupComputeArgs, opts ...pulumi.InvokeOption) (*LookupComputeResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupComputeResult
	err := ctx.Invoke("exoscale:index/getCompute:getCompute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getCompute.
type LookupComputeArgs struct {
	// The instance hostname to match.
	Hostname *string `pulumi:"hostname"`
	// The compute instance ID to match.
	Id *string `pulumi:"id"`
	// The instance tags to match (map of key/value).
	Tags map[string]string `pulumi:"tags"`
}

// A collection of values returned by getCompute.
type LookupComputeResult struct {
	// Number of cpu the Compute instance is running with
	Cpu int `pulumi:"cpu"`
	// Date when the Compute instance was created
	Created string `pulumi:"created"`
	// Size of the Compute instance disk
	DiskSize int `pulumi:"diskSize"`
	// The instance hostname to match.
	Hostname *string `pulumi:"hostname"`
	// The compute instance ID to match.
	Id *string `pulumi:"id"`
	// Compute instance public ipv6 address (if ipv6 is enabled)
	Ip6Address string `pulumi:"ip6Address"`
	// Compute instance public ipv4 address
	IpAddress string `pulumi:"ipAddress"`
	// Memory allocated for the Compute instance
	Memory int `pulumi:"memory"`
	// List of Compute instance private IP addresses (in managed Private Networks only)
	PrivateNetworkIpAddresses []string `pulumi:"privateNetworkIpAddresses"`
	// Current size of the Compute instance
	Size string `pulumi:"size"`
	// State of the Compute instance
	State string `pulumi:"state"`
	// The instance tags to match (map of key/value).
	Tags map[string]string `pulumi:"tags"`
	// Name of the template for the Compute instance
	Template string `pulumi:"template"`
	// Name of the availability zone for the Compute instance
	Zone string `pulumi:"zone"`
}

func LookupComputeOutput(ctx *pulumi.Context, args LookupComputeOutputArgs, opts ...pulumi.InvokeOption) LookupComputeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComputeResult, error) {
			args := v.(LookupComputeArgs)
			r, err := LookupCompute(ctx, &args, opts...)
			var s LookupComputeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComputeResultOutput)
}

// A collection of arguments for invoking getCompute.
type LookupComputeOutputArgs struct {
	// The instance hostname to match.
	Hostname pulumi.StringPtrInput `pulumi:"hostname"`
	// The compute instance ID to match.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The instance tags to match (map of key/value).
	Tags pulumi.StringMapInput `pulumi:"tags"`
}

func (LookupComputeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeArgs)(nil)).Elem()
}

// A collection of values returned by getCompute.
type LookupComputeResultOutput struct{ *pulumi.OutputState }

func (LookupComputeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeResult)(nil)).Elem()
}

func (o LookupComputeResultOutput) ToLookupComputeResultOutput() LookupComputeResultOutput {
	return o
}

func (o LookupComputeResultOutput) ToLookupComputeResultOutputWithContext(ctx context.Context) LookupComputeResultOutput {
	return o
}

func (o LookupComputeResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupComputeResult] {
	return pulumix.Output[LookupComputeResult]{
		OutputState: o.OutputState,
	}
}

// Number of cpu the Compute instance is running with
func (o LookupComputeResultOutput) Cpu() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeResult) int { return v.Cpu }).(pulumi.IntOutput)
}

// Date when the Compute instance was created
func (o LookupComputeResultOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Created }).(pulumi.StringOutput)
}

// Size of the Compute instance disk
func (o LookupComputeResultOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeResult) int { return v.DiskSize }).(pulumi.IntOutput)
}

// The instance hostname to match.
func (o LookupComputeResultOutput) Hostname() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeResult) *string { return v.Hostname }).(pulumi.StringPtrOutput)
}

// The compute instance ID to match.
func (o LookupComputeResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Compute instance public ipv6 address (if ipv6 is enabled)
func (o LookupComputeResultOutput) Ip6Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Ip6Address }).(pulumi.StringOutput)
}

// Compute instance public ipv4 address
func (o LookupComputeResultOutput) IpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.IpAddress }).(pulumi.StringOutput)
}

// Memory allocated for the Compute instance
func (o LookupComputeResultOutput) Memory() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeResult) int { return v.Memory }).(pulumi.IntOutput)
}

// List of Compute instance private IP addresses (in managed Private Networks only)
func (o LookupComputeResultOutput) PrivateNetworkIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeResult) []string { return v.PrivateNetworkIpAddresses }).(pulumi.StringArrayOutput)
}

// Current size of the Compute instance
func (o LookupComputeResultOutput) Size() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Size }).(pulumi.StringOutput)
}

// State of the Compute instance
func (o LookupComputeResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.State }).(pulumi.StringOutput)
}

// The instance tags to match (map of key/value).
func (o LookupComputeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

// Name of the template for the Compute instance
func (o LookupComputeResultOutput) Template() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Template }).(pulumi.StringOutput)
}

// Name of the availability zone for the Compute instance
func (o LookupComputeResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeResultOutput{})
}
