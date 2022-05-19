// *** WARNING: this file was generated by Pulumi SDK Generator. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

// Export members:
export * from "./argocdApp";
export * from "./clusterBootstrap";
export * from "./eks";
export * from "./observabilityDependencies";
export * from "./provider";
export * from "./veleroDependencies";
export * from "./vpc";

// Export sub-modules:
import * as types from "./types";

export {
    types,
};

// Import resources to register:
import { ArgocdApp } from "./argocdApp";
import { ClusterBootstrap } from "./clusterBootstrap";
import { Eks } from "./eks";
import { ObservabilityDependencies } from "./observabilityDependencies";
import { VeleroDependencies } from "./veleroDependencies";
import { Vpc } from "./vpc";

const _module = {
    version: utilities.getVersion(),
    construct: (name: string, type: string, urn: string): pulumi.Resource => {
        switch (type) {
            case "catalystsquad-platform:index:ArgocdApp":
                return new ArgocdApp(name, <any>undefined, { urn })
            case "catalystsquad-platform:index:ClusterBootstrap":
                return new ClusterBootstrap(name, <any>undefined, { urn })
            case "catalystsquad-platform:index:Eks":
                return new Eks(name, <any>undefined, { urn })
            case "catalystsquad-platform:index:ObservabilityDependencies":
                return new ObservabilityDependencies(name, <any>undefined, { urn })
            case "catalystsquad-platform:index:VeleroDependencies":
                return new VeleroDependencies(name, <any>undefined, { urn })
            case "catalystsquad-platform:index:Vpc":
                return new Vpc(name, <any>undefined, { urn })
            default:
                throw new Error(`unknown resource type ${type}`);
        }
    },
};
pulumi.runtime.registerResourceModule("catalystsquad-platform", "index", _module)

import { Provider } from "./provider";

pulumi.runtime.registerResourcePackage("catalystsquad-platform", {
    version: utilities.getVersion(),
    constructProvider: (name: string, type: string, urn: string): pulumi.ProviderResource => {
        if (type !== "pulumi:providers:catalystsquad-platform") {
            throw new Error(`unknown provider type ${type}`);
        }
        return new Provider(name, <any>undefined, { urn });
    },
});
