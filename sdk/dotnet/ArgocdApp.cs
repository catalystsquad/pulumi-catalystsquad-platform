// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CatalystsquadPlatform
{
    [CatalystsquadPlatformResourceType("catalystsquad-platform:index:ArgocdApp")]
    public partial class ArgocdApp : Pulumi.ComponentResource
    {
        /// <summary>
        /// Create a ArgocdApp resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ArgocdApp(string name, ArgocdAppArgs args, ComponentResourceOptions? options = null)
            : base("catalystsquad-platform:index:ArgocdApp", name, args ?? new ArgocdAppArgs(), MakeResourceOptions(options, ""), remote: true)
        {
        }

        private static ComponentResourceOptions MakeResourceOptions(ComponentResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new ComponentResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "https://github.com/catalystsquad/pulumi-catalystsquad-platform/releases/download/v${VERSION}",
            };
            var merged = ComponentResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
    }

    public sealed class ArgocdAppArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional, apiVersion of the Argocd Application. Default: v1alpha1
        /// </summary>
        [Input("apiVersion")]
        public Input<string>? ApiVersion { get; set; }

        /// <summary>
        /// Required, name of the Argocd Application
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        /// <summary>
        /// Optional, namespace to deploy Argocd Application to. Should be the namespace where the argocd server runs. Default: "argo-cd"
        /// </summary>
        [Input("namespace")]
        public Input<string>? Namespace { get; set; }

        /// <summary>
        /// Required, spec of the Argocd Application
        /// </summary>
        [Input("spec")]
        public Input<Inputs.ArgocdApplicationArgs>? Spec { get; set; }

        public ArgocdAppArgs()
        {
        }
    }
}
