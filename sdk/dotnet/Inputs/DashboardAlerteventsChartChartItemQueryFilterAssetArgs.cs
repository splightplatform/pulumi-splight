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

    public sealed class DashboardAlerteventsChartChartItemQueryFilterAssetArgs : global::Pulumi.ResourceArgs
    {
        /// <summary>
        /// Id of the resource
        /// </summary>
        [Input("id")]
        public Input<string>? Id { get; set; }

        /// <summary>
        /// name of the resource
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        public DashboardAlerteventsChartChartItemQueryFilterAssetArgs()
        {
        }
        public static new DashboardAlerteventsChartChartItemQueryFilterAssetArgs Empty => new DashboardAlerteventsChartChartItemQueryFilterAssetArgs();
    }
}
