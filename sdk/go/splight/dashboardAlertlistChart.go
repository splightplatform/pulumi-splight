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
//			_, err = splight.NewDashboardAlertlistChart(ctx, "dashboardChartTest", &splight.DashboardAlertlistChartArgs{
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
//				FilterName:        pulumi.String("some name"),
//				FilterStatuses: pulumi.StringArray{
//					pulumi.String("healthy"),
//				},
//				AlertListType: pulumi.String("table"),
//				ChartItems: splight.DashboardAlertlistChartChartItemArray{
//					&splight.DashboardAlertlistChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardAlertlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardAlertlistChartChartItemQueryFilterAttributeArgs{
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
//					&splight.DashboardAlertlistChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardAlertlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardAlertlistChartChartItemQueryFilterAttributeArgs{
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
//				Thresholds: splight.DashboardAlertlistChartThresholdArray{
//					&splight.DashboardAlertlistChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardAlertlistChartValueMappingArray{
//					&splight.DashboardAlertlistChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardAlertlistChart:DashboardAlertlistChart [options] splight_dashboard_alertlist_chart.<name> <dashboard_chart_id>
// ```
type DashboardAlertlistChart struct {
	pulumi.CustomResourceState

	// alert list type
	AlertListType pulumi.StringPtrOutput `pulumi:"alertListType"`
	// chart traces to be included
	ChartItems DashboardAlertlistChartChartItemArrayOutput `pulumi:"chartItems"`
	Collection pulumi.StringPtrOutput                      `pulumi:"collection"`
	// chart description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrOutput `pulumi:"displayTimeRange"`
	// filter name
	FilterName pulumi.StringPtrOutput `pulumi:"filterName"`
	// filter status list
	FilterStatuses pulumi.StringArrayOutput `pulumi:"filterStatuses"`
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
	Thresholds DashboardAlertlistChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardAlertlistChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardAlertlistChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardAlertlistChart(ctx *pulumi.Context,
	name string, args *DashboardAlertlistChartArgs, opts ...pulumi.ResourceOption) (*DashboardAlertlistChart, error) {
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
	var resource DashboardAlertlistChart
	err := ctx.RegisterResource("splight:index/dashboardAlertlistChart:DashboardAlertlistChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardAlertlistChart gets an existing DashboardAlertlistChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardAlertlistChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardAlertlistChartState, opts ...pulumi.ResourceOption) (*DashboardAlertlistChart, error) {
	var resource DashboardAlertlistChart
	err := ctx.ReadResource("splight:index/dashboardAlertlistChart:DashboardAlertlistChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardAlertlistChart resources.
type dashboardAlertlistChartState struct {
	// alert list type
	AlertListType *string `pulumi:"alertListType"`
	// chart traces to be included
	ChartItems []DashboardAlertlistChartChartItem `pulumi:"chartItems"`
	Collection *string                            `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// filter name
	FilterName *string `pulumi:"filterName"`
	// filter status list
	FilterStatuses []string `pulumi:"filterStatuses"`
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
	Thresholds []DashboardAlertlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardAlertlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardAlertlistChartState struct {
	// alert list type
	AlertListType pulumi.StringPtrInput
	// chart traces to be included
	ChartItems DashboardAlertlistChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// filter name
	FilterName pulumi.StringPtrInput
	// filter status list
	FilterStatuses pulumi.StringArrayInput
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
	Thresholds DashboardAlertlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardAlertlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardAlertlistChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardAlertlistChartState)(nil)).Elem()
}

type dashboardAlertlistChartArgs struct {
	// alert list type
	AlertListType *string `pulumi:"alertListType"`
	// chart traces to be included
	ChartItems []DashboardAlertlistChartChartItem `pulumi:"chartItems"`
	Collection *string                            `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// filter name
	FilterName *string `pulumi:"filterName"`
	// filter status list
	FilterStatuses []string `pulumi:"filterStatuses"`
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
	Thresholds []DashboardAlertlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardAlertlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardAlertlistChart resource.
type DashboardAlertlistChartArgs struct {
	// alert list type
	AlertListType pulumi.StringPtrInput
	// chart traces to be included
	ChartItems DashboardAlertlistChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// filter name
	FilterName pulumi.StringPtrInput
	// filter status list
	FilterStatuses pulumi.StringArrayInput
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
	Thresholds DashboardAlertlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardAlertlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardAlertlistChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardAlertlistChartArgs)(nil)).Elem()
}

type DashboardAlertlistChartInput interface {
	pulumi.Input

	ToDashboardAlertlistChartOutput() DashboardAlertlistChartOutput
	ToDashboardAlertlistChartOutputWithContext(ctx context.Context) DashboardAlertlistChartOutput
}

func (*DashboardAlertlistChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardAlertlistChart)(nil)).Elem()
}

func (i *DashboardAlertlistChart) ToDashboardAlertlistChartOutput() DashboardAlertlistChartOutput {
	return i.ToDashboardAlertlistChartOutputWithContext(context.Background())
}

func (i *DashboardAlertlistChart) ToDashboardAlertlistChartOutputWithContext(ctx context.Context) DashboardAlertlistChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAlertlistChartOutput)
}

