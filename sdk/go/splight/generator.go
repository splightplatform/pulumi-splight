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
// $ pulumi import splight:index/generator:Generator [options] splight_generator.<name> <generator_id>
// ```
type Generator struct {
	pulumi.CustomResourceState

	// attribute of the resource
	ActivePowers GeneratorActivePowerArrayOutput `pulumi:"activePowers"`
	// attribute of the resource
	Co2Coefficient GeneratorCo2CoefficientOutput `pulumi:"co2Coefficient"`
	// attribute of the resource
	DailyEmissionAvoideds GeneratorDailyEmissionAvoidedArrayOutput `pulumi:"dailyEmissionAvoideds"`
	// attribute of the resource
	DailyEnergies GeneratorDailyEnergyArrayOutput `pulumi:"dailyEnergies"`
	// description of the resource
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// geo position and shape of the resource
	Geometry pulumi.StringPtrOutput `pulumi:"geometry"`
	// kind of the resource
	Kinds GeneratorKindArrayOutput `pulumi:"kinds"`
	// name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// attribute of the resource
	ReactivePowers GeneratorReactivePowerArrayOutput `pulumi:"reactivePowers"`
	// tags of the resource
	Tags GeneratorTagArrayOutput `pulumi:"tags"`
}

// NewGenerator registers a new resource with the given unique name, arguments, and options.
func NewGenerator(ctx *pulumi.Context,
	name string, args *GeneratorArgs, opts ...pulumi.ResourceOption) (*Generator, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Co2Coefficient == nil {
		return nil, errors.New("invalid value for required argument 'Co2Coefficient'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Generator
	err := ctx.RegisterResource("splight:index/generator:Generator", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetGenerator gets an existing Generator resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetGenerator(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GeneratorState, opts ...pulumi.ResourceOption) (*Generator, error) {
	var resource Generator
	err := ctx.ReadResource("splight:index/generator:Generator", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Generator resources.
type generatorState struct {
	// attribute of the resource
	ActivePowers []GeneratorActivePower `pulumi:"activePowers"`
	// attribute of the resource
	Co2Coefficient *GeneratorCo2Coefficient `pulumi:"co2Coefficient"`
	// attribute of the resource
	DailyEmissionAvoideds []GeneratorDailyEmissionAvoided `pulumi:"dailyEmissionAvoideds"`
	// attribute of the resource
	DailyEnergies []GeneratorDailyEnergy `pulumi:"dailyEnergies"`
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// kind of the resource
	Kinds []GeneratorKind `pulumi:"kinds"`
	// name of the resource
	Name *string `pulumi:"name"`
	// attribute of the resource
	ReactivePowers []GeneratorReactivePower `pulumi:"reactivePowers"`
	// tags of the resource
	Tags []GeneratorTag `pulumi:"tags"`
}

type GeneratorState struct {
	// attribute of the resource
	ActivePowers GeneratorActivePowerArrayInput
	// attribute of the resource
	Co2Coefficient GeneratorCo2CoefficientPtrInput
	// attribute of the resource
	DailyEmissionAvoideds GeneratorDailyEmissionAvoidedArrayInput
	// attribute of the resource
	DailyEnergies GeneratorDailyEnergyArrayInput
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// kind of the resource
	Kinds GeneratorKindArrayInput
	// name of the resource
	Name pulumi.StringPtrInput
	// attribute of the resource
	ReactivePowers GeneratorReactivePowerArrayInput
	// tags of the resource
	Tags GeneratorTagArrayInput
}

func (GeneratorState) ElementType() reflect.Type {
	return reflect.TypeOf((*generatorState)(nil)).Elem()
}

type generatorArgs struct {
	// attribute of the resource
	Co2Coefficient GeneratorCo2Coefficient `pulumi:"co2Coefficient"`
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// name of the resource
	Name *string `pulumi:"name"`
	// tags of the resource
	Tags []GeneratorTag `pulumi:"tags"`
}

// The set of arguments for constructing a Generator resource.
type GeneratorArgs struct {
	// attribute of the resource
	Co2Coefficient GeneratorCo2CoefficientInput
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// name of the resource
	Name pulumi.StringPtrInput
	// tags of the resource
	Tags GeneratorTagArrayInput
}

func (GeneratorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*generatorArgs)(nil)).Elem()
}

type GeneratorInput interface {
	pulumi.Input

	ToGeneratorOutput() GeneratorOutput
	ToGeneratorOutputWithContext(ctx context.Context) GeneratorOutput
}

func (*Generator) ElementType() reflect.Type {
	return reflect.TypeOf((**Generator)(nil)).Elem()
}

func (i *Generator) ToGeneratorOutput() GeneratorOutput {
	return i.ToGeneratorOutputWithContext(context.Background())
}

func (i *Generator) ToGeneratorOutputWithContext(ctx context.Context) GeneratorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneratorOutput)
}

// GeneratorArrayInput is an input type that accepts GeneratorArray and GeneratorArrayOutput values.
// You can construct a concrete instance of `GeneratorArrayInput` via:
//
//	GeneratorArray{ GeneratorArgs{...} }
type GeneratorArrayInput interface {
	pulumi.Input

	ToGeneratorArrayOutput() GeneratorArrayOutput
	ToGeneratorArrayOutputWithContext(context.Context) GeneratorArrayOutput
}

type GeneratorArray []GeneratorInput

func (GeneratorArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Generator)(nil)).Elem()
}

