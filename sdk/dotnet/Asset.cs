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
    /// $ pulumi import splight:index/asset:Asset [options] splight_asset.&lt;name&gt; &lt;asset_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/asset:Asset")]
    public partial class Asset : global::Pulumi.CustomResource
    {
        /// <summary>
        /// description of the resource
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// GeoJSON GeomtryCollection
        /// </summary>
        [Output("geometry")]
        public Output<string> Geometry { get; private set; } = null!;

        /// <summary>
        /// kind of the resource
        /// </summary>
        [Output("kind")]
        public Output<Outputs.AssetKind?> Kind { get; private set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// tags of the resource
        /// </summary>
        [Output("tags")]
        public Output<ImmutableArray<Outputs.AssetTag>> Tags { get; private set; } = null!;

        /// <summary>
        /// timezone of the resource (overriden by the location if set)
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;


        /// <summary>
        /// Create a Asset resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public Asset(string name, AssetArgs args, CustomResourceOptions? options = null)
            : base("splight:index/asset:Asset", name, args ?? new AssetArgs(), MakeResourceOptions(options, ""))
        {
        }

        private Asset(string name, Input<string> id, AssetState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/asset:Asset", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing Asset resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static Asset Get(string name, Input<string> id, AssetState? state = null, CustomResourceOptions? options = null)
        {
            return new Asset(name, id, state, options);
        }
    }

    public sealed class AssetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// GeoJSON GeomtryCollection
        /// </summary>
        [Input("geometry", required: true)]
        public Input<string> Geometry { get; set; } = null!;

        /// <summary>
        /// kind of the resource
        /// </summary>
        [Input("kind")]
        public Input<Inputs.AssetKindArgs>? Kind { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.AssetTagArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.AssetTagArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AssetTagArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// timezone of the resource (overriden by the location if set)
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public AssetArgs()
        {
        }
        public static new AssetArgs Empty => new AssetArgs();
    }

    public sealed class AssetState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// description of the resource
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// GeoJSON GeomtryCollection
        /// </summary>
        [Input("geometry")]
        public Input<string>? Geometry { get; set; }

        /// <summary>
        /// kind of the resource
        /// </summary>
        [Input("kind")]
        public Input<Inputs.AssetKindGetArgs>? Kind { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        [Input("tags")]
        private InputList<Inputs.AssetTagGetArgs>? _tags;

        /// <summary>
        /// tags of the resource
        /// </summary>
        public InputList<Inputs.AssetTagGetArgs> Tags
        {
            get => _tags ?? (_tags = new InputList<Inputs.AssetTagGetArgs>());
            set => _tags = value;
        }

        /// <summary>
        /// timezone of the resource (overriden by the location if set)
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        public AssetState()
        {
        }
        public static new AssetState Empty => new AssetState();
    }
}
