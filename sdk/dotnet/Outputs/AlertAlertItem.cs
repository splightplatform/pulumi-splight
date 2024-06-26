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
    public sealed class AlertAlertItem
    {
        public readonly string ExpressionPlain;
        public readonly string QueryPlain;
        public readonly string RefId;
        public readonly string Type;

        [OutputConstructor]
        private AlertAlertItem(
            string expressionPlain,

            string queryPlain,

            string refId,

            string type)
        {
            ExpressionPlain = expressionPlain;
            QueryPlain = queryPlain;
            RefId = refId;
            Type = type;
        }
    }
}
