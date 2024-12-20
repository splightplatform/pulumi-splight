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
// ## Import
//
// ```sh
// $ pulumi import splight:index/transformer:Transformer [options] splight_transformer.<name> <transformer_id>
// ```
type Transformer struct {
	pulumi.CustomResourceState

	// attribute of the resource
	ActivePowerHvs TransformerActivePowerHvArrayOutput `pulumi:"activePowerHvs"`
	// attribute of the resource
	ActivePowerLosses TransformerActivePowerLossArrayOutput `pulumi:"activePowerLosses"`
	// attribute of the resource
	ActivePowerLvs TransformerActivePowerLvArrayOutput `pulumi:"activePowerLvs"`
	// attribute of the resource
	Capacitance TransformerCapacitanceOutput `pulumi:"capacitance"`
	// attribute of the resource
	Conductance TransformerConductanceOutput `pulumi:"conductance"`
	// attribute of the resource
	Contingencies TransformerContingencyArrayOutput `pulumi:"contingencies"`
	// attribute of the resource
	CurrentHvs TransformerCurrentHvArrayOutput `pulumi:"currentHvs"`
	// attribute of the resource
	CurrentLvs TransformerCurrentLvArrayOutput `pulumi:"currentLvs"`
	// description of the resource
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// geo position and shape of the resource
	Geometry pulumi.StringPtrOutput `pulumi:"geometry"`
	// kind of the resource
	Kinds TransformerKindArrayOutput `pulumi:"kinds"`
	// attribute of the resource
	MaximumAllowedCurrent TransformerMaximumAllowedCurrentOutput `pulumi:"maximumAllowedCurrent"`
	// attribute of the resource
	MaximumAllowedPower TransformerMaximumAllowedPowerOutput `pulumi:"maximumAllowedPower"`
	// name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// attribute of the resource
	Reactance TransformerReactanceOutput `pulumi:"reactance"`
	// attribute of the resource
	ReactivePowerHvs TransformerReactivePowerHvArrayOutput `pulumi:"reactivePowerHvs"`
	// attribute of the resource
	ReactivePowerLosses TransformerReactivePowerLossArrayOutput `pulumi:"reactivePowerLosses"`
	// attribute of the resource
	ReactivePowerLvs TransformerReactivePowerLvArrayOutput `pulumi:"reactivePowerLvs"`
	// attribute of the resource
	Resistance TransformerResistanceOutput `pulumi:"resistance"`
	// attribute of the resource
	SafetyMarginForPower TransformerSafetyMarginForPowerOutput `pulumi:"safetyMarginForPower"`
	// attribute of the resource
	StandardType TransformerStandardTypeOutput `pulumi:"standardType"`
	// attribute of the resource
	SwitchStatusHvs TransformerSwitchStatusHvArrayOutput `pulumi:"switchStatusHvs"`
	// attribute of the resource
	SwitchStatusLvs TransformerSwitchStatusLvArrayOutput `pulumi:"switchStatusLvs"`
	// tags of the resource
	Tags TransformerTagArrayOutput `pulumi:"tags"`
	// attribute of the resource
	TapPos TransformerTapPosOutput `pulumi:"tapPos"`
	// timezone that overrides location-based timezone of the resource
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
	// attribute of the resource
	VoltageHvs TransformerVoltageHvArrayOutput `pulumi:"voltageHvs"`
	// attribute of the resource
	VoltageLvs TransformerVoltageLvArrayOutput `pulumi:"voltageLvs"`
	// attribute of the resource
	XnOhm TransformerXnOhmOutput `pulumi:"xnOhm"`
}

