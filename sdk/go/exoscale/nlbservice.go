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

// ## Import
//
// An existing NLB service may be imported by `<nlb-ID>/<service-ID>@<zone>`
//
// ```sh
//
//	$ pulumi import exoscale:index/nLBService:NLBService \
//
// ```
//
//	exoscale_nlb_service.my_nlb_service \
//
//	f81d4fae-7dec-11d0-a765-00a0c91e6bf6/9ecc6b8b-73d4-4211-8ced-f7f29bb79524@ch-gva-2
type NLBService struct {
	pulumi.CustomResourceState

	// A free-form text describing the NLB service.
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The service health checking configuration.
	Healthchecks NLBServiceHealthcheckArrayOutput `pulumi:"healthchecks"`
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId pulumi.StringOutput `pulumi:"instancePoolId"`
	// The NLB service name.
	Name pulumi.StringOutput `pulumi:"name"`
	// ❗ The parent NLB ID.
	NlbId pulumi.StringOutput `pulumi:"nlbId"`
	// The NLB service (TCP/UDP) port.
	Port pulumi.IntOutput `pulumi:"port"`
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol pulumi.StringPtrOutput `pulumi:"protocol"`
	State    pulumi.StringOutput    `pulumi:"state"`
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy pulumi.StringPtrOutput `pulumi:"strategy"`
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort pulumi.IntOutput `pulumi:"targetPort"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringOutput `pulumi:"zone"`
}

// NewNLBService registers a new resource with the given unique name, arguments, and options.
func NewNLBService(ctx *pulumi.Context,
	name string, args *NLBServiceArgs, opts ...pulumi.ResourceOption) (*NLBService, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Healthchecks == nil {
		return nil, errors.New("invalid value for required argument 'Healthchecks'")
	}
	if args.InstancePoolId == nil {
		return nil, errors.New("invalid value for required argument 'InstancePoolId'")
	}
	if args.NlbId == nil {
		return nil, errors.New("invalid value for required argument 'NlbId'")
	}
	if args.Port == nil {
		return nil, errors.New("invalid value for required argument 'Port'")
	}
	if args.TargetPort == nil {
		return nil, errors.New("invalid value for required argument 'TargetPort'")
	}
	if args.Zone == nil {
		return nil, errors.New("invalid value for required argument 'Zone'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource NLBService
	err := ctx.RegisterResource("exoscale:index/nLBService:NLBService", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetNLBService gets an existing NLBService resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetNLBService(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *NLBServiceState, opts ...pulumi.ResourceOption) (*NLBService, error) {
	var resource NLBService
	err := ctx.ReadResource("exoscale:index/nLBService:NLBService", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering NLBService resources.
type nlbserviceState struct {
	// A free-form text describing the NLB service.
	Description *string `pulumi:"description"`
	// The service health checking configuration.
	Healthchecks []NLBServiceHealthcheck `pulumi:"healthchecks"`
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId *string `pulumi:"instancePoolId"`
	// The NLB service name.
	Name *string `pulumi:"name"`
	// ❗ The parent NLB ID.
	NlbId *string `pulumi:"nlbId"`
	// The NLB service (TCP/UDP) port.
	Port *int `pulumi:"port"`
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol *string `pulumi:"protocol"`
	State    *string `pulumi:"state"`
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy *string `pulumi:"strategy"`
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort *int `pulumi:"targetPort"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone *string `pulumi:"zone"`
}

type NLBServiceState struct {
	// A free-form text describing the NLB service.
	Description pulumi.StringPtrInput
	// The service health checking configuration.
	Healthchecks NLBServiceHealthcheckArrayInput
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId pulumi.StringPtrInput
	// The NLB service name.
	Name pulumi.StringPtrInput
	// ❗ The parent NLB ID.
	NlbId pulumi.StringPtrInput
	// The NLB service (TCP/UDP) port.
	Port pulumi.IntPtrInput
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol pulumi.StringPtrInput
	State    pulumi.StringPtrInput
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy pulumi.StringPtrInput
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort pulumi.IntPtrInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringPtrInput
}

func (NLBServiceState) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbserviceState)(nil)).Elem()
}

