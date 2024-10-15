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
    /// $ pulumi import splight:index/grid:Grid [options] splight_grid.&lt;name&gt; &lt;grid_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/grid:Grid")]
    public partial class Grid : global::Pulumi.CustomResource
    {
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
        public Output<ImmutableArray<Outputs.GridKind>> Kinds { get; private set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.GridTag>> Tags { get; private set; } = null!;


        /// <summary>
        /// Create a Grid resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Grid(string name, GridArgs? args = null, CustomResourceOptions? options = null)
            : base("splight:index/grid:Grid", name, args ?? new GridArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Grid(string name, Input<string> id, GridState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/grid:Grid", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Grid resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Grid Get(string name, Input<string> id, GridState? state = null, CustomResourceOptions? options = null)
        {
            return new Grid(name, id, state, options);
        }
    }

    public sealed class GridArgs : global::Pulumi.ResourceArgs
    {
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
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.GridTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.GridTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GridTagArgs>());
            set => _tags = value;
        }

        public GridArgs()
        {
        }
        public static new GridArgs Empty => new GridArgs();
    }

    public sealed class GridState : global::Pulumi.ResourceArgs
    {
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
        private InputList<Inputs.GridKindGetArgs>? _kinds;

        /// <summary>
        /// kind of the resource
        /// </summary>
        public InputList<Inputs.GridKindGetArgs> Kinds
        {
            get => _kinds ?? (_kinds = new InputList<Inputs.GridKindGetArgs>());
            set => _kinds = value;
        }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.GridTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.GridTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.GridTagGetArgs>());
            set => _tags = value;
        }

        public GridState()
        {
        }
        public static new GridState Empty => new GridState();
    }
}
