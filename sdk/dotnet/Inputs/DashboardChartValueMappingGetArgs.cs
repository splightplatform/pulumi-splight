// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splight.Inputs
{

    public sealed class DashboardChartValueMappingGetArgs : global::Pulumi.ResourceArgs
    {
        [Input("displayText", required: true)]
        public Input<string> DisplayText { get; set; } = null!;

        [Input("matchValue", required: true)]
        public Input<string> MatchValue { get; set; } = null!;

        [Input("order", required: true)]
        public Input<int> Order { get; set; } = null!;

        [Input("type", required: true)]
        public Input<string> Type { get; set; } = null!;

        public DashboardChartValueMappingGetArgs()
        {
        }
        public static new DashboardChartValueMappingGetArgs Empty => new DashboardChartValueMappingGetArgs();
    }
}
