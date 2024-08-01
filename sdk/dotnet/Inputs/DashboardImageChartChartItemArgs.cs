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

    public sealed class DashboardImageChartChartItemArgs : global::Pulumi.ResourceArgs
    {
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        [Input("expressionPlain", required: true)]
        public Input<string> ExpressionPlain { get; set; } = null!;

        [Input("hidden")]
        public Input<bool>? Hidden { get; set; }

        [Input("label")]
        public Input<string>? Label { get; set; }

        /// <summary>
        /// Asset/Attribute filter
        /// </summary>
        [Input("queryFilterAsset", required: true)]
        public Input<Inputs.DashboardImageChartChartItemQueryFilterAssetArgs> QueryFilterAsset { get; set; } = null!;

        /// <summary>
        /// Asset/Attribute filter
        /// </summary>
        [Input("queryFilterAttribute", required: true)]
        public Input<Inputs.DashboardImageChartChartItemQueryFilterAttributeArgs> QueryFilterAttribute { get; set; } = null!;

        [Input("queryGroupFunction")]
        public Input<string>? QueryGroupFunction { get; set; }

        [Input("queryGroupUnit")]
        public Input<string>? QueryGroupUnit { get; set; }

        [Input("queryLimit")]
        public Input<int>? QueryLimit { get; set; }

        [Input("queryPlain", required: true)]
        public Input<string> QueryPlain { get; set; } = null!;

        [Input("querySortDirection")]
        public Input<int>? QuerySortDirection { get; set; }

        [Input("refId", required: true)]
        public Input<string> RefId { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DashboardImageChartChartItemArgs()
        {
        }
        public static new DashboardImageChartChartItemArgs Empty => new DashboardImageChartChartItemArgs();
    }
}
