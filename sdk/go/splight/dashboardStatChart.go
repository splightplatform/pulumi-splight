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
//			_, err = splight.NewDashboardStatChart(ctx, "dashboardChartTest", &splight.DashboardStatChartArgs{
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
//				YAxisUnit:         pulumi.String("MW"),
//				Border:            pulumi.Bool(false),
//				NumberOfDecimals:  pulumi.Int(4),
//				ChartItems: splight.DashboardStatChartChartItemArray{
//					&splight.DashboardStatChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardStatChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardStatChartChartItemQueryFilterAttributeArgs{
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
//					&splight.DashboardStatChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardStatChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardStatChartChartItemQueryFilterAttributeArgs{
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
//				Thresholds: splight.DashboardStatChartThresholdArray{
//					&splight.DashboardStatChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardStatChartValueMappingArray{
//					&splight.DashboardStatChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardStatChart:DashboardStatChart [options] splight_dashboard_stat_chart.<name> <dashboard_chart_id>
// ```
type DashboardStatChart struct {
	pulumi.CustomResourceState

	// whether to show the border or not
	Border pulumi.BoolPtrOutput `pulumi:"border"`
	// chart traces to be included
	ChartItems DashboardStatChartChartItemArrayOutput `pulumi:"chartItems"`
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
	// number of decimals
	NumberOfDecimals pulumi.IntPtrOutput `pulumi:"numberOfDecimals"`
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
	Thresholds DashboardStatChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardStatChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
	// y axis units
	YAxisUnit pulumi.StringPtrOutput `pulumi:"yAxisUnit"`
}

// NewDashboardStatChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardStatChart(ctx *pulumi.Context,
	name string, args *DashboardStatChartArgs, opts ...pulumi.ResourceOption) (*DashboardStatChart, error) {
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
	var resource DashboardStatChart
	err := ctx.RegisterResource("splight:index/dashboardStatChart:DashboardStatChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardStatChart gets an existing DashboardStatChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardStatChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardStatChartState, opts ...pulumi.ResourceOption) (*DashboardStatChart, error) {
	var resource DashboardStatChart
	err := ctx.ReadResource("splight:index/dashboardStatChart:DashboardStatChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardStatChart resources.
type dashboardStatChartState struct {
	// whether to show the border or not
	Border *bool `pulumi:"border"`
	// chart traces to be included
	ChartItems []DashboardStatChartChartItem `pulumi:"chartItems"`
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
	// number of decimals
	NumberOfDecimals *int `pulumi:"numberOfDecimals"`
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
	Thresholds []DashboardStatChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardStatChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
	// y axis units
	YAxisUnit *string `pulumi:"yAxisUnit"`
}

type DashboardStatChartState struct {
	// whether to show the border or not
	Border pulumi.BoolPtrInput
	// chart traces to be included
	ChartItems DashboardStatChartChartItemArrayInput
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
	// number of decimals
	NumberOfDecimals pulumi.IntPtrInput
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
	Thresholds DashboardStatChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardStatChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
	// y axis units
	YAxisUnit pulumi.StringPtrInput
}

func (DashboardStatChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardStatChartState)(nil)).Elem()
}

type dashboardStatChartArgs struct {
	// whether to show the border or not
	Border *bool `pulumi:"border"`
	// chart traces to be included
	ChartItems []DashboardStatChartChartItem `pulumi:"chartItems"`
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
	// number of decimals
	NumberOfDecimals *int `pulumi:"numberOfDecimals"`
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
	Thresholds []DashboardStatChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardStatChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
	// y axis units
	YAxisUnit *string `pulumi:"yAxisUnit"`
}

// The set of arguments for constructing a DashboardStatChart resource.
type DashboardStatChartArgs struct {
	// whether to show the border or not
	Border pulumi.BoolPtrInput
	// chart traces to be included
	ChartItems DashboardStatChartChartItemArrayInput
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
	// number of decimals
	NumberOfDecimals pulumi.IntPtrInput
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
	Thresholds DashboardStatChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardStatChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
	// y axis units
	YAxisUnit pulumi.StringPtrInput
}

func (DashboardStatChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardStatChartArgs)(nil)).Elem()
}

type DashboardStatChartInput interface {
	pulumi.Input

	ToDashboardStatChartOutput() DashboardStatChartOutput
	ToDashboardStatChartOutputWithContext(ctx context.Context) DashboardStatChartOutput
}

func (*DashboardStatChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardStatChart)(nil)).Elem()
}

func (i *DashboardStatChart) ToDashboardStatChartOutput() DashboardStatChartOutput {
	return i.ToDashboardStatChartOutputWithContext(context.Background())
}

func (i *DashboardStatChart) ToDashboardStatChartOutputWithContext(ctx context.Context) DashboardStatChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardStatChartOutput)
}

// DashboardStatChartArrayInput is an input type that accepts DashboardStatChartArray and DashboardStatChartArrayOutput values.
// You can construct a concrete instance of `DashboardStatChartArrayInput` via:
//
//	DashboardStatChartArray{ DashboardStatChartArgs{...} }
type DashboardStatChartArrayInput interface {
	pulumi.Input

	ToDashboardStatChartArrayOutput() DashboardStatChartArrayOutput
	ToDashboardStatChartArrayOutputWithContext(context.Context) DashboardStatChartArrayOutput
}

