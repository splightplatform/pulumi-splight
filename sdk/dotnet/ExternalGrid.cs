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
    /// $ pulumi import splight:index/externalGrid:ExternalGrid [options] splight_external_grid.&lt;name&gt; &lt;external_grid_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/externalGrid:ExternalGrid")]
    public partial class ExternalGrid : global::Pulumi.CustomResource
    {
        /// <summary>
        /// id of the related Bus object
        /// </summary>
        [Output("bus")]
        public Output<string?> Bus { get; private set; } = null!;

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
        /// id of the related Grid object
        /// </summary>
        [Output("grid")]
        public Output<string?> Grid { get; private set; } = null!;

        /// <summary>
        /// kind of the resource
        /// </summary>
        [Output("kinds")]
        public Output<ImmutableArray<Outputs.ExternalGridKind>> Kinds { get; private set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.ExternalGridTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// timezone that overrides location-based timezone of the resource
        /// </summary>
        [Output("timezone")]
        public Output<string> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a ExternalGrid resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public ExternalGrid(string name, ExternalGridArgs? args = null, CustomResourceOptions? options = null)
            : base("splight:index/externalGrid:ExternalGrid", name, args ?? new ExternalGridArgs(), MakeResourceOptions(options, ""))
        {
        }

        private ExternalGrid(string name, Input<string> id, ExternalGridState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/externalGrid:ExternalGrid", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing ExternalGrid resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static ExternalGrid Get(string name, Input<string> id, ExternalGridState? state = null, CustomResourceOptions? options = null)
        {
            return new ExternalGrid(name, id, state, options);
        }
    }

    public sealed class ExternalGridArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// id of the related Bus object
        /// </summary>
        [Input("bus")]
        public Input<string>? Bus { get; set; }

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
        /// id of the related Grid object
        /// </summary>
        [Input("grid")]
        public Input<string>? Grid { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.ExternalGridTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.ExternalGridTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ExternalGridTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// timezone that overrides location-based timezone of the resource
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public ExternalGridArgs()
        {
        }
        public static new ExternalGridArgs Empty => new ExternalGridArgs();
    }

    public sealed class ExternalGridState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// id of the related Bus object
        /// </summary>
        [Input("bus")]
        public Input<string>? Bus { get; set; }

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
        /// id of the related Grid object
        /// </summary>
        [Input("grid")]
        public Input<string>? Grid { get; set; }

        [Input("kinds")]
        private InputList<Inputs.ExternalGridKindGetArgs>? _kinds;

        /// <summary>
        /// kind of the resource
        /// </summary>
        public InputList<Inputs.ExternalGridKindGetArgs> Kinds
        {
            get => _kinds ?? (_kinds = new InputList<Inputs.ExternalGridKindGetArgs>());
            set => _kinds = value;
        }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.ExternalGridTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.ExternalGridTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.ExternalGridTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// timezone that overrides location-based timezone of the resource
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public ExternalGridState()
        {
        }
        public static new ExternalGridState Empty => new ExternalGridState();
    }
}
