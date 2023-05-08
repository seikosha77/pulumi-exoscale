// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package exoscale

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSKSCluster(ctx *pulumi.Context, args *LookupSKSClusterArgs, opts ...pulumi.InvokeOption) (*LookupSKSClusterResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv LookupSKSClusterResult
	err := ctx.Invoke("exoscale:index/getSKSCluster:getSKSCluster", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getSKSCluster.
type LookupSKSClusterArgs struct {
	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons []string `pulumi:"addons"`
	// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
	AggregationCa *string `pulumi:"aggregationCa"`
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade *bool   `pulumi:"autoUpgrade"`
	Cni         *string `pulumi:"cni"`
	// The CA certificate (in PEM format) for TLS communications between control plane components.
	ControlPlaneCa *string `pulumi:"controlPlaneCa"`
	// The cluster creation date.
	CreatedAt *string `pulumi:"createdAt"`
	// A free-form text describing the cluster.
	Description *string `pulumi:"description"`
	// The cluster API endpoint.
	Endpoint *string `pulumi:"endpoint"`
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm *bool `pulumi:"exoscaleCcm"`
	// The ID of this resource.
	Id *string `pulumi:"id"`
	// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	KubeletCa *string `pulumi:"kubeletCa"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer *bool   `pulumi:"metricsServer"`
	Name          *string `pulumi:"name"`
	// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
	Nodepools []string `pulumi:"nodepools"`
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc *GetSKSClusterOidc `pulumi:"oidc"`
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel *string `pulumi:"serviceLevel"`
	// The cluster state.
	State *string `pulumi:"state"`
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version *string `pulumi:"version"`
	Zone    string  `pulumi:"zone"`
}

// A collection of values returned by getSKSCluster.
type LookupSKSClusterResult struct {
	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons []string `pulumi:"addons"`
	// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
	AggregationCa string `pulumi:"aggregationCa"`
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade *bool   `pulumi:"autoUpgrade"`
	Cni         *string `pulumi:"cni"`
	// The CA certificate (in PEM format) for TLS communications between control plane components.
	ControlPlaneCa string `pulumi:"controlPlaneCa"`
	// The cluster creation date.
	CreatedAt string `pulumi:"createdAt"`
	// A free-form text describing the cluster.
	Description *string `pulumi:"description"`
	// The cluster API endpoint.
	Endpoint string `pulumi:"endpoint"`
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm *bool `pulumi:"exoscaleCcm"`
	// The ID of this resource.
	Id *string `pulumi:"id"`
	// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	KubeletCa string `pulumi:"kubeletCa"`
	// A map of key/value labels.
	Labels map[string]string `pulumi:"labels"`
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer *bool   `pulumi:"metricsServer"`
	Name          *string `pulumi:"name"`
	// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
	Nodepools []string `pulumi:"nodepools"`
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc GetSKSClusterOidc `pulumi:"oidc"`
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel *string `pulumi:"serviceLevel"`
	// The cluster state.
	State string `pulumi:"state"`
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version string `pulumi:"version"`
	Zone    string `pulumi:"zone"`
}

func LookupSKSClusterOutput(ctx *pulumi.Context, args LookupSKSClusterOutputArgs, opts ...pulumi.InvokeOption) LookupSKSClusterResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSKSClusterResult, error) {
			args := v.(LookupSKSClusterArgs)
			r, err := LookupSKSCluster(ctx, &args, opts...)
			var s LookupSKSClusterResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSKSClusterResultOutput)
}

// A collection of arguments for invoking getSKSCluster.
type LookupSKSClusterOutputArgs struct {
	// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
	Addons pulumi.StringArrayInput `pulumi:"addons"`
	// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
	AggregationCa pulumi.StringPtrInput `pulumi:"aggregationCa"`
	// Enable automatic upgrading of the control plane version.
	AutoUpgrade pulumi.BoolPtrInput   `pulumi:"autoUpgrade"`
	Cni         pulumi.StringPtrInput `pulumi:"cni"`
	// The CA certificate (in PEM format) for TLS communications between control plane components.
	ControlPlaneCa pulumi.StringPtrInput `pulumi:"controlPlaneCa"`
	// The cluster creation date.
	CreatedAt pulumi.StringPtrInput `pulumi:"createdAt"`
	// A free-form text describing the cluster.
	Description pulumi.StringPtrInput `pulumi:"description"`
	// The cluster API endpoint.
	Endpoint pulumi.StringPtrInput `pulumi:"endpoint"`
	// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
	ExoscaleCcm pulumi.BoolPtrInput `pulumi:"exoscaleCcm"`
	// The ID of this resource.
	Id pulumi.StringPtrInput `pulumi:"id"`
	// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
	KubeletCa pulumi.StringPtrInput `pulumi:"kubeletCa"`
	// A map of key/value labels.
	Labels pulumi.StringMapInput `pulumi:"labels"`
	// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
	MetricsServer pulumi.BoolPtrInput   `pulumi:"metricsServer"`
	Name          pulumi.StringPtrInput `pulumi:"name"`
	// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
	Nodepools pulumi.StringArrayInput `pulumi:"nodepools"`
	// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
	Oidc GetSKSClusterOidcPtrInput `pulumi:"oidc"`
	// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
	ServiceLevel pulumi.StringPtrInput `pulumi:"serviceLevel"`
	// The cluster state.
	State pulumi.StringPtrInput `pulumi:"state"`
	// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
	Version pulumi.StringPtrInput `pulumi:"version"`
	Zone    pulumi.StringInput    `pulumi:"zone"`
}

