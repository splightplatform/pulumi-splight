// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splight

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/splightplatform/pulumi-splight/sdk/go/splight/internal"
)

// ## Example Usage
//
// ## Import
//
// ```sh
// $ pulumi import splight:index/externalGrid:ExternalGrid [options] splight_external_grid.<name> <external_grid_id>
// ```
type ExternalGrid struct {
	pulumi.CustomResourceState

	// description of the resource
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// geo position and shape of the resource
	Geometry pulumi.StringPtrOutput `pulumi:"geometry"`
	// kind of the resource
	Kinds ExternalGridKindArrayOutput `pulumi:"kinds"`
	// name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// tags of the resource
	Tags ExternalGridTagArrayOutput `pulumi:"tags"`
}

// NewExternalGrid registers a new resource with the given unique name, arguments, and options.
func NewExternalGrid(ctx *pulumi.Context,
	name string, args *ExternalGridArgs, opts ...pulumi.ResourceOption) (*ExternalGrid, error) {
	if args == nil {
		args = &ExternalGridArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource ExternalGrid
	err := ctx.RegisterResource("splight:index/externalGrid:ExternalGrid", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetExternalGrid gets an existing ExternalGrid resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetExternalGrid(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ExternalGridState, opts ...pulumi.ResourceOption) (*ExternalGrid, error) {
	var resource ExternalGrid
	err := ctx.ReadResource("splight:index/externalGrid:ExternalGrid", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ExternalGrid resources.
type externalGridState struct {
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// kind of the resource
	Kinds []ExternalGridKind `pulumi:"kinds"`
	// name of the resource
	Name *string `pulumi:"name"`
	// tags of the resource
	Tags []ExternalGridTag `pulumi:"tags"`
}

type ExternalGridState struct {
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// kind of the resource
	Kinds ExternalGridKindArrayInput
	// name of the resource
	Name pulumi.StringPtrInput
	// tags of the resource
	Tags ExternalGridTagArrayInput
}

func (ExternalGridState) ElementType() reflect.Type {
	return reflect.TypeOf((*externalGridState)(nil)).Elem()
}

type externalGridArgs struct {
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// name of the resource
	Name *string `pulumi:"name"`
	// tags of the resource
	Tags []ExternalGridTag `pulumi:"tags"`
}

// The set of arguments for constructing a ExternalGrid resource.
type ExternalGridArgs struct {
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// name of the resource
	Name pulumi.StringPtrInput
	// tags of the resource
	Tags ExternalGridTagArrayInput
}

func (ExternalGridArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*externalGridArgs)(nil)).Elem()
}

type ExternalGridInput interface {
	pulumi.Input

	ToExternalGridOutput() ExternalGridOutput
	ToExternalGridOutputWithContext(ctx context.Context) ExternalGridOutput
}

func (*ExternalGrid) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalGrid)(nil)).Elem()
}

func (i *ExternalGrid) ToExternalGridOutput() ExternalGridOutput {
	return i.ToExternalGridOutputWithContext(context.Background())
}

func (i *ExternalGrid) ToExternalGridOutputWithContext(ctx context.Context) ExternalGridOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalGridOutput)
}

// ExternalGridArrayInput is an input type that accepts ExternalGridArray and ExternalGridArrayOutput values.
// You can construct a concrete instance of `ExternalGridArrayInput` via:
//
//	ExternalGridArray{ ExternalGridArgs{...} }
type ExternalGridArrayInput interface {
	pulumi.Input

	ToExternalGridArrayOutput() ExternalGridArrayOutput
	ToExternalGridArrayOutputWithContext(context.Context) ExternalGridArrayOutput
}

type ExternalGridArray []ExternalGridInput

func (ExternalGridArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalGrid)(nil)).Elem()
}

func (i ExternalGridArray) ToExternalGridArrayOutput() ExternalGridArrayOutput {
	return i.ToExternalGridArrayOutputWithContext(context.Background())
}

