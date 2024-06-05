// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splight
{
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Splight = Pulumi.Splight;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var componentTestRoutine = new Splight.ComponentRoutine("componentTestRoutine", new()
    ///     {
    ///         Description = "Created with Terraform",
    ///         Type = "IncomingRoutine",
    ///         ComponentId = "1234-1234-1234-1234",
    ///         Configs = new[]
    ///         {
    ///             new Splight.Inputs.ComponentRoutineConfigArgs
    ///             {
    ///                 Name = "config_param",
    ///                 Type = "bool",
    ///                 Value = "true",
    ///                 Multiple = false,
    ///                 Required = true,
    ///                 Sensitive = false,
    ///                 Description = JsonSerializer.Serialize("Created with Terraform123123"),
    ///             },
    ///         },
    ///         Outputs = new[]
    ///         {
    ///             new Splight.Inputs.ComponentRoutineOutputArgs
    ///             {
    ///                 Name = "address",
    ///                 Description = "destination address for data to be pushed",
    ///                 Type = "DataAddress",
    ///                 ValueType = "Number",
    ///                 Multiple = false,
    ///                 Required = true,
    ///                 Value = new Splight.Inputs.ComponentRoutineOutputValueArgs
    ///                 {
    ///                     Asset = "1234-1234-1234-1234",
    ///                     Attribute = "1234-1234-1234-1234",
    ///                 },
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import splight:index/componentRoutine:ComponentRoutine [options] splight_component_routine.&lt;name&gt; &lt;component_routine_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/componentRoutine:ComponentRoutine")]
    public partial class ComponentRoutine : global::Pulumi.CustomResource
    {
        /// <summary>
        /// reference to component to be attached
        /// </summary>
        [Output("componentId")]
        public Output<string> ComponentId { get; private set; } = null!;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        [Output("configs")]
        public Output<ImmutableArray<Outputs.ComponentRoutineConfig>> Configs { get; private set; } = null!;

        /// <summary>
        /// optional complementary information about the routine
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// asset attribute where to read data. Only valid for OutgoingRoutine
        /// </summary>
        [Output("inputs")]
        public Output<ImmutableArray<Outputs.ComponentRoutineInput>> Inputs { get; private set; } = null!;

        /// <summary>
        /// name of the routine
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// asset attribute where to ingest data. Only valid for IncomingRoutine
        /// </summary>
        [Output("outputs")]
        public Output<ImmutableArray<Outputs.ComponentRoutineOutput>> Outputs { get; private set; } = null!;

        /// <summary>
        /// [IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;


        /// <summary>
        /// Create a ComponentRoutine resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ComponentRoutine(string name, ComponentRoutineArgs args, CustomResourceOptions? options = null)
            : base("splight:index/componentRoutine:ComponentRoutine", name, args ?? new ComponentRoutineArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ComponentRoutine(string name, Input<string> id, ComponentRoutineState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/componentRoutine:ComponentRoutine", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing ComponentRoutine resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ComponentRoutine Get(string name, Input<string> id, ComponentRoutineState? state = null, CustomResourceOptions? options = null)
        {
            return new ComponentRoutine(name, id, state, options);
        }
    }

    public sealed class ComponentRoutineArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// reference to component to be attached
        /// </summary>
        [Input("componentId", required: true)]
        public Input<string> ComponentId { get; set; } = null!;

        [Input("configs")]
        private InputList<Inputs.ComponentRoutineConfigArgs>? _configs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ComponentRoutineConfigArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ComponentRoutineConfigArgs>());
            set => _configs = value;
        }

        /// <summary>
        /// optional complementary information about the routine
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputs")]
        private InputList<Inputs.ComponentRoutineInputArgs>? _inputs;

        /// <summary>
        /// asset attribute where to read data. Only valid for OutgoingRoutine
        /// </summary>
        public InputList<Inputs.ComponentRoutineInputArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.ComponentRoutineInputArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// name of the routine
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outputs")]
        private InputList<Inputs.ComponentRoutineOutputArgs>? _outputs;

        /// <summary>
        /// asset attribute where to ingest data. Only valid for IncomingRoutine
        /// </summary>
        public InputList<Inputs.ComponentRoutineOutputArgs> Outputs
        {
            get => _outputs ?? (_outputs = new InputList<Inputs.ComponentRoutineOutputArgs>());
            set => _outputs = value;
        }

        /// <summary>
        /// [IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public ComponentRoutineArgs()
        {
        }
        public static new ComponentRoutineArgs Empty => new ComponentRoutineArgs();
    }

    public sealed class ComponentRoutineState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// reference to component to be attached
        /// </summary>
        [Input("componentId")]
        public Input<string>? ComponentId { get; set; }

        [Input("configs")]
        private InputList<Inputs.ComponentRoutineConfigGetArgs>? _configs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ComponentRoutineConfigGetArgs> Configs
        {
            get => _configs ?? (_configs = new InputList<Inputs.ComponentRoutineConfigGetArgs>());
            set => _configs = value;
        }

        /// <summary>
        /// optional complementary information about the routine
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputs")]
        private InputList<Inputs.ComponentRoutineInputGetArgs>? _inputs;

        /// <summary>
        /// asset attribute where to read data. Only valid for OutgoingRoutine
        /// </summary>
        public InputList<Inputs.ComponentRoutineInputGetArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.ComponentRoutineInputGetArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// name of the routine
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("outputs")]
        private InputList<Inputs.ComponentRoutineOutputGetArgs>? _outputs;

        /// <summary>
        /// asset attribute where to ingest data. Only valid for IncomingRoutine
        /// </summary>
        public InputList<Inputs.ComponentRoutineOutputGetArgs> Outputs
        {
            get => _outputs ?? (_outputs = new InputList<Inputs.ComponentRoutineOutputGetArgs>());
            set => _outputs = value;
        }

        /// <summary>
        /// [IncomingRoutine|OutgoingRoutine] direction of the data flow (from device to system or from system to device)
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        public ComponentRoutineState()
        {
        }
        public static new ComponentRoutineState Empty => new ComponentRoutineState();
    }
}