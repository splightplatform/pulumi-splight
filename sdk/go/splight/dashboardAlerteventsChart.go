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
//			_, err = splight.NewDashboardAlerteventsChart(ctx, "dashboardChartTest", &splight.DashboardAlerteventsChartArgs{
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
//				FilterOldStatuses: pulumi.StringArray{
//					pulumi.String("warning"),
//				},
//				FilterNewStatuses: pulumi.StringArray{
//					pulumi.String("no_alert"),
//					pulumi.String("warning"),
//				},
//				ChartItems: splight.DashboardAlerteventsChartChartItemArray{
//					&splight.DashboardAlerteventsChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardAlerteventsChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardAlerteventsChartChartItemQueryFilterAttributeArgs{
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
//					&splight.DashboardAlerteventsChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardAlerteventsChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardAlerteventsChartChartItemQueryFilterAttributeArgs{
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
//				Thresholds: splight.DashboardAlerteventsChartThresholdArray{
//					&splight.DashboardAlerteventsChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardAlerteventsChartValueMappingArray{
//					&splight.DashboardAlerteventsChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardAlerteventsChart:DashboardAlerteventsChart [options] splight_dashboard_alertevents_chart.<name> <dashboard_chart_id>
// ```
type DashboardAlerteventsChart struct {
	pulumi.CustomResourceState

	// chart traces to be included
	ChartItems DashboardAlerteventsChartChartItemArrayOutput `pulumi:"chartItems"`
	Collection pulumi.StringPtrOutput                        `pulumi:"collection"`
	// chart description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrOutput `pulumi:"displayTimeRange"`
	// filter name
	FilterName pulumi.StringPtrOutput `pulumi:"filterName"`
	// filter new status
	FilterNewStatuses pulumi.StringArrayOutput `pulumi:"filterNewStatuses"`
	// filter old status
	FilterOldStatuses pulumi.StringArrayOutput `pulumi:"filterOldStatuses"`
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
	Thresholds DashboardAlerteventsChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardAlerteventsChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardAlerteventsChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardAlerteventsChart(ctx *pulumi.Context,
	name string, args *DashboardAlerteventsChartArgs, opts ...pulumi.ResourceOption) (*DashboardAlerteventsChart, error) {
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
	var resource DashboardAlerteventsChart
	err := ctx.RegisterResource("splight:index/dashboardAlerteventsChart:DashboardAlerteventsChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardAlerteventsChart gets an existing DashboardAlerteventsChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardAlerteventsChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardAlerteventsChartState, opts ...pulumi.ResourceOption) (*DashboardAlerteventsChart, error) {
	var resource DashboardAlerteventsChart
	err := ctx.ReadResource("splight:index/dashboardAlerteventsChart:DashboardAlerteventsChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardAlerteventsChart resources.
type dashboardAlerteventsChartState struct {
	// chart traces to be included
	ChartItems []DashboardAlerteventsChartChartItem `pulumi:"chartItems"`
	Collection *string                              `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// filter name
	FilterName *string `pulumi:"filterName"`
	// filter new status
	FilterNewStatuses []string `pulumi:"filterNewStatuses"`
	// filter old status
	FilterOldStatuses []string `pulumi:"filterOldStatuses"`
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
	Thresholds []DashboardAlerteventsChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardAlerteventsChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardAlerteventsChartState struct {
	// chart traces to be included
	ChartItems DashboardAlerteventsChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// filter name
	FilterName pulumi.StringPtrInput
	// filter new status
	FilterNewStatuses pulumi.StringArrayInput
	// filter old status
	FilterOldStatuses pulumi.StringArrayInput
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
	Thresholds DashboardAlerteventsChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardAlerteventsChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardAlerteventsChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardAlerteventsChartState)(nil)).Elem()
}

type dashboardAlerteventsChartArgs struct {
	// chart traces to be included
	ChartItems []DashboardAlerteventsChartChartItem `pulumi:"chartItems"`
	Collection *string                              `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// filter name
	FilterName *string `pulumi:"filterName"`
	// filter new status
	FilterNewStatuses []string `pulumi:"filterNewStatuses"`
	// filter old status
	FilterOldStatuses []string `pulumi:"filterOldStatuses"`
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
	Thresholds []DashboardAlerteventsChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardAlerteventsChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardAlerteventsChart resource.
type DashboardAlerteventsChartArgs struct {
	// chart traces to be included
	ChartItems DashboardAlerteventsChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// filter name
	FilterName pulumi.StringPtrInput
	// filter new status
	FilterNewStatuses pulumi.StringArrayInput
	// filter old status
	FilterOldStatuses pulumi.StringArrayInput
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
	Thresholds DashboardAlerteventsChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardAlerteventsChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardAlerteventsChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardAlerteventsChartArgs)(nil)).Elem()
}

type DashboardAlerteventsChartInput interface {
	pulumi.Input

	ToDashboardAlerteventsChartOutput() DashboardAlerteventsChartOutput
	ToDashboardAlerteventsChartOutputWithContext(ctx context.Context) DashboardAlerteventsChartOutput
}

func (*DashboardAlerteventsChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardAlerteventsChart)(nil)).Elem()
}

func (i *DashboardAlerteventsChart) ToDashboardAlerteventsChartOutput() DashboardAlerteventsChartOutput {
	return i.ToDashboardAlerteventsChartOutputWithContext(context.Background())
}

func (i *DashboardAlerteventsChart) ToDashboardAlerteventsChartOutputWithContext(ctx context.Context) DashboardAlerteventsChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAlerteventsChartOutput)
}