// NewTransformer registers a new resource with the given unique name, arguments, and options.
func NewTransformer(ctx *pulumi.Context,
	name string, args *TransformerArgs, opts ...pulumi.ResourceOption) (*Transformer, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Capacitance == nil {
		return nil, errors.New("invalid value for required argument 'Capacitance'")
	}
	if args.Conductance == nil {
		return nil, errors.New("invalid value for required argument 'Conductance'")
	}
	if args.MaximumAllowedCurrent == nil {
		return nil, errors.New("invalid value for required argument 'MaximumAllowedCurrent'")
	}
	if args.MaximumAllowedPower == nil {
		return nil, errors.New("invalid value for required argument 'MaximumAllowedPower'")
	}
	if args.Reactance == nil {
		return nil, errors.New("invalid value for required argument 'Reactance'")
	}
	if args.Resistance == nil {
		return nil, errors.New("invalid value for required argument 'Resistance'")
	}
	if args.SafetyMarginForPower == nil {
		return nil, errors.New("invalid value for required argument 'SafetyMarginForPower'")
	}
	if args.StandardType == nil {
		return nil, errors.New("invalid value for required argument 'StandardType'")
	}
	if args.TapPos == nil {
		return nil, errors.New("invalid value for required argument 'TapPos'")
	}
	if args.XnOhm == nil {
		return nil, errors.New("invalid value for required argument 'XnOhm'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Transformer
	err := ctx.RegisterResource("splight:index/transformer:Transformer", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetTransformer gets an existing Transformer resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetTransformer(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TransformerState, opts ...pulumi.ResourceOption) (*Transformer, error) {
	var resource Transformer
	err := ctx.ReadResource("splight:index/transformer:Transformer", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Transformer resources.
type transformerState struct {
	// attribute of the resource
	ActivePowerHvs []TransformerActivePowerHv `pulumi:"activePowerHvs"`
	// attribute of the resource
	ActivePowerLosses []TransformerActivePowerLoss `pulumi:"activePowerLosses"`
	// attribute of the resource
	ActivePowerLvs []TransformerActivePowerLv `pulumi:"activePowerLvs"`
	// attribute of the resource
	Capacitance *TransformerCapacitance `pulumi:"capacitance"`
	// attribute of the resource
	Conductance *TransformerConductance `pulumi:"conductance"`
	// attribute of the resource
	Contingencies []TransformerContingency `pulumi:"contingencies"`
	// attribute of the resource
	CurrentHvs []TransformerCurrentHv `pulumi:"currentHvs"`
	// attribute of the resource
	CurrentLvs []TransformerCurrentLv `pulumi:"currentLvs"`
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// kind of the resource
	Kinds []TransformerKind `pulumi:"kinds"`
	// attribute of the resource
	MaximumAllowedCurrent *TransformerMaximumAllowedCurrent `pulumi:"maximumAllowedCurrent"`
	// attribute of the resource
	MaximumAllowedPower *TransformerMaximumAllowedPower `pulumi:"maximumAllowedPower"`
	// name of the resource
	Name *string `pulumi:"name"`
	// attribute of the resource
	Reactance *TransformerReactance `pulumi:"reactance"`
	// attribute of the resource
	ReactivePowerHvs []TransformerReactivePowerHv `pulumi:"reactivePowerHvs"`
	// attribute of the resource
	ReactivePowerLosses []TransformerReactivePowerLoss `pulumi:"reactivePowerLosses"`
	// attribute of the resource
	ReactivePowerLvs []TransformerReactivePowerLv `pulumi:"reactivePowerLvs"`
	// attribute of the resource
	Resistance *TransformerResistance `pulumi:"resistance"`
	// attribute of the resource
	SafetyMarginForPower *TransformerSafetyMarginForPower `pulumi:"safetyMarginForPower"`
	// attribute of the resource
	StandardType *TransformerStandardType `pulumi:"standardType"`
	// attribute of the resource
	SwitchStatusHvs []TransformerSwitchStatusHv `pulumi:"switchStatusHvs"`
	// attribute of the resource
	SwitchStatusLvs []TransformerSwitchStatusLv `pulumi:"switchStatusLvs"`
	// tags of the resource
	Tags []TransformerTag `pulumi:"tags"`
	// attribute of the resource
	TapPos *TransformerTapPos `pulumi:"tapPos"`
	// timezone that overrides location-based timezone of the resource
	Timezone *string `pulumi:"timezone"`
	// attribute of the resource
	VoltageHvs []TransformerVoltageHv `pulumi:"voltageHvs"`
	// attribute of the resource
	VoltageLvs []TransformerVoltageLv `pulumi:"voltageLvs"`
	// attribute of the resource
	XnOhm *TransformerXnOhm `pulumi:"xnOhm"`
}

type TransformerState struct {
	// attribute of the resource
	ActivePowerHvs TransformerActivePowerHvArrayInput
	// attribute of the resource
	ActivePowerLosses TransformerActivePowerLossArrayInput
	// attribute of the resource
	ActivePowerLvs TransformerActivePowerLvArrayInput
	// attribute of the resource
	Capacitance TransformerCapacitancePtrInput
	// attribute of the resource
	Conductance TransformerConductancePtrInput
	// attribute of the resource
	Contingencies TransformerContingencyArrayInput
	// attribute of the resource
	CurrentHvs TransformerCurrentHvArrayInput
	// attribute of the resource
	CurrentLvs TransformerCurrentLvArrayInput
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// kind of the resource
	Kinds TransformerKindArrayInput
	// attribute of the resource
	MaximumAllowedCurrent TransformerMaximumAllowedCurrentPtrInput
	// attribute of the resource
	MaximumAllowedPower TransformerMaximumAllowedPowerPtrInput
	// name of the resource
	Name pulumi.StringPtrInput
	// attribute of the resource
	Reactance TransformerReactancePtrInput
	// attribute of the resource
	ReactivePowerHvs TransformerReactivePowerHvArrayInput
	// attribute of the resource
	ReactivePowerLosses TransformerReactivePowerLossArrayInput
	// attribute of the resource
	ReactivePowerLvs TransformerReactivePowerLvArrayInput
	// attribute of the resource
	Resistance TransformerResistancePtrInput
	// attribute of the resource
	SafetyMarginForPower TransformerSafetyMarginForPowerPtrInput
	// attribute of the resource
	StandardType TransformerStandardTypePtrInput
	// attribute of the resource
	SwitchStatusHvs TransformerSwitchStatusHvArrayInput
	// attribute of the resource
	SwitchStatusLvs TransformerSwitchStatusLvArrayInput
	// tags of the resource
	Tags TransformerTagArrayInput
	// attribute of the resource
	TapPos TransformerTapPosPtrInput
	// timezone that overrides location-based timezone of the resource
	Timezone pulumi.StringPtrInput
	// attribute of the resource
	VoltageHvs TransformerVoltageHvArrayInput
	// attribute of the resource
	VoltageLvs TransformerVoltageLvArrayInput
	// attribute of the resource
	XnOhm TransformerXnOhmPtrInput
}

func (TransformerState) ElementType() reflect.Type {
	return reflect.TypeOf((*transformerState)(nil)).Elem()
}

type transformerArgs struct {
	// attribute of the resource
	Capacitance TransformerCapacitance `pulumi:"capacitance"`
	// attribute of the resource
	Conductance TransformerConductance `pulumi:"conductance"`
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// attribute of the resource
	MaximumAllowedCurrent TransformerMaximumAllowedCurrent `pulumi:"maximumAllowedCurrent"`
	// attribute of the resource
	MaximumAllowedPower TransformerMaximumAllowedPower `pulumi:"maximumAllowedPower"`
	// name of the resource
	Name *string `pulumi:"name"`
	// attribute of the resource
	Reactance TransformerReactance `pulumi:"reactance"`
	// attribute of the resource
	Resistance TransformerResistance `pulumi:"resistance"`
	// attribute of the resource
	SafetyMarginForPower TransformerSafetyMarginForPower `pulumi:"safetyMarginForPower"`
	// attribute of the resource
	StandardType TransformerStandardType `pulumi:"standardType"`
	// tags of the resource
	Tags []TransformerTag `pulumi:"tags"`
	// attribute of the resource
	TapPos TransformerTapPos `pulumi:"tapPos"`
	// timezone that overrides location-based timezone of the resource
	Timezone *string `pulumi:"timezone"`
	// attribute of the resource
	XnOhm TransformerXnOhm `pulumi:"xnOhm"`
}

// The set of arguments for constructing a Transformer resource.
type TransformerArgs struct {
	// attribute of the resource
	Capacitance TransformerCapacitanceInput
	// attribute of the resource
	Conductance TransformerConductanceInput
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// attribute of the resource
	MaximumAllowedCurrent TransformerMaximumAllowedCurrentInput
	// attribute of the resource
	MaximumAllowedPower TransformerMaximumAllowedPowerInput
	// name of the resource
	Name pulumi.StringPtrInput
	// attribute of the resource
	Reactance TransformerReactanceInput
	// attribute of the resource
	Resistance TransformerResistanceInput
	// attribute of the resource
	SafetyMarginForPower TransformerSafetyMarginForPowerInput
	// attribute of the resource
	StandardType TransformerStandardTypeInput
	// tags of the resource
	Tags TransformerTagArrayInput
	// attribute of the resource
	TapPos TransformerTapPosInput
	// timezone that overrides location-based timezone of the resource
	Timezone pulumi.StringPtrInput
	// attribute of the resource
	XnOhm TransformerXnOhmInput
}

func (TransformerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*transformerArgs)(nil)).Elem()
}

type TransformerInput interface {
	pulumi.Input

	ToTransformerOutput() TransformerOutput
	ToTransformerOutputWithContext(ctx context.Context) TransformerOutput
}

func (*Transformer) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformer)(nil)).Elem()
}

func (i *Transformer) ToTransformerOutput() TransformerOutput {
	return i.ToTransformerOutputWithContext(context.Background())
}

func (i *Transformer) ToTransformerOutputWithContext(ctx context.Context) TransformerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformerOutput)
}

