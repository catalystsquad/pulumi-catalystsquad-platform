package argocd

import "github.com/pulumi/pulumi/sdk/v3/go/pulumi"

// ArgocdApplication is a struct that marshalls into valid argocd application
// yaml. We could use the argo types but we have had problems with the yaml
// marshalling, and that also requires depending on argo, and nearly the entire
// k8s api.  This is let DRY and less direct but more simple and
// straightforward. We'll need to keep this in sync with their spec though.
// see spec at
// https://github.com/argoproj/argo-cd/blob/master/pkg/apis/application/v1alpha1/types.go
type ArgocdApplication struct {
	ApiVersion string                 `yaml:"apiVersion" pulumi:"apiVersion"`
	Kind       string                 `yaml:"kind" pulumi:"kind"`
	Metadata   map[string]interface{} `yaml:"metadata" pulumi:"metadata"`
	Spec       ArgocdApplicationSpec  `yaml:"spec" pulumi:"spec"`
}

type ArgocdApplicationSpec struct {
	Source            ArgocdApplicationSpecSource          `yaml:"source" pulumi:"source"`
	Destination       ArgocdApplicationSpecDestination     `yaml:"destination" pulumi:"destination"`
	Project           string                               `yaml:"project" pulumi:"project"`
	SyncPolicy        ArgocdApplicationSyncPolicy          `yaml:"syncPolicy,omitempty" pulumi:"syncPolicy"`
	IgnoreDifferences []ArgocdApplicationIgnoreDifferences `yaml:"ignoreDifferences,omitempty" pulumi:"ignoreDifferences"`
}

type ArgocdApplicationSpecSource struct {
	RepoUrl        string          `yaml:"repoURL" pulumi:"repoURL"`
	Path           string          `yaml:"path,omitempty" pulumi:"path"`
	TargetRevision string          `yaml:"targetRevision,omitempty" pulumi:"targetRevision"`
	Helm           HelmSource      `yaml:"helm,omitempty" pulumi:"helm"`
	Kustomize      KustomizeSource `yaml:"kustomize,omitempty" pulumi:"kustomize"`
	Directory      DirectorySource `yaml:"directory,omitempty" pulumi:"directory"`
	Plugin         PluginSource    `yaml:"plugin,omitempty" pulumi:"plugin"`
	Chart          string          `yaml:"chart,omitempty" pulumi:"chart"`
}

type HelmSource struct {
	ValueFiles              []string                  `yaml:"valueFiles,omitempty" pulumi:"valueFiles"`
	Parameters              []HelmSourceParameter     `yaml:"parameters,omitempty" pulumi:"parameters"`
	ReleaseName             string                    `yaml:"releaseName,omitempty" pulumi:"releaseName"`
	Values                  pulumi.StringInput        `yaml:"values,omitempty" pulumi:"values"`
	FileParameters          []HelmSourceFileParameter `yaml:"fileParameters,omitempty" pulumi:"fileParameters"`
	Version                 string                    `yaml:"version,omitempty" pulumi:"version"`
	PassCredentials         bool                      `yaml:"passCredentials,omitempty" pulumi:"passCredentials"`
	IgnoreMissingValueFiles bool                      `yaml:"ignoreMissingValueFiles,omitempty" pulumi:"ignoreMissingValueFiles"`
	SkipCrds                bool                      `yaml:"skipCrds,omitempty" pulumi:"skipCrds"`
}

type HelmSourceParameter struct {
	Name        string `yaml:"name,omitempty" pulumi:"name"`
	Value       string `yaml:"value,omitempty" pulumi:"value"`
	ForceString bool   `yaml:"forceString,omitempty" pulumi:"forceString"`
}

type HelmSourceFileParameter struct {
	Name string `pulumi:"name"`
	Path string `pulumi:"path"`
}

