// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CatalystsquadPlatform.Inputs
{

    public sealed class ArgocdApplicationSpecDestinationArgs : Pulumi.ResourceArgs
    {
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        [Input("server")]
        public Input<string>? Server { get; set; }

        public ArgocdApplicationSpecDestinationArgs()
        {
        }
    }
}