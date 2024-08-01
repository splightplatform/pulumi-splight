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
//			_, err = splight.NewDashboardHistogramChart(ctx, "dashboardChartTest", &splight.DashboardHistogramChartArgs{
//				Tab:                   dashboardTabTest.ID(),
//				TimestampGte:          pulumi.String("now - 7d"),
//				TimestampLte:          pulumi.String("now"),
//				Description:           pulumi.String("Chart description"),
//				MinHeight:             pulumi.Int(1),
//				MinWidth:              pulumi.Int(4),
//				DisplayTimeRange:      pulumi.Bool(true),
//				LabelsDisplay:         pulumi.Bool(true),
//				LabelsAggregation:     pulumi.String("last"),
//				LabelsPlacement:       pulumi.String("bottom"),
//				ShowBeyondData:        pulumi.Bool(true),
//				Height:                pulumi.Int(10),
//				Width:                 pulumi.Int(20),
//				Collection:            pulumi.String("default"),
//				NumberOfDecimals:      pulumi.Int(2),
//				BucketCount:           pulumi.Int(20),
//				BucketSize:            pulumi.Int(2),
//				HistogramType:         pulumi.String("categorical"),
//				Sorting:               pulumi.String("count"),
//				Stacked:               pulumi.Bool(true),
//				CategoriesTopMaxLimit: nil,
//				ChartItems: splight.DashboardHistogramChartChartItemArray{
//					&splight.DashboardHistogramChartChartItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Color:           pulumi.String("red"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardHistogramChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardHistogramChartChartItemQueryFilterAttributeArgs{
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
//					&splight.DashboardHistogramChartChartItemArgs{
//						RefId:           pulumi.String("B"),
//						Color:           pulumi.String("blue"),
//						Type:            pulumi.String("QUERY"),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.DashboardHistogramChartChartItemQueryFilterAssetArgs{
//							Id:   assetTest.ID(),
//							Name: assetTest.Name,
//						},
//						QueryFilterAttribute: &splight.DashboardHistogramChartChartItemQueryFilterAttributeArgs{
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
//				Thresholds: splight.DashboardHistogramChartThresholdArray{
//					&splight.DashboardHistogramChartThresholdArgs{
//						Color:       pulumi.String("#00edcf"),
//						DisplayText: pulumi.String("T1Test"),
//						Value:       pulumi.Float64(13.1),
//					},
//				},
//				ValueMappings: splight.DashboardHistogramChartValueMappingArray{
//					&splight.DashboardHistogramChartValueMappingArgs{
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
// $ pulumi import splight:index/dashboardHistogramChart:DashboardHistogramChart [options] splight_dashboard_histogram_chart.<name> <dashboard_chart_id>
// ```
type DashboardHistogramChart struct {
	pulumi.CustomResourceState

	// bucket count
	BucketCount pulumi.IntPtrOutput `pulumi:"bucketCount"`
	// bucket size
	BucketSize pulumi.IntPtrOutput `pulumi:"bucketSize"`
	// categories top max limit
	CategoriesTopMaxLimit pulumi.IntPtrOutput `pulumi:"categoriesTopMaxLimit"`
	// chart traces to be included
	ChartItems DashboardHistogramChartChartItemArrayOutput `pulumi:"chartItems"`
	Collection pulumi.StringPtrOutput                      `pulumi:"collection"`
	// chart description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrOutput `pulumi:"displayTimeRange"`
	// chart height in px
	Height pulumi.IntPtrOutput `pulumi:"height"`
	// histogram type
	HistogramType pulumi.StringPtrOutput `pulumi:"histogramType"`
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
	// sorting type
	Sorting pulumi.StringPtrOutput `pulumi:"sorting"`
	// whether to stack or not the histogram
	Stacked pulumi.BoolPtrOutput `pulumi:"stacked"`
	// id for the tab where to place the chart
	Tab pulumi.StringOutput `pulumi:"tab"`
	// optional static lines to be added to the chart as references
	Thresholds DashboardHistogramChartThresholdArrayOutput `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringOutput `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringOutput `pulumi:"timestampLte"`
	// chart timezone
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings DashboardHistogramChartValueMappingArrayOutput `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width pulumi.IntPtrOutput `pulumi:"width"`
}

// NewDashboardHistogramChart registers a new resource with the given unique name, arguments, and options.
func NewDashboardHistogramChart(ctx *pulumi.Context,
	name string, args *DashboardHistogramChartArgs, opts ...pulumi.ResourceOption) (*DashboardHistogramChart, error) {
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
	var resource DashboardHistogramChart
	err := ctx.RegisterResource("splight:index/dashboardHistogramChart:DashboardHistogramChart", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardHistogramChart gets an existing DashboardHistogramChart resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardHistogramChart(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardHistogramChartState, opts ...pulumi.ResourceOption) (*DashboardHistogramChart, error) {
	var resource DashboardHistogramChart
	err := ctx.ReadResource("splight:index/dashboardHistogramChart:DashboardHistogramChart", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardHistogramChart resources.
type dashboardHistogramChartState struct {
	// bucket count
	BucketCount *int `pulumi:"bucketCount"`
	// bucket size
	BucketSize *int `pulumi:"bucketSize"`
	// categories top max limit
	CategoriesTopMaxLimit *int `pulumi:"categoriesTopMaxLimit"`
	// chart traces to be included
	ChartItems []DashboardHistogramChartChartItem `pulumi:"chartItems"`
	Collection *string                            `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// chart height in px
	Height *int `pulumi:"height"`
	// histogram type
	HistogramType *string `pulumi:"histogramType"`
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
	// sorting type
	Sorting *string `pulumi:"sorting"`
	// whether to stack or not the histogram
	Stacked *bool `pulumi:"stacked"`
	// id for the tab where to place the chart
	Tab *string `pulumi:"tab"`
	// optional static lines to be added to the chart as references
	Thresholds []DashboardHistogramChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte *string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte *string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardHistogramChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

type DashboardHistogramChartState struct {
	// bucket count
	BucketCount pulumi.IntPtrInput
	// bucket size
	BucketSize pulumi.IntPtrInput
	// categories top max limit
	CategoriesTopMaxLimit pulumi.IntPtrInput
	// chart traces to be included
	ChartItems DashboardHistogramChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// chart height in px
	Height pulumi.IntPtrInput
	// histogram type
	HistogramType pulumi.StringPtrInput
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
	// sorting type
	Sorting pulumi.StringPtrInput
	// whether to stack or not the histogram
	Stacked pulumi.BoolPtrInput
	// id for the tab where to place the chart
	Tab pulumi.StringPtrInput
	// optional static lines to be added to the chart as references
	Thresholds DashboardHistogramChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringPtrInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringPtrInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardHistogramChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardHistogramChartState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardHistogramChartState)(nil)).Elem()
}

