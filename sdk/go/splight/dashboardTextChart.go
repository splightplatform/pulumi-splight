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
//			_, err = splight.NewDashboardTextChart(ctx, "dashboardChartTest", &splight.DashboardTextChartArgs{
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
//				Text:              pulumi.String("text to show"),
//				ChartItems: splight.DashboardTextChartChartItemArray{
//					&splight.DashboardTextChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardTextChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardTextChartChartItemQueryFilterAttributeArgs{
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
//					&splight.DashboardTextChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardTextChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardTextChartChartItemQueryFilterAttributeArgs{
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
//				Thresholds: splight.DashboardTextChartThresholdArray{
//					&splight.DashboardTextChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardTextChartValueMappingArray{
//					&splight.DashboardTextChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardTextChart:DashboardTextChart [options] splight_dashboard_text_chart.<name> <dashboard_chart_id>
// ```
type DashboardTextChart struct {
	pulumi.CustomResourceState

	// chart traces to be included
	ChartItems DashboardTextChartChartItemArrayOutput `pulumi:"chartItems"`
	Collection pulumi.StringPtrOutput                 `pulumi:"collection"`
	// chart description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrOutput `pulumi:"displayTimeRange"`
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
	// text to display
	Text pulumi.StringPtrOutput `pulumi:"text"`
	// optional static lines to be added to the chart as references
	Thresholds DashboardTextChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardTextChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardTextChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardTextChart(ctx *pulumi.Context,
	name string, args *DashboardTextChartArgs, opts ...pulumi.ResourceOption) (*DashboardTextChart, error) {
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
	var resource DashboardTextChart
	err := ctx.RegisterResource("splight:index/dashboardTextChart:DashboardTextChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardTextChart gets an existing DashboardTextChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardTextChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardTextChartState, opts ...pulumi.ResourceOption) (*DashboardTextChart, error) {
	var resource DashboardTextChart
	err := ctx.ReadResource("splight:index/dashboardTextChart:DashboardTextChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardTextChart resources.
type dashboardTextChartState struct {
	// chart traces to be included
	ChartItems []DashboardTextChartChartItem `pulumi:"chartItems"`
	Collection *string                       `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
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
	// text to display
	Text *string `pulumi:"text"`
	// optional static lines to be added to the chart as references
	Thresholds []DashboardTextChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardTextChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardTextChartState struct {
	// chart traces to be included
	ChartItems DashboardTextChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
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
	// text to display
	Text pulumi.StringPtrInput
	// optional static lines to be added to the chart as references
	Thresholds DashboardTextChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardTextChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardTextChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardTextChartState)(nil)).Elem()
}

type dashboardTextChartArgs struct {
	// chart traces to be included
	ChartItems []DashboardTextChartChartItem `pulumi:"chartItems"`
	Collection *string                       `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
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
	// text to display
	Text *string `pulumi:"text"`
	// optional static lines to be added to the chart as references
	Thresholds []DashboardTextChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardTextChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardTextChart resource.
type DashboardTextChartArgs struct {
	// chart traces to be included
	ChartItems DashboardTextChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
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
	// text to display
	Text pulumi.StringPtrInput
	// optional static lines to be added to the chart as references
	Thresholds DashboardTextChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardTextChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardTextChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardTextChartArgs)(nil)).Elem()
}

type DashboardTextChartInput interface {
	pulumi.Input

	ToDashboardTextChartOutput() DashboardTextChartOutput
	ToDashboardTextChartOutputWithContext(ctx context.Context) DashboardTextChartOutput
}

func (*DashboardTextChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardTextChart)(nil)).Elem()
}

func (i *DashboardTextChart) ToDashboardTextChartOutput() DashboardTextChartOutput {
	return i.ToDashboardTextChartOutputWithContext(context.Background())
}

func (i *DashboardTextChart) ToDashboardTextChartOutputWithContext(ctx context.Context) DashboardTextChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardTextChartOutput)
}

// DashboardTextChartArrayInput is an input type that accepts DashboardTextChartArray and DashboardTextChartArrayOutput values.
// You can construct a concrete instance of `DashboardTextChartArrayInput` via:
//
//	DashboardTextChartArray{ DashboardTextChartArgs{...} }
type DashboardTextChartArrayInput interface {
	pulumi.Input

	ToDashboardTextChartArrayOutput() DashboardTextChartArrayOutput
	ToDashboardTextChartArrayOutputWithContext(context.Context) DashboardTextChartArrayOutput
}

type DashboardTextChartArray []DashboardTextChartInput

func (DashboardTextChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardTextChart)(nil)).Elem()
}

