// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Splight.Splight.Outputs
{

    [OutputType]
    public sealed class DashboardTimeseriesChartChartItem
    {
        public readonly string Color;
        public readonly string ExpressionPlain;
        public readonly bool? Hidden;
        public readonly string? Label;
        /// <summary>
        /// Asset filter
        /// </summary>
        public readonly Outputs.DashboardTimeseriesChartChartItemQueryFilterAsset QueryFilterAsset;
        /// <summary>
        /// Attribute filter
        /// </summary>
        public readonly Outputs.DashboardTimeseriesChartChartItemQueryFilterAttribute QueryFilterAttribute;
        public readonly string? QueryGroupFunction;
        public readonly string? QueryGroupUnit;
        public readonly int? QueryLimit;
        public readonly string QueryPlain;
        public readonly int? QuerySortDirection;
        public readonly string RefId;
        public readonly string Type;

        [OutputConstructor]
        private DashboardTimeseriesChartChartItem(
            string color,

            string expressionPlain,

            bool? hidden,

            string? label,

            Outputs.DashboardTimeseriesChartChartItemQueryFilterAsset queryFilterAsset,

            Outputs.DashboardTimeseriesChartChartItemQueryFilterAttribute queryFilterAttribute,

            string? queryGroupFunction,

            string? queryGroupUnit,

            int? queryLimit,

            string queryPlain,

            int? querySortDirection,

            string refId,

            string type)
        {
            Color = color;
            ExpressionPlain = expressionPlain;
            Hidden = hidden;
            Label = label;
            QueryFilterAsset = queryFilterAsset;
            QueryFilterAttribute = queryFilterAttribute;
            QueryGroupFunction = queryGroupFunction;
            QueryGroupUnit = queryGroupUnit;
            QueryLimit = queryLimit;
            QueryPlain = queryPlain;
            QuerySortDirection = querySortDirection;
            RefId = refId;
            Type = type;
        }
    }
}
