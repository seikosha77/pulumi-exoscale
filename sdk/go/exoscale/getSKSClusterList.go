// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

func GetSKSClusterList(ctx *pulumi.Context, args *GetSKSClusterListArgs, opts ...pulumi.InvokeOption) (*GetSKSClusterListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSKSClusterListResult
	err := ctx.Invoke("exoscale:index/getSKSClusterList:getSKSClusterList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSKSClusterList.
type GetSKSClusterListArgs struct {
	AggregationCa  *string           `pulumi:"aggregationCa"`
	AutoUpgrade    *bool             `pulumi:"autoUpgrade"`
	Cni            *string           `pulumi:"cni"`
	ControlPlaneCa *string           `pulumi:"controlPlaneCa"`
	CreatedAt      *string           `pulumi:"createdAt"`
	Description    *string           `pulumi:"description"`
	Endpoint       *string           `pulumi:"endpoint"`
	ExoscaleCcm    *bool             `pulumi:"exoscaleCcm"`
	Id             *string           `pulumi:"id"`
	KubeletCa      *string           `pulumi:"kubeletCa"`
	Labels         map[string]string `pulumi:"labels"`
	MetricsServer  *bool             `pulumi:"metricsServer"`
	Name           *string           `pulumi:"name"`
	ServiceLevel   *string           `pulumi:"serviceLevel"`
	State          *string           `pulumi:"state"`
	Version        *string           `pulumi:"version"`
	Zone           string            `pulumi:"zone"`
}

// A collection of values returned by getSKSClusterList.
type GetSKSClusterListResult struct {
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	AggregationCa *string `pulumi:"aggregationCa"`
	// Match against this bool
	AutoUpgrade *bool                      `pulumi:"autoUpgrade"`
	Clusters    []GetSKSClusterListCluster `pulumi:"clusters"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Cni *string `pulumi:"cni"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ControlPlaneCa *string `pulumi:"controlPlaneCa"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	CreatedAt *string `pulumi:"createdAt"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Description *string `pulumi:"description"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Endpoint *string `pulumi:"endpoint"`
	// Match against this bool
	ExoscaleCcm *bool `pulumi:"exoscaleCcm"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Id *string `pulumi:"id"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	KubeletCa *string `pulumi:"kubeletCa"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Labels map[string]string `pulumi:"labels"`
	// Match against this bool
	MetricsServer *bool `pulumi:"metricsServer"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Name *string `pulumi:"name"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ServiceLevel *string `pulumi:"serviceLevel"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	State *string `pulumi:"state"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Version *string `pulumi:"version"`
	// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

func GetSKSClusterListOutput(ctx *pulumi.Context, args GetSKSClusterListOutputArgs, opts ...pulumi.InvokeOption) GetSKSClusterListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSKSClusterListResult, error) {
			args := v.(GetSKSClusterListArgs)
			r, err := GetSKSClusterList(ctx, &args, opts...)
			var s GetSKSClusterListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSKSClusterListResultOutput)
}

// A collection of arguments for invoking getSKSClusterList.
type GetSKSClusterListOutputArgs struct {
	AggregationCa  pulumi.StringPtrInput `pulumi:"aggregationCa"`
	AutoUpgrade    pulumi.BoolPtrInput   `pulumi:"autoUpgrade"`
	Cni            pulumi.StringPtrInput `pulumi:"cni"`
	ControlPlaneCa pulumi.StringPtrInput `pulumi:"controlPlaneCa"`
	CreatedAt      pulumi.StringPtrInput `pulumi:"createdAt"`
	Description    pulumi.StringPtrInput `pulumi:"description"`
	Endpoint       pulumi.StringPtrInput `pulumi:"endpoint"`
	ExoscaleCcm    pulumi.BoolPtrInput   `pulumi:"exoscaleCcm"`
	Id             pulumi.StringPtrInput `pulumi:"id"`
	KubeletCa      pulumi.StringPtrInput `pulumi:"kubeletCa"`
	Labels         pulumi.StringMapInput `pulumi:"labels"`
	MetricsServer  pulumi.BoolPtrInput   `pulumi:"metricsServer"`
	Name           pulumi.StringPtrInput `pulumi:"name"`
	ServiceLevel   pulumi.StringPtrInput `pulumi:"serviceLevel"`
	State          pulumi.StringPtrInput `pulumi:"state"`
	Version        pulumi.StringPtrInput `pulumi:"version"`
	Zone           pulumi.StringInput    `pulumi:"zone"`
}

func (GetSKSClusterListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSKSClusterListArgs)(nil)).Elem()
}

// A collection of values returned by getSKSClusterList.
type GetSKSClusterListResultOutput struct{ *pulumi.OutputState }

func (GetSKSClusterListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSKSClusterListResult)(nil)).Elem()
}

func (o GetSKSClusterListResultOutput) ToGetSKSClusterListResultOutput() GetSKSClusterListResultOutput {
	return o
}

func (o GetSKSClusterListResultOutput) ToGetSKSClusterListResultOutputWithContext(ctx context.Context) GetSKSClusterListResultOutput {
	return o
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) AggregationCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.AggregationCa }).(pulumi.StringPtrOutput)
}

// Match against this bool
func (o GetSKSClusterListResultOutput) AutoUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *bool { return v.AutoUpgrade }).(pulumi.BoolPtrOutput)
}

func (o GetSKSClusterListResultOutput) Clusters() GetSKSClusterListClusterArrayOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) []GetSKSClusterListCluster { return v.Clusters }).(GetSKSClusterListClusterArrayOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) Cni() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.Cni }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) ControlPlaneCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.ControlPlaneCa }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.Endpoint }).(pulumi.StringPtrOutput)
}

// Match against this bool
func (o GetSKSClusterListResultOutput) ExoscaleCcm() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *bool { return v.ExoscaleCcm }).(pulumi.BoolPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) KubeletCa() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.KubeletCa }).(pulumi.StringPtrOutput)
}

// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
func (o GetSKSClusterListResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Match against this bool
func (o GetSKSClusterListResultOutput) MetricsServer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *bool { return v.MetricsServer }).(pulumi.BoolPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSClusterListResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o GetSKSClusterListResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetSKSClusterListResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSKSClusterListResultOutput{})
}
