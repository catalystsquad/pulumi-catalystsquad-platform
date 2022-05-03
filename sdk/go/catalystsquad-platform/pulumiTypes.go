// Code generated by Pulumi SDK Generator DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package catalystsquadplatform

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// Configuration supplied to AvailabilityZone list in VpcArgs to specify which availability zones to deploy to and what subnet configuration for each availability zone. Supports one private and public subnet per AZ.
type AvailabilityZone struct {
	// Name of the availability zone to deploy subnets to.
	AzName string `pulumi:"azName"`
	// CIDR for private subnets in the availability zone. If not supplied, the subnet is not created.
	PrivateSubnetCidr *string `pulumi:"privateSubnetCidr"`
	// CIDR for private subnets in the availability zone. If not supplied the subnet is not created.
	PublicSubnetCidr *string `pulumi:"publicSubnetCidr"`
}

// AvailabilityZoneInput is an input type that accepts AvailabilityZoneArgs and AvailabilityZoneOutput values.
// You can construct a concrete instance of `AvailabilityZoneInput` via:
//
//          AvailabilityZoneArgs{...}
type AvailabilityZoneInput interface {
	pulumi.Input

	ToAvailabilityZoneOutput() AvailabilityZoneOutput
	ToAvailabilityZoneOutputWithContext(context.Context) AvailabilityZoneOutput
}

// Configuration supplied to AvailabilityZone list in VpcArgs to specify which availability zones to deploy to and what subnet configuration for each availability zone. Supports one private and public subnet per AZ.
type AvailabilityZoneArgs struct {
	// Name of the availability zone to deploy subnets to.
	AzName pulumi.StringInput `pulumi:"azName"`
	// CIDR for private subnets in the availability zone. If not supplied, the subnet is not created.
	PrivateSubnetCidr pulumi.StringPtrInput `pulumi:"privateSubnetCidr"`
	// CIDR for private subnets in the availability zone. If not supplied the subnet is not created.
	PublicSubnetCidr pulumi.StringPtrInput `pulumi:"publicSubnetCidr"`
}

func (AvailabilityZoneArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityZone)(nil)).Elem()
}

func (i AvailabilityZoneArgs) ToAvailabilityZoneOutput() AvailabilityZoneOutput {
	return i.ToAvailabilityZoneOutputWithContext(context.Background())
}

func (i AvailabilityZoneArgs) ToAvailabilityZoneOutputWithContext(ctx context.Context) AvailabilityZoneOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityZoneOutput)
}

// AvailabilityZoneArrayInput is an input type that accepts AvailabilityZoneArray and AvailabilityZoneArrayOutput values.
// You can construct a concrete instance of `AvailabilityZoneArrayInput` via:
//
//          AvailabilityZoneArray{ AvailabilityZoneArgs{...} }
type AvailabilityZoneArrayInput interface {
	pulumi.Input

	ToAvailabilityZoneArrayOutput() AvailabilityZoneArrayOutput
	ToAvailabilityZoneArrayOutputWithContext(context.Context) AvailabilityZoneArrayOutput
}

type AvailabilityZoneArray []AvailabilityZoneInput

func (AvailabilityZoneArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AvailabilityZone)(nil)).Elem()
}

func (i AvailabilityZoneArray) ToAvailabilityZoneArrayOutput() AvailabilityZoneArrayOutput {
	return i.ToAvailabilityZoneArrayOutputWithContext(context.Background())
}

func (i AvailabilityZoneArray) ToAvailabilityZoneArrayOutputWithContext(ctx context.Context) AvailabilityZoneArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityZoneArrayOutput)
}

// Configuration supplied to AvailabilityZone list in VpcArgs to specify which availability zones to deploy to and what subnet configuration for each availability zone. Supports one private and public subnet per AZ.
type AvailabilityZoneOutput struct{ *pulumi.OutputState }

func (AvailabilityZoneOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityZone)(nil)).Elem()
}

func (o AvailabilityZoneOutput) ToAvailabilityZoneOutput() AvailabilityZoneOutput {
	return o
}

func (o AvailabilityZoneOutput) ToAvailabilityZoneOutputWithContext(ctx context.Context) AvailabilityZoneOutput {
	return o
}

// Name of the availability zone to deploy subnets to.
func (o AvailabilityZoneOutput) AzName() pulumi.StringOutput {
	return o.ApplyT(func(v AvailabilityZone) string { return v.AzName }).(pulumi.StringOutput)
}

// CIDR for private subnets in the availability zone. If not supplied, the subnet is not created.
func (o AvailabilityZoneOutput) PrivateSubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityZone) *string { return v.PrivateSubnetCidr }).(pulumi.StringPtrOutput)
}

// CIDR for private subnets in the availability zone. If not supplied the subnet is not created.
func (o AvailabilityZoneOutput) PublicSubnetCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AvailabilityZone) *string { return v.PublicSubnetCidr }).(pulumi.StringPtrOutput)
}

type AvailabilityZoneArrayOutput struct{ *pulumi.OutputState }

func (AvailabilityZoneArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AvailabilityZone)(nil)).Elem()
}

func (o AvailabilityZoneArrayOutput) ToAvailabilityZoneArrayOutput() AvailabilityZoneArrayOutput {
	return o
}

func (o AvailabilityZoneArrayOutput) ToAvailabilityZoneArrayOutputWithContext(ctx context.Context) AvailabilityZoneArrayOutput {
	return o
}

func (o AvailabilityZoneArrayOutput) Index(i pulumi.IntInput) AvailabilityZoneOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AvailabilityZone {
		return vs[0].([]AvailabilityZone)[vs[1].(int)]
	}).(AvailabilityZoneOutput)
}