// DashboardAlerteventsChartArrayInput is an input type that accepts DashboardAlerteventsChartArray and DashboardAlerteventsChartArrayOutput values.
// You can construct a concrete instance of `DashboardAlerteventsChartArrayInput` via:
//
//	DashboardAlerteventsChartArray{ DashboardAlerteventsChartArgs{...} }
type DashboardAlerteventsChartArrayInput interface {
	pulumi.Input

	ToDashboardAlerteventsChartArrayOutput() DashboardAlerteventsChartArrayOutput
	ToDashboardAlerteventsChartArrayOutputWithContext(context.Context) DashboardAlerteventsChartArrayOutput
}

type DashboardAlerteventsChartArray []DashboardAlerteventsChartInput

func (DashboardAlerteventsChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardAlerteventsChart)(nil)).Elem()
}

func (i DashboardAlerteventsChartArray) ToDashboardAlerteventsChartArrayOutput() DashboardAlerteventsChartArrayOutput {
	return i.ToDashboardAlerteventsChartArrayOutputWithContext(context.Background())
}

func (i DashboardAlerteventsChartArray) ToDashboardAlerteventsChartArrayOutputWithContext(ctx context.Context) DashboardAlerteventsChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAlerteventsChartArrayOutput)
}

// DashboardAlerteventsChartMapInput is an input type that accepts DashboardAlerteventsChartMap and DashboardAlerteventsChartMapOutput values.
// You can construct a concrete instance of `DashboardAlerteventsChartMapInput` via:
//
//	DashboardAlerteventsChartMap{ "key": DashboardAlerteventsChartArgs{...} }
type DashboardAlerteventsChartMapInput interface {
	pulumi.Input

	ToDashboardAlerteventsChartMapOutput() DashboardAlerteventsChartMapOutput
	ToDashboardAlerteventsChartMapOutputWithContext(context.Context) DashboardAlerteventsChartMapOutput
}

type DashboardAlerteventsChartMap map[string]DashboardAlerteventsChartInput

func (DashboardAlerteventsChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardAlerteventsChart)(nil)).Elem()
}

func (i DashboardAlerteventsChartMap) ToDashboardAlerteventsChartMapOutput() DashboardAlerteventsChartMapOutput {
	return i.ToDashboardAlerteventsChartMapOutputWithContext(context.Background())
}

func (i DashboardAlerteventsChartMap) ToDashboardAlerteventsChartMapOutputWithContext(ctx context.Context) DashboardAlerteventsChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAlerteventsChartMapOutput)
}

type DashboardAlerteventsChartOutput struct{ *pulumi.OutputState }

func (DashboardAlerteventsChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardAlerteventsChart)(nil)).Elem()
}

