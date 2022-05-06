package main

import (
	csp "github.com/catalystsquad/pulumi-catalystsquad-platform/sdk/go/catalystsquad-platform"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// this example assumes that there is already a kubeconfig present
		// during runtime

		// install bootstrap to cluster with default values for platform
		// application chart
		_, err := csp.NewClusterBootstrap(ctx, "cluster-boostrap", &csp.ClusterBootstrapArgs{
			PlatformApplicationConfig: &csp.PlatformApplicationConfigArgs{
				CertManagerDnsSolverSecret: pulumi.String("someValue"),
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
}
