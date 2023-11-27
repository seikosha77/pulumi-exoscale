// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumiverse/pulumi-exoscale/sdk/go/exoscale/internal"
)

func GetDatabaseURI(ctx *pulumi.Context, args *GetDatabaseURIArgs, opts ...pulumi.InvokeOption) (*GetDatabaseURIResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetDatabaseURIResult
	err := ctx.Invoke("exoscale:index/getDatabaseURI:getDatabaseURI", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getDatabaseURI.
type GetDatabaseURIArgs struct {
	// The database name to match.
	Name     string                  `pulumi:"name"`
	Timeouts *GetDatabaseURITimeouts `pulumi:"timeouts"`
	// The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
	Type string `pulumi:"type"`
	// The Exoscale Zone name.
	Zone string `pulumi:"zone"`
}

// A collection of values returned by getDatabaseURI.
type GetDatabaseURIResult struct {
	// The ID of this resource.
	Id string `pulumi:"id"`
	// The database name to match.
	Name     string                  `pulumi:"name"`
	Timeouts *GetDatabaseURITimeouts `pulumi:"timeouts"`
	// The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
	Type string `pulumi:"type"`
	// The database service connection URI.
	Uri string `pulumi:"uri"`
	// The Exoscale Zone name.
	Zone string `pulumi:"zone"`
}

func GetDatabaseURIOutput(ctx *pulumi.Context, args GetDatabaseURIOutputArgs, opts ...pulumi.InvokeOption) GetDatabaseURIResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetDatabaseURIResult, error) {
			args := v.(GetDatabaseURIArgs)
			r, err := GetDatabaseURI(ctx, &args, opts...)
			var s GetDatabaseURIResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetDatabaseURIResultOutput)
}

// A collection of arguments for invoking getDatabaseURI.
type GetDatabaseURIOutputArgs struct {
	// The database name to match.
	Name     pulumi.StringInput             `pulumi:"name"`
	Timeouts GetDatabaseURITimeoutsPtrInput `pulumi:"timeouts"`
	// The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
	Type pulumi.StringInput `pulumi:"type"`
	// The Exoscale Zone name.
	Zone pulumi.StringInput `pulumi:"zone"`
}

func (GetDatabaseURIOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseURIArgs)(nil)).Elem()
}

// A collection of values returned by getDatabaseURI.
type GetDatabaseURIResultOutput struct{ *pulumi.OutputState }

func (GetDatabaseURIResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetDatabaseURIResult)(nil)).Elem()
}

func (o GetDatabaseURIResultOutput) ToGetDatabaseURIResultOutput() GetDatabaseURIResultOutput {
	return o
}

func (o GetDatabaseURIResultOutput) ToGetDatabaseURIResultOutputWithContext(ctx context.Context) GetDatabaseURIResultOutput {
	return o
}

// The ID of this resource.
func (o GetDatabaseURIResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseURIResult) string { return v.Id }).(pulumi.StringOutput)
}

// The database name to match.
func (o GetDatabaseURIResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseURIResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetDatabaseURIResultOutput) Timeouts() GetDatabaseURITimeoutsPtrOutput {
	return o.ApplyT(func(v GetDatabaseURIResult) *GetDatabaseURITimeouts { return v.Timeouts }).(GetDatabaseURITimeoutsPtrOutput)
}

// The type of the database service (`kafka`, `mysql`, `opensearch`, `pg`, `redis`).
func (o GetDatabaseURIResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseURIResult) string { return v.Type }).(pulumi.StringOutput)
}

// The database service connection URI.
func (o GetDatabaseURIResultOutput) Uri() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseURIResult) string { return v.Uri }).(pulumi.StringOutput)
}

// The Exoscale Zone name.
func (o GetDatabaseURIResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v GetDatabaseURIResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetDatabaseURIResultOutput{})
}