func (i ExternalGridArray) ToExternalGridArrayOutputWithContext(ctx context.Context) ExternalGridArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalGridArrayOutput)
}

// ExternalGridMapInput is an input type that accepts ExternalGridMap and ExternalGridMapOutput values.
// You can construct a concrete instance of `ExternalGridMapInput` via:
//
//	ExternalGridMap{ "key": ExternalGridArgs{...} }
type ExternalGridMapInput interface {
	pulumi.Input

	ToExternalGridMapOutput() ExternalGridMapOutput
	ToExternalGridMapOutputWithContext(context.Context) ExternalGridMapOutput
}

type ExternalGridMap map[string]ExternalGridInput

func (ExternalGridMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalGrid)(nil)).Elem()
}

func (i ExternalGridMap) ToExternalGridMapOutput() ExternalGridMapOutput {
	return i.ToExternalGridMapOutputWithContext(context.Background())
}

func (i ExternalGridMap) ToExternalGridMapOutputWithContext(ctx context.Context) ExternalGridMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExternalGridMapOutput)
}

type ExternalGridOutput struct{ *pulumi.OutputState }

func (ExternalGridOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExternalGrid)(nil)).Elem()
}

func (o ExternalGridOutput) ToExternalGridOutput() ExternalGridOutput {
	return o
}

func (o ExternalGridOutput) ToExternalGridOutputWithContext(ctx context.Context) ExternalGridOutput {
	return o
}

// description of the resource
func (o ExternalGridOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalGrid) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// geo position and shape of the resource
func (o ExternalGridOutput) Geometry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExternalGrid) pulumi.StringPtrOutput { return v.Geometry }).(pulumi.StringPtrOutput)
}

// kind of the resource
func (o ExternalGridOutput) Kinds() ExternalGridKindArrayOutput {
	return o.ApplyT(func(v *ExternalGrid) ExternalGridKindArrayOutput { return v.Kinds }).(ExternalGridKindArrayOutput)
}

// name of the resource
func (o ExternalGridOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ExternalGrid) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// tags of the resource
func (o ExternalGridOutput) Tags() ExternalGridTagArrayOutput {
	return o.ApplyT(func(v *ExternalGrid) ExternalGridTagArrayOutput { return v.Tags }).(ExternalGridTagArrayOutput)
}

type ExternalGridArrayOutput struct{ *pulumi.OutputState }

func (ExternalGridArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ExternalGrid)(nil)).Elem()
}

func (o ExternalGridArrayOutput) ToExternalGridArrayOutput() ExternalGridArrayOutput {
	return o
}

func (o ExternalGridArrayOutput) ToExternalGridArrayOutputWithContext(ctx context.Context) ExternalGridArrayOutput {
	return o
}

func (o ExternalGridArrayOutput) Index(i pulumi.IntInput) ExternalGridOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ExternalGrid {
		return vs[0].([]*ExternalGrid)[vs[1].(int)]
	}).(ExternalGridOutput)
}

type ExternalGridMapOutput struct{ *pulumi.OutputState }

func (ExternalGridMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ExternalGrid)(nil)).Elem()
}

func (o ExternalGridMapOutput) ToExternalGridMapOutput() ExternalGridMapOutput {
	return o
}

func (o ExternalGridMapOutput) ToExternalGridMapOutputWithContext(ctx context.Context) ExternalGridMapOutput {
	return o
}

func (o ExternalGridMapOutput) MapIndex(k pulumi.StringInput) ExternalGridOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ExternalGrid {
		return vs[0].(map[string]*ExternalGrid)[vs[1].(string)]
	}).(ExternalGridOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalGridInput)(nil)).Elem(), &ExternalGrid{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalGridArrayInput)(nil)).Elem(), ExternalGridArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ExternalGridMapInput)(nil)).Elem(), ExternalGridMap{})
	pulumi.RegisterOutputType(ExternalGridOutput{})
	pulumi.RegisterOutputType(ExternalGridArrayOutput{})
	pulumi.RegisterOutputType(ExternalGridMapOutput{})
}