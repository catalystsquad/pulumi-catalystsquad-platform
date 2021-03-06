// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CatalystsquadPlatform
{
    [CatalystsquadPlatformResourceType("catalystsquad-platform:index:ClusterBootstrap")]
    public partial class ClusterBootstrap : Pulumi.ComponentResource
    {
        /// <summary>
        /// Create a ClusterBootstrap resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ClusterBootstrap(string name, ClusterBootstrapArgs? args = null, ComponentResourceOptions? options = null)
            : base("catalystsquad-platform:index:ClusterBootstrap", name, args ?? new ClusterBootstrapArgs(), MakeResourceOptions(options, ""), remote: true)
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

    public sealed class ClusterBootstrapArgs : Pulumi.ResourceArgs
    {
        /// <summary>
        /// Optional, configures the argocd helm release.
        /// </summary>
        [Input("argocdHelmConfig")]
        public Input<Inputs.HelmReleaseConfigArgs>? ArgocdHelmConfig { get; set; }

        /// <summary>
        /// Optional, configures the kube-prometheus-stack helm release.
        /// </summary>
        [Input("kubePrometheusStackHelmConfig")]
        public Input<Inputs.HelmReleaseConfigArgs>? KubePrometheusStackHelmConfig { get; set; }

        /// <summary>
        /// Optional, configures the platform application release. Does not deploy if not specified.
        /// </summary>
        [Input("platformApplicationConfig")]
        public Input<Inputs.PlatformApplicationConfigArgs>? PlatformApplicationConfig { get; set; }

        /// <summary>
        /// Optional, configuration for a prometheus remoteWrite secret. Does not deploy if not specified.
        /// </summary>
        [Input("prometheusRemoteWriteConfig")]
        public Input<Inputs.PrometheusRemoteWriteConfigArgs>? PrometheusRemoteWriteConfig { get; set; }

        public ClusterBootstrapArgs()
        {
        }
    }
}