type nlbserviceArgs struct {
	// A free-form text describing the NLB service.
	Description *string `pulumi:"description"`
	// The service health checking configuration.
	Healthchecks []NLBServiceHealthcheck `pulumi:"healthchecks"`
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId string `pulumi:"instancePoolId"`
	// The NLB service name.
	Name *string `pulumi:"name"`
	// ❗ The parent NLB ID.
	NlbId string `pulumi:"nlbId"`
	// The NLB service (TCP/UDP) port.
	Port int `pulumi:"port"`
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol *string `pulumi:"protocol"`
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy *string `pulumi:"strategy"`
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort int `pulumi:"targetPort"`
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone string `pulumi:"zone"`
}

// The set of arguments for constructing a NLBService resource.
type NLBServiceArgs struct {
	// A free-form text describing the NLB service.
	Description pulumi.StringPtrInput
	// The service health checking configuration.
	Healthchecks NLBServiceHealthcheckArrayInput
	// ❗ The exoscale*instance*pool (ID) to forward traffic to.
	InstancePoolId pulumi.StringInput
	// The NLB service name.
	Name pulumi.StringPtrInput
	// ❗ The parent NLB ID.
	NlbId pulumi.StringInput
	// The NLB service (TCP/UDP) port.
	Port pulumi.IntInput
	// The protocol (`tcp`|`udp`; default: `tcp`).
	Protocol pulumi.StringPtrInput
	// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
	Strategy pulumi.StringPtrInput
	// The (TCP/UDP) port to forward traffic to (on target instance pool members).
	TargetPort pulumi.IntInput
	// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
	Zone pulumi.StringInput
}

func (NLBServiceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*nlbserviceArgs)(nil)).Elem()
}

type NLBServiceInput interface {
	pulumi.Input

	ToNLBServiceOutput() NLBServiceOutput
	ToNLBServiceOutputWithContext(ctx context.Context) NLBServiceOutput
}

func (*NLBService) ElementType() reflect.Type {
	return reflect.TypeOf((**NLBService)(nil)).Elem()
}

func (i *NLBService) ToNLBServiceOutput() NLBServiceOutput {
	return i.ToNLBServiceOutputWithContext(context.Background())
}

func (i *NLBService) ToNLBServiceOutputWithContext(ctx context.Context) NLBServiceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NLBServiceOutput)
}

// NLBServiceArrayInput is an input type that accepts NLBServiceArray and NLBServiceArrayOutput values.
// You can construct a concrete instance of `NLBServiceArrayInput` via:
//
//	NLBServiceArray{ NLBServiceArgs{...} }
type NLBServiceArrayInput interface {
	pulumi.Input

	ToNLBServiceArrayOutput() NLBServiceArrayOutput
	ToNLBServiceArrayOutputWithContext(context.Context) NLBServiceArrayOutput
}

type NLBServiceArray []NLBServiceInput

func (NLBServiceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NLBService)(nil)).Elem()
}

func (i NLBServiceArray) ToNLBServiceArrayOutput() NLBServiceArrayOutput {
	return i.ToNLBServiceArrayOutputWithContext(context.Background())
}

func (i NLBServiceArray) ToNLBServiceArrayOutputWithContext(ctx context.Context) NLBServiceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NLBServiceArrayOutput)
}

// NLBServiceMapInput is an input type that accepts NLBServiceMap and NLBServiceMapOutput values.
// You can construct a concrete instance of `NLBServiceMapInput` via:
//
//	NLBServiceMap{ "key": NLBServiceArgs{...} }
type NLBServiceMapInput interface {
	pulumi.Input

	ToNLBServiceMapOutput() NLBServiceMapOutput
	ToNLBServiceMapOutputWithContext(context.Context) NLBServiceMapOutput
}

type NLBServiceMap map[string]NLBServiceInput

func (NLBServiceMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NLBService)(nil)).Elem()
}

func (i NLBServiceMap) ToNLBServiceMapOutput() NLBServiceMapOutput {
	return i.ToNLBServiceMapOutputWithContext(context.Background())
}

func (i NLBServiceMap) ToNLBServiceMapOutputWithContext(ctx context.Context) NLBServiceMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NLBServiceMapOutput)
}

