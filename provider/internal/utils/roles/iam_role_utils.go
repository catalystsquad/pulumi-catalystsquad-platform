package roles

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateIrsaAssumeRolePolicy creates an iam assume role policy for IRSA
// https://docs.aws.amazon.com/eks/latest/userguide/create-service-account-iam-policy-and-role.html
func CreateIrsaAssumeRolePolicy(oidcProviderArn pulumi.StringInput, oidcProviderURL pulumi.StringInput,
	namespace string, serviceAccount string) pulumi.Output {

	return pulumi.All(oidcProviderArn, oidcProviderURL).ApplyT(func(args []interface{}) (string, error) {
		arn := args[0].(string)
		provider := strings.TrimLeft(args[1].(string), "https://")
		policyByteArray, err := json.Marshal(map[string]interface{}{
			"Version": "2012-10-17",
			"Statement": []map[string]interface{}{
				{
					"Action": "sts:AssumeRoleWithWebIdentity",
					"Effect": "Allow",
					"Principal": map[string]interface{}{
						"Federated": arn,
					},
					"Condition": map[string]interface{}{
						"StringEquals": map[string]string{
							fmt.Sprintf("%s:sub", provider): fmt.Sprintf("system:serviceaccount:%s:%s", namespace, serviceAccount),
						},
					},
				},
			},
		})
		return string(policyByteArray), err
	})
}

// CreateRoleWithPolicy creates an IAM policy and associates it to a new IAM role.
func CreateRoleWithPolicy(ctx *pulumi.Context, resourceNamePrefix string, policyName string,
	roleName string, policyJSON string, assumeRolePolicyJSON pulumi.Output,
	roleDescription string, parent pulumi.Resource) error {

	policy, err := iam.NewPolicy(ctx, fmt.Sprintf("%s-policy", resourceNamePrefix), &iam.PolicyArgs{
		Name:        pulumi.String(policyName),
		Description: pulumi.String(roleDescription),
		Policy:      pulumi.String(policyJSON),
	}, pulumi.Parent(parent))
	if err != nil {
		return err
	}

	role, err := iam.NewRole(ctx, fmt.Sprintf("%s-role", resourceNamePrefix), &iam.RoleArgs{
		Name:             pulumi.String(roleName),
		AssumeRolePolicy: assumeRolePolicyJSON,
	}, pulumi.Parent(parent))
	if err != nil {
		return err
	}

	_, err = iam.NewRolePolicyAttachment(ctx, fmt.Sprintf("%s-role-policy-attachment", resourceNamePrefix), &iam.RolePolicyAttachmentArgs{
		Role:      role,
		PolicyArn: policy.Arn,
	}, pulumi.Parent(parent))
	return err
}
