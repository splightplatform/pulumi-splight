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
//			tmpJSON0, err := json.Marshal(10)
//			if err != nil {
//				return err
//			}
//			json0 := string(tmpJSON0)
//			tmpJSON1, err := json.Marshal(1)
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			tmpJSON2, err := json.Marshal(150)
//			if err != nil {
//				return err
//			}
//			json2 := string(tmpJSON2)
//			tmpJSON3, err := json.Marshal(3)
//			if err != nil {
//				return err
//			}
//			json3 := string(tmpJSON3)
//			tmpJSON4, err := json.Marshal("true")
//			if err != nil {
//				return err
//			}
//			json4 := string(tmpJSON4)
//			_, err = splight.NewComponent(ctx, "componentTest", &splight.ComponentArgs{
//				Description: pulumi.String("Created with Terraform"),
//				Version:     pulumi.String("Random-3.1.0"),
//				Inputs: splight.ComponentInputTypeArray{
//					&splight.ComponentInputTypeArgs{
//						Name:        pulumi.String("period"),
//						Type:        pulumi.String("int"),
//						Value:       pulumi.String(json0),
//						Multiple:    pulumi.Bool(false),
//						Required:    pulumi.Bool(false),
//						Sensitive:   pulumi.Bool(false),
//						Description: pulumi.String(""),
//					},
//					&splight.ComponentInputTypeArgs{
//						Name:        pulumi.String("min"),
//						Type:        pulumi.String("int"),
//						Value:       pulumi.String(json1),
//						Multiple:    pulumi.Bool(false),
//						Required:    pulumi.Bool(false),
//						Sensitive:   pulumi.Bool(false),
//						Description: pulumi.String(""),
//					},
//					&splight.ComponentInputTypeArgs{
//						Name:        pulumi.String("max"),
//						Type:        pulumi.String("int"),
//						Value:       pulumi.String(json2),
//						Multiple:    pulumi.Bool(false),
//						Required:    pulumi.Bool(false),
//						Sensitive:   pulumi.Bool(false),
//						Description: pulumi.String(""),
//					},
//					&splight.ComponentInputTypeArgs{
//						Name:        pulumi.String("max_iterations"),
//						Type:        pulumi.String("int"),
//						Value:       pulumi.String(json3),
//						Multiple:    pulumi.Bool(false),
//						Required:    pulumi.Bool(false),
//						Sensitive:   pulumi.Bool(false),
//						Description: pulumi.String(""),
//					},
//					&splight.ComponentInputTypeArgs{
//						Name:        pulumi.String("should_crash"),
//						Type:        pulumi.String("bool"),
//						Value:       pulumi.String(json4),
//						Multiple:    pulumi.Bool(false),
//						Required:    pulumi.Bool(false),
//						Sensitive:   pulumi.Bool(false),
//						Description: pulumi.String(""),
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
// $ pulumi import splight:index/component:Component [options] splight_component.<name> <component_id>
// ```
type Component struct {
	pulumi.CustomResourceState

	// optinal description to add details of the resource
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// The input parameters based on hubcomponent spec
	Inputs ComponentInputTypeArrayOutput `pulumi:"inputs"`
	// the name of the component to be created
	Name pulumi.StringOutput `pulumi:"name"`
	// [NAME-VERSION] the version of the hub component
	Version pulumi.StringOutput `pulumi:"version"`
}

