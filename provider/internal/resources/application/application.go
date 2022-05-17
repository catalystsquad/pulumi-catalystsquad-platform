package application

import (
	"errors"

	"github.com/catalystsquad/pulumi-catalystsquad-platform/internal/argocd"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ArgocdAppArgs supplies input for configuring an ArgocdApplication
type ArgocdAppArgs struct {
	// Required, name of the Argocd Application
	Name string `pulumi:"name"`
	// Required, spec of the Argocd Application
	Spec *argocd.ArgocdApplication `pulumi:"spec"`
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

	application := argocd.ArgocdApplication{
		ApiVersion: "argoproj.io/v1alpha1",
		Kind:       "Application",
		Metadata: map[string]interface{}{
			"name":      args.Name,
			"namespace": argocdNamespace,
		},
	}

	_, err = argocd.SyncArgocdApplication(ctx, name, application, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	return component, nil
}
