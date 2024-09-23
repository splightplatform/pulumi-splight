// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splight

import (
	"context"
	"reflect"

	"errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/splightplatform/pulumi-splight/sdk/go/splight/internal"
)

// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"encoding/json"
//
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/splightplatform/pulumi-splight/sdk/go/splight"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			tmpJSON0, err := json.Marshal(map[string]interface{}{
//				"type":       "GeometryCollection",
//				"geometries": []interface{}{},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			assetTest, err := splight.NewAsset(ctx, "assetTest", &splight.AssetArgs{
//				Description: pulumi.String("Created with Terraform"),
//				Geometry:    pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			attributeTest1, err := splight.NewAssetAttribute(ctx, "attributeTest1", &splight.AssetAttributeArgs{
//				Type:  pulumi.String("Number"),
//				Unit:  pulumi.String("meters"),
//				Asset: assetTest.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			attributeTest2, err := splight.NewAssetAttribute(ctx, "attributeTest2", &splight.AssetAttributeArgs{
//				Type:  pulumi.String("Number"),
//				Unit:  pulumi.String("seconds"),
//				Asset: assetTest.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			dashboardTest, err := splight.NewDashboard(ctx, "dashboardTest", &splight.DashboardArgs{
//				RelatedAssets: splight.DashboardRelatedAssetArray{},
//			})
//			if err != nil {
//				return err
//			}
//			dashboardTabTest, err := splight.NewDashboardTab(ctx, "dashboardTabTest", &splight.DashboardTabArgs{
//				Order:     pulumi.Int(0),
//				Dashboard: dashboardTest.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = splight.NewDashboardActionlistChart(ctx, "dashboardChartTest", &splight.DashboardActionlistChartArgs{
//				Tab:               dashboardTabTest.ID(),
//				TimestampGte:      pulumi.String("now - 7d"),
//				TimestampLte:      pulumi.String("now"),
//				Description:       pulumi.String("Chart description"),
//				MinHeight:         pulumi.Int(1),
//				MinWidth:          pulumi.Int(4),
//				DisplayTimeRange:  pulumi.Bool(true),
//				LabelsDisplay:     pulumi.Bool(true),
//				LabelsAggregation: pulumi.String("last"),
//				LabelsPlacement:   pulumi.String("bottom"),
//				ShowBeyondData:    pulumi.Bool(true),
//				Height:            pulumi.Int(10),
//				Width:             pulumi.Int(20),
//				Collection:        pulumi.String("default"),
//				ActionListType:    pulumi.String("table"),
//				FilterName:        pulumi.String("some name"),
//				FilterAssetName:   pulumi.String("some asset name"),
//				ChartItems: splight.DashboardActionlistChartChartItemArray{
//					&splight.DashboardActionlistChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardActionlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardActionlistChartChartItemQueryFilterAttributeArgs{
//							Id:   attributeTest1.ID(),
//							Name: attributeTest1.Name,
//						},
//						QueryPlain: pulumi.All(assetTest.ID(), attributeTest1.ID()).ApplyT(func(_args []interface{}) (string, error) {
//							assetTestId := _args[0].(string)
//							attributeTest1Id := _args[1].(string)
//							var _zero string
//							tmpJSON1, err := json.Marshal([]interface{}{
//								map[string]interface{}{
//									"$match": map[string]interface{}{
//										"asset":     assetTestId,
//										"attribute": attributeTest1Id,
//									},
//								},
//								map[string]interface{}{
//									"$addFields": map[string]interface{}{
//										"timestamp": map[string]interface{}{
//											"$dateTrunc": map[string]interface{}{
//												"date":    "$timestamp",
//												"unit":    "day",
//												"binSize": 1,
//											},
//										},
//									},
//								},
//								map[string]interface{}{
//									"$group": map[string]interface{}{
//										"_id": "$timestamp",
//										"value": map[string]interface{}{
//											"$last": "$value",
//										},
//										"timestamp": map[string]interface{}{
//											"$last": "$timestamp",
//										},
//									},
//								},
//							})
//							if err != nil {
//								return _zero, err
//							}
//							json1 := string(tmpJSON1)
//							return json1, nil
//						}).(pulumi.StringOutput),
//					},
//					&splight.DashboardActionlistChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardActionlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardActionlistChartChartItemQueryFilterAttributeArgs{
//							Id:   attributeTest2.ID(),
//							Name: attributeTest2.Name,
//						},
//						QueryPlain: pulumi.All(assetTest.ID(), attributeTest2.ID()).ApplyT(func(_args []interface{}) (string, error) {
//							assetTestId := _args[0].(string)
//							attributeTest2Id := _args[1].(string)
//							var _zero string
//							tmpJSON2, err := json.Marshal([]interface{}{
//								map[string]interface{}{
//									"$match": map[string]interface{}{
//										"asset":     assetTestId,
//										"attribute": attributeTest2Id,
//									},
//								},
//								map[string]interface{}{
//									"$addFields": map[string]interface{}{
//										"timestamp": map[string]interface{}{
//											"$dateTrunc": map[string]interface{}{
//												"date":    "$timestamp",
//												"unit":    "hour",
//												"binSize": 1,
//											},
//										},
//									},
//								},
//								map[string]interface{}{
//									"$group": map[string]interface{}{
//										"_id": "$timestamp",
//										"value": map[string]interface{}{
//											"$last": "$value",
//										},
//										"timestamp": map[string]interface{}{
//											"$last": "$timestamp",
//										},
//									},
//								},
//							})
//							if err != nil {
//								return _zero, err
//							}
//							json2 := string(tmpJSON2)
//							return json2, nil
//						}).(pulumi.StringOutput),
//					},
//				},
//				Thresholds: splight.DashboardActionlistChartThresholdArray{
//					&splight.DashboardActionlistChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardActionlistChartValueMappingArray{
//					&splight.DashboardActionlistChartValueMappingArgs{
//						DisplayText: pulumi.String("MODIFICADO"),
//						MatchValue:  pulumi.String("123.3"),
//						Type:        pulumi.String("exact_match"),
//						Order:       pulumi.Int(0),
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
//
// ## Import
//
// ```sh
// $ pulumi import splight:index/dashboardActionlistChart:DashboardActionlistChart [options] splight_dashboard_actionlist_chart.<name> <dashboard_chart_id>
// ```
type DashboardActionlistChart struct {
	pulumi.CustomResourceState

	// action list type
	ActionListType pulumi.StringPtrOutput `pulumi:"actionListType"`
	// chart traces to be included
	ChartItems DashboardActionlistChartChartItemArrayOutput `pulumi:"chartItems"`
	Collection pulumi.StringPtrOutput                       `pulumi:"collection"`
	// chart description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrOutput `pulumi:"displayTimeRange"`
	// filter asset name
	FilterAssetName pulumi.StringPtrOutput `pulumi:"filterAssetName"`
	// filter name
	FilterName pulumi.StringPtrOutput `pulumi:"filterName"`
	// chart height in px
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// [last|avg|...] aggregation
	LabelsAggregation pulumi.StringPtrOutput `pulumi:"labelsAggregation"`
	// whether to display the labels or not
	LabelsDisplay pulumi.BoolPtrOutput `pulumi:"labelsDisplay"`
	// [right|bottom] placement
	LabelsPlacement pulumi.StringPtrOutput `pulumi:"labelsPlacement"`
	// minimum chart height
	MinHeight pulumi.IntPtrOutput `pulumi:"minHeight"`
	// minimum chart width
	MinWidth pulumi.IntPtrOutput `pulumi:"minWidth"`
	// name of the chart
	Name pulumi.StringOutput `pulumi:"name"`
	// chart x position
	PositionX pulumi.IntPtrOutput `pulumi:"positionX"`
	// chart y position
	PositionY pulumi.IntPtrOutput `pulumi:"positionY"`
	// refresh interval
	RefreshInterval pulumi.StringPtrOutput `pulumi:"refreshInterval"`
	// relative window time
	RelativeWindowTime pulumi.StringPtrOutput `pulumi:"relativeWindowTime"`
	// whether to show data which is beyond timestampLte or not
	ShowBeyondData pulumi.BoolPtrOutput `pulumi:"showBeyondData"`
	// id for the tab where to place the chart
	Tab pulumi.StringOutput `pulumi:"tab"`
	// optional static lines to be added to the chart as references
	Thresholds DashboardActionlistChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardActionlistChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardActionlistChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardActionlistChart(ctx *pulumi.Context,
	name string, args *DashboardActionlistChartArgs, opts ...pulumi.ResourceOption) (*DashboardActionlistChart, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ChartItems == nil {
		return nil, errors.New("invalid value for required argument 'ChartItems'")
	}
	if args.Tab == nil {
		return nil, errors.New("invalid value for required argument 'Tab'")
	}
	if args.TimestampGte == nil {
		return nil, errors.New("invalid value for required argument 'TimestampGte'")
	}
	if args.TimestampLte == nil {
		return nil, errors.New("invalid value for required argument 'TimestampLte'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DashboardActionlistChart
	err := ctx.RegisterResource("splight:index/dashboardActionlistChart:DashboardActionlistChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardActionlistChart gets an existing DashboardActionlistChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardActionlistChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardActionlistChartState, opts ...pulumi.ResourceOption) (*DashboardActionlistChart, error) {
	var resource DashboardActionlistChart
	err := ctx.ReadResource("splight:index/dashboardActionlistChart:DashboardActionlistChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardActionlistChart resources.
type dashboardActionlistChartState struct {
	// action list type
	ActionListType *string `pulumi:"actionListType"`
	// chart traces to be included
	ChartItems []DashboardActionlistChartChartItem `pulumi:"chartItems"`
	Collection *string                             `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// filter asset name
	FilterAssetName *string `pulumi:"filterAssetName"`
	// filter name
	FilterName *string `pulumi:"filterName"`
	// chart height in px
	Height *int `pulumi:"height"`
	// [last|avg|...] aggregation
	LabelsAggregation *string `pulumi:"labelsAggregation"`
	// whether to display the labels or not
	LabelsDisplay *bool `pulumi:"labelsDisplay"`
	// [right|bottom] placement
	LabelsPlacement *string `pulumi:"labelsPlacement"`
	// minimum chart height
	MinHeight *int `pulumi:"minHeight"`
	// minimum chart width
	MinWidth *int `pulumi:"minWidth"`
	// name of the chart
	Name *string `pulumi:"name"`
	// chart x position
	PositionX *int `pulumi:"positionX"`
	// chart y position
	PositionY *int `pulumi:"positionY"`
	// refresh interval
	RefreshInterval *string `pulumi:"refreshInterval"`
	// relative window time
	RelativeWindowTime *string `pulumi:"relativeWindowTime"`
	// whether to show data which is beyond timestampLte or not
	ShowBeyondData *bool `pulumi:"showBeyondData"`
	// id for the tab where to place the chart
	Tab *string `pulumi:"tab"`
	// optional static lines to be added to the chart as references
	Thresholds []DashboardActionlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardActionlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardActionlistChartState struct {
	// action list type
	ActionListType pulumi.StringPtrInput
	// chart traces to be included
	ChartItems DashboardActionlistChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// filter asset name
	FilterAssetName pulumi.StringPtrInput
	// filter name
	FilterName pulumi.StringPtrInput
	// chart height in px
	Height pulumi.IntPtrInput
	// [last|avg|...] aggregation
	LabelsAggregation pulumi.StringPtrInput
	// whether to display the labels or not
	LabelsDisplay pulumi.BoolPtrInput
	// [right|bottom] placement
	LabelsPlacement pulumi.StringPtrInput
	// minimum chart height
	MinHeight pulumi.IntPtrInput
	// minimum chart width
	MinWidth pulumi.IntPtrInput
	// name of the chart
	Name pulumi.StringPtrInput
	// chart x position
	PositionX pulumi.IntPtrInput
	// chart y position
	PositionY pulumi.IntPtrInput
	// refresh interval
	RefreshInterval pulumi.StringPtrInput
	// relative window time
	RelativeWindowTime pulumi.StringPtrInput
	// whether to show data which is beyond timestampLte or not
	ShowBeyondData pulumi.BoolPtrInput
	// id for the tab where to place the chart
	Tab pulumi.StringPtrInput
	// optional static lines to be added to the chart as references
	Thresholds DashboardActionlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardActionlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardActionlistChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardActionlistChartState)(nil)).Elem()
}

type dashboardActionlistChartArgs struct {
	// action list type
	ActionListType *string `pulumi:"actionListType"`
	// chart traces to be included
	ChartItems []DashboardActionlistChartChartItem `pulumi:"chartItems"`
	Collection *string                             `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// filter asset name
	FilterAssetName *string `pulumi:"filterAssetName"`
	// filter name
	FilterName *string `pulumi:"filterName"`
	// chart height in px
	Height *int `pulumi:"height"`
	// [last|avg|...] aggregation
	LabelsAggregation *string `pulumi:"labelsAggregation"`
	// whether to display the labels or not
	LabelsDisplay *bool `pulumi:"labelsDisplay"`
	// [right|bottom] placement
	LabelsPlacement *string `pulumi:"labelsPlacement"`
	// minimum chart height
	MinHeight *int `pulumi:"minHeight"`
	// minimum chart width
	MinWidth *int `pulumi:"minWidth"`
	// name of the chart
	Name *string `pulumi:"name"`
	// chart x position
	PositionX *int `pulumi:"positionX"`
	// chart y position
	PositionY *int `pulumi:"positionY"`
	// refresh interval
	RefreshInterval *string `pulumi:"refreshInterval"`
	// relative window time
	RelativeWindowTime *string `pulumi:"relativeWindowTime"`
	// whether to show data which is beyond timestampLte or not
	ShowBeyondData *bool `pulumi:"showBeyondData"`
	// id for the tab where to place the chart
	Tab string `pulumi:"tab"`
	// optional static lines to be added to the chart as references
	Thresholds []DashboardActionlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardActionlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardActionlistChart resource.
type DashboardActionlistChartArgs struct {
	// action list type
	ActionListType pulumi.StringPtrInput
	// chart traces to be included
	ChartItems DashboardActionlistChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// filter asset name
	FilterAssetName pulumi.StringPtrInput
	// filter name
	FilterName pulumi.StringPtrInput
	// chart height in px
	Height pulumi.IntPtrInput
	// [last|avg|...] aggregation
	LabelsAggregation pulumi.StringPtrInput
	// whether to display the labels or not
	LabelsDisplay pulumi.BoolPtrInput
	// [right|bottom] placement
	LabelsPlacement pulumi.StringPtrInput
	// minimum chart height
	MinHeight pulumi.IntPtrInput
	// minimum chart width
	MinWidth pulumi.IntPtrInput
	// name of the chart
	Name pulumi.StringPtrInput
	// chart x position
	PositionX pulumi.IntPtrInput
	// chart y position
	PositionY pulumi.IntPtrInput
	// refresh interval
	RefreshInterval pulumi.StringPtrInput
	// relative window time
	RelativeWindowTime pulumi.StringPtrInput
	// whether to show data which is beyond timestampLte or not
	ShowBeyondData pulumi.BoolPtrInput
	// id for the tab where to place the chart
	Tab pulumi.StringInput
	// optional static lines to be added to the chart as references
	Thresholds DashboardActionlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardActionlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardActionlistChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardActionlistChartArgs)(nil)).Elem()
}

type DashboardActionlistChartInput interface {
	pulumi.Input

	ToDashboardActionlistChartOutput() DashboardActionlistChartOutput
	ToDashboardActionlistChartOutputWithContext(ctx context.Context) DashboardActionlistChartOutput
}

func (*DashboardActionlistChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardActionlistChart)(nil)).Elem()
}

func (i *DashboardActionlistChart) ToDashboardActionlistChartOutput() DashboardActionlistChartOutput {
	return i.ToDashboardActionlistChartOutputWithContext(context.Background())
}

func (i *DashboardActionlistChart) ToDashboardActionlistChartOutputWithContext(ctx context.Context) DashboardActionlistChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardActionlistChartOutput)
}

// DashboardActionlistChartArrayInput is an input type that accepts DashboardActionlistChartArray and DashboardActionlistChartArrayOutput values.
// You can construct a concrete instance of `DashboardActionlistChartArrayInput` via:
//
//	DashboardActionlistChartArray{ DashboardActionlistChartArgs{...} }
type DashboardActionlistChartArrayInput interface {
	pulumi.Input

	ToDashboardActionlistChartArrayOutput() DashboardActionlistChartArrayOutput
	ToDashboardActionlistChartArrayOutputWithContext(context.Context) DashboardActionlistChartArrayOutput
}

type DashboardActionlistChartArray []DashboardActionlistChartInput

func (DashboardActionlistChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardActionlistChart)(nil)).Elem()
}

