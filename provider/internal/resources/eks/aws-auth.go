package eks

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/iam"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"gopkg.in/yaml.v2"
)

type mapRolesElement struct {
	Groups   []string `yaml:"groups"`
	RoleArn  string   `yaml:"rolearn"`
	Username string   `yaml:"username"`
}

type mapUsersElement struct {
	Groups   []string `yaml:"groups"`
	UserArn  string   `yaml:"userarn"`
	Username string   `yaml:"username"`
}

func generateAwsAuthConfigmap(ctx *pulumi.Context, nodegroupRoleArn string,
	autoDiscoverSSORoles []SSORolePermissionSetConfig, iamRoles []IAMIdentityConfig,
	iamUsers []IAMIdentityConfig) (map[string]string, error) {

	var mapRoles []mapRolesElement
	var mapUsers []mapUsersElement
	data := make(map[string]string)

	// add nodegroup iam role to mapRoles
	mapRoles = append(mapRoles, mapRolesElement{
		RoleArn:  nodegroupRoleArn,
		Username: "system:node:{{EC2PrivateDNSName}}",
		Groups: []string{
			"system:bootstrappers",
			"system:nodes",
		},
	})

	// add all sso autodiscovery roles
	for _, ssoRoleConfig := range autoDiscoverSSORoles {
		// default username to the permissionset name
		username := ssoRoleConfig.Name
		if ssoRoleConfig.Username != "" {
			username = ssoRoleConfig.Username
		}

		roleArn, err := discoverSSORole(ctx, ssoRoleConfig.Name)
		if err != nil {
			return nil, err
		}

		mapRoles = append(mapRoles, mapRolesElement{
			RoleArn:  removeArnPath(roleArn),
			Username: username,
			Groups:   ssoRoleConfig.PermissionGroups,
		})
	}

	// add all iam roles
	for _, roleConfig := range iamRoles {
		// default username to the role name, derived from the role arn
		username := arnToUsername(roleConfig.Arn)
		if roleConfig.Username != "" {
			username = roleConfig.Username
		}

		mapRoles = append(mapRoles, mapRolesElement{
			RoleArn:  removeArnPath(roleConfig.Arn),
			Username: username,
			Groups:   roleConfig.PermissionGroups,
		})
	}

	// add all iam users
	for _, userConfig := range iamUsers {
		// default username to the role name, derived from the role arn
		username := arnToUsername(userConfig.Arn)
		if userConfig.Username != "" {
			username = userConfig.Username
		}

		mapUsers = append(mapUsers, mapUsersElement{
			UserArn:  removeArnPath(userConfig.Arn),
			Username: username,
			Groups:   userConfig.PermissionGroups,
		})
	}

	mapRolesBytes, err := yaml.Marshal(&mapRoles)
	if err != nil {
		return nil, err
	}
	data["mapRoles"] = string(mapRolesBytes)

	// omit mapUsers if empty
	if len(mapUsers) != 0 {
		mapUsersBytes, err := yaml.Marshal(&mapUsers)
		if err != nil {
			return nil, err
		}
		data["mapUsers"] = string(mapUsersBytes)
	}

	return data, nil
}

func discoverSSORole(ctx *pulumi.Context, permissionSetName string) (roleArn string, err error) {
	ssoRoleRegex := fmt.Sprintf("AWSReservedSSO_%s_.*", permissionSetName)

	discoverSSORole, err := iam.GetRoles(ctx, &iam.GetRolesArgs{
		NameRegex:  pulumi.StringRef(ssoRoleRegex),
		PathPrefix: &ssoRolePathPrefix,
	})
	if err != nil {
		return
	}

	// fail if we don't discover just 1 role
	if len(discoverSSORole.Arns) != 1 {
		err = errors.New(fmt.Sprintf(
			"role auto discovery failed, discovered %d",
			len(discoverSSORole.Arns),
		))
		return
	}

	roleArn = discoverSSORole.Arns[0]
	return
}

// auth configmap doesn't support arns with paths, so we have to remove them
// https://docs.aws.amazon.com/eks/latest/userguide/troubleshooting_iam.html#security-iam-troubleshoot-ConfigMap
func removeArnPath(arn string) string {
	a := strings.Split(arn, "/")
	return strings.Join([]string{a[0], a[len(a)-1]}, "/")
}

// trim an ARN to use in the aws-auth configmap username field
func arnToUsername(i string) string {
	a := strings.Split(i, "/")
	return a[len(a)-1]
}
