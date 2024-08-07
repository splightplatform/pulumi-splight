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
//				"type": "GeometryCollection",
//				"geometries": []map[string]interface{}{
//					map[string]interface{}{
//						"type": "Point",
//						"coordinates": []float64{
//							0,
//							0,
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			myAsset, err := splight.NewAsset(ctx, "myAsset", &splight.AssetArgs{
//				Description: pulumi.String("My Asset Description"),
//				Geometry:    pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			myAttribute, err := splight.NewAssetAttribute(ctx, "myAttribute", &splight.AssetAttributeArgs{
//				Type:  pulumi.String("Number"),
//				Asset: myAsset.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			_, err = splight.NewAlert(ctx, "myAlert", &splight.AlertArgs{
//				Description: pulumi.String("My Alert Description"),
//				Type:        pulumi.String("rate"),
//				RateUnit:    pulumi.String("minute"),
//				RateValue:   pulumi.Int(10),
//				TimeWindow:  3600 * 12,
//				Thresholds: splight.AlertThresholdArray{
//					&splight.AlertThresholdArgs{
//						Value:      pulumi.Float64(1),
//						Status:     pulumi.String("alert"),
//						StatusText: pulumi.String("Some warning!"),
//					},
//				},
//				Severity:       pulumi.String("sev1"),
//				Operator:       pulumi.String("lt"),
//				Aggregation:    pulumi.String("max"),
//				TargetVariable: pulumi.String("A"),
//				AlertItems: splight.AlertAlertItemArray{
//					&splight.AlertAlertItemArgs{
//						RefId:           pulumi.String("A"),
//						Type:            pulumi.String("QUERY"),
//						Expression:      pulumi.String(""),
//						ExpressionPlain: pulumi.String(""),
//						QueryFilterAsset: &splight.AlertAlertItemQueryFilterAssetArgs{
//							Id:   myAsset.ID(),
//							Name: myAsset.Name,
//						},
//						QueryFilterAttribute: &splight.AlertAlertItemQueryFilterAttributeArgs{
//							Id:   myAttribute.ID(),
//							Name: myAttribute.Name,
//						},
//						QueryPlain: pulumi.All(myAsset.ID(), myAttribute.ID()).ApplyT(func(_args []interface{}) (string, error) {
//							myAssetId := _args[0].(string)
//							myAttributeId := _args[1].(string)
//							var _zero string
//							tmpJSON1, err := json.Marshal([]map[string]interface{}{
//								map[string]interface{}{
//									"$match": map[string]interface{}{
//										"asset":     myAssetId,
//										"attribute": myAttributeId,
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
// $ pulumi import splight:index/alert:Alert [options] splight_alert.<name> <alert_id>
// ```
type Alert struct {
	pulumi.CustomResourceState

	// aggregation to be applied to reads before comparisson
	Aggregation pulumi.StringOutput `pulumi:"aggregation"`
	// traces to be used to compute the results
	AlertItems AlertAlertItemArrayOutput `pulumi:"alertItems"`
	// schedule value for cron
	CronDom pulumi.IntOutput `pulumi:"cronDom"`
	// schedule value for cron
	CronDow pulumi.IntOutput `pulumi:"cronDow"`
	// schedule value for cron
	CronHours pulumi.IntOutput `pulumi:"cronHours"`
	// schedule value for cron
	CronMinutes pulumi.IntOutput `pulumi:"cronMinutes"`
	// schedule value for cron
	CronMonth pulumi.IntOutput `pulumi:"cronMonth"`
	// schedule value for cron
	CronYear pulumi.IntOutput `pulumi:"cronYear"`
	// The description of the resource
	Description pulumi.StringOutput `pulumi:"description"`
	// The name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// operator to be used to compare the read value with the threshold value
	Operator pulumi.StringOutput `pulumi:"operator"`
	// [day|hour|minute] schedule unit
	RateUnit pulumi.StringOutput `pulumi:"rateUnit"`
	// schedule value
	RateValue pulumi.IntOutput `pulumi:"rateValue"`
	// related assets to be linked. In case one of these alerts triggers it will be reflected on each of these assets.
	RelatedAssets pulumi.StringArrayOutput `pulumi:"relatedAssets"`
	// [sev1,...,sev8] severity for the alert
	Severity pulumi.StringOutput `pulumi:"severity"`
	// variable to be used to compare with thresholds
	TargetVariable pulumi.StringOutput       `pulumi:"targetVariable"`
	Thresholds     AlertThresholdArrayOutput `pulumi:"thresholds"`
	// window to fetch data from. Data out of that window will not be considered for evaluation
	TimeWindow pulumi.IntOutput `pulumi:"timeWindow"`
	// [cron|rate] type for the cron
	Type pulumi.StringOutput `pulumi:"type"`
}

// NewAlert registers a new resource with the given unique name, arguments, and options.
func NewAlert(ctx *pulumi.Context,
	name string, args *AlertArgs, opts ...pulumi.ResourceOption) (*Alert, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Aggregation == nil {
		return nil, errors.New("invalid value for required argument 'Aggregation'")
	}
	if args.AlertItems == nil {
		return nil, errors.New("invalid value for required argument 'AlertItems'")
	}
	if args.Description == nil {
		return nil, errors.New("invalid value for required argument 'Description'")
	}
	if args.Operator == nil {
		return nil, errors.New("invalid value for required argument 'Operator'")
	}
	if args.Severity == nil {
		return nil, errors.New("invalid value for required argument 'Severity'")
	}
	if args.TargetVariable == nil {
		return nil, errors.New("invalid value for required argument 'TargetVariable'")
	}
	if args.Thresholds == nil {
		return nil, errors.New("invalid value for required argument 'Thresholds'")
	}
	if args.TimeWindow == nil {
		return nil, errors.New("invalid value for required argument 'TimeWindow'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Alert
	err := ctx.RegisterResource("splight:index/alert:Alert", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAlert gets an existing Alert resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAlert(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertState, opts ...pulumi.ResourceOption) (*Alert, error) {
	var resource Alert
	err := ctx.ReadResource("splight:index/alert:Alert", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Alert resources.
type alertState struct {
	// aggregation to be applied to reads before comparisson
	Aggregation *string `pulumi:"aggregation"`
	// traces to be used to compute the results
	AlertItems []AlertAlertItem `pulumi:"alertItems"`
	// schedule value for cron
	CronDom *int `pulumi:"cronDom"`
	// schedule value for cron
	CronDow *int `pulumi:"cronDow"`
	// schedule value for cron
	CronHours *int `pulumi:"cronHours"`
	// schedule value for cron
	CronMinutes *int `pulumi:"cronMinutes"`
	// schedule value for cron
	CronMonth *int `pulumi:"cronMonth"`
	// schedule value for cron
	CronYear *int `pulumi:"cronYear"`
	// The description of the resource
	Description *string `pulumi:"description"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// operator to be used to compare the read value with the threshold value
	Operator *string `pulumi:"operator"`
	// [day|hour|minute] schedule unit
	RateUnit *string `pulumi:"rateUnit"`
	// schedule value
	RateValue *int `pulumi:"rateValue"`
	// related assets to be linked. In case one of these alerts triggers it will be reflected on each of these assets.
	RelatedAssets []string `pulumi:"relatedAssets"`
	// [sev1,...,sev8] severity for the alert
	Severity *string `pulumi:"severity"`
	// variable to be used to compare with thresholds
	TargetVariable *string          `pulumi:"targetVariable"`
	Thresholds     []AlertThreshold `pulumi:"thresholds"`
	// window to fetch data from. Data out of that window will not be considered for evaluation
	TimeWindow *int `pulumi:"timeWindow"`
	// [cron|rate] type for the cron
	Type *string `pulumi:"type"`
}

type AlertState struct {
	// aggregation to be applied to reads before comparisson
	Aggregation pulumi.StringPtrInput
	// traces to be used to compute the results
	AlertItems AlertAlertItemArrayInput
	// schedule value for cron
	CronDom pulumi.IntPtrInput
	// schedule value for cron
	CronDow pulumi.IntPtrInput
	// schedule value for cron
	CronHours pulumi.IntPtrInput
	// schedule value for cron
	CronMinutes pulumi.IntPtrInput
	// schedule value for cron
	CronMonth pulumi.IntPtrInput
	// schedule value for cron
	CronYear pulumi.IntPtrInput
	// The description of the resource
	Description pulumi.StringPtrInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// operator to be used to compare the read value with the threshold value
	Operator pulumi.StringPtrInput
	// [day|hour|minute] schedule unit
	RateUnit pulumi.StringPtrInput
	// schedule value
	RateValue pulumi.IntPtrInput
	// related assets to be linked. In case one of these alerts triggers it will be reflected on each of these assets.
	RelatedAssets pulumi.StringArrayInput
	// [sev1,...,sev8] severity for the alert
	Severity pulumi.StringPtrInput
	// variable to be used to compare with thresholds
	TargetVariable pulumi.StringPtrInput
	Thresholds     AlertThresholdArrayInput
	// window to fetch data from. Data out of that window will not be considered for evaluation
	TimeWindow pulumi.IntPtrInput
	// [cron|rate] type for the cron
	Type pulumi.StringPtrInput
}

func (AlertState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertState)(nil)).Elem()
}

type alertArgs struct {
	// aggregation to be applied to reads before comparisson
	Aggregation string `pulumi:"aggregation"`
	// traces to be used to compute the results
	AlertItems []AlertAlertItem `pulumi:"alertItems"`
	// schedule value for cron
	CronDom *int `pulumi:"cronDom"`
	// schedule value for cron
	CronDow *int `pulumi:"cronDow"`
	// schedule value for cron
	CronHours *int `pulumi:"cronHours"`
	// schedule value for cron
	CronMinutes *int `pulumi:"cronMinutes"`
	// schedule value for cron
	CronMonth *int `pulumi:"cronMonth"`
	// schedule value for cron
	CronYear *int `pulumi:"cronYear"`
	// The description of the resource
	Description string `pulumi:"description"`
	// The name of the resource
	Name *string `pulumi:"name"`
	// operator to be used to compare the read value with the threshold value
	Operator string `pulumi:"operator"`
	// [day|hour|minute] schedule unit
	RateUnit *string `pulumi:"rateUnit"`
	// schedule value
	RateValue *int `pulumi:"rateValue"`
	// related assets to be linked. In case one of these alerts triggers it will be reflected on each of these assets.
	RelatedAssets []string `pulumi:"relatedAssets"`
	// [sev1,...,sev8] severity for the alert
	Severity string `pulumi:"severity"`
	// variable to be used to compare with thresholds
	TargetVariable string           `pulumi:"targetVariable"`
	Thresholds     []AlertThreshold `pulumi:"thresholds"`
	// window to fetch data from. Data out of that window will not be considered for evaluation
	TimeWindow int `pulumi:"timeWindow"`
	// [cron|rate] type for the cron
	Type string `pulumi:"type"`
}

// The set of arguments for constructing a Alert resource.
type AlertArgs struct {
	// aggregation to be applied to reads before comparisson
	Aggregation pulumi.StringInput
	// traces to be used to compute the results
	AlertItems AlertAlertItemArrayInput
	// schedule value for cron
	CronDom pulumi.IntPtrInput
	// schedule value for cron
	CronDow pulumi.IntPtrInput
	// schedule value for cron
	CronHours pulumi.IntPtrInput
	// schedule value for cron
	CronMinutes pulumi.IntPtrInput
	// schedule value for cron
	CronMonth pulumi.IntPtrInput
	// schedule value for cron
	CronYear pulumi.IntPtrInput
	// The description of the resource
	Description pulumi.StringInput
	// The name of the resource
	Name pulumi.StringPtrInput
	// operator to be used to compare the read value with the threshold value
	Operator pulumi.StringInput
	// [day|hour|minute] schedule unit
	RateUnit pulumi.StringPtrInput
	// schedule value
	RateValue pulumi.IntPtrInput
	// related assets to be linked. In case one of these alerts triggers it will be reflected on each of these assets.
	RelatedAssets pulumi.StringArrayInput
	// [sev1,...,sev8] severity for the alert
	Severity pulumi.StringInput
	// variable to be used to compare with thresholds
	TargetVariable pulumi.StringInput
	Thresholds     AlertThresholdArrayInput
	// window to fetch data from. Data out of that window will not be considered for evaluation
	TimeWindow pulumi.IntInput
	// [cron|rate] type for the cron
	Type pulumi.StringInput
}

func (AlertArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertArgs)(nil)).Elem()
}

type AlertInput interface {
	pulumi.Input

	ToAlertOutput() AlertOutput
	ToAlertOutputWithContext(ctx context.Context) AlertOutput
}

func (*Alert) ElementType() reflect.Type {
	return reflect.TypeOf((**Alert)(nil)).Elem()
}

func (i *Alert) ToAlertOutput() AlertOutput {
	return i.ToAlertOutputWithContext(context.Background())
}

func (i *Alert) ToAlertOutputWithContext(ctx context.Context) AlertOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertOutput)
}

// AlertArrayInput is an input type that accepts AlertArray and AlertArrayOutput values.
// You can construct a concrete instance of `AlertArrayInput` via:
//
//	AlertArray{ AlertArgs{...} }
type AlertArrayInput interface {
	pulumi.Input

	ToAlertArrayOutput() AlertArrayOutput
	ToAlertArrayOutputWithContext(context.Context) AlertArrayOutput
}

type AlertArray []AlertInput

func (AlertArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alert)(nil)).Elem()
}

func (i AlertArray) ToAlertArrayOutput() AlertArrayOutput {
	return i.ToAlertArrayOutputWithContext(context.Background())
}

func (i AlertArray) ToAlertArrayOutputWithContext(ctx context.Context) AlertArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertArrayOutput)
}

// AlertMapInput is an input type that accepts AlertMap and AlertMapOutput values.
// You can construct a concrete instance of `AlertMapInput` via:
//
//	AlertMap{ "key": AlertArgs{...} }
type AlertMapInput interface {
	pulumi.Input

	ToAlertMapOutput() AlertMapOutput
	ToAlertMapOutputWithContext(context.Context) AlertMapOutput
}

type AlertMap map[string]AlertInput

func (AlertMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alert)(nil)).Elem()
}

func (i AlertMap) ToAlertMapOutput() AlertMapOutput {
	return i.ToAlertMapOutputWithContext(context.Background())
}

func (i AlertMap) ToAlertMapOutputWithContext(ctx context.Context) AlertMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertMapOutput)
}

type AlertOutput struct{ *pulumi.OutputState }

func (AlertOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Alert)(nil)).Elem()
}

func (o AlertOutput) ToAlertOutput() AlertOutput {
	return o
}

func (o AlertOutput) ToAlertOutputWithContext(ctx context.Context) AlertOutput {
	return o
}

// aggregation to be applied to reads before comparisson
func (o AlertOutput) Aggregation() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Aggregation }).(pulumi.StringOutput)
}

// traces to be used to compute the results
func (o AlertOutput) AlertItems() AlertAlertItemArrayOutput {
	return o.ApplyT(func(v *Alert) AlertAlertItemArrayOutput { return v.AlertItems }).(AlertAlertItemArrayOutput)
}

// schedule value for cron
func (o AlertOutput) CronDom() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.CronDom }).(pulumi.IntOutput)
}

// schedule value for cron
func (o AlertOutput) CronDow() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.CronDow }).(pulumi.IntOutput)
}

// schedule value for cron
func (o AlertOutput) CronHours() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.CronHours }).(pulumi.IntOutput)
}

// schedule value for cron
func (o AlertOutput) CronMinutes() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.CronMinutes }).(pulumi.IntOutput)
}

// schedule value for cron
func (o AlertOutput) CronMonth() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.CronMonth }).(pulumi.IntOutput)
}

// schedule value for cron
func (o AlertOutput) CronYear() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.CronYear }).(pulumi.IntOutput)
}

// The description of the resource
func (o AlertOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

// The name of the resource
func (o AlertOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// operator to be used to compare the read value with the threshold value
func (o AlertOutput) Operator() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Operator }).(pulumi.StringOutput)
}

// [day|hour|minute] schedule unit
func (o AlertOutput) RateUnit() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.RateUnit }).(pulumi.StringOutput)
}

// schedule value
func (o AlertOutput) RateValue() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.RateValue }).(pulumi.IntOutput)
}

// related assets to be linked. In case one of these alerts triggers it will be reflected on each of these assets.
func (o AlertOutput) RelatedAssets() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringArrayOutput { return v.RelatedAssets }).(pulumi.StringArrayOutput)
}

// [sev1,...,sev8] severity for the alert
func (o AlertOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

// variable to be used to compare with thresholds
func (o AlertOutput) TargetVariable() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.TargetVariable }).(pulumi.StringOutput)
}

func (o AlertOutput) Thresholds() AlertThresholdArrayOutput {
	return o.ApplyT(func(v *Alert) AlertThresholdArrayOutput { return v.Thresholds }).(AlertThresholdArrayOutput)
}

// window to fetch data from. Data out of that window will not be considered for evaluation
func (o AlertOutput) TimeWindow() pulumi.IntOutput {
	return o.ApplyT(func(v *Alert) pulumi.IntOutput { return v.TimeWindow }).(pulumi.IntOutput)
}

// [cron|rate] type for the cron
func (o AlertOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Alert) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

type AlertArrayOutput struct{ *pulumi.OutputState }

func (AlertArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Alert)(nil)).Elem()
}

func (o AlertArrayOutput) ToAlertArrayOutput() AlertArrayOutput {
	return o
}

func (o AlertArrayOutput) ToAlertArrayOutputWithContext(ctx context.Context) AlertArrayOutput {
	return o
}

func (o AlertArrayOutput) Index(i pulumi.IntInput) AlertOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Alert {
		return vs[0].([]*Alert)[vs[1].(int)]
	}).(AlertOutput)
}

type AlertMapOutput struct{ *pulumi.OutputState }

func (AlertMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Alert)(nil)).Elem()
}

func (o AlertMapOutput) ToAlertMapOutput() AlertMapOutput {
	return o
}

func (o AlertMapOutput) ToAlertMapOutputWithContext(ctx context.Context) AlertMapOutput {
	return o
}

func (o AlertMapOutput) MapIndex(k pulumi.StringInput) AlertOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Alert {
		return vs[0].(map[string]*Alert)[vs[1].(string)]
	}).(AlertOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlertInput)(nil)).Elem(), &Alert{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertArrayInput)(nil)).Elem(), AlertArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertMapInput)(nil)).Elem(), AlertMap{})
	pulumi.RegisterOutputType(AlertOutput{})
	pulumi.RegisterOutputType(AlertArrayOutput{})
	pulumi.RegisterOutputType(AlertMapOutput{})
}