type dashboardHistogramChartArgs struct {
	// bucket count
	BucketCount *int `pulumi:"bucketCount"`
	// bucket size
	BucketSize *int `pulumi:"bucketSize"`
	// categories top max limit
	CategoriesTopMaxLimit *int `pulumi:"categoriesTopMaxLimit"`
	// chart traces to be included
	ChartItems []DashboardHistogramChartChartItem `pulumi:"chartItems"`
	Collection *string                            `pulumi:"collection"`
	// chart description
	Description *string `pulumi:"description"`
	// whether to display the time range or not
	DisplayTimeRange *bool `pulumi:"displayTimeRange"`
	// chart height in px
	Height *int `pulumi:"height"`
	// histogram type
	HistogramType *string `pulumi:"histogramType"`
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
	// sorting type
	Sorting *string `pulumi:"sorting"`
	// whether to stack or not the histogram
	Stacked *bool `pulumi:"stacked"`
	// id for the tab where to place the chart
	Tab string `pulumi:"tab"`
	// optional static lines to be added to the chart as references
	Thresholds []DashboardHistogramChartThreshold `pulumi:"thresholds"`
	// date in isoformat or shortcut string where to end reading
	TimestampGte string `pulumi:"timestampGte"`
	// date in isoformat or shortcut string where to start reading
	TimestampLte string `pulumi:"timestampLte"`
	// chart timezone
	Timezone *string `pulumi:"timezone"`
	// optional mappings to transform data with rules
	ValueMappings []DashboardHistogramChartValueMapping `pulumi:"valueMappings"`
	// chart width in cols (max 20)
	Width *int `pulumi:"width"`
}

