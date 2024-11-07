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
// $ pulumi import splight:index/bus:Bus [options] splight_bus.<name> <bus_id>
// ```
type Bus struct {
	pulumi.CustomResourceState

	// description of the resource
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// geo position and shape of the resource
	Geometry pulumi.StringPtrOutput `pulumi:"geometry"`
	// kind of the resource
	Kinds BusKindArrayOutput `pulumi:"kinds"`
	// name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// attribute of the resource
	NominalVoltage BusNominalVoltageOutput `pulumi:"nominalVoltage"`
	// tags of the resource
	Tags BusTagArrayOutput `pulumi:"tags"`
	// timezone that overrides location-based timezone of the resource
	Timezone pulumi.StringPtrOutput `pulumi:"timezone"`
}

// NewBus registers a new resource with the given unique name, arguments, and options.
func NewBus(ctx *pulumi.Context,
	name string, args *BusArgs, opts ...pulumi.ResourceOption) (*Bus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NominalVoltage == nil {
		return nil, errors.New("invalid value for required argument 'NominalVoltage'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Bus
	err := ctx.RegisterResource("splight:index/bus:Bus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetBus gets an existing Bus resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetBus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *BusState, opts ...pulumi.ResourceOption) (*Bus, error) {
	var resource Bus
	err := ctx.ReadResource("splight:index/bus:Bus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Bus resources.
type busState struct {
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// kind of the resource
	Kinds []BusKind `pulumi:"kinds"`
	// name of the resource
	Name *string `pulumi:"name"`
	// attribute of the resource
	NominalVoltage *BusNominalVoltage `pulumi:"nominalVoltage"`
	// tags of the resource
	Tags []BusTag `pulumi:"tags"`
	// timezone that overrides location-based timezone of the resource
	Timezone *string `pulumi:"timezone"`
}

type BusState struct {
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// kind of the resource
	Kinds BusKindArrayInput
	// name of the resource
	Name pulumi.StringPtrInput
	// attribute of the resource
	NominalVoltage BusNominalVoltagePtrInput
	// tags of the resource
	Tags BusTagArrayInput
	// timezone that overrides location-based timezone of the resource
	Timezone pulumi.StringPtrInput
}

func (BusState) ElementType() reflect.Type {
	return reflect.TypeOf((*busState)(nil)).Elem()
}

type busArgs struct {
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// name of the resource
	Name *string `pulumi:"name"`
	// attribute of the resource
	NominalVoltage BusNominalVoltage `pulumi:"nominalVoltage"`
	// tags of the resource
	Tags []BusTag `pulumi:"tags"`
	// timezone that overrides location-based timezone of the resource
	Timezone *string `pulumi:"timezone"`
}

// The set of arguments for constructing a Bus resource.
type BusArgs struct {
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// name of the resource
	Name pulumi.StringPtrInput
	// attribute of the resource
	NominalVoltage BusNominalVoltageInput
	// tags of the resource
	Tags BusTagArrayInput
	// timezone that overrides location-based timezone of the resource
	Timezone pulumi.StringPtrInput
}

func (BusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*busArgs)(nil)).Elem()
}

type BusInput interface {
	pulumi.Input

	ToBusOutput() BusOutput
	ToBusOutputWithContext(ctx context.Context) BusOutput
}

func (*Bus) ElementType() reflect.Type {
	return reflect.TypeOf((**Bus)(nil)).Elem()
}

func (i *Bus) ToBusOutput() BusOutput {
	return i.ToBusOutputWithContext(context.Background())
}

func (i *Bus) ToBusOutputWithContext(ctx context.Context) BusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusOutput)
}

// BusArrayInput is an input type that accepts BusArray and BusArrayOutput values.
// You can construct a concrete instance of `BusArrayInput` via:
//
//	BusArray{ BusArgs{...} }
type BusArrayInput interface {
	pulumi.Input

	ToBusArrayOutput() BusArrayOutput
	ToBusArrayOutputWithContext(context.Context) BusArrayOutput
}

type BusArray []BusInput

func (BusArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bus)(nil)).Elem()
}

