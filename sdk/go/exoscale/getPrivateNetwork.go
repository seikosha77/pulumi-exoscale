// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrivateNetwork(ctx *pulumi.Context, args *LookupPrivateNetworkArgs, opts ...pulumi.InvokeOption) (*LookupPrivateNetworkResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupPrivateNetworkResult
	err := ctx.Invoke("exoscale:index/getPrivateNetwork:getPrivateNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getPrivateNetwork.
type LookupPrivateNetworkArgs struct {
	// The private network description.
	Description *string `pulumi:"description"`
	// The private network ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The network name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getPrivateNetwork.
type LookupPrivateNetworkResult struct {
	// The private network description.
	Description *string `pulumi:"description"`
	// The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	EndIp string `pulumi:"endIp"`
	// The private network ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The network name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The network mask defining the IPv4 network allowed for static leases.
	Netmask string `pulumi:"netmask"`
	// The first/last IPv4 addresses used by the DHCP service for dynamic leases.
	StartIp string `pulumi:"startIp"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func LookupPrivateNetworkOutput(ctx *pulumi.Context, args LookupPrivateNetworkOutputArgs, opts ...pulumi.InvokeOption) LookupPrivateNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrivateNetworkResult, error) {
			args := v.(LookupPrivateNetworkArgs)
			r, err := LookupPrivateNetwork(ctx, &args, opts...)
			var s LookupPrivateNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrivateNetworkResultOutput)
}

// A collection of arguments for invoking getPrivateNetwork.
type LookupPrivateNetworkOutputArgs struct {
	// The private network description.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The private network ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The network name to match (conflicts with `id`).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupPrivateNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateNetworkArgs)(nil)).Elem()
}

// A collection of values returned by getPrivateNetwork.
type LookupPrivateNetworkResultOutput struct{ *pulumi.OutputState }

func (LookupPrivateNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrivateNetworkResult)(nil)).Elem()
}

func (o LookupPrivateNetworkResultOutput) ToLookupPrivateNetworkResultOutput() LookupPrivateNetworkResultOutput {
	return o
}

func (o LookupPrivateNetworkResultOutput) ToLookupPrivateNetworkResultOutputWithContext(ctx context.Context) LookupPrivateNetworkResultOutput {
	return o
}

// The private network description.
func (o LookupPrivateNetworkResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The first/last IPv4 addresses used by the DHCP service for dynamic leases.
func (o LookupPrivateNetworkResultOutput) EndIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.EndIp }).(pulumi.StringOutput)
}

// The private network ID to match (conflicts with `name`).
func (o LookupPrivateNetworkResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The network name to match (conflicts with `id`).
func (o LookupPrivateNetworkResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The network mask defining the IPv4 network allowed for static leases.
func (o LookupPrivateNetworkResultOutput) Netmask() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.Netmask }).(pulumi.StringOutput)
}

// The first/last IPv4 addresses used by the DHCP service for dynamic leases.
func (o LookupPrivateNetworkResultOutput) StartIp() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.StartIp }).(pulumi.StringOutput)
}

// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o LookupPrivateNetworkResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrivateNetworkResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrivateNetworkResultOutput{})
}
