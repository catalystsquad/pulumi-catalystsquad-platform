# catalystsquad-platform Pulumi Component Provider (Go)

The Catalyst Squad Platform is a combination of cloud resources and open-source
Kubernetes tooling, that enables quickly provisioning a Kubernetes environment
for deploying production code.

The platform consists of cloud provider networking resources, a Kubernetes
cluster, and development tooling deployed into Kubernetes. 

- [AWS VPC](#aws-vpc)
- [AWS EKS cluster](#aws-eks-cluster)
- [Bootstrap cluster](#bootstrap-cluster)
- [Argocd application](#argocd-application)


## Prerequisites

- Go 1.17
- Pulumi CLI
- Node.js (to build the Node.js SDK)
- Yarn (to build the Node.js SDK)
- Python 3.6+ (to build the Python SDK)
- .NET Core SDK (to build the .NET SDK)

## Build and Test

```bash
# Build and install the provider (plugin copied to $GOPATH/bin)
make install_provider

# Regenerate SDKs
make generate

# Test Go SDK
$ make install_go_sdk
$ cd examples/complex-full-stack-go
$ pulumi stack init test
$ pulumi config set aws:region us-east-1
$ pulumi up
```

## Resources

Many component resources included fields that default to the stack name,
simplifying what is required to be supplied to the resources. Be sure to
overwrite these fields if this is not desired.


### AWS VPC

The VPC component provisions a VPC with an optional set of public and private
subnets. All that is required is to specify which availability zones and what
subnet CIDR blocks are desired.

For an example implementation of the VPC component, see [examples/simple-vpc-go/main.go](examples/simple-vpc-go/main.go)

Public subnets will be provisioned with NAT gateways that the private subnets
will automatically use.

Automatic assignment of necessary tags for operating EKS will be applied to
private and public subnets. This can be disabled by specifying the
`enableEksClusterTags` input property. The EKS cluster name will default to the
stack name, which can be overwritten via the `eksClusterName` input property.

Vpc component input properties:

| input property         | type               | description                                                                      |
| ---                    | ---                | ---                                                                              |
| availabilityZoneConfig | []AvailabilityZone | Optional, list of AvailabilityZones to create subnets in. Default: []            |
| cidr                   | string             | Optional, CIDR block of the VPC. Default: 10.0.0.0/16                            |
| eksClusterName         | string             | Optional, EKS cluster name, if VPC is used for EKS. Default: <stack name>        |
| enableEksClusterTags   | boolean            | Optional, whether to enable required EKS cluster tags to subnets.  Default: true |
| name                   | string             | Optional, Name tag value for VPC resource. Default: <stack name>                 |
| tags                   | map[string]string  | Optional, tags to add to all resources. Default: {}                              |


AvailabilityZone input properties:

| input property    | type   | description                                                                                    |
| ---               | ---    | ---                                                                                            |
| azName            | string | Name of the availability zone to deploy subnets to.                                            |
| privateSubnetCidr | string | CIDR for private subnets in the availability zone. If not supplied, the subnet is not created. |
| publicSubnetCidr  | string | CIDR for private subnets in the availability zone. If not supplied the subnet is not created.  |


### AWS EKS cluster

The EKS component provisions an EKS cluster with optional parameters for
configuring node groups, enabling access to ECR, and access for the cluster
autoscaler via IRSA. For a full list of possible configuration, see the input
properties section.

For an example implementation of the EKS component, see [examples/simple-eks-go/main.go](examples/simple-eks-go/main.go)

Eks component input properties:

| input property                   | type                | description                                                                                                                         |
| ---                              | ---                 | ---                                                                                                                                 |
| clusterName                      | string              | Optional, name of the EKS cluster. Default: <stack name>                                                                            |
| k8sVersion                       | string              | Optional, k8s version of the EKS cluster. Default: 1.22.6                                                                           |
| nodeGroupVersion                 | string              | Optional, k8s version of all node groups. Allows for upgrading the control plane before upgrading nodegroups. Default: <k8sVersion> |
| nodeGroupConfig                  | []EksNodeGroup      | Required, list of nodegroup configurations to create.                                                                               |
| authConfigmapConfig:             | AuthConfigMapConfig | Optional, configures management of the eks auth configmap.                                                                          |
| enableECRAccess                  | boolean             | Optional, whether to enable ECR access policy on nodegroups. Default: true                                                          |
| enableClusterAutoscalerResources | boolean             | Optional, whether to enable cluster autoscaler IRSA resources. Default: true                                                        |
| clusterAutoscalerServiceAccount  | string              | Optional, cluster autoscaler service account name for IRSA. Default: cluster-autoscaler                                             |
| clusterAutoscalerNamespace       | string              | Optional, cluster autoscaler namespace for IRSA. Default: cluster-autoscaler                                                        |
| enabledClusterLogTypes           | string              | Optional, list of log types to enable on the cluster. Default: []                                                                   |
| subnetIDs                        | []string            | Required, list of subnet IDs to deploy the cluster and nodegroups to                                                                |
| kubeConfigAssumeRoleArn          | string              | Optional, assume role arn to add to the kubeconfig.                                                                                 |
| kubeConfigAwsProfile             | string              | Optional, AWS profile to add to the kubeconfig.                                                                                     |


EksNodeGroup input properties:

| input property | type     | description                                                                              |
| ---            | ---      | ---                                                                                      |
| namePrefix     | string   | Name prefix of node group                                                                |
| desiredSize    | integer  | Desired size of node group                                                               |
| maxSize        | integer  | Max size of node group                                                                   |
| minSize        | integer  | Min size of node group                                                                   |
| instanceTypes: | []string | List of instance types to allow the nodegroup to use                                     |
| subnetIDs      | []string | Optional, list of subnet IDs to deploy the nodegroup in. Defaults to EKS cluster subnets |


AuthConfigMapConfig input properties:

| input property       | type                         | description                                                       |
| ---                  | ---                          | ---                                                               |
| autoDiscoverSSORoles | []SSORolePermissionSetConfig | Optional, list of AWS SSO permission set roles to autodiscover.   |
| iamRoles             | []IAMIdentityConfig          | Optional, list of IAM roles to grant access in the auth configmap |
| iamUsers             | []IAMIdentityConfig          | Optional, list of IAM users to grant access in the auth configmap |


### Bootstrap cluster

The BootstrapCluster component deploys common use-case configuration and
services to Kubernetes, including the kube-prometheus-stack, argo-cd, the
catalystsquad [chart-platform-services](https://github.com/catalystsquad/chart-platform-services),
and management of the EKS auth configmap

Management of the EKS auth configmap includes auto discover of AWS SSO roles
based on regex, greatly simplifying what needs to be hardcoded, because SSO
role names have generated random IDs.


BootstrapCluster input properties:

| input property                | type                        | description                                                                                                |
| ---                           | ---                         | ---                                                                                                        |
| argocdHelmConfig              | HelmReleaseConfig           | Optional, configures the argocd helm release.                                                              |
| kubePrometheusStackHelmConfig | HelmReleaseConfig           | Optional, configures the kube-prometheus-stack helm release.                                               |
| prometheusRemoteWriteConfig   | PrometheusRemoteWriteConfig | Optional, configuration for a prometheus remoteWrite secret. Does not deploy if not specified.             |
| platformApplicationConfig     | PlatformApplicationConfig   | Optional, configures the platform application release. Does not deploy if not specified.                   |


HelmReleaseConfig input properties:

| input property | type        | description                                                            |
| ---            | ---         | ---                                                                    |
| version        | string      | Optional for each implementation, defaults specific to each helm chart |
| valuesFiles    | []string    | Optional for each implementation, empty by default                     |
| values         | map[string] | Optional for each implementation, empty by default                     |


SSORolePermissionSetConfig input properties:

| input property   | type     | description                                                                                   |
| ---              | ---      | ---                                                                                           |
| name             | string   | Name of the permission set. Will use for autodiscovery using regex "AWSReservedSSO_<name>_.*" |
| permissionGroups | []string | List of permission groups to add to each identity. Ex: system:masters                         |
| username         | string   | Optional username field, defaults to the name of the SSO role.                                |


IAMIdentityConfig input properties:

| input property   | type     | description                                                    |
| ---              | ---      | ---                                                            |
| arn              | string   | Required, ARN of IAM role to use in configmap                  |
| permissionGroups | []string | Required, permission groups to add role to. Ex: system:masters |
| username         | string   | Optional username field, defaults to role name                 |


PrometheusRemoteWriteConfig input properties:

| input property    | type   | description                                                                   |
| ---               | ---    | ---                                                                           |
| basicAuthUsername | string | Optional, basic auth username. Default: <stack name>                          |
| basicAuthPassword | string | Required, basic auth password.                                                |
| secretName        | string | Optional, basic auth secret name. Default: prometheus-remote-write-basic-auth |


PlatformApplicationConfig input properties:

| input property             | type                        | description                                                                      |
| ---                        | ---                         | ---                                                                              |
| targetRevision             | string                      | Optional, target revision of platform application config. Deafult: >=1.0.0-alpha |
| syncPolicy                 | ArgocdApplicationSyncPolicy | Optional, sync policy of platform application config.                            |
| values                     | string                      | Optional, platform application values                                            |
| certManagerDnsSolverSecret | string                      | Optional, value of certmanager dns resolver secret                               |


### Argocd application

The ArgocdApp component deploys an ArgoCD application custom resource manifest
to Kubernetes.

ArgocdApp input properties:

| input property | type              | description                                                                                                                   |
| ---            | ---               | ---                                                                                                                           |
| name           | string            | Required, name of the Argocd Application                                                                                      |
| spec           | ArgocdApplication | Required, spec of the Argocd Application                                                                                      |
| namespace      | string            | Optional, namespace to deploy Argocd Application to. Should be the namespace where the argocd server runs. Default: "argo-cd" |
| apiVersion     | string            | Optional, apiVersion of the Argocd Application. Default: v1alpha1                                                             |


ArgocdApplication input properties:

| input property | type                  | description |
| ---            | ---                   | ---         |
| apiVersion     | string                |             |
| kind           | string                |             |
| metadata       | map[string]           |             |
| spec           | ArgocdApplicationSpec |             |


ArgocdApplicationSpec input properties:

| input property    | type                                 | description |
| ---               | ---                                  | ---         |
| source:           | ArgocdApplicationSpecSource          |             |
| destination       | ArgocdApplicationSpecDestination     |             |
| project           | string                               |             |
| syncPolicy        | ArgocdApplicationSyncPolicy          |             |
| ignoreDifferences | []ArgocdApplicationIgnoreDifferences |             |


### Observability Dependencies

Observability dependencies are AWS resources necessary for operating Cortex and
Loki. Including S3 buckets and IAM resources for EKS to authenticate to those
buckets via IRSA.

| input property       | type   | description                                                                                                                           |
| ---                  | ---    | ---                                                                                                                                   |
| oidcProviderArn      | string | Required, Arn of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.                                             |
| oidcProviderUrl      | string | Required, URL of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.                                             |
| cortexBucketName     | string | Optional, name of bucket to create for Cortex. Default: <account-id>-<stack-name>-cortex                                              |
| cortexIAMPolicyName  | string | Optional, Cortex's IAM policy name. Default: cortex-policy                                                                            |
| cortexIAMRoleName    | string | Optional, Cortex's IAM role name. Default: cortex-role                                                                                |
| cortexNamespace      | string | Optional, kubernetes namespace where Cortex will exist, for configuring the IRSA IAM role trust relationship. Default: cortex         |
| cortexServiceAccount | string | Optional, kubernetes service account name that Cortex will use, for configuring the IRSA IAM role trust relationship. Default: cortex |
| lokiBucketName       | string | Optional, name of bucket to create for Loki.  Default: <account-id>-<stack-name>-loki                                                 |
| LokiIAMPolicyName    | string | Optional, Loki's IAM policy name. Default: loki-policy                                                                                |
| lokiIAMRoleName      | string | Optional, Loki's IAM role name. Default: loki-role                                                                                    |
| lokiNamespace        | string | Optional, kubernetes namespace where Loki will exist, for configuring the IRSA IAM role trust relationship. Default: loki             |
| lokiServiceAccount   | string | Optional, kubernetes service account name that Loki will use, for configuring the IRSA IAM role trust relationship. Default: loki     |


## Velero Dependencies

Velero dependencies are AWS resources necessary for operating Velero in EKS. It
includes an S3 bucket which is retained on destroy, and IAM resources for
authentication to the bucket via IRSA.

| input property       | type    | description                                                                                                                           |
| ---                  | ---     | ---                                                                                                                                   |
| oidcProviderArn      | string  | Required, Arn of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.                                             |
| oidcProviderUrl      | string  | Required, URL of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.                                             |
| createBucket         | boolean | Optional, whether to create the Velero S3 bucket. Allows the bucket to exist outside of pulumi. Default: true                         |
| veleroBucketName     | string  | Optional, Velero's bucket name. Default: <account-id>-<stack-name>-velero                                                             |
| veleroIAMPolicyName  | string  | Optional, Velero's IAM policy name. Default: <stack-name>-velero-policy                                                               |
| veleroIAMRoleName    | string  | Optional, Velero's IAM role name. Default: <stack-name>-velero-role                                                                   |
| veleroNamespace      | string  | Optional, kubernetes namespace where Velero will exist, for configuring the IRSA IAM role trust relationship. Default: velero         |
| veleroServiceAccount | string  | Optional, kubernetes service account name that Velero will use, for configuring the IRSA IAM role trust relationship. Default: velero |
