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

__all__ = ['GeneratorArgs', 'Generator']

@pulumi.input_type
class GeneratorArgs:
    def __init__(__self__, *,
                 co2_coefficient: pulumi.Input['GeneratorCo2CoefficientArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]]] = None):
        """
        The set of arguments for constructing a Generator resource.
        :param pulumi.Input['GeneratorCo2CoefficientArgs'] co2_coefficient: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]] tags: tags of the resource
        """
        pulumi.set(__self__, "co2_coefficient", co2_coefficient)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="co2Coefficient")
    def co2_coefficient(self) -> pulumi.Input['GeneratorCo2CoefficientArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "co2_coefficient")

    @co2_coefficient.setter
    def co2_coefficient(self, value: pulumi.Input['GeneratorCo2CoefficientArgs']):
        pulumi.set(self, "co2_coefficient", value)

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
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _GeneratorState:
    def __init__(__self__, *,
                 active_powers: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorActivePowerArgs']]]] = None,
                 co2_coefficient: Optional[pulumi.Input['GeneratorCo2CoefficientArgs']] = None,
                 daily_emission_avoideds: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEmissionAvoidedArgs']]]] = None,
                 daily_energies: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEnergyArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 kinds: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorKindArgs']]]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 reactive_powers: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorReactivePowerArgs']]]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]]] = None):
        """
        Input properties used for looking up and filtering Generator resources.
        :param pulumi.Input[Sequence[pulumi.Input['GeneratorActivePowerArgs']]] active_powers: attribute of the resource
        :param pulumi.Input['GeneratorCo2CoefficientArgs'] co2_coefficient: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEmissionAvoidedArgs']]] daily_emission_avoideds: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEnergyArgs']]] daily_energies: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[Sequence[pulumi.Input['GeneratorKindArgs']]] kinds: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input['GeneratorReactivePowerArgs']]] reactive_powers: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]] tags: tags of the resource
        """
        if active_powers is not None:
            pulumi.set(__self__, "active_powers", active_powers)
        if co2_coefficient is not None:
            pulumi.set(__self__, "co2_coefficient", co2_coefficient)
        if daily_emission_avoideds is not None:
            pulumi.set(__self__, "daily_emission_avoideds", daily_emission_avoideds)
        if daily_energies is not None:
            pulumi.set(__self__, "daily_energies", daily_energies)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if kinds is not None:
            pulumi.set(__self__, "kinds", kinds)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if reactive_powers is not None:
            pulumi.set(__self__, "reactive_powers", reactive_powers)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="activePowers")
    def active_powers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorActivePowerArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "active_powers")

    @active_powers.setter
    def active_powers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorActivePowerArgs']]]]):
        pulumi.set(self, "active_powers", value)

    @property
    @pulumi.getter(name="co2Coefficient")
    def co2_coefficient(self) -> Optional[pulumi.Input['GeneratorCo2CoefficientArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "co2_coefficient")

    @co2_coefficient.setter
    def co2_coefficient(self, value: Optional[pulumi.Input['GeneratorCo2CoefficientArgs']]):
        pulumi.set(self, "co2_coefficient", value)

    @property
    @pulumi.getter(name="dailyEmissionAvoideds")
    def daily_emission_avoideds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEmissionAvoidedArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "daily_emission_avoideds")

    @daily_emission_avoideds.setter
    def daily_emission_avoideds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEmissionAvoidedArgs']]]]):
        pulumi.set(self, "daily_emission_avoideds", value)

    @property
    @pulumi.getter(name="dailyEnergies")
    def daily_energies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEnergyArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "daily_energies")

    @daily_energies.setter
    def daily_energies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorDailyEnergyArgs']]]]):
        pulumi.set(self, "daily_energies", value)

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
    def kinds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorKindArgs']]]]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kinds")

    @kinds.setter
    def kinds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorKindArgs']]]]):
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
    @pulumi.getter(name="reactivePowers")
    def reactive_powers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorReactivePowerArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "reactive_powers")

    @reactive_powers.setter
    def reactive_powers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorReactivePowerArgs']]]]):
        pulumi.set(self, "reactive_powers", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['GeneratorTagArgs']]]]):
        pulumi.set(self, "tags", value)


