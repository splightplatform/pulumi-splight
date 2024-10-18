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

__all__ = ['InverterArgs', 'Inverter']

@pulumi.input_type
class InverterArgs:
    def __init__(__self__, *,
                 energy_measurement_type: pulumi.Input['InverterEnergyMeasurementTypeArgs'],
                 make: pulumi.Input['InverterMakeArgs'],
                 max_active_power: pulumi.Input['InverterMaxActivePowerArgs'],
                 model: pulumi.Input['InverterModelArgs'],
                 serial_number: pulumi.Input['InverterSerialNumberArgs'],
                 description: Optional[pulumi.Input[str]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]]] = None):
        """
        The set of arguments for constructing a Inverter resource.
        :param pulumi.Input['InverterEnergyMeasurementTypeArgs'] energy_measurement_type: attribute of the resource
        :param pulumi.Input['InverterMakeArgs'] make: attribute of the resource
        :param pulumi.Input['InverterMaxActivePowerArgs'] max_active_power: attribute of the resource
        :param pulumi.Input['InverterModelArgs'] model: attribute of the resource
        :param pulumi.Input['InverterSerialNumberArgs'] serial_number: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]] tags: tags of the resource
        """
        pulumi.set(__self__, "energy_measurement_type", energy_measurement_type)
        pulumi.set(__self__, "make", make)
        pulumi.set(__self__, "max_active_power", max_active_power)
        pulumi.set(__self__, "model", model)
        pulumi.set(__self__, "serial_number", serial_number)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)

    @property
    @pulumi.getter(name="energyMeasurementType")
    def energy_measurement_type(self) -> pulumi.Input['InverterEnergyMeasurementTypeArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "energy_measurement_type")

    @energy_measurement_type.setter
    def energy_measurement_type(self, value: pulumi.Input['InverterEnergyMeasurementTypeArgs']):
        pulumi.set(self, "energy_measurement_type", value)

    @property
    @pulumi.getter
    def make(self) -> pulumi.Input['InverterMakeArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "make")

    @make.setter
    def make(self, value: pulumi.Input['InverterMakeArgs']):
        pulumi.set(self, "make", value)

    @property
    @pulumi.getter(name="maxActivePower")
    def max_active_power(self) -> pulumi.Input['InverterMaxActivePowerArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "max_active_power")

    @max_active_power.setter
    def max_active_power(self, value: pulumi.Input['InverterMaxActivePowerArgs']):
        pulumi.set(self, "max_active_power", value)

    @property
    @pulumi.getter
    def model(self) -> pulumi.Input['InverterModelArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: pulumi.Input['InverterModelArgs']):
        pulumi.set(self, "model", value)

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Input['InverterSerialNumberArgs']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: pulumi.Input['InverterSerialNumberArgs']):
        pulumi.set(self, "serial_number", value)

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
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]]]):
        pulumi.set(self, "tags", value)