func (i DashboardActionlistChartArray) ToDashboardActionlistChartArrayOutput() DashboardActionlistChartArrayOutput {
	return i.ToDashboardActionlistChartArrayOutputWithContext(context.Background())
}

func (i DashboardActionlistChartArray) ToDashboardActionlistChartArrayOutputWithContext(ctx context.Context) DashboardActionlistChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardActionlistChartArrayOutput)
}

// DashboardActionlistChartMapInput is an input type that accepts DashboardActionlistChartMap and DashboardActionlistChartMapOutput values.
// You can construct a concrete instance of `DashboardActionlistChartMapInput` via:
//
//	DashboardActionlistChartMap{ "key": DashboardActionlistChartArgs{...} }
type DashboardActionlistChartMapInput interface {
	pulumi.Input

	ToDashboardActionlistChartMapOutput() DashboardActionlistChartMapOutput
	ToDashboardActionlistChartMapOutputWithContext(context.Context) DashboardActionlistChartMapOutput
}

type DashboardActionlistChartMap map[string]DashboardActionlistChartInput

func (DashboardActionlistChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardActionlistChart)(nil)).Elem()
}

func (i DashboardActionlistChartMap) ToDashboardActionlistChartMapOutput() DashboardActionlistChartMapOutput {
	return i.ToDashboardActionlistChartMapOutputWithContext(context.Background())
}

