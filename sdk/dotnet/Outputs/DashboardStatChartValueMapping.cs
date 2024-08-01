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
    public sealed class DashboardStatChartValueMapping
    {
        public readonly string DisplayText;
        public readonly string MatchValue;
        public readonly int Order;
        public readonly string Type;

        [OutputConstructor]
        private DashboardStatChartValueMapping(
            string displayText,

            string matchValue,

            int order,

            string type)
        {
            DisplayText = displayText;
            MatchValue = matchValue;
            Order = order;
            Type = type;
        }
    }
}
