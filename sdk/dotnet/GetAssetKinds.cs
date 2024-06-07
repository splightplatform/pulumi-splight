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
    public static class GetAssetKinds
    {
        /// <summary>
        /// Data source to fetch all asset kinds defined in the platform
        /// 
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
        ///     var kinds = Splight.GetAssetKinds.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Task<GetAssetKindsResult> InvokeAsync(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.InvokeAsync<GetAssetKindsResult>("splight:index/getAssetKinds:getAssetKinds", InvokeArgs.Empty, options.WithDefaults());

        /// <summary>
        /// Data source to fetch all asset kinds defined in the platform
        /// 
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
        ///     var kinds = Splight.GetAssetKinds.Invoke();
        /// 
        /// });
        /// ```
        /// </summary>
        public static Output<GetAssetKindsResult> Invoke(InvokeOptions? options = null)
            => global::Pulumi.Deployment.Instance.Invoke<GetAssetKindsResult>("splight:index/getAssetKinds:getAssetKinds", InvokeArgs.Empty, options.WithDefaults());
    }


    [OutputType]
    public sealed class GetAssetKindsResult
    {
        /// <summary>
        /// The provider-assigned unique ID for this managed resource.
        /// </summary>
        public readonly string Id;
        public readonly ImmutableArray<Outputs.GetAssetKindsKindResult> Kinds;

        [OutputConstructor]
        private GetAssetKindsResult(
            string id,

            ImmutableArray<Outputs.GetAssetKindsKindResult> kinds)
        {
            Id = id;
            Kinds = kinds;
        }
    }
}
