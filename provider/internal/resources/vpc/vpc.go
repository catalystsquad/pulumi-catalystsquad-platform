package vpc

import (
	"fmt"

	"github.com/pulumi/pulumi-aws/sdk/v5/go/aws/ec2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// VpcArgs supplies input for configuring vpc resources
type VpcArgs struct {
	// Optional, list of AvailabilityZones to create subnets in. Default: []
	AvailabilityZoneConfig []AvailabilityZone `pulumi:"availabilityZoneConfig"`
	// Optional, CIDR block of the VPC. Default: 10.0.0.0/16
	Cidr string `pulumi:"cidr"`
	// Optional, EKS cluster name, if VPC is used for EKS. Default: <stack name>
	EksClusterName string `pulumi:"eksClusterName"`
	// Optional, whether to enable required EKS cluster tags to subnets. Default: true
	EnableEksClusterTags *bool `pulumi:"enableEksClusterTags"`
	// Optional, Name tag value for VPC resource. Default: <stack name>
	Name string `pulumi:"name"`
	// Optional, tags to add to all resources. Default: {}
	Tags map[string]string `pulumi:"tags"`
}

// AvailabilityZone specifies subnet CIDR ranges for creating a public and
// private subnet inside an AZ.
type AvailabilityZone struct {
	// Required, name of availability zone, ex: us-east-1a
	AzName string `pulumi:"azName"`
	// Optional, CIDR block of private subnet, if not supplied subnet is not created
	PrivateSubnetCidr string `pulumi:"privateSubnetCidr"`
	// Optional, CIDR block of public subnet, if not supplied subnet is not created
	PublicSubnetCidr string `pulumi:"publicSubnetCidr"`
}

// Vpc pulumi component resource
type Vpc struct {
	pulumi.ResourceState

	VpcID            pulumi.StringOutput      `pulumi:"vpcID"`
	PublicSubnetIDs  pulumi.StringArrayOutput `pulumi:"publicSubnetIDs"`
	PrivateSubnetIDs pulumi.StringArrayOutput `pulumi:"privateSubnetIDs"`
	NatGatewayIPs    pulumi.StringArrayOutput `pulumi:"natGatewayIPs"`
}

// NewVpc creates a simple Vpc
func NewVpc(ctx *pulumi.Context, name string, args *VpcArgs, opts ...pulumi.ResourceOption) (*Vpc, error) {
	if args == nil {
		args = &VpcArgs{}
	}

	component := &Vpc{}
	err := ctx.RegisterComponentResource("catalystsquad-platform:index:Vpc", name, component, opts...)
	if err != nil {
		return nil, err
	}

	// default vpc arguments
	vpcName := ctx.Stack()
	if args.Name == "" {
		vpcName = args.Name
	}

	vpcCidr := "10.0.0.0/16"
	if args.Cidr != "" {
		vpcCidr = args.Cidr
	}

	availabilityZones := []AvailabilityZone{}
	if args.AvailabilityZoneConfig != nil {
		availabilityZones = args.AvailabilityZoneConfig
	}

	enableEksClusterTags := true
	if args.EnableEksClusterTags != nil {
		enableEksClusterTags = *args.EnableEksClusterTags
	}

	eksClusterName := vpcName
	if args.EksClusterName == "" {
		eksClusterName = args.EksClusterName
	}

	// create tag map for adding to all resources
	tags := make(map[string]string)

	// tags for individual resources
	vpcTags := make(map[string]string)
	publicSubnetTags := make(map[string]string)
	privateSubnetTags := make(map[string]string)

	// add user provided tags to all resources
	if args.Tags != nil {
		for k, v := range args.Tags {
			tags[k] = v
		}
		vpcTags = tags
		publicSubnetTags = tags
	}

	// set vpc Name tag
	vpcTags["Name"] = vpcName

	if enableEksClusterTags {
		// required eks tag for public load balancers
		publicSubnetTags["kubernetes.io/role/elb"] = "1"
		// required eks tag for all subnets
		publicSubnetTags[fmt.Sprintf("kubernetes.io/cluster/%s", eksClusterName)] = "owned"
		privateSubnetTags[fmt.Sprintf("kubernetes.io/cluster/%s", eksClusterName)] = "owned"
	}

	vpc, err := ec2.NewVpc(ctx, "vpc", &ec2.VpcArgs{
		CidrBlock: pulumi.String(vpcCidr),
		Tags:      pulumi.ToStringMap(vpcTags),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	// create internet gateway
	internetGateway, err := ec2.NewInternetGateway(ctx, "internet-gateway", &ec2.InternetGatewayArgs{
		VpcId: vpc.ID(),
	}, pulumi.Parent(component))
	if err != nil {
		return nil, err
	}

	var publicSubnetIDs []pulumi.StringOutput
	var privateSubnetIDs []pulumi.StringOutput
	var natGatewayIPs []pulumi.StringOutput
	for i, az := range availabilityZones {
		// declare natgateway id outside of public subnet if statement so that
		// it can be left out on private subnet if not created
		var natgatewayID pulumi.IDOutput

		if az.PublicSubnetCidr != "" {
			// create public subnets
			publicSubnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("public-subnet-%d", i), &ec2.SubnetArgs{
				VpcId:            vpc.ID(),
				CidrBlock:        pulumi.String(az.PublicSubnetCidr),
				AvailabilityZone: pulumi.String(az.AzName),
				Tags:             pulumi.ToStringMap(publicSubnetTags),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			publicSubnetIDs = append(publicSubnetIDs, publicSubnet.ID().ToStringOutput())

			// create public subnet route tables
			publicRouteTable, err := ec2.NewRouteTable(ctx, fmt.Sprintf("public-route-table-%d", i), &ec2.RouteTableArgs{
				VpcId: vpc.ID(),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			// default public route
			_, err = ec2.NewRoute(ctx, fmt.Sprintf("public-route-%d", i), &ec2.RouteArgs{
				RouteTableId:         publicRouteTable.ID(),
				DestinationCidrBlock: pulumi.String("0.0.0.0/0"),
				GatewayId:            internetGateway.ID(),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			// associate route table to new subnet
			_, err = ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("public-route-table-association-%d", i), &ec2.RouteTableAssociationArgs{
				SubnetId:     publicSubnet.ID(),
				RouteTableId: publicRouteTable.ID(),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			// create nat gateway public ip
			natGatewayIP, err := ec2.NewEip(ctx, fmt.Sprintf("elastic-ip-%d", i), &ec2.EipArgs{
				Vpc: pulumi.Bool(true),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			natGatewayIPs = append(natGatewayIPs, natGatewayIP.ID().ToStringOutput())

			// create nat gateway
			natGateway, err := ec2.NewNatGateway(ctx, fmt.Sprintf("nat-gateway-%d", i), &ec2.NatGatewayArgs{
				AllocationId: natGatewayIP.ID(),
				SubnetId:     publicSubnet.ID(),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}
			natgatewayID = natGateway.ID()
		}

		if az.PrivateSubnetCidr != "" {
			// create private subnets
			privateSubnet, err := ec2.NewSubnet(ctx, fmt.Sprintf("private-subnet-%d", i), &ec2.SubnetArgs{
				VpcId:            vpc.ID(),
				CidrBlock:        pulumi.String(az.PrivateSubnetCidr),
				AvailabilityZone: pulumi.String(az.AzName),
				Tags:             pulumi.ToStringMap(privateSubnetTags),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			privateSubnetIDs = append(privateSubnetIDs, privateSubnet.ID().ToStringOutput())

			// create private subnet route tables
			privateRouteTable, err := ec2.NewRouteTable(ctx, fmt.Sprintf("private-route-table-%d", i), &ec2.RouteTableArgs{
				VpcId: vpc.ID(),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			// default private route
			_, err = ec2.NewRoute(ctx, fmt.Sprintf("private-route-%d", i), &ec2.RouteArgs{
				RouteTableId:         privateRouteTable.ID(),
				DestinationCidrBlock: pulumi.String("0.0.0.0/0"),
				NatGatewayId:         natgatewayID,
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}

			// associate route table to new subnet
			_, err = ec2.NewRouteTableAssociation(ctx, fmt.Sprintf("private-route-table-association-%d", i), &ec2.RouteTableAssociationArgs{
				SubnetId:     privateSubnet.ID(),
				RouteTableId: privateRouteTable.ID(),
			}, pulumi.Parent(component))
			if err != nil {
				return nil, err
			}
		}
	}

	component.VpcID = vpc.ID().ToStringOutput()
	component.PublicSubnetIDs = pulumi.ToStringArrayOutput(publicSubnetIDs)
	component.PrivateSubnetIDs = pulumi.ToStringArrayOutput(privateSubnetIDs)
	component.NatGatewayIPs = pulumi.ToStringArrayOutput(natGatewayIPs)

	err = ctx.RegisterResourceOutputs(component, pulumi.Map{
		"vpcId":            vpc.ID().ToStringOutput(),
		"publicSubnetIDs":  pulumi.ToStringArrayOutput(publicSubnetIDs),
		"privateSubnetIDs": pulumi.ToStringArrayOutput(privateSubnetIDs),
		"natGatewayIPs":    pulumi.ToStringArrayOutput(natGatewayIPs),
	})
	if err != nil {
		return nil, err
	}

	return component, nil
}
