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

    public sealed class ComponentRoutineOutputValueArgs : global::Pulumi.ResourceArgs
    {
        [Input("asset", required: true)]
        public Input<string> Asset { get; set; } = null!;

        [Input("attribute", required: true)]
        public Input<string> Attribute { get; set; } = null!;

        public ComponentRoutineOutputValueArgs()
        {
        }
        public static new ComponentRoutineOutputValueArgs Empty => new ComponentRoutineOutputValueArgs();
    }
}
