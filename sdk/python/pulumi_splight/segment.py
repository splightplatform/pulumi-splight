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

__all__ = ['SegmentArgs', 'Segment']

@pulumi.input_type
class SegmentArgs:
    def __init__(__self__, *,
                 altitude: pulumi.Input['SegmentAltitudeArgs'],
                 azimuth: pulumi.Input['SegmentAzimuthArgs'],
                 cumulative_distance: pulumi.Input['SegmentCumulativeDistanceArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]]] = None):
        """
        The set of arguments for constructing a Segment resource.
        :param pulumi.Input['SegmentAltitudeArgs'] altitude: attribute of the resource
        :param pulumi.Input['SegmentAzimuthArgs'] azimuth: attribute of the resource
        :param pulumi.Input['SegmentCumulativeDistanceArgs'] cumulative_distance: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]] tags: tags of the resource
        """
        pulumi.set(__self__, "altitude", altitude)
        pulumi.set(__self__, "azimuth", azimuth)
        pulumi.set(__self__, "cumulative_distance", cumulative_distance)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter
    def altitude(self) -> pulumi.Input['SegmentAltitudeArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "altitude")

    @altitude.setter
    def altitude(self, value: pulumi.Input['SegmentAltitudeArgs']):
        pulumi.set(self, "altitude", value)

    @property
    @pulumi.getter
    def azimuth(self) -> pulumi.Input['SegmentAzimuthArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "azimuth")

    @azimuth.setter
    def azimuth(self, value: pulumi.Input['SegmentAzimuthArgs']):
        pulumi.set(self, "azimuth", value)

    @property
    @pulumi.getter(name="cumulativeDistance")
    def cumulative_distance(self) -> pulumi.Input['SegmentCumulativeDistanceArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "cumulative_distance")

    @cumulative_distance.setter
    def cumulative_distance(self, value: pulumi.Input['SegmentCumulativeDistanceArgs']):
        pulumi.set(self, "cumulative_distance", value)

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
        geo position and shape of the resource
        """
        return pulumi.get(self, "geometry")

    @geometry.setter
    def geometry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geometry", value)

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
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _SegmentState:
    def __init__(__self__, *,
                 altitude: Optional[pulumi.Input['SegmentAltitudeArgs']] = None,
                 azimuth: Optional[pulumi.Input['SegmentAzimuthArgs']] = None,
                 cumulative_distance: Optional[pulumi.Input['SegmentCumulativeDistanceArgs']] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 kinds: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentKindArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]]] = None,
                 temperatures: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTemperatureArgs']]]] = None,
                 wind_directions: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentWindDirectionArgs']]]] = None,
                 wind_speeds: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentWindSpeedArgs']]]] = None):
        """
        Input properties used for looking up and filtering Segment resources.
        :param pulumi.Input['SegmentAltitudeArgs'] altitude: attribute of the resource
        :param pulumi.Input['SegmentAzimuthArgs'] azimuth: attribute of the resource
        :param pulumi.Input['SegmentCumulativeDistanceArgs'] cumulative_distance: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[Sequence[pulumi.Input['SegmentKindArgs']]] kinds: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]] tags: tags of the resource
        :param pulumi.Input[Sequence[pulumi.Input['SegmentTemperatureArgs']]] temperatures: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['SegmentWindDirectionArgs']]] wind_directions: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['SegmentWindSpeedArgs']]] wind_speeds: attribute of the resource
        """
        if altitude is not None:
            pulumi.set(__self__, "altitude", altitude)
        if azimuth is not None:
            pulumi.set(__self__, "azimuth", azimuth)
        if cumulative_distance is not None:
            pulumi.set(__self__, "cumulative_distance", cumulative_distance)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if kinds is not None:
            pulumi.set(__self__, "kinds", kinds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if temperatures is not None:
            pulumi.set(__self__, "temperatures", temperatures)
        if wind_directions is not None:
            pulumi.set(__self__, "wind_directions", wind_directions)
        if wind_speeds is not None:
            pulumi.set(__self__, "wind_speeds", wind_speeds)

    @property
    @pulumi.getter
    def altitude(self) -> Optional[pulumi.Input['SegmentAltitudeArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "altitude")

    @altitude.setter
    def altitude(self, value: Optional[pulumi.Input['SegmentAltitudeArgs']]):
        pulumi.set(self, "altitude", value)

    @property
    @pulumi.getter
    def azimuth(self) -> Optional[pulumi.Input['SegmentAzimuthArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "azimuth")

    @azimuth.setter
    def azimuth(self, value: Optional[pulumi.Input['SegmentAzimuthArgs']]):
        pulumi.set(self, "azimuth", value)

    @property
    @pulumi.getter(name="cumulativeDistance")
    def cumulative_distance(self) -> Optional[pulumi.Input['SegmentCumulativeDistanceArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "cumulative_distance")

    @cumulative_distance.setter
    def cumulative_distance(self, value: Optional[pulumi.Input['SegmentCumulativeDistanceArgs']]):
        pulumi.set(self, "cumulative_distance", value)

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
        geo position and shape of the resource
        """
        return pulumi.get(self, "geometry")

    @geometry.setter
    def geometry(self, value: Optional[pulumi.Input[str]]):
        pulumi.set(self, "geometry", value)

    @property
    @pulumi.getter
    def kinds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentKindArgs']]]]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kinds")

    @kinds.setter
    def kinds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentKindArgs']]]]):
        pulumi.set(self, "kinds", value)

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
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def temperatures(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTemperatureArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "temperatures")

    @temperatures.setter
    def temperatures(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentTemperatureArgs']]]]):
        pulumi.set(self, "temperatures", value)

    @property
    @pulumi.getter(name="windDirections")
    def wind_directions(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentWindDirectionArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "wind_directions")

    @wind_directions.setter
    def wind_directions(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentWindDirectionArgs']]]]):
        pulumi.set(self, "wind_directions", value)

    @property
    @pulumi.getter(name="windSpeeds")
    def wind_speeds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['SegmentWindSpeedArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "wind_speeds")

    @wind_speeds.setter
    def wind_speeds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['SegmentWindSpeedArgs']]]]):
        pulumi.set(self, "wind_speeds", value)


