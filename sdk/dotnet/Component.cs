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
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Splight = Splight.Splight;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var componentTest = new Splight.Component("componentTest", new()
    ///     {
    ///         Description = "Created with Terraform",
    ///         Version = "Random-3.1.0",
    ///         Inputs = new[]
    ///         {
    ///             new Splight.Inputs.ComponentInputArgs
    ///             {
    ///                 Name = "period",
    ///                 Type = "int",
    ///                 Value = JsonSerializer.Serialize(10),
    ///                 Multiple = false,
    ///                 Required = false,
    ///                 Sensitive = false,
    ///                 Description = "",
    ///             },
    ///             new Splight.Inputs.ComponentInputArgs
    ///             {
    ///                 Name = "min",
    ///                 Type = "int",
    ///                 Value = JsonSerializer.Serialize(1),
    ///                 Multiple = false,
    ///                 Required = false,
    ///                 Sensitive = false,
    ///                 Description = "",
    ///             },
    ///             new Splight.Inputs.ComponentInputArgs
    ///             {
    ///                 Name = "max",
    ///                 Type = "int",
    ///                 Value = JsonSerializer.Serialize(150),
    ///                 Multiple = false,
    ///                 Required = false,
    ///                 Sensitive = false,
    ///                 Description = "",
    ///             },
    ///             new Splight.Inputs.ComponentInputArgs
    ///             {
    ///                 Name = "max_iterations",
    ///                 Type = "int",
    ///                 Value = JsonSerializer.Serialize(3),
    ///                 Multiple = false,
    ///                 Required = false,
    ///                 Sensitive = false,
    ///                 Description = "",
    ///             },
    ///             new Splight.Inputs.ComponentInputArgs
    ///             {
    ///                 Name = "should_crash",
    ///                 Type = "bool",
    ///                 Value = JsonSerializer.Serialize("true"),
    ///                 Multiple = false,
    ///                 Required = false,
    ///                 Sensitive = false,
    ///                 Description = "",
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
    /// $ pulumi import splight:index/component:Component [options] splight_component.&lt;name&gt; &lt;component_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/component:Component")]
    public partial class Component : global::Pulumi.CustomResource
    {
        /// <summary>
        /// optinal description to add details of the resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        [Output("inputs")]
        public Output<ImmutableArray<Outputs.ComponentInput>> Inputs { get; private set; } = null!;

        /// <summary>
        /// the name of the component to be created
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [NAME-VERSION] the version of the hub component
        /// </summary>
        [Output("version")]
        public Output<string> Version { get; private set; } = null!;


        /// <summary>
        /// Create a Component resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Component(string name, ComponentArgs args, CustomResourceOptions? options = null)
            : base("splight:index/component:Component", name, args ?? new ComponentArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Component(string name, Input<string> id, ComponentState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/component:Component", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Component resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Component Get(string name, Input<string> id, ComponentState? state = null, CustomResourceOptions? options = null)
        {
            return new Component(name, id, state, options);
        }
    }

    public sealed class ComponentArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// optinal description to add details of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputs")]
        private InputList<Inputs.ComponentInputArgs>? _inputs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ComponentInputArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.ComponentInputArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// the name of the component to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [NAME-VERSION] the version of the hub component
        /// </summary>
        [Input("version", required: true)]
        public Input<string> Version { get; set; } = null!;

        public ComponentArgs()
        {
        }
        public static new ComponentArgs Empty => new ComponentArgs();
    }

    public sealed class ComponentState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// optinal description to add details of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("inputs")]
        private InputList<Inputs.ComponentInputGetArgs>? _inputs;

        /// <summary>
        /// static config parameters of the routine
        /// </summary>
        public InputList<Inputs.ComponentInputGetArgs> Inputs
        {
            get => _inputs ?? (_inputs = new InputList<Inputs.ComponentInputGetArgs>());
            set => _inputs = value;
        }

        /// <summary>
        /// the name of the component to be created
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [NAME-VERSION] the version of the hub component
        /// </summary>
        [Input("version")]
        public Input<string>? Version { get; set; }

        public ComponentState()
        {
        }
        public static new ComponentState Empty => new ComponentState();
    }
}