// The set of arguments for constructing a DashboardHistogramChart resource.
type DashboardHistogramChartArgs struct {
	// bucket count
	BucketCount pulumi.IntPtrInput
	// bucket size
	BucketSize pulumi.IntPtrInput
	// categories top max limit
	CategoriesTopMaxLimit pulumi.IntPtrInput
	// chart traces to be included
	ChartItems DashboardHistogramChartChartItemArrayInput
	Collection pulumi.StringPtrInput
	// chart description
	Description pulumi.StringPtrInput
	// whether to display the time range or not
	DisplayTimeRange pulumi.BoolPtrInput
	// chart height in px
	Height pulumi.IntPtrInput
	// histogram type
	HistogramType pulumi.StringPtrInput
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
	// sorting type
	Sorting pulumi.StringPtrInput
	// whether to stack or not the histogram
	Stacked pulumi.BoolPtrInput
	// id for the tab where to place the chart
	Tab pulumi.StringInput
	// optional static lines to be added to the chart as references
	Thresholds DashboardHistogramChartThresholdArrayInput
	// date in isoformat or shortcut string where to end reading
	TimestampGte pulumi.StringInput
	// date in isoformat or shortcut string where to start reading
	TimestampLte pulumi.StringInput
	// chart timezone
	Timezone pulumi.StringPtrInput
	// optional mappings to transform data with rules
	ValueMappings DashboardHistogramChartValueMappingArrayInput
	// chart width in cols (max 20)
	Width pulumi.IntPtrInput
}

func (DashboardHistogramChartArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardHistogramChartArgs)(nil)).Elem()
}

type DashboardHistogramChartInput interface {
	pulumi.Input

	ToDashboardHistogramChartOutput() DashboardHistogramChartOutput
	ToDashboardHistogramChartOutputWithContext(ctx context.Context) DashboardHistogramChartOutput
}

func (*DashboardHistogramChart) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardHistogramChart)(nil)).Elem()
}

func (i *DashboardHistogramChart) ToDashboardHistogramChartOutput() DashboardHistogramChartOutput {
	return i.ToDashboardHistogramChartOutputWithContext(context.Background())
}

func (i *DashboardHistogramChart) ToDashboardHistogramChartOutputWithContext(ctx context.Context) DashboardHistogramChartOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardHistogramChartOutput)
}

// DashboardHistogramChartArrayInput is an input type that accepts DashboardHistogramChartArray and DashboardHistogramChartArrayOutput values.
// You can construct a concrete instance of `DashboardHistogramChartArrayInput` via:
//
//	DashboardHistogramChartArray{ DashboardHistogramChartArgs{...} }
type DashboardHistogramChartArrayInput interface {
	pulumi.Input

	ToDashboardHistogramChartArrayOutput() DashboardHistogramChartArrayOutput
	ToDashboardHistogramChartArrayOutputWithContext(context.Context) DashboardHistogramChartArrayOutput
}

type DashboardHistogramChartArray []DashboardHistogramChartInput

func (DashboardHistogramChartArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardHistogramChart)(nil)).Elem()
}

func (i DashboardHistogramChartArray) ToDashboardHistogramChartArrayOutput() DashboardHistogramChartArrayOutput {
	return i.ToDashboardHistogramChartArrayOutputWithContext(context.Background())
}

func (i DashboardHistogramChartArray) ToDashboardHistogramChartArrayOutputWithContext(ctx context.Context) DashboardHistogramChartArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardHistogramChartArrayOutput)
}

// DashboardHistogramChartMapInput is an input type that accepts DashboardHistogramChartMap and DashboardHistogramChartMapOutput values.
// You can construct a concrete instance of `DashboardHistogramChartMapInput` via:
//
//	DashboardHistogramChartMap{ "key": DashboardHistogramChartArgs{...} }
type DashboardHistogramChartMapInput interface {
	pulumi.Input

	ToDashboardHistogramChartMapOutput() DashboardHistogramChartMapOutput
	ToDashboardHistogramChartMapOutputWithContext(context.Context) DashboardHistogramChartMapOutput
}

