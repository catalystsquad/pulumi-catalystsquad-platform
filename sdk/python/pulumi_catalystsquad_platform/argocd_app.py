# coding=utf-8
# *** WARNING: this file was generated by Pulumi SDK Generator. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from ._inputs import *

__all__ = ['ArgocdAppArgs', 'ArgocdApp']

@pulumi.input_type
class ArgocdAppArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 api_version: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input['ArgocdApplicationArgs']] = None):
        """
        The set of arguments for constructing a ArgocdApp resource.
        :param pulumi.Input[str] name: Required, name of the Argocd Application
        :param pulumi.Input[str] api_version: Optional, apiVersion of the Argocd Application. Default: v1alpha1
        :param pulumi.Input[str] namespace: Optional, namespace to deploy Argocd Application to. Should be the namespace where the argocd server runs. Default: "argo-cd"
        :param pulumi.Input['ArgocdApplicationArgs'] spec: Required, spec of the Argocd Application
        """
        pulumi.set(__self__, "name", name)
        if api_version is not None:
            pulumi.set(__self__, "api_version", api_version)
        if namespace is not None:
            pulumi.set(__self__, "namespace", namespace)
        if spec is not None:
            pulumi.set(__self__, "spec", spec)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        Required, name of the Argocd Application
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="apiVersion")
    def api_version(self) -> Optional[pulumi.Input[str]]:
        """
        Optional, apiVersion of the Argocd Application. Default: v1alpha1
        """
        return pulumi.get(self, "api_version")

    @api_version.setter
    def api_version(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "api_version", value)

    @property
    @pulumi.getter
    def namespace(self) -> Optional[pulumi.Input[str]]:
        """
        Optional, namespace to deploy Argocd Application to. Should be the namespace where the argocd server runs. Default: "argo-cd"
        """
        return pulumi.get(self, "namespace")

    @namespace.setter
    def namespace(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "namespace", value)

    @property
    @pulumi.getter
    def spec(self) -> Optional[pulumi.Input['ArgocdApplicationArgs']]:
        """
        Required, spec of the Argocd Application
        """
        return pulumi.get(self, "spec")

    @spec.setter
    def spec(self, value: Optional[pulumi.Input['ArgocdApplicationArgs']]):
        pulumi.set(self, "spec", value)


class ArgocdApp(pulumi.ComponentResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['ArgocdApplicationArgs']]] = None,
                 __props__=None):
        """
        Create a ArgocdApp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] api_version: Optional, apiVersion of the Argocd Application. Default: v1alpha1
        :param pulumi.Input[str] name: Required, name of the Argocd Application
        :param pulumi.Input[str] namespace: Optional, namespace to deploy Argocd Application to. Should be the namespace where the argocd server runs. Default: "argo-cd"
        :param pulumi.Input[pulumi.InputType['ArgocdApplicationArgs']] spec: Required, spec of the Argocd Application
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: ArgocdAppArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        Create a ArgocdApp resource with the given unique name, props, and options.
        :param str resource_name: The name of the resource.
        :param ArgocdAppArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(ArgocdAppArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 api_version: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 namespace: Optional[pulumi.Input[str]] = None,
                 spec: Optional[pulumi.Input[pulumi.InputType['ArgocdApplicationArgs']]] = None,
                 __props__=None):
        if opts is None:
            opts = pulumi.ResourceOptions()
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.version is None:
            opts.version = _utilities.get_version()
        if opts.plugin_download_url is None:
            opts.plugin_download_url = _utilities.get_plugin_download_url()
        if opts.id is not None:
            raise ValueError('ComponentResource classes do not support opts.id')
        else:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = ArgocdAppArgs.__new__(ArgocdAppArgs)

            __props__.__dict__["api_version"] = api_version
            if name is None and not opts.urn:
                raise TypeError("Missing required property 'name'")
            __props__.__dict__["name"] = name
            __props__.__dict__["namespace"] = namespace
            __props__.__dict__["spec"] = spec
        super(ArgocdApp, __self__).__init__(
            'catalystsquad-platform:index:ArgocdApp',
            resource_name,
            __props__,
            opts,
            remote=True)

