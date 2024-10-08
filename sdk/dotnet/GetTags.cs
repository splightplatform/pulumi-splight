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
    public static class GetTags
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
        ///     var tags = Splight.GetTags.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetTagsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetTagsResult>("splight:index/getTags:getTags", InvokeArgs.Empty, options.WithDefaults());

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
        ///     var tags = Splight.GetTags.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetTagsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetTagsResult>("splight:index/getTags:getTags", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetTagsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetTagsTagResult> Tags;

        [OutputConstructor]
        private GetTagsResult(
            string id,

            ImmutableArray<Outputs.GetTagsTagResult> tags)
        {
            Id = id;
            Tags = tags;
        }
    }
}
