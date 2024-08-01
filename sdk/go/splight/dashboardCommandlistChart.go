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
//				RelatedAssets: pulumi.StringArray{},
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
//			_, err = splight.NewDashboardCommandlistChart(ctx, "dashboardChartTest", &splight.DashboardCommandlistChartArgs{
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
//				CommandListType:   pulumi.String("table"),
//				FilterName:        pulumi.String("some name"),
//				ChartItems: splight.DashboardCommandlistChartChartItemArray{
//					&splight.DashboardCommandlistChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardCommandlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardCommandlistChartChartItemQueryFilterAttributeArgs{
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
//					&splight.DashboardCommandlistChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardCommandlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardCommandlistChartChartItemQueryFilterAttributeArgs{
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
//				Thresholds: splight.DashboardCommandlistChartThresholdArray{
//					&splight.DashboardCommandlistChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardCommandlistChartValueMappingArray{
//					&splight.DashboardCommandlistChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardCommandlistChart:DashboardCommandlistChart [options] splight_dashboard_commandlist_chart.<name> <dashboard_chart_id>
// ```
type DashboardCommandlistChart struct {
	pulumi.CustomResourceState

	// chart traces to be included
	ChartItems DashboardCommandlistChartChartItemArrayOutput `pulumi:"chartItems"`
	Collection pulumi.StringPtrOutput                        `pulumi:"collection"`
	// [table|button_list]command list type
	CommandListType pulumi.StringPtrOutput `pulumi:"commandListType"`
	// chart description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrOutput `pulumi:"displayTimeRange"`
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
	Thresholds DashboardCommandlistChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardCommandlistChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardCommandlistChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardCommandlistChart(ctx *pulumi.Context,
	name string, args *DashboardCommandlistChartArgs, opts ...pulumi.ResourceOption) (*DashboardCommandlistChart, error) {
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
	var resource DashboardCommandlistChart
	err := ctx.RegisterResource("splight:index/dashboardCommandlistChart:DashboardCommandlistChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardCommandlistChart gets an existing DashboardCommandlistChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardCommandlistChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardCommandlistChartState, opts ...pulumi.ResourceOption) (*DashboardCommandlistChart, error) {
	var resource DashboardCommandlistChart
	err := ctx.ReadResource("splight:index/dashboardCommandlistChart:DashboardCommandlistChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardCommandlistChart resources.
type dashboardCommandlistChartState struct {
	// chart traces to be included
	ChartItems []DashboardCommandlistChartChartItem `pulumi:"chartItems"`
	Collection *string                              `pulumi:"collection"`
	// [table|button_list]command list type
	CommandListType *string `pulumi:"commandListType"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
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
	Thresholds []DashboardCommandlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardCommandlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardCommandlistChartState struct {
	// chart traces to be included
	ChartItems DashboardCommandlistChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// [table|button_list]command list type
	CommandListType pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
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
	Thresholds DashboardCommandlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardCommandlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardCommandlistChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardCommandlistChartState)(nil)).Elem()
}

type dashboardCommandlistChartArgs struct {
	// chart traces to be included
	ChartItems []DashboardCommandlistChartChartItem `pulumi:"chartItems"`
	Collection *string                              `pulumi:"collection"`
	// [table|button_list]command list type
	CommandListType *string `pulumi:"commandListType"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
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
	Thresholds []DashboardCommandlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardCommandlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardCommandlistChart resource.
type DashboardCommandlistChartArgs struct {
	// chart traces to be included
	ChartItems DashboardCommandlistChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// [table|button_list]command list type
	CommandListType pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
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
	Thresholds DashboardCommandlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardCommandlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardCommandlistChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardCommandlistChartArgs)(nil)).Elem()
}

type DashboardCommandlistChartInput interface {
	pulumi.Input

	ToDashboardCommandlistChartOutput() DashboardCommandlistChartOutput
	ToDashboardCommandlistChartOutputWithContext(ctx context.Context) DashboardCommandlistChartOutput
}

func (*DashboardCommandlistChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardCommandlistChart)(nil)).Elem()
}

func (i *DashboardCommandlistChart) ToDashboardCommandlistChartOutput() DashboardCommandlistChartOutput {
	return i.ToDashboardCommandlistChartOutputWithContext(context.Background())
}

func (i *DashboardCommandlistChart) ToDashboardCommandlistChartOutputWithContext(ctx context.Context) DashboardCommandlistChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardCommandlistChartOutput)
}

// DashboardCommandlistChartArrayInput is an input type that accepts DashboardCommandlistChartArray and DashboardCommandlistChartArrayOutput values.
// You can construct a concrete instance of `DashboardCommandlistChartArrayInput` via:
//
//	DashboardCommandlistChartArray{ DashboardCommandlistChartArgs{...} }
type DashboardCommandlistChartArrayInput interface {
	pulumi.Input

	ToDashboardCommandlistChartArrayOutput() DashboardCommandlistChartArrayOutput
	ToDashboardCommandlistChartArrayOutputWithContext(context.Context) DashboardCommandlistChartArrayOutput
}

type DashboardCommandlistChartArray []DashboardCommandlistChartInput

func (DashboardCommandlistChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardCommandlistChart)(nil)).Elem()
}

func (i DashboardCommandlistChartArray) ToDashboardCommandlistChartArrayOutput() DashboardCommandlistChartArrayOutput {
	return i.ToDashboardCommandlistChartArrayOutputWithContext(context.Background())
}

