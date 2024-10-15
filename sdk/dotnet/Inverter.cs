// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Splight.Splight
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import splight:index/inverter:Inverter [options] splight_inverter.&lt;name&gt; &lt;inverter_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/inverter:Inverter")]
    public partial class Inverter : global::Pulumi.CustomResource
    {
        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("accumulatedEnergies")]
        public Output<ImmutableArray<Outputs.InverterAccumulatedEnergy>> AccumulatedEnergies { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("activePowers")]
        public Output<ImmutableArray<Outputs.InverterActivePower>> ActivePowers { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("dailyEnergies")]
        public Output<ImmutableArray<Outputs.InverterDailyEnergy>> DailyEnergies { get; private set; } = null!;

        /// <summary>
        /// description of the resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("energyMeasurementType")]
        public Output<Outputs.InverterEnergyMeasurementType> EnergyMeasurementType { get; private set; } = null!;

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Output("geometry")]
        public Output<string?> Geometry { get; private set; } = null!;

        /// <summary>
        /// kind of the resource
        /// </summary>
        [Output("kinds")]
        public Output<ImmutableArray<Outputs.InverterKind>> Kinds { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("make")]
        public Output<Outputs.InverterMake> Make { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maxActivePower")]
        public Output<Outputs.InverterMaxActivePower> MaxActivePower { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("model")]
        public Output<Outputs.InverterModel> Model { get; private set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("serialNumber")]
        public Output<Outputs.InverterSerialNumber> SerialNumber { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.InverterTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("temperatures")]
        public Output<ImmutableArray<Outputs.InverterTemperature>> Temperatures { get; private set; } = null!;


        /// <summary>
        /// Create a Inverter resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Inverter(string name, InverterArgs args, CustomResourceOptions? options = null)
            : base("splight:index/inverter:Inverter", name, args ?? new InverterArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Inverter(string name, Input<string> id, InverterState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/inverter:Inverter", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/splightplatform",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing Inverter resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Inverter Get(string name, Input<string> id, InverterState? state = null, CustomResourceOptions? options = null)
        {
            return new Inverter(name, id, state, options);
        }
    }

    public sealed class InverterArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("energyMeasurementType", required: true)]
        public Input<Inputs.InverterEnergyMeasurementTypeArgs> EnergyMeasurementType { get; set; } = null!;

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Input("geometry")]
        public Input<string>? Geometry { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("make", required: true)]
        public Input<Inputs.InverterMakeArgs> Make { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maxActivePower", required: true)]
        public Input<Inputs.InverterMaxActivePowerArgs> MaxActivePower { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("model", required: true)]
        public Input<Inputs.InverterModelArgs> Model { get; set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("serialNumber", required: true)]
        public Input<Inputs.InverterSerialNumberArgs> SerialNumber { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.InverterTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.InverterTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InverterTagArgs>());
            set => _tags = value;
        }

        public InverterArgs()
        {
        }
        public static new InverterArgs Empty => new InverterArgs();
    }

    public sealed class InverterState : global::Pulumi.ResourceArgs
    {
        [Input("accumulatedEnergies")]
        private InputList<Inputs.InverterAccumulatedEnergyGetArgs>? _accumulatedEnergies;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.InverterAccumulatedEnergyGetArgs> AccumulatedEnergies
        {
            get => _accumulatedEnergies ?? (_accumulatedEnergies = new InputList<Inputs.InverterAccumulatedEnergyGetArgs>());
            set => _accumulatedEnergies = value;
        }

        [Input("activePowers")]
        private InputList<Inputs.InverterActivePowerGetArgs>? _activePowers;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.InverterActivePowerGetArgs> ActivePowers
        {
            get => _activePowers ?? (_activePowers = new InputList<Inputs.InverterActivePowerGetArgs>());
            set => _activePowers = value;
        }

        [Input("dailyEnergies")]
        private InputList<Inputs.InverterDailyEnergyGetArgs>? _dailyEnergies;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.InverterDailyEnergyGetArgs> DailyEnergies
        {
            get => _dailyEnergies ?? (_dailyEnergies = new InputList<Inputs.InverterDailyEnergyGetArgs>());
            set => _dailyEnergies = value;
        }

        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("energyMeasurementType")]
        public Input<Inputs.InverterEnergyMeasurementTypeGetArgs>? EnergyMeasurementType { get; set; }

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Input("geometry")]
        public Input<string>? Geometry { get; set; }

        [Input("kinds")]
        private InputList<Inputs.InverterKindGetArgs>? _kinds;

        /// <summary>
        /// kind of the resource
        /// </summary>
        public InputList<Inputs.InverterKindGetArgs> Kinds
        {
            get => _kinds ?? (_kinds = new InputList<Inputs.InverterKindGetArgs>());
            set => _kinds = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("make")]
        public Input<Inputs.InverterMakeGetArgs>? Make { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maxActivePower")]
        public Input<Inputs.InverterMaxActivePowerGetArgs>? MaxActivePower { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("model")]
        public Input<Inputs.InverterModelGetArgs>? Model { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("serialNumber")]
        public Input<Inputs.InverterSerialNumberGetArgs>? SerialNumber { get; set; }

        [Input("tags")]
        private InputList<Inputs.InverterTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.InverterTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.InverterTagGetArgs>());
            set => _tags = value;
        }

        [Input("temperatures")]
        private InputList<Inputs.InverterTemperatureGetArgs>? _temperatures;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.InverterTemperatureGetArgs> Temperatures
        {
            get => _temperatures ?? (_temperatures = new InputList<Inputs.InverterTemperatureGetArgs>());
            set => _temperatures = value;
        }

        public InverterState()
        {
        }
        public static new InverterState Empty => new InverterState();
    }
}
