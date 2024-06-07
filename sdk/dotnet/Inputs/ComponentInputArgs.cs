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

    public sealed class ComponentInputArgs : global::Pulumi.ResourceArgs
    {
        [Input("description", required: true)]
        public Input<string> Description { get; set; } = null!;

        [Input("multiple", required: true)]
        public Input<bool> Multiple { get; set; } = null!;

        [Input("name", required: true)]
        public Input<string> Name { get; set; } = null!;

        [Input("required", required: true)]
        public Input<bool> Required { get; set; } = null!;

        [Input("sensitive", required: true)]
        public Input<bool> Sensitive { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        [Input("value", required: true)]
        public Input<string> Value { get; set; } = null!;

        public ComponentInputArgs()
        {
        }
        public static new ComponentInputArgs Empty => new ComponentInputArgs();
    }
}
