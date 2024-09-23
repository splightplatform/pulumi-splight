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
    /// $ pulumi import splight:index/line:Line [options] splight_line.&lt;name&gt; &lt;line_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/line:Line")]
    public partial class Line : global::Pulumi.CustomResource
    {
        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("absorptivity")]
        public Output<Outputs.LineAbsorptivity> Absorptivity { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("activePowerEnds")]
        public Output<ImmutableArray<Outputs.LineActivePowerEnd>> ActivePowerEnds { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("activePowers")]
        public Output<ImmutableArray<Outputs.LineActivePower>> ActivePowers { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("ampacities")]
        public Output<ImmutableArray<Outputs.LineAmpacity>> Ampacities { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("atmosphere")]
        public Output<Outputs.LineAtmosphere> Atmosphere { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("capacitance")]
        public Output<Outputs.LineCapacitance> Capacitance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("conductance")]
        public Output<Outputs.LineConductance> Conductance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("currentRs")]
        public Output<ImmutableArray<Outputs.LineCurrentR>> CurrentRs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("currentS")]
        public Output<ImmutableArray<Outputs.LineCurrent>> CurrentS { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("currentTs")]
        public Output<ImmutableArray<Outputs.LineCurrentT>> CurrentTs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("currents")]
        public Output<ImmutableArray<Outputs.LineCurrent>> Currents { get; private set; } = null!;

        /// <summary>
        /// description of the resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("diameter")]
        public Output<Outputs.LineDiameter> Diameter { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("emissivity")]
        public Output<Outputs.LineEmissivity> Emissivity { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("energies")]
        public Output<ImmutableArray<Outputs.LineEnergy>> Energies { get; private set; } = null!;

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Output("geometry")]
        public Output<string?> Geometry { get; private set; } = null!;

        /// <summary>
        /// kind of the resource
        /// </summary>
        [Output("kinds")]
        public Output<ImmutableArray<Outputs.LineKind>> Kinds { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("length")]
        public Output<Outputs.LineLength> Length { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maxTemperatures")]
        public Output<ImmutableArray<Outputs.LineMaxTemperature>> MaxTemperatures { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maximumAllowedCurrent")]
        public Output<Outputs.LineMaximumAllowedCurrent> MaximumAllowedCurrent { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maximumAllowedPower")]
        public Output<Outputs.LineMaximumAllowedPower> MaximumAllowedPower { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maximumAllowedTemperature")]
        public Output<Outputs.LineMaximumAllowedTemperature> MaximumAllowedTemperature { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maximumAllowedTemperatureLte")]
        public Output<Outputs.LineMaximumAllowedTemperatureLte> MaximumAllowedTemperatureLte { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maximumAllowedTemperatureSte")]
        public Output<Outputs.LineMaximumAllowedTemperatureSte> MaximumAllowedTemperatureSte { get; private set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("numberOfConductors")]
        public Output<Outputs.LineNumberOfConductors> NumberOfConductors { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("reactance")]
        public Output<Outputs.LineReactance> Reactance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("reactivePowers")]
        public Output<ImmutableArray<Outputs.LineReactivePower>> ReactivePowers { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("referenceResistance")]
        public Output<Outputs.LineReferenceResistance> ReferenceResistance { get; private set; } = null!;

        /// <summary>
        /// related assets of the resource
        /// </summary>
        [Output("relatedAssets")]
        public Output<ImmutableArray<Outputs.LineRelatedAsset>> RelatedAssets { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("resistance")]
        public Output<Outputs.LineResistance> Resistance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("safetyMarginForPower")]
        public Output<Outputs.LineSafetyMarginForPower> SafetyMarginForPower { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("susceptance")]
        public Output<Outputs.LineSusceptance> Susceptance { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.LineTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("temperatureCoeffResistance")]
        public Output<Outputs.LineTemperatureCoeffResistance> TemperatureCoeffResistance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("voltageRs")]
        public Output<ImmutableArray<Outputs.LineVoltageR>> VoltageRs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("voltageSts")]
        public Output<ImmutableArray<Outputs.LineVoltageSt>> VoltageSts { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("voltageTrs")]
        public Output<ImmutableArray<Outputs.LineVoltageTr>> VoltageTrs { get; private set; } = null!;


        /// <summary>
        /// Create a Line resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Line(string name, LineArgs args, CustomResourceOptions? options = null)
            : base("splight:index/line:Line", name, args ?? new LineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Line(string name, Input<string> id, LineState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/line:Line", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Line resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Line Get(string name, Input<string> id, LineState? state = null, CustomResourceOptions? options = null)
        {
            return new Line(name, id, state, options);
        }
    }

    public sealed class LineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("absorptivity", required: true)]
        public Input<Inputs.LineAbsorptivityArgs> Absorptivity { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("atmosphere", required: true)]
        public Input<Inputs.LineAtmosphereArgs> Atmosphere { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("capacitance", required: true)]
        public Input<Inputs.LineCapacitanceArgs> Capacitance { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("conductance", required: true)]
        public Input<Inputs.LineConductanceArgs> Conductance { get; set; } = null!;

        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("diameter", required: true)]
        public Input<Inputs.LineDiameterArgs> Diameter { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("emissivity", required: true)]
        public Input<Inputs.LineEmissivityArgs> Emissivity { get; set; } = null!;

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Input("geometry")]
        public Input<string>? Geometry { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("length", required: true)]
        public Input<Inputs.LineLengthArgs> Length { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedCurrent", required: true)]
        public Input<Inputs.LineMaximumAllowedCurrentArgs> MaximumAllowedCurrent { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedPower", required: true)]
        public Input<Inputs.LineMaximumAllowedPowerArgs> MaximumAllowedPower { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedTemperature", required: true)]
        public Input<Inputs.LineMaximumAllowedTemperatureArgs> MaximumAllowedTemperature { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedTemperatureLte", required: true)]
        public Input<Inputs.LineMaximumAllowedTemperatureLteArgs> MaximumAllowedTemperatureLte { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedTemperatureSte", required: true)]
        public Input<Inputs.LineMaximumAllowedTemperatureSteArgs> MaximumAllowedTemperatureSte { get; set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("numberOfConductors", required: true)]
        public Input<Inputs.LineNumberOfConductorsArgs> NumberOfConductors { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("reactance", required: true)]
        public Input<Inputs.LineReactanceArgs> Reactance { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("referenceResistance", required: true)]
        public Input<Inputs.LineReferenceResistanceArgs> ReferenceResistance { get; set; } = null!;

        [Input("relatedAssets")]
        private InputList<Inputs.LineRelatedAssetArgs>? _relatedAssets;

        /// <summary>
        /// related assets of the resource
        /// </summary>
        public InputList<Inputs.LineRelatedAssetArgs> RelatedAssets
        {
            get => _relatedAssets ?? (_relatedAssets = new InputList<Inputs.LineRelatedAssetArgs>());
            set => _relatedAssets = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("resistance", required: true)]
        public Input<Inputs.LineResistanceArgs> Resistance { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("safetyMarginForPower", required: true)]
        public Input<Inputs.LineSafetyMarginForPowerArgs> SafetyMarginForPower { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("susceptance", required: true)]
        public Input<Inputs.LineSusceptanceArgs> Susceptance { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.LineTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.LineTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LineTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("temperatureCoeffResistance", required: true)]
        public Input<Inputs.LineTemperatureCoeffResistanceArgs> TemperatureCoeffResistance { get; set; } = null!;

        public LineArgs()
        {
        }
        public static new LineArgs Empty => new LineArgs();
    }

    public sealed class LineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("absorptivity")]
        public Input<Inputs.LineAbsorptivityGetArgs>? Absorptivity { get; set; }

        [Input("activePowerEnds")]
        private InputList<Inputs.LineActivePowerEndGetArgs>? _activePowerEnds;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineActivePowerEndGetArgs> ActivePowerEnds
        {
            get => _activePowerEnds ?? (_activePowerEnds = new InputList<Inputs.LineActivePowerEndGetArgs>());
            set => _activePowerEnds = value;
        }

        [Input("activePowers")]
        private InputList<Inputs.LineActivePowerGetArgs>? _activePowers;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineActivePowerGetArgs> ActivePowers
        {
            get => _activePowers ?? (_activePowers = new InputList<Inputs.LineActivePowerGetArgs>());
            set => _activePowers = value;
        }

        [Input("ampacities")]
        private InputList<Inputs.LineAmpacityGetArgs>? _ampacities;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineAmpacityGetArgs> Ampacities
        {
            get => _ampacities ?? (_ampacities = new InputList<Inputs.LineAmpacityGetArgs>());
            set => _ampacities = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("atmosphere")]
        public Input<Inputs.LineAtmosphereGetArgs>? Atmosphere { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("capacitance")]
        public Input<Inputs.LineCapacitanceGetArgs>? Capacitance { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("conductance")]
        public Input<Inputs.LineConductanceGetArgs>? Conductance { get; set; }

        [Input("currentRs")]
        private InputList<Inputs.LineCurrentRGetArgs>? _currentRs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineCurrentRGetArgs> CurrentRs
        {
            get => _currentRs ?? (_currentRs = new InputList<Inputs.LineCurrentRGetArgs>());
            set => _currentRs = value;
        }

        [Input("currentS")]
        private InputList<Inputs.LineCurrentGetArgs>? _currentS;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineCurrentGetArgs> CurrentS
        {
            get => _currentS ?? (_currentS = new InputList<Inputs.LineCurrentGetArgs>());
            set => _currentS = value;
        }

        [Input("currentTs")]
        private InputList<Inputs.LineCurrentTGetArgs>? _currentTs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineCurrentTGetArgs> CurrentTs
        {
            get => _currentTs ?? (_currentTs = new InputList<Inputs.LineCurrentTGetArgs>());
            set => _currentTs = value;
        }

        [Input("currents")]
        private InputList<Inputs.LineCurrentGetArgs>? _currents;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineCurrentGetArgs> Currents
        {
            get => _currents ?? (_currents = new InputList<Inputs.LineCurrentGetArgs>());
            set => _currents = value;
        }

        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("diameter")]
        public Input<Inputs.LineDiameterGetArgs>? Diameter { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("emissivity")]
        public Input<Inputs.LineEmissivityGetArgs>? Emissivity { get; set; }

        [Input("energies")]
        private InputList<Inputs.LineEnergyGetArgs>? _energies;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineEnergyGetArgs> Energies
        {
            get => _energies ?? (_energies = new InputList<Inputs.LineEnergyGetArgs>());
            set => _energies = value;
        }

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Input("geometry")]
        public Input<string>? Geometry { get; set; }

        [Input("kinds")]
        private InputList<Inputs.LineKindGetArgs>? _kinds;

        /// <summary>
        /// kind of the resource
        /// </summary>
        public InputList<Inputs.LineKindGetArgs> Kinds
        {
            get => _kinds ?? (_kinds = new InputList<Inputs.LineKindGetArgs>());
            set => _kinds = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("length")]
        public Input<Inputs.LineLengthGetArgs>? Length { get; set; }

        [Input("maxTemperatures")]
        private InputList<Inputs.LineMaxTemperatureGetArgs>? _maxTemperatures;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineMaxTemperatureGetArgs> MaxTemperatures
        {
            get => _maxTemperatures ?? (_maxTemperatures = new InputList<Inputs.LineMaxTemperatureGetArgs>());
            set => _maxTemperatures = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedCurrent")]
        public Input<Inputs.LineMaximumAllowedCurrentGetArgs>? MaximumAllowedCurrent { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedPower")]
        public Input<Inputs.LineMaximumAllowedPowerGetArgs>? MaximumAllowedPower { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedTemperature")]
        public Input<Inputs.LineMaximumAllowedTemperatureGetArgs>? MaximumAllowedTemperature { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedTemperatureLte")]
        public Input<Inputs.LineMaximumAllowedTemperatureLteGetArgs>? MaximumAllowedTemperatureLte { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedTemperatureSte")]
        public Input<Inputs.LineMaximumAllowedTemperatureSteGetArgs>? MaximumAllowedTemperatureSte { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("numberOfConductors")]
        public Input<Inputs.LineNumberOfConductorsGetArgs>? NumberOfConductors { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("reactance")]
        public Input<Inputs.LineReactanceGetArgs>? Reactance { get; set; }

        [Input("reactivePowers")]
        private InputList<Inputs.LineReactivePowerGetArgs>? _reactivePowers;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineReactivePowerGetArgs> ReactivePowers
        {
            get => _reactivePowers ?? (_reactivePowers = new InputList<Inputs.LineReactivePowerGetArgs>());
            set => _reactivePowers = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("referenceResistance")]
        public Input<Inputs.LineReferenceResistanceGetArgs>? ReferenceResistance { get; set; }

        [Input("relatedAssets")]
        private InputList<Inputs.LineRelatedAssetGetArgs>? _relatedAssets;

        /// <summary>
        /// related assets of the resource
        /// </summary>
        public InputList<Inputs.LineRelatedAssetGetArgs> RelatedAssets
        {
            get => _relatedAssets ?? (_relatedAssets = new InputList<Inputs.LineRelatedAssetGetArgs>());
            set => _relatedAssets = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("resistance")]
        public Input<Inputs.LineResistanceGetArgs>? Resistance { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("safetyMarginForPower")]
        public Input<Inputs.LineSafetyMarginForPowerGetArgs>? SafetyMarginForPower { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("susceptance")]
        public Input<Inputs.LineSusceptanceGetArgs>? Susceptance { get; set; }

        [Input("tags")]
        private InputList<Inputs.LineTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.LineTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.LineTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("temperatureCoeffResistance")]
        public Input<Inputs.LineTemperatureCoeffResistanceGetArgs>? TemperatureCoeffResistance { get; set; }

        [Input("voltageRs")]
        private InputList<Inputs.LineVoltageRGetArgs>? _voltageRs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineVoltageRGetArgs> VoltageRs
        {
            get => _voltageRs ?? (_voltageRs = new InputList<Inputs.LineVoltageRGetArgs>());
            set => _voltageRs = value;
        }

        [Input("voltageSts")]
        private InputList<Inputs.LineVoltageStGetArgs>? _voltageSts;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineVoltageStGetArgs> VoltageSts
        {
            get => _voltageSts ?? (_voltageSts = new InputList<Inputs.LineVoltageStGetArgs>());
            set => _voltageSts = value;
        }

        [Input("voltageTrs")]
        private InputList<Inputs.LineVoltageTrGetArgs>? _voltageTrs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.LineVoltageTrGetArgs> VoltageTrs
        {
            get => _voltageTrs ?? (_voltageTrs = new InputList<Inputs.LineVoltageTrGetArgs>());
            set => _voltageTrs = value;
        }

        public LineState()
        {
        }
        public static new LineState Empty => new LineState();
    }
}