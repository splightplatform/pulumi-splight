# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import sys
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
if sys.version_info >= (3, 11):
    from typing import NotRequired, TypedDict, TypeAlias
else:
    from typing_extensions import NotRequired, TypedDict, TypeAlias
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AssetArgs', 'Asset']

@pulumi.input_type
class AssetArgs:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input['AssetKindArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Asset resource.
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: GeoJSON GeomtryCollection
        :param pulumi.Input['AssetKindArgs'] kind: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]] tags: tags of the resource
        :param pulumi.Input[str] timezone: timezone of the resource (overriden by the location if set)
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        description of the resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def geometry(self) -> Optional[pulumi.Input[str]]:
        """
        GeoJSON GeomtryCollection
        """
        return pulumi.get(self, "geometry")

    @geometry.setter
    def geometry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geometry", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input['AssetKindArgs']]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input['AssetKindArgs']]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        name of the resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        timezone of the resource (overriden by the location if set)
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


@pulumi.input_type
class _AssetState:
    def __init__(__self__, *,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input['AssetKindArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Asset resources.
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: GeoJSON GeomtryCollection
        :param pulumi.Input['AssetKindArgs'] kind: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]] tags: tags of the resource
        :param pulumi.Input[str] timezone: timezone of the resource (overriden by the location if set)
        """
        if description is not None:
            pulumi.set(__self__, "description", description)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if kind is not None:
            pulumi.set(__self__, "kind", kind)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if timezone is not None:
            pulumi.set(__self__, "timezone", timezone)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        description of the resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def geometry(self) -> Optional[pulumi.Input[str]]:
        """
        GeoJSON GeomtryCollection
        """
        return pulumi.get(self, "geometry")

    @geometry.setter
    def geometry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geometry", value)

    @property
    @pulumi.getter
    def kind(self) -> Optional[pulumi.Input['AssetKindArgs']]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kind")

    @kind.setter
    def kind(self, value: Optional[pulumi.Input['AssetKindArgs']]):
        pulumi.set(self, "kind", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        name of the resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['AssetTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def timezone(self) -> Optional[pulumi.Input[str]]:
        """
        timezone of the resource (overriden by the location if set)
        """
        return pulumi.get(self, "timezone")

    @timezone.setter
    def timezone(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "timezone", value)


class Asset(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[Union['AssetKindArgs', 'AssetKindArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AssetTagArgs', 'AssetTagArgsDict']]]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/asset:Asset [options] splight_asset.<name> <asset_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: GeoJSON GeomtryCollection
        :param pulumi.Input[Union['AssetKindArgs', 'AssetKindArgsDict']] kind: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['AssetTagArgs', 'AssetTagArgsDict']]]] tags: tags of the resource
        :param pulumi.Input[str] timezone: timezone of the resource (overriden by the location if set)
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: Optional[AssetArgs] = None,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/asset:Asset [options] splight_asset.<name> <asset_id>
        ```

        :param str resource_name: The name of the resource.
        :param AssetArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssetArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 kind: Optional[pulumi.Input[Union['AssetKindArgs', 'AssetKindArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AssetTagArgs', 'AssetTagArgsDict']]]]] = None,
                 timezone: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AssetArgs.__new__(AssetArgs)

            __props__.__dict__["description"] = description
            __props__.__dict__["geometry"] = geometry
            __props__.__dict__["kind"] = kind
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["timezone"] = timezone
        super(Asset, __self__).__init__(
            'splight:index/asset:Asset',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            description: Optional[pulumi.Input[str]] = None,
            geometry: Optional[pulumi.Input[str]] = None,
            kind: Optional[pulumi.Input[Union['AssetKindArgs', 'AssetKindArgsDict']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['AssetTagArgs', 'AssetTagArgsDict']]]]] = None,
            timezone: Optional[pulumi.Input[str]] = None) -> 'Asset':
        """
        Get an existing Asset resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: GeoJSON GeomtryCollection
        :param pulumi.Input[Union['AssetKindArgs', 'AssetKindArgsDict']] kind: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['AssetTagArgs', 'AssetTagArgsDict']]]] tags: tags of the resource
        :param pulumi.Input[str] timezone: timezone of the resource (overriden by the location if set)
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AssetState.__new__(_AssetState)

        __props__.__dict__["description"] = description
        __props__.__dict__["geometry"] = geometry
        __props__.__dict__["kind"] = kind
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["timezone"] = timezone
        return Asset(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        description of the resource
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def geometry(self) -> pulumi.Output[Optional[str]]:
        """
        GeoJSON GeomtryCollection
        """
        return pulumi.get(self, "geometry")

    @property
    @pulumi.getter
    def kind(self) -> pulumi.Output[Optional['outputs.AssetKind']]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kind")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.AssetTag']]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def timezone(self) -> pulumi.Output[str]:
        """
        timezone of the resource (overriden by the location if set)
        """
        return pulumi.get(self, "timezone")