func (LookupSKSClusterOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSKSClusterArgs)(nil)).Elem()
}

// A collection of values returned by getSKSCluster.
type LookupSKSClusterResultOutput struct{ *pulumi.OutputState }

func (LookupSKSClusterResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSKSClusterResult)(nil)).Elem()
}

func (o LookupSKSClusterResultOutput) ToLookupSKSClusterResultOutput() LookupSKSClusterResultOutput {
	return o
}

func (o LookupSKSClusterResultOutput) ToLookupSKSClusterResultOutputWithContext(ctx context.Context) LookupSKSClusterResultOutput {
	return o
}

// Deprecated: This attribute has been replaced by `exoscale_ccm`/`metrics_server` attributes, it will be removed in a future release.
func (o LookupSKSClusterResultOutput) Addons() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) []string { return v.Addons }).(pulumi.StringArrayOutput)
}

// The CA certificate (in PEM format) for TLS communications between the control plane and the aggregation layer (e.g. `metrics-server`).
func (o LookupSKSClusterResultOutput) AggregationCa() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.AggregationCa }).(pulumi.StringOutput)
}

// Enable automatic upgrading of the control plane version.
func (o LookupSKSClusterResultOutput) AutoUpgrade() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *bool { return v.AutoUpgrade }).(pulumi.BoolPtrOutput)
}

func (o LookupSKSClusterResultOutput) Cni() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *string { return v.Cni }).(pulumi.StringPtrOutput)
}

// The CA certificate (in PEM format) for TLS communications between control plane components.
func (o LookupSKSClusterResultOutput) ControlPlaneCa() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.ControlPlaneCa }).(pulumi.StringOutput)
}

// The cluster creation date.
func (o LookupSKSClusterResultOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.CreatedAt }).(pulumi.StringOutput)
}

// A free-form text describing the cluster.
func (o LookupSKSClusterResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

// The cluster API endpoint.
func (o LookupSKSClusterResultOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.Endpoint }).(pulumi.StringOutput)
}

// Deploy the Exoscale [Cloud Controller Manager](https://github.com/exoscale/exoscale-cloud-controller-manager/) in the control plane (boolean; default: `true`; may only be set at creation time).
func (o LookupSKSClusterResultOutput) ExoscaleCcm() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *bool { return v.ExoscaleCcm }).(pulumi.BoolPtrOutput)
}

// The ID of this resource.
func (o LookupSKSClusterResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

// The CA certificate (in PEM format) for TLS communications between kubelets and the control plane.
func (o LookupSKSClusterResultOutput) KubeletCa() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.KubeletCa }).(pulumi.StringOutput)
}

// A map of key/value labels.
func (o LookupSKSClusterResultOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

// Deploy the [Kubernetes Metrics Server](https://github.com/kubernetes-sigs/metrics-server/) in the control plane (boolean; default: `true`; may only be set at creation time).
func (o LookupSKSClusterResultOutput) MetricsServer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *bool { return v.MetricsServer }).(pulumi.BoolPtrOutput)
}

func (o LookupSKSClusterResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

// The list of exoscale*sks*nodepool (IDs) attached to the cluster.
func (o LookupSKSClusterResultOutput) Nodepools() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) []string { return v.Nodepools }).(pulumi.StringArrayOutput)
}

// An OpenID Connect configuration to provide to the Kubernetes API server (may only be set at creation time). Structure is documented below.
func (o LookupSKSClusterResultOutput) Oidc() GetSKSClusterOidcOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) GetSKSClusterOidc { return v.Oidc }).(GetSKSClusterOidcOutput)
}

// The service level of the control plane (`pro` or `starter`; default: `pro`; may only be set at creation time).
func (o LookupSKSClusterResultOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

// The cluster state.
func (o LookupSKSClusterResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.State }).(pulumi.StringOutput)
}

// The version of the control plane (default: latest version available from the API; see `exo compute sks versions` for reference; may only be set at creation time).
func (o LookupSKSClusterResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.Version }).(pulumi.StringOutput)
}

func (o LookupSKSClusterResultOutput) Zone() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSKSClusterResult) string { return v.Zone }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSKSClusterResultOutput{})
}