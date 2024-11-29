// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Splight.Splight.Outputs
{

    [OutputType]
    public sealed class TransformerActivePowerHv
    {
        /// <summary>
        /// reference to the asset to be linked to
        /// </summary>
        public readonly string? Asset;
        /// <summary>
        /// id of the resource
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// name of the resource
        /// </summary>
        public readonly string? Name;
        /// <summary>
        /// [String|Boolean|Number] type of the data to be ingested in this attribute
        /// </summary>
        public readonly string? Type;
        /// <summary>
        /// unit of measure
        /// </summary>
        public readonly string? Unit;

        [OutputConstructor]
        private TransformerActivePowerHv(
            string? asset,

            string? id,

            string? name,

            string? type,

            string? unit)
        {
            Asset = asset;
            Id = id;
            Name = name;
            Type = type;
            Unit = unit;
        }
    }
}
