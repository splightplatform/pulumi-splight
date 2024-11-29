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
    /// $ pulumi import splight:index/transformer:Transformer [options] splight_transformer.&lt;name&gt; &lt;transformer_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/transformer:Transformer")]
    public partial class Transformer : global::Pulumi.CustomResource
    {
        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("activePowerHvs")]
        public Output<ImmutableArray<Outputs.TransformerActivePowerHv>> ActivePowerHvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("activePowerLosses")]
        public Output<ImmutableArray<Outputs.TransformerActivePowerLoss>> ActivePowerLosses { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("activePowerLvs")]
        public Output<ImmutableArray<Outputs.TransformerActivePowerLv>> ActivePowerLvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("capacitance")]
        public Output<Outputs.TransformerCapacitance> Capacitance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("conductance")]
        public Output<Outputs.TransformerConductance> Conductance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("contingencies")]
        public Output<ImmutableArray<Outputs.TransformerContingency>> Contingencies { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("currentHvs")]
        public Output<ImmutableArray<Outputs.TransformerCurrentHv>> CurrentHvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("currentLvs")]
        public Output<ImmutableArray<Outputs.TransformerCurrentLv>> CurrentLvs { get; private set; } = null!;

        /// <summary>
        /// description of the resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Output("geometry")]
        public Output<string?> Geometry { get; private set; } = null!;

        /// <summary>
        /// kind of the resource
        /// </summary>
        [Output("kinds")]
        public Output<ImmutableArray<Outputs.TransformerKind>> Kinds { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maximumAllowedCurrent")]
        public Output<Outputs.TransformerMaximumAllowedCurrent> MaximumAllowedCurrent { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("maximumAllowedPower")]
        public Output<Outputs.TransformerMaximumAllowedPower> MaximumAllowedPower { get; private set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("reactance")]
        public Output<Outputs.TransformerReactance> Reactance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("reactivePowerHvs")]
        public Output<ImmutableArray<Outputs.TransformerReactivePowerHv>> ReactivePowerHvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("reactivePowerLosses")]
        public Output<ImmutableArray<Outputs.TransformerReactivePowerLoss>> ReactivePowerLosses { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("reactivePowerLvs")]
        public Output<ImmutableArray<Outputs.TransformerReactivePowerLv>> ReactivePowerLvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("resistance")]
        public Output<Outputs.TransformerResistance> Resistance { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("safetyMarginForPower")]
        public Output<Outputs.TransformerSafetyMarginForPower> SafetyMarginForPower { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("standardType")]
        public Output<Outputs.TransformerStandardType> StandardType { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("switchStatusHvs")]
        public Output<ImmutableArray<Outputs.TransformerSwitchStatusHv>> SwitchStatusHvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("switchStatusLvs")]
        public Output<ImmutableArray<Outputs.TransformerSwitchStatusLv>> SwitchStatusLvs { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.TransformerTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("tapPos")]
        public Output<Outputs.TransformerTapPos> TapPos { get; private set; } = null!;

        /// <summary>
        /// timezone that overrides location-based timezone of the resource
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("voltageHvs")]
        public Output<ImmutableArray<Outputs.TransformerVoltageHv>> VoltageHvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("voltageLvs")]
        public Output<ImmutableArray<Outputs.TransformerVoltageLv>> VoltageLvs { get; private set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Output("xnOhm")]
        public Output<Outputs.TransformerXnOhm> XnOhm { get; private set; } = null!;


        /// <summary>
        /// Create a Transformer resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Transformer(string name, TransformerArgs args, CustomResourceOptions? options = null)
            : base("splight:index/transformer:Transformer", name, args ?? new TransformerArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Transformer(string name, Input<string> id, TransformerState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/transformer:Transformer", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Transformer resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Transformer Get(string name, Input<string> id, TransformerState? state = null, CustomResourceOptions? options = null)
        {
            return new Transformer(name, id, state, options);
        }
    }

    public sealed class TransformerArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("capacitance", required: true)]
        public Input<Inputs.TransformerCapacitanceArgs> Capacitance { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("conductance", required: true)]
        public Input<Inputs.TransformerConductanceArgs> Conductance { get; set; } = null!;

        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Input("geometry")]
        public Input<string>? Geometry { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedCurrent", required: true)]
        public Input<Inputs.TransformerMaximumAllowedCurrentArgs> MaximumAllowedCurrent { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedPower", required: true)]
        public Input<Inputs.TransformerMaximumAllowedPowerArgs> MaximumAllowedPower { get; set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("reactance", required: true)]
        public Input<Inputs.TransformerReactanceArgs> Reactance { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("resistance", required: true)]
        public Input<Inputs.TransformerResistanceArgs> Resistance { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("safetyMarginForPower", required: true)]
        public Input<Inputs.TransformerSafetyMarginForPowerArgs> SafetyMarginForPower { get; set; } = null!;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("standardType", required: true)]
        public Input<Inputs.TransformerStandardTypeArgs> StandardType { get; set; } = null!;

        [Input("tags")]
        private InputList<Inputs.TransformerTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.TransformerTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TransformerTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("tapPos", required: true)]
        public Input<Inputs.TransformerTapPosArgs> TapPos { get; set; } = null!;

        /// <summary>
        /// timezone that overrides location-based timezone of the resource
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("xnOhm", required: true)]
        public Input<Inputs.TransformerXnOhmArgs> XnOhm { get; set; } = null!;

        public TransformerArgs()
        {
        }
        public static new TransformerArgs Empty => new TransformerArgs();
    }

    public sealed class TransformerState : global::Pulumi.ResourceArgs
    {
        [Input("activePowerHvs")]
        private InputList<Inputs.TransformerActivePowerHvGetArgs>? _activePowerHvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerActivePowerHvGetArgs> ActivePowerHvs
        {
            get => _activePowerHvs ?? (_activePowerHvs = new InputList<Inputs.TransformerActivePowerHvGetArgs>());
            set => _activePowerHvs = value;
        }

        [Input("activePowerLosses")]
        private InputList<Inputs.TransformerActivePowerLossGetArgs>? _activePowerLosses;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerActivePowerLossGetArgs> ActivePowerLosses
        {
            get => _activePowerLosses ?? (_activePowerLosses = new InputList<Inputs.TransformerActivePowerLossGetArgs>());
            set => _activePowerLosses = value;
        }

        [Input("activePowerLvs")]
        private InputList<Inputs.TransformerActivePowerLvGetArgs>? _activePowerLvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerActivePowerLvGetArgs> ActivePowerLvs
        {
            get => _activePowerLvs ?? (_activePowerLvs = new InputList<Inputs.TransformerActivePowerLvGetArgs>());
            set => _activePowerLvs = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("capacitance")]
        public Input<Inputs.TransformerCapacitanceGetArgs>? Capacitance { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("conductance")]
        public Input<Inputs.TransformerConductanceGetArgs>? Conductance { get; set; }

        [Input("contingencies")]
        private InputList<Inputs.TransformerContingencyGetArgs>? _contingencies;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerContingencyGetArgs> Contingencies
        {
            get => _contingencies ?? (_contingencies = new InputList<Inputs.TransformerContingencyGetArgs>());
            set => _contingencies = value;
        }

        [Input("currentHvs")]
        private InputList<Inputs.TransformerCurrentHvGetArgs>? _currentHvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerCurrentHvGetArgs> CurrentHvs
        {
            get => _currentHvs ?? (_currentHvs = new InputList<Inputs.TransformerCurrentHvGetArgs>());
            set => _currentHvs = value;
        }

        [Input("currentLvs")]
        private InputList<Inputs.TransformerCurrentLvGetArgs>? _currentLvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerCurrentLvGetArgs> CurrentLvs
        {
            get => _currentLvs ?? (_currentLvs = new InputList<Inputs.TransformerCurrentLvGetArgs>());
            set => _currentLvs = value;
        }

        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// geo position and shape of the resource
        /// </summary>
        [Input("geometry")]
        public Input<string>? Geometry { get; set; }

        [Input("kinds")]
        private InputList<Inputs.TransformerKindGetArgs>? _kinds;

        /// <summary>
        /// kind of the resource
        /// </summary>
        public InputList<Inputs.TransformerKindGetArgs> Kinds
        {
            get => _kinds ?? (_kinds = new InputList<Inputs.TransformerKindGetArgs>());
            set => _kinds = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedCurrent")]
        public Input<Inputs.TransformerMaximumAllowedCurrentGetArgs>? MaximumAllowedCurrent { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("maximumAllowedPower")]
        public Input<Inputs.TransformerMaximumAllowedPowerGetArgs>? MaximumAllowedPower { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("reactance")]
        public Input<Inputs.TransformerReactanceGetArgs>? Reactance { get; set; }

        [Input("reactivePowerHvs")]
        private InputList<Inputs.TransformerReactivePowerHvGetArgs>? _reactivePowerHvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerReactivePowerHvGetArgs> ReactivePowerHvs
        {
            get => _reactivePowerHvs ?? (_reactivePowerHvs = new InputList<Inputs.TransformerReactivePowerHvGetArgs>());
            set => _reactivePowerHvs = value;
        }

        [Input("reactivePowerLosses")]
        private InputList<Inputs.TransformerReactivePowerLossGetArgs>? _reactivePowerLosses;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerReactivePowerLossGetArgs> ReactivePowerLosses
        {
            get => _reactivePowerLosses ?? (_reactivePowerLosses = new InputList<Inputs.TransformerReactivePowerLossGetArgs>());
            set => _reactivePowerLosses = value;
        }

        [Input("reactivePowerLvs")]
        private InputList<Inputs.TransformerReactivePowerLvGetArgs>? _reactivePowerLvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerReactivePowerLvGetArgs> ReactivePowerLvs
        {
            get => _reactivePowerLvs ?? (_reactivePowerLvs = new InputList<Inputs.TransformerReactivePowerLvGetArgs>());
            set => _reactivePowerLvs = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("resistance")]
        public Input<Inputs.TransformerResistanceGetArgs>? Resistance { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("safetyMarginForPower")]
        public Input<Inputs.TransformerSafetyMarginForPowerGetArgs>? SafetyMarginForPower { get; set; }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("standardType")]
        public Input<Inputs.TransformerStandardTypeGetArgs>? StandardType { get; set; }

        [Input("switchStatusHvs")]
        private InputList<Inputs.TransformerSwitchStatusHvGetArgs>? _switchStatusHvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerSwitchStatusHvGetArgs> SwitchStatusHvs
        {
            get => _switchStatusHvs ?? (_switchStatusHvs = new InputList<Inputs.TransformerSwitchStatusHvGetArgs>());
            set => _switchStatusHvs = value;
        }

        [Input("switchStatusLvs")]
        private InputList<Inputs.TransformerSwitchStatusLvGetArgs>? _switchStatusLvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerSwitchStatusLvGetArgs> SwitchStatusLvs
        {
            get => _switchStatusLvs ?? (_switchStatusLvs = new InputList<Inputs.TransformerSwitchStatusLvGetArgs>());
            set => _switchStatusLvs = value;
        }

        [Input("tags")]
        private InputList<Inputs.TransformerTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.TransformerTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.TransformerTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("tapPos")]
        public Input<Inputs.TransformerTapPosGetArgs>? TapPos { get; set; }

        /// <summary>
        /// timezone that overrides location-based timezone of the resource
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        [Input("voltageHvs")]
        private InputList<Inputs.TransformerVoltageHvGetArgs>? _voltageHvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerVoltageHvGetArgs> VoltageHvs
        {
            get => _voltageHvs ?? (_voltageHvs = new InputList<Inputs.TransformerVoltageHvGetArgs>());
            set => _voltageHvs = value;
        }

        [Input("voltageLvs")]
        private InputList<Inputs.TransformerVoltageLvGetArgs>? _voltageLvs;

        /// <summary>
        /// attribute of the resource
        /// </summary>
        public InputList<Inputs.TransformerVoltageLvGetArgs> VoltageLvs
        {
            get => _voltageLvs ?? (_voltageLvs = new InputList<Inputs.TransformerVoltageLvGetArgs>());
            set => _voltageLvs = value;
        }

        /// <summary>
        /// attribute of the resource
        /// </summary>
        [Input("xnOhm")]
        public Input<Inputs.TransformerXnOhmGetArgs>? XnOhm { get; set; }

        public TransformerState()
        {
        }
        public static new TransformerState Empty => new TransformerState();
    }
}
