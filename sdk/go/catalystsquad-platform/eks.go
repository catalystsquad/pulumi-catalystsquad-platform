// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package catalystsquadplatform

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Eks struct {
	pulumi.ResourceState

	Cluster             eks.ClusterOutput               `pulumi:"cluster"`
	KubeConfig          pulumi.StringOutput             `pulumi:"kubeConfig"`
	KubernetesProvider  kubernetes.ProviderOutput       `pulumi:"kubernetesProvider"`
	NodeGroupIAMRoleArn pulumi.StringPtrOutput          `pulumi:"nodeGroupIAMRoleArn"`
	OidcProvider        iam.OpenIdConnectProviderOutput `pulumi:"oidcProvider"`
}

// NewEks registers a new resource with the given unique name, arguments, and options.
func NewEks(ctx *pulumi.Context,
	name string, args *EksArgs, opts ...pulumi.ResourceOption) (*Eks, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NodeGroupConfig == nil {
		return nil, errors.New("invalid value for required argument 'NodeGroupConfig'")
	}
	if args.SubnetIDs == nil {
		return nil, errors.New("invalid value for required argument 'SubnetIDs'")
	}
	opts = pkgResourceDefaultOpts(opts)
	var resource Eks
	err := ctx.RegisterRemoteComponentResource("catalystsquad-platform:index:Eks", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type eksArgs struct {
	// Optional, configures management of the eks auth configmap.
	AuthConfigmapConfig *AuthConfigMapConfig `pulumi:"authConfigmapConfig"`
	// Optional, cluster autoscaler namespace for IRSA. Default: cluster-autoscaler
	ClusterAutoscalerNamespace *string `pulumi:"clusterAutoscalerNamespace"`
	// Optional, cluster autoscaler service account name for IRSA. Default: cluster-autoscaler
	ClusterAutoscalerServiceAccount *string `pulumi:"clusterAutoscalerServiceAccount"`
	// Optional, name of the EKS cluster. Default: <stack name>
	ClusterName *string `pulumi:"clusterName"`
	// Optional, whether to enable cluster autoscaler IRSA resources. Default: true
	EnableClusterAutoscalerResources *bool `pulumi:"enableClusterAutoscalerResources"`
	// Optional, whether to enable ECR access policy on nodegroups. Default: true
	EnableECRAccess *bool `pulumi:"enableECRAccess"`
	// Optional, list of log types to enable on the cluster. Default: []
	EnabledClusterLogTypes *string `pulumi:"enabledClusterLogTypes"`
	// Optional, k8s version of the EKS cluster. Default: 1.22.6
	K8sVersion *string `pulumi:"k8sVersion"`
	// Optional, assume role arn to add to the kubeconfig.
	KubeConfigAssumeRoleArn *string `pulumi:"kubeConfigAssumeRoleArn"`
	// Optional, AWS profile to add to the kubeconfig.
	KubeConfigAwsProfile *string `pulumi:"kubeConfigAwsProfile"`
	// Required, list of nodegroup configurations to create.
	NodeGroupConfig []EksNodeGroup `pulumi:"nodeGroupConfig"`
	// Optional, k8s version of all node groups. Allows for upgrading the control plane before upgrading nodegroups. Default: <k8sVersion>
	NodeGroupVersion *string `pulumi:"nodeGroupVersion"`
	// Required, list of subnet IDs to deploy the cluster and nodegroups to
	SubnetIDs []string `pulumi:"subnetIDs"`
}

// The set of arguments for constructing a Eks resource.
type EksArgs struct {
	// Optional, configures management of the eks auth configmap.
	AuthConfigmapConfig AuthConfigMapConfigPtrInput
	// Optional, cluster autoscaler namespace for IRSA. Default: cluster-autoscaler
	ClusterAutoscalerNamespace pulumi.StringPtrInput
	// Optional, cluster autoscaler service account name for IRSA. Default: cluster-autoscaler
	ClusterAutoscalerServiceAccount pulumi.StringPtrInput
	// Optional, name of the EKS cluster. Default: <stack name>
	ClusterName pulumi.StringPtrInput
	// Optional, whether to enable cluster autoscaler IRSA resources. Default: true
	EnableClusterAutoscalerResources pulumi.BoolPtrInput
	// Optional, whether to enable ECR access policy on nodegroups. Default: true
	EnableECRAccess pulumi.BoolPtrInput
	// Optional, list of log types to enable on the cluster. Default: []
	EnabledClusterLogTypes pulumi.StringPtrInput
	// Optional, k8s version of the EKS cluster. Default: 1.22.6
	K8sVersion pulumi.StringPtrInput
	// Optional, assume role arn to add to the kubeconfig.
	KubeConfigAssumeRoleArn pulumi.StringPtrInput
	// Optional, AWS profile to add to the kubeconfig.
	KubeConfigAwsProfile pulumi.StringPtrInput
	// Required, list of nodegroup configurations to create.
	NodeGroupConfig EksNodeGroupArrayInput
	// Optional, k8s version of all node groups. Allows for upgrading the control plane before upgrading nodegroups. Default: <k8sVersion>
	NodeGroupVersion pulumi.StringPtrInput
	// Required, list of subnet IDs to deploy the cluster and nodegroups to
	SubnetIDs pulumi.StringArrayInput
}

func (EksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*eksArgs)(nil)).Elem()
}