func (i DashboardActionlistChartMap) ToDashboardActionlistChartMapOutputWithContext(ctx context.Context) DashboardActionlistChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardActionlistChartMapOutput)
}

type DashboardActionlistChartOutput struct{ *pulumi.OutputState }

func (DashboardActionlistChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardActionlistChart)(nil)).Elem()
}

func (o DashboardActionlistChartOutput) ToDashboardActionlistChartOutput() DashboardActionlistChartOutput {
	return o
}

func (o DashboardActionlistChartOutput) ToDashboardActionlistChartOutputWithContext(ctx context.Context) DashboardActionlistChartOutput {
	return o
}

// action list type
func (o DashboardActionlistChartOutput) ActionListType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.ActionListType }).(pulumi.StringPtrOutput)
}

// chart traces to be included
func (o DashboardActionlistChartOutput) ChartItems() DashboardActionlistChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) DashboardActionlistChartChartItemArrayOutput { return v.ChartItems }).(DashboardActionlistChartChartItemArrayOutput)
}

func (o DashboardActionlistChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardActionlistChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardActionlistChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// filter asset name
func (o DashboardActionlistChartOutput) FilterAssetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.FilterAssetName }).(pulumi.StringPtrOutput)
}

// filter name
func (o DashboardActionlistChartOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.FilterName }).(pulumi.StringPtrOutput)
}