@pulumi.input_type
class _InverterState:
    def __init__(__self__, *,
                 accumulated_energies: Optional[pulumi.Input[Sequence[pulumi.Input['InverterAccumulatedEnergyArgs']]]] = None,
                 active_powers: Optional[pulumi.Input[Sequence[pulumi.Input['InverterActivePowerArgs']]]] = None,
                 daily_energies: Optional[pulumi.Input[Sequence[pulumi.Input['InverterDailyEnergyArgs']]]] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 energy_measurement_type: Optional[pulumi.Input['InverterEnergyMeasurementTypeArgs']] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 kinds: Optional[pulumi.Input[Sequence[pulumi.Input['InverterKindArgs']]]] = None,
                 make: Optional[pulumi.Input['InverterMakeArgs']] = None,
                 max_active_power: Optional[pulumi.Input['InverterMaxActivePowerArgs']] = None,
                 model: Optional[pulumi.Input['InverterModelArgs']] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input['InverterSerialNumberArgs']] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]]] = None,
                 temperatures: Optional[pulumi.Input[Sequence[pulumi.Input['InverterTemperatureArgs']]]] = None):
        """
        Input properties used for looking up and filtering Inverter resources.
        :param pulumi.Input[Sequence[pulumi.Input['InverterAccumulatedEnergyArgs']]] accumulated_energies: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['InverterActivePowerArgs']]] active_powers: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['InverterDailyEnergyArgs']]] daily_energies: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input['InverterEnergyMeasurementTypeArgs'] energy_measurement_type: attribute of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[Sequence[pulumi.Input['InverterKindArgs']]] kinds: kind of the resource
        :param pulumi.Input['InverterMakeArgs'] make: attribute of the resource
        :param pulumi.Input['InverterMaxActivePowerArgs'] max_active_power: attribute of the resource
        :param pulumi.Input['InverterModelArgs'] model: attribute of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input['InverterSerialNumberArgs'] serial_number: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]] tags: tags of the resource
        :param pulumi.Input[Sequence[pulumi.Input['InverterTemperatureArgs']]] temperatures: attribute of the resource
        """
        if accumulated_energies is not None:
            pulumi.set(__self__, "accumulated_energies", accumulated_energies)
        if active_powers is not None:
            pulumi.set(__self__, "active_powers", active_powers)
        if daily_energies is not None:
            pulumi.set(__self__, "daily_energies", daily_energies)
        if description is not None:
            pulumi.set(__self__, "description", description)
        if energy_measurement_type is not None:
            pulumi.set(__self__, "energy_measurement_type", energy_measurement_type)
        if geometry is not None:
            pulumi.set(__self__, "geometry", geometry)
        if kinds is not None:
            pulumi.set(__self__, "kinds", kinds)
        if make is not None:
            pulumi.set(__self__, "make", make)
        if max_active_power is not None:
            pulumi.set(__self__, "max_active_power", max_active_power)
        if model is not None:
            pulumi.set(__self__, "model", model)
        if name is not None:
            pulumi.set(__self__, "name", name)
        if serial_number is not None:
            pulumi.set(__self__, "serial_number", serial_number)
        if tags is not None:
            pulumi.set(__self__, "tags", tags)
        if temperatures is not None:
            pulumi.set(__self__, "temperatures", temperatures)

    @property
    @pulumi.getter(name="accumulatedEnergies")
    def accumulated_energies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InverterAccumulatedEnergyArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "accumulated_energies")

    @accumulated_energies.setter
    def accumulated_energies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InverterAccumulatedEnergyArgs']]]]):
        pulumi.set(self, "accumulated_energies", value)

    @property
    @pulumi.getter(name="activePowers")
    def active_powers(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InverterActivePowerArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "active_powers")

    @active_powers.setter
    def active_powers(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InverterActivePowerArgs']]]]):
        pulumi.set(self, "active_powers", value)

    @property
    @pulumi.getter(name="dailyEnergies")
    def daily_energies(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InverterDailyEnergyArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "daily_energies")

    @daily_energies.setter
    def daily_energies(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InverterDailyEnergyArgs']]]]):
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
    @pulumi.getter(name="energyMeasurementType")
    def energy_measurement_type(self) -> Optional[pulumi.Input['InverterEnergyMeasurementTypeArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "energy_measurement_type")

    @energy_measurement_type.setter
    def energy_measurement_type(self, value: Optional[pulumi.Input['InverterEnergyMeasurementTypeArgs']]):
        pulumi.set(self, "energy_measurement_type", value)

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
    def kinds(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InverterKindArgs']]]]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kinds")

    @kinds.setter
    def kinds(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InverterKindArgs']]]]):
        pulumi.set(self, "kinds", value)

    @property
    @pulumi.getter
    def make(self) -> Optional[pulumi.Input['InverterMakeArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "make")

    @make.setter
    def make(self, value: Optional[pulumi.Input['InverterMakeArgs']]):
        pulumi.set(self, "make", value)

    @property
    @pulumi.getter(name="maxActivePower")
    def max_active_power(self) -> Optional[pulumi.Input['InverterMaxActivePowerArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "max_active_power")

    @max_active_power.setter
    def max_active_power(self, value: Optional[pulumi.Input['InverterMaxActivePowerArgs']]):
        pulumi.set(self, "max_active_power", value)

    @property
    @pulumi.getter
    def model(self) -> Optional[pulumi.Input['InverterModelArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "model")

    @model.setter
    def model(self, value: Optional[pulumi.Input['InverterModelArgs']]):
        pulumi.set(self, "model", value)

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
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> Optional[pulumi.Input['InverterSerialNumberArgs']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "serial_number")

    @serial_number.setter
    def serial_number(self, value: Optional[pulumi.Input['InverterSerialNumberArgs']]):
        pulumi.set(self, "serial_number", value)

    @property
    @pulumi.getter
    def tags(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @tags.setter
    def tags(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InverterTagArgs']]]]):
        pulumi.set(self, "tags", value)

    @property
    @pulumi.getter
    def temperatures(self) -> Optional[pulumi.Input[Sequence[pulumi.Input['InverterTemperatureArgs']]]]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "temperatures")

    @temperatures.setter
    def temperatures(self, value: Optional[pulumi.Input[Sequence[pulumi.Input['InverterTemperatureArgs']]]]):
        pulumi.set(self, "temperatures", value)


class Inverter(pulumi.CustomResource):
    @overload
    def __init__(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 energy_measurement_type: Optional[pulumi.Input[Union['InverterEnergyMeasurementTypeArgs', 'InverterEnergyMeasurementTypeArgsDict']]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 make: Optional[pulumi.Input[Union['InverterMakeArgs', 'InverterMakeArgsDict']]] = None,
                 max_active_power: Optional[pulumi.Input[Union['InverterMaxActivePowerArgs', 'InverterMaxActivePowerArgsDict']]] = None,
                 model: Optional[pulumi.Input[Union['InverterModelArgs', 'InverterModelArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[Union['InverterSerialNumberArgs', 'InverterSerialNumberArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterTagArgs', 'InverterTagArgsDict']]]]] = None,
                 __props__=None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/inverter:Inverter [options] splight_inverter.<name> <inverter_id>
        ```

        :param str resource_name: The name of the resource.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[Union['InverterEnergyMeasurementTypeArgs', 'InverterEnergyMeasurementTypeArgsDict']] energy_measurement_type: attribute of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[Union['InverterMakeArgs', 'InverterMakeArgsDict']] make: attribute of the resource
        :param pulumi.Input[Union['InverterMaxActivePowerArgs', 'InverterMaxActivePowerArgsDict']] max_active_power: attribute of the resource
        :param pulumi.Input[Union['InverterModelArgs', 'InverterModelArgsDict']] model: attribute of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Union['InverterSerialNumberArgs', 'InverterSerialNumberArgsDict']] serial_number: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['InverterTagArgs', 'InverterTagArgsDict']]]] tags: tags of the resource
        """
        ...
    @overload
    def __init__(__self__,
                 resource_name: str,
                 args: InverterArgs,
                 opts: Optional[pulumi.ResourceOptions] = None):
        """
        ## Example Usage

        ## Import

        ```sh
        $ pulumi import splight:index/inverter:Inverter [options] splight_inverter.<name> <inverter_id>
        ```

        :param str resource_name: The name of the resource.
        :param InverterArgs args: The arguments to use to populate this resource's properties.
        :param pulumi.ResourceOptions opts: Options for the resource.
        """
        ...
    def __init__(__self__, resource_name: str, *args, **kwargs):
        resource_args, opts = _utilities.get_resource_args_opts(InverterArgs, pulumi.ResourceOptions, *args, **kwargs)
        if resource_args is not None:
            __self__._internal_init(resource_name, opts, **resource_args.__dict__)
        else:
            __self__._internal_init(resource_name, *args, **kwargs)

    def _internal_init(__self__,
                 resource_name: str,
                 opts: Optional[pulumi.ResourceOptions] = None,
                 description: Optional[pulumi.Input[str]] = None,
                 energy_measurement_type: Optional[pulumi.Input[Union['InverterEnergyMeasurementTypeArgs', 'InverterEnergyMeasurementTypeArgsDict']]] = None,
                 geometry: Optional[pulumi.Input[str]] = None,
                 make: Optional[pulumi.Input[Union['InverterMakeArgs', 'InverterMakeArgsDict']]] = None,
                 max_active_power: Optional[pulumi.Input[Union['InverterMaxActivePowerArgs', 'InverterMaxActivePowerArgsDict']]] = None,
                 model: Optional[pulumi.Input[Union['InverterModelArgs', 'InverterModelArgsDict']]] = None,
                 name: Optional[pulumi.Input[str]] = None,
                 serial_number: Optional[pulumi.Input[Union['InverterSerialNumberArgs', 'InverterSerialNumberArgsDict']]] = None,
                 tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterTagArgs', 'InverterTagArgsDict']]]]] = None,
                 __props__=None):
        opts = pulumi.ResourceOptions.merge(_utilities.get_resource_opts_defaults(), opts)
        if not isinstance(opts, pulumi.ResourceOptions):
            raise TypeError('Expected resource options to be a ResourceOptions instance')
        if opts.id is None:
            if __props__ is not None:
                raise TypeError('__props__ is only valid when passed in combination with a valid opts.id to get an existing resource')
            __props__ = InverterArgs.__new__(InverterArgs)

            __props__.__dict__["description"] = description
            if energy_measurement_type is None and not opts.urn:
                raise TypeError("Missing required property 'energy_measurement_type'")
            __props__.__dict__["energy_measurement_type"] = energy_measurement_type
            __props__.__dict__["geometry"] = geometry
            if make is None and not opts.urn:
                raise TypeError("Missing required property 'make'")
            __props__.__dict__["make"] = make
            if max_active_power is None and not opts.urn:
                raise TypeError("Missing required property 'max_active_power'")
            __props__.__dict__["max_active_power"] = max_active_power
            if model is None and not opts.urn:
                raise TypeError("Missing required property 'model'")
            __props__.__dict__["model"] = model
            __props__.__dict__["name"] = name
            if serial_number is None and not opts.urn:
                raise TypeError("Missing required property 'serial_number'")
            __props__.__dict__["serial_number"] = serial_number
            __props__.__dict__["tags"] = tags
            __props__.__dict__["accumulated_energies"] = None
            __props__.__dict__["active_powers"] = None
            __props__.__dict__["daily_energies"] = None
            __props__.__dict__["kinds"] = None
            __props__.__dict__["temperatures"] = None
        super(Inverter, __self__).__init__(
            'splight:index/inverter:Inverter',
            resource_name,
            __props__,
            opts)

    @staticmethod
    def get(resource_name: str,
            id: pulumi.Input[str],
            opts: Optional[pulumi.ResourceOptions] = None,
            accumulated_energies: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterAccumulatedEnergyArgs', 'InverterAccumulatedEnergyArgsDict']]]]] = None,
            active_powers: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterActivePowerArgs', 'InverterActivePowerArgsDict']]]]] = None,
            daily_energies: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterDailyEnergyArgs', 'InverterDailyEnergyArgsDict']]]]] = None,
            description: Optional[pulumi.Input[str]] = None,
            energy_measurement_type: Optional[pulumi.Input[Union['InverterEnergyMeasurementTypeArgs', 'InverterEnergyMeasurementTypeArgsDict']]] = None,
            geometry: Optional[pulumi.Input[str]] = None,
            kinds: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterKindArgs', 'InverterKindArgsDict']]]]] = None,
            make: Optional[pulumi.Input[Union['InverterMakeArgs', 'InverterMakeArgsDict']]] = None,
            max_active_power: Optional[pulumi.Input[Union['InverterMaxActivePowerArgs', 'InverterMaxActivePowerArgsDict']]] = None,
            model: Optional[pulumi.Input[Union['InverterModelArgs', 'InverterModelArgsDict']]] = None,
            name: Optional[pulumi.Input[str]] = None,
            serial_number: Optional[pulumi.Input[Union['InverterSerialNumberArgs', 'InverterSerialNumberArgsDict']]] = None,
            tags: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterTagArgs', 'InverterTagArgsDict']]]]] = None,
            temperatures: Optional[pulumi.Input[Sequence[pulumi.Input[Union['InverterTemperatureArgs', 'InverterTemperatureArgsDict']]]]] = None) -> 'Inverter':
        """
        Get an existing Inverter resource's state with the given name, id, and optional extra
        properties used to qualify the lookup.

        :param str resource_name: The unique name of the resulting resource.
        :param pulumi.Input[str] id: The unique provider ID of the resource to lookup.
        :param pulumi.ResourceOptions opts: Options for the resource.
        :param pulumi.Input[Sequence[pulumi.Input[Union['InverterAccumulatedEnergyArgs', 'InverterAccumulatedEnergyArgsDict']]]] accumulated_energies: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['InverterActivePowerArgs', 'InverterActivePowerArgsDict']]]] active_powers: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['InverterDailyEnergyArgs', 'InverterDailyEnergyArgsDict']]]] daily_energies: attribute of the resource
        :param pulumi.Input[str] description: description of the resource
        :param pulumi.Input[Union['InverterEnergyMeasurementTypeArgs', 'InverterEnergyMeasurementTypeArgsDict']] energy_measurement_type: attribute of the resource
        :param pulumi.Input[str] geometry: geo position and shape of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['InverterKindArgs', 'InverterKindArgsDict']]]] kinds: kind of the resource
        :param pulumi.Input[Union['InverterMakeArgs', 'InverterMakeArgsDict']] make: attribute of the resource
        :param pulumi.Input[Union['InverterMaxActivePowerArgs', 'InverterMaxActivePowerArgsDict']] max_active_power: attribute of the resource
        :param pulumi.Input[Union['InverterModelArgs', 'InverterModelArgsDict']] model: attribute of the resource
        :param pulumi.Input[str] name: name of the resource
        :param pulumi.Input[Union['InverterSerialNumberArgs', 'InverterSerialNumberArgsDict']] serial_number: attribute of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['InverterTagArgs', 'InverterTagArgsDict']]]] tags: tags of the resource
        :param pulumi.Input[Sequence[pulumi.Input[Union['InverterTemperatureArgs', 'InverterTemperatureArgsDict']]]] temperatures: attribute of the resource
        """
        opts = pulumi.ResourceOptions.merge(opts, pulumi.ResourceOptions(id=id))

        __props__ = _InverterState.__new__(_InverterState)

        __props__.__dict__["accumulated_energies"] = accumulated_energies
        __props__.__dict__["active_powers"] = active_powers
        __props__.__dict__["daily_energies"] = daily_energies
        __props__.__dict__["description"] = description
        __props__.__dict__["energy_measurement_type"] = energy_measurement_type
        __props__.__dict__["geometry"] = geometry
        __props__.__dict__["kinds"] = kinds
        __props__.__dict__["make"] = make
        __props__.__dict__["max_active_power"] = max_active_power
        __props__.__dict__["model"] = model
        __props__.__dict__["name"] = name
        __props__.__dict__["serial_number"] = serial_number
        __props__.__dict__["tags"] = tags
        __props__.__dict__["temperatures"] = temperatures
        return Inverter(resource_name, opts=opts, __props__=__props__)

    @property
    @pulumi.getter(name="accumulatedEnergies")
    def accumulated_energies(self) -> pulumi.Output[Sequence['outputs.InverterAccumulatedEnergy']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "accumulated_energies")

    @property
    @pulumi.getter(name="activePowers")
    def active_powers(self) -> pulumi.Output[Sequence['outputs.InverterActivePower']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "active_powers")

    @property
    @pulumi.getter(name="dailyEnergies")
    def daily_energies(self) -> pulumi.Output[Sequence['outputs.InverterDailyEnergy']]:
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
    @pulumi.getter(name="energyMeasurementType")
    def energy_measurement_type(self) -> pulumi.Output['outputs.InverterEnergyMeasurementType']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "energy_measurement_type")

    @property
    @pulumi.getter
    def geometry(self) -> pulumi.Output[Optional[str]]:
        """
        geo position and shape of the resource
        """
        return pulumi.get(self, "geometry")

    @property
    @pulumi.getter
    def kinds(self) -> pulumi.Output[Sequence['outputs.InverterKind']]:
        """
        kind of the resource
        """
        return pulumi.get(self, "kinds")

    @property
    @pulumi.getter
    def make(self) -> pulumi.Output['outputs.InverterMake']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "make")

    @property
    @pulumi.getter(name="maxActivePower")
    def max_active_power(self) -> pulumi.Output['outputs.InverterMaxActivePower']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "max_active_power")

    @property
    @pulumi.getter
    def model(self) -> pulumi.Output['outputs.InverterModel']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "model")

    @property
    @pulumi.getter
    def name(self) -> pulumi.Output[str]:
        """
        name of the resource
        """
        return pulumi.get(self, "name")

    @property
    @pulumi.getter(name="serialNumber")
    def serial_number(self) -> pulumi.Output['outputs.InverterSerialNumber']:
        """
        attribute of the resource
        """
        return pulumi.get(self, "serial_number")

    @property
    @pulumi.getter
    def tags(self) -> pulumi.Output[Optional[Sequence['outputs.InverterTag']]]:
        """
        tags of the resource
        """
        return pulumi.get(self, "tags")

    @property
    @pulumi.getter
    def temperatures(self) -> pulumi.Output[Sequence['outputs.InverterTemperature']]:
        """
        attribute of the resource
        """
        return pulumi.get(self, "temperatures")