func (o DashboardAlerteventsChartOutput) ToDashboardAlerteventsChartOutput() DashboardAlerteventsChartOutput {
	return o
}

func (o DashboardAlerteventsChartOutput) ToDashboardAlerteventsChartOutputWithContext(ctx context.Context) DashboardAlerteventsChartOutput {
	return o
}

// chart traces to be included
func (o DashboardAlerteventsChartOutput) ChartItems() DashboardAlerteventsChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) DashboardAlerteventsChartChartItemArrayOutput { return v.ChartItems }).(DashboardAlerteventsChartChartItemArrayOutput)
}

func (o DashboardAlerteventsChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardAlerteventsChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardAlerteventsChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// filter name
func (o DashboardAlerteventsChartOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.FilterName }).(pulumi.StringPtrOutput)
}

// filter new status
func (o DashboardAlerteventsChartOutput) FilterNewStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringArrayOutput { return v.FilterNewStatuses }).(pulumi.StringArrayOutput)
}

// filter old status
func (o DashboardAlerteventsChartOutput) FilterOldStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringArrayOutput { return v.FilterOldStatuses }).(pulumi.StringArrayOutput)
}

// chart height in px
func (o DashboardAlerteventsChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardAlerteventsChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardAlerteventsChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardAlerteventsChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardAlerteventsChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardAlerteventsChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardAlerteventsChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// chart x position
func (o DashboardAlerteventsChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardAlerteventsChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardAlerteventsChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardAlerteventsChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardAlerteventsChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardAlerteventsChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardAlerteventsChartOutput) Thresholds() DashboardAlerteventsChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) DashboardAlerteventsChartThresholdArrayOutput { return v.Thresholds }).(DashboardAlerteventsChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardAlerteventsChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardAlerteventsChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardAlerteventsChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardAlerteventsChartOutput) ValueMappings() DashboardAlerteventsChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) DashboardAlerteventsChartValueMappingArrayOutput {
		return v.ValueMappings
	}).(DashboardAlerteventsChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardAlerteventsChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAlerteventsChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardAlerteventsChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardAlerteventsChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardAlerteventsChart)(nil)).Elem()
}

func (o DashboardAlerteventsChartArrayOutput) ToDashboardAlerteventsChartArrayOutput() DashboardAlerteventsChartArrayOutput {
	return o
}

func (o DashboardAlerteventsChartArrayOutput) ToDashboardAlerteventsChartArrayOutputWithContext(ctx context.Context) DashboardAlerteventsChartArrayOutput {
	return o
}

func (o DashboardAlerteventsChartArrayOutput) Index(i pulumi.IntInput) DashboardAlerteventsChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardAlerteventsChart {
		return vs[0].([]*DashboardAlerteventsChart)[vs[1].(int)]
	}).(DashboardAlerteventsChartOutput)
}

type DashboardAlerteventsChartMapOutput struct{ *pulumi.OutputState }

func (DashboardAlerteventsChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardAlerteventsChart)(nil)).Elem()
}

func (o DashboardAlerteventsChartMapOutput) ToDashboardAlerteventsChartMapOutput() DashboardAlerteventsChartMapOutput {
	return o
}

func (o DashboardAlerteventsChartMapOutput) ToDashboardAlerteventsChartMapOutputWithContext(ctx context.Context) DashboardAlerteventsChartMapOutput {
	return o
}

func (o DashboardAlerteventsChartMapOutput) MapIndex(k pulumi.StringInput) DashboardAlerteventsChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardAlerteventsChart {
		return vs[0].(map[string]*DashboardAlerteventsChart)[vs[1].(string)]
	}).(DashboardAlerteventsChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAlerteventsChartInput)(nil)).Elem(), &DashboardAlerteventsChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAlerteventsChartArrayInput)(nil)).Elem(), DashboardAlerteventsChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAlerteventsChartMapInput)(nil)).Elem(), DashboardAlerteventsChartMap{})
	pulumi.RegisterOutputType(DashboardAlerteventsChartOutput{})
	pulumi.RegisterOutputType(DashboardAlerteventsChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardAlerteventsChartMapOutput{})
}
