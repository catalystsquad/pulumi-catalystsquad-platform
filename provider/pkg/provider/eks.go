package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// EksArgs supplies input for configuring EKS
type EksArgs struct {
	// Optional, name of the EKS cluster. Default: <stack name>
	ClusterName string `pulumi:"clusterName"`
	// Optional, k8s version of the EKS cluster. Default: 1.22.6
	K8sVersion string `pulumi:"k8sVersion"`
	// Optional, k8s version of all node groups. Allows for upgrading the
	// control plane before upgrading nodegroups. Default: <k8sVersion>
	NodeGroupVersion string `pulumi:"nodeGroupVersion"`
	// Required, list of nodegroup configurations to create.
	NodeGroupConfig []EksNodeGroup `pulumi:"nodeGroupConfig"`
	// Optional, whether to enable ECR access policy on nodegroups. Default: true
	EnableECRAccess *bool `pulumi:"enableECRAccess"`
	// Optional, whether to enable cluster autoscaler IRSA resources. Default: true
	EnableClusterAutoscalerResources *bool `pulumi:"enableClusterAutoscalerResources"`
	// Optional, cluster autoscaler service account name for IRSA. Default: cluster-autoscaler
	ClusterAutoscalerServiceAccount string `pulumi:"clusterAutoscalerServiceAccount"`
	// Optional, cluster autoscaler namespace for IRSA. Default: cluster-autoscaler
	ClusterAutoscalerNamespace string `pulumi:"clusterAutoscalerNamespace"`
	// Optional, list of log types to enable on the cluster. Default: []
	EnabledClusterLogTypes []string `pulumi:"enabledClusterLogTypes"`
	// Required, list of subnet IDs to deploy the cluster and nodegroups to
	SubnetIDs pulumi.StringArrayInput `pulumi:"subnetIDs"`
}

// EksNodeGroup allows configuring multiple nodegroups
type EksNodeGroup struct {
	// Required, name prefix of the nodegroup
	NamePrefix string `pulumi:"namePrefix"`
	// Required, initial desired size of nodegroup, ignored after creation
	DesiredSize int `pulumi:"desiredSize"`
	// Required, maximum size of nodegroup
	MaxSize int `pulumi:"maxSize"`
	// Required, minimum size of nodegroup
	MinSize int `pulumi:"minSize"`
	// Required, list of instance types for the nodegroup
	InstanceTypes []string `pulumi:"instanceTypes"`
}

// Eks pulumi component resource
type Eks struct {
	pulumi.ResourceState

	EksCluster   *eks.Cluster               `pulumi:"eksCluster"`
	OidcProvider *iam.OpenIdConnectProvider `pulumi:"oidcProvider"`
}

// https://github.com/hashicorp/terraform-provider-aws/issues/10104#issuecomment-545264374
// TODO: generate this instead
var awsRootCAThumbprint string = "9e99a48a9960b14926bb7f3b02e22da2b0ab7280"