// TransformerArrayInput is an input type that accepts TransformerArray and TransformerArrayOutput values.
// You can construct a concrete instance of `TransformerArrayInput` via:
//
//	TransformerArray{ TransformerArgs{...} }
type TransformerArrayInput interface {
	pulumi.Input

	ToTransformerArrayOutput() TransformerArrayOutput
	ToTransformerArrayOutputWithContext(context.Context) TransformerArrayOutput
}

type TransformerArray []TransformerInput

func (TransformerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Transformer)(nil)).Elem()
}

func (i TransformerArray) ToTransformerArrayOutput() TransformerArrayOutput {
	return i.ToTransformerArrayOutputWithContext(context.Background())
}

func (i TransformerArray) ToTransformerArrayOutputWithContext(ctx context.Context) TransformerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformerArrayOutput)
}

// TransformerMapInput is an input type that accepts TransformerMap and TransformerMapOutput values.
// You can construct a concrete instance of `TransformerMapInput` via:
//
//	TransformerMap{ "key": TransformerArgs{...} }
type TransformerMapInput interface {
	pulumi.Input

	ToTransformerMapOutput() TransformerMapOutput
	ToTransformerMapOutputWithContext(context.Context) TransformerMapOutput
}

type TransformerMap map[string]TransformerInput

