// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splight.Inputs
{

    public sealed class ComponentRoutineOutputGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("multiple", required: true)]
        public Input<bool> Multiple { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("required", required: true)]
        public Input<bool> Required { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value", required: true)]
        public Input<Inputs.ComponentRoutineOutputValueGetArgs> Value { get; set; } = null!;

        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        public ComponentRoutineOutputGetArgs()
        {
        }
        public static new ComponentRoutineOutputGetArgs Empty => new ComponentRoutineOutputGetArgs();
    }
}