func (i GeneratorArray) ToGeneratorArrayOutput() GeneratorArrayOutput {
	return i.ToGeneratorArrayOutputWithContext(context.Background())
}

func (i GeneratorArray) ToGeneratorArrayOutputWithContext(ctx context.Context) GeneratorArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneratorArrayOutput)
}

// GeneratorMapInput is an input type that accepts GeneratorMap and GeneratorMapOutput values.
// You can construct a concrete instance of `GeneratorMapInput` via:
//
//	GeneratorMap{ "key": GeneratorArgs{...} }
type GeneratorMapInput interface {
	pulumi.Input

	ToGeneratorMapOutput() GeneratorMapOutput
	ToGeneratorMapOutputWithContext(context.Context) GeneratorMapOutput
}

type GeneratorMap map[string]GeneratorInput

func (GeneratorMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Generator)(nil)).Elem()
}

func (i GeneratorMap) ToGeneratorMapOutput() GeneratorMapOutput {
	return i.ToGeneratorMapOutputWithContext(context.Background())
}

func (i GeneratorMap) ToGeneratorMapOutputWithContext(ctx context.Context) GeneratorMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GeneratorMapOutput)
}

type GeneratorOutput struct{ *pulumi.OutputState }

func (GeneratorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Generator)(nil)).Elem()
}

func (o GeneratorOutput) ToGeneratorOutput() GeneratorOutput {
	return o
}

func (o GeneratorOutput) ToGeneratorOutputWithContext(ctx context.Context) GeneratorOutput {
	return o
}

// attribute of the resource
func (o GeneratorOutput) ActivePowers() GeneratorActivePowerArrayOutput {
	return o.ApplyT(func(v *Generator) GeneratorActivePowerArrayOutput { return v.ActivePowers }).(GeneratorActivePowerArrayOutput)
}

// attribute of the resource
func (o GeneratorOutput) Co2Coefficient() GeneratorCo2CoefficientOutput {
	return o.ApplyT(func(v *Generator) GeneratorCo2CoefficientOutput { return v.Co2Coefficient }).(GeneratorCo2CoefficientOutput)
}

// attribute of the resource
func (o GeneratorOutput) DailyEmissionAvoideds() GeneratorDailyEmissionAvoidedArrayOutput {
	return o.ApplyT(func(v *Generator) GeneratorDailyEmissionAvoidedArrayOutput { return v.DailyEmissionAvoideds }).(GeneratorDailyEmissionAvoidedArrayOutput)
}

// attribute of the resource
func (o GeneratorOutput) DailyEnergies() GeneratorDailyEnergyArrayOutput {
	return o.ApplyT(func(v *Generator) GeneratorDailyEnergyArrayOutput { return v.DailyEnergies }).(GeneratorDailyEnergyArrayOutput)
}

// description of the resource
func (o GeneratorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Generator) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// geo position and shape of the resource
func (o GeneratorOutput) Geometry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Generator) pulumi.StringPtrOutput { return v.Geometry }).(pulumi.StringPtrOutput)
}

// kind of the resource
func (o GeneratorOutput) Kinds() GeneratorKindArrayOutput {
	return o.ApplyT(func(v *Generator) GeneratorKindArrayOutput { return v.Kinds }).(GeneratorKindArrayOutput)
}

// name of the resource
func (o GeneratorOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Generator) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// attribute of the resource
func (o GeneratorOutput) ReactivePowers() GeneratorReactivePowerArrayOutput {
	return o.ApplyT(func(v *Generator) GeneratorReactivePowerArrayOutput { return v.ReactivePowers }).(GeneratorReactivePowerArrayOutput)
}

// tags of the resource
func (o GeneratorOutput) Tags() GeneratorTagArrayOutput {
	return o.ApplyT(func(v *Generator) GeneratorTagArrayOutput { return v.Tags }).(GeneratorTagArrayOutput)
}

type GeneratorArrayOutput struct{ *pulumi.OutputState }

func (GeneratorArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Generator)(nil)).Elem()
}

func (o GeneratorArrayOutput) ToGeneratorArrayOutput() GeneratorArrayOutput {
	return o
}

func (o GeneratorArrayOutput) ToGeneratorArrayOutputWithContext(ctx context.Context) GeneratorArrayOutput {
	return o
}

func (o GeneratorArrayOutput) Index(i pulumi.IntInput) GeneratorOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Generator {
		return vs[0].([]*Generator)[vs[1].(int)]
	}).(GeneratorOutput)
}

type GeneratorMapOutput struct{ *pulumi.OutputState }

func (GeneratorMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Generator)(nil)).Elem()
}

func (o GeneratorMapOutput) ToGeneratorMapOutput() GeneratorMapOutput {
	return o
}

func (o GeneratorMapOutput) ToGeneratorMapOutputWithContext(ctx context.Context) GeneratorMapOutput {
	return o
}

func (o GeneratorMapOutput) MapIndex(k pulumi.StringInput) GeneratorOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Generator {
		return vs[0].(map[string]*Generator)[vs[1].(string)]
	}).(GeneratorOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*GeneratorInput)(nil)).Elem(), &Generator{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneratorArrayInput)(nil)).Elem(), GeneratorArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GeneratorMapInput)(nil)).Elem(), GeneratorMap{})
	pulumi.RegisterOutputType(GeneratorOutput{})
	pulumi.RegisterOutputType(GeneratorArrayOutput{})
	pulumi.RegisterOutputType(GeneratorMapOutput{})
}
