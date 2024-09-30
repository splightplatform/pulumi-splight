# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities
from . import outputs
from ._inputs import *

__all__ = ['AssetRelationArgs', 'AssetRelation']

@pulumi.input_type
class AssetRelationArgs:
    def __init__(__self__, *,
                 asset: pulumi.Input['AssetRelationAssetArgs'],
                 related_asset: pulumi.Input['AssetRelationRelatedAssetArgs'],
                 related_asset_kind: pulumi.Input['AssetRelationRelatedAssetKindArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a AssetRelation resource.
        :param pulumi.Input['AssetRelationAssetArgs'] asset: asset where the relation origins
        :param pulumi.Input['AssetRelationRelatedAssetArgs'] related_asset: target asset of the relation
        :param pulumi.Input['AssetRelationRelatedAssetKindArgs'] related_asset_kind: kind of the target relation asset
        :param pulumi.Input[str] description: relation description
        :param pulumi.Input[str] name: relation name
        """
        pulumi.set(__self__, "asset", asset)
        pulumi.set(__self__, "related_asset", related_asset)
        pulumi.set(__self__, "related_asset_kind", related_asset_kind)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def asset(self) -> pulumi.Input['AssetRelationAssetArgs']:
        """
        asset where the relation origins
        """
        return pulumi.get(self, "asset")

    @asset.setter
    def asset(self, value: pulumi.Input['AssetRelationAssetArgs']):
        pulumi.set(self, "asset", value)

    @property
    @pulumi.getter(name="relatedAsset")
    def related_asset(self) -> pulumi.Input['AssetRelationRelatedAssetArgs']:
        """
        target asset of the relation
        """
        return pulumi.get(self, "related_asset")

    @related_asset.setter
    def related_asset(self, value: pulumi.Input['AssetRelationRelatedAssetArgs']):
        pulumi.set(self, "related_asset", value)

    @property
    @pulumi.getter(name="relatedAssetKind")
    def related_asset_kind(self) -> pulumi.Input['AssetRelationRelatedAssetKindArgs']:
        """
        kind of the target relation asset
        """
        return pulumi.get(self, "related_asset_kind")

    @related_asset_kind.setter
    def related_asset_kind(self, value: pulumi.Input['AssetRelationRelatedAssetKindArgs']):
        pulumi.set(self, "related_asset_kind", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        relation description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        relation name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _AssetRelationState:
    def __init__(__self__, *,
                 asset: Optional[pulumi.Input['AssetRelationAssetArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 related_asset: Optional[pulumi.Input['AssetRelationRelatedAssetArgs']] = None,
                 related_asset_kind: Optional[pulumi.Input['AssetRelationRelatedAssetKindArgs']] = None):
        """
        Input properties used for looking up and filtering AssetRelation resources.
        :param pulumi.Input['AssetRelationAssetArgs'] asset: asset where the relation origins
        :param pulumi.Input[str] description: relation description
        :param pulumi.Input[str] name: relation name
        :param pulumi.Input['AssetRelationRelatedAssetArgs'] related_asset: target asset of the relation
        :param pulumi.Input['AssetRelationRelatedAssetKindArgs'] related_asset_kind: kind of the target relation asset
        """
        if asset is not None:
            pulumi.set(__self__, "asset", asset)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if related_asset is not None:
            pulumi.set(__self__, "related_asset", related_asset)
        if related_asset_kind is not None:
            pulumi.set(__self__, "related_asset_kind", related_asset_kind)

    @property
    @pulumi.getter
    def asset(self) -> Optional[pulumi.Input['AssetRelationAssetArgs']]:
        """
        asset where the relation origins
        """
        return pulumi.get(self, "asset")

    @asset.setter
    def asset(self, value: Optional[pulumi.Input['AssetRelationAssetArgs']]):
        pulumi.set(self, "asset", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        relation description
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        relation name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="relatedAsset")
    def related_asset(self) -> Optional[pulumi.Input['AssetRelationRelatedAssetArgs']]:
        """
        target asset of the relation
        """
        return pulumi.get(self, "related_asset")

    @related_asset.setter
    def related_asset(self, value: Optional[pulumi.Input['AssetRelationRelatedAssetArgs']]):
        pulumi.set(self, "related_asset", value)

    @property
    @pulumi.getter(name="relatedAssetKind")
    def related_asset_kind(self) -> Optional[pulumi.Input['AssetRelationRelatedAssetKindArgs']]:
        """
        kind of the target relation asset
        """
        return pulumi.get(self, "related_asset_kind")

    @related_asset_kind.setter
    def related_asset_kind(self, value: Optional[pulumi.Input['AssetRelationRelatedAssetKindArgs']]):
        pulumi.set(self, "related_asset_kind", value)


class AssetRelation(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset: Optional[pulumi.Input[Union['AssetRelationAssetArgs', 'AssetRelationAssetArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 related_asset: Optional[pulumi.Input[Union['AssetRelationRelatedAssetArgs', 'AssetRelationRelatedAssetArgsDict']]] = None,
                 related_asset_kind: Optional[pulumi.Input[Union['AssetRelationRelatedAssetKindArgs', 'AssetRelationRelatedAssetKindArgsDict']]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/assetRelation:AssetRelation [options] splight_relation.<name> <relation_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['AssetRelationAssetArgs', 'AssetRelationAssetArgsDict']] asset: asset where the relation origins
        :param pulumi.Input[str] description: relation description
        :param pulumi.Input[str] name: relation name
        :param pulumi.Input[Union['AssetRelationRelatedAssetArgs', 'AssetRelationRelatedAssetArgsDict']] related_asset: target asset of the relation
        :param pulumi.Input[Union['AssetRelationRelatedAssetKindArgs', 'AssetRelationRelatedAssetKindArgsDict']] related_asset_kind: kind of the target relation asset
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: AssetRelationArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/assetRelation:AssetRelation [options] splight_relation.<name> <relation_id>
        ```

        :param str resource_name: The name of the resource.
        :param AssetRelationArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(AssetRelationArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 asset: Optional[pulumi.Input[Union['AssetRelationAssetArgs', 'AssetRelationAssetArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 related_asset: Optional[pulumi.Input[Union['AssetRelationRelatedAssetArgs', 'AssetRelationRelatedAssetArgsDict']]] = None,
                 related_asset_kind: Optional[pulumi.Input[Union['AssetRelationRelatedAssetKindArgs', 'AssetRelationRelatedAssetKindArgsDict']]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = AssetRelationArgs.__new__(AssetRelationArgs)

            if asset is None and not opts.urn:
                raise TypeError("Missing required property 'asset'")
            __props__.__dict__["asset"] = asset
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
            if related_asset is None and not opts.urn:
                raise TypeError("Missing required property 'related_asset'")
            __props__.__dict__["related_asset"] = related_asset
            if related_asset_kind is None and not opts.urn:
                raise TypeError("Missing required property 'related_asset_kind'")
            __props__.__dict__["related_asset_kind"] = related_asset_kind
        super(AssetRelation, __self__).__init__(
            'splight:index/assetRelation:AssetRelation',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            asset: Optional[pulumi.Input[Union['AssetRelationAssetArgs', 'AssetRelationAssetArgsDict']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None,
            related_asset: Optional[pulumi.Input[Union['AssetRelationRelatedAssetArgs', 'AssetRelationRelatedAssetArgsDict']]] = None,
            related_asset_kind: Optional[pulumi.Input[Union['AssetRelationRelatedAssetKindArgs', 'AssetRelationRelatedAssetKindArgsDict']]] = None) -> 'AssetRelation':
        """
        Get an existing AssetRelation resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['AssetRelationAssetArgs', 'AssetRelationAssetArgsDict']] asset: asset where the relation origins
        :param pulumi.Input[str] description: relation description
        :param pulumi.Input[str] name: relation name
        :param pulumi.Input[Union['AssetRelationRelatedAssetArgs', 'AssetRelationRelatedAssetArgsDict']] related_asset: target asset of the relation
        :param pulumi.Input[Union['AssetRelationRelatedAssetKindArgs', 'AssetRelationRelatedAssetKindArgsDict']] related_asset_kind: kind of the target relation asset
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _AssetRelationState.__new__(_AssetRelationState)

        __props__.__dict__["asset"] = asset
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        __props__.__dict__["related_asset"] = related_asset
        __props__.__dict__["related_asset_kind"] = related_asset_kind
        return AssetRelation(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def asset(self) -> pulumi.Output['outputs.AssetRelationAsset']:
        """
        asset where the relation origins
        """
        return pulumi.get(self, "asset")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        relation description
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        relation name
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="relatedAsset")
    def related_asset(self) -> pulumi.Output['outputs.AssetRelationRelatedAsset']:
        """
        target asset of the relation
        """
        return pulumi.get(self, "related_asset")

    @property
    @pulumi.getter(name="relatedAssetKind")
    def related_asset_kind(self) -> pulumi.Output['outputs.AssetRelationRelatedAssetKind']:
        """
        kind of the target relation asset
        """
        return pulumi.get(self, "related_asset_kind")
