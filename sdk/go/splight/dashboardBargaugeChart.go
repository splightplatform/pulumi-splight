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
//				Timezone:    pulumi.String("America/Los_Angeles"),
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
//			_, err = splight.NewAssetAttribute(ctx, "attributeTest2", &splight.AssetAttributeArgs{
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
//			tmpJSON1, err := json.Marshal(map[string]interface{}{
//				"$function": map[string]interface{}{
//					"body": "function ($A) { return $A/50 }",
//					"args": []string{
//						"$A",
//					},
//					"lang": "js",
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			_, err = splight.NewDashboardBargaugeChart(ctx, "dashboardChartTest", &splight.DashboardBargaugeChartArgs{
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
//				MaxLimit:          pulumi.Int(100),
//				NumberOfDecimals:  pulumi.Int(2),
//				Orientation:       pulumi.String("vertical"),
//				ChartItems: splight.DashboardBargaugeChartChartItemArray{
//					&splight.DashboardBargaugeChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardBargaugeChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardBargaugeChartChartItemQueryFilterAttributeArgs{
//							Id:   attributeTest1.ID(),
//							Name: attributeTest1.Name,
//						},
//						QueryPlain: pulumi.All(assetTest.ID(), attributeTest1.ID()).ApplyT(func(_args []interface{}) (string, error) {
//							assetTestId := _args[0].(string)
//							attributeTest1Id := _args[1].(string)
//							var _zero string
//							tmpJSON2, err := json.Marshal([]interface{}{
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
//							json2 := string(tmpJSON2)
//							return json2, nil
//						}).(pulumi.StringOutput),
//					},
//					&splight.DashboardBargaugeChartChartItemArgs{
//						RefId:                pulumi.String("B"),
//						Color:                pulumi.String("blue"),
//						Type:                 pulumi.String("EXPRESSION"),
//						QueryPlain:           pulumi.String(""),
//						ExpressionPlain:      pulumi.String(json1),
//						QueryFilterAsset:     &splight.DashboardBargaugeChartChartItemQueryFilterAssetArgs{},
//						QueryFilterAttribute: &splight.DashboardBargaugeChartChartItemQueryFilterAttributeArgs{},
//					},
//				},
//				Thresholds: splight.DashboardBargaugeChartThresholdArray{
//					&splight.DashboardBargaugeChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardBargaugeChartValueMappingArray{
//					&splight.DashboardBargaugeChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardBargaugeChart:DashboardBargaugeChart [options] splight_dashboard_bargauge_chart.<name> <dashboard_chart_id>
// ```
type DashboardBargaugeChart struct {
	pulumi.CustomResourceState

	// chart traces to be included
	ChartItems DashboardBargaugeChartChartItemArrayOutput `pulumi:"chartItems"`
	Collection pulumi.StringPtrOutput                     `pulumi:"collection"`
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
	// bar gauge max limit
	MaxLimit pulumi.IntPtrOutput `pulumi:"maxLimit"`
	// minimum chart height
	MinHeight pulumi.IntPtrOutput `pulumi:"minHeight"`
	// minimum chart width
	MinWidth pulumi.IntPtrOutput `pulumi:"minWidth"`
	// name of the chart
	Name pulumi.StringOutput `pulumi:"name"`
	// number of decimals
	NumberOfDecimals pulumi.IntPtrOutput `pulumi:"numberOfDecimals"`
	// orientation
	Orientation pulumi.StringPtrOutput `pulumi:"orientation"`
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
	Thresholds DashboardBargaugeChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardBargaugeChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardBargaugeChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardBargaugeChart(ctx *pulumi.Context,
	name string, args *DashboardBargaugeChartArgs, opts ...pulumi.ResourceOption) (*DashboardBargaugeChart, error) {
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
	var resource DashboardBargaugeChart
	err := ctx.RegisterResource("splight:index/dashboardBargaugeChart:DashboardBargaugeChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardBargaugeChart gets an existing DashboardBargaugeChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardBargaugeChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardBargaugeChartState, opts ...pulumi.ResourceOption) (*DashboardBargaugeChart, error) {
	var resource DashboardBargaugeChart
	err := ctx.ReadResource("splight:index/dashboardBargaugeChart:DashboardBargaugeChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardBargaugeChart resources.
type dashboardBargaugeChartState struct {
	// chart traces to be included
	ChartItems []DashboardBargaugeChartChartItem `pulumi:"chartItems"`
	Collection *string                           `pulumi:"collection"`
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
	// bar gauge max limit
	MaxLimit *int `pulumi:"maxLimit"`
	// minimum chart height
	MinHeight *int `pulumi:"minHeight"`
	// minimum chart width
	MinWidth *int `pulumi:"minWidth"`
	// name of the chart
	Name *string `pulumi:"name"`
	// number of decimals
	NumberOfDecimals *int `pulumi:"numberOfDecimals"`
	// orientation
	Orientation *string `pulumi:"orientation"`
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
	Thresholds []DashboardBargaugeChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardBargaugeChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardBargaugeChartState struct {
	// chart traces to be included
	ChartItems DashboardBargaugeChartChartItemArrayInput
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
	// bar gauge max limit
	MaxLimit pulumi.IntPtrInput
	// minimum chart height
	MinHeight pulumi.IntPtrInput
	// minimum chart width
	MinWidth pulumi.IntPtrInput
	// name of the chart
	Name pulumi.StringPtrInput
	// number of decimals
	NumberOfDecimals pulumi.IntPtrInput
	// orientation
	Orientation pulumi.StringPtrInput
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
	Thresholds DashboardBargaugeChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardBargaugeChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardBargaugeChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardBargaugeChartState)(nil)).Elem()
}

type dashboardBargaugeChartArgs struct {
	// chart traces to be included
	ChartItems []DashboardBargaugeChartChartItem `pulumi:"chartItems"`
	Collection *string                           `pulumi:"collection"`
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
	// bar gauge max limit
	MaxLimit *int `pulumi:"maxLimit"`
	// minimum chart height
	MinHeight *int `pulumi:"minHeight"`
	// minimum chart width
	MinWidth *int `pulumi:"minWidth"`
	// name of the chart
	Name *string `pulumi:"name"`
	// number of decimals
	NumberOfDecimals *int `pulumi:"numberOfDecimals"`
	// orientation
	Orientation *string `pulumi:"orientation"`
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
	Thresholds []DashboardBargaugeChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardBargaugeChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardBargaugeChart resource.
type DashboardBargaugeChartArgs struct {
	// chart traces to be included
	ChartItems DashboardBargaugeChartChartItemArrayInput
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
	// bar gauge max limit
	MaxLimit pulumi.IntPtrInput
	// minimum chart height
	MinHeight pulumi.IntPtrInput
	// minimum chart width
	MinWidth pulumi.IntPtrInput
	// name of the chart
	Name pulumi.StringPtrInput
	// number of decimals
	NumberOfDecimals pulumi.IntPtrInput
	// orientation
	Orientation pulumi.StringPtrInput
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
	Thresholds DashboardBargaugeChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardBargaugeChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardBargaugeChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardBargaugeChartArgs)(nil)).Elem()
}

type DashboardBargaugeChartInput interface {
	pulumi.Input

	ToDashboardBargaugeChartOutput() DashboardBargaugeChartOutput
	ToDashboardBargaugeChartOutputWithContext(ctx context.Context) DashboardBargaugeChartOutput
}

func (*DashboardBargaugeChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardBargaugeChart)(nil)).Elem()
}

func (i *DashboardBargaugeChart) ToDashboardBargaugeChartOutput() DashboardBargaugeChartOutput {
	return i.ToDashboardBargaugeChartOutputWithContext(context.Background())
}

func (i *DashboardBargaugeChart) ToDashboardBargaugeChartOutputWithContext(ctx context.Context) DashboardBargaugeChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardBargaugeChartOutput)
}

