// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CatalystsquadPlatform.Inputs
{

    public sealed class DirectorySourceArgs : Pulumi.ResourceArgs
    {
        [Input("exclude")]
        public Input<string>? Exclude { get; set; }

        [Input("include")]
        public Input<string>? Include { get; set; }

        [Input("jsonnet")]
        public Input<Inputs.DirectorySourceJsonnetArgs>? Jsonnet { get; set; }

        [Input("recurse")]
        public Input<bool>? Recurse { get; set; }

        public DirectorySourceArgs()
        {
        }
    }
}