func (i BusArray) ToBusArrayOutput() BusArrayOutput {
	return i.ToBusArrayOutputWithContext(context.Background())
}

func (i BusArray) ToBusArrayOutputWithContext(ctx context.Context) BusArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusArrayOutput)
}

// BusMapInput is an input type that accepts BusMap and BusMapOutput values.
// You can construct a concrete instance of `BusMapInput` via:
//
//	BusMap{ "key": BusArgs{...} }
type BusMapInput interface {
	pulumi.Input

	ToBusMapOutput() BusMapOutput
	ToBusMapOutputWithContext(context.Context) BusMapOutput
}

type BusMap map[string]BusInput

func (BusMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bus)(nil)).Elem()
}

func (i BusMap) ToBusMapOutput() BusMapOutput {
	return i.ToBusMapOutputWithContext(context.Background())
}

func (i BusMap) ToBusMapOutputWithContext(ctx context.Context) BusMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BusMapOutput)
}

type BusOutput struct{ *pulumi.OutputState }

func (BusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Bus)(nil)).Elem()
}

func (o BusOutput) ToBusOutput() BusOutput {
	return o
}

func (o BusOutput) ToBusOutputWithContext(ctx context.Context) BusOutput {
	return o
}

// description of the resource
func (o BusOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bus) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// geo position and shape of the resource
func (o BusOutput) Geometry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bus) pulumi.StringPtrOutput { return v.Geometry }).(pulumi.StringPtrOutput)
}

// kind of the resource
func (o BusOutput) Kinds() BusKindArrayOutput {
	return o.ApplyT(func(v *Bus) BusKindArrayOutput { return v.Kinds }).(BusKindArrayOutput)
}

// name of the resource
func (o BusOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Bus) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// attribute of the resource
func (o BusOutput) NominalVoltage() BusNominalVoltageOutput {
	return o.ApplyT(func(v *Bus) BusNominalVoltageOutput { return v.NominalVoltage }).(BusNominalVoltageOutput)
}

// tags of the resource
func (o BusOutput) Tags() BusTagArrayOutput {
	return o.ApplyT(func(v *Bus) BusTagArrayOutput { return v.Tags }).(BusTagArrayOutput)
}

// timezone that overrides location-based timezone of the resource
func (o BusOutput) Timezone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Bus) pulumi.StringPtrOutput { return v.Timezone }).(pulumi.StringPtrOutput)
}

type BusArrayOutput struct{ *pulumi.OutputState }

func (BusArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Bus)(nil)).Elem()
}

func (o BusArrayOutput) ToBusArrayOutput() BusArrayOutput {
	return o
}

func (o BusArrayOutput) ToBusArrayOutputWithContext(ctx context.Context) BusArrayOutput {
	return o
}

func (o BusArrayOutput) Index(i pulumi.IntInput) BusOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Bus {
		return vs[0].([]*Bus)[vs[1].(int)]
	}).(BusOutput)
}

type BusMapOutput struct{ *pulumi.OutputState }

func (BusMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Bus)(nil)).Elem()
}

func (o BusMapOutput) ToBusMapOutput() BusMapOutput {
	return o
}

func (o BusMapOutput) ToBusMapOutputWithContext(ctx context.Context) BusMapOutput {
	return o
}

func (o BusMapOutput) MapIndex(k pulumi.StringInput) BusOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Bus {
		return vs[0].(map[string]*Bus)[vs[1].(string)]
	}).(BusOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*BusInput)(nil)).Elem(), &Bus{})
	pulumi.RegisterInputType(reflect.TypeOf((*BusArrayInput)(nil)).Elem(), BusArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*BusMapInput)(nil)).Elem(), BusMap{})
	pulumi.RegisterOutputType(BusOutput{})
	pulumi.RegisterOutputType(BusArrayOutput{})
	pulumi.RegisterOutputType(BusMapOutput{})
}
