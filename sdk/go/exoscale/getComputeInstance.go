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

func LookupComputeInstance(ctx *pulumi.Context, args *LookupComputeInstanceArgs, opts ...pulumi.InvokeOption) (*LookupComputeInstanceResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv LookupComputeInstanceResult
	err := ctx.Invoke("exoscale:index/getComputeInstance:getComputeInstance", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getComputeInstance.
type LookupComputeInstanceArgs struct {
	// The compute instance ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// The instance name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getComputeInstance.
type LookupComputeInstanceResult struct {
	// The list of attached exoscale*anti*affinity_group (IDs).
	AntiAffinityGroupIds []string `pulumi:"antiAffinityGroupIds"`
	// The compute instance creation date.
	CreatedAt string `pulumi:"createdAt"`
	// A deploy target ID.
	DeployTargetId string `pulumi:"deployTargetId"`
	// The instance disk size (GiB).
	DiskSize int `pulumi:"diskSize"`
	// The list of attached exoscale*elastic*ip (IDs).
	ElasticIpIds []string `pulumi:"elasticIpIds"`
	// The compute instance ID to match (conflicts with `name`).
	Id *string `pulumi:"id"`
	// Whether IPv6 is enabled on the instance.
	Ipv6 bool `pulumi:"ipv6"`
	// The instance (main network interface) IPv6 address (if enabled).
	Ipv6Address string `pulumi:"ipv6Address"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// The instance manager ID, if any.
	ManagerId string `pulumi:"managerId"`
	// The instance manager type (instance pool, SKS node pool, etc.), if any.
	ManagerType string `pulumi:"managerType"`
	// The instance name to match (conflicts with `id`).
	Name *string `pulumi:"name"`
	// The list of attached exoscale*private*network (IDs).
	PrivateNetworkIds []string `pulumi:"privateNetworkIds"`
	// The instance (main network interface) IPv4 address.
	PublicIpAddress string `pulumi:"publicIpAddress"`
	// Domain name for reverse DNS record.
	ReverseDns string `pulumi:"reverseDns"`
	// The list of attached exoscale*security*group (IDs).
	SecurityGroupIds []string `pulumi:"securityGroupIds"`
	// The exoscale*ssh*key (name) authorized on the instance.
	SshKey string `pulumi:"sshKey"`
	// The instance state.
	State string `pulumi:"state"`
	// The instance exoscale*compute*template ID.
	TemplateId string `pulumi:"templateId"`
	// The instance type.
	Type string `pulumi:"type"`
	// The instance [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
	UserData string `pulumi:"userData"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func LookupComputeInstanceOutput(ctx *pulumi.Context, args LookupComputeInstanceOutputArgs, opts ...pulumi.InvokeOption) LookupComputeInstanceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupComputeInstanceResult, error) {
			args := v.(LookupComputeInstanceArgs)
			r, err := LookupComputeInstance(ctx, &args, opts...)
			var s LookupComputeInstanceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupComputeInstanceResultOutput)
}

// A collection of arguments for invoking getComputeInstance.
type LookupComputeInstanceOutputArgs struct {
	// The compute instance ID to match (conflicts with `name`).
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The instance name to match (conflicts with `id`).
	Name pulumi.StringPtrInput `pulumi:"name"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (LookupComputeInstanceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeInstanceArgs)(nil)).Elem()
}

// A collection of values returned by getComputeInstance.
type LookupComputeInstanceResultOutput struct{ *pulumi.OutputState }

func (LookupComputeInstanceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupComputeInstanceResult)(nil)).Elem()
}

func (o LookupComputeInstanceResultOutput) ToLookupComputeInstanceResultOutput() LookupComputeInstanceResultOutput {
	return o
}

func (o LookupComputeInstanceResultOutput) ToLookupComputeInstanceResultOutputWithContext(ctx context.Context) LookupComputeInstanceResultOutput {
	return o
}

func (o LookupComputeInstanceResultOutput) ToOutput(ctx context.Context) pulumix.Output[LookupComputeInstanceResult] {
	return pulumix.Output[LookupComputeInstanceResult]{
		OutputState: o.OutputState,
	}
}

// The list of attached exoscale*anti*affinity_group (IDs).
func (o LookupComputeInstanceResultOutput) AntiAffinityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []string { return v.AntiAffinityGroupIds }).(pulumi.StringArrayOutput)
}

// The compute instance creation date.
func (o LookupComputeInstanceResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A deploy target ID.
func (o LookupComputeInstanceResultOutput) DeployTargetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.DeployTargetId }).(pulumi.StringOutput)
}

// The instance disk size (GiB).
func (o LookupComputeInstanceResultOutput) DiskSize() pulumi.IntOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) int { return v.DiskSize }).(pulumi.IntOutput)
}

// The list of attached exoscale*elastic*ip (IDs).
func (o LookupComputeInstanceResultOutput) ElasticIpIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []string { return v.ElasticIpIds }).(pulumi.StringArrayOutput)
}

// The compute instance ID to match (conflicts with `name`).
func (o LookupComputeInstanceResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Whether IPv6 is enabled on the instance.
func (o LookupComputeInstanceResultOutput) Ipv6() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) bool { return v.Ipv6 }).(pulumi.BoolOutput)
}

// The instance (main network interface) IPv6 address (if enabled).
func (o LookupComputeInstanceResultOutput) Ipv6Address() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Ipv6Address }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o LookupComputeInstanceResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// The instance manager ID, if any.
func (o LookupComputeInstanceResultOutput) ManagerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.ManagerId }).(pulumi.StringOutput)
}

// The instance manager type (instance pool, SKS node pool, etc.), if any.
func (o LookupComputeInstanceResultOutput) ManagerType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.ManagerType }).(pulumi.StringOutput)
}

// The instance name to match (conflicts with `id`).
func (o LookupComputeInstanceResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The list of attached exoscale*private*network (IDs).
func (o LookupComputeInstanceResultOutput) PrivateNetworkIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []string { return v.PrivateNetworkIds }).(pulumi.StringArrayOutput)
}

// The instance (main network interface) IPv4 address.
func (o LookupComputeInstanceResultOutput) PublicIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.PublicIpAddress }).(pulumi.StringOutput)
}

// Domain name for reverse DNS record.
func (o LookupComputeInstanceResultOutput) ReverseDns() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.ReverseDns }).(pulumi.StringOutput)
}

// The list of attached exoscale*security*group (IDs).
func (o LookupComputeInstanceResultOutput) SecurityGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) []string { return v.SecurityGroupIds }).(pulumi.StringArrayOutput)
}

// The exoscale*ssh*key (name) authorized on the instance.
func (o LookupComputeInstanceResultOutput) SshKey() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.SshKey }).(pulumi.StringOutput)
}

// The instance state.
func (o LookupComputeInstanceResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.State }).(pulumi.StringOutput)
}

// The instance exoscale*compute*template ID.
func (o LookupComputeInstanceResultOutput) TemplateId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.TemplateId }).(pulumi.StringOutput)
}

// The instance type.
func (o LookupComputeInstanceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Type }).(pulumi.StringOutput)
}

// The instance [cloud-init](http://cloudinit.readthedocs.io/en/latest/) configuration.
func (o LookupComputeInstanceResultOutput) UserData() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.UserData }).(pulumi.StringOutput)
}

// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o LookupComputeInstanceResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupComputeInstanceResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupComputeInstanceResultOutput{})
}