func (i DashboardCommandlistChartArray) ToDashboardCommandlistChartArrayOutputWithContext(ctx context.Context) DashboardCommandlistChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardCommandlistChartArrayOutput)
}

// DashboardCommandlistChartMapInput is an input type that accepts DashboardCommandlistChartMap and DashboardCommandlistChartMapOutput values.
// You can construct a concrete instance of `DashboardCommandlistChartMapInput` via:
//
//	DashboardCommandlistChartMap{ "key": DashboardCommandlistChartArgs{...} }
type DashboardCommandlistChartMapInput interface {
	pulumi.Input

	ToDashboardCommandlistChartMapOutput() DashboardCommandlistChartMapOutput
	ToDashboardCommandlistChartMapOutputWithContext(context.Context) DashboardCommandlistChartMapOutput
}

type DashboardCommandlistChartMap map[string]DashboardCommandlistChartInput

func (DashboardCommandlistChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardCommandlistChart)(nil)).Elem()
}

func (i DashboardCommandlistChartMap) ToDashboardCommandlistChartMapOutput() DashboardCommandlistChartMapOutput {
	return i.ToDashboardCommandlistChartMapOutputWithContext(context.Background())
}

func (i DashboardCommandlistChartMap) ToDashboardCommandlistChartMapOutputWithContext(ctx context.Context) DashboardCommandlistChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardCommandlistChartMapOutput)
}

type DashboardCommandlistChartOutput struct{ *pulumi.OutputState }

func (DashboardCommandlistChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardCommandlistChart)(nil)).Elem()
}

func (o DashboardCommandlistChartOutput) ToDashboardCommandlistChartOutput() DashboardCommandlistChartOutput {
	return o
}

func (o DashboardCommandlistChartOutput) ToDashboardCommandlistChartOutputWithContext(ctx context.Context) DashboardCommandlistChartOutput {
	return o
}

// chart traces to be included
func (o DashboardCommandlistChartOutput) ChartItems() DashboardCommandlistChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) DashboardCommandlistChartChartItemArrayOutput { return v.ChartItems }).(DashboardCommandlistChartChartItemArrayOutput)
}

func (o DashboardCommandlistChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// [table|button_list]command list type
func (o DashboardCommandlistChartOutput) CommandListType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.CommandListType }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardCommandlistChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardCommandlistChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// filter name
func (o DashboardCommandlistChartOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.FilterName }).(pulumi.StringPtrOutput)
}

// chart height in px
func (o DashboardCommandlistChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardCommandlistChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardCommandlistChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardCommandlistChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardCommandlistChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardCommandlistChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardCommandlistChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// chart x position
func (o DashboardCommandlistChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardCommandlistChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardCommandlistChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardCommandlistChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardCommandlistChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardCommandlistChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardCommandlistChartOutput) Thresholds() DashboardCommandlistChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) DashboardCommandlistChartThresholdArrayOutput { return v.Thresholds }).(DashboardCommandlistChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardCommandlistChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardCommandlistChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardCommandlistChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardCommandlistChartOutput) ValueMappings() DashboardCommandlistChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) DashboardCommandlistChartValueMappingArrayOutput {
		return v.ValueMappings
	}).(DashboardCommandlistChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardCommandlistChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardCommandlistChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardCommandlistChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardCommandlistChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardCommandlistChart)(nil)).Elem()
}

func (o DashboardCommandlistChartArrayOutput) ToDashboardCommandlistChartArrayOutput() DashboardCommandlistChartArrayOutput {
	return o
}

func (o DashboardCommandlistChartArrayOutput) ToDashboardCommandlistChartArrayOutputWithContext(ctx context.Context) DashboardCommandlistChartArrayOutput {
	return o
}

func (o DashboardCommandlistChartArrayOutput) Index(i pulumi.IntInput) DashboardCommandlistChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardCommandlistChart {
		return vs[0].([]*DashboardCommandlistChart)[vs[1].(int)]
	}).(DashboardCommandlistChartOutput)
}

type DashboardCommandlistChartMapOutput struct{ *pulumi.OutputState }

func (DashboardCommandlistChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardCommandlistChart)(nil)).Elem()
}

func (o DashboardCommandlistChartMapOutput) ToDashboardCommandlistChartMapOutput() DashboardCommandlistChartMapOutput {
	return o
}

func (o DashboardCommandlistChartMapOutput) ToDashboardCommandlistChartMapOutputWithContext(ctx context.Context) DashboardCommandlistChartMapOutput {
	return o
}

func (o DashboardCommandlistChartMapOutput) MapIndex(k pulumi.StringInput) DashboardCommandlistChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardCommandlistChart {
		return vs[0].(map[string]*DashboardCommandlistChart)[vs[1].(string)]
	}).(DashboardCommandlistChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardCommandlistChartInput)(nil)).Elem(), &DashboardCommandlistChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardCommandlistChartArrayInput)(nil)).Elem(), DashboardCommandlistChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardCommandlistChartMapInput)(nil)).Elem(), DashboardCommandlistChartMap{})
	pulumi.RegisterOutputType(DashboardCommandlistChartOutput{})
	pulumi.RegisterOutputType(DashboardCommandlistChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardCommandlistChartMapOutput{})
}