// chart height in px
func (o DashboardActionlistChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardActionlistChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardActionlistChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardActionlistChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardActionlistChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardActionlistChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardActionlistChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// chart x position
func (o DashboardActionlistChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardActionlistChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardActionlistChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardActionlistChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardActionlistChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardActionlistChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardActionlistChartOutput) Thresholds() DashboardActionlistChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) DashboardActionlistChartThresholdArrayOutput { return v.Thresholds }).(DashboardActionlistChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardActionlistChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardActionlistChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardActionlistChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardActionlistChartOutput) ValueMappings() DashboardActionlistChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) DashboardActionlistChartValueMappingArrayOutput {
		return v.ValueMappings
	}).(DashboardActionlistChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardActionlistChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardActionlistChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardActionlistChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardActionlistChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardActionlistChart)(nil)).Elem()
}

func (o DashboardActionlistChartArrayOutput) ToDashboardActionlistChartArrayOutput() DashboardActionlistChartArrayOutput {
	return o
}

func (o DashboardActionlistChartArrayOutput) ToDashboardActionlistChartArrayOutputWithContext(ctx context.Context) DashboardActionlistChartArrayOutput {
	return o
}