type KustomizeSource struct {
	NamePrefix             string            `yaml:"namePrefix,omitempty" pulumi:"namePrefix"`
	NameSuffix             string            `yaml:"nameSuffix,omitempty" pulumi:"nameSuffix"`
	Images                 []string          `yaml:"images,omitempty" pulumi:"images"`
	CommonLabels           map[string]string `yaml:"commonLabels,omitempty" pulumi:"commonLabels"`
	Version                string            `yaml:"version,omitempty" pulumi:"version"`
	CommonAnnotations      map[string]string `yaml:"commonAnnotations,omitempty" pulumi:"commonAnnotations"`
	ForceCommonLabels      bool              `yaml:"forceCommonLabels,omitempty" pulumi:"forceCommonLabels"`
	ForceCommonAnnotations bool              `yaml:"forceCommonAnnotations,omitempty" pulumi:"forceCommonAnnotations"`
}

type DirectorySource struct {
	Recurse bool                   `yaml:"recurse,omitempty" pulumi:"recurse"`
	Jsonnet DirectorySourceJsonnet `yaml:"jsonnet,omitempty" pulumi:"jsonnet"`
	Exclude string                 `yaml:"exclude,omitempty" pulumi:"exclude"`
	Include string                 `yaml:"include,omitempty" pulumi:"include"`
}

type DirectorySourceJsonnet struct {
	ExtVars []JsonnetVar `yaml:"extVars,omitempty" pulumi:"extVars"`
	TLAs    []JsonnetVar `yaml:"TLAs,omitempty" pulumi:"TLAs"`
	Libs    []string     `yaml:"libs,omitempty" pulumi:"libs"`
}

type JsonnetVar struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
	Code  bool   `yaml:"code,omitempty" pulumi:"code"`
}

type PluginSource struct {
	Name string            `yaml:"name,omitempty" pulumi:"name"`
	Env  []PluginSourceEnv `yaml:"env,omitempty" pulumi:"env"`
}

type PluginSourceEnv struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ArgocdApplicationSpecDestination struct {
	Server    string `yaml:"server,omitempty" pulumi:"server"`
	Namespace string `yaml:"namespace,omitempty" pulumi:"namespace"`
	Name      string `yaml:"name,omitempty" pulumi:"name"`
}

type ArgocdApplicationSyncPolicy struct {
	Automated   SyncPolicyAutomated `yaml:"automated,omitempty" pulumi:"automated"`
	Retry       SyncPolicyRetry     `yaml:"retry,omitempty" pulumi:"retry"`
	SyncOptions []string            `yaml:"syncOptions,omitempty" pulumi:"syncOptions"`
}

type SyncPolicyAutomated struct {
	Prune      bool `yaml:"prune,omitempty" pulumi:"prune"`
	SelfHeal   bool `yaml:"selfHeal,omitempty" pulumi:"selfHeal"`
	AllowEmpty bool `yaml:"allowEmpty,omitempty" pulumi:"allowEmpty"`
}

type SyncPolicyRetry struct {
	Limit   int          `yaml:"limit,omitempty" pulumi:"limit"`
	Backoff RetryBackoff `yaml:"backoff,omitempty" pulumi:"backoff"`
}

type RetryBackoff struct {
	Duration    string `yaml:"duration,omitempty" pulumi:"duration"`
	Factor      int    `yaml:"factor,omitempty" pulumi:"factor"`
	MaxDuration string `yaml:"maxDuration,omitempty" pulumi:"maxDuration"`
}

type ArgocdApplicationIgnoreDifferences struct {
	Group                 string   `yaml:"group,omitempty" pulumi:"group"`
	Kind                  string   `yaml:"kind,omitempty" pulumi:"kind"`
	Name                  string   `yaml:"name,omitempty" pulumi:"name"`
	Namespace             string   `yaml:"namespace,omitempty" pulumi:"namespace"`
	JsonPointers          []string `yaml:"jsonPointers,omitempty" pulumi:"jsonPointers"`
	JQPathExpressions     []string `yaml:"jqPathExpressions,omitempty" pulumi:"jqPathExpressions"`
	ManagedFieldsManagers []string `yaml:"managedFieldsManagers,omitempty" pulumi:"managedFieldsManagers"`
}