// DashboardAlertlistChartArrayInput is an input type that accepts DashboardAlertlistChartArray and DashboardAlertlistChartArrayOutput values.
// You can construct a concrete instance of `DashboardAlertlistChartArrayInput` via:
//
//	DashboardAlertlistChartArray{ DashboardAlertlistChartArgs{...} }
type DashboardAlertlistChartArrayInput interface {
	pulumi.Input

	ToDashboardAlertlistChartArrayOutput() DashboardAlertlistChartArrayOutput
	ToDashboardAlertlistChartArrayOutputWithContext(context.Context) DashboardAlertlistChartArrayOutput
}

type DashboardAlertlistChartArray []DashboardAlertlistChartInput

func (DashboardAlertlistChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardAlertlistChart)(nil)).Elem()
}

func (i DashboardAlertlistChartArray) ToDashboardAlertlistChartArrayOutput() DashboardAlertlistChartArrayOutput {
	return i.ToDashboardAlertlistChartArrayOutputWithContext(context.Background())
}

func (i DashboardAlertlistChartArray) ToDashboardAlertlistChartArrayOutputWithContext(ctx context.Context) DashboardAlertlistChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAlertlistChartArrayOutput)
}

// DashboardAlertlistChartMapInput is an input type that accepts DashboardAlertlistChartMap and DashboardAlertlistChartMapOutput values.
// You can construct a concrete instance of `DashboardAlertlistChartMapInput` via:
//
//	DashboardAlertlistChartMap{ "key": DashboardAlertlistChartArgs{...} }
type DashboardAlertlistChartMapInput interface {
	pulumi.Input

	ToDashboardAlertlistChartMapOutput() DashboardAlertlistChartMapOutput
	ToDashboardAlertlistChartMapOutputWithContext(context.Context) DashboardAlertlistChartMapOutput
}

type DashboardAlertlistChartMap map[string]DashboardAlertlistChartInput

func (DashboardAlertlistChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardAlertlistChart)(nil)).Elem()
}

func (i DashboardAlertlistChartMap) ToDashboardAlertlistChartMapOutput() DashboardAlertlistChartMapOutput {
	return i.ToDashboardAlertlistChartMapOutputWithContext(context.Background())
}

func (i DashboardAlertlistChartMap) ToDashboardAlertlistChartMapOutputWithContext(ctx context.Context) DashboardAlertlistChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAlertlistChartMapOutput)
}

type DashboardAlertlistChartOutput struct{ *pulumi.OutputState }

func (DashboardAlertlistChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardAlertlistChart)(nil)).Elem()
}

