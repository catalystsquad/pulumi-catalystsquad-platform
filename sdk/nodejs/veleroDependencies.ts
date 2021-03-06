// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

export class VeleroDependencies extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'catalystsquad-platform:index:VeleroDependencies';

    /**
     * Returns true if the given object is an instance of VeleroDependencies.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is VeleroDependencies {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === VeleroDependencies.__pulumiType;
    }


    /**
     * Create a VeleroDependencies resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: VeleroDependenciesArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.oidcProviderArn === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcProviderArn'");
            }
            if ((!args || args.oidcProviderUrl === undefined) && !opts.urn) {
                throw new Error("Missing required property 'oidcProviderUrl'");
            }
            resourceInputs["createBucket"] = args ? args.createBucket : undefined;
            resourceInputs["oidcProviderArn"] = args ? args.oidcProviderArn : undefined;
            resourceInputs["oidcProviderUrl"] = args ? args.oidcProviderUrl : undefined;
            resourceInputs["veleroBucketName"] = args ? args.veleroBucketName : undefined;
            resourceInputs["veleroIAMPolicyName"] = args ? args.veleroIAMPolicyName : undefined;
            resourceInputs["veleroIAMRoleName"] = args ? args.veleroIAMRoleName : undefined;
            resourceInputs["veleroNamespace"] = args ? args.veleroNamespace : undefined;
            resourceInputs["veleroServiceAccount"] = args ? args.veleroServiceAccount : undefined;
        } else {
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(VeleroDependencies.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a VeleroDependencies resource.
 */
export interface VeleroDependenciesArgs {
    /**
     * Optional, whether to create the Velero S3 bucket. Allows the bucket to exist outside of pulumi. Default: true
     */
    createBucket?: pulumi.Input<boolean>;
    /**
     * Required, Arn of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.
     */
    oidcProviderArn: pulumi.Input<string>;
    /**
     * Required, URL of EKS OIDC Provider for configuring the IRSA  IAM role trust relationship.
     */
    oidcProviderUrl: pulumi.Input<string>;
    /**
     * Optional, Velero's bucket name. Default: <account-id>-<stack-name>-velero
     */
    veleroBucketName?: pulumi.Input<string>;
    /**
     * Optional, Velero's IAM policy name. Default: <stack-name>-velero-policy
     */
    veleroIAMPolicyName?: pulumi.Input<string>;
    /**
     * Optional, Velero's IAM role name. Default: <stack-name>-velero-role
     */
    veleroIAMRoleName?: pulumi.Input<string>;
    /**
     * Optional, kubernetes namespace where Velero will exist, for configuring the IRSA IAM role trust relationship. Default: velero
     */
    veleroNamespace?: pulumi.Input<string>;
    /**
     * Optional, kubernetes service account name that Velero will use, for configuring the IRSA IAM role trust relationship. Default: velero
     */
    veleroServiceAccount?: pulumi.Input<string>;
}
