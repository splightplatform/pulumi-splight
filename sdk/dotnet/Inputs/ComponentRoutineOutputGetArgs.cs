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

    public sealed class ComponentRoutineOutputGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("multiple")]
        public Input<bool>? Multiple { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("required")]
        public Input<bool>? Required { get; set; }

        [Input("type")]
        public Input<string>? Type { get; set; }

        [Input("valueType", required: true)]
        public Input<string> ValueType { get; set; } = null!;

        [Input("values")]
        private InputList<Inputs.ComponentRoutineOutputValueGetArgs>? _values;
        public InputList<Inputs.ComponentRoutineOutputValueGetArgs> Values
        {
            get => _values ?? (_values = new InputList<Inputs.ComponentRoutineOutputValueGetArgs>());
            set => _values = value;
        }

        public ComponentRoutineOutputGetArgs()
        {
        }
        public static new ComponentRoutineOutputGetArgs Empty => new ComponentRoutineOutputGetArgs();
    }
}
