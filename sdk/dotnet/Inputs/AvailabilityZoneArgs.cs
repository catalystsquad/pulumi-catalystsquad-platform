// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CatalystsquadPlatform.Inputs
{

    /// <summary>
    /// Configuration supplied to AvailabilityZone list in VpcArgs to specify which availability zones to deploy to and what subnet configuration for each availability zone. Supports one private and public subnet per AZ.
    /// </summary>
    public sealed class AvailabilityZoneArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the availability zone to deploy subnets to.
        /// </summary>
        [Input("azName", required: true)]
        public Input<string> AzName { get; set; } = null!;

        /// <summary>
        /// CIDR for private subnets in the availability zone. If not supplied, the subnet is not created.
        /// </summary>
        [Input("privateSubnetCidr")]
        public Input<string>? PrivateSubnetCidr { get; set; }

        /// <summary>
        /// CIDR for private subnets in the availability zone. If not supplied the subnet is not created.
        /// </summary>
        [Input("publicSubnetCidr")]
        public Input<string>? PublicSubnetCidr { get; set; }

        public AvailabilityZoneArgs()
        {
        }
    }
}
