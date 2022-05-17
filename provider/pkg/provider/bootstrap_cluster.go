package provider

import (
	"github.com/catalystsquad/pulumi-catalystsquad-platform/internal/templates"
	"github.com/pkg/errors"
	corev1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/core/v1"
	"github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/helm/v3"
	metav1 "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/meta/v1"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	yamlv3 "gopkg.in/yaml.v3"
)

// ClusterBootstrapArgs supplies input for configuring ClusterBootstrap component
type ClusterBootstrapArgs struct {
	// Optional, configures the argocd helm release.
	ArgocdHelmConfig *HelmReleaseConfig `pulumi:"argocdHelmConfig"`
	// Optional, configures the kube-prometheus-stack helm release.
	KubePrometheusStackHelmConfig *HelmReleaseConfig `pulumi:"kubePrometheusStackHelmConfig"`
	// Optional, configuration for a prometheus remoteWrite secret. Does not
	// deploy if not specified.
	PrometheusRemoteWriteConfig *PrometheusRemoteWriteConfig `pulumi:"prometheusRemoteWriteConfig"`
	// Optional, configures the platform application release. Does not deploy
	// if not specified.
	PlatformApplicationConfig *PlatformApplicationConfig `pulumi:"platformApplicationConfig"`
}

type HelmReleaseConfig struct {
	// Optional for each implementation, defaults specific to each helm chart
	Version string `pulumi:"version"`
	// Optional for each implementation, empty by default
	ValuesFiles []string `pulumi:"valuesFiles"`
	// Optional for each implementation, empty by default
	Values map[string]interface{} `pulumi:"values"`
}

type PrometheusRemoteWriteConfig struct {
	// Optional, basic auth username. Default: <stack name>
	BasicAuthUsername string `pulumi:"basicAuthUsername"`
	// Required, basic auth password.
	BasicAuthPassword string `pulumi:"basicAuthPassword"`
	// Optional, basic auth secret name. Default: prometheus-remote-write-basic-auth
	SecretName string `pulumi:"secretName"`
}

type PlatformApplicationConfig struct {
	// Optional, target revision of platform application config. Deafult: >=1.0.0-alpha
	TargetRevision string `pulumi:"targetRevision"`
	// Optional, sync policy of platform application config.
	SyncPolicy *ArgocdApplicationSyncPolicy `pulumi:"syncPolicy"`
	// Optional, platform application values
	Values pulumi.StringInput `pulumi:"values"`
	// Optional, value of certmanager dns resolver secret
	CertManagerDnsSolverSecret pulumi.StringInput `pulumi:"certManagerDnsSolverSecret"`
}

// ClusterBootstrap pulumi component resource
type ClusterBootstrap struct {
	pulumi.ResourceState
}