func (o DashboardAlertlistChartOutput) ToDashboardAlertlistChartOutput() DashboardAlertlistChartOutput {
	return o
}

func (o DashboardAlertlistChartOutput) ToDashboardAlertlistChartOutputWithContext(ctx context.Context) DashboardAlertlistChartOutput {
	return o
}

// alert list type
func (o DashboardAlertlistChartOutput) AlertListType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.AlertListType }).(pulumi.StringPtrOutput)
}

// chart traces to be included
func (o DashboardAlertlistChartOutput) ChartItems() DashboardAlertlistChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) DashboardAlertlistChartChartItemArrayOutput { return v.ChartItems }).(DashboardAlertlistChartChartItemArrayOutput)
}

func (o DashboardAlertlistChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardAlertlistChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardAlertlistChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// filter name
func (o DashboardAlertlistChartOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.FilterName }).(pulumi.StringPtrOutput)
}

// filter status list
func (o DashboardAlertlistChartOutput) FilterStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringArrayOutput { return v.FilterStatuses }).(pulumi.StringArrayOutput)
}

// chart height in px
func (o DashboardAlertlistChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardAlertlistChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardAlertlistChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardAlertlistChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardAlertlistChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardAlertlistChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardAlertlistChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// chart x position
func (o DashboardAlertlistChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardAlertlistChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardAlertlistChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardAlertlistChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardAlertlistChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardAlertlistChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardAlertlistChartOutput) Thresholds() DashboardAlertlistChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) DashboardAlertlistChartThresholdArrayOutput { return v.Thresholds }).(DashboardAlertlistChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardAlertlistChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardAlertlistChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardAlertlistChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardAlertlistChartOutput) ValueMappings() DashboardAlertlistChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) DashboardAlertlistChartValueMappingArrayOutput {
		return v.ValueMappings
	}).(DashboardAlertlistChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardAlertlistChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlertlistChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardAlertlistChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardAlertlistChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardAlertlistChart)(nil)).Elem()
}

func (o DashboardAlertlistChartArrayOutput) ToDashboardAlertlistChartArrayOutput() DashboardAlertlistChartArrayOutput {
	return o
}

func (o DashboardAlertlistChartArrayOutput) ToDashboardAlertlistChartArrayOutputWithContext(ctx context.Context) DashboardAlertlistChartArrayOutput {
	return o
}

func (o DashboardAlertlistChartArrayOutput) Index(i pulumi.IntInput) DashboardAlertlistChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardAlertlistChart {
		return vs[0].([]*DashboardAlertlistChart)[vs[1].(int)]
	}).(DashboardAlertlistChartOutput)
}

type DashboardAlertlistChartMapOutput struct{ *pulumi.OutputState }

func (DashboardAlertlistChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardAlertlistChart)(nil)).Elem()
}

func (o DashboardAlertlistChartMapOutput) ToDashboardAlertlistChartMapOutput() DashboardAlertlistChartMapOutput {
	return o
}

func (o DashboardAlertlistChartMapOutput) ToDashboardAlertlistChartMapOutputWithContext(ctx context.Context) DashboardAlertlistChartMapOutput {
	return o
}

func (o DashboardAlertlistChartMapOutput) MapIndex(k pulumi.StringInput) DashboardAlertlistChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardAlertlistChart {
		return vs[0].(map[string]*DashboardAlertlistChart)[vs[1].(string)]
	}).(DashboardAlertlistChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAlertlistChartInput)(nil)).Elem(), &DashboardAlertlistChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAlertlistChartArrayInput)(nil)).Elem(), DashboardAlertlistChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAlertlistChartMapInput)(nil)).Elem(), DashboardAlertlistChartMap{})
	pulumi.RegisterOutputType(DashboardAlertlistChartOutput{})
	pulumi.RegisterOutputType(DashboardAlertlistChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardAlertlistChartMapOutput{})
}
