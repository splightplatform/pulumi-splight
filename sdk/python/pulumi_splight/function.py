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

__all__ = ['FunctionArgs', 'Function']

@pulumi.input_type
class FunctionArgs:
    def __init__(__self__, *,
                 description: pulumi.Input[str],
                 function_items: pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]],
                 target_asset: pulumi.Input['FunctionTargetAssetArgs'],
                 target_attribute: pulumi.Input['FunctionTargetAttributeArgs'],
                 target_variable: pulumi.Input[str],
                 time_window: pulumi.Input[int],
                 type: pulumi.Input[str],
                 cron_dom: Optional[pulumi.Input[int]] = None,
                 cron_dow: Optional[pulumi.Input[int]] = None,
                 cron_hours: Optional[pulumi.Input[int]] = None,
                 cron_minutes: Optional[pulumi.Input[int]] = None,
                 cron_month: Optional[pulumi.Input[int]] = None,
                 cron_year: Optional[pulumi.Input[int]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_unit: Optional[pulumi.Input[str]] = None,
                 rate_value: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]]] = None):
        """
        The set of arguments for constructing a Function resource.
        :param pulumi.Input[str] description: The description of the resource
        :param pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]] function_items: traces to be used to compute the results
        :param pulumi.Input['FunctionTargetAssetArgs'] target_asset: Asset filter
        :param pulumi.Input['FunctionTargetAttributeArgs'] target_attribute: Attribute filter
        :param pulumi.Input[str] target_variable: variable to be considered to be ingested
        :param pulumi.Input[int] time_window: window to fetch data from. Data out of that window will not be considered for evaluation
        :param pulumi.Input[str] type: [cron|rate] type for the cron
        :param pulumi.Input[int] cron_dom: schedule value for cron
        :param pulumi.Input[int] cron_dow: schedule value for cron
        :param pulumi.Input[int] cron_hours: schedule value for cron
        :param pulumi.Input[int] cron_minutes: schedule value for cron
        :param pulumi.Input[int] cron_month: schedule value for cron
        :param pulumi.Input[int] cron_year: schedule value for cron
        :param pulumi.Input[str] name: The name of the resource
        :param pulumi.Input[str] rate_unit: [day|hour|minute] schedule unit
        :param pulumi.Input[int] rate_value: schedule value
        :param pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]] tags: tags of the resource
        """
        pulumi.set(__self__, "description", description)
        pulumi.set(__self__, "function_items", function_items)
        pulumi.set(__self__, "target_asset", target_asset)
        pulumi.set(__self__, "target_attribute", target_attribute)
        pulumi.set(__self__, "target_variable", target_variable)
        pulumi.set(__self__, "time_window", time_window)
        pulumi.set(__self__, "type", type)
        if cron_dom is not None:
            pulumi.set(__self__, "cron_dom", cron_dom)
        if cron_dow is not None:
            pulumi.set(__self__, "cron_dow", cron_dow)
        if cron_hours is not None:
            pulumi.set(__self__, "cron_hours", cron_hours)
        if cron_minutes is not None:
            pulumi.set(__self__, "cron_minutes", cron_minutes)
        if cron_month is not None:
            pulumi.set(__self__, "cron_month", cron_month)
        if cron_year is not None:
            pulumi.set(__self__, "cron_year", cron_year)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate_unit is not None:
            pulumi.set(__self__, "rate_unit", rate_unit)
        if rate_value is not None:
            pulumi.set(__self__, "rate_value", rate_value)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def description(self) -> pulumi.Input[str]:
        """
        The description of the resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: pulumi.Input[str]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="functionItems")
    def function_items(self) -> pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]]:
        """
        traces to be used to compute the results
        """
        return pulumi.get(self, "function_items")

    @function_items.setter
    def function_items(self, value: pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]]):
        pulumi.set(self, "function_items", value)

    @property
    @pulumi.getter(name="targetAsset")
    def target_asset(self) -> pulumi.Input['FunctionTargetAssetArgs']:
        """
        Asset filter
        """
        return pulumi.get(self, "target_asset")

    @target_asset.setter
    def target_asset(self, value: pulumi.Input['FunctionTargetAssetArgs']):
        pulumi.set(self, "target_asset", value)

    @property
    @pulumi.getter(name="targetAttribute")
    def target_attribute(self) -> pulumi.Input['FunctionTargetAttributeArgs']:
        """
        Attribute filter
        """
        return pulumi.get(self, "target_attribute")

    @target_attribute.setter
    def target_attribute(self, value: pulumi.Input['FunctionTargetAttributeArgs']):
        pulumi.set(self, "target_attribute", value)

    @property
    @pulumi.getter(name="targetVariable")
    def target_variable(self) -> pulumi.Input[str]:
        """
        variable to be considered to be ingested
        """
        return pulumi.get(self, "target_variable")

    @target_variable.setter
    def target_variable(self, value: pulumi.Input[str]):
        pulumi.set(self, "target_variable", value)

    @property
    @pulumi.getter(name="timeWindow")
    def time_window(self) -> pulumi.Input[int]:
        """
        window to fetch data from. Data out of that window will not be considered for evaluation
        """
        return pulumi.get(self, "time_window")

    @time_window.setter
    def time_window(self, value: pulumi.Input[int]):
        pulumi.set(self, "time_window", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        """
        [cron|rate] type for the cron
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter(name="cronDom")
    def cron_dom(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_dom")

    @cron_dom.setter
    def cron_dom(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_dom", value)

    @property
    @pulumi.getter(name="cronDow")
    def cron_dow(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_dow")

    @cron_dow.setter
    def cron_dow(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_dow", value)

    @property
    @pulumi.getter(name="cronHours")
    def cron_hours(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_hours")

    @cron_hours.setter
    def cron_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_hours", value)

    @property
    @pulumi.getter(name="cronMinutes")
    def cron_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_minutes")

    @cron_minutes.setter
    def cron_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_minutes", value)

    @property
    @pulumi.getter(name="cronMonth")
    def cron_month(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_month")

    @cron_month.setter
    def cron_month(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_month", value)

    @property
    @pulumi.getter(name="cronYear")
    def cron_year(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_year")

    @cron_year.setter
    def cron_year(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_year", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rateUnit")
    def rate_unit(self) -> Optional[pulumi.Input[str]]:
        """
        [day|hour|minute] schedule unit
        """
        return pulumi.get(self, "rate_unit")

    @rate_unit.setter
    def rate_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rate_unit", value)

    @property
    @pulumi.getter(name="rateValue")
    def rate_value(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value
        """
        return pulumi.get(self, "rate_value")

    @rate_value.setter
    def rate_value(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate_value", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _FunctionState:
    def __init__(__self__, *,
                 cron_dom: Optional[pulumi.Input[int]] = None,
                 cron_dow: Optional[pulumi.Input[int]] = None,
                 cron_hours: Optional[pulumi.Input[int]] = None,
                 cron_minutes: Optional[pulumi.Input[int]] = None,
                 cron_month: Optional[pulumi.Input[int]] = None,
                 cron_year: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_items: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_unit: Optional[pulumi.Input[str]] = None,
                 rate_value: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]]] = None,
                 target_asset: Optional[pulumi.Input['FunctionTargetAssetArgs']] = None,
                 target_attribute: Optional[pulumi.Input['FunctionTargetAttributeArgs']] = None,
                 target_variable: Optional[pulumi.Input[str]] = None,
                 time_window: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None):
        """
        Input properties used for looking up and filtering Function resources.
        :param pulumi.Input[int] cron_dom: schedule value for cron
        :param pulumi.Input[int] cron_dow: schedule value for cron
        :param pulumi.Input[int] cron_hours: schedule value for cron
        :param pulumi.Input[int] cron_minutes: schedule value for cron
        :param pulumi.Input[int] cron_month: schedule value for cron
        :param pulumi.Input[int] cron_year: schedule value for cron
        :param pulumi.Input[str] description: The description of the resource
        :param pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]] function_items: traces to be used to compute the results
        :param pulumi.Input[str] name: The name of the resource
        :param pulumi.Input[str] rate_unit: [day|hour|minute] schedule unit
        :param pulumi.Input[int] rate_value: schedule value
        :param pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]] tags: tags of the resource
        :param pulumi.Input['FunctionTargetAssetArgs'] target_asset: Asset filter
        :param pulumi.Input['FunctionTargetAttributeArgs'] target_attribute: Attribute filter
        :param pulumi.Input[str] target_variable: variable to be considered to be ingested
        :param pulumi.Input[int] time_window: window to fetch data from. Data out of that window will not be considered for evaluation
        :param pulumi.Input[str] type: [cron|rate] type for the cron
        """
        if cron_dom is not None:
            pulumi.set(__self__, "cron_dom", cron_dom)
        if cron_dow is not None:
            pulumi.set(__self__, "cron_dow", cron_dow)
        if cron_hours is not None:
            pulumi.set(__self__, "cron_hours", cron_hours)
        if cron_minutes is not None:
            pulumi.set(__self__, "cron_minutes", cron_minutes)
        if cron_month is not None:
            pulumi.set(__self__, "cron_month", cron_month)
        if cron_year is not None:
            pulumi.set(__self__, "cron_year", cron_year)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if function_items is not None:
            pulumi.set(__self__, "function_items", function_items)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if rate_unit is not None:
            pulumi.set(__self__, "rate_unit", rate_unit)
        if rate_value is not None:
            pulumi.set(__self__, "rate_value", rate_value)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if target_asset is not None:
            pulumi.set(__self__, "target_asset", target_asset)
        if target_attribute is not None:
            pulumi.set(__self__, "target_attribute", target_attribute)
        if target_variable is not None:
            pulumi.set(__self__, "target_variable", target_variable)
        if time_window is not None:
            pulumi.set(__self__, "time_window", time_window)
        if type is not None:
            pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="cronDom")
    def cron_dom(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_dom")

    @cron_dom.setter
    def cron_dom(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_dom", value)

    @property
    @pulumi.getter(name="cronDow")
    def cron_dow(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_dow")

    @cron_dow.setter
    def cron_dow(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_dow", value)

    @property
    @pulumi.getter(name="cronHours")
    def cron_hours(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_hours")

    @cron_hours.setter
    def cron_hours(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_hours", value)

    @property
    @pulumi.getter(name="cronMinutes")
    def cron_minutes(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_minutes")

    @cron_minutes.setter
    def cron_minutes(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_minutes", value)

    @property
    @pulumi.getter(name="cronMonth")
    def cron_month(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_month")

    @cron_month.setter
    def cron_month(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_month", value)

    @property
    @pulumi.getter(name="cronYear")
    def cron_year(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_year")

    @cron_year.setter
    def cron_year(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "cron_year", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        """
        The description of the resource
        """
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter(name="functionItems")
    def function_items(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]]]:
        """
        traces to be used to compute the results
        """
        return pulumi.get(self, "function_items")

    @function_items.setter
    def function_items(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionFunctionItemArgs']]]]):
        pulumi.set(self, "function_items", value)

    @property
    @pulumi.getter
    def name(self) -> Optional[pulumi.Input[str]]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="rateUnit")
    def rate_unit(self) -> Optional[pulumi.Input[str]]:
        """
        [day|hour|minute] schedule unit
        """
        return pulumi.get(self, "rate_unit")

    @rate_unit.setter
    def rate_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "rate_unit", value)

    @property
    @pulumi.getter(name="rateValue")
    def rate_value(self) -> Optional[pulumi.Input[int]]:
        """
        schedule value
        """
        return pulumi.get(self, "rate_value")

    @rate_value.setter
    def rate_value(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "rate_value", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['FunctionTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter(name="targetAsset")
    def target_asset(self) -> Optional[pulumi.Input['FunctionTargetAssetArgs']]:
        """
        Asset filter
        """
        return pulumi.get(self, "target_asset")

    @target_asset.setter
    def target_asset(self, value: Optional[pulumi.Input['FunctionTargetAssetArgs']]):
        pulumi.set(self, "target_asset", value)

    @property
    @pulumi.getter(name="targetAttribute")
    def target_attribute(self) -> Optional[pulumi.Input['FunctionTargetAttributeArgs']]:
        """
        Attribute filter
        """
        return pulumi.get(self, "target_attribute")

    @target_attribute.setter
    def target_attribute(self, value: Optional[pulumi.Input['FunctionTargetAttributeArgs']]):
        pulumi.set(self, "target_attribute", value)

    @property
    @pulumi.getter(name="targetVariable")
    def target_variable(self) -> Optional[pulumi.Input[str]]:
        """
        variable to be considered to be ingested
        """
        return pulumi.get(self, "target_variable")

    @target_variable.setter
    def target_variable(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "target_variable", value)

    @property
    @pulumi.getter(name="timeWindow")
    def time_window(self) -> Optional[pulumi.Input[int]]:
        """
        window to fetch data from. Data out of that window will not be considered for evaluation
        """
        return pulumi.get(self, "time_window")

    @time_window.setter
    def time_window(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "time_window", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        """
        [cron|rate] type for the cron
        """
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)


class Function(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_dom: Optional[pulumi.Input[int]] = None,
                 cron_dow: Optional[pulumi.Input[int]] = None,
                 cron_hours: Optional[pulumi.Input[int]] = None,
                 cron_minutes: Optional[pulumi.Input[int]] = None,
                 cron_month: Optional[pulumi.Input[int]] = None,
                 cron_year: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_items: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FunctionFunctionItemArgs', 'FunctionFunctionItemArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_unit: Optional[pulumi.Input[str]] = None,
                 rate_value: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FunctionTagArgs', 'FunctionTagArgsDict']]]]] = None,
                 target_asset: Optional[pulumi.Input[Union['FunctionTargetAssetArgs', 'FunctionTargetAssetArgsDict']]] = None,
                 target_attribute: Optional[pulumi.Input[Union['FunctionTargetAttributeArgs', 'FunctionTargetAttributeArgsDict']]] = None,
                 target_variable: Optional[pulumi.Input[str]] = None,
                 time_window: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/function:Function [options] splight_function.<name> <function_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cron_dom: schedule value for cron
        :param pulumi.Input[int] cron_dow: schedule value for cron
        :param pulumi.Input[int] cron_hours: schedule value for cron
        :param pulumi.Input[int] cron_minutes: schedule value for cron
        :param pulumi.Input[int] cron_month: schedule value for cron
        :param pulumi.Input[int] cron_year: schedule value for cron
        :param pulumi.Input[str] description: The description of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['FunctionFunctionItemArgs', 'FunctionFunctionItemArgsDict']]]] function_items: traces to be used to compute the results
        :param pulumi.Input[str] name: The name of the resource
        :param pulumi.Input[str] rate_unit: [day|hour|minute] schedule unit
        :param pulumi.Input[int] rate_value: schedule value
        :param pulumi.Input[Sequence[pulumi.Input[Union['FunctionTagArgs', 'FunctionTagArgsDict']]]] tags: tags of the resource
        :param pulumi.Input[Union['FunctionTargetAssetArgs', 'FunctionTargetAssetArgsDict']] target_asset: Asset filter
        :param pulumi.Input[Union['FunctionTargetAttributeArgs', 'FunctionTargetAttributeArgsDict']] target_attribute: Attribute filter
        :param pulumi.Input[str] target_variable: variable to be considered to be ingested
        :param pulumi.Input[int] time_window: window to fetch data from. Data out of that window will not be considered for evaluation
        :param pulumi.Input[str] type: [cron|rate] type for the cron
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: FunctionArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/function:Function [options] splight_function.<name> <function_id>
        ```

        :param str resource_name: The name of the resource.
        :param FunctionArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(FunctionArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 cron_dom: Optional[pulumi.Input[int]] = None,
                 cron_dow: Optional[pulumi.Input[int]] = None,
                 cron_hours: Optional[pulumi.Input[int]] = None,
                 cron_minutes: Optional[pulumi.Input[int]] = None,
                 cron_month: Optional[pulumi.Input[int]] = None,
                 cron_year: Optional[pulumi.Input[int]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 function_items: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FunctionFunctionItemArgs', 'FunctionFunctionItemArgsDict']]]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 rate_unit: Optional[pulumi.Input[str]] = None,
                 rate_value: Optional[pulumi.Input[int]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FunctionTagArgs', 'FunctionTagArgsDict']]]]] = None,
                 target_asset: Optional[pulumi.Input[Union['FunctionTargetAssetArgs', 'FunctionTargetAssetArgsDict']]] = None,
                 target_attribute: Optional[pulumi.Input[Union['FunctionTargetAttributeArgs', 'FunctionTargetAttributeArgsDict']]] = None,
                 target_variable: Optional[pulumi.Input[str]] = None,
                 time_window: Optional[pulumi.Input[int]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = FunctionArgs.__new__(FunctionArgs)

            __props__.__dict__["cron_dom"] = cron_dom
            __props__.__dict__["cron_dow"] = cron_dow
            __props__.__dict__["cron_hours"] = cron_hours
            __props__.__dict__["cron_minutes"] = cron_minutes
            __props__.__dict__["cron_month"] = cron_month
            __props__.__dict__["cron_year"] = cron_year
            if description is None and not opts.urn:
                raise TypeError("Missing required property 'description'")
            __props__.__dict__["description"] = description
            if function_items is None and not opts.urn:
                raise TypeError("Missing required property 'function_items'")
            __props__.__dict__["function_items"] = function_items
            __props__.__dict__["name"] = name
            __props__.__dict__["rate_unit"] = rate_unit
            __props__.__dict__["rate_value"] = rate_value
            __props__.__dict__["tags"] = tags
            if target_asset is None and not opts.urn:
                raise TypeError("Missing required property 'target_asset'")
            __props__.__dict__["target_asset"] = target_asset
            if target_attribute is None and not opts.urn:
                raise TypeError("Missing required property 'target_attribute'")
            __props__.__dict__["target_attribute"] = target_attribute
            if target_variable is None and not opts.urn:
                raise TypeError("Missing required property 'target_variable'")
            __props__.__dict__["target_variable"] = target_variable
            if time_window is None and not opts.urn:
                raise TypeError("Missing required property 'time_window'")
            __props__.__dict__["time_window"] = time_window
            if type is None and not opts.urn:
                raise TypeError("Missing required property 'type'")
            __props__.__dict__["type"] = type
        super(Function, __self__).__init__(
            'splight:index/function:Function',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            cron_dom: Optional[pulumi.Input[int]] = None,
            cron_dow: Optional[pulumi.Input[int]] = None,
            cron_hours: Optional[pulumi.Input[int]] = None,
            cron_minutes: Optional[pulumi.Input[int]] = None,
            cron_month: Optional[pulumi.Input[int]] = None,
            cron_year: Optional[pulumi.Input[int]] = None,
            description: Optional[pulumi.Input[str]] = None,
            function_items: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FunctionFunctionItemArgs', 'FunctionFunctionItemArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            rate_unit: Optional[pulumi.Input[str]] = None,
            rate_value: Optional[pulumi.Input[int]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['FunctionTagArgs', 'FunctionTagArgsDict']]]]] = None,
            target_asset: Optional[pulumi.Input[Union['FunctionTargetAssetArgs', 'FunctionTargetAssetArgsDict']]] = None,
            target_attribute: Optional[pulumi.Input[Union['FunctionTargetAttributeArgs', 'FunctionTargetAttributeArgsDict']]] = None,
            target_variable: Optional[pulumi.Input[str]] = None,
            time_window: Optional[pulumi.Input[int]] = None,
            type: Optional[pulumi.Input[str]] = None) -> 'Function':
        """
        Get an existing Function resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[int] cron_dom: schedule value for cron
        :param pulumi.Input[int] cron_dow: schedule value for cron
        :param pulumi.Input[int] cron_hours: schedule value for cron
        :param pulumi.Input[int] cron_minutes: schedule value for cron
        :param pulumi.Input[int] cron_month: schedule value for cron
        :param pulumi.Input[int] cron_year: schedule value for cron
        :param pulumi.Input[str] description: The description of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['FunctionFunctionItemArgs', 'FunctionFunctionItemArgsDict']]]] function_items: traces to be used to compute the results
        :param pulumi.Input[str] name: The name of the resource
        :param pulumi.Input[str] rate_unit: [day|hour|minute] schedule unit
        :param pulumi.Input[int] rate_value: schedule value
        :param pulumi.Input[Sequence[pulumi.Input[Union['FunctionTagArgs', 'FunctionTagArgsDict']]]] tags: tags of the resource
        :param pulumi.Input[Union['FunctionTargetAssetArgs', 'FunctionTargetAssetArgsDict']] target_asset: Asset filter
        :param pulumi.Input[Union['FunctionTargetAttributeArgs', 'FunctionTargetAttributeArgsDict']] target_attribute: Attribute filter
        :param pulumi.Input[str] target_variable: variable to be considered to be ingested
        :param pulumi.Input[int] time_window: window to fetch data from. Data out of that window will not be considered for evaluation
        :param pulumi.Input[str] type: [cron|rate] type for the cron
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _FunctionState.__new__(_FunctionState)

        __props__.__dict__["cron_dom"] = cron_dom
        __props__.__dict__["cron_dow"] = cron_dow
        __props__.__dict__["cron_hours"] = cron_hours
        __props__.__dict__["cron_minutes"] = cron_minutes
        __props__.__dict__["cron_month"] = cron_month
        __props__.__dict__["cron_year"] = cron_year
        __props__.__dict__["description"] = description
        __props__.__dict__["function_items"] = function_items
        __props__.__dict__["name"] = name
        __props__.__dict__["rate_unit"] = rate_unit
        __props__.__dict__["rate_value"] = rate_value
        __props__.__dict__["tags"] = tags
        __props__.__dict__["target_asset"] = target_asset
        __props__.__dict__["target_attribute"] = target_attribute
        __props__.__dict__["target_variable"] = target_variable
        __props__.__dict__["time_window"] = time_window
        __props__.__dict__["type"] = type
        return Function(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="cronDom")
    def cron_dom(self) -> pulumi.Output[int]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_dom")

    @property
    @pulumi.getter(name="cronDow")
    def cron_dow(self) -> pulumi.Output[int]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_dow")

    @property
    @pulumi.getter(name="cronHours")
    def cron_hours(self) -> pulumi.Output[int]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_hours")

    @property
    @pulumi.getter(name="cronMinutes")
    def cron_minutes(self) -> pulumi.Output[int]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_minutes")

    @property
    @pulumi.getter(name="cronMonth")
    def cron_month(self) -> pulumi.Output[int]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_month")

    @property
    @pulumi.getter(name="cronYear")
    def cron_year(self) -> pulumi.Output[int]:
        """
        schedule value for cron
        """
        return pulumi.get(self, "cron_year")

    @property
    @pulumi.getter
    def description(self) -> pulumi.Output[str]:
        """
        The description of the resource
        """
        return pulumi.get(self, "description")

    @property
    @pulumi.getter(name="functionItems")
    def function_items(self) -> pulumi.Output[Sequence['outputs.FunctionFunctionItem']]:
        """
        traces to be used to compute the results
        """
        return pulumi.get(self, "function_items")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        The name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="rateUnit")
    def rate_unit(self) -> pulumi.Output[str]:
        """
        [day|hour|minute] schedule unit
        """
        return pulumi.get(self, "rate_unit")

    @property
    @pulumi.getter(name="rateValue")
    def rate_value(self) -> pulumi.Output[int]:
        """
        schedule value
        """
        return pulumi.get(self, "rate_value")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.FunctionTag']]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter(name="targetAsset")
    def target_asset(self) -> pulumi.Output['outputs.FunctionTargetAsset']:
        """
        Asset filter
        """
        return pulumi.get(self, "target_asset")

    @property
    @pulumi.getter(name="targetAttribute")
    def target_attribute(self) -> pulumi.Output['outputs.FunctionTargetAttribute']:
        """
        Attribute filter
        """
        return pulumi.get(self, "target_attribute")

    @property
    @pulumi.getter(name="targetVariable")
    def target_variable(self) -> pulumi.Output[str]:
        """
        variable to be considered to be ingested
        """
        return pulumi.get(self, "target_variable")

    @property
    @pulumi.getter(name="timeWindow")
    def time_window(self) -> pulumi.Output[int]:
        """
        window to fetch data from. Data out of that window will not be considered for evaluation
        """
        return pulumi.get(self, "time_window")

    @property
    @pulumi.getter
    def type(self) -> pulumi.Output[str]:
        """
        [cron|rate] type for the cron
        """
        return pulumi.get(self, "type")