func NewClusterBootstrap(ctx *pulumi.Context, name string, args *ClusterBootstrapArgs, opts ...pulumi.ResourceOption) (*ClusterBootstrap, error) {
	if args == nil {
		args = &ClusterBootstrapArgs{}
	}

	component := &ClusterBootstrap{}
	err := ctx.RegisterComponentResource("catalystsquad-platform:index:ClusterBootstrap", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// deploy kube-prometheus-stack remote-write basic auth secret
	prometheusRemoteWriteSecret, err := deployPrometheusRemoteWriteBasicAuthSecret(ctx, component, args)
	if err != nil {
		return nil, err
	}

	// dynamic depends on for the prometheusRemoteWriteSecret
	var prometheusDependsOn pulumi.ResourceOption
	if prometheusRemoteWriteSecret != nil {
		prometheusDependsOn = pulumi.DependsOn([]pulumi.Resource{prometheusRemoteWriteSecret})
	}

	// deploy kube-prometheus-stack, this happens first so that CRDs exist for
	// service monitors in other helm releases
	prometheus, err := deployKubePrometheusStack(ctx, component, args, prometheusDependsOn)
	if err != nil {
		return nil, err
	}

	// deploy argocd, this helm chart installs service monitors, so it depends
	// on kube-prometheus-stack
	argocd, err := deployArgocd(ctx, component, args, pulumi.DependsOn([]pulumi.Resource{prometheus}))
	if err != nil {
		return nil, err
	}

	// deploy cluster argocd application
	if args.PlatformApplicationConfig != nil {
		platformApplication, err := deployPlatformApplicationManifest(
			ctx, component, args.PlatformApplicationConfig, pulumi.DependsOn([]pulumi.Resource{argocd}),
		) // depend on argocd for application CRDs
		if err != nil {
			return nil, err
		}

		// create cert-manager dns secret
		err = deployCertManagerDnsSolverSecret(
			ctx, component, args.PlatformApplicationConfig, pulumi.DependsOn([]pulumi.Resource{platformApplication}),
		)
		if err != nil {
			return nil, err
		}
	}

	return component, nil
}

func deployPrometheusRemoteWriteBasicAuthSecret(ctx *pulumi.Context, parent pulumi.Resource,
	args *ClusterBootstrapArgs) (pulumi.Resource, error) {

	// only enable if remoteWriteConfiguration is supplied
	if args.PrometheusRemoteWriteConfig != nil {

		// default argument values
		secretName := "prometheus-remote-write-basic-auth"
		if args.PrometheusRemoteWriteConfig.SecretName != "" {
			secretName = args.PrometheusRemoteWriteConfig.SecretName
		}

		username := ctx.Stack()
		if args.PrometheusRemoteWriteConfig.BasicAuthUsername != "" {
			username = args.PrometheusRemoteWriteConfig.BasicAuthUsername
		}

		// return an error if a remoteWriteConfig was supplied without a password
		if args.PrometheusRemoteWriteConfig.BasicAuthPassword == "" {
			return nil, errors.New("prometheusRemoteWriteConfig was supplied without a BasicAuthPassword value")
		}
		password := args.PrometheusRemoteWriteConfig.BasicAuthPassword

		// create secret
		secret, err := corev1.NewSecret(ctx, "prometheus-remote-write-basic-auth-secret", &corev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String(secretName),
				Namespace: pulumi.String("kube-prometheus-stack"),
			},
			StringData: pulumi.StringMap{
				"username": pulumi.String(username),
				"password": pulumi.String(password),
			},
		}, pulumi.Parent(parent))
		return secret, err
	}

	return nil, nil
}

func deployKubePrometheusStack(ctx *pulumi.Context, parent pulumi.Resource, args *ClusterBootstrapArgs,
	opts ...pulumi.ResourceOption) (pulumi.Resource, error) {

	// default kube-prometheus-stack helm release configuration
	version := "33.1.0"
	values := make(map[string]interface{})
	var valuesFiles []string
	if args.KubePrometheusStackHelmConfig != nil {
		if args.KubePrometheusStackHelmConfig.Version != "" {
			version = args.KubePrometheusStackHelmConfig.Version
		}
		if args.KubePrometheusStackHelmConfig.Values != nil {
			values = args.KubePrometheusStackHelmConfig.Values
		}
		if args.KubePrometheusStackHelmConfig.ValuesFiles != nil {
			valuesFiles = args.KubePrometheusStackHelmConfig.ValuesFiles
		}
	}

	opts = append(opts, pulumi.Parent(parent))
	return helm.NewRelease(ctx, "kube-prometheus-stack", &helm.ReleaseArgs{
		Chart:           pulumi.String("kube-prometheus-stack"),
		Name:            pulumi.String("kube-prometheus-stack"),
		Namespace:       pulumi.String("kube-prometheus-stack"),
		CreateNamespace: pulumi.Bool(true),
		Version:         pulumi.String(version),
		RepositoryOpts: helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://prometheus-community.github.io/helm-charts"),
		},
		ValueYamlFiles: stringArrayToAssetOrArchiveArrayOutput(valuesFiles),
		Values:         pulumi.ToMap(values),
	}, opts...)
}