func (i DashboardTextChartArray) ToDashboardTextChartArrayOutput() DashboardTextChartArrayOutput {
	return i.ToDashboardTextChartArrayOutputWithContext(context.Background())
}

func (i DashboardTextChartArray) ToDashboardTextChartArrayOutputWithContext(ctx context.Context) DashboardTextChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardTextChartArrayOutput)
}

// DashboardTextChartMapInput is an input type that accepts DashboardTextChartMap and DashboardTextChartMapOutput values.
// You can construct a concrete instance of `DashboardTextChartMapInput` via:
//
//	DashboardTextChartMap{ "key": DashboardTextChartArgs{...} }
type DashboardTextChartMapInput interface {
	pulumi.Input

	ToDashboardTextChartMapOutput() DashboardTextChartMapOutput
	ToDashboardTextChartMapOutputWithContext(context.Context) DashboardTextChartMapOutput
}

type DashboardTextChartMap map[string]DashboardTextChartInput

func (DashboardTextChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardTextChart)(nil)).Elem()
}

func (i DashboardTextChartMap) ToDashboardTextChartMapOutput() DashboardTextChartMapOutput {
	return i.ToDashboardTextChartMapOutputWithContext(context.Background())
}

func (i DashboardTextChartMap) ToDashboardTextChartMapOutputWithContext(ctx context.Context) DashboardTextChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardTextChartMapOutput)
}

type DashboardTextChartOutput struct{ *pulumi.OutputState }

func (DashboardTextChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardTextChart)(nil)).Elem()
}

func (o DashboardTextChartOutput) ToDashboardTextChartOutput() DashboardTextChartOutput {
	return o
}

func (o DashboardTextChartOutput) ToDashboardTextChartOutputWithContext(ctx context.Context) DashboardTextChartOutput {
	return o
}

// chart traces to be included
func (o DashboardTextChartOutput) ChartItems() DashboardTextChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardTextChart) DashboardTextChartChartItemArrayOutput { return v.ChartItems }).(DashboardTextChartChartItemArrayOutput)
}

func (o DashboardTextChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardTextChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardTextChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// chart height in px
func (o DashboardTextChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardTextChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardTextChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardTextChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardTextChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardTextChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardTextChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// chart x position
func (o DashboardTextChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardTextChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardTextChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardTextChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardTextChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardTextChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// text to display
func (o DashboardTextChartOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.Text }).(pulumi.StringPtrOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardTextChartOutput) Thresholds() DashboardTextChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardTextChart) DashboardTextChartThresholdArrayOutput { return v.Thresholds }).(DashboardTextChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardTextChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardTextChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardTextChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardTextChartOutput) ValueMappings() DashboardTextChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardTextChart) DashboardTextChartValueMappingArrayOutput { return v.ValueMappings }).(DashboardTextChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardTextChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardTextChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardTextChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardTextChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardTextChart)(nil)).Elem()
}

func (o DashboardTextChartArrayOutput) ToDashboardTextChartArrayOutput() DashboardTextChartArrayOutput {
	return o
}

func (o DashboardTextChartArrayOutput) ToDashboardTextChartArrayOutputWithContext(ctx context.Context) DashboardTextChartArrayOutput {
	return o
}

func (o DashboardTextChartArrayOutput) Index(i pulumi.IntInput) DashboardTextChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardTextChart {
		return vs[0].([]*DashboardTextChart)[vs[1].(int)]
	}).(DashboardTextChartOutput)
}

type DashboardTextChartMapOutput struct{ *pulumi.OutputState }

func (DashboardTextChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardTextChart)(nil)).Elem()
}

func (o DashboardTextChartMapOutput) ToDashboardTextChartMapOutput() DashboardTextChartMapOutput {
	return o
}

func (o DashboardTextChartMapOutput) ToDashboardTextChartMapOutputWithContext(ctx context.Context) DashboardTextChartMapOutput {
	return o
}

func (o DashboardTextChartMapOutput) MapIndex(k pulumi.StringInput) DashboardTextChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardTextChart {
		return vs[0].(map[string]*DashboardTextChart)[vs[1].(string)]
	}).(DashboardTextChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardTextChartInput)(nil)).Elem(), &DashboardTextChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardTextChartArrayInput)(nil)).Elem(), DashboardTextChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardTextChartMapInput)(nil)).Elem(), DashboardTextChartMap{})
	pulumi.RegisterOutputType(DashboardTextChartOutput{})
	pulumi.RegisterOutputType(DashboardTextChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardTextChartMapOutput{})
}