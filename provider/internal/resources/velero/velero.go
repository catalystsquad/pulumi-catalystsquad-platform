package velero

import (
	"encoding/json"
	"fmt"

	"github.com/catalystsquad/pulumi-catalystsquad-platform/internal/utils/roles"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/s3"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VeleroDependenciesArgs struct {
	// Required, Arn of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.
	OidcProviderArn pulumi.StringInput `pulumi:"oidcProviderArn"`
	// Required, URL of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.
	OidcProviderURL pulumi.StringInput `pulumi:"oidcProviderUrl"`
	// Optional, whether to create the Velero S3 bucket. Allows the bucket to
	// exist outside of pulumi. Default: true
	CreateBucket *bool `pulumi:"createBucket"`
	// Optional, Velero's bucket name. Default: <account-id>-<stack-name>-velero
	VeleroBucketName string `pulumi:"veleroBucketName"`
	// Optional, Velero's IAM policy name. Default: <stack-name>-velero-policy
	VeleroIAMPolicyName string `pulumi:"veleroIAMPolicyName"`
	// Optional, Velero's IAM role name. Default: <stack-name>-velero-role
	VeleroIAMRoleName string `pulumi:"veleroIAMRoleName"`
	// Optional, kubernetes namespace where Velero will exist, for configuring
	// the IRSA IAM role trust relationship. Default: velero
	VeleroNamespace string `pulumi:"veleroNamespace"`
	// Optional, kubernetes service account name that Velero will use, for
	// configuring the IRSA IAM role trust relationship. Default: velero
	VeleroServiceAccount string `pulumi:"veleroServiceAccount"`
}

type VeleroDependencies struct {
	pulumi.ResourceState
}

func NewVeleroDependencies(ctx *pulumi.Context, name string, args *VeleroDependenciesArgs,
	opts ...pulumi.ResourceOption) (*VeleroDependencies, error) {

	if args == nil {
		args = &VeleroDependenciesArgs{}
	}

	component := &VeleroDependencies{}
	err := ctx.RegisterComponentResource("catalystsquad-platform:index:VeleroDependencies",
		name, component, opts...)
	if err != nil {
		return nil, err
	}

	// used to create generic bucket names if a bucket name override is not
	// specified
	current, err := aws.GetCallerIdentity(ctx, nil, nil)
	// used to
	stackName := ctx.Stack()

	// default velero specific configuration if not supplied
	createBucket := true
	if args.CreateBucket != nil {
		createBucket = *args.CreateBucket
	}
	veleroBucketName := fmt.Sprintf("%s-%s-velero", current.AccountId, stackName)
	if args.VeleroBucketName != "" {
		veleroBucketName = args.VeleroBucketName
	}
	veleroIAMPolicyName := fmt.Sprintf("%s-velero-policy", stackName)
	if args.VeleroIAMPolicyName != "" {
		veleroIAMPolicyName = args.VeleroIAMPolicyName
	}
	veleroIAMRoleName := fmt.Sprintf("%s-velero-role", stackName)
	if args.VeleroIAMRoleName != "" {
		veleroIAMRoleName = args.VeleroIAMRoleName
	}
	veleroServiceAccount := "velero"
	if args.VeleroServiceAccount != "" {
		veleroServiceAccount = args.VeleroServiceAccount
	}
	veleroNamespace := "velero"
	if args.VeleroNamespace != "" {
		veleroNamespace = args.VeleroNamespace
	}

	// create velero iam policy
	veleroPolicyJSON, err := json.Marshal(map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []map[string]interface{}{
			{
				"Action": []string{
					"ec2:DescribeVolumes",
					"ec2:DescribeSnapshots",
					"ec2:CreateTags",
					"ec2:CreateVolume",
					"ec2:CreateSnapshot",
					"ec2:DeleteSnapshot",
				},
				"Effect":   "Allow",
				"Resource": "*",
			},
			{
				"Action": []string{
					"s3:GetObject",
					"s3:DeleteObject",
					"s3:PutObject",
					"s3:AbortMultipartUpload",
					"s3:ListMultipartUploadParts",
				},
				"Effect":   "Allow",
				"Resource": fmt.Sprintf("arn:aws:s3:::%s/*", veleroBucketName),
			},
			{
				"Action": []string{
					"s3:ListBucket",
				},
				"Effect":   "Allow",
				"Resource": fmt.Sprintf("arn:aws:s3:::%s", veleroBucketName),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	err = roles.CreateRoleWithPolicy(
		ctx, "velero", veleroIAMPolicyName, veleroIAMRoleName, string(veleroPolicyJSON),
		roles.CreateIrsaAssumeRolePolicy(args.OidcProviderArn, args.OidcProviderURL,
			veleroNamespace, veleroServiceAccount),
		"Velero S3 object and EC2 EBS snapshot access", component,
	)
	if err != nil {
		return nil, err
	}

	// create velero s3 bucket, optional to allow the bucket to exist outside
	// of pulumi. if the stack is ever destroyed, we want the backup data to
	// continue to exist. configuration allows disabling creation of the bucket
	// so that recreating a stack won't conflict with the existing bucket
	if createBucket {
		_, err = s3.NewBucket(ctx, "velero-bucket", &s3.BucketArgs{
			Bucket: pulumi.String(veleroBucketName),
			Acl:    pulumi.String("private"),
			ServerSideEncryptionConfiguration: &s3.BucketServerSideEncryptionConfigurationArgs{
				Rule: &s3.BucketServerSideEncryptionConfigurationRuleArgs{
					ApplyServerSideEncryptionByDefault: &s3.BucketServerSideEncryptionConfigurationRuleApplyServerSideEncryptionByDefaultArgs{
						SseAlgorithm: pulumi.String("AES256"),
					},
				},
			},
		}, pulumi.Parent(component), pulumi.RetainOnDelete(true))
		if err != nil {
			return nil, err
		}
	}

	return component, nil
}
