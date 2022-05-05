package eksauth

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/eks"
	"github.com/pulumi/pulumi-aws/sdk/v4/go/aws/iam"
	"github.com/pulumi/pulumi-command/sdk/go/command/local"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"

	"gopkg.in/yaml.v2"
)

type AuthConfigMapConfig struct {
	// required if nodegroup IAM role autodiscovery not enabled
	NodeGroupIamRole string `pulumi:"nodeGroupIamRole"`
	// required if nodegroup IAM role not supplied
	EnableNodeGroupIamRoleAutoDiscover bool `pulumi:"enableNodeGroupIamRoleAutoDiscover"`
	// required if nodegroup IAM role not supplied
	EKSClusterName string `pulumi:"eksClusterName"`
	// optional list of AWS SSO permission set roles to autodiscover
	AutoDiscoverSSORoles []SSORolePermissionSetConfig `pulumi:"autoDiscoverSSORoles"`
	// optional list of IAM roles and users
	IAMRoles []IAMIdentityConfig `pulumi:"iamRoles"`
	IAMUsers []IAMIdentityConfig `pulumi:"iamUsers"`
}

type SSORolePermissionSetConfig struct {
	// name of permission set to discover for use in configmap
	Name string `pulumi:"name"`
	// required groups to add role to
	PermissionGroups []string `pulumi:"permissionGroups"`
	// optional username field, defaults to name field
	Username string `pulumi:"username"`
}

type IAMIdentityConfig struct {
	// arn of IAM role to use in configmap
	Arn string `pulumi:"arn"`
	// required groups to add role to
	PermissionGroups []string `pulumi:"permissionGroups"`
	// optional username field, defaults to role name
	Username string `pulumi:"username"`
}

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

type configMap struct {
	ApiVersion string            `yaml:"apiVersion"`
	Data       map[string]string `yaml:"data"`
	Kind       string            `yaml:"kind"`
	Metadata   configMapMetadata `yaml:"metadata"`
}

type configMapMetadata struct {
	Name      string `yaml:"name"`
	Namespace string `yaml:"namespace"`
}

var ssoRolePathPrefix string = "/aws-reserved/sso.amazonaws.com/"

func SyncAuthConfigMap(ctx *pulumi.Context, config AuthConfigMapConfig, opts ...pulumi.ResourceOption) error {
	var authConfigMap configMap = configMap{
		ApiVersion: "v1",
		Data:       map[string]string{},
		Kind:       "ConfigMap",
		Metadata: configMapMetadata{
			Name:      "aws-auth",
			Namespace: "kube-system",
		},
	}
	var mapRoles []mapRolesElement
	var mapUsers []mapUsersElement

	var nodeRoleArn string
	var err error
	if config.EnableNodeGroupIamRoleAutoDiscover {
		if config.EKSClusterName != "" {
			nodeRoleArn, err = discoverNodeIAMRole(ctx, config.EKSClusterName)
			if err != nil {
				return err
			}
		} else {
			return errors.New("Node Group IAM Role auto discover enabled, but EKS cluster name not supplied")
		}
	} else {
		if config.NodeGroupIamRole != "" {
			nodeRoleArn = config.NodeGroupIamRole
		} else {
			return errors.New("Node Group IAM Role not supplied, auto discover not enabled")
		}
	}

	// add nodegroup iam role to mapRoles
	mapRoles = append(mapRoles, mapRolesElement{
		RoleArn:  nodeRoleArn,
		Username: "system:node:{{EC2PrivateDNSName}}",
		Groups: []string{
			"system:bootstrappers",
			"system:nodes",
		},
	})

	// add all sso autodiscovery roles
	for _, ssoRoleConfig := range config.AutoDiscoverSSORoles {
		// default username to the permissionset name
		username := ssoRoleConfig.Name
		if ssoRoleConfig.Username != "" {
			username = ssoRoleConfig.Username
		}

		roleArn, err := discoverSSORole(ctx, ssoRoleConfig.Name)
		if err != nil {
			return err
		}

		mapRoles = append(mapRoles, mapRolesElement{
			RoleArn:  removeArnPath(roleArn),
			Username: username,
			Groups:   ssoRoleConfig.PermissionGroups,
		})
	}

	// add all iam roles
	for _, roleConfig := range config.IAMRoles {
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
	for _, userConfig := range config.IAMUsers {
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

	// marshal all the data fields
	mapRolesBytes, err := yaml.Marshal(&mapRoles)
	if err != nil {
		return err
	}
	authConfigMap.Data["mapRoles"] = string(mapRolesBytes)

	// omit mapUsers if empty
	if len(mapUsers) != 0 {
		mapUsersBytes, err := yaml.Marshal(&mapUsers)
		if err != nil {
			return err
		}
		authConfigMap.Data["mapUsers"] = string(mapUsersBytes)
	}

	// marshal configmap
	configMapYaml, err := yaml.Marshal(&authConfigMap)
	applyKubernetesManifest(ctx, "aws-auth-configmap", configMapYaml, opts...)
	return err
}

// assumes that all nodegroups have the same IAM role, so only finds the first
// roleArn of the first nodegroup discovered
func discoverNodeIAMRole(ctx *pulumi.Context, clusterName string) (roleArn string, err error) {
	nodegroups, err := eks.GetNodeGroups(ctx, &eks.GetNodeGroupsArgs{
		ClusterName: clusterName,
	})
	if err != nil {
		return
	}

	nodegroup, err := eks.LookupNodeGroup(ctx, &eks.LookupNodeGroupArgs{
		ClusterName:   clusterName,
		NodeGroupName: nodegroups.Names[0],
	})
	if err != nil {
		return
	}

	roleArn = nodegroup.NodeRoleArn
	return
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

// trim an ARN to use in the username field
func arnToUsername(i string) string {
	a := strings.Split(i, "/")
	return a[len(a)-1]
}

func applyKubernetesManifest(ctx *pulumi.Context, pulumiResourceName string, manifest []byte, opts ...pulumi.ResourceOption) error {
	// write bytes to file
	tempFileName := fmt.Sprintf("/tmp/%s.yaml", pulumiResourceName)
	err := os.WriteFile(tempFileName, manifest, 0644)
	if err != nil {
		return err
	}
	// execute kubectl apply
	_, err = local.NewCommand(ctx, pulumiResourceName, &local.CommandArgs{
		Create:   pulumi.String(fmt.Sprintf("kubectl apply -f %s; rm %s", tempFileName, tempFileName)),
		Triggers: pulumi.ToArrayOutput([]pulumi.Output{pulumi.ToOutput(string(manifest))}),
	}, opts...)
	return err
}
