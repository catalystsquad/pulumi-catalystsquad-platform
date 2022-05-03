package main

import (
	csp "github.com/catalystsquad/pulumi-catalystsquad-platform/sdk/go/catalystsquad-platform"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		// create a vpc for the eks cluster
		vpc, err := csp.NewVpc(ctx, "vpc", &csp.VpcArgs{
			AvailabilityZoneConfig: csp.AvailabilityZoneArray{
				csp.AvailabilityZoneArgs{
					AzName:            pulumi.String("us-east-1a"),
					PrivateSubnetCidr: pulumi.String("10.1.0.0/18"),
					PublicSubnetCidr:  pulumi.String("10.1.128.0/23"),
				},
				csp.AvailabilityZoneArgs{
					AzName:            pulumi.String("us-east-1b"),
					PrivateSubnetCidr: pulumi.String("10.1.64.0/18"),
					PublicSubnetCidr:  pulumi.String("10.1.130.0/23"),
				},
			},
			Cidr:                 pulumi.String("10.1.0.0/16"),
			EnableEksClusterTags: pulumi.Bool(true),
		})
		if err != nil {
			return err
		}

		// create the eks cluster
		_, err = csp.NewEks(ctx, "eks", &csp.EksArgs{
			NodeGroupConfig: csp.EksNodeGroupArray{
				csp.EksNodeGroupArgs{
					NamePrefix:  pulumi.String("default"),
					DesiredSize: pulumi.Int(1),
					MaxSize:     pulumi.Int(2),
					MinSize:     pulumi.Int(1),
					InstanceTypes: pulumi.StringArray{
						pulumi.String("t3.small"),
					},
				},
			},
			SubnetIDs: vpc.PrivateSubnetIDs,
		})
		if err != nil {
			return err
		}

		return nil
	})
}