type DashboardHistogramChartMap map[string]DashboardHistogramChartInput

func (DashboardHistogramChartMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardHistogramChart)(nil)).Elem()
}

func (i DashboardHistogramChartMap) ToDashboardHistogramChartMapOutput() DashboardHistogramChartMapOutput {
	return i.ToDashboardHistogramChartMapOutputWithContext(context.Background())
}

func (i DashboardHistogramChartMap) ToDashboardHistogramChartMapOutputWithContext(ctx context.Context) DashboardHistogramChartMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardHistogramChartMapOutput)
}

type DashboardHistogramChartOutput struct{ *pulumi.OutputState }

func (DashboardHistogramChartOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardHistogramChart)(nil)).Elem()
}

func (o DashboardHistogramChartOutput) ToDashboardHistogramChartOutput() DashboardHistogramChartOutput {
	return o
}

func (o DashboardHistogramChartOutput) ToDashboardHistogramChartOutputWithContext(ctx context.Context) DashboardHistogramChartOutput {
	return o
}

// bucket count
func (o DashboardHistogramChartOutput) BucketCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.BucketCount }).(pulumi.IntPtrOutput)
}

// bucket size
func (o DashboardHistogramChartOutput) BucketSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.BucketSize }).(pulumi.IntPtrOutput)
}

// categories top max limit
func (o DashboardHistogramChartOutput) CategoriesTopMaxLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.CategoriesTopMaxLimit }).(pulumi.IntPtrOutput)
}

// chart traces to be included
func (o DashboardHistogramChartOutput) ChartItems() DashboardHistogramChartChartItemArrayOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) DashboardHistogramChartChartItemArrayOutput { return v.ChartItems }).(DashboardHistogramChartChartItemArrayOutput)
}

func (o DashboardHistogramChartOutput) Collection() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.Collection }).(pulumi.StringPtrOutput)
}

// chart description
func (o DashboardHistogramChartOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// whether to display the time range or not
func (o DashboardHistogramChartOutput) DisplayTimeRange() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.BoolPtrOutput { return v.DisplayTimeRange }).(pulumi.BoolPtrOutput)
}

// chart height in px
func (o DashboardHistogramChartOutput) Height() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.Height }).(pulumi.IntPtrOutput)
}

// histogram type
func (o DashboardHistogramChartOutput) HistogramType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.HistogramType }).(pulumi.StringPtrOutput)
}

// [last|avg|...] aggregation
func (o DashboardHistogramChartOutput) LabelsAggregation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.LabelsAggregation }).(pulumi.StringPtrOutput)
}

// whether to display the labels or not
func (o DashboardHistogramChartOutput) LabelsDisplay() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.BoolPtrOutput { return v.LabelsDisplay }).(pulumi.BoolPtrOutput)
}

// [right|bottom] placement
func (o DashboardHistogramChartOutput) LabelsPlacement() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.LabelsPlacement }).(pulumi.StringPtrOutput)
}

// minimum chart height
func (o DashboardHistogramChartOutput) MinHeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.MinHeight }).(pulumi.IntPtrOutput)
}

// minimum chart width
func (o DashboardHistogramChartOutput) MinWidth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.MinWidth }).(pulumi.IntPtrOutput)
}

// name of the chart
func (o DashboardHistogramChartOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// number of decimals
func (o DashboardHistogramChartOutput) NumberOfDecimals() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.NumberOfDecimals }).(pulumi.IntPtrOutput)
}

// chart x position
func (o DashboardHistogramChartOutput) PositionX() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.PositionX }).(pulumi.IntPtrOutput)
}

// chart y position
func (o DashboardHistogramChartOutput) PositionY() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.PositionY }).(pulumi.IntPtrOutput)
}

// refresh interval
func (o DashboardHistogramChartOutput) RefreshInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.RefreshInterval }).(pulumi.StringPtrOutput)
}

