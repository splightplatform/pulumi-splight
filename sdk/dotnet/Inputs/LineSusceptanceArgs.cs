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

    public sealed class LineSusceptanceArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// reference to the asset to be linked to
        /// </summary>
        [Input("asset")]
        public Input<string>? Asset { get; set; }

        /// <summary>
        /// id of the resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// [String|Boolean|Number] type of the data to be ingested in this attribute
        /// </summary>
        [Input("type")]
        public Input<string>? Type { get; set; }

        /// <summary>
        /// unit of measure
        /// </summary>
        [Input("unit")]
        public Input<string>? Unit { get; set; }

        /// <summary>
        /// metadata value
        /// </summary>
        [Input("value")]
        public Input<string>? Value { get; set; }

        public LineSusceptanceArgs()
        {
        }
        public static new LineSusceptanceArgs Empty => new LineSusceptanceArgs();
    }
}
