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

    public sealed class FunctionFunctionItemGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("expressionPlain", required: true)]
        public Input<string> ExpressionPlain { get; set; } = null!;

        /// <summary>
        /// optional id
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        [Input("queryPlain", required: true)]
        public Input<string> QueryPlain { get; set; } = null!;

        [Input("refId", required: true)]
        public Input<string> RefId { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public FunctionFunctionItemGetArgs()
        {
        }
        public static new FunctionFunctionItemGetArgs Empty => new FunctionFunctionItemGetArgs();
    }
}