class Segment(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 altitude: Optional[pulumi.Input[Union['SegmentAltitudeArgs', 'SegmentAltitudeArgsDict']]] = None,
                 azimuth: Optional[pulumi.Input[Union['SegmentAzimuthArgs', 'SegmentAzimuthArgsDict']]] = None,
                 cumulative_distance: Optional[pulumi.Input[Union['SegmentCumulativeDistanceArgs', 'SegmentCumulativeDistanceArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentTagArgs', 'SegmentTagArgsDict']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/segment:Segment [options] splight_segment.<name> <segment_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['SegmentAltitudeArgs', 'SegmentAltitudeArgsDict']] altitude: attribute of the resource
        :param pulumi.Input[Union['SegmentAzimuthArgs', 'SegmentAzimuthArgsDict']] azimuth: attribute of the resource
        :param pulumi.Input[Union['SegmentCumulativeDistanceArgs', 'SegmentCumulativeDistanceArgsDict']] cumulative_distance: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['SegmentTagArgs', 'SegmentTagArgsDict']]]] tags: tags of the resource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: SegmentArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/segment:Segment [options] splight_segment.<name> <segment_id>
        ```

        :param str resource_name: The name of the resource.
        :param SegmentArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(SegmentArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 altitude: Optional[pulumi.Input[Union['SegmentAltitudeArgs', 'SegmentAltitudeArgsDict']]] = None,
                 azimuth: Optional[pulumi.Input[Union['SegmentAzimuthArgs', 'SegmentAzimuthArgsDict']]] = None,
                 cumulative_distance: Optional[pulumi.Input[Union['SegmentCumulativeDistanceArgs', 'SegmentCumulativeDistanceArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentTagArgs', 'SegmentTagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = SegmentArgs.__new__(SegmentArgs)

            if altitude is None and not opts.urn:
                raise TypeError("Missing required property 'altitude'")
            __props__.__dict__["altitude"] = altitude
            if azimuth is None and not opts.urn:
                raise TypeError("Missing required property 'azimuth'")
            __props__.__dict__["azimuth"] = azimuth
            if cumulative_distance is None and not opts.urn:
                raise TypeError("Missing required property 'cumulative_distance'")
            __props__.__dict__["cumulative_distance"] = cumulative_distance
            __props__.__dict__["description"] = description
            __props__.__dict__["geometry"] = geometry
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["kinds"] = None
            __props__.__dict__["temperatures"] = None
            __props__.__dict__["wind_directions"] = None
            __props__.__dict__["wind_speeds"] = None
        super(Segment, __self__).__init__(
            'splight:index/segment:Segment',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            altitude: Optional[pulumi.Input[Union['SegmentAltitudeArgs', 'SegmentAltitudeArgsDict']]] = None,
            azimuth: Optional[pulumi.Input[Union['SegmentAzimuthArgs', 'SegmentAzimuthArgsDict']]] = None,
            cumulative_distance: Optional[pulumi.Input[Union['SegmentCumulativeDistanceArgs', 'SegmentCumulativeDistanceArgsDict']]] = None,
            description: Optional[pulumi.Input[str]] = None,
            geometry: Optional[pulumi.Input[str]] = None,
            kinds: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentKindArgs', 'SegmentKindArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentTagArgs', 'SegmentTagArgsDict']]]]] = None,
            temperatures: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentTemperatureArgs', 'SegmentTemperatureArgsDict']]]]] = None,
            wind_directions: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentWindDirectionArgs', 'SegmentWindDirectionArgsDict']]]]] = None,
            wind_speeds: Optional[pulumi.Input[Sequence[pulumi.Input[Union['SegmentWindSpeedArgs', 'SegmentWindSpeedArgsDict']]]]] = None) -> 'Segment':
        """
        Get an existing Segment resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['SegmentAltitudeArgs', 'SegmentAltitudeArgsDict']] altitude: attribute of the resource
        :param pulumi.Input[Union['SegmentAzimuthArgs', 'SegmentAzimuthArgsDict']] azimuth: attribute of the resource
        :param pulumi.Input[Union['SegmentCumulativeDistanceArgs', 'SegmentCumulativeDistanceArgsDict']] cumulative_distance: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['SegmentKindArgs', 'SegmentKindArgsDict']]]] kinds: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['SegmentTagArgs', 'SegmentTagArgsDict']]]] tags: tags of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['SegmentTemperatureArgs', 'SegmentTemperatureArgsDict']]]] temperatures: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['SegmentWindDirectionArgs', 'SegmentWindDirectionArgsDict']]]] wind_directions: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['SegmentWindSpeedArgs', 'SegmentWindSpeedArgsDict']]]] wind_speeds: attribute of the resource
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _SegmentState.__new__(_SegmentState)

        __props__.__dict__["altitude"] = altitude
        __props__.__dict__["azimuth"] = azimuth
        __props__.__dict__["cumulative_distance"] = cumulative_distance
        __props__.__dict__["description"] = description
        __props__.__dict__["geometry"] = geometry
        __props__.__dict__["kinds"] = kinds
        __props__.__dict__["name"] = name
        __props__.__dict__["tags"] = tags
        __props__.__dict__["temperatures"] = temperatures
        __props__.__dict__["wind_directions"] = wind_directions
        __props__.__dict__["wind_speeds"] = wind_speeds
        return Segment(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter
    def altitude(self) -> pulumi.Output['outputs.SegmentAltitude']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "altitude")

    @property
    @pulumi.getter
    def azimuth(self) -> pulumi.Output['outputs.SegmentAzimuth']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "azimuth")

    @property
    @pulumi.getter(name="cumulativeDistance")
    def cumulative_distance(self) -> pulumi.Output['outputs.SegmentCumulativeDistance']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "cumulative_distance")

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
        geo position and shape of the resource
        """
        return pulumi.get(self, "geometry")

    @property
    @pulumi.getter
    def kinds(self) -> pulumi.Output[Sequence['outputs.SegmentKind']]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kinds")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.SegmentTag']]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def temperatures(self) -> pulumi.Output[Sequence['outputs.SegmentTemperature']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "temperatures")

    @property
    @pulumi.getter(name="windDirections")
    def wind_directions(self) -> pulumi.Output[Sequence['outputs.SegmentWindDirection']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "wind_directions")

    @property
    @pulumi.getter(name="windSpeeds")
    def wind_speeds(self) -> pulumi.Output[Sequence['outputs.SegmentWindSpeed']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "wind_speeds")

