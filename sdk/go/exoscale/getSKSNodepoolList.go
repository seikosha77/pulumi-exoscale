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

func GetSKSNodepoolList(ctx *pulumi.Context, args *GetSKSNodepoolListArgs, opts ...pulumi.InvokeOption) (*GetSKSNodepoolListResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetSKSNodepoolListResult
	err := ctx.Invoke("exoscale:index/getSKSNodepoolList:getSKSNodepoolList", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSKSNodepoolList.
type GetSKSNodepoolListArgs struct {
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ClusterId *string `pulumi:"clusterId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	CreatedAt *string `pulumi:"createdAt"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Description *string `pulumi:"description"`
	// Match against this int
	DiskSize *int `pulumi:"diskSize"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Id *string `pulumi:"id"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstancePoolId *string `pulumi:"instancePoolId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstancePrefix *string `pulumi:"instancePrefix"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstanceType *string `pulumi:"instanceType"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Labels map[string]string `pulumi:"labels"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Name *string `pulumi:"name"`
	// Match against this int
	Size *int `pulumi:"size"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	State *string `pulumi:"state"`
	// Match against this bool
	StorageLvm *bool `pulumi:"storageLvm"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Taints map[string]string `pulumi:"taints"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	TemplateId *string `pulumi:"templateId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Version *string `pulumi:"version"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Zone *string `pulumi:"zone"`
}

// A collection of values returned by getSKSNodepoolList.
type GetSKSNodepoolListResult struct {
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ClusterId *string `pulumi:"clusterId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	CreatedAt *string `pulumi:"createdAt"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	DeployTargetId *string `pulumi:"deployTargetId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Description *string `pulumi:"description"`
	// Match against this int
	DiskSize *int `pulumi:"diskSize"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Id *string `pulumi:"id"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstancePoolId *string `pulumi:"instancePoolId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstancePrefix *string `pulumi:"instancePrefix"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstanceType *string `pulumi:"instanceType"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Labels map[string]string `pulumi:"labels"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Name      *string                      `pulumi:"name"`
	Nodepools []GetSKSNodepoolListNodepool `pulumi:"nodepools"`
	// Match against this int
	Size *int `pulumi:"size"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	State *string `pulumi:"state"`
	// Match against this bool
	StorageLvm *bool `pulumi:"storageLvm"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Taints map[string]string `pulumi:"taints"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	TemplateId *string `pulumi:"templateId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Version *string `pulumi:"version"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Zone *string `pulumi:"zone"`
}

func GetSKSNodepoolListOutput(ctx *pulumi.Context, args GetSKSNodepoolListOutputArgs, opts ...pulumi.InvokeOption) GetSKSNodepoolListResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetSKSNodepoolListResult, error) {
			args := v.(GetSKSNodepoolListArgs)
			r, err := GetSKSNodepoolList(ctx, &args, opts...)
			var s GetSKSNodepoolListResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetSKSNodepoolListResultOutput)
}

// A collection of arguments for invoking getSKSNodepoolList.
type GetSKSNodepoolListOutputArgs struct {
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	ClusterId pulumi.StringPtrInput `pulumi:"clusterId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	DeployTargetId pulumi.StringPtrInput `pulumi:"deployTargetId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// Match against this int
	DiskSize pulumi.IntPtrInput `pulumi:"diskSize"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstancePoolId pulumi.StringPtrInput `pulumi:"instancePoolId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstancePrefix pulumi.StringPtrInput `pulumi:"instancePrefix"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	InstanceType pulumi.StringPtrInput `pulumi:"instanceType"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Name pulumi.StringPtrInput `pulumi:"name"`
	// Match against this int
	Size pulumi.IntPtrInput `pulumi:"size"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	State pulumi.StringPtrInput `pulumi:"state"`
	// Match against this bool
	StorageLvm pulumi.BoolPtrInput `pulumi:"storageLvm"`
	// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
	Taints pulumi.StringMapInput `pulumi:"taints"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	TemplateId pulumi.StringPtrInput `pulumi:"templateId"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Version pulumi.StringPtrInput `pulumi:"version"`
	// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
	Zone pulumi.StringPtrInput `pulumi:"zone"`
}

func (GetSKSNodepoolListOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSKSNodepoolListArgs)(nil)).Elem()
}

// A collection of values returned by getSKSNodepoolList.
type GetSKSNodepoolListResultOutput struct{ *pulumi.OutputState }

func (GetSKSNodepoolListResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetSKSNodepoolListResult)(nil)).Elem()
}

func (o GetSKSNodepoolListResultOutput) ToGetSKSNodepoolListResultOutput() GetSKSNodepoolListResultOutput {
	return o
}

func (o GetSKSNodepoolListResultOutput) ToGetSKSNodepoolListResultOutputWithContext(ctx context.Context) GetSKSNodepoolListResultOutput {
	return o
}

func (o GetSKSNodepoolListResultOutput) ToOutput(ctx context.Context) pulumix.Output[GetSKSNodepoolListResult] {
	return pulumix.Output[GetSKSNodepoolListResult]{
		OutputState: o.OutputState,
	}
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) ClusterId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.ClusterId }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) DeployTargetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.DeployTargetId }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// Match against this int
func (o GetSKSNodepoolListResultOutput) DiskSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *int { return v.DiskSize }).(pulumi.IntPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) InstancePoolId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.InstancePoolId }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) InstancePrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.InstancePrefix }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) InstanceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.InstanceType }).(pulumi.StringPtrOutput)
}

// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
func (o GetSKSNodepoolListResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetSKSNodepoolListResultOutput) Nodepools() GetSKSNodepoolListNodepoolArrayOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) []GetSKSNodepoolListNodepool { return v.Nodepools }).(GetSKSNodepoolListNodepoolArrayOutput)
}

// Match against this int
func (o GetSKSNodepoolListResultOutput) Size() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *int { return v.Size }).(pulumi.IntPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

// Match against this bool
func (o GetSKSNodepoolListResultOutput) StorageLvm() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *bool { return v.StorageLvm }).(pulumi.BoolPtrOutput)
}

// Match against key/values. Keys are matched exactly, while values may be matched as a regex if you supply a string that begins and ends with "/"
func (o GetSKSNodepoolListResultOutput) Taints() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) map[string]string { return v.Taints }).(pulumi.StringMapOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) TemplateId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.TemplateId }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// Match against this string. If you supply a string that begins and ends with a "/" it will be matched as a regex.
func (o GetSKSNodepoolListResultOutput) Zone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetSKSNodepoolListResult) *string { return v.Zone }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetSKSNodepoolListResultOutput{})
}
