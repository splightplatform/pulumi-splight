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
    public sealed class DashboardActionlistChartThreshold
    {
        public readonly string Color;
        public readonly string DisplayText;
        public readonly double Value;

        [OutputConstructor]
        private DashboardActionlistChartThreshold(
            string color,

            string displayText,

            double value)
        {
            Color = color;
            DisplayText = displayText;
            Value = value;
        }
    }
}