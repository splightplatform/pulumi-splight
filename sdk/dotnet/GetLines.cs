// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

using System;
using System.Collections.Generic;
using System.Collections.Immutable;
using System.Threading.Tasks;
using Pulumi.Serialization;
using Pulumi;

namespace Splight.Splight
{
    public static class GetLines
    {
        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Splight = Pulumi.Splight;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var lines = Splight.GetLines.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetLinesResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetLinesResult>("splight:index/getLines:getLines", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// ## Example Usage
        /// 
        /// ```csharp
        /// using System.Collections.Generic;
        /// using System.Linq;
        /// using Pulumi;
        /// using Splight = Pulumi.Splight;
        /// 
        /// return await Deployment.RunAsync(() =&gt; 
        /// {
        ///     var lines = Splight.GetLines.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetLinesResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetLinesResult>("splight:index/getLines:getLines", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetLinesResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetLinesTagResult> Tags;

        [OutputConstructor]
        private GetLinesResult(
            string id,

            ImmutableArray<Outputs.GetLinesTagResult> tags)
        {
            Id = id;
            Tags = tags;
        }
    }
}
