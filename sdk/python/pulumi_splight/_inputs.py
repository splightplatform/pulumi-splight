# coding=utf-8
# *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
# *** Do not edit by hand unless you're certain you know what you are doing! ***

import copy
import warnings
import pulumi
import pulumi.runtime
from typing import Any, Mapping, Optional, Sequence, Union, overload
from . import _utilities

__all__ = [
    'AlertAlertItemArgs',
    'AlertThresholdArgs',
    'AssetKindArgs',
    'ComponentInputArgs',
    'ComponentRoutineConfigArgs',
    'ComponentRoutineInputArgs',
    'ComponentRoutineInputValueArgs',
    'ComponentRoutineOutputArgs',
    'ComponentRoutineOutputValueArgs',
    'DashboardChartChartItemArgs',
    'DashboardChartChartItemQueryFilterAssetArgs',
    'DashboardChartChartItemQueryFilterAttributeArgs',
    'DashboardChartThresholdArgs',
    'DashboardChartValueMappingArgs',
    'FunctionFunctionItemArgs',
]

@pulumi.input_type
class AlertAlertItemArgs:
    def __init__(__self__, *,
                 expression_plain: pulumi.Input[str],
                 query_plain: pulumi.Input[str],
                 ref_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: optional id
        """
        pulumi.set(__self__, "expression_plain", expression_plain)
        pulumi.set(__self__, "query_plain", query_plain)
        pulumi.set(__self__, "ref_id", ref_id)
        pulumi.set(__self__, "type", type)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="expressionPlain")
    def expression_plain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression_plain")

    @expression_plain.setter
    def expression_plain(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression_plain", value)

    @property
    @pulumi.getter(name="queryPlain")
    def query_plain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "query_plain")

    @query_plain.setter
    def query_plain(self, value: pulumi.Input[str]):
        pulumi.set(self, "query_plain", value)

    @property
    @pulumi.getter(name="refId")
    def ref_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ref_id")

    @ref_id.setter
    def ref_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ref_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        optional id
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


@pulumi.input_type
class AlertThresholdArgs:
    def __init__(__self__, *,
                 status: pulumi.Input[str],
                 value: pulumi.Input[float],
                 status_text: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] status: [alert|warning|no_alert] status value for the threshold
        :param pulumi.Input[float] value: value to be considered to compare
        :param pulumi.Input[str] status_text: optional custom value to be displayed in the platform.
        """
        pulumi.set(__self__, "status", status)
        pulumi.set(__self__, "value", value)
        if status_text is not None:
            pulumi.set(__self__, "status_text", status_text)

    @property
    @pulumi.getter
    def status(self) -> pulumi.Input[str]:
        """
        [alert|warning|no_alert] status value for the threshold
        """
        return pulumi.get(self, "status")

    @status.setter
    def status(self, value: pulumi.Input[str]):
        pulumi.set(self, "status", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[float]:
        """
        value to be considered to compare
        """
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[float]):
        pulumi.set(self, "value", value)

    @property
    @pulumi.getter(name="statusText")
    def status_text(self) -> Optional[pulumi.Input[str]]:
        """
        optional custom value to be displayed in the platform.
        """
        return pulumi.get(self, "status_text")

    @status_text.setter
    def status_text(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "status_text", value)


@pulumi.input_type
class AssetKindArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] id: kind id
        :param pulumi.Input[str] name: kind name
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        kind id
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        """
        kind name
        """
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class ComponentInputArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 multiple: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if multiple is not None:
            pulumi.set(__self__, "multiple", multiple)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if sensitive is not None:
            pulumi.set(__self__, "sensitive", sensitive)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def multiple(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multiple")

    @multiple.setter
    def multiple(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multiple", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter
    def sensitive(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sensitive")

    @sensitive.setter
    def sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sensitive", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ComponentRoutineConfigArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 multiple: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 value: Optional[pulumi.Input[str]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "type", type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if multiple is not None:
            pulumi.set(__self__, "multiple", multiple)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if sensitive is not None:
            pulumi.set(__self__, "sensitive", sensitive)
        if value is not None:
            pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def multiple(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multiple")

    @multiple.setter
    def multiple(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multiple", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter
    def sensitive(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sensitive")

    @sensitive.setter
    def sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sensitive", value)

    @property
    @pulumi.getter
    def value(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class ComponentRoutineInputArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 multiple: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentRoutineInputValueArgs']]]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value_type", value_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if multiple is not None:
            pulumi.set(__self__, "multiple", multiple)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if sensitive is not None:
            pulumi.set(__self__, "sensitive", sensitive)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value_type")

    @value_type.setter
    def value_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "value_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def multiple(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multiple")

    @multiple.setter
    def multiple(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multiple", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter
    def sensitive(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sensitive")

    @sensitive.setter
    def sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sensitive", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ComponentRoutineInputValueArgs']]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentRoutineInputValueArgs']]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class ComponentRoutineInputValueArgs:
    def __init__(__self__, *,
                 asset: pulumi.Input[str],
                 attribute: pulumi.Input[str]):
        pulumi.set(__self__, "asset", asset)
        pulumi.set(__self__, "attribute", attribute)

    @property
    @pulumi.getter
    def asset(self) -> pulumi.Input[str]:
        return pulumi.get(self, "asset")

    @asset.setter
    def asset(self, value: pulumi.Input[str]):
        pulumi.set(self, "asset", value)

    @property
    @pulumi.getter
    def attribute(self) -> pulumi.Input[str]:
        return pulumi.get(self, "attribute")

    @attribute.setter
    def attribute(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute", value)


@pulumi.input_type
class ComponentRoutineOutputArgs:
    def __init__(__self__, *,
                 name: pulumi.Input[str],
                 value_type: pulumi.Input[str],
                 description: Optional[pulumi.Input[str]] = None,
                 multiple: Optional[pulumi.Input[bool]] = None,
                 required: Optional[pulumi.Input[bool]] = None,
                 sensitive: Optional[pulumi.Input[bool]] = None,
                 type: Optional[pulumi.Input[str]] = None,
                 values: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentRoutineOutputValueArgs']]]] = None):
        pulumi.set(__self__, "name", name)
        pulumi.set(__self__, "value_type", value_type)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if multiple is not None:
            pulumi.set(__self__, "multiple", multiple)
        if required is not None:
            pulumi.set(__self__, "required", required)
        if sensitive is not None:
            pulumi.set(__self__, "sensitive", sensitive)
        if type is not None:
            pulumi.set(__self__, "type", type)
        if values is not None:
            pulumi.set(__self__, "values", values)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)

    @property
    @pulumi.getter(name="valueType")
    def value_type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "value_type")

    @value_type.setter
    def value_type(self, value: pulumi.Input[str]):
        pulumi.set(self, "value_type", value)

    @property
    @pulumi.getter
    def description(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "description")

    @description.setter
    def description(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "description", value)

    @property
    @pulumi.getter
    def multiple(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "multiple")

    @multiple.setter
    def multiple(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "multiple", value)

    @property
    @pulumi.getter
    def required(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "required")

    @required.setter
    def required(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "required", value)

    @property
    @pulumi.getter
    def sensitive(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "sensitive")

    @sensitive.setter
    def sensitive(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "sensitive", value)

    @property
    @pulumi.getter
    def type(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def values(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['ComponentRoutineOutputValueArgs']]]]:
        return pulumi.get(self, "values")

    @values.setter
    def values(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['ComponentRoutineOutputValueArgs']]]]):
        pulumi.set(self, "values", value)


@pulumi.input_type
class ComponentRoutineOutputValueArgs:
    def __init__(__self__, *,
                 asset: pulumi.Input[str],
                 attribute: pulumi.Input[str]):
        pulumi.set(__self__, "asset", asset)
        pulumi.set(__self__, "attribute", attribute)

    @property
    @pulumi.getter
    def asset(self) -> pulumi.Input[str]:
        return pulumi.get(self, "asset")

    @asset.setter
    def asset(self, value: pulumi.Input[str]):
        pulumi.set(self, "asset", value)

    @property
    @pulumi.getter
    def attribute(self) -> pulumi.Input[str]:
        return pulumi.get(self, "attribute")

    @attribute.setter
    def attribute(self, value: pulumi.Input[str]):
        pulumi.set(self, "attribute", value)


@pulumi.input_type
class DashboardChartChartItemArgs:
    def __init__(__self__, *,
                 color: pulumi.Input[str],
                 expression_plain: pulumi.Input[str],
                 query_plain: pulumi.Input[str],
                 ref_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 hidden: Optional[pulumi.Input[bool]] = None,
                 label: Optional[pulumi.Input[str]] = None,
                 query_filter_asset: Optional[pulumi.Input['DashboardChartChartItemQueryFilterAssetArgs']] = None,
                 query_filter_attribute: Optional[pulumi.Input['DashboardChartChartItemQueryFilterAttributeArgs']] = None,
                 query_group_function: Optional[pulumi.Input[str]] = None,
                 query_group_unit: Optional[pulumi.Input[str]] = None,
                 query_limit: Optional[pulumi.Input[int]] = None,
                 query_sort_direction: Optional[pulumi.Input[int]] = None):
        """
        :param pulumi.Input['DashboardChartChartItemQueryFilterAssetArgs'] query_filter_asset: asset filter
        :param pulumi.Input['DashboardChartChartItemQueryFilterAttributeArgs'] query_filter_attribute: Attribute filter
        """
        pulumi.set(__self__, "color", color)
        pulumi.set(__self__, "expression_plain", expression_plain)
        pulumi.set(__self__, "query_plain", query_plain)
        pulumi.set(__self__, "ref_id", ref_id)
        pulumi.set(__self__, "type", type)
        if hidden is not None:
            pulumi.set(__self__, "hidden", hidden)
        if label is not None:
            pulumi.set(__self__, "label", label)
        if query_filter_asset is not None:
            pulumi.set(__self__, "query_filter_asset", query_filter_asset)
        if query_filter_attribute is not None:
            pulumi.set(__self__, "query_filter_attribute", query_filter_attribute)
        if query_group_function is not None:
            pulumi.set(__self__, "query_group_function", query_group_function)
        if query_group_unit is not None:
            pulumi.set(__self__, "query_group_unit", query_group_unit)
        if query_limit is not None:
            pulumi.set(__self__, "query_limit", query_limit)
        if query_sort_direction is not None:
            pulumi.set(__self__, "query_sort_direction", query_sort_direction)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Input[str]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: pulumi.Input[str]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter(name="expressionPlain")
    def expression_plain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression_plain")

    @expression_plain.setter
    def expression_plain(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression_plain", value)

    @property
    @pulumi.getter(name="queryPlain")
    def query_plain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "query_plain")

    @query_plain.setter
    def query_plain(self, value: pulumi.Input[str]):
        pulumi.set(self, "query_plain", value)

    @property
    @pulumi.getter(name="refId")
    def ref_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ref_id")

    @ref_id.setter
    def ref_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ref_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def hidden(self) -> Optional[pulumi.Input[bool]]:
        return pulumi.get(self, "hidden")

    @hidden.setter
    def hidden(self, value: Optional[pulumi.Input[bool]]):
        pulumi.set(self, "hidden", value)

    @property
    @pulumi.getter
    def label(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "label")

    @label.setter
    def label(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "label", value)

    @property
    @pulumi.getter(name="queryFilterAsset")
    def query_filter_asset(self) -> Optional[pulumi.Input['DashboardChartChartItemQueryFilterAssetArgs']]:
        """
        asset filter
        """
        return pulumi.get(self, "query_filter_asset")

    @query_filter_asset.setter
    def query_filter_asset(self, value: Optional[pulumi.Input['DashboardChartChartItemQueryFilterAssetArgs']]):
        pulumi.set(self, "query_filter_asset", value)

    @property
    @pulumi.getter(name="queryFilterAttribute")
    def query_filter_attribute(self) -> Optional[pulumi.Input['DashboardChartChartItemQueryFilterAttributeArgs']]:
        """
        Attribute filter
        """
        return pulumi.get(self, "query_filter_attribute")

    @query_filter_attribute.setter
    def query_filter_attribute(self, value: Optional[pulumi.Input['DashboardChartChartItemQueryFilterAttributeArgs']]):
        pulumi.set(self, "query_filter_attribute", value)

    @property
    @pulumi.getter(name="queryGroupFunction")
    def query_group_function(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "query_group_function")

    @query_group_function.setter
    def query_group_function(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_group_function", value)

    @property
    @pulumi.getter(name="queryGroupUnit")
    def query_group_unit(self) -> Optional[pulumi.Input[str]]:
        return pulumi.get(self, "query_group_unit")

    @query_group_unit.setter
    def query_group_unit(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "query_group_unit", value)

    @property
    @pulumi.getter(name="queryLimit")
    def query_limit(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "query_limit")

    @query_limit.setter
    def query_limit(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "query_limit", value)

    @property
    @pulumi.getter(name="querySortDirection")
    def query_sort_direction(self) -> Optional[pulumi.Input[int]]:
        return pulumi.get(self, "query_sort_direction")

    @query_sort_direction.setter
    def query_sort_direction(self, value: Optional[pulumi.Input[int]]):
        pulumi.set(self, "query_sort_direction", value)


@pulumi.input_type
class DashboardChartChartItemQueryFilterAssetArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DashboardChartChartItemQueryFilterAttributeArgs:
    def __init__(__self__, *,
                 id: pulumi.Input[str],
                 name: pulumi.Input[str]):
        """
        :param pulumi.Input[str] id: The ID of this resource.
        """
        pulumi.set(__self__, "id", id)
        pulumi.set(__self__, "name", name)

    @property
    @pulumi.getter
    def id(self) -> pulumi.Input[str]:
        """
        The ID of this resource.
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: pulumi.Input[str]):
        pulumi.set(self, "id", value)

    @property
    @pulumi.getter
    def name(self) -> pulumi.Input[str]:
        return pulumi.get(self, "name")

    @name.setter
    def name(self, value: pulumi.Input[str]):
        pulumi.set(self, "name", value)


@pulumi.input_type
class DashboardChartThresholdArgs:
    def __init__(__self__, *,
                 color: pulumi.Input[str],
                 display_text: pulumi.Input[str],
                 value: pulumi.Input[float]):
        pulumi.set(__self__, "color", color)
        pulumi.set(__self__, "display_text", display_text)
        pulumi.set(__self__, "value", value)

    @property
    @pulumi.getter
    def color(self) -> pulumi.Input[str]:
        return pulumi.get(self, "color")

    @color.setter
    def color(self, value: pulumi.Input[str]):
        pulumi.set(self, "color", value)

    @property
    @pulumi.getter(name="displayText")
    def display_text(self) -> pulumi.Input[str]:
        return pulumi.get(self, "display_text")

    @display_text.setter
    def display_text(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_text", value)

    @property
    @pulumi.getter
    def value(self) -> pulumi.Input[float]:
        return pulumi.get(self, "value")

    @value.setter
    def value(self, value: pulumi.Input[float]):
        pulumi.set(self, "value", value)


@pulumi.input_type
class DashboardChartValueMappingArgs:
    def __init__(__self__, *,
                 display_text: pulumi.Input[str],
                 match_value: pulumi.Input[str],
                 order: pulumi.Input[int],
                 type: pulumi.Input[str]):
        pulumi.set(__self__, "display_text", display_text)
        pulumi.set(__self__, "match_value", match_value)
        pulumi.set(__self__, "order", order)
        pulumi.set(__self__, "type", type)

    @property
    @pulumi.getter(name="displayText")
    def display_text(self) -> pulumi.Input[str]:
        return pulumi.get(self, "display_text")

    @display_text.setter
    def display_text(self, value: pulumi.Input[str]):
        pulumi.set(self, "display_text", value)

    @property
    @pulumi.getter(name="matchValue")
    def match_value(self) -> pulumi.Input[str]:
        return pulumi.get(self, "match_value")

    @match_value.setter
    def match_value(self, value: pulumi.Input[str]):
        pulumi.set(self, "match_value", value)

    @property
    @pulumi.getter
    def order(self) -> pulumi.Input[int]:
        return pulumi.get(self, "order")

    @order.setter
    def order(self, value: pulumi.Input[int]):
        pulumi.set(self, "order", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)


@pulumi.input_type
class FunctionFunctionItemArgs:
    def __init__(__self__, *,
                 expression_plain: pulumi.Input[str],
                 query_plain: pulumi.Input[str],
                 ref_id: pulumi.Input[str],
                 type: pulumi.Input[str],
                 id: Optional[pulumi.Input[str]] = None):
        """
        :param pulumi.Input[str] id: optional id
        """
        pulumi.set(__self__, "expression_plain", expression_plain)
        pulumi.set(__self__, "query_plain", query_plain)
        pulumi.set(__self__, "ref_id", ref_id)
        pulumi.set(__self__, "type", type)
        if id is not None:
            pulumi.set(__self__, "id", id)

    @property
    @pulumi.getter(name="expressionPlain")
    def expression_plain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "expression_plain")

    @expression_plain.setter
    def expression_plain(self, value: pulumi.Input[str]):
        pulumi.set(self, "expression_plain", value)

    @property
    @pulumi.getter(name="queryPlain")
    def query_plain(self) -> pulumi.Input[str]:
        return pulumi.get(self, "query_plain")

    @query_plain.setter
    def query_plain(self, value: pulumi.Input[str]):
        pulumi.set(self, "query_plain", value)

    @property
    @pulumi.getter(name="refId")
    def ref_id(self) -> pulumi.Input[str]:
        return pulumi.get(self, "ref_id")

    @ref_id.setter
    def ref_id(self, value: pulumi.Input[str]):
        pulumi.set(self, "ref_id", value)

    @property
    @pulumi.getter
    def type(self) -> pulumi.Input[str]:
        return pulumi.get(self, "type")

    @type.setter
    def type(self, value: pulumi.Input[str]):
        pulumi.set(self, "type", value)

    @property
    @pulumi.getter
    def id(self) -> Optional[pulumi.Input[str]]:
        """
        optional id
        """
        return pulumi.get(self, "id")

    @id.setter
    def id(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "id", value)


