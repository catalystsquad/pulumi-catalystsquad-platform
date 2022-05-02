package provider

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pkg/errors"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/samber/lo"
)

// EksArgs supplies input for configuring EKS
type EksArgs struct {
	// TODO document options, add pulumi tags
	ClusterName      string         `pulumi:"clusterName"`
	K8sVersion       string         `pulumi:"k8sVersion"`
	NodeGroupVersion string         `pulumi:"nodeGroupVersion"`
	NodeGroupConfig  []EksNodeGroup `pulumi:"nodeGroupConfig"`

	// optional
	EnableECRAccess *bool `pulumi:"enableECRAccess"`

	// optional cluster autoscaler IRSA configuration
	EnableClusterAutoscalerResources *bool  `pulumi:"enableClusterAutoscalerResources"`
	ClusterAutoscalerServiceAccount  string `pulumi:"clusterAutoscalerServiceAccount"`
	ClusterAutoscalerNamespace       string `pulumi:"clusterAutoscalerNamespace"`

	// optional
	EnabledClusterLogTypes []string `pulumi:"enabledClusterLogTypes"`

	// required
	SubnetIDs []string `pulumi:"subnetIDs"`
}

// EksNodeGroup allows configuring multiple nodegroups
type EksNodeGroup struct {
	NamePrefix    string   `pulumi:"namePrefix"`
	DesiredSize   int      `pulumi:"desiredSize"`
	MaxSize       int      `pulumi:"maxSize"`
	MinSize       int      `pulumi:"minSize"`
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
	subnetIDs := args.SubnetIDs

	// default eks arguments
	clusterName := lo.Ternary(args.ClusterName == "", ctx.Stack(), args.ClusterName)
	k8sVersion := lo.Ternary(args.K8sVersion == "", "1.22.6", args.K8sVersion)
	nodeGroupVersion := lo.Ternary(args.NodeGroupVersion == "", k8sVersion, args.NodeGroupVersion)
	enableECRAccess := lo.Ternary(args.EnableECRAccess == nil, true, *args.EnableECRAccess)
	enableClusterAutoscalerResources := lo.Ternary(args.EnableClusterAutoscalerResources == nil, true, *args.EnableClusterAutoscalerResources)
	clusterAutoscalerServiceAccount := lo.Ternary(args.ClusterAutoscalerServiceAccount == "", "cluster-autoscaler", args.ClusterAutoscalerServiceAccount)
	clusterAutoscalerNamespace := lo.Ternary(args.ClusterAutoscalerNamespace == "", "cluster-autoscaler", args.ClusterAutoscalerNamespace)

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
	})
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
		})
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
	})
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
		})
		if err != nil {
			return nil, err
		}
	}

	if enableECRAccess {
		// ecr access policy
		ecrAccessPolicy, err := iam.NewPolicy(ctx, "policy", &iam.PolicyArgs{
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
		})
		if err != nil {
			return nil, err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "nodegroup-ecr-policy-attachment", &iam.RolePolicyAttachmentArgs{
			Role:      nodeGroupRole.Name,
			PolicyArn: ecrAccessPolicy.Arn,
		})
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
			SubnetIds:            pulumi.ToStringArray(subnetIDs),
			EndpointPublicAccess: pulumi.Bool(true),
			PublicAccessCidrs: pulumi.StringArray{
				pulumi.String("0.0.0.0/0"),
			},
		},
	})
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
			SubnetIds:           pulumi.ToStringArray(subnetIDs),
			Version:             pulumi.String(nodeGroupVersion),
			ScalingConfig: &eks.NodeGroupScalingConfigArgs{
				DesiredSize: pulumi.Int(nodeGroupConfig.DesiredSize),
				MaxSize:     pulumi.Int(nodeGroupConfig.MaxSize),
				MinSize:     pulumi.Int(nodeGroupConfig.MinSize),
			},
		}, pulumi.IgnoreChanges([]string{"scalingConfig.desiredSize"}))
		if err != nil {
			return nil, err
		}

		nodeGroups = append(nodeGroups, nodeGroup)
	}

	// create oidc provider for IRSA https://docs.aws.amazon.com/eks/latest/userguide/iam-roles-for-service-accounts.html
	oidcProvider, err := iam.NewOpenIdConnectProvider(ctx, "eks-oidc-provider", &iam.OpenIdConnectProviderArgs{
		ClientIdLists:   pulumi.StringArray{pulumi.String("sts.amazonaws.com")},
		ThumbprintLists: pulumi.StringArray{pulumi.String(awsRootCAThumbprint)},
		Url:             eksCluster.Identities.Index(pulumi.Int(0)).Oidcs().Index(pulumi.Int(0)).Issuer().Elem(), // what the fuck
	})
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
		})
		if err != nil {
			return nil, err
		}

		// create cluster autoscaler iam role with IRSA
		clusterAutoscalerRole, err := iam.NewRole(ctx, "cluster-autoscaler-role", &iam.RoleArgs{
			Name:             pulumi.String(fmt.Sprintf("cluster-autoscaler-role-%s", clusterName)),
			AssumeRolePolicy: createIrsaAssumeRolePolicy(oidcProvider, clusterAutoscalerNamespace, clusterAutoscalerServiceAccount),
		})
		if err != nil {
			return nil, err
		}
		_, err = iam.NewRolePolicyAttachment(ctx, "cluster-autoscaler-role-policy-attachment", &iam.RolePolicyAttachmentArgs{
			Role:      clusterAutoscalerRole.Name,
			PolicyArn: clusterAutoscalerPolicy.Arn,
		})
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
