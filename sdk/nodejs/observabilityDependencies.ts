// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class ObservabilityDependencies extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'catalystsquad-platform:index:ObservabilityDependencies';

    /**
     * Returns true if the given object is an instance of ObservabilityDependencies.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is ObservabilityDependencies {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === ObservabilityDependencies.__pulumiType;
    }


    /**
     * Create a ObservabilityDependencies resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: ObservabilityDependenciesArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.oidcProviderArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcProviderArn'");
            }
            if ((!args || args.oidcProviderUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcProviderUrl'");
            }
            resourceInputs["cortexBucketName"] = args ? args.cortexBucketName : undefined;
            resourceInputs["cortexNamespace"] = args ? args.cortexNamespace : undefined;
            resourceInputs["cortexServiceAccount"] = args ? args.cortexServiceAccount : undefined;
            resourceInputs["lokiBucketName"] = args ? args.lokiBucketName : undefined;
            resourceInputs["lokiNamespace"] = args ? args.lokiNamespace : undefined;
            resourceInputs["lokiServiceAccount"] = args ? args.lokiServiceAccount : undefined;
            resourceInputs["oidcProviderArn"] = args ? args.oidcProviderArn : undefined;
            resourceInputs["oidcProviderUrl"] = args ? args.oidcProviderUrl : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(ObservabilityDependencies.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a ObservabilityDependencies resource.
 */
export interface ObservabilityDependenciesArgs {
    /**
     * Optional, name of bucket to create for Cortex. Default: <account-id>-<stack-name>-cortex
     */
    cortexBucketName?: pulumi.Input<string>;
    /**
     * Optional, kubernetes namespace where Cortex will exist, for configuring the IRSA IAM role trust relationship. Default: cortex
     */
    cortexNamespace?: pulumi.Input<string>;
    /**
     * Optional, kubernetes service account name that Cortex will use, for configuring the IRSA IAM role trust relationship. Default: cortex
     */
    cortexServiceAccount?: pulumi.Input<string>;
    /**
     * Optional, name of bucket to create for Loki. Default: <account-id>-<stack-name>-loki
     */
    lokiBucketName?: pulumi.Input<string>;
    /**
     * Optional, kubernetes namespace where Loki will exist, for configuring the IRSA IAM role trust relationship. Default: loki
     */
    lokiNamespace?: pulumi.Input<string>;
    /**
     * Optional, kubernetes service account name that Loki will use, for configuring the IRSA IAM role trust relationship. Default: loki
     */
    lokiServiceAccount?: pulumi.Input<string>;
    /**
     * TODO FIXME
     */
    oidcProviderArn: pulumi.Input<string>;
    /**
     * TODO FIXME
     */
    oidcProviderUrl: pulumi.Input<string>;
}
