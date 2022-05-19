// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package catalystsquadplatform

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "catalystsquad-platform:index:ArgocdApp":
		r = &ArgocdApp{}
	case "catalystsquad-platform:index:ClusterBootstrap":
		r = &ClusterBootstrap{}
	case "catalystsquad-platform:index:Eks":
		r = &Eks{}
	case "catalystsquad-platform:index:ObservabilityDependencies":
		r = &ObservabilityDependencies{}
	case "catalystsquad-platform:index:VeleroDependencies":
		r = &VeleroDependencies{}
	case "catalystsquad-platform:index:Vpc":
		r = &Vpc{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:catalystsquad-platform" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, _ := PkgVersion()
	pulumi.RegisterResourceModule(
		"catalystsquad-platform",
		"index",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"catalystsquad-platform",
		&pkg{version},
	)
}