func (o DashboardActionlistChartArrayOutput) Index(i pulumi.IntInput) DashboardActionlistChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardActionlistChart {
		return vs[0].([]*DashboardActionlistChart)[vs[1].(int)]
	}).(DashboardActionlistChartOutput)
}

type DashboardActionlistChartMapOutput struct{ *pulumi.OutputState }

func (DashboardActionlistChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardActionlistChart)(nil)).Elem()
}

func (o DashboardActionlistChartMapOutput) ToDashboardActionlistChartMapOutput() DashboardActionlistChartMapOutput {
	return o
}

func (o DashboardActionlistChartMapOutput) ToDashboardActionlistChartMapOutputWithContext(ctx context.Context) DashboardActionlistChartMapOutput {
	return o
}

func (o DashboardActionlistChartMapOutput) MapIndex(k pulumi.StringInput) DashboardActionlistChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardActionlistChart {
		return vs[0].(map[string]*DashboardActionlistChart)[vs[1].(string)]
	}).(DashboardActionlistChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardActionlistChartInput)(nil)).Elem(), &DashboardActionlistChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardActionlistChartArrayInput)(nil)).Elem(), DashboardActionlistChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardActionlistChartMapInput)(nil)).Elem(), DashboardActionlistChartMap{})
	pulumi.RegisterOutputType(DashboardActionlistChartOutput{})
	pulumi.RegisterOutputType(DashboardActionlistChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardActionlistChartMapOutput{})
}
