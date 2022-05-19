package observability

import (
	"encoding/json"
	"fmt"

	"github.com/catalystsquad/pulumi-catalystsquad-platform/internal/utils/roles"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ObservabilityDependenciesArgs struct {
	// Required, Arn of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.
	OidcProviderArn pulumi.StringInput `pulumi:"oidcProviderArn"`
	// Required, URL of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.
	OidcProviderURL pulumi.StringInput `pulumi:"oidcProviderUrl"`
	// Optional, name of bucket to create for Cortex.
	// Default: <account-id>-<stack-name>-cortex
	CortexBucketName string `pulumi:"cortexBucketName"`
	// Optional, Cortex's IAM policy name. Default: cortex-policy
	CortexIAMPolicyName string `pulumi:"cortexIAMPolicyName"`
	// Optional, Cortex's IAM role name. Default: cortex-role
	CortexIAMRoleName string `pulumi:"cortexIAMRoleName"`
	// Optional, kubernetes namespace where Cortex will exist, for configuring
	// the IRSA IAM role trust relationship. Default: cortex
	CortexNamespace string `pulumi:"cortexNamespace"`
	// Optional, kubernetes service account name that Cortex will use, for
	// configuring the IRSA IAM role trust relationship. Default: cortex
	CortexServiceAccount string `pulumi:"cortexServiceAccount"`
	// Optional, name of bucket to create for Loki.
	// Default: <account-id>-<stack-name>-loki
	LokiBucketName string `pulumi:"lokiBucketName"`
	// Optional, Loki's IAM policy name. Default: loki-policy
	LokiIAMPolicyName string `pulumi:"LokiIAMPolicyName"`
	// Optional, Loki's IAM role name. Default: loki-role
	LokiIAMRoleName string `pulumi:"lokiIAMRoleName"`
	// Optional, kubernetes namespace where Loki will exist, for configuring
	// the IRSA IAM role trust relationship. Default: loki
	LokiNamespace string `pulumi:"lokiNamespace"`
	// Optional, kubernetes service account name that Loki will use, for
	// configuring the IRSA IAM role trust relationship. Default: loki
	LokiServiceAccount string `pulumi:"lokiServiceAccount"`
}

type ObservabilityDependencies struct {
	pulumi.ResourceState
}

func NewObservabilityDependencies(ctx *pulumi.Context, name string, args *ObservabilityDependenciesArgs,
	opts ...pulumi.ResourceOption) (*ObservabilityDependencies, error) {

	if args == nil {
		args = &ObservabilityDependenciesArgs{}
	}

	component := &ObservabilityDependencies{}
	err := ctx.RegisterComponentResource("catalystsquad-platform:index:ObservabilityDependencies",
		name, component, opts...)
	if err != nil {
		return nil, err
	}

	// retrieve account information for uniquely naming resources
	current, err := aws.GetCallerIdentity(ctx, nil, nil)
	if err != nil {
		return nil, err
	}

	// use the stack name for naming resources if overrides are not supplied
	stackName := ctx.Stack()

	// set defaults for optional parameters for cortex
	cortexNamespace := "cortex"
	if args.CortexNamespace != "" {
		cortexNamespace = args.CortexNamespace
	}
	cortexIAMPolicyName := "cortex-policy"
	if args.CortexIAMPolicyName != "" {
		cortexIAMPolicyName = args.CortexIAMPolicyName
	}
	cortexIAMRoleName := "cortex-role"
	if args.CortexIAMRoleName != "" {
		cortexIAMRoleName = args.CortexIAMRoleName
	}
	cortexServiceAccount := "cortex"
	if args.CortexServiceAccount != "" {
		cortexServiceAccount = args.CortexServiceAccount
	}
	cortexBucketName := fmt.Sprintf("%s-%s-cortex", current.AccountId, stackName)
	if args.CortexBucketName != "" {
		cortexBucketName = args.CortexBucketName
	}

	// set defaults for optional parameters for loki
	lokiNamespace := "loki"
	if args.LokiNamespace != "" {
		lokiNamespace = args.LokiNamespace
	}
	lokiIAMPolicyName := "loki-policy"
	if args.LokiIAMPolicyName != "" {
		lokiIAMPolicyName = args.LokiIAMPolicyName
	}
	lokiIAMRoleName := "loki-role"
	if args.LokiIAMRoleName != "" {
		lokiIAMRoleName = args.LokiIAMRoleName
	}
	lokiServiceAccount := "loki"
	if args.LokiServiceAccount != "" {
		lokiServiceAccount = args.LokiServiceAccount
	}
	lokiBucketName := fmt.Sprintf("%s-%s-loki", current.AccountId, stackName)
	if args.CortexBucketName != "" {
		lokiBucketName = args.LokiBucketName
	}

	// create all cortex aws resources
	_, err = createObservabilityBucket(ctx, "cortex-bucket", cortexBucketName, component)
	if err != nil {
		return nil, err
	}

	cortexPolicyBytes, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Action": []string{
					"s3:ListBucket",
					"s3:PutObject",
					"s3:GetObject",
					"s3:DeleteObject",
				},
				"Effect": "Allow",
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s", cortexBucketName),
					fmt.Sprintf("arn:aws:s3:::%s/*", cortexBucketName),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	err = roles.CreateRoleWithPolicy(
		ctx, "cortex", cortexIAMPolicyName, cortexIAMRoleName, string(cortexPolicyBytes),
		roles.CreateIrsaAssumeRolePolicy(args.OidcProviderArn, args.OidcProviderURL,
			cortexNamespace, cortexServiceAccount),
		"Grants access to Cortex S3 Bucket", component,
	)
	if err != nil {
		return nil, err
	}

	// create all loki aws resources
	_, err = createObservabilityBucket(ctx, "loki-bucket", lokiBucketName, component)
	if err != nil {
		return nil, err
	}

	lokiPolicyBytes, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Action": []string{
					"s3:ListBucket",
					"s3:PutObject",
					"s3:GetObject",
					"s3:DeleteObject",
				},
				"Effect": "Allow",
				"Resource": []string{
					fmt.Sprintf("arn:aws:s3:::%s", lokiBucketName),
					fmt.Sprintf("arn:aws:s3:::%s/*", lokiBucketName),
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}

	err = roles.CreateRoleWithPolicy(
		ctx, "loki", lokiIAMPolicyName, lokiIAMRoleName, string(lokiPolicyBytes),
		roles.CreateIrsaAssumeRolePolicy(args.OidcProviderArn, args.OidcProviderURL,
			lokiNamespace, lokiServiceAccount),
		"Grants access to Loki S3 Bucket", component,
	)
	if err != nil {
		return nil, err
	}

	return component, nil
}

func createObservabilityBucket(ctx *pulumi.Context, resourceName string,
	bucketName string, parent pulumi.Resource) (pulumi.Resource, error) {

	bucket, err := s3.NewBucket(ctx, resourceName, &s3.BucketArgs{
		Bucket: pulumi.String(bucketName),
		Acl:    pulumi.String("private"),
		ServerSideEncryptionConfiguration: &s3.BucketServerSideEncryptionConfigurationArgs{
			Rule: &s3.BucketServerSideEncryptionConfigurationRuleArgs{
				ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
					SseAlgorithm: pulumi.String("AES256"),
				},
			},
		},
	}, pulumi.Parent(parent))
	return bucket, err
}
