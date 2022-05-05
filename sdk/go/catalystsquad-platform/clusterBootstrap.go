// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package catalystsquadplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ClusterBootstrap struct {
	pulumi.ResourceState
}

// NewClusterBootstrap registers a new resource with the given unique name, arguments, and options.
func NewClusterBootstrap(ctx *pulumi.Context,
	name string, args *ClusterBootstrapArgs, opts ...pulumi.ResourceOption) (*ClusterBootstrap, error) {
	if args == nil {
		args = &ClusterBootstrapArgs{}
	}

	opts = pkgResourceDefaultOpts(opts)
	var resource ClusterBootstrap
	err := ctx.RegisterRemoteComponentResource("catalystsquad-platform:index:ClusterBootstrap", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type clusterBootstrapArgs struct {
	// Optional, configures the argocd helm release.
	ArgocdHelmConfig *HelmReleaseConfig `pulumi:"argocdHelmConfig"`
	// Optional, configures management of the eks auth configmap. Does not manage the configmap if not specified.
	EksAuthConfigmapConfig *AuthConfigMapConfig `pulumi:"eksAuthConfigmapConfig"`
	// Optional, configures the kube-prometheus-stack helm release.
	KubePrometheusStackHelmConfig *HelmReleaseConfig `pulumi:"kubePrometheusStackHelmConfig"`
	// Optional, configures the platform application release. Does not deploy if not specified.
	PlatformApplicationConfig *PlatformApplicationConfig `pulumi:"platformApplicationConfig"`
	// Optional, configuration for a prometheus remoteWrite secret. Does not deploy if not specified.
	PrometheusRemoteWriteConfig *PrometheusRemoteWriteConfig `pulumi:"prometheusRemoteWriteConfig"`
}

// The set of arguments for constructing a ClusterBootstrap resource.
type ClusterBootstrapArgs struct {
	// Optional, configures the argocd helm release.
	ArgocdHelmConfig HelmReleaseConfigPtrInput
	// Optional, configures management of the eks auth configmap. Does not manage the configmap if not specified.
	EksAuthConfigmapConfig AuthConfigMapConfigPtrInput
	// Optional, configures the kube-prometheus-stack helm release.
	KubePrometheusStackHelmConfig HelmReleaseConfigPtrInput
	// Optional, configures the platform application release. Does not deploy if not specified.
	PlatformApplicationConfig PlatformApplicationConfigPtrInput
	// Optional, configuration for a prometheus remoteWrite secret. Does not deploy if not specified.
	PrometheusRemoteWriteConfig PrometheusRemoteWriteConfigPtrInput
}

func (ClusterBootstrapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*clusterBootstrapArgs)(nil)).Elem()
}

type ClusterBootstrapInput interface {
	pulumi.Input

	ToClusterBootstrapOutput() ClusterBootstrapOutput
	ToClusterBootstrapOutputWithContext(ctx context.Context) ClusterBootstrapOutput
}

func (*ClusterBootstrap) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterBootstrap)(nil)).Elem()
}

func (i *ClusterBootstrap) ToClusterBootstrapOutput() ClusterBootstrapOutput {
	return i.ToClusterBootstrapOutputWithContext(context.Background())
}

func (i *ClusterBootstrap) ToClusterBootstrapOutputWithContext(ctx context.Context) ClusterBootstrapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBootstrapOutput)
}

// ClusterBootstrapArrayInput is an input type that accepts ClusterBootstrapArray and ClusterBootstrapArrayOutput values.
// You can construct a concrete instance of `ClusterBootstrapArrayInput` via:
//
//          ClusterBootstrapArray{ ClusterBootstrapArgs{...} }
type ClusterBootstrapArrayInput interface {
	pulumi.Input

	ToClusterBootstrapArrayOutput() ClusterBootstrapArrayOutput
	ToClusterBootstrapArrayOutputWithContext(context.Context) ClusterBootstrapArrayOutput
}

type ClusterBootstrapArray []ClusterBootstrapInput