type EksInput interface {
	pulumi.Input

	ToEksOutput() EksOutput
	ToEksOutputWithContext(ctx context.Context) EksOutput
}

func (*Eks) ElementType() reflect.Type {
	return reflect.TypeOf((**Eks)(nil)).Elem()
}

func (i *Eks) ToEksOutput() EksOutput {
	return i.ToEksOutputWithContext(context.Background())
}

func (i *Eks) ToEksOutputWithContext(ctx context.Context) EksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EksOutput)
}

// EksArrayInput is an input type that accepts EksArray and EksArrayOutput values.
// You can construct a concrete instance of `EksArrayInput` via:
//
//          EksArray{ EksArgs{...} }
type EksArrayInput interface {
	pulumi.Input

	ToEksArrayOutput() EksArrayOutput
	ToEksArrayOutputWithContext(context.Context) EksArrayOutput
}

type EksArray []EksInput

func (EksArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eks)(nil)).Elem()
}

func (i EksArray) ToEksArrayOutput() EksArrayOutput {
	return i.ToEksArrayOutputWithContext(context.Background())
}

func (i EksArray) ToEksArrayOutputWithContext(ctx context.Context) EksArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EksArrayOutput)
}

// EksMapInput is an input type that accepts EksMap and EksMapOutput values.
// You can construct a concrete instance of `EksMapInput` via:
//
//          EksMap{ "key": EksArgs{...} }
type EksMapInput interface {
	pulumi.Input

	ToEksMapOutput() EksMapOutput
	ToEksMapOutputWithContext(context.Context) EksMapOutput
}

type EksMap map[string]EksInput

func (EksMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eks)(nil)).Elem()
}

func (i EksMap) ToEksMapOutput() EksMapOutput {
	return i.ToEksMapOutputWithContext(context.Background())
}

func (i EksMap) ToEksMapOutputWithContext(ctx context.Context) EksMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EksMapOutput)
}

type EksOutput struct{ *pulumi.OutputState }

func (EksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Eks)(nil)).Elem()
}

func (o EksOutput) ToEksOutput() EksOutput {
	return o
}

func (o EksOutput) ToEksOutputWithContext(ctx context.Context) EksOutput {
	return o
}

type EksArrayOutput struct{ *pulumi.OutputState }

func (EksArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Eks)(nil)).Elem()
}

func (o EksArrayOutput) ToEksArrayOutput() EksArrayOutput {
	return o
}

func (o EksArrayOutput) ToEksArrayOutputWithContext(ctx context.Context) EksArrayOutput {
	return o
}

func (o EksArrayOutput) Index(i pulumi.IntInput) EksOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Eks {
		return vs[0].([]*Eks)[vs[1].(int)]
	}).(EksOutput)
}

type EksMapOutput struct{ *pulumi.OutputState }

func (EksMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Eks)(nil)).Elem()
}

func (o EksMapOutput) ToEksMapOutput() EksMapOutput {
	return o
}

func (o EksMapOutput) ToEksMapOutputWithContext(ctx context.Context) EksMapOutput {
	return o
}

func (o EksMapOutput) MapIndex(k pulumi.StringInput) EksOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Eks {
		return vs[0].(map[string]*Eks)[vs[1].(string)]
	}).(EksOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*EksInput)(nil)).Elem(), &Eks{})
	pulumi.RegisterInputType(reflect.TypeOf((*EksArrayInput)(nil)).Elem(), EksArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EksMapInput)(nil)).Elem(), EksMap{})
	pulumi.RegisterOutputType(EksOutput{})
	pulumi.RegisterOutputType(EksArrayOutput{})
	pulumi.RegisterOutputType(EksMapOutput{})
}