type DashboardStatChartArray []DashboardStatChartInput

func (DashboardStatChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardStatChart)(nil)).Elem()
}

func (i DashboardStatChartArray) ToDashboardStatChartArrayOutput() DashboardStatChartArrayOutput {
	return i.ToDashboardStatChartArrayOutputWithContext(context.Background())
}

func (i DashboardStatChartArray) ToDashboardStatChartArrayOutputWithContext(ctx context.Context) DashboardStatChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardStatChartArrayOutput)
}

// DashboardStatChartMapInput is an input type that accepts DashboardStatChartMap and DashboardStatChartMapOutput values.
// You can construct a concrete instance of `DashboardStatChartMapInput` via:
//
//	DashboardStatChartMap{ "key": DashboardStatChartArgs{...} }
type DashboardStatChartMapInput interface {
	pulumi.Input

	ToDashboardStatChartMapOutput() DashboardStatChartMapOutput
	ToDashboardStatChartMapOutputWithContext(context.Context) DashboardStatChartMapOutput
}

type DashboardStatChartMap map[string]DashboardStatChartInput

func (DashboardStatChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardStatChart)(nil)).Elem()
}

func (i DashboardStatChartMap) ToDashboardStatChartMapOutput() DashboardStatChartMapOutput {
	return i.ToDashboardStatChartMapOutputWithContext(context.Background())
}

func (i DashboardStatChartMap) ToDashboardStatChartMapOutputWithContext(ctx context.Context) DashboardStatChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardStatChartMapOutput)
}

type DashboardStatChartOutput struct{ *pulumi.OutputState }

func (DashboardStatChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardStatChart)(nil)).Elem()
}

func (o DashboardStatChartOutput) ToDashboardStatChartOutput() DashboardStatChartOutput {
	return o
}

func (o DashboardStatChartOutput) ToDashboardStatChartOutputWithContext(ctx context.Context) DashboardStatChartOutput {
	return o
}

// whether to show the border or not
func (o DashboardStatChartOutput) Border() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.BoolPtrOutput { return v.Border }).(pulumi.BoolPtrOutput)
}

// chart traces to be included
func (o DashboardStatChartOutput) ChartItems() DashboardStatChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardStatChart) DashboardStatChartChartItemArrayOutput { return v.ChartItems }).(DashboardStatChartChartItemArrayOutput)
}

func (o DashboardStatChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardStatChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardStatChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// chart height in px
func (o DashboardStatChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardStatChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardStatChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardStatChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardStatChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardStatChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardStatChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// number of decimals
func (o DashboardStatChartOutput) NumberOfDecimals() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.IntPtrOutput { return v.NumberOfDecimals }).(pulumi.IntPtrOutput)
}

// chart x position
func (o DashboardStatChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardStatChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardStatChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardStatChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardStatChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardStatChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardStatChartOutput) Thresholds() DashboardStatChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardStatChart) DashboardStatChartThresholdArrayOutput { return v.Thresholds }).(DashboardStatChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardStatChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardStatChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardStatChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardStatChartOutput) ValueMappings() DashboardStatChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardStatChart) DashboardStatChartValueMappingArrayOutput { return v.ValueMappings }).(DashboardStatChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardStatChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

// y axis units
func (o DashboardStatChartOutput) YAxisUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardStatChart) pulumi.StringPtrOutput { return v.YAxisUnit }).(pulumi.StringPtrOutput)
}

type DashboardStatChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardStatChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardStatChart)(nil)).Elem()
}

func (o DashboardStatChartArrayOutput) ToDashboardStatChartArrayOutput() DashboardStatChartArrayOutput {
	return o
}

func (o DashboardStatChartArrayOutput) ToDashboardStatChartArrayOutputWithContext(ctx context.Context) DashboardStatChartArrayOutput {
	return o
}

func (o DashboardStatChartArrayOutput) Index(i pulumi.IntInput) DashboardStatChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardStatChart {
		return vs[0].([]*DashboardStatChart)[vs[1].(int)]
	}).(DashboardStatChartOutput)
}

type DashboardStatChartMapOutput struct{ *pulumi.OutputState }

func (DashboardStatChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardStatChart)(nil)).Elem()
}

func (o DashboardStatChartMapOutput) ToDashboardStatChartMapOutput() DashboardStatChartMapOutput {
	return o
}

func (o DashboardStatChartMapOutput) ToDashboardStatChartMapOutputWithContext(ctx context.Context) DashboardStatChartMapOutput {
	return o
}

func (o DashboardStatChartMapOutput) MapIndex(k pulumi.StringInput) DashboardStatChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardStatChart {
		return vs[0].(map[string]*DashboardStatChart)[vs[1].(string)]
	}).(DashboardStatChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardStatChartInput)(nil)).Elem(), &DashboardStatChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardStatChartArrayInput)(nil)).Elem(), DashboardStatChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardStatChartMapInput)(nil)).Elem(), DashboardStatChartMap{})
	pulumi.RegisterOutputType(DashboardStatChartOutput{})
	pulumi.RegisterOutputType(DashboardStatChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardStatChartMapOutput{})
}
