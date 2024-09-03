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

    public sealed class AlertAlertItemGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// how the expression is shown (i.e 'A * 2')
        /// </summary>
        [Input("expression", required: true)]
        public Input<string> Expression { get; set; } = null!;

        /// <summary>
        /// actual mongo query containing the expression
        /// </summary>
        [Input("expressionPlain", required: true)]
        public Input<string> ExpressionPlain { get; set; } = null!;

        /// <summary>
        /// ID of the function item
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// Asset/Attribute filter
        /// </summary>
        [Input("queryFilterAsset", required: true)]
        public Input<Inputs.AlertAlertItemQueryFilterAssetGetArgs> QueryFilterAsset { get; set; } = null!;

        /// <summary>
        /// Asset/Attribute filter
        /// </summary>
        [Input("queryFilterAttribute", required: true)]
        public Input<Inputs.AlertAlertItemQueryFilterAttributeGetArgs> QueryFilterAttribute { get; set; } = null!;

        /// <summary>
        /// function used to aggregate data
        /// </summary>
        [Input("queryGroupFunction", required: true)]
        public Input<string> QueryGroupFunction { get; set; } = null!;

        /// <summary>
        /// time window to apply the aggregation
        /// </summary>
        [Input("queryGroupUnit", required: true)]
        public Input<string> QueryGroupUnit { get; set; } = null!;

        /// <summary>
        /// actual mongo query
        /// </summary>
        [Input("queryPlain", required: true)]
        public Input<string> QueryPlain { get; set; } = null!;

        /// <summary>
        /// identifier of the variable (i.e 'A')
        /// </summary>
        [Input("refId", required: true)]
        public Input<string> RefId { get; set; } = null!;

        /// <summary>
        /// either QUERY or EXPRESSION
        /// </summary>
        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public AlertAlertItemGetArgs()
        {
        }
        public static new AlertAlertItemGetArgs Empty => new AlertAlertItemGetArgs();
    }
}
