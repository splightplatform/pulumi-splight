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

    public sealed class ActionSetpointAttributeArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// attribute id
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// attribute name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public ActionSetpointAttributeArgs()
        {
        }
        public static new ActionSetpointAttributeArgs Empty => new ActionSetpointAttributeArgs();
    }
}
