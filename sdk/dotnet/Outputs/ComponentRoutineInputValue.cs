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
    public sealed class ComponentRoutineInputValue
    {
        public readonly string Asset;
        public readonly string Attribute;

        [OutputConstructor]
        private ComponentRoutineInputValue(
            string asset,

            string attribute)
        {
            Asset = asset;
            Attribute = attribute;
        }
    }
}