// DashboardBargaugeChartArrayInput is an input type that accepts DashboardBargaugeChartArray and DashboardBargaugeChartArrayOutput values.
// You can construct a concrete instance of `DashboardBargaugeChartArrayInput` via:
//
//	DashboardBargaugeChartArray{ DashboardBargaugeChartArgs{...} }
type DashboardBargaugeChartArrayInput interface {
	pulumi.Input

	ToDashboardBargaugeChartArrayOutput() DashboardBargaugeChartArrayOutput
	ToDashboardBargaugeChartArrayOutputWithContext(context.Context) DashboardBargaugeChartArrayOutput
}

type DashboardBargaugeChartArray []DashboardBargaugeChartInput

func (DashboardBargaugeChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardBargaugeChart)(nil)).Elem()
}

func (i DashboardBargaugeChartArray) ToDashboardBargaugeChartArrayOutput() DashboardBargaugeChartArrayOutput {
	return i.ToDashboardBargaugeChartArrayOutputWithContext(context.Background())
}

func (i DashboardBargaugeChartArray) ToDashboardBargaugeChartArrayOutputWithContext(ctx context.Context) DashboardBargaugeChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardBargaugeChartArrayOutput)
}

// DashboardBargaugeChartMapInput is an input type that accepts DashboardBargaugeChartMap and DashboardBargaugeChartMapOutput values.
// You can construct a concrete instance of `DashboardBargaugeChartMapInput` via:
//
//	DashboardBargaugeChartMap{ "key": DashboardBargaugeChartArgs{...} }
type DashboardBargaugeChartMapInput interface {
	pulumi.Input

	ToDashboardBargaugeChartMapOutput() DashboardBargaugeChartMapOutput
	ToDashboardBargaugeChartMapOutputWithContext(context.Context) DashboardBargaugeChartMapOutput
}

