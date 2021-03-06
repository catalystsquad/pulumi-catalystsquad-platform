// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.CatalystsquadPlatform.Inputs
{

    public sealed class HelmSourceArgs : Pulumi.ResourceArgs
    {
        [Input("fileParameters")]
        private InputList<Inputs.HelmSourceFileParameterArgs>? _fileParameters;
        public InputList<Inputs.HelmSourceFileParameterArgs> FileParameters
        {
            get => _fileParameters ?? (_fileParameters = new InputList<Inputs.HelmSourceFileParameterArgs>());
            set => _fileParameters = value;
        }

        [Input("ignoreMissingValueFiles")]
        public Input<bool>? IgnoreMissingValueFiles { get; set; }

        [Input("parameters")]
        private InputList<Inputs.HelmSourceParameterArgs>? _parameters;
        public InputList<Inputs.HelmSourceParameterArgs> Parameters
        {
            get => _parameters ?? (_parameters = new InputList<Inputs.HelmSourceParameterArgs>());
            set => _parameters = value;
        }

        [Input("passCredentials")]
        public Input<bool>? PassCredentials { get; set; }

        [Input("releaseName")]
        public Input<string>? ReleaseName { get; set; }

        [Input("skipCrds")]
        public Input<string>? SkipCrds { get; set; }

        [Input("valueFiles")]
        private InputList<string>? _valueFiles;
        public InputList<string> ValueFiles
        {
            get => _valueFiles ?? (_valueFiles = new InputList<string>());
            set => _valueFiles = value;
        }

        [Input("values")]
        public Input<string>? Values { get; set; }

        [Input("version")]
        public Input<string>? Version { get; set; }

        public HelmSourceArgs()
        {
        }
    }
}