// NewEks creates an EKS cluster
func NewEks(ctx *pulumi.Context, name string, args *EksArgs, opts ...pulumi.ResourceOption) (*Eks, error) {
	if args == nil {
		args = &EksArgs{}
	}

	component := &Eks{}
	err := ctx.RegisterComponentResource("catalystsquad-platform:index:Eks", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// throw an error if we don't have required arguments
	if args.SubnetIDs == nil {
		return component, errors.New("Missing SubnetID argument")
	}

	// default eks arguments
	clusterName := ctx.Stack()
	if args.ClusterName != "" {
		clusterName = args.ClusterName
	}

	k8sVersion := "1.22.6"
	if args.K8sVersion != "" {
		k8sVersion = args.K8sVersion
	}

	nodeGroupVersion := k8sVersion
	if args.NodeGroupVersion != "" {
		nodeGroupVersion = args.NodeGroupVersion
	}

	enableECRAccess := true
	if args.EnableECRAccess != nil {
		enableECRAccess = *args.EnableECRAccess
	}

	enableClusterAutoscalerResources := true
	if args.EnableClusterAutoscalerResources != nil {
		enableClusterAutoscalerResources = *args.EnableClusterAutoscalerResources
	}

	clusterAutoscalerServiceAccount := "cluster-autoscaler"
	if args.ClusterAutoscalerServiceAccount != "" {
		clusterAutoscalerServiceAccount = args.ClusterAutoscalerServiceAccount
	}

	clusterAutoscalerNamespace := "cluster-autoscaler"
	if args.ClusterAutoscalerNamespace != "" {
		clusterAutoscalerNamespace = args.ClusterAutoscalerNamespace
	}

	// create eks service role
	eksServiceRole, err := iam.NewRole(ctx, "eks-service-role", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
			"Version": "2008-10-17",
			"Statement": [{
				"Sid": "",
				"Effect": "Allow",
				"Principal": {
					"Service": "eks.amazonaws.com"
				},
				"Action": "sts:AssumeRole"
			}]
		}`),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// attach aws managed policies to service role
	eksPolicyArns := []string{
		"arn:aws:iam::aws:policy/AmazonEKSServicePolicy",
		"arn:aws:iam::aws:policy/AmazonEKSClusterPolicy",
	}
	for _, policyArn := range eksPolicyArns {
		policyName := strings.TrimPrefix(policyArn, "arn:aws:iam::aws:policy/")
		_, err := iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("eks-service-role-%s-policy-attachment", policyName), &iam.RolePolicyAttachmentArgs{
			Role:      eksServiceRole.Name,
			PolicyArn: pulumi.String(policyArn),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// create default nodegroup role
	nodeGroupRole, err := iam.NewRole(ctx, "nodegroup-role", &iam.RoleArgs{
		AssumeRolePolicy: pulumi.String(`{
			"Version": "2012-10-17",
			"Statement": [{
				"Sid": "",
				"Effect": "Allow",
				"Principal": {
					"Service": "ec2.amazonaws.com"
				},
				"Action": "sts:AssumeRole"
			}]
		}`),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// attach aws managed nodegroup policies
	nodeGroupPolicyArns := []string{
		"arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy",
		"arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy",
		"arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly",
	}
	for _, policyArn := range nodeGroupPolicyArns {
		policyName := strings.TrimPrefix(policyArn, "arn:aws:iam::aws:policy/")
		_, err := iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("nodegroup-role-%s-policy-attachment", policyName), &iam.RolePolicyAttachmentArgs{
			Role:      nodeGroupRole.Name,
			PolicyArn: pulumi.String(policyArn),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	if enableECRAccess {
		// ecr access policy
		ecrAccessPolicy, err := iam.NewPolicy(ctx, "ecr-access-policy", &iam.PolicyArgs{
			Name:        pulumi.String(fmt.Sprintf("ecr-access-policy-%s", clusterName)),
			Description: pulumi.String("Grants access to ECR"),
			Policy: pulumi.String(`{
			"Version": "2012-10-17",
			"Statement": [
				{
					"Effect": "Allow",
					"Action": [
						"ecr:BatchCheckLayerAvailability",
						"ecr:BatchGetImage",
						"ecr:GetDownloadUrlForLayer",
						"ecr:GetAuthorizationToken"
					],
					"Resource": "*"
				}
			]
		}`),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "nodegroup-ecr-policy-attachment", &iam.RolePolicyAttachmentArgs{
			Role:      nodeGroupRole.Name,
			PolicyArn: ecrAccessPolicy.Arn,
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// create eks cluster
	eksCluster, err := eks.NewCluster(ctx, "eks-cluster", &eks.ClusterArgs{
		Name:                   pulumi.String(clusterName),
		Version:                pulumi.String(k8sVersion),
		RoleArn:                pulumi.StringInput(eksServiceRole.Arn),
		EnabledClusterLogTypes: pulumi.ToStringArray(args.EnabledClusterLogTypes),
		VpcConfig: &eks.ClusterVpcConfigArgs{
			SubnetIds:            args.SubnetIDs,
			EndpointPublicAccess: pulumi.Bool(true),
			PublicAccessCidrs: pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
			},
		},
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// create node groups
	var nodeGroups []pulumi.Resource
	for _, nodeGroupConfig := range args.NodeGroupConfig {
		nodeGroup, err := eks.NewNodeGroup(ctx, fmt.Sprintf("node-group-%s", nodeGroupConfig.NamePrefix), &eks.NodeGroupArgs{
			ClusterName:         eksCluster.Name,
			InstanceTypes:       pulumi.ToStringArray(nodeGroupConfig.InstanceTypes),
			NodeGroupNamePrefix: pulumi.String(nodeGroupConfig.NamePrefix),
			NodeRoleArn:         pulumi.StringInput(nodeGroupRole.Arn),
			SubnetIds:           args.SubnetIDs,
			Version:             pulumi.String(nodeGroupVersion),
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: pulumi.Int(nodeGroupConfig.DesiredSize),
				MaxSize:     pulumi.Int(nodeGroupConfig.MaxSize),
				MinSize:     pulumi.Int(nodeGroupConfig.MinSize),
			},
		}, pulumi.Parent(component), pulumi.IgnoreChanges([]string{"scalingConfig.desiredSize"}))
		if err != nil {
			return nil, err
		}

		nodeGroups = append(nodeGroups, nodeGroup)
	}

	// create oidc provider for IRSA https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
	oidcProvider, err := iam.NewOpenIdConnectProvider(ctx, "eks-oidc-provider", &iam.OpenIdConnectProviderArgs{
		ClientIdLists:   pulumi.StringArray{pulumi.String("sts.amazonaws.com")},
		ThumbprintLists: pulumi.StringArray{pulumi.String(awsRootCAThumbprint)},
		Url:             eksCluster.Identities.Index(pulumi.Int(0)).Oidcs().Index(pulumi.Int(0)).Issuer().Elem(),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	if enableClusterAutoscalerResources {
		// create cluster autoscaler iam policy
		clusterAutoscalerPolicyJSON, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				// allow read only actions
				{
					"Action": []string{
						"autoscaling:DescribeAutoScalingGroups",
						"autoscaling:DescribeAutoScalingInstances",
						"autoscaling:DescribeLaunchConfigurations",
						"autoscaling:DescribeTags",
						"ec2:DescribeLaunchTemplateVersions",
						"ec2:DescribeInstanceTypes",
					},
					"Effect":   "Allow",
					"Resource": "*",
				},
				// allow autoscaling for only this specific eks cluster
				{
					"Action": []string{
						"autoscaling:SetDesiredCapacity",
						"autoscaling:TerminateInstanceInAutoScalingGroup",
						"autoscaling:UpdateAutoScalingGroup",
					},
					"Effect":   "Allow",
					"Resource": "*",
					"Condition": map[string]interface{}{
						"StringEquals": map[string]string{
							fmt.Sprintf("autoscaling:ResourceTag/kubernetes.io/cluster/%s", clusterName): "owned",
						},
					},
				},
			},
		})
		if err != nil {
			return nil, err
		}

		clusterAutoscalerPolicy, err := iam.NewPolicy(ctx, "cluster-autoscaler-policy", &iam.PolicyArgs{
			Name:        pulumi.String(fmt.Sprintf("cluster-autoscaler-policy-%s", clusterName)),
			Description: pulumi.String(fmt.Sprintf("cluster autoscaler policy for %s eks cluster", clusterName)),
			Policy:      pulumi.String(clusterAutoscalerPolicyJSON),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}

		// create cluster autoscaler iam role with IRSA
		clusterAutoscalerRole, err := iam.NewRole(ctx, "cluster-autoscaler-role", &iam.RoleArgs{
			Name:             pulumi.String(fmt.Sprintf("cluster-autoscaler-role-%s", clusterName)),
			AssumeRolePolicy: createIrsaAssumeRolePolicy(oidcProvider, clusterAutoscalerNamespace, clusterAutoscalerServiceAccount),
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "cluster-autoscaler-role-policy-attachment", &iam.RolePolicyAttachmentArgs{
			Role:      clusterAutoscalerRole.Name,
			PolicyArn: clusterAutoscalerPolicy.Arn,
		}, pulumi.Parent(component))
		if err != nil {
			return nil, err
		}
	}

	// set outputs for component
	component.EksCluster = eksCluster
	component.OidcProvider = oidcProvider

	err = ctx.RegisterResourceOutputs(component, pulumi.Map{
		"eksCluster":   eksCluster,
		"oidcProvider": oidcProvider,
	})
	if err != nil {
		return nil, err
	}

	return component, nil
}