type DashboardBargaugeChartMap map[string]DashboardBargaugeChartInput

func (DashboardBargaugeChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardBargaugeChart)(nil)).Elem()
}

func (i DashboardBargaugeChartMap) ToDashboardBargaugeChartMapOutput() DashboardBargaugeChartMapOutput {
	return i.ToDashboardBargaugeChartMapOutputWithContext(context.Background())
}

func (i DashboardBargaugeChartMap) ToDashboardBargaugeChartMapOutputWithContext(ctx context.Context) DashboardBargaugeChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardBargaugeChartMapOutput)
}

type DashboardBargaugeChartOutput struct{ *pulumi.OutputState }

func (DashboardBargaugeChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardBargaugeChart)(nil)).Elem()
}

func (o DashboardBargaugeChartOutput) ToDashboardBargaugeChartOutput() DashboardBargaugeChartOutput {
	return o
}

func (o DashboardBargaugeChartOutput) ToDashboardBargaugeChartOutputWithContext(ctx context.Context) DashboardBargaugeChartOutput {
	return o
}

// chart traces to be included
func (o DashboardBargaugeChartOutput) ChartItems() DashboardBargaugeChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) DashboardBargaugeChartChartItemArrayOutput { return v.ChartItems }).(DashboardBargaugeChartChartItemArrayOutput)
}

func (o DashboardBargaugeChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardBargaugeChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardBargaugeChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// chart height in px
func (o DashboardBargaugeChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardBargaugeChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardBargaugeChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardBargaugeChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// bar gauge max limit
func (o DashboardBargaugeChartOutput) MaxLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.MaxLimit }).(pulumi.IntPtrOutput)
}

// minimum chart height
func (o DashboardBargaugeChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardBargaugeChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardBargaugeChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// number of decimals
func (o DashboardBargaugeChartOutput) NumberOfDecimals() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.NumberOfDecimals }).(pulumi.IntPtrOutput)
}

// orientation
func (o DashboardBargaugeChartOutput) Orientation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.Orientation }).(pulumi.StringPtrOutput)
}

// chart x position
func (o DashboardBargaugeChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardBargaugeChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardBargaugeChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardBargaugeChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardBargaugeChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardBargaugeChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardBargaugeChartOutput) Thresholds() DashboardBargaugeChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) DashboardBargaugeChartThresholdArrayOutput { return v.Thresholds }).(DashboardBargaugeChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardBargaugeChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardBargaugeChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardBargaugeChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardBargaugeChartOutput) ValueMappings() DashboardBargaugeChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) DashboardBargaugeChartValueMappingArrayOutput { return v.ValueMappings }).(DashboardBargaugeChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardBargaugeChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardBargaugeChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardBargaugeChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardBargaugeChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardBargaugeChart)(nil)).Elem()
}

func (o DashboardBargaugeChartArrayOutput) ToDashboardBargaugeChartArrayOutput() DashboardBargaugeChartArrayOutput {
	return o
}

func (o DashboardBargaugeChartArrayOutput) ToDashboardBargaugeChartArrayOutputWithContext(ctx context.Context) DashboardBargaugeChartArrayOutput {
	return o
}

func (o DashboardBargaugeChartArrayOutput) Index(i pulumi.IntInput) DashboardBargaugeChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardBargaugeChart {
		return vs[0].([]*DashboardBargaugeChart)[vs[1].(int)]
	}).(DashboardBargaugeChartOutput)
}

type DashboardBargaugeChartMapOutput struct{ *pulumi.OutputState }

func (DashboardBargaugeChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardBargaugeChart)(nil)).Elem()
}

func (o DashboardBargaugeChartMapOutput) ToDashboardBargaugeChartMapOutput() DashboardBargaugeChartMapOutput {
	return o
}

func (o DashboardBargaugeChartMapOutput) ToDashboardBargaugeChartMapOutputWithContext(ctx context.Context) DashboardBargaugeChartMapOutput {
	return o
}

func (o DashboardBargaugeChartMapOutput) MapIndex(k pulumi.StringInput) DashboardBargaugeChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardBargaugeChart {
		return vs[0].(map[string]*DashboardBargaugeChart)[vs[1].(string)]
	}).(DashboardBargaugeChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardBargaugeChartInput)(nil)).Elem(), &DashboardBargaugeChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardBargaugeChartArrayInput)(nil)).Elem(), DashboardBargaugeChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardBargaugeChartMapInput)(nil)).Elem(), DashboardBargaugeChartMap{})
	pulumi.RegisterOutputType(DashboardBargaugeChartOutput{})
	pulumi.RegisterOutputType(DashboardBargaugeChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardBargaugeChartMapOutput{})
}
