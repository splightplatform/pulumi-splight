// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;

namespace Pulumi.Splight.Inputs
{

    public sealed class DashboardChartThresholdArgs : global::Pulumi.ResourceArgs
    {
        [Input("color", required: true)]
        public Input<string> Color { get; set; } = null!;

        [Input("displayText", required: true)]
        public Input<string> DisplayText { get; set; } = null!;

        [Input("value", required: true)]
        public Input<double> Value { get; set; } = null!;

        public DashboardChartThresholdArgs()
        {
        }
        public static new DashboardChartThresholdArgs Empty => new DashboardChartThresholdArgs();
    }
}
