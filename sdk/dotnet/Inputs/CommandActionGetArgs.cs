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

    public sealed class CommandActionGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// asset associated with the action (to be deprecated)
        /// </summary>
        [Input("asset", required: true)]
        public Input<Inputs.CommandActionAssetGetArgs> Asset { get; set; } = null!;

        /// <summary>
        /// action ID
        /// </summary>
        [Input("id", required: true)]
        public Input<string> Id { get; set; } = null!;

        /// <summary>
        /// setpoint name
        /// </summary>
        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        public CommandActionGetArgs()
        {
        }
        public static new CommandActionGetArgs Empty => new CommandActionGetArgs();
    }
}