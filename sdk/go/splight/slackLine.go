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
// $ pulumi import splight:index/slackLine:SlackLine [options] splight_slack_line.<name> <line_id>
// ```
type SlackLine struct {
	pulumi.CustomResourceState

	// description of the resource
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// geo position and shape of the resource
	Geometry pulumi.StringPtrOutput `pulumi:"geometry"`
	// kind of the resource
	Kinds SlackLineKindArrayOutput `pulumi:"kinds"`
	// name of the resource
	Name pulumi.StringOutput `pulumi:"name"`
	// tags of the resource
	Tags SlackLineTagArrayOutput `pulumi:"tags"`
}

// NewSlackLine registers a new resource with the given unique name, arguments, and options.
func NewSlackLine(ctx *pulumi.Context,
	name string, args *SlackLineArgs, opts ...pulumi.ResourceOption) (*SlackLine, error) {
	if args == nil {
		args = &SlackLineArgs{}
	}

	opts = internal.PkgResourceDefaultOpts(opts)
	var resource SlackLine
	err := ctx.RegisterResource("splight:index/slackLine:SlackLine", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetSlackLine gets an existing SlackLine resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetSlackLine(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SlackLineState, opts ...pulumi.ResourceOption) (*SlackLine, error) {
	var resource SlackLine
	err := ctx.ReadResource("splight:index/slackLine:SlackLine", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering SlackLine resources.
type slackLineState struct {
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// kind of the resource
	Kinds []SlackLineKind `pulumi:"kinds"`
	// name of the resource
	Name *string `pulumi:"name"`
	// tags of the resource
	Tags []SlackLineTag `pulumi:"tags"`
}

type SlackLineState struct {
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// kind of the resource
	Kinds SlackLineKindArrayInput
	// name of the resource
	Name pulumi.StringPtrInput
	// tags of the resource
	Tags SlackLineTagArrayInput
}

func (SlackLineState) ElementType() reflect.Type {
	return reflect.TypeOf((*slackLineState)(nil)).Elem()
}

type slackLineArgs struct {
	// description of the resource
	Description *string `pulumi:"description"`
	// geo position and shape of the resource
	Geometry *string `pulumi:"geometry"`
	// name of the resource
	Name *string `pulumi:"name"`
	// tags of the resource
	Tags []SlackLineTag `pulumi:"tags"`
}

// The set of arguments for constructing a SlackLine resource.
type SlackLineArgs struct {
	// description of the resource
	Description pulumi.StringPtrInput
	// geo position and shape of the resource
	Geometry pulumi.StringPtrInput
	// name of the resource
	Name pulumi.StringPtrInput
	// tags of the resource
	Tags SlackLineTagArrayInput
}

func (SlackLineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*slackLineArgs)(nil)).Elem()
}

type SlackLineInput interface {
	pulumi.Input

	ToSlackLineOutput() SlackLineOutput
	ToSlackLineOutputWithContext(ctx context.Context) SlackLineOutput
}

func (*SlackLine) ElementType() reflect.Type {
	return reflect.TypeOf((**SlackLine)(nil)).Elem()
}

func (i *SlackLine) ToSlackLineOutput() SlackLineOutput {
	return i.ToSlackLineOutputWithContext(context.Background())
}

func (i *SlackLine) ToSlackLineOutputWithContext(ctx context.Context) SlackLineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackLineOutput)
}

// SlackLineArrayInput is an input type that accepts SlackLineArray and SlackLineArrayOutput values.
// You can construct a concrete instance of `SlackLineArrayInput` via:
//
//	SlackLineArray{ SlackLineArgs{...} }
type SlackLineArrayInput interface {
	pulumi.Input

	ToSlackLineArrayOutput() SlackLineArrayOutput
	ToSlackLineArrayOutputWithContext(context.Context) SlackLineArrayOutput
}

type SlackLineArray []SlackLineInput

func (SlackLineArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SlackLine)(nil)).Elem()
}