class Generator(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 co2_coefficient: Optional[pulumi.Input[Union['GeneratorCo2CoefficientArgs', 'GeneratorCo2CoefficientArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorTagArgs', 'GeneratorTagArgsDict']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/generator:Generator [options] splight_generator.<name> <generator_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Union['GeneratorCo2CoefficientArgs', 'GeneratorCo2CoefficientArgsDict']] co2_coefficient: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['GeneratorTagArgs', 'GeneratorTagArgsDict']]]] tags: tags of the resource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: GeneratorArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/generator:Generator [options] splight_generator.<name> <generator_id>
        ```

        :param str resource_name: The name of the resource.
        :param GeneratorArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(GeneratorArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 co2_coefficient: Optional[pulumi.Input[Union['GeneratorCo2CoefficientArgs', 'GeneratorCo2CoefficientArgsDict']]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorTagArgs', 'GeneratorTagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = GeneratorArgs.__new__(GeneratorArgs)

            if co2_coefficient is None and not opts.urn:
                raise TypeError("Missing required property 'co2_coefficient'")
            __props__.__dict__["co2_coefficient"] = co2_coefficient
            __props__.__dict__["description"] = description
            __props__.__dict__["geometry"] = geometry
            __props__.__dict__["name"] = name
            __props__.__dict__["tags"] = tags
            __props__.__dict__["active_powers"] = None
            __props__.__dict__["daily_emission_avoideds"] = None
            __props__.__dict__["daily_energies"] = None
            __props__.__dict__["kinds"] = None
            __props__.__dict__["reactive_powers"] = None
        super(Generator, __self__).__init__(
            'splight:index/generator:Generator',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            active_powers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorActivePowerArgs', 'GeneratorActivePowerArgsDict']]]]] = None,
            co2_coefficient: Optional[pulumi.Input[Union['GeneratorCo2CoefficientArgs', 'GeneratorCo2CoefficientArgsDict']]] = None,
            daily_emission_avoideds: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorDailyEmissionAvoidedArgs', 'GeneratorDailyEmissionAvoidedArgsDict']]]]] = None,
            daily_energies: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorDailyEnergyArgs', 'GeneratorDailyEnergyArgsDict']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            geometry: Optional[pulumi.Input[str]] = None,
            kinds: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorKindArgs', 'GeneratorKindArgsDict']]]]] = None,
            name: Optional[pulumi.Input[str]] = None,
            reactive_powers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorReactivePowerArgs', 'GeneratorReactivePowerArgsDict']]]]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['GeneratorTagArgs', 'GeneratorTagArgsDict']]]]] = None) -> 'Generator':
        """
        Get an existing Generator resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['GeneratorActivePowerArgs', 'GeneratorActivePowerArgsDict']]]] active_powers: attribute of the resource
        :param pulumi.Input[Union['GeneratorCo2CoefficientArgs', 'GeneratorCo2CoefficientArgsDict']] co2_coefficient: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['GeneratorDailyEmissionAvoidedArgs', 'GeneratorDailyEmissionAvoidedArgsDict']]]] daily_emission_avoideds: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['GeneratorDailyEnergyArgs', 'GeneratorDailyEnergyArgsDict']]]] daily_energies: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['GeneratorKindArgs', 'GeneratorKindArgsDict']]]] kinds: kind of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['GeneratorReactivePowerArgs', 'GeneratorReactivePowerArgsDict']]]] reactive_powers: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['GeneratorTagArgs', 'GeneratorTagArgsDict']]]] tags: tags of the resource
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _GeneratorState.__new__(_GeneratorState)

        __props__.__dict__["active_powers"] = active_powers
        __props__.__dict__["co2_coefficient"] = co2_coefficient
        __props__.__dict__["daily_emission_avoideds"] = daily_emission_avoideds
        __props__.__dict__["daily_energies"] = daily_energies
        __props__.__dict__["description"] = description
        __props__.__dict__["geometry"] = geometry
        __props__.__dict__["kinds"] = kinds
        __props__.__dict__["name"] = name
        __props__.__dict__["reactive_powers"] = reactive_powers
        __props__.__dict__["tags"] = tags
        return Generator(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="activePowers")
    def active_powers(self) -> pulumi.Output[Sequence['outputs.GeneratorActivePower']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "active_powers")

    @property
    @pulumi.getter(name="co2Coefficient")
    def co2_coefficient(self) -> pulumi.Output['outputs.GeneratorCo2Coefficient']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "co2_coefficient")

    @property
    @pulumi.getter(name="dailyEmissionAvoideds")
    def daily_emission_avoideds(self) -> pulumi.Output[Sequence['outputs.GeneratorDailyEmissionAvoided']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "daily_emission_avoideds")

    @property
    @pulumi.getter(name="dailyEnergies")
    def daily_energies(self) -> pulumi.Output[Sequence['outputs.GeneratorDailyEnergy']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "daily_energies")

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
    def kinds(self) -> pulumi.Output[Sequence['outputs.GeneratorKind']]:
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
    @pulumi.getter(name="reactivePowers")
    def reactive_powers(self) -> pulumi.Output[Sequence['outputs.GeneratorReactivePower']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "reactive_powers")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.GeneratorTag']]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")