type NLBServiceOutput struct{ *pulumi.OutputState }

func (NLBServiceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NLBService)(nil)).Elem()
}

func (o NLBServiceOutput) ToNLBServiceOutput() NLBServiceOutput {
	return o
}

func (o NLBServiceOutput) ToNLBServiceOutputWithContext(ctx context.Context) NLBServiceOutput {
	return o
}

// A free-form text describing the NLB service.
func (o NLBServiceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The service health checking configuration.
func (o NLBServiceOutput) Healthchecks() NLBServiceHealthcheckArrayOutput {
	return o.ApplyT(func(v *NLBService) NLBServiceHealthcheckArrayOutput { return v.Healthchecks }).(NLBServiceHealthcheckArrayOutput)
}

// ❗ The exoscale*instance*pool (ID) to forward traffic to.
func (o NLBServiceOutput) InstancePoolId() pulumi.StringOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringOutput { return v.InstancePoolId }).(pulumi.StringOutput)
}

// The NLB service name.
func (o NLBServiceOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// ❗ The parent NLB ID.
func (o NLBServiceOutput) NlbId() pulumi.StringOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringOutput { return v.NlbId }).(pulumi.StringOutput)
}

// The NLB service (TCP/UDP) port.
func (o NLBServiceOutput) Port() pulumi.IntOutput {
	return o.ApplyT(func(v *NLBService) pulumi.IntOutput { return v.Port }).(pulumi.IntOutput)
}

// The protocol (`tcp`|`udp`; default: `tcp`).
func (o NLBServiceOutput) Protocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringPtrOutput { return v.Protocol }).(pulumi.StringPtrOutput)
}

func (o NLBServiceOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

// The strategy (`round-robin`|`source-hash`; default: `round-robin`).
func (o NLBServiceOutput) Strategy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringPtrOutput { return v.Strategy }).(pulumi.StringPtrOutput)
}

// The (TCP/UDP) port to forward traffic to (on target instance pool members).
func (o NLBServiceOutput) TargetPort() pulumi.IntOutput {
	return o.ApplyT(func(v *NLBService) pulumi.IntOutput { return v.TargetPort }).(pulumi.IntOutput)
}

// ❗ The Exoscale [Zone](https://www.exoscale.com/datacenters/) name.
func (o NLBServiceOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v *NLBService) pulumi.StringOutput { return v.Zone }).(pulumi.StringOutput)
}

type NLBServiceArrayOutput struct{ *pulumi.OutputState }

func (NLBServiceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*NLBService)(nil)).Elem()
}

func (o NLBServiceArrayOutput) ToNLBServiceArrayOutput() NLBServiceArrayOutput {
	return o
}

func (o NLBServiceArrayOutput) ToNLBServiceArrayOutputWithContext(ctx context.Context) NLBServiceArrayOutput {
	return o
}

func (o NLBServiceArrayOutput) Index(i pulumi.IntInput) NLBServiceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *NLBService {
		return vs[0].([]*NLBService)[vs[1].(int)]
	}).(NLBServiceOutput)
}

type NLBServiceMapOutput struct{ *pulumi.OutputState }

func (NLBServiceMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*NLBService)(nil)).Elem()
}

func (o NLBServiceMapOutput) ToNLBServiceMapOutput() NLBServiceMapOutput {
	return o
}

func (o NLBServiceMapOutput) ToNLBServiceMapOutputWithContext(ctx context.Context) NLBServiceMapOutput {
	return o
}

func (o NLBServiceMapOutput) MapIndex(k pulumi.StringInput) NLBServiceOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *NLBService {
		return vs[0].(map[string]*NLBService)[vs[1].(string)]
	}).(NLBServiceOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*NLBServiceInput)(nil)).Elem(), &NLBService{})
	pulumi.RegisterInputType(reflect.TypeOf((*NLBServiceArrayInput)(nil)).Elem(), NLBServiceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*NLBServiceMapInput)(nil)).Elem(), NLBServiceMap{})
	pulumi.RegisterOutputType(NLBServiceOutput{})
	pulumi.RegisterOutputType(NLBServiceArrayOutput{})
	pulumi.RegisterOutputType(NLBServiceMapOutput{})
}