func (i SlackLineArray) ToSlackLineArrayOutput() SlackLineArrayOutput {
	return i.ToSlackLineArrayOutputWithContext(context.Background())
}

func (i SlackLineArray) ToSlackLineArrayOutputWithContext(ctx context.Context) SlackLineArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackLineArrayOutput)
}

// SlackLineMapInput is an input type that accepts SlackLineMap and SlackLineMapOutput values.
// You can construct a concrete instance of `SlackLineMapInput` via:
//
//	SlackLineMap{ "key": SlackLineArgs{...} }
type SlackLineMapInput interface {
	pulumi.Input

	ToSlackLineMapOutput() SlackLineMapOutput
	ToSlackLineMapOutputWithContext(context.Context) SlackLineMapOutput
}

type SlackLineMap map[string]SlackLineInput

func (SlackLineMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SlackLine)(nil)).Elem()
}

func (i SlackLineMap) ToSlackLineMapOutput() SlackLineMapOutput {
	return i.ToSlackLineMapOutputWithContext(context.Background())
}

func (i SlackLineMap) ToSlackLineMapOutputWithContext(ctx context.Context) SlackLineMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackLineMapOutput)
}

type SlackLineOutput struct{ *pulumi.OutputState }

func (SlackLineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlackLine)(nil)).Elem()
}

func (o SlackLineOutput) ToSlackLineOutput() SlackLineOutput {
	return o
}

func (o SlackLineOutput) ToSlackLineOutputWithContext(ctx context.Context) SlackLineOutput {
	return o
}

// description of the resource
func (o SlackLineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackLine) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// geo position and shape of the resource
func (o SlackLineOutput) Geometry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackLine) pulumi.StringPtrOutput { return v.Geometry }).(pulumi.StringPtrOutput)
}

// kind of the resource
func (o SlackLineOutput) Kinds() SlackLineKindArrayOutput {
	return o.ApplyT(func(v *SlackLine) SlackLineKindArrayOutput { return v.Kinds }).(SlackLineKindArrayOutput)
}

// name of the resource
func (o SlackLineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SlackLine) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// tags of the resource
func (o SlackLineOutput) Tags() SlackLineTagArrayOutput {
	return o.ApplyT(func(v *SlackLine) SlackLineTagArrayOutput { return v.Tags }).(SlackLineTagArrayOutput)
}

type SlackLineArrayOutput struct{ *pulumi.OutputState }

func (SlackLineArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*SlackLine)(nil)).Elem()
}

func (o SlackLineArrayOutput) ToSlackLineArrayOutput() SlackLineArrayOutput {
	return o
}

func (o SlackLineArrayOutput) ToSlackLineArrayOutputWithContext(ctx context.Context) SlackLineArrayOutput {
	return o
}

func (o SlackLineArrayOutput) Index(i pulumi.IntInput) SlackLineOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *SlackLine {
		return vs[0].([]*SlackLine)[vs[1].(int)]
	}).(SlackLineOutput)
}

type SlackLineMapOutput struct{ *pulumi.OutputState }

func (SlackLineMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*SlackLine)(nil)).Elem()
}

func (o SlackLineMapOutput) ToSlackLineMapOutput() SlackLineMapOutput {
	return o
}

func (o SlackLineMapOutput) ToSlackLineMapOutputWithContext(ctx context.Context) SlackLineMapOutput {
	return o
}

func (o SlackLineMapOutput) MapIndex(k pulumi.StringInput) SlackLineOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *SlackLine {
		return vs[0].(map[string]*SlackLine)[vs[1].(string)]
	}).(SlackLineOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SlackLineInput)(nil)).Elem(), &SlackLine{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackLineArrayInput)(nil)).Elem(), SlackLineArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackLineMapInput)(nil)).Elem(), SlackLineMap{})
	pulumi.RegisterOutputType(SlackLineOutput{})
	pulumi.RegisterOutputType(SlackLineArrayOutput{})
	pulumi.RegisterOutputType(SlackLineMapOutput{})
}
