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

    public sealed class AlertAlertItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("expressionPlain", required: true)]
        public Input<string> ExpressionPlain { get; set; } = null!;

        [Input("queryPlain", required: true)]
        public Input<string> QueryPlain { get; set; } = null!;

        [Input("refId", required: true)]
        public Input<string> RefId { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AlertAlertItemArgs()
        {
        }
        public static new AlertAlertItemArgs Empty => new AlertAlertItemArgs();
    }
}