func (TransformerMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Transformer)(nil)).Elem()
}

func (i TransformerMap) ToTransformerMapOutput() TransformerMapOutput {
	return i.ToTransformerMapOutputWithContext(context.Background())
}

func (i TransformerMap) ToTransformerMapOutputWithContext(ctx context.Context) TransformerMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TransformerMapOutput)
}

type TransformerOutput struct{ *pulumi.OutputState }

func (TransformerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Transformer)(nil)).Elem()
}

func (o TransformerOutput) ToTransformerOutput() TransformerOutput {
	return o
}

func (o TransformerOutput) ToTransformerOutputWithContext(ctx context.Context) TransformerOutput {
	return o
}

// attribute of the resource
func (o TransformerOutput) ActivePowerHvs() TransformerActivePowerHvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerActivePowerHvArrayOutput { return v.ActivePowerHvs }).(TransformerActivePowerHvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) ActivePowerLosses() TransformerActivePowerLossArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerActivePowerLossArrayOutput { return v.ActivePowerLosses }).(TransformerActivePowerLossArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) ActivePowerLvs() TransformerActivePowerLvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerActivePowerLvArrayOutput { return v.ActivePowerLvs }).(TransformerActivePowerLvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) Capacitance() TransformerCapacitanceOutput {
	return o.ApplyT(func(v *Transformer) TransformerCapacitanceOutput { return v.Capacitance }).(TransformerCapacitanceOutput)
}

// attribute of the resource
func (o TransformerOutput) Conductance() TransformerConductanceOutput {
	return o.ApplyT(func(v *Transformer) TransformerConductanceOutput { return v.Conductance }).(TransformerConductanceOutput)
}

// attribute of the resource
func (o TransformerOutput) Contingencies() TransformerContingencyArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerContingencyArrayOutput { return v.Contingencies }).(TransformerContingencyArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) CurrentHvs() TransformerCurrentHvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerCurrentHvArrayOutput { return v.CurrentHvs }).(TransformerCurrentHvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) CurrentLvs() TransformerCurrentLvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerCurrentLvArrayOutput { return v.CurrentLvs }).(TransformerCurrentLvArrayOutput)
}

// description of the resource
func (o TransformerOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformer) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// geo position and shape of the resource
func (o TransformerOutput) Geometry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformer) pulumi.StringPtrOutput { return v.Geometry }).(pulumi.StringPtrOutput)
}

// kind of the resource
func (o TransformerOutput) Kinds() TransformerKindArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerKindArrayOutput { return v.Kinds }).(TransformerKindArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) MaximumAllowedCurrent() TransformerMaximumAllowedCurrentOutput {
	return o.ApplyT(func(v *Transformer) TransformerMaximumAllowedCurrentOutput { return v.MaximumAllowedCurrent }).(TransformerMaximumAllowedCurrentOutput)
}

// attribute of the resource
func (o TransformerOutput) MaximumAllowedPower() TransformerMaximumAllowedPowerOutput {
	return o.ApplyT(func(v *Transformer) TransformerMaximumAllowedPowerOutput { return v.MaximumAllowedPower }).(TransformerMaximumAllowedPowerOutput)
}