// Configuration for an EKS node group
type EksNodeGroup struct {
	DesiredSize   int      `pulumi:"desiredSize"`
	InstanceTypes []string `pulumi:"instanceTypes"`
	MaxSize       int      `pulumi:"maxSize"`
	MinSize       int      `pulumi:"minSize"`
	NamePrefix    string   `pulumi:"namePrefix"`
}

// EksNodeGroupInput is an input type that accepts EksNodeGroupArgs and EksNodeGroupOutput values.
// You can construct a concrete instance of `EksNodeGroupInput` via:
//
//          EksNodeGroupArgs{...}
type EksNodeGroupInput interface {
	pulumi.Input

	ToEksNodeGroupOutput() EksNodeGroupOutput
	ToEksNodeGroupOutputWithContext(context.Context) EksNodeGroupOutput
}

// Configuration for an EKS node group
type EksNodeGroupArgs struct {
	DesiredSize   pulumi.IntInput         `pulumi:"desiredSize"`
	InstanceTypes pulumi.StringArrayInput `pulumi:"instanceTypes"`
	MaxSize       pulumi.IntInput         `pulumi:"maxSize"`
	MinSize       pulumi.IntInput         `pulumi:"minSize"`
	NamePrefix    pulumi.StringInput      `pulumi:"namePrefix"`
}

func (EksNodeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EksNodeGroup)(nil)).Elem()
}

func (i EksNodeGroupArgs) ToEksNodeGroupOutput() EksNodeGroupOutput {
	return i.ToEksNodeGroupOutputWithContext(context.Background())
}

func (i EksNodeGroupArgs) ToEksNodeGroupOutputWithContext(ctx context.Context) EksNodeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EksNodeGroupOutput)
}

// EksNodeGroupArrayInput is an input type that accepts EksNodeGroupArray and EksNodeGroupArrayOutput values.
// You can construct a concrete instance of `EksNodeGroupArrayInput` via:
//
//          EksNodeGroupArray{ EksNodeGroupArgs{...} }
type EksNodeGroupArrayInput interface {
	pulumi.Input

	ToEksNodeGroupArrayOutput() EksNodeGroupArrayOutput
	ToEksNodeGroupArrayOutputWithContext(context.Context) EksNodeGroupArrayOutput
}

type EksNodeGroupArray []EksNodeGroupInput

func (EksNodeGroupArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EksNodeGroup)(nil)).Elem()
}

func (i EksNodeGroupArray) ToEksNodeGroupArrayOutput() EksNodeGroupArrayOutput {
	return i.ToEksNodeGroupArrayOutputWithContext(context.Background())
}

func (i EksNodeGroupArray) ToEksNodeGroupArrayOutputWithContext(ctx context.Context) EksNodeGroupArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EksNodeGroupArrayOutput)
}

// Configuration for an EKS node group
type EksNodeGroupOutput struct{ *pulumi.OutputState }

func (EksNodeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EksNodeGroup)(nil)).Elem()
}

func (o EksNodeGroupOutput) ToEksNodeGroupOutput() EksNodeGroupOutput {
	return o
}

func (o EksNodeGroupOutput) ToEksNodeGroupOutputWithContext(ctx context.Context) EksNodeGroupOutput {
	return o
}

func (o EksNodeGroupOutput) DesiredSize() pulumi.IntOutput {
	return o.ApplyT(func(v EksNodeGroup) int { return v.DesiredSize }).(pulumi.IntOutput)
}

func (o EksNodeGroupOutput) InstanceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EksNodeGroup) []string { return v.InstanceTypes }).(pulumi.StringArrayOutput)
}

func (o EksNodeGroupOutput) MaxSize() pulumi.IntOutput {
	return o.ApplyT(func(v EksNodeGroup) int { return v.MaxSize }).(pulumi.IntOutput)
}

func (o EksNodeGroupOutput) MinSize() pulumi.IntOutput {
	return o.ApplyT(func(v EksNodeGroup) int { return v.MinSize }).(pulumi.IntOutput)
}

func (o EksNodeGroupOutput) NamePrefix() pulumi.StringOutput {
	return o.ApplyT(func(v EksNodeGroup) string { return v.NamePrefix }).(pulumi.StringOutput)
}

type EksNodeGroupArrayOutput struct{ *pulumi.OutputState }

func (EksNodeGroupArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EksNodeGroup)(nil)).Elem()
}

func (o EksNodeGroupArrayOutput) ToEksNodeGroupArrayOutput() EksNodeGroupArrayOutput {
	return o
}

func (o EksNodeGroupArrayOutput) ToEksNodeGroupArrayOutputWithContext(ctx context.Context) EksNodeGroupArrayOutput {
	return o
}

func (o EksNodeGroupArrayOutput) Index(i pulumi.IntInput) EksNodeGroupOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EksNodeGroup {
		return vs[0].([]EksNodeGroup)[vs[1].(int)]
	}).(EksNodeGroupOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AvailabilityZoneInput)(nil)).Elem(), AvailabilityZoneArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AvailabilityZoneArrayInput)(nil)).Elem(), AvailabilityZoneArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EksNodeGroupInput)(nil)).Elem(), EksNodeGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EksNodeGroupArrayInput)(nil)).Elem(), EksNodeGroupArray{})
	pulumi.RegisterOutputType(AvailabilityZoneOutput{})
	pulumi.RegisterOutputType(AvailabilityZoneArrayOutput{})
	pulumi.RegisterOutputType(EksNodeGroupOutput{})
	pulumi.RegisterOutputType(EksNodeGroupArrayOutput{})
}