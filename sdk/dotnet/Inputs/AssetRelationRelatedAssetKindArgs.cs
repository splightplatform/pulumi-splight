// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Splight.Splight.Inputs
{

    public sealed class AssetRelationRelatedAssetKindArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// kind id
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// kind name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public AssetRelationRelatedAssetKindArgs()
        {
        }
        public static new AssetRelationRelatedAssetKindArgs Empty => new AssetRelationRelatedAssetKindArgs();
    }
}