func deployArgocd(ctx *pulumi.Context, parent pulumi.Resource, args *ClusterBootstrapArgs,
	opts ...pulumi.ResourceOption) (pulumi.Resource, error) {

	// default argo-cd helm release configuration
	version := "3.33.8"
	values := make(map[string]interface{})
	var valuesFiles []string
	if args.ArgocdHelmConfig != nil {
		if args.ArgocdHelmConfig.Version != "" {
			version = args.ArgocdHelmConfig.Version
		}
		if args.ArgocdHelmConfig.Values != nil {
			values = args.ArgocdHelmConfig.Values
		}
		if args.ArgocdHelmConfig.ValuesFiles != nil {
			valuesFiles = args.ArgocdHelmConfig.ValuesFiles
		}
	}

	opts = append(opts, pulumi.Parent(parent))
	// deploy argo using helm
	argocd, err := helm.NewRelease(ctx, "argo-cd", &helm.ReleaseArgs{
		Chart:           pulumi.String("argo-cd"),
		Name:            pulumi.String("argo-cd"),
		Namespace:       pulumi.String("argo-cd"),
		CreateNamespace: pulumi.Bool(true),
		Version:         pulumi.String(version),
		RepositoryOpts: helm.RepositoryOptsArgs{
			Repo: pulumi.String("https://argoproj.github.io/argo-helm"),
		},
		ValueYamlFiles: stringArrayToAssetOrArchiveArrayOutput(valuesFiles),
		Values:         pulumi.ToMap(values),
	}, opts...)
	return argocd, err
}

func deployPlatformApplicationManifest(ctx *pulumi.Context, parent pulumi.Resource, args *PlatformApplicationConfig,
	opts ...pulumi.ResourceOption) (pulumi.Resource, error) {

	// default platform application config
	targetRevision := ">=1.0.0-alpha"
	if args.TargetRevision != "" {
		targetRevision = args.TargetRevision
	}

	var values pulumi.StringInput
	if args.Values != nil {
		values = args.Values
	}

	syncPolicy := ArgocdApplicationSyncPolicy{
		Automated: SyncPolicyAutomated{
			AllowEmpty: false,
			Prune:      true,
			SelfHeal:   true,
		},
		Retry: SyncPolicyRetry{
			Backoff: RetryBackoff{
				Duration:    "5s",
				Factor:      2,
				MaxDuration: "3m",
			},
			Limit: 3,
		},
		SyncOptions: []string{
			"CreateNamespace=true",
			"PrunePropagationPolicy=foreground",
			"PruneLast=true",
		},
	}
	if args.SyncPolicy != nil {
		syncPolicy = *args.SyncPolicy
	}

	// get application from template
	application, err := newApplicationFromBytes(templates.PlatformApplicationBytes)
	if err != nil {
		return nil, err
	}

	// set variables from stack config
	application.Spec.Source.TargetRevision = targetRevision
	application.Spec.Source.Helm.Values = values
	application.Spec.SyncPolicy = syncPolicy

	// sync
	opts = append(opts, pulumi.Parent(parent))
	resource, err := SyncArgocdApplication(ctx, "cluster-platform-application-services", application, opts...)

	return resource, err
}

func deployCertManagerDnsSolverSecret(ctx *pulumi.Context, parent pulumi.Resource, args *PlatformApplicationConfig, opts ...pulumi.ResourceOption) error {

	if args.CertManagerDnsSolverSecret != nil {
		secretValue := args.CertManagerDnsSolverSecret

		opts = append(opts, pulumi.Parent(parent))
		_, err := corev1.NewSecret(ctx, "cert-manager-cloudflare-api-token-secret", &corev1.SecretArgs{
			Metadata: &metav1.ObjectMetaArgs{
				Name:      pulumi.String("cloudflare-api-token-secret"),
				Namespace: pulumi.String("cert-manager"),
			},
			StringData: pulumi.StringMap{
				"api-token": secretValue,
			},
			Type: pulumi.String("Opaque"),
		}, opts...)
		return err
	}
	return nil
}

func stringArrayToAssetOrArchiveArrayOutput(in []string) pulumi.AssetOrArchiveArrayOutput {
	var o pulumi.AssetOrArchiveArray
	for _, i := range in {
		o = append(o, pulumi.NewFileAsset(i))
	}
	return o.ToAssetOrArchiveArrayOutput()
}

// newApplicationFromBytes transforms yaml formatted byte array into an
// ArgocdApplication struct
func newApplicationFromBytes(bytes []byte) (ArgocdApplication, error) {
	var application ArgocdApplication
	// marshall template into map[string]interface{}
	err := yamlv3.Unmarshal(bytes, &application)
	return application, err
}
