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
    public sealed class DashboardBarChartChartItemQueryFilterAsset
    {
        /// <summary>
        /// ID of the resource
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// name of the resource
        /// </summary>
        public readonly string? Name;

        [OutputConstructor]
        private DashboardBarChartChartItemQueryFilterAsset(
            string? id,

            string? name)
        {
            Id = id;
            Name = name;
        }
    }
}
