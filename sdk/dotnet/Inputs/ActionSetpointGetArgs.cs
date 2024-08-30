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

    public sealed class ActionSetpointGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// the target attribute of the setpoint which should also be an attribute of the specified asset
        /// </summary>
        [Input("attribute", required: true)]
        public Input<Inputs.ActionSetpointAttributeGetArgs> Attribute { get; set; } = null!;

        /// <summary>
        /// setpoint ID
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// setpoint name
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// JSON encoded scalar value
        /// </summary>
        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ActionSetpointGetArgs()
        {
        }
        public static new ActionSetpointGetArgs Empty => new ActionSetpointGetArgs();
    }
}