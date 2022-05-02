// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "../types";

/**
 * Configuration supplied to AvailabilityZone list in VpcArgs to specify which availability zones to deploy to and what subnet configuration for each availability zone. Supports one private and public subnet per AZ.
 */
export interface AvailabilityZoneArgs {
    /**
     * Name of the availability zone to deploy subnets to.
     */
    azName: pulumi.Input<string>;
    /**
     * CIDR for private subnets in the availability zone. If not supplied, the subnet is not created.
     */
    privateSubnetCidr?: pulumi.Input<string>;
    /**
     * CIDR for private subnets in the availability zone. If not supplied the subnet is not created.
     */
    publicSubnetCidr?: pulumi.Input<string>;
}

/**
 * Configuration for an EKS node group
 */
export interface EksNodeGroupArgs {
    desiredSize: pulumi.Input<number>;
    instanceTypes: pulumi.Input<pulumi.Input<string>[]>;
    maxSize: pulumi.Input<number>;
    minSize: pulumi.Input<number>;
    namePrefix: pulumi.Input<string>;
}
