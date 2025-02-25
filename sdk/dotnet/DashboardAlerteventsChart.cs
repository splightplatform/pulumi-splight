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
    /// <summary>
    /// ## Example Usage
    /// 
    /// ```csharp
    /// using System.Collections.Generic;
    /// using System.Linq;
    /// using System.Text.Json;
    /// using Pulumi;
    /// using Splight = Splight.Splight;
    /// 
    /// return await Deployment.RunAsync(() =&gt; 
    /// {
    ///     var assetTest = new Splight.Asset("assetTest", new()
    ///     {
    ///         Description = "Created with Terraform",
    ///         Timezone = "America/Los_Angeles",
    ///         Geometry = JsonSerializer.Serialize(new Dictionary&lt;string, object?&gt;
    ///         {
    ///             ["type"] = "GeometryCollection",
    ///             ["geometries"] = new[]
    ///             {
    ///             },
    ///         }),
    ///     });
    /// 
    ///     var attributeTest1 = new Splight.AssetAttribute("attributeTest1", new()
    ///     {
    ///         Type = "Number",
    ///         Unit = "meters",
    ///         Asset = assetTest.Id,
    ///     });
    /// 
    ///     var attributeTest2 = new Splight.AssetAttribute("attributeTest2", new()
    ///     {
    ///         Type = "Number",
    ///         Unit = "seconds",
    ///         Asset = assetTest.Id,
    ///     });
    /// 
    ///     var dashboardTest = new Splight.Dashboard("dashboardTest", new()
    ///     {
    ///         RelatedAssets = new[] {},
    ///     });
    /// 
    ///     var dashboardTabTest = new Splight.DashboardTab("dashboardTabTest", new()
    ///     {
    ///         Order = 0,
    ///         Dashboard = dashboardTest.Id,
    ///     });
    /// 
    ///     var dashboardChartTest = new Splight.DashboardAlerteventsChart("dashboardChartTest", new()
    ///     {
    ///         Tab = dashboardTabTest.Id,
    ///         TimestampGte = "now - 7d",
    ///         TimestampLte = "now",
    ///         Description = "Chart description",
    ///         MinHeight = 1,
    ///         MinWidth = 4,
    ///         DisplayTimeRange = true,
    ///         LabelsDisplay = true,
    ///         LabelsAggregation = "last",
    ///         LabelsPlacement = "bottom",
    ///         ShowBeyondData = true,
    ///         Height = 10,
    ///         Width = 20,
    ///         Collection = "default",
    ///         FilterName = "some name",
    ///         FilterOldStatuses = new[]
    ///         {
    ///             "warning",
    ///         },
    ///         FilterNewStatuses = new[]
    ///         {
    ///             "no_alert",
    ///             "warning",
    ///         },
    ///         ChartItems = new[]
    ///         {
    ///             new Splight.Inputs.DashboardAlerteventsChartChartItemArgs
    ///             {
    ///                 RefId = "A",
    ///                 Type = "QUERY",
    ///                 Color = "red",
    ///                 ExpressionPlain = "",
    ///                 QueryFilterAsset = new Splight.Inputs.DashboardAlerteventsChartChartItemQueryFilterAssetArgs
    ///                 {
    ///                     Id = assetTest.Id,
    ///                     Name = assetTest.Name,
    ///                 },
    ///                 QueryFilterAttribute = new Splight.Inputs.DashboardAlerteventsChartChartItemQueryFilterAttributeArgs
    ///                 {
    ///                     Id = attributeTest1.Id,
    ///                     Name = attributeTest1.Name,
    ///                 },
    ///                 QueryPlain = Output.JsonSerialize(Output.Create(new[]
    ///                 {
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["$match"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["asset"] = assetTest.Id,
    ///                             ["attribute"] = attributeTest1.Id,
    ///                         },
    ///                     },
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["$addFields"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["timestamp"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["$dateTrunc"] = new Dictionary&lt;string, object?&gt;
    ///                                 {
    ///                                     ["date"] = "$timestamp",
    ///                                     ["unit"] = "day",
    ///                                     ["binSize"] = 1,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["$group"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["_id"] = "$timestamp",
    ///                             ["value"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["$last"] = "$value",
    ///                             },
    ///                             ["timestamp"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["$last"] = "$timestamp",
    ///                             },
    ///                         },
    ///                     },
    ///                 })),
    ///             },
    ///             new Splight.Inputs.DashboardAlerteventsChartChartItemArgs
    ///             {
    ///                 RefId = "B",
    ///                 Color = "blue",
    ///                 Type = "QUERY",
    ///                 ExpressionPlain = "",
    ///                 QueryFilterAsset = new Splight.Inputs.DashboardAlerteventsChartChartItemQueryFilterAssetArgs
    ///                 {
    ///                     Id = assetTest.Id,
    ///                     Name = assetTest.Name,
    ///                 },
    ///                 QueryFilterAttribute = new Splight.Inputs.DashboardAlerteventsChartChartItemQueryFilterAttributeArgs
    ///                 {
    ///                     Id = attributeTest2.Id,
    ///                     Name = attributeTest2.Name,
    ///                 },
    ///                 QueryPlain = Output.JsonSerialize(Output.Create(new[]
    ///                 {
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["$match"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["asset"] = assetTest.Id,
    ///                             ["attribute"] = attributeTest2.Id,
    ///                         },
    ///                     },
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["$addFields"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["timestamp"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["$dateTrunc"] = new Dictionary&lt;string, object?&gt;
    ///                                 {
    ///                                     ["date"] = "$timestamp",
    ///                                     ["unit"] = "hour",
    ///                                     ["binSize"] = 1,
    ///                                 },
    ///                             },
    ///                         },
    ///                     },
    ///                     new Dictionary&lt;string, object?&gt;
    ///                     {
    ///                         ["$group"] = new Dictionary&lt;string, object?&gt;
    ///                         {
    ///                             ["_id"] = "$timestamp",
    ///                             ["value"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["$last"] = "$value",
    ///                             },
    ///                             ["timestamp"] = new Dictionary&lt;string, object?&gt;
    ///                             {
    ///                                 ["$last"] = "$timestamp",
    ///                             },
    ///                         },
    ///                     },
    ///                 })),
    ///             },
    ///         },
    ///         Thresholds = new[]
    ///         {
    ///             new Splight.Inputs.DashboardAlerteventsChartThresholdArgs
    ///             {
    ///                 Color = "#00edcf",
    ///                 DisplayText = "T1Test",
    ///                 Value = 13.1,
    ///             },
    ///         },
    ///         ValueMappings = new[]
    ///         {
    ///             new Splight.Inputs.DashboardAlerteventsChartValueMappingArgs
    ///             {
    ///                 DisplayText = "MODIFICADO",
    ///                 MatchValue = "123.3",
    ///                 Type = "exact_match",
    ///                 Order = 0,
    ///             },
    ///         },
    ///     });
    /// 
    /// });
    /// ```
    /// 
    /// ## Import
    /// 
    /// ```sh
    /// $ pulumi import splight:index/dashboardAlerteventsChart:DashboardAlerteventsChart [options] splight_dashboard_alertevents_chart.&lt;name&gt; &lt;dashboard_chart_id&gt;
    /// ```
    /// </summary>
    [SplightResourceType("splight:index/dashboardAlerteventsChart:DashboardAlerteventsChart")]
    public partial class DashboardAlerteventsChart : global::Pulumi.CustomResource
    {
        /// <summary>
        /// chart traces to be included
        /// </summary>
        [Output("chartItems")]
        public Output<ImmutableArray<Outputs.DashboardAlerteventsChartChartItem>> ChartItems { get; private set; } = null!;

        [Output("collection")]
        public Output<string?> Collection { get; private set; } = null!;

        /// <summary>
        /// chart description
        /// </summary>
        [Output("description")]
        public Output<string?> Description { get; private set; } = null!;

        /// <summary>
        /// whether to display the time range or not
        /// </summary>
        [Output("displayTimeRange")]
        public Output<bool?> DisplayTimeRange { get; private set; } = null!;

        /// <summary>
        /// filter name
        /// </summary>
        [Output("filterName")]
        public Output<string?> FilterName { get; private set; } = null!;

        /// <summary>
        /// filter new status
        /// </summary>
        [Output("filterNewStatuses")]
        public Output<ImmutableArray<string>> FilterNewStatuses { get; private set; } = null!;

        /// <summary>
        /// filter old status
        /// </summary>
        [Output("filterOldStatuses")]
        public Output<ImmutableArray<string>> FilterOldStatuses { get; private set; } = null!;

        /// <summary>
        /// chart height in px
        /// </summary>
        [Output("height")]
        public Output<int?> Height { get; private set; } = null!;

        /// <summary>
        /// [last|avg|...] aggregation
        /// </summary>
        [Output("labelsAggregation")]
        public Output<string?> LabelsAggregation { get; private set; } = null!;

        /// <summary>
        /// whether to display the labels or not
        /// </summary>
        [Output("labelsDisplay")]
        public Output<bool?> LabelsDisplay { get; private set; } = null!;

        /// <summary>
        /// [right|bottom] placement
        /// </summary>
        [Output("labelsPlacement")]
        public Output<string?> LabelsPlacement { get; private set; } = null!;

        /// <summary>
        /// minimum chart height
        /// </summary>
        [Output("minHeight")]
        public Output<int?> MinHeight { get; private set; } = null!;

        /// <summary>
        /// minimum chart width
        /// </summary>
        [Output("minWidth")]
        public Output<int?> MinWidth { get; private set; } = null!;

        /// <summary>
        /// name of the chart
        /// </summary>
        [Output("name")]
        public Output<string> Name { get; private set; } = null!;

        /// <summary>
        /// chart x position
        /// </summary>
        [Output("positionX")]
        public Output<int?> PositionX { get; private set; } = null!;

        /// <summary>
        /// chart y position
        /// </summary>
        [Output("positionY")]
        public Output<int?> PositionY { get; private set; } = null!;

        /// <summary>
        /// refresh interval
        /// </summary>
        [Output("refreshInterval")]
        public Output<string?> RefreshInterval { get; private set; } = null!;

        /// <summary>
        /// relative window time
        /// </summary>
        [Output("relativeWindowTime")]
        public Output<string?> RelativeWindowTime { get; private set; } = null!;

        /// <summary>
        /// whether to show data which is beyond timestamp_lte or not
        /// </summary>
        [Output("showBeyondData")]
        public Output<bool?> ShowBeyondData { get; private set; } = null!;

        /// <summary>
        /// id for the tab where to place the chart
        /// </summary>
        [Output("tab")]
        public Output<string> Tab { get; private set; } = null!;

        /// <summary>
        /// optional static lines to be added to the chart as references
        /// </summary>
        [Output("thresholds")]
        public Output<ImmutableArray<Outputs.DashboardAlerteventsChartThreshold>> Thresholds { get; private set; } = null!;

        /// <summary>
        /// date in isoformat or shortcut string where to end reading
        /// </summary>
        [Output("timestampGte")]
        public Output<string> TimestampGte { get; private set; } = null!;

        /// <summary>
        /// date in isoformat or shortcut string where to start reading
        /// </summary>
        [Output("timestampLte")]
        public Output<string> TimestampLte { get; private set; } = null!;

        /// <summary>
        /// chart timezone
        /// </summary>
        [Output("timezone")]
        public Output<string?> Timezone { get; private set; } = null!;

        /// <summary>
        /// optional mappings to transform data with rules
        /// </summary>
        [Output("valueMappings")]
        public Output<ImmutableArray<Outputs.DashboardAlerteventsChartValueMapping>> ValueMappings { get; private set; } = null!;

        /// <summary>
        /// chart width in cols (max 20)
        /// </summary>
        [Output("width")]
        public Output<int?> Width { get; private set; } = null!;


        /// <summary>
        /// Create a DashboardAlerteventsChart resource with the given unique name, arguments, and options.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resource</param>
        /// <param name="args">The arguments used to populate this resource's properties</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public DashboardAlerteventsChart(string name, DashboardAlerteventsChartArgs args, CustomResourceOptions? options = null)
            : base("splight:index/dashboardAlerteventsChart:DashboardAlerteventsChart", name, args ?? new DashboardAlerteventsChartArgs(), MakeResourceOptions(options, ""))
        {
        }

        private DashboardAlerteventsChart(string name, Input<string> id, DashboardAlerteventsChartState? state = null, CustomResourceOptions? options = null)
            : base("splight:index/dashboardAlerteventsChart:DashboardAlerteventsChart", name, state, MakeResourceOptions(options, id))
        {
        }

        private static CustomResourceOptions MakeResourceOptions(CustomResourceOptions? options, Input<string>? id)
        {
            var defaultOptions = new CustomResourceOptions
            {
                Version = Utilities.Version,
                PluginDownloadURL = "github://api.github.com/splightplatform",
            };
            var merged = CustomResourceOptions.Merge(defaultOptions, options);
            // Override the ID if one was specified for consistency with other language SDKs.
            merged.Id = id ?? merged.Id;
            return merged;
        }
        /// <summary>
        /// Get an existing DashboardAlerteventsChart resource's state with the given name, ID, and optional extra
        /// properties used to qualify the lookup.
        /// </summary>
        ///
        /// <param name="name">The unique name of the resulting resource.</param>
        /// <param name="id">The unique provider ID of the resource to lookup.</param>
        /// <param name="state">Any extra arguments used during the lookup.</param>
        /// <param name="options">A bag of options that control this resource's behavior</param>
        public static DashboardAlerteventsChart Get(string name, Input<string> id, DashboardAlerteventsChartState? state = null, CustomResourceOptions? options = null)
        {
            return new DashboardAlerteventsChart(name, id, state, options);
        }
    }

    public sealed class DashboardAlerteventsChartArgs : global::Pulumi.ResourceArgs
    {
        [Input("chartItems", required: true)]
        private InputList<Inputs.DashboardAlerteventsChartChartItemArgs>? _chartItems;

        /// <summary>
        /// chart traces to be included
        /// </summary>
        public InputList<Inputs.DashboardAlerteventsChartChartItemArgs> ChartItems
        {
            get => _chartItems ?? (_chartItems = new InputList<Inputs.DashboardAlerteventsChartChartItemArgs>());
            set => _chartItems = value;
        }

        [Input("collection")]
        public Input<string>? Collection { get; set; }

        /// <summary>
        /// chart description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// whether to display the time range or not
        /// </summary>
        [Input("displayTimeRange")]
        public Input<bool>? DisplayTimeRange { get; set; }

        /// <summary>
        /// filter name
        /// </summary>
        [Input("filterName")]
        public Input<string>? FilterName { get; set; }

        [Input("filterNewStatuses")]
        private InputList<string>? _filterNewStatuses;

        /// <summary>
        /// filter new status
        /// </summary>
        public InputList<string> FilterNewStatuses
        {
            get => _filterNewStatuses ?? (_filterNewStatuses = new InputList<string>());
            set => _filterNewStatuses = value;
        }

        [Input("filterOldStatuses")]
        private InputList<string>? _filterOldStatuses;

        /// <summary>
        /// filter old status
        /// </summary>
        public InputList<string> FilterOldStatuses
        {
            get => _filterOldStatuses ?? (_filterOldStatuses = new InputList<string>());
            set => _filterOldStatuses = value;
        }

        /// <summary>
        /// chart height in px
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// [last|avg|...] aggregation
        /// </summary>
        [Input("labelsAggregation")]
        public Input<string>? LabelsAggregation { get; set; }

        /// <summary>
        /// whether to display the labels or not
        /// </summary>
        [Input("labelsDisplay")]
        public Input<bool>? LabelsDisplay { get; set; }

        /// <summary>
        /// [right|bottom] placement
        /// </summary>
        [Input("labelsPlacement")]
        public Input<string>? LabelsPlacement { get; set; }

        /// <summary>
        /// minimum chart height
        /// </summary>
        [Input("minHeight")]
        public Input<int>? MinHeight { get; set; }

        /// <summary>
        /// minimum chart width
        /// </summary>
        [Input("minWidth")]
        public Input<int>? MinWidth { get; set; }

        /// <summary>
        /// name of the chart
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// chart x position
        /// </summary>
        [Input("positionX")]
        public Input<int>? PositionX { get; set; }

        /// <summary>
        /// chart y position
        /// </summary>
        [Input("positionY")]
        public Input<int>? PositionY { get; set; }

        /// <summary>
        /// refresh interval
        /// </summary>
        [Input("refreshInterval")]
        public Input<string>? RefreshInterval { get; set; }

        /// <summary>
        /// relative window time
        /// </summary>
        [Input("relativeWindowTime")]
        public Input<string>? RelativeWindowTime { get; set; }

        /// <summary>
        /// whether to show data which is beyond timestamp_lte or not
        /// </summary>
        [Input("showBeyondData")]
        public Input<bool>? ShowBeyondData { get; set; }

        /// <summary>
        /// id for the tab where to place the chart
        /// </summary>
        [Input("tab", required: true)]
        public Input<string> Tab { get; set; } = null!;

        [Input("thresholds")]
        private InputList<Inputs.DashboardAlerteventsChartThresholdArgs>? _thresholds;

        /// <summary>
        /// optional static lines to be added to the chart as references
        /// </summary>
        public InputList<Inputs.DashboardAlerteventsChartThresholdArgs> Thresholds
        {
            get => _thresholds ?? (_thresholds = new InputList<Inputs.DashboardAlerteventsChartThresholdArgs>());
            set => _thresholds = value;
        }

        /// <summary>
        /// date in isoformat or shortcut string where to end reading
        /// </summary>
        [Input("timestampGte", required: true)]
        public Input<string> TimestampGte { get; set; } = null!;

        /// <summary>
        /// date in isoformat or shortcut string where to start reading
        /// </summary>
        [Input("timestampLte", required: true)]
        public Input<string> TimestampLte { get; set; } = null!;

        /// <summary>
        /// chart timezone
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        [Input("valueMappings")]
        private InputList<Inputs.DashboardAlerteventsChartValueMappingArgs>? _valueMappings;

        /// <summary>
        /// optional mappings to transform data with rules
        /// </summary>
        public InputList<Inputs.DashboardAlerteventsChartValueMappingArgs> ValueMappings
        {
            get => _valueMappings ?? (_valueMappings = new InputList<Inputs.DashboardAlerteventsChartValueMappingArgs>());
            set => _valueMappings = value;
        }

        /// <summary>
        /// chart width in cols (max 20)
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public DashboardAlerteventsChartArgs()
        {
        }
        public static new DashboardAlerteventsChartArgs Empty => new DashboardAlerteventsChartArgs();
    }

    public sealed class DashboardAlerteventsChartState : global::Pulumi.ResourceArgs
    {
        [Input("chartItems")]
        private InputList<Inputs.DashboardAlerteventsChartChartItemGetArgs>? _chartItems;

        /// <summary>
        /// chart traces to be included
        /// </summary>
        public InputList<Inputs.DashboardAlerteventsChartChartItemGetArgs> ChartItems
        {
            get => _chartItems ?? (_chartItems = new InputList<Inputs.DashboardAlerteventsChartChartItemGetArgs>());
            set => _chartItems = value;
        }

        [Input("collection")]
        public Input<string>? Collection { get; set; }

        /// <summary>
        /// chart description
        /// </summary>
        [Input("description")]
        public Input<string>? Description { get; set; }

        /// <summary>
        /// whether to display the time range or not
        /// </summary>
        [Input("displayTimeRange")]
        public Input<bool>? DisplayTimeRange { get; set; }

        /// <summary>
        /// filter name
        /// </summary>
        [Input("filterName")]
        public Input<string>? FilterName { get; set; }

        [Input("filterNewStatuses")]
        private InputList<string>? _filterNewStatuses;

        /// <summary>
        /// filter new status
        /// </summary>
        public InputList<string> FilterNewStatuses
        {
            get => _filterNewStatuses ?? (_filterNewStatuses = new InputList<string>());
            set => _filterNewStatuses = value;
        }

        [Input("filterOldStatuses")]
        private InputList<string>? _filterOldStatuses;

        /// <summary>
        /// filter old status
        /// </summary>
        public InputList<string> FilterOldStatuses
        {
            get => _filterOldStatuses ?? (_filterOldStatuses = new InputList<string>());
            set => _filterOldStatuses = value;
        }

        /// <summary>
        /// chart height in px
        /// </summary>
        [Input("height")]
        public Input<int>? Height { get; set; }

        /// <summary>
        /// [last|avg|...] aggregation
        /// </summary>
        [Input("labelsAggregation")]
        public Input<string>? LabelsAggregation { get; set; }

        /// <summary>
        /// whether to display the labels or not
        /// </summary>
        [Input("labelsDisplay")]
        public Input<bool>? LabelsDisplay { get; set; }

        /// <summary>
        /// [right|bottom] placement
        /// </summary>
        [Input("labelsPlacement")]
        public Input<string>? LabelsPlacement { get; set; }

        /// <summary>
        /// minimum chart height
        /// </summary>
        [Input("minHeight")]
        public Input<int>? MinHeight { get; set; }

        /// <summary>
        /// minimum chart width
        /// </summary>
        [Input("minWidth")]
        public Input<int>? MinWidth { get; set; }

        /// <summary>
        /// name of the chart
        /// </summary>
        [Input("name")]
        public Input<string>? Name { get; set; }

        /// <summary>
        /// chart x position
        /// </summary>
        [Input("positionX")]
        public Input<int>? PositionX { get; set; }

        /// <summary>
        /// chart y position
        /// </summary>
        [Input("positionY")]
        public Input<int>? PositionY { get; set; }

        /// <summary>
        /// refresh interval
        /// </summary>
        [Input("refreshInterval")]
        public Input<string>? RefreshInterval { get; set; }

        /// <summary>
        /// relative window time
        /// </summary>
        [Input("relativeWindowTime")]
        public Input<string>? RelativeWindowTime { get; set; }

        /// <summary>
        /// whether to show data which is beyond timestamp_lte or not
        /// </summary>
        [Input("showBeyondData")]
        public Input<bool>? ShowBeyondData { get; set; }

        /// <summary>
        /// id for the tab where to place the chart
        /// </summary>
        [Input("tab")]
        public Input<string>? Tab { get; set; }

        [Input("thresholds")]
        private InputList<Inputs.DashboardAlerteventsChartThresholdGetArgs>? _thresholds;

        /// <summary>
        /// optional static lines to be added to the chart as references
        /// </summary>
        public InputList<Inputs.DashboardAlerteventsChartThresholdGetArgs> Thresholds
        {
            get => _thresholds ?? (_thresholds = new InputList<Inputs.DashboardAlerteventsChartThresholdGetArgs>());
            set => _thresholds = value;
        }

        /// <summary>
        /// date in isoformat or shortcut string where to end reading
        /// </summary>
        [Input("timestampGte")]
        public Input<string>? TimestampGte { get; set; }

        /// <summary>
        /// date in isoformat or shortcut string where to start reading
        /// </summary>
        [Input("timestampLte")]
        public Input<string>? TimestampLte { get; set; }

        /// <summary>
        /// chart timezone
        /// </summary>
        [Input("timezone")]
        public Input<string>? Timezone { get; set; }

        [Input("valueMappings")]
        private InputList<Inputs.DashboardAlerteventsChartValueMappingGetArgs>? _valueMappings;

        /// <summary>
        /// optional mappings to transform data with rules
        /// </summary>
        public InputList<Inputs.DashboardAlerteventsChartValueMappingGetArgs> ValueMappings
        {
            get => _valueMappings ?? (_valueMappings = new InputList<Inputs.DashboardAlerteventsChartValueMappingGetArgs>());
            set => _valueMappings = value;
        }

        /// <summary>
        /// chart width in cols (max 20)
        /// </summary>
        [Input("width")]
        public Input<int>? Width { get; set; }

        public DashboardAlerteventsChartState()
        {
        }
        public static new DashboardAlerteventsChartState Empty => new DashboardAlerteventsChartState();
    }
}
