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
//			_, err = splight.NewDashboardAssetlistChart(ctx, "dashboardChartTest", &splight.DashboardAssetlistChartArgs{
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
//				AssetListType: pulumi.String("table"),
//				ChartItems: splight.DashboardAssetlistChartChartItemArray{
//					&splight.DashboardAssetlistChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardAssetlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardAssetlistChartChartItemQueryFilterAttributeArgs{
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
//					&splight.DashboardAssetlistChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardAssetlistChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardAssetlistChartChartItemQueryFilterAttributeArgs{
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
//				Thresholds: splight.DashboardAssetlistChartThresholdArray{
//					&splight.DashboardAssetlistChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardAssetlistChartValueMappingArray{
//					&splight.DashboardAssetlistChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardAssetlistChart:DashboardAssetlistChart [options] splight_dashboard_assetlist_chart.<name> <dashboard_chart_id>
// ```
type DashboardAssetlistChart struct {
	pulumi.CustomResourceState

	// asset list type
	AssetListType pulumi.StringPtrOutput `pulumi:"assetListType"`
	// chart traces to be included
	ChartItems DashboardAssetlistChartChartItemArrayOutput `pulumi:"chartItems"`
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
	Thresholds DashboardAssetlistChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardAssetlistChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardAssetlistChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardAssetlistChart(ctx *pulumi.Context,
	name string, args *DashboardAssetlistChartArgs, opts ...pulumi.ResourceOption) (*DashboardAssetlistChart, error) {
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
	var resource DashboardAssetlistChart
	err := ctx.RegisterResource("splight:index/dashboardAssetlistChart:DashboardAssetlistChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardAssetlistChart gets an existing DashboardAssetlistChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardAssetlistChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardAssetlistChartState, opts ...pulumi.ResourceOption) (*DashboardAssetlistChart, error) {
	var resource DashboardAssetlistChart
	err := ctx.ReadResource("splight:index/dashboardAssetlistChart:DashboardAssetlistChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardAssetlistChart resources.
type dashboardAssetlistChartState struct {
	// asset list type
	AssetListType *string `pulumi:"assetListType"`
	// chart traces to be included
	ChartItems []DashboardAssetlistChartChartItem `pulumi:"chartItems"`
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
	Thresholds []DashboardAssetlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardAssetlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardAssetlistChartState struct {
	// asset list type
	AssetListType pulumi.StringPtrInput
	// chart traces to be included
	ChartItems DashboardAssetlistChartChartItemArrayInput
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
	Thresholds DashboardAssetlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardAssetlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardAssetlistChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardAssetlistChartState)(nil)).Elem()
}

type dashboardAssetlistChartArgs struct {
	// asset list type
	AssetListType *string `pulumi:"assetListType"`
	// chart traces to be included
	ChartItems []DashboardAssetlistChartChartItem `pulumi:"chartItems"`
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
	Thresholds []DashboardAssetlistChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardAssetlistChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardAssetlistChart resource.
type DashboardAssetlistChartArgs struct {
	// asset list type
	AssetListType pulumi.StringPtrInput
	// chart traces to be included
	ChartItems DashboardAssetlistChartChartItemArrayInput
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
	Thresholds DashboardAssetlistChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardAssetlistChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardAssetlistChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardAssetlistChartArgs)(nil)).Elem()
}

type DashboardAssetlistChartInput interface {
	pulumi.Input

	ToDashboardAssetlistChartOutput() DashboardAssetlistChartOutput
	ToDashboardAssetlistChartOutputWithContext(ctx context.Context) DashboardAssetlistChartOutput
}

func (*DashboardAssetlistChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardAssetlistChart)(nil)).Elem()
}

func (i *DashboardAssetlistChart) ToDashboardAssetlistChartOutput() DashboardAssetlistChartOutput {
	return i.ToDashboardAssetlistChartOutputWithContext(context.Background())
}

func (i *DashboardAssetlistChart) ToDashboardAssetlistChartOutputWithContext(ctx context.Context) DashboardAssetlistChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAssetlistChartOutput)
}

// DashboardAssetlistChartArrayInput is an input type that accepts DashboardAssetlistChartArray and DashboardAssetlistChartArrayOutput values.
// You can construct a concrete instance of `DashboardAssetlistChartArrayInput` via:
//
//	DashboardAssetlistChartArray{ DashboardAssetlistChartArgs{...} }
type DashboardAssetlistChartArrayInput interface {
	pulumi.Input

	ToDashboardAssetlistChartArrayOutput() DashboardAssetlistChartArrayOutput
	ToDashboardAssetlistChartArrayOutputWithContext(context.Context) DashboardAssetlistChartArrayOutput
}

type DashboardAssetlistChartArray []DashboardAssetlistChartInput

func (DashboardAssetlistChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardAssetlistChart)(nil)).Elem()
}

func (i DashboardAssetlistChartArray) ToDashboardAssetlistChartArrayOutput() DashboardAssetlistChartArrayOutput {
	return i.ToDashboardAssetlistChartArrayOutputWithContext(context.Background())
}

func (i DashboardAssetlistChartArray) ToDashboardAssetlistChartArrayOutputWithContext(ctx context.Context) DashboardAssetlistChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAssetlistChartArrayOutput)
}

