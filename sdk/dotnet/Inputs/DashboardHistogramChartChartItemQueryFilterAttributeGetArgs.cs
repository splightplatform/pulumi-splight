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

    public sealed class DashboardHistogramChartChartItemQueryFilterAttributeGetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// ID of the resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DashboardHistogramChartChartItemQueryFilterAttributeGetArgs()
        {
        }
        public static new DashboardHistogramChartChartItemQueryFilterAttributeGetArgs Empty => new DashboardHistogramChartChartItemQueryFilterAttributeGetArgs();
    }
}
