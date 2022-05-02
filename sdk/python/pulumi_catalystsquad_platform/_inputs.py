# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'AvailabilityZoneArgs',
    'EksNodeGroupArgs',
]

@pulumi.input_type
class AvailabilityZoneArgs:
    def __init__(__self__, *,
                 az_name: pulumi.Input[str],
                 private_subnet_cidr: Optional[pulumi.Input[str]] = None,
                 public_subnet_cidr: Optional[pulumi.Input[str]] = None):
        """
        Configuration supplied to AvailabilityZone list in VpcArgs to specify which availability zones to deploy to and what subnet configuration for each availability zone. Supports one private and public subnet per AZ.
        :param pulumi.Input[str] az_name: Name of the availability zone to deploy subnets to.
        :param pulumi.Input[str] private_subnet_cidr: CIDR for private subnets in the availability zone. If not supplied, the subnet is not created.
        :param pulumi.Input[str] public_subnet_cidr: CIDR for private subnets in the availability zone. If not supplied the subnet is not created.
        """
        pulumi.set(__self__, "az_name", az_name)
        if private_subnet_cidr is not None:
            pulumi.set(__self__, "private_subnet_cidr", private_subnet_cidr)
        if public_subnet_cidr is not None:
            pulumi.set(__self__, "public_subnet_cidr", public_subnet_cidr)

    @property
    @pulumi.getter(name="azName")
    def az_name(self) -> pulumi.Input[str]:
        """
        Name of the availability zone to deploy subnets to.
        """
        return pulumi.get(self, "az_name")

    @az_name.setter
    def az_name(self, value: pulumi.Input[str]):
        pulumi.set(self, "az_name", value)

    @property
    @pulumi.getter(name="privateSubnetCidr")
    def private_subnet_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        CIDR for private subnets in the availability zone. If not supplied, the subnet is not created.
        """
        return pulumi.get(self, "private_subnet_cidr")

    @private_subnet_cidr.setter
    def private_subnet_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "private_subnet_cidr", value)

    @property
    @pulumi.getter(name="publicSubnetCidr")
    def public_subnet_cidr(self) -> Optional[pulumi.Input[str]]:
        """
        CIDR for private subnets in the availability zone. If not supplied the subnet is not created.
        """
        return pulumi.get(self, "public_subnet_cidr")

    @public_subnet_cidr.setter
    def public_subnet_cidr(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "public_subnet_cidr", value)


@pulumi.input_type
class EksNodeGroupArgs:
    def __init__(__self__, *,
                 desired_size: pulumi.Input[int],
                 instance_types: pulumi.Input[Sequence[pulumi.Input[str]]],
                 max_size: pulumi.Input[int],
                 min_size: pulumi.Input[int],
                 name_prefix: pulumi.Input[str]):
        """
        Configuration for an EKS node group
        """
        pulumi.set(__self__, "desired_size", desired_size)
        pulumi.set(__self__, "instance_types", instance_types)
        pulumi.set(__self__, "max_size", max_size)
        pulumi.set(__self__, "min_size", min_size)
        pulumi.set(__self__, "name_prefix", name_prefix)

    @property
    @pulumi.getter(name="desiredSize")
    def desired_size(self) -> pulumi.Input[int]:
        return pulumi.get(self, "desired_size")

    @desired_size.setter
    def desired_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "desired_size", value)

    @property
    @pulumi.getter(name="instanceTypes")
    def instance_types(self) -> pulumi.Input[Sequence[pulumi.Input[str]]]:
        return pulumi.get(self, "instance_types")

    @instance_types.setter
    def instance_types(self, value: pulumi.Input[Sequence[pulumi.Input[str]]]):
        pulumi.set(self, "instance_types", value)

    @property
    @pulumi.getter(name="maxSize")
    def max_size(self) -> pulumi.Input[int]:
        return pulumi.get(self, "max_size")

    @max_size.setter
    def max_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "max_size", value)

    @property
    @pulumi.getter(name="minSize")
    def min_size(self) -> pulumi.Input[int]:
        return pulumi.get(self, "min_size")

    @min_size.setter
    def min_size(self, value: pulumi.Input[int]):
        pulumi.set(self, "min_size", value)

    @property
    @pulumi.getter(name="namePrefix")
    def name_prefix(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name_prefix")

    @name_prefix.setter
    def name_prefix(self, value: pulumi.Input[str]):
        pulumi.set(self, "name_prefix", value)


