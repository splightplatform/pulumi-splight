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
    public sealed class FunctionFunctionItem
    {
        /// <summary>
        /// how the expression is shown (i.e 'A * 2')
        /// </summary>
        public readonly string Expression;
        /// <summary>
        /// actual mongo query containing the expression
        /// </summary>
        public readonly string ExpressionPlain;
        /// <summary>
        /// ID of the function item
        /// </summary>
        public readonly string? Id;
        /// <summary>
        /// Asset/Attribute filter
        /// </summary>
        public readonly Outputs.FunctionFunctionItemQueryFilterAsset QueryFilterAsset;
        /// <summary>
        /// Asset/Attribute filter
        /// </summary>
        public readonly Outputs.FunctionFunctionItemQueryFilterAttribute QueryFilterAttribute;
        /// <summary>
        /// function used to aggregate data
        /// </summary>
        public readonly string QueryGroupFunction;
        /// <summary>
        /// time window to apply the aggregation
        /// </summary>
        public readonly string QueryGroupUnit;
        /// <summary>
        /// actual mongo query
        /// </summary>
        public readonly string QueryPlain;
        /// <summary>
        /// identifier of the variable (i.e 'A')
        /// </summary>
        public readonly string RefId;
        /// <summary>
        /// either QUERY or EXPRESSION
        /// </summary>
        public readonly string Type;

        [OutputConstructor]
        private FunctionFunctionItem(
            string expression,

            string expressionPlain,

            string? id,

            Outputs.FunctionFunctionItemQueryFilterAsset queryFilterAsset,

            Outputs.FunctionFunctionItemQueryFilterAttribute queryFilterAttribute,

            string queryGroupFunction,

            string queryGroupUnit,

            string queryPlain,

            string refId,

            string type)
        {
            Expression = expression;
            ExpressionPlain = expressionPlain;
            Id = id;
            QueryFilterAsset = queryFilterAsset;
            QueryFilterAttribute = queryFilterAttribute;
            QueryGroupFunction = queryGroupFunction;
            QueryGroupUnit = queryGroupUnit;
            QueryPlain = queryPlain;
            RefId = refId;
            Type = type;
        }
    }
}