// DashboardAssetlistChartMapInput is an input type that accepts DashboardAssetlistChartMap and DashboardAssetlistChartMapOutput values.
// You can construct a concrete instance of `DashboardAssetlistChartMapInput` via:
//
//	DashboardAssetlistChartMap{ "key": DashboardAssetlistChartArgs{...} }
type DashboardAssetlistChartMapInput interface {
	pulumi.Input

	ToDashboardAssetlistChartMapOutput() DashboardAssetlistChartMapOutput
	ToDashboardAssetlistChartMapOutputWithContext(context.Context) DashboardAssetlistChartMapOutput
}

type DashboardAssetlistChartMap map[string]DashboardAssetlistChartInput

func (DashboardAssetlistChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardAssetlistChart)(nil)).Elem()
}

func (i DashboardAssetlistChartMap) ToDashboardAssetlistChartMapOutput() DashboardAssetlistChartMapOutput {
	return i.ToDashboardAssetlistChartMapOutputWithContext(context.Background())
}

func (i DashboardAssetlistChartMap) ToDashboardAssetlistChartMapOutputWithContext(ctx context.Context) DashboardAssetlistChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardAssetlistChartMapOutput)
}

type DashboardAssetlistChartOutput struct{ *pulumi.OutputState }

func (DashboardAssetlistChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardAssetlistChart)(nil)).Elem()
}

func (o DashboardAssetlistChartOutput) ToDashboardAssetlistChartOutput() DashboardAssetlistChartOutput {
	return o
}

func (o DashboardAssetlistChartOutput) ToDashboardAssetlistChartOutputWithContext(ctx context.Context) DashboardAssetlistChartOutput {
	return o
}

// asset list type
func (o DashboardAssetlistChartOutput) AssetListType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.AssetListType }).(pulumi.StringPtrOutput)
}

// chart traces to be included
func (o DashboardAssetlistChartOutput) ChartItems() DashboardAssetlistChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) DashboardAssetlistChartChartItemArrayOutput { return v.ChartItems }).(DashboardAssetlistChartChartItemArrayOutput)
}

func (o DashboardAssetlistChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardAssetlistChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardAssetlistChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// filter name
func (o DashboardAssetlistChartOutput) FilterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.FilterName }).(pulumi.StringPtrOutput)
}

// filter status list
func (o DashboardAssetlistChartOutput) FilterStatuses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringArrayOutput { return v.FilterStatuses }).(pulumi.StringArrayOutput)
}

// chart height in px
func (o DashboardAssetlistChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardAssetlistChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardAssetlistChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardAssetlistChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardAssetlistChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardAssetlistChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardAssetlistChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// chart x position
func (o DashboardAssetlistChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardAssetlistChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardAssetlistChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardAssetlistChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardAssetlistChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardAssetlistChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardAssetlistChartOutput) Thresholds() DashboardAssetlistChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) DashboardAssetlistChartThresholdArrayOutput { return v.Thresholds }).(DashboardAssetlistChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardAssetlistChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardAssetlistChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardAssetlistChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardAssetlistChartOutput) ValueMappings() DashboardAssetlistChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) DashboardAssetlistChartValueMappingArrayOutput {
		return v.ValueMappings
	}).(DashboardAssetlistChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardAssetlistChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardAssetlistChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardAssetlistChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardAssetlistChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardAssetlistChart)(nil)).Elem()
}

func (o DashboardAssetlistChartArrayOutput) ToDashboardAssetlistChartArrayOutput() DashboardAssetlistChartArrayOutput {
	return o
}

func (o DashboardAssetlistChartArrayOutput) ToDashboardAssetlistChartArrayOutputWithContext(ctx context.Context) DashboardAssetlistChartArrayOutput {
	return o
}

func (o DashboardAssetlistChartArrayOutput) Index(i pulumi.IntInput) DashboardAssetlistChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardAssetlistChart {
		return vs[0].([]*DashboardAssetlistChart)[vs[1].(int)]
	}).(DashboardAssetlistChartOutput)
}

type DashboardAssetlistChartMapOutput struct{ *pulumi.OutputState }

func (DashboardAssetlistChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardAssetlistChart)(nil)).Elem()
}

func (o DashboardAssetlistChartMapOutput) ToDashboardAssetlistChartMapOutput() DashboardAssetlistChartMapOutput {
	return o
}

func (o DashboardAssetlistChartMapOutput) ToDashboardAssetlistChartMapOutputWithContext(ctx context.Context) DashboardAssetlistChartMapOutput {
	return o
}

func (o DashboardAssetlistChartMapOutput) MapIndex(k pulumi.StringInput) DashboardAssetlistChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardAssetlistChart {
		return vs[0].(map[string]*DashboardAssetlistChart)[vs[1].(string)]
	}).(DashboardAssetlistChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAssetlistChartInput)(nil)).Elem(), &DashboardAssetlistChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAssetlistChartArrayInput)(nil)).Elem(), DashboardAssetlistChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardAssetlistChartMapInput)(nil)).Elem(), DashboardAssetlistChartMap{})
	pulumi.RegisterOutputType(DashboardAssetlistChartOutput{})
	pulumi.RegisterOutputType(DashboardAssetlistChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardAssetlistChartMapOutput{})
}
