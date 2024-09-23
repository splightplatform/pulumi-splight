// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as inputs from "./types/input";
import * as outputs from "./types/output";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as splight from "@splightplatform/pulumi-splight";
 *
 * const assetTest = new splight.Asset("assetTest", {
 *     description: "Created with Terraform",
 *     geometry: JSON.stringify({
 *         type: "GeometryCollection",
 *         geometries: [],
 *     }),
 * });
 * const attributeTest1 = new splight.AssetAttribute("attributeTest1", {
 *     type: "Number",
 *     unit: "meters",
 *     asset: assetTest.id,
 * });
 * const attributeTest2 = new splight.AssetAttribute("attributeTest2", {
 *     type: "Number",
 *     unit: "seconds",
 *     asset: assetTest.id,
 * });
 * const dashboardTest = new splight.Dashboard("dashboardTest", {relatedAssets: []});
 * const dashboardTabTest = new splight.DashboardTab("dashboardTabTest", {
 *     order: 0,
 *     dashboard: dashboardTest.id,
 * });
 * const dashboardChartTest = new splight.DashboardBargaugeChart("dashboardChartTest", {
 *     tab: dashboardTabTest.id,
 *     timestampGte: "now - 7d",
 *     timestampLte: "now",
 *     description: "Chart description",
 *     minHeight: 1,
 *     minWidth: 4,
 *     displayTimeRange: true,
 *     labelsDisplay: true,
 *     labelsAggregation: "last",
 *     labelsPlacement: "bottom",
 *     showBeyondData: true,
 *     height: 10,
 *     width: 20,
 *     collection: "default",
 *     maxLimit: 100,
 *     numberOfDecimals: 2,
 *     orientation: "vertical",
 *     chartItems: [
 *         {
 *             refId: "A",
 *             type: "QUERY",
 *             color: "red",
 *             expressionPlain: "",
 *             queryFilterAsset: {
 *                 id: assetTest.id,
 *                 name: assetTest.name,
 *             },
 *             queryFilterAttribute: {
 *                 id: attributeTest1.id,
 *                 name: attributeTest1.name,
 *             },
 *             queryPlain: pulumi.jsonStringify([
 *                 {
 *                     $match: {
 *                         asset: assetTest.id,
 *                         attribute: attributeTest1.id,
 *                     },
 *                 },
 *                 {
 *                     $addFields: {
 *                         timestamp: {
 *                             $dateTrunc: {
 *                                 date: "$timestamp",
 *                                 unit: "day",
 *                                 binSize: 1,
 *                             },
 *                         },
 *                     },
 *                 },
 *                 {
 *                     $group: {
 *                         _id: "$timestamp",
 *                         value: {
 *                             $last: "$value",
 *                         },
 *                         timestamp: {
 *                             $last: "$timestamp",
 *                         },
 *                     },
 *                 },
 *             ]),
 *         },
 *         {
 *             refId: "B",
 *             color: "blue",
 *             type: "EXPRESSION",
 *             queryPlain: "",
 *             expressionPlain: JSON.stringify({
 *                 $function: {
 *                     body: "function ($A) { return $A/50 }",
 *                     args: ["$A"],
 *                     lang: "js",
 *                 },
 *             }),
 *             queryFilterAsset: {},
 *             queryFilterAttribute: {},
 *         },
 *     ],
 *     thresholds: [{
 *         color: "#00edcf",
 *         displayText: "T1Test",
 *         value: 13.1,
 *     }],
 *     valueMappings: [{
 *         displayText: "MODIFICADO",
 *         matchValue: "123.3",
 *         type: "exact_match",
 *         order: 0,
 *     }],
 * });
 * ```
 *
 * ## Import
 *
 * ```sh
 * $ pulumi import splight:index/dashboardBargaugeChart:DashboardBargaugeChart [options] splight_dashboard_bargauge_chart.<name> <dashboard_chart_id>
 * ```
 */
export class DashboardBargaugeChart extends pulumi.CustomResource {
    /**
     * Get an existing DashboardBargaugeChart resource's state with the given name, ID, and optional extra
     * properties used to qualify the lookup.
     *
     * @param name The _unique_ name of the resulting resource.
     * @param id The _unique_ provider ID of the resource to lookup.
     * @param state Any extra arguments used during the lookup.
     * @param opts Optional settings to control the behavior of the CustomResource.
     */
    public static get(name: string, id: pulumi.Input<pulumi.ID>, state?: DashboardBargaugeChartState, opts?: pulumi.CustomResourceOptions): DashboardBargaugeChart {
        return new DashboardBargaugeChart(name, <any>state, { ...opts, id: id });
    }

    /** @internal */
    public static readonly __pulumiType = 'splight:index/dashboardBargaugeChart:DashboardBargaugeChart';

    /**
     * Returns true if the given object is an instance of DashboardBargaugeChart.  This is designed to work even
     * when multiple copies of the Pulumi SDK have been loaded into the same process.
     */
    public static isInstance(obj: any): obj is DashboardBargaugeChart {
        if (obj === undefined || obj === null) {
            return false;
        }
        return obj['__pulumiType'] === DashboardBargaugeChart.__pulumiType;
    }

    /**
     * chart traces to be included
     */
    public readonly chartItems!: pulumi.Output<outputs.DashboardBargaugeChartChartItem[]>;
    public readonly collection!: pulumi.Output<string | undefined>;
    /**
     * chart description
     */
    public readonly description!: pulumi.Output<string | undefined>;
    /**
     * whether to display the time range or not
     */
    public readonly displayTimeRange!: pulumi.Output<boolean | undefined>;
    /**
     * chart height in px
     */
    public readonly height!: pulumi.Output<number | undefined>;
    /**
     * [last|avg|...] aggregation
     */
    public readonly labelsAggregation!: pulumi.Output<string | undefined>;
    /**
     * whether to display the labels or not
     */
    public readonly labelsDisplay!: pulumi.Output<boolean | undefined>;
    /**
     * [right|bottom] placement
     */
    public readonly labelsPlacement!: pulumi.Output<string | undefined>;
    /**
     * bar gauge max limit
     */
    public readonly maxLimit!: pulumi.Output<number | undefined>;
    /**
     * minimum chart height
     */
    public readonly minHeight!: pulumi.Output<number | undefined>;
    /**
     * minimum chart width
     */
    public readonly minWidth!: pulumi.Output<number | undefined>;
    /**
     * name of the chart
     */
    public readonly name!: pulumi.Output<string>;
    /**
     * number of decimals
     */
    public readonly numberOfDecimals!: pulumi.Output<number | undefined>;
    /**
     * orientation
     */
    public readonly orientation!: pulumi.Output<string | undefined>;
    /**
     * chart x position
     */
    public readonly positionX!: pulumi.Output<number | undefined>;
    /**
     * chart y position
     */
    public readonly positionY!: pulumi.Output<number | undefined>;
    /**
     * refresh interval
     */
    public readonly refreshInterval!: pulumi.Output<string | undefined>;
    /**
     * relative window time
     */
    public readonly relativeWindowTime!: pulumi.Output<string | undefined>;
    /**
     * whether to show data which is beyond timestampLte or not
     */
    public readonly showBeyondData!: pulumi.Output<boolean | undefined>;
    /**
     * id for the tab where to place the chart
     */
    public readonly tab!: pulumi.Output<string>;
    /**
     * optional static lines to be added to the chart as references
     */
    public readonly thresholds!: pulumi.Output<outputs.DashboardBargaugeChartThreshold[] | undefined>;
    /**
     * date in isoformat or shortcut string where to end reading
     */
    public readonly timestampGte!: pulumi.Output<string>;
    /**
     * date in isoformat or shortcut string where to start reading
     */
    public readonly timestampLte!: pulumi.Output<string>;
    /**
     * chart timezone
     */
    public readonly timezone!: pulumi.Output<string | undefined>;
    /**
     * optional mappings to transform data with rules
     */
    public readonly valueMappings!: pulumi.Output<outputs.DashboardBargaugeChartValueMapping[] | undefined>;
    /**
     * chart width in cols (max 20)
     */
    public readonly width!: pulumi.Output<number | undefined>;

    /**
     * Create a DashboardBargaugeChart resource with the given unique name, arguments, and options.
     *
     * @param name The _unique_ name of the resource.
     * @param args The arguments to use to populate this resource's properties.
     * @param opts A bag of options that control this resource's behavior.
     */
    constructor(name: string, args: DashboardBargaugeChartArgs, opts?: pulumi.CustomResourceOptions)
    constructor(name: string, argsOrState?: DashboardBargaugeChartArgs | DashboardBargaugeChartState, opts?: pulumi.CustomResourceOptions) {
        let resourceInputs: pulumi.Inputs = {};
        opts = opts || {};
        if (opts.id) {
            const state = argsOrState as DashboardBargaugeChartState | undefined;
            resourceInputs["chartItems"] = state ? state.chartItems : undefined;
            resourceInputs["collection"] = state ? state.collection : undefined;
            resourceInputs["description"] = state ? state.description : undefined;
            resourceInputs["displayTimeRange"] = state ? state.displayTimeRange : undefined;
            resourceInputs["height"] = state ? state.height : undefined;
            resourceInputs["labelsAggregation"] = state ? state.labelsAggregation : undefined;
            resourceInputs["labelsDisplay"] = state ? state.labelsDisplay : undefined;
            resourceInputs["labelsPlacement"] = state ? state.labelsPlacement : undefined;
            resourceInputs["maxLimit"] = state ? state.maxLimit : undefined;
            resourceInputs["minHeight"] = state ? state.minHeight : undefined;
            resourceInputs["minWidth"] = state ? state.minWidth : undefined;
            resourceInputs["name"] = state ? state.name : undefined;
            resourceInputs["numberOfDecimals"] = state ? state.numberOfDecimals : undefined;
            resourceInputs["orientation"] = state ? state.orientation : undefined;
            resourceInputs["positionX"] = state ? state.positionX : undefined;
            resourceInputs["positionY"] = state ? state.positionY : undefined;
            resourceInputs["refreshInterval"] = state ? state.refreshInterval : undefined;
            resourceInputs["relativeWindowTime"] = state ? state.relativeWindowTime : undefined;
            resourceInputs["showBeyondData"] = state ? state.showBeyondData : undefined;
            resourceInputs["tab"] = state ? state.tab : undefined;
            resourceInputs["thresholds"] = state ? state.thresholds : undefined;
            resourceInputs["timestampGte"] = state ? state.timestampGte : undefined;
            resourceInputs["timestampLte"] = state ? state.timestampLte : undefined;
            resourceInputs["timezone"] = state ? state.timezone : undefined;
            resourceInputs["valueMappings"] = state ? state.valueMappings : undefined;
            resourceInputs["width"] = state ? state.width : undefined;
        } else {
            const args = argsOrState as DashboardBargaugeChartArgs | undefined;
            if ((!args || args.chartItems === undefined) && !opts.urn) {
                throw new Error("Missing required property 'chartItems'");
            }
            if ((!args || args.tab === undefined) && !opts.urn) {
                throw new Error("Missing required property 'tab'");
            }
            if ((!args || args.timestampGte === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timestampGte'");
            }
            if ((!args || args.timestampLte === undefined) && !opts.urn) {
                throw new Error("Missing required property 'timestampLte'");
            }
            resourceInputs["chartItems"] = args ? args.chartItems : undefined;
            resourceInputs["collection"] = args ? args.collection : undefined;
            resourceInputs["description"] = args ? args.description : undefined;
            resourceInputs["displayTimeRange"] = args ? args.displayTimeRange : undefined;
            resourceInputs["height"] = args ? args.height : undefined;
            resourceInputs["labelsAggregation"] = args ? args.labelsAggregation : undefined;
            resourceInputs["labelsDisplay"] = args ? args.labelsDisplay : undefined;
            resourceInputs["labelsPlacement"] = args ? args.labelsPlacement : undefined;
            resourceInputs["maxLimit"] = args ? args.maxLimit : undefined;
            resourceInputs["minHeight"] = args ? args.minHeight : undefined;
            resourceInputs["minWidth"] = args ? args.minWidth : undefined;
            resourceInputs["name"] = args ? args.name : undefined;
            resourceInputs["numberOfDecimals"] = args ? args.numberOfDecimals : undefined;
            resourceInputs["orientation"] = args ? args.orientation : undefined;
            resourceInputs["positionX"] = args ? args.positionX : undefined;
            resourceInputs["positionY"] = args ? args.positionY : undefined;
            resourceInputs["refreshInterval"] = args ? args.refreshInterval : undefined;
            resourceInputs["relativeWindowTime"] = args ? args.relativeWindowTime : undefined;
            resourceInputs["showBeyondData"] = args ? args.showBeyondData : undefined;
            resourceInputs["tab"] = args ? args.tab : undefined;
            resourceInputs["thresholds"] = args ? args.thresholds : undefined;
            resourceInputs["timestampGte"] = args ? args.timestampGte : undefined;
            resourceInputs["timestampLte"] = args ? args.timestampLte : undefined;
            resourceInputs["timezone"] = args ? args.timezone : undefined;
            resourceInputs["valueMappings"] = args ? args.valueMappings : undefined;
            resourceInputs["width"] = args ? args.width : undefined;
        }
        opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
        super(DashboardBargaugeChart.__pulumiType, name, resourceInputs, opts);
    }
}

/**
 * Input properties used for looking up and filtering DashboardBargaugeChart resources.
 */
export interface DashboardBargaugeChartState {
    /**
     * chart traces to be included
     */
    chartItems?: pulumi.Input<pulumi.Input<inputs.DashboardBargaugeChartChartItem>[]>;
    collection?: pulumi.Input<string>;
    /**
     * chart description
     */
    description?: pulumi.Input<string>;
    /**
     * whether to display the time range or not
     */
    displayTimeRange?: pulumi.Input<boolean>;
    /**
     * chart height in px
     */
    height?: pulumi.Input<number>;
    /**
     * [last|avg|...] aggregation
     */
    labelsAggregation?: pulumi.Input<string>;
    /**
     * whether to display the labels or not
     */
    labelsDisplay?: pulumi.Input<boolean>;
    /**
     * [right|bottom] placement
     */
    labelsPlacement?: pulumi.Input<string>;
    /**
     * bar gauge max limit
     */
    maxLimit?: pulumi.Input<number>;
    /**
     * minimum chart height
     */
    minHeight?: pulumi.Input<number>;
    /**
     * minimum chart width
     */
    minWidth?: pulumi.Input<number>;
    /**
     * name of the chart
     */
    name?: pulumi.Input<string>;
    /**
     * number of decimals
     */
    numberOfDecimals?: pulumi.Input<number>;
    /**
     * orientation
     */
    orientation?: pulumi.Input<string>;
    /**
     * chart x position
     */
    positionX?: pulumi.Input<number>;
    /**
     * chart y position
     */
    positionY?: pulumi.Input<number>;
    /**
     * refresh interval
     */
    refreshInterval?: pulumi.Input<string>;
    /**
     * relative window time
     */
    relativeWindowTime?: pulumi.Input<string>;
    /**
     * whether to show data which is beyond timestampLte or not
     */
    showBeyondData?: pulumi.Input<boolean>;
    /**
     * id for the tab where to place the chart
     */
    tab?: pulumi.Input<string>;
    /**
     * optional static lines to be added to the chart as references
     */
    thresholds?: pulumi.Input<pulumi.Input<inputs.DashboardBargaugeChartThreshold>[]>;
    /**
     * date in isoformat or shortcut string where to end reading
     */
    timestampGte?: pulumi.Input<string>;
    /**
     * date in isoformat or shortcut string where to start reading
     */
    timestampLte?: pulumi.Input<string>;
    /**
     * chart timezone
     */
    timezone?: pulumi.Input<string>;
    /**
     * optional mappings to transform data with rules
     */
    valueMappings?: pulumi.Input<pulumi.Input<inputs.DashboardBargaugeChartValueMapping>[]>;
    /**
     * chart width in cols (max 20)
     */
    width?: pulumi.Input<number>;
}

/**
 * The set of arguments for constructing a DashboardBargaugeChart resource.
 */
export interface DashboardBargaugeChartArgs {
    /**
     * chart traces to be included
     */
    chartItems: pulumi.Input<pulumi.Input<inputs.DashboardBargaugeChartChartItem>[]>;
    collection?: pulumi.Input<string>;
    /**
     * chart description
     */
    description?: pulumi.Input<string>;
    /**
     * whether to display the time range or not
     */
    displayTimeRange?: pulumi.Input<boolean>;
    /**
     * chart height in px
     */
    height?: pulumi.Input<number>;
    /**
     * [last|avg|...] aggregation
     */
    labelsAggregation?: pulumi.Input<string>;
    /**
     * whether to display the labels or not
     */
    labelsDisplay?: pulumi.Input<boolean>;
    /**
     * [right|bottom] placement
     */
    labelsPlacement?: pulumi.Input<string>;
    /**
     * bar gauge max limit
     */
    maxLimit?: pulumi.Input<number>;
    /**
     * minimum chart height
     */
    minHeight?: pulumi.Input<number>;
    /**
     * minimum chart width
     */
    minWidth?: pulumi.Input<number>;
    /**
     * name of the chart
     */
    name?: pulumi.Input<string>;
    /**
     * number of decimals
     */
    numberOfDecimals?: pulumi.Input<number>;
    /**
     * orientation
     */
    orientation?: pulumi.Input<string>;
    /**
     * chart x position
     */
    positionX?: pulumi.Input<number>;
    /**
     * chart y position
     */
    positionY?: pulumi.Input<number>;
    /**
     * refresh interval
     */
    refreshInterval?: pulumi.Input<string>;
    /**
     * relative window time
     */
    relativeWindowTime?: pulumi.Input<string>;
    /**
     * whether to show data which is beyond timestampLte or not
     */
    showBeyondData?: pulumi.Input<boolean>;
    /**
     * id for the tab where to place the chart
     */
    tab: pulumi.Input<string>;
    /**
     * optional static lines to be added to the chart as references
     */
    thresholds?: pulumi.Input<pulumi.Input<inputs.DashboardBargaugeChartThreshold>[]>;
    /**
     * date in isoformat or shortcut string where to end reading
     */
    timestampGte: pulumi.Input<string>;
    /**
     * date in isoformat or shortcut string where to start reading
     */
    timestampLte: pulumi.Input<string>;
    /**
     * chart timezone
     */
    timezone?: pulumi.Input<string>;
    /**
     * optional mappings to transform data with rules
     */
    valueMappings?: pulumi.Input<pulumi.Input<inputs.DashboardBargaugeChartValueMapping>[]>;
    /**
     * chart width in cols (max 20)
     */
    width?: pulumi.Input<number>;
}
