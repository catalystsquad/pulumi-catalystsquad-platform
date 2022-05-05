package provider

import (
	"errors"
	"fmt"
	"os"

	k8syaml "github.com/pulumi/pulumi-kubernetes/sdk/v3/go/kubernetes/yaml"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	yamlv3 "gopkg.in/yaml.v3"
)

// ArgocdAppArgs supplies input for configuring an ArgocdApplication
type ArgocdAppArgs struct {
	// Required, name of the Argocd Application
	Name string `pulumi:"name"`
	// Required, spec of the Argocd Application
	Spec *ArgocdApplication `pulumi:"spec"`
	// Optional, namespace to deploy Argocd Application to. Should be the
	// namespace where the argocd server runs. Default: "argo-cd"
	Namespace string `pulumi:"namespace"`
	// Optional, apiVersion of the Argocd Application. Default: v1alpha1
	ApiVersion string `pulumi:"apiVersion"`
}

// ArgocdApp pulumi component resource
type ArgocdApp struct {
	pulumi.ResourceState
}

func NewArgocdApp(ctx *pulumi.Context, name string, args *ArgocdAppArgs, opts ...pulumi.ResourceOption) (*ArgocdApp, error) {
	if args == nil {
		args = &ArgocdAppArgs{}
	}

	component := &ArgocdApp{}
	err := ctx.RegisterComponentResource("catalystsquad-platform:index:ArgocdApp", name, component, opts...)
	if err != nil {
		return nil, err
	}

	if args.Name == "" {
		return nil, errors.New("name argument not supplied")
	}
	if args.Spec == nil {
		return nil, errors.New("spec argument not supplied")
	}

	argocdNamespace := "argo-cd"
	if args.Namespace != "" {
		argocdNamespace = args.Namespace
	}

	application := ArgocdApplication{
		ApiVersion: "argoproj.io/v1alpha1",
		Kind:       "Application",
		Metadata: map[string]interface{}{
			"name":      args.Name,
			"namespace": argocdNamespace,
		},
	}

	_, err = SyncArgocdApplication(ctx, name, application, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	return component, nil
}

// SyncArgocdApplication takes in a pulumi resource name, an argocd
// application, and any pulumi options. It will replace secrets in the
// spec.source.helm.values with the configured secrets provider, then sync the
// resulting yaml to k8s
func SyncArgocdApplication(ctx *pulumi.Context, pulumiResourceName string, application ArgocdApplication, opts ...pulumi.ResourceOption) (pulumi.Resource, error) {
	// marshall application to yaml
	bytes, err := yamlv3.Marshal(application)
	if err != nil {
		return nil, err
	}
	return SyncKubernetesManifest(ctx, pulumiResourceName, bytes, opts...)
}

// SyncKubernetesManifest takes in a pulumi resource name, and a yaml
// kubernetes manifest as byte array.  It writes the manifest to a file, defers
// deletion of said file, and creates a pulumi config file from it.  Pulumi
// creates the k8s resources from the config file. Recommended use is to store
// your manifests in yaml file, embed them, template them with pulumi secrets,
// or variables, and then pass them to this method to sync the kubernetes
// resource, whatever it may be.
func SyncKubernetesManifest(ctx *pulumi.Context, pulumiResourceName string, manifest []byte, opts ...pulumi.ResourceOption) (pulumi.Resource, error) {
	// write bytes to file
	tempFileName := fmt.Sprintf("/tmp/%s.yaml", pulumiResourceName)
	err := os.WriteFile(tempFileName, manifest, 0644)
	if err != nil {
		return nil, err
	}
	// defer file deletion
	defer func() {
		err = os.Remove(tempFileName)
	}()
	// get pulumi configfile from written manifest
	resource, err := k8syaml.NewConfigFile(ctx, pulumiResourceName, &k8syaml.ConfigFileArgs{
		File: tempFileName,
	}, opts...)
	return resource, err
}