func (ClusterBootstrapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterBootstrap)(nil)).Elem()
}

func (i ClusterBootstrapArray) ToClusterBootstrapArrayOutput() ClusterBootstrapArrayOutput {
	return i.ToClusterBootstrapArrayOutputWithContext(context.Background())
}

func (i ClusterBootstrapArray) ToClusterBootstrapArrayOutputWithContext(ctx context.Context) ClusterBootstrapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBootstrapArrayOutput)
}

// ClusterBootstrapMapInput is an input type that accepts ClusterBootstrapMap and ClusterBootstrapMapOutput values.
// You can construct a concrete instance of `ClusterBootstrapMapInput` via:
//
//          ClusterBootstrapMap{ "key": ClusterBootstrapArgs{...} }
type ClusterBootstrapMapInput interface {
	pulumi.Input

	ToClusterBootstrapMapOutput() ClusterBootstrapMapOutput
	ToClusterBootstrapMapOutputWithContext(context.Context) ClusterBootstrapMapOutput
}

type ClusterBootstrapMap map[string]ClusterBootstrapInput

func (ClusterBootstrapMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterBootstrap)(nil)).Elem()
}

func (i ClusterBootstrapMap) ToClusterBootstrapMapOutput() ClusterBootstrapMapOutput {
	return i.ToClusterBootstrapMapOutputWithContext(context.Background())
}

func (i ClusterBootstrapMap) ToClusterBootstrapMapOutputWithContext(ctx context.Context) ClusterBootstrapMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterBootstrapMapOutput)
}

type ClusterBootstrapOutput struct{ *pulumi.OutputState }

func (ClusterBootstrapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterBootstrap)(nil)).Elem()
}

func (o ClusterBootstrapOutput) ToClusterBootstrapOutput() ClusterBootstrapOutput {
	return o
}

func (o ClusterBootstrapOutput) ToClusterBootstrapOutputWithContext(ctx context.Context) ClusterBootstrapOutput {
	return o
}

type ClusterBootstrapArrayOutput struct{ *pulumi.OutputState }

func (ClusterBootstrapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ClusterBootstrap)(nil)).Elem()
}

func (o ClusterBootstrapArrayOutput) ToClusterBootstrapArrayOutput() ClusterBootstrapArrayOutput {
	return o
}

func (o ClusterBootstrapArrayOutput) ToClusterBootstrapArrayOutputWithContext(ctx context.Context) ClusterBootstrapArrayOutput {
	return o
}

func (o ClusterBootstrapArrayOutput) Index(i pulumi.IntInput) ClusterBootstrapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ClusterBootstrap {
		return vs[0].([]*ClusterBootstrap)[vs[1].(int)]
	}).(ClusterBootstrapOutput)
}

type ClusterBootstrapMapOutput struct{ *pulumi.OutputState }

func (ClusterBootstrapMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ClusterBootstrap)(nil)).Elem()
}

func (o ClusterBootstrapMapOutput) ToClusterBootstrapMapOutput() ClusterBootstrapMapOutput {
	return o
}

func (o ClusterBootstrapMapOutput) ToClusterBootstrapMapOutputWithContext(ctx context.Context) ClusterBootstrapMapOutput {
	return o
}

func (o ClusterBootstrapMapOutput) MapIndex(k pulumi.StringInput) ClusterBootstrapOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ClusterBootstrap {
		return vs[0].(map[string]*ClusterBootstrap)[vs[1].(string)]
	}).(ClusterBootstrapOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBootstrapInput)(nil)).Elem(), &ClusterBootstrap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBootstrapArrayInput)(nil)).Elem(), ClusterBootstrapArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClusterBootstrapMapInput)(nil)).Elem(), ClusterBootstrapMap{})
	pulumi.RegisterOutputType(ClusterBootstrapOutput{})
	pulumi.RegisterOutputType(ClusterBootstrapArrayOutput{})
	pulumi.RegisterOutputType(ClusterBootstrapMapOutput{})
}