// NewComponent registers a new resource with the given unique name, arguments, and options.
func NewComponent(ctx *pulumi.Context,
	name string, args *ComponentArgs, opts ...pulumi.ResourceOption) (*Component, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Inputs == nil {
		return nil, errors.New("invalid value for required argument 'Inputs'")
	}
	if args.Version == nil {
		return nil, errors.New("invalid value for required argument 'Version'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Component
	err := ctx.RegisterResource("splight:index/component:Component", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetComponent gets an existing Component resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetComponent(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ComponentState, opts ...pulumi.ResourceOption) (*Component, error) {
	var resource Component
	err := ctx.ReadResource("splight:index/component:Component", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Component resources.
type componentState struct {
	// optinal description to add details of the resource
	Description *string `pulumi:"description"`
	// The input parameters based on hubcomponent spec
	Inputs []ComponentInputType `pulumi:"inputs"`
	// the name of the component to be created
	Name *string `pulumi:"name"`
	// [NAME-VERSION] the version of the hub component
	Version *string `pulumi:"version"`
}

type ComponentState struct {
	// optinal description to add details of the resource
	Description pulumi.StringPtrInput
	// The input parameters based on hubcomponent spec
	Inputs ComponentInputTypeArrayInput
	// the name of the component to be created
	Name pulumi.StringPtrInput
	// [NAME-VERSION] the version of the hub component
	Version pulumi.StringPtrInput
}

func (ComponentState) ElementType() reflect.Type {
	return reflect.TypeOf((*componentState)(nil)).Elem()
}

type componentArgs struct {
	// optinal description to add details of the resource
	Description *string `pulumi:"description"`
	// The input parameters based on hubcomponent spec
	Inputs []ComponentInputType `pulumi:"inputs"`
	// the name of the component to be created
	Name *string `pulumi:"name"`
	// [NAME-VERSION] the version of the hub component
	Version string `pulumi:"version"`
}

// The set of arguments for constructing a Component resource.
type ComponentArgs struct {
	// optinal description to add details of the resource
	Description pulumi.StringPtrInput
	// The input parameters based on hubcomponent spec
	Inputs ComponentInputTypeArrayInput
	// the name of the component to be created
	Name pulumi.StringPtrInput
	// [NAME-VERSION] the version of the hub component
	Version pulumi.StringInput
}

func (ComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*componentArgs)(nil)).Elem()
}

type ComponentInput interface {
	pulumi.Input

	ToComponentOutput() ComponentOutput
	ToComponentOutputWithContext(ctx context.Context) ComponentOutput
}

func (*Component) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (i *Component) ToComponentOutput() ComponentOutput {
	return i.ToComponentOutputWithContext(context.Background())
}

func (i *Component) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentOutput)
}

// ComponentArrayInput is an input type that accepts ComponentArray and ComponentArrayOutput values.
// You can construct a concrete instance of `ComponentArrayInput` via:
//
//	ComponentArray{ ComponentArgs{...} }
type ComponentArrayInput interface {
	pulumi.Input

	ToComponentArrayOutput() ComponentArrayOutput
	ToComponentArrayOutputWithContext(context.Context) ComponentArrayOutput
}

type ComponentArray []ComponentInput

func (ComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (i ComponentArray) ToComponentArrayOutput() ComponentArrayOutput {
	return i.ToComponentArrayOutputWithContext(context.Background())
}

func (i ComponentArray) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentArrayOutput)
}

// ComponentMapInput is an input type that accepts ComponentMap and ComponentMapOutput values.
// You can construct a concrete instance of `ComponentMapInput` via:
//
//	ComponentMap{ "key": ComponentArgs{...} }
type ComponentMapInput interface {
	pulumi.Input

	ToComponentMapOutput() ComponentMapOutput
	ToComponentMapOutputWithContext(context.Context) ComponentMapOutput
}

type ComponentMap map[string]ComponentInput

func (ComponentMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (i ComponentMap) ToComponentMapOutput() ComponentMapOutput {
	return i.ToComponentMapOutputWithContext(context.Background())
}

func (i ComponentMap) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ComponentMapOutput)
}

type ComponentOutput struct{ *pulumi.OutputState }

func (ComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Component)(nil)).Elem()
}

func (o ComponentOutput) ToComponentOutput() ComponentOutput {
	return o
}

func (o ComponentOutput) ToComponentOutputWithContext(ctx context.Context) ComponentOutput {
	return o
}

// optinal description to add details of the resource
func (o ComponentOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Component) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// The input parameters based on hubcomponent spec
func (o ComponentOutput) Inputs() ComponentInputTypeArrayOutput {
	return o.ApplyT(func(v *Component) ComponentInputTypeArrayOutput { return v.Inputs }).(ComponentInputTypeArrayOutput)
}

// the name of the component to be created
func (o ComponentOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// [NAME-VERSION] the version of the hub component
func (o ComponentOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Component) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

type ComponentArrayOutput struct{ *pulumi.OutputState }

func (ComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Component)(nil)).Elem()
}

func (o ComponentArrayOutput) ToComponentArrayOutput() ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) ToComponentArrayOutputWithContext(ctx context.Context) ComponentArrayOutput {
	return o
}

func (o ComponentArrayOutput) Index(i pulumi.IntInput) ComponentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Component {
		return vs[0].([]*Component)[vs[1].(int)]
	}).(ComponentOutput)
}

type ComponentMapOutput struct{ *pulumi.OutputState }

func (ComponentMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Component)(nil)).Elem()
}

func (o ComponentMapOutput) ToComponentMapOutput() ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) ToComponentMapOutputWithContext(ctx context.Context) ComponentMapOutput {
	return o
}

func (o ComponentMapOutput) MapIndex(k pulumi.StringInput) ComponentOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Component {
		return vs[0].(map[string]*Component)[vs[1].(string)]
	}).(ComponentOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentInput)(nil)).Elem(), &Component{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentArrayInput)(nil)).Elem(), ComponentArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ComponentMapInput)(nil)).Elem(), ComponentMap{})
	pulumi.RegisterOutputType(ComponentOutput{})
	pulumi.RegisterOutputType(ComponentArrayOutput{})
	pulumi.RegisterOutputType(ComponentMapOutput{})
}