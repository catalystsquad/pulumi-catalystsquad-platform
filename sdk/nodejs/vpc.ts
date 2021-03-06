// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

export class Vpc extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'catalystsquad-platform:index:Vpc';

    /**
     * Returns true if the given object is an instance of Vpc.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Vpc {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Vpc.__pulumiType;
    }

    /**
     * Static IPs for all NAT gateways
     */
    public /*out*/ readonly natGatewayIPs!: pulumi.Output<string[] | undefined>;
    /**
     * IDs for all private subnets
     */
    public /*out*/ readonly privateSubnetIDs!: pulumi.Output<string[] | undefined>;
    /**
     * IDs for all public subnets
     */
    public /*out*/ readonly publicSubnetIDs!: pulumi.Output<string[] | undefined>;
    /**
     * CIDR for the VPC
     */
    public /*out*/ readonly vpcID!: pulumi.Output<string>;

    /**
     * Create a Vpc resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args?: VpcArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            resourceInputs["availabilityZoneConfig"] = args ? args.availabilityZoneConfig : undefined;
            resourceInputs["cidr"] = args ? args.cidr : undefined;
            resourceInputs["eksClusterName"] = args ? args.eksClusterName : undefined;
            resourceInputs["enableEksClusterTags"] = args ? args.enableEksClusterTags : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["tags"] = args ? args.tags : undefined;
            resourceInputs["natGatewayIPs"] = undefined /*out*/;
            resourceInputs["privateSubnetIDs"] = undefined /*out*/;
            resourceInputs["publicSubnetIDs"] = undefined /*out*/;
            resourceInputs["vpcID"] = undefined /*out*/;
        } else {
            resourceInputs["natGatewayIPs"] = undefined /*out*/;
            resourceInputs["privateSubnetIDs"] = undefined /*out*/;
            resourceInputs["publicSubnetIDs"] = undefined /*out*/;
            resourceInputs["vpcID"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Vpc.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Vpc resource.
 */
export interface VpcArgs {
    /**
     * Optional, list of AvailabilityZones to create subnets in. Default: []
     */
    availabilityZoneConfig?: pulumi.Input<pulumi.Input<inputs.AvailabilityZoneArgs>[]>;
    /**
     * Optional, CIDR block of the VPC. Default: 10.0.0.0/16
     */
    cidr?: pulumi.Input<string>;
    /**
     * Optional, EKS cluster name, if VPC is used for EKS. Default: <stack name>
     */
    eksClusterName?: pulumi.Input<string>;
    /**
     * Optional, whether to enable required EKS cluster tags to subnets. Default: true
     */
    enableEksClusterTags?: pulumi.Input<boolean>;
    /**
     * Optional, Name tag value for VPC resource. Default: <stack name>
     */
    name?: pulumi.Input<string>;
    /**
     * Optional, tags to add to all resources. Default: {}
     */
    tags?: pulumi.Input<{[key: string]: pulumi.Input<string>}>;
}
