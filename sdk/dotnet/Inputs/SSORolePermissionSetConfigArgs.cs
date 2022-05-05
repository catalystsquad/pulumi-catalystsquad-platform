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
    /// Configuration of SSO IAM Roles to auto discover.
    /// </summary>
    public sealed class SSORolePermissionSetConfigArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Name of the permission set. Will use for autodiscovery using regex "AWSReservedSSO_&lt;name&gt;_.*"
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("permissionGroups")]
        private InputList<string>? _permissionGroups;

        /// <summary>
        /// List of permission groups to add to each identity. Ex: system:masters
        /// </summary>
        public InputList<string> PermissionGroups
        {
            get => _permissionGroups ?? (_permissionGroups = new InputList<string>());
            set => _permissionGroups = value;
        }

        /// <summary>
        /// Optional username field, defaults to the name of the SSO role.
        /// </summary>
        [Input("username")]
        public Input<string>? Username { get; set; }

        public SSORolePermissionSetConfigArgs()
        {
        }
    }
}
