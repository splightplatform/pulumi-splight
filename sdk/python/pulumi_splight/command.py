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

__all__ = ['CommandArgs', 'Command']

@pulumi.input_type
class CommandArgs:
    def __init__(__self__, *,
                 actions: pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]],
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        The set of arguments for constructing a Command resource.
        :param pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]] actions: command actions
        :param pulumi.Input[str] description: the description of the command to be created
        :param pulumi.Input[str] name: the name of the command to be created
        """
        pulumi.set(__self__, "actions", actions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]]:
        """
        command actions
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        the description of the command to be created
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        the name of the command to be created
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class _CommandState:
    def __init__(__self__, *,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Command resources.
        :param pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]] actions: command actions
        :param pulumi.Input[str] description: the description of the command to be created
        :param pulumi.Input[str] name: the name of the command to be created
        """
        if actions is not None:
            pulumi.set(__self__, "actions", actions)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if name is not None:
            pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def actions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]]]:
        """
        command actions
        """
        return pulumi.get(self, "actions")

    @actions.setter
    def actions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['CommandActionArgs']]]]):
        pulumi.set(self, "actions", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        the description of the command to be created
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        the name of the command to be created
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)


class Command(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CommandActionArgs', 'CommandActionArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_splight as splight

        my_asset = splight.Asset("myAsset",
            description="My Asset Description",
            geometry=json.dumps({
                "type": "GeometryCollection",
                "geometries": [{
                    "type": "Point",
                    "coordinates": [
                        0,
                        0,
                    ],
                }],
            }))
        my_attribute = splight.AssetAttribute("myAttribute",
            type="Number",
            unit="meters",
            asset=my_asset.id)
        my_action = splight.Action("myAction",
            asset={
                "id": my_asset.id,
                "name": my_asset.name,
            },
            setpoints=[{
                "value": json.dumps(1),
                "attribute": {
                    "id": my_attribute.id,
                    "name": my_attribute.name,
                },
            }])
        my_command = splight.Command("myCommand", actions=[{
            "id": my_action.id,
            "name": my_action.name,
            "asset": {
                "id": my_asset.id,
                "name": my_asset.name,
            },
        }])
        ```

        ## Import

        ```sh
        $ pulumi import splight:index/command:Command [options] splight_command.<name> <command_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['CommandActionArgs', 'CommandActionArgsDict']]]] actions: command actions
        :param pulumi.Input[str] description: the description of the command to be created
        :param pulumi.Input[str] name: the name of the command to be created
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: CommandArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ```python
        import pulumi
        import json
        import pulumi_splight as splight

        my_asset = splight.Asset("myAsset",
            description="My Asset Description",
            geometry=json.dumps({
                "type": "GeometryCollection",
                "geometries": [{
                    "type": "Point",
                    "coordinates": [
                        0,
                        0,
                    ],
                }],
            }))
        my_attribute = splight.AssetAttribute("myAttribute",
            type="Number",
            unit="meters",
            asset=my_asset.id)
        my_action = splight.Action("myAction",
            asset={
                "id": my_asset.id,
                "name": my_asset.name,
            },
            setpoints=[{
                "value": json.dumps(1),
                "attribute": {
                    "id": my_attribute.id,
                    "name": my_attribute.name,
                },
            }])
        my_command = splight.Command("myCommand", actions=[{
            "id": my_action.id,
            "name": my_action.name,
            "asset": {
                "id": my_asset.id,
                "name": my_asset.name,
            },
        }])
        ```

        ## Import

        ```sh
        $ pulumi import splight:index/command:Command [options] splight_command.<name> <command_id>
        ```

        :param str resource_name: The name of the resource.
        :param CommandArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(CommandArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 actions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CommandActionArgs', 'CommandActionArgsDict']]]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = CommandArgs.__new__(CommandArgs)

            if actions is None and not opts.urn:
                raise TypeError("Missing required property 'actions'")
            __props__.__dict__["actions"] = actions
            __props__.__dict__["description"] = description
            __props__.__dict__["name"] = name
        super(Command, __self__).__init__(
            'splight:index/command:Command',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            actions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['CommandActionArgs', 'CommandActionArgsDict']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            name: Optional[pulumi.Input[str]] = None) -> 'Command':
        """
        Get an existing Command resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['CommandActionArgs', 'CommandActionArgsDict']]]] actions: command actions
        :param pulumi.Input[str] description: the description of the command to be created
        :param pulumi.Input[str] name: the name of the command to be created
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _CommandState.__new__(_CommandState)

        __props__.__dict__["actions"] = actions
        __props__.__dict__["description"] = description
        __props__.__dict__["name"] = name
        return Command(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def actions(self) -> pulumi.Output[Sequence['outputs.CommandAction']]:
        """
        command actions
        """
        return pulumi.get(self, "actions")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[Optional[str]]:
        """
        the description of the command to be created
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        the name of the command to be created
        """
        return pulumi.get(self, "name")
