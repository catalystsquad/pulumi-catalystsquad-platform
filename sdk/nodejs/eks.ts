// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import { input as inputs, output as outputs } from "./types";
import * as utilities from "./utilities";

import * as pulumiAws from "@pulumi/aws";
import * as pulumiKubernetes from "@pulumi/kubernetes";

export class Eks extends pulumi.ComponentResource {
    /** @internal */
    public static readonly __pulumiType = 'catalystsquad-platform:index:Eks';

    /**
     * Returns true if the given object is an instance of Eks.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is Eks {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === Eks.__pulumiType;
    }

    public /*out*/ readonly cluster!: pulumi.Output<pulumiAws.eks.Cluster>;
    public /*out*/ readonly kubeConfig!: pulumi.Output<string>;
    public /*out*/ readonly kubernetesProvider!: pulumi.Output<pulumiKubernetes.Provider | undefined>;
    public /*out*/ readonly nodeGroupIAMRoleArn!: pulumi.Output<string | undefined>;
    public /*out*/ readonly oidcProvider!: pulumi.Output<pulumiAws.iam.OpenIdConnectProvider>;

    /**
     * Create a Eks resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: EksArgs, opts?: pulumi.ComponentResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (!opts.id) {
            if ((!args || args.nodeGroupConfig === undefined) && !opts.urn) {
                throw new Error("Missing required property 'nodeGroupConfig'");
            }
            if ((!args || args.subnetIDs === undefined) && !opts.urn) {
                throw new Error("Missing required property 'subnetIDs'");
            }
            resourceInputs["authConfigmapConfig"] = args ? args.authConfigmapConfig : undefined;
            resourceInputs["clusterAutoscalerNamespace"] = args ? args.clusterAutoscalerNamespace : undefined;
            resourceInputs["clusterAutoscalerServiceAccount"] = args ? args.clusterAutoscalerServiceAccount : undefined;
            resourceInputs["clusterName"] = args ? args.clusterName : undefined;
            resourceInputs["enableClusterAutoscalerResources"] = args ? args.enableClusterAutoscalerResources : undefined;
            resourceInputs["enableECRAccess"] = args ? args.enableECRAccess : undefined;
            resourceInputs["enabledClusterLogTypes"] = args ? args.enabledClusterLogTypes : undefined;
            resourceInputs["k8sVersion"] = args ? args.k8sVersion : undefined;
            resourceInputs["kubeConfigAssumeRoleArn"] = args ? args.kubeConfigAssumeRoleArn : undefined;
            resourceInputs["kubeConfigAwsProfile"] = args ? args.kubeConfigAwsProfile : undefined;
            resourceInputs["nodeGroupConfig"] = args ? args.nodeGroupConfig : undefined;
            resourceInputs["nodeGroupVersion"] = args ? args.nodeGroupVersion : undefined;
            resourceInputs["subnetIDs"] = args ? args.subnetIDs : undefined;
            resourceInputs["cluster"] = undefined /*out*/;
            resourceInputs["kubeConfig"] = undefined /*out*/;
            resourceInputs["kubernetesProvider"] = undefined /*out*/;
            resourceInputs["nodeGroupIAMRoleArn"] = undefined /*out*/;
            resourceInputs["oidcProvider"] = undefined /*out*/;
        } else {
            resourceInputs["cluster"] = undefined /*out*/;
            resourceInputs["kubeConfig"] = undefined /*out*/;
            resourceInputs["kubernetesProvider"] = undefined /*out*/;
            resourceInputs["nodeGroupIAMRoleArn"] = undefined /*out*/;
            resourceInputs["oidcProvider"] = undefined /*out*/;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(Eks.__pulumiType, name, resourceInputs, opts, true /*remote*/);
    }
}

/**
 * The set of arguments for constructing a Eks resource.
 */
export interface EksArgs {
    /**
     * Optional, configures management of the eks auth configmap.
     */
    authConfigmapConfig?: pulumi.Input<inputs.AuthConfigMapConfigArgs>;
    /**
     * Optional, cluster autoscaler namespace for IRSA. Default: cluster-autoscaler
     */
    clusterAutoscalerNamespace?: pulumi.Input<string>;
    /**
     * Optional, cluster autoscaler service account name for IRSA. Default: cluster-autoscaler
     */
    clusterAutoscalerServiceAccount?: pulumi.Input<string>;
    /**
     * Optional, name of the EKS cluster. Default: <stack name>
     */
    clusterName?: pulumi.Input<string>;
    /**
     * Optional, whether to enable cluster autoscaler IRSA resources. Default: true
     */
    enableClusterAutoscalerResources?: pulumi.Input<boolean>;
    /**
     * Optional, whether to enable ECR access policy on nodegroups. Default: true
     */
    enableECRAccess?: pulumi.Input<boolean>;
    /**
     * Optional, list of log types to enable on the cluster. Default: []
     */
    enabledClusterLogTypes?: pulumi.Input<string>;
    /**
     * Optional, k8s version of the EKS cluster. Default: 1.22.6
     */
    k8sVersion?: pulumi.Input<string>;
    /**
     * Optional, assume role arn to add to the kubeconfig.
     */
    kubeConfigAssumeRoleArn?: pulumi.Input<string>;
    /**
     * Optional, AWS profile to add to the kubeconfig.
     */
    kubeConfigAwsProfile?: pulumi.Input<string>;
    /**
     * Required, list of nodegroup configurations to create.
     */
    nodeGroupConfig: pulumi.Input<pulumi.Input<inputs.EksNodeGroupArgs>[]>;
    /**
     * Optional, k8s version of all node groups. Allows for upgrading the control plane before upgrading nodegroups. Default: <k8sVersion>
     */
    nodeGroupVersion?: pulumi.Input<string>;
    /**
     * Required, list of subnet IDs to deploy the cluster and nodegroups to
     */
    subnetIDs: pulumi.Input<pulumi.Input<string>[]>;
}