// name of the resource
func (o TransformerOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Transformer) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// attribute of the resource
func (o TransformerOutput) Reactance() TransformerReactanceOutput {
	return o.ApplyT(func(v *Transformer) TransformerReactanceOutput { return v.Reactance }).(TransformerReactanceOutput)
}

// attribute of the resource
func (o TransformerOutput) ReactivePowerHvs() TransformerReactivePowerHvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerReactivePowerHvArrayOutput { return v.ReactivePowerHvs }).(TransformerReactivePowerHvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) ReactivePowerLosses() TransformerReactivePowerLossArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerReactivePowerLossArrayOutput { return v.ReactivePowerLosses }).(TransformerReactivePowerLossArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) ReactivePowerLvs() TransformerReactivePowerLvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerReactivePowerLvArrayOutput { return v.ReactivePowerLvs }).(TransformerReactivePowerLvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) Resistance() TransformerResistanceOutput {
	return o.ApplyT(func(v *Transformer) TransformerResistanceOutput { return v.Resistance }).(TransformerResistanceOutput)
}

// attribute of the resource
func (o TransformerOutput) SafetyMarginForPower() TransformerSafetyMarginForPowerOutput {
	return o.ApplyT(func(v *Transformer) TransformerSafetyMarginForPowerOutput { return v.SafetyMarginForPower }).(TransformerSafetyMarginForPowerOutput)
}

// attribute of the resource
func (o TransformerOutput) StandardType() TransformerStandardTypeOutput {
	return o.ApplyT(func(v *Transformer) TransformerStandardTypeOutput { return v.StandardType }).(TransformerStandardTypeOutput)
}

// attribute of the resource
func (o TransformerOutput) SwitchStatusHvs() TransformerSwitchStatusHvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerSwitchStatusHvArrayOutput { return v.SwitchStatusHvs }).(TransformerSwitchStatusHvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) SwitchStatusLvs() TransformerSwitchStatusLvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerSwitchStatusLvArrayOutput { return v.SwitchStatusLvs }).(TransformerSwitchStatusLvArrayOutput)
}

// tags of the resource
func (o TransformerOutput) Tags() TransformerTagArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerTagArrayOutput { return v.Tags }).(TransformerTagArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) TapPos() TransformerTapPosOutput {
	return o.ApplyT(func(v *Transformer) TransformerTapPosOutput { return v.TapPos }).(TransformerTapPosOutput)
}

// timezone that overrides location-based timezone of the resource
func (o TransformerOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Transformer) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

// attribute of the resource
func (o TransformerOutput) VoltageHvs() TransformerVoltageHvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerVoltageHvArrayOutput { return v.VoltageHvs }).(TransformerVoltageHvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) VoltageLvs() TransformerVoltageLvArrayOutput {
	return o.ApplyT(func(v *Transformer) TransformerVoltageLvArrayOutput { return v.VoltageLvs }).(TransformerVoltageLvArrayOutput)
}

// attribute of the resource
func (o TransformerOutput) XnOhm() TransformerXnOhmOutput {
	return o.ApplyT(func(v *Transformer) TransformerXnOhmOutput { return v.XnOhm }).(TransformerXnOhmOutput)
}

type TransformerArrayOutput struct{ *pulumi.OutputState }

func (TransformerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Transformer)(nil)).Elem()
}

func (o TransformerArrayOutput) ToTransformerArrayOutput() TransformerArrayOutput {
	return o
}

func (o TransformerArrayOutput) ToTransformerArrayOutputWithContext(ctx context.Context) TransformerArrayOutput {
	return o
}

func (o TransformerArrayOutput) Index(i pulumi.IntInput) TransformerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Transformer {
		return vs[0].([]*Transformer)[vs[1].(int)]
	}).(TransformerOutput)
}

type TransformerMapOutput struct{ *pulumi.OutputState }

func (TransformerMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Transformer)(nil)).Elem()
}

func (o TransformerMapOutput) ToTransformerMapOutput() TransformerMapOutput {
	return o
}

func (o TransformerMapOutput) ToTransformerMapOutputWithContext(ctx context.Context) TransformerMapOutput {
	return o
}

func (o TransformerMapOutput) MapIndex(k pulumi.StringInput) TransformerOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Transformer {
		return vs[0].(map[string]*Transformer)[vs[1].(string)]
	}).(TransformerOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TransformerInput)(nil)).Elem(), &Transformer{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformerArrayInput)(nil)).Elem(), TransformerArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TransformerMapInput)(nil)).Elem(), TransformerMap{})
	pulumi.RegisterOutputType(TransformerOutput{})
	pulumi.RegisterOutputType(TransformerArrayOutput{})
	pulumi.RegisterOutputType(TransformerMapOutput{})
}