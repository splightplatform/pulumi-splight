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

    public sealed class ComponentInputGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("description")]
        public Input<string>? Description { get; set; }

        [Input("multiple")]
        public Input<bool>? Multiple { get; set; }

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("required")]
        public Input<bool>? Required { get; set; }

        [Input("sensitive")]
        public Input<bool>? Sensitive { get; set; }

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value")]
        public Input<string>? Value { get; set; }

        public ComponentInputGetArgs()
        {
        }
        public static new ComponentInputGetArgs Empty => new ComponentInputGetArgs();
    }
}