// relative window time
func (o DashboardHistogramChartOutput) RelativeWindowTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.RelativeWindowTime }).(pulumi.StringPtrOutput)
}

// whether to show data which is beyond timestampLte or not
func (o DashboardHistogramChartOutput) ShowBeyondData() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.BoolPtrOutput { return v.ShowBeyondData }).(pulumi.BoolPtrOutput)
}

// sorting type
func (o DashboardHistogramChartOutput) Sorting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.Sorting }).(pulumi.StringPtrOutput)
}

// whether to stack or not the histogram
func (o DashboardHistogramChartOutput) Stacked() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.BoolPtrOutput { return v.Stacked }).(pulumi.BoolPtrOutput)
}

// id for the tab where to place the chart
func (o DashboardHistogramChartOutput) Tab() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringOutput { return v.Tab }).(pulumi.StringOutput)
}

// optional static lines to be added to the chart as references
func (o DashboardHistogramChartOutput) Thresholds() DashboardHistogramChartThresholdArrayOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) DashboardHistogramChartThresholdArrayOutput { return v.Thresholds }).(DashboardHistogramChartThresholdArrayOutput)
}

// date in isoformat or shortcut string where to end reading
func (o DashboardHistogramChartOutput) TimestampGte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringOutput { return v.TimestampGte }).(pulumi.StringOutput)
}

// date in isoformat or shortcut string where to start reading
func (o DashboardHistogramChartOutput) TimestampLte() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringOutput { return v.TimestampLte }).(pulumi.StringOutput)
}

// chart timezone
func (o DashboardHistogramChartOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// optional mappings to transform data with rules
func (o DashboardHistogramChartOutput) ValueMappings() DashboardHistogramChartValueMappingArrayOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) DashboardHistogramChartValueMappingArrayOutput {
		return v.ValueMappings
	}).(DashboardHistogramChartValueMappingArrayOutput)
}

// chart width in cols (max 20)
func (o DashboardHistogramChartOutput) Width() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardHistogramChart) pulumi.IntPtrOutput { return v.Width }).(pulumi.IntPtrOutput)
}

type DashboardHistogramChartArrayOutput struct{ *pulumi.OutputState }

func (DashboardHistogramChartArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardHistogramChart)(nil)).Elem()
}

func (o DashboardHistogramChartArrayOutput) ToDashboardHistogramChartArrayOutput() DashboardHistogramChartArrayOutput {
	return o
}

func (o DashboardHistogramChartArrayOutput) ToDashboardHistogramChartArrayOutputWithContext(ctx context.Context) DashboardHistogramChartArrayOutput {
	return o
}

func (o DashboardHistogramChartArrayOutput) Index(i pulumi.IntInput) DashboardHistogramChartOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardHistogramChart {
		return vs[0].([]*DashboardHistogramChart)[vs[1].(int)]
	}).(DashboardHistogramChartOutput)
}

type DashboardHistogramChartMapOutput struct{ *pulumi.OutputState }

func (DashboardHistogramChartMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardHistogramChart)(nil)).Elem()
}

func (o DashboardHistogramChartMapOutput) ToDashboardHistogramChartMapOutput() DashboardHistogramChartMapOutput {
	return o
}

func (o DashboardHistogramChartMapOutput) ToDashboardHistogramChartMapOutputWithContext(ctx context.Context) DashboardHistogramChartMapOutput {
	return o
}

func (o DashboardHistogramChartMapOutput) MapIndex(k pulumi.StringInput) DashboardHistogramChartOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardHistogramChart {
		return vs[0].(map[string]*DashboardHistogramChart)[vs[1].(string)]
	}).(DashboardHistogramChartOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardHistogramChartInput)(nil)).Elem(), &DashboardHistogramChart{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardHistogramChartArrayInput)(nil)).Elem(), DashboardHistogramChartArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardHistogramChartMapInput)(nil)).Elem(), DashboardHistogramChartMap{})
	pulumi.RegisterOutputType(DashboardHistogramChartOutput{})
	pulumi.RegisterOutputType(DashboardHistogramChartArrayOutput{})
	pulumi.RegisterOutputType(DashboardHistogramChartMapOutput{})
}
