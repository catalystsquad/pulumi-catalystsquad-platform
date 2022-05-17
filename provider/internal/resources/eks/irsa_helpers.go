package eks

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// CreateIrsaAssumeRolePolicy creates an iam assume role policy for IRSA
// https://docs.aws.amazon.com/eks/latest/userguide/create-service-account-iam-policy-and-role.html
func createIrsaAssumeRolePolicy(oidcProvider *iam.OpenIdConnectProvider, namespace string, serviceAccount string) pulumi.Output {
	return pulumi.All(oidcProvider.Arn, oidcProvider.Url).ApplyT(func(args []interface{}) (string, error) {
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
