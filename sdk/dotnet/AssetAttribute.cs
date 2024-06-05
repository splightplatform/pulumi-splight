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
    /// using Pulumi;
    /// using Splight = Pulumi.Splight;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var assetTestFunctionAttribute = new Splight.AssetAttribute("assetTestFunctionAttribute", new()
    ///     {
    ///         Asset = "1234-1234-1234-1234",
    ///         Type = "Number",
    ///         Unit = "meters",
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import splight:index/assetAttribute:AssetAttribute [options] splight_asset_attribute.&lt;name&gt; &lt;asset_attribute_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/assetAttribute:AssetAttribute")]
    public partial class AssetAttribute : global::Pulumi.CustomResource
    {
        /// <summary>
        /// reference to the asset to be linked to
        /// </summary>
        [Output("asset")]
        public Output<string> Asset { get; private set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// [string|boolean|number] type of the data to be ingested in this attribute
        /// </summary>
        [Output("type")]
        public Output<string> Type { get; private set; } = null!;

        /// <summary>
        /// optional reference to the unit of the measure
        /// </summary>
        [Output("unit")]
        public Output<string?> Unit { get; private set; } = null!;


        /// <summary>
        /// Create a AssetAttribute resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public AssetAttribute(string name, AssetAttributeArgs args, CustomResourceOptions? options = null)
            : base("splight:index/assetAttribute:AssetAttribute", name, args ?? new AssetAttributeArgs(), MakeResourceOptions(options, ""))
        {
        }

        private AssetAttribute(string name, Input<string> id, AssetAttributeState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/assetAttribute:AssetAttribute", name, state, MakeResourceOptions(options, id))
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
        /// Get an existing AssetAttribute resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static AssetAttribute Get(string name, Input<string> id, AssetAttributeState? state = null, CustomResourceOptions? options = null)
        {
            return new AssetAttribute(name, id, state, options);
        }
    }

    public sealed class AssetAttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// reference to the asset to be linked to
        /// </summary>
        [Input("asset", required: true)]
        public Input<string> Asset { get; set; } = null!;

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [string|boolean|number] type of the data to be ingested in this attribute
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        /// <summary>
        /// optional reference to the unit of the measure
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public AssetAttributeArgs()
        {
        }
        public static new AssetAttributeArgs Empty => new AssetAttributeArgs();
    }

    public sealed class AssetAttributeState : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// reference to the asset to be linked to
        /// </summary>
        [Input("asset")]
        public Input<string>? Asset { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [string|boolean|number] type of the data to be ingested in this attribute
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// optional reference to the unit of the measure
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        public AssetAttributeState()
        {
        }
        public static new AssetAttributeState Empty => new AssetAttributeState();
    }
}