package main

import (
	csp "github.com/catalystsquad/pulumi-catalystsquad-platform/sdk/go/catalystsquad-platform"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := csp.NewVpc(ctx, "vpc", &csp.VpcArgs{
			AvailabilityZoneConfig: csp.AvailabilityZoneArray{
				csp.AvailabilityZoneArgs{
					AzName:            pulumi.String("us-east-1a"),
					PrivateSubnetCidr: pulumi.String("10.1.0.0/18"),
					PublicSubnetCidr:  pulumi.String("10.1.128.0/23"),
				},
			},
			Cidr:                 pulumi.String("10.1.0.0/16"),
			EnableEksClusterTags: pulumi.Bool(true),
		})
		return err
	})
}
