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
//				Timezone:    pulumi.String("America/Los_Angeles"),
//				Geometry:    pulumi.String(json0),
//			})
//			if err != nil {
//				return err
//			}
//			myAttribute, err := splight.NewAssetAttribute(ctx, "myAttribute", &splight.AssetAttributeArgs{
//				Type:  pulumi.String("Number"),
//				Unit:  pulumi.String("meters"),
//				Asset: myAsset.ID(),
//			})
//			if err != nil {
//				return err
//			}
//			tmpJSON1, err := json.Marshal(1)
//			if err != nil {
//				return err
//			}
//			json1 := string(tmpJSON1)
//			myAction, err := splight.NewAction(ctx, "myAction", &splight.ActionArgs{
//				Asset: &splight.ActionAssetArgs{
//					Id:   myAsset.ID(),
//					Name: myAsset.Name,
//				},
//				Setpoints: splight.ActionSetpointArray{
//					&splight.ActionSetpointArgs{
//						Value: pulumi.String(json1),
//						Attribute: &splight.ActionSetpointAttributeArgs{
//							Id:   myAttribute.ID(),
//							Name: myAttribute.Name,
//						},
//					},
//				},
//			})
//			if err != nil {
//				return err
//			}
//			_, err = splight.NewCommand(ctx, "myCommand", &splight.CommandArgs{
//				Actions: splight.CommandActionArray{
//					&splight.CommandActionArgs{
//						Id:   myAction.ID(),
//						Name: myAction.Name,
//						Asset: &splight.CommandActionAssetArgs{
//							Id:   myAsset.ID(),
//							Name: myAsset.Name,
//						},
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
// $ pulumi import splight:index/command:Command [options] splight_command.<name> <command_id>
// ```
type Command struct {
	pulumi.CustomResourceState

	// command actions
	Actions CommandActionArrayOutput `pulumi:"actions"`
	// the description of the command to be created
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// the name of the command to be created
	Name pulumi.StringOutput `pulumi:"name"`
}

// NewCommand registers a new resource with the given unique name, arguments, and options.
func NewCommand(ctx *pulumi.Context,
	name string, args *CommandArgs, opts ...pulumi.ResourceOption) (*Command, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Actions == nil {
		return nil, errors.New("invalid value for required argument 'Actions'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Command
	err := ctx.RegisterResource("splight:index/command:Command", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetCommand gets an existing Command resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetCommand(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CommandState, opts ...pulumi.ResourceOption) (*Command, error) {
	var resource Command
	err := ctx.ReadResource("splight:index/command:Command", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Command resources.
type commandState struct {
	// command actions
	Actions []CommandAction `pulumi:"actions"`
	// the description of the command to be created
	Description *string `pulumi:"description"`
	// the name of the command to be created
	Name *string `pulumi:"name"`
}

type CommandState struct {
	// command actions
	Actions CommandActionArrayInput
	// the description of the command to be created
	Description pulumi.StringPtrInput
	// the name of the command to be created
	Name pulumi.StringPtrInput
}

func (CommandState) ElementType() reflect.Type {
	return reflect.TypeOf((*commandState)(nil)).Elem()
}

type commandArgs struct {
	// command actions
	Actions []CommandAction `pulumi:"actions"`
	// the description of the command to be created
	Description *string `pulumi:"description"`
	// the name of the command to be created
	Name *string `pulumi:"name"`
}

// The set of arguments for constructing a Command resource.
type CommandArgs struct {
	// command actions
	Actions CommandActionArrayInput
	// the description of the command to be created
	Description pulumi.StringPtrInput
	// the name of the command to be created
	Name pulumi.StringPtrInput
}

func (CommandArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*commandArgs)(nil)).Elem()
}

type CommandInput interface {
	pulumi.Input

	ToCommandOutput() CommandOutput
	ToCommandOutputWithContext(ctx context.Context) CommandOutput
}

func (*Command) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (i *Command) ToCommandOutput() CommandOutput {
	return i.ToCommandOutputWithContext(context.Background())
}

func (i *Command) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandOutput)
}

// CommandArrayInput is an input type that accepts CommandArray and CommandArrayOutput values.
// You can construct a concrete instance of `CommandArrayInput` via:
//
//	CommandArray{ CommandArgs{...} }
type CommandArrayInput interface {
	pulumi.Input

	ToCommandArrayOutput() CommandArrayOutput
	ToCommandArrayOutputWithContext(context.Context) CommandArrayOutput
}

type CommandArray []CommandInput

func (CommandArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Command)(nil)).Elem()
}

func (i CommandArray) ToCommandArrayOutput() CommandArrayOutput {
	return i.ToCommandArrayOutputWithContext(context.Background())
}

func (i CommandArray) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandArrayOutput)
}

// CommandMapInput is an input type that accepts CommandMap and CommandMapOutput values.
// You can construct a concrete instance of `CommandMapInput` via:
//
//	CommandMap{ "key": CommandArgs{...} }
type CommandMapInput interface {
	pulumi.Input

	ToCommandMapOutput() CommandMapOutput
	ToCommandMapOutputWithContext(context.Context) CommandMapOutput
}

type CommandMap map[string]CommandInput

func (CommandMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Command)(nil)).Elem()
}

func (i CommandMap) ToCommandMapOutput() CommandMapOutput {
	return i.ToCommandMapOutputWithContext(context.Background())
}

func (i CommandMap) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CommandMapOutput)
}

type CommandOutput struct{ *pulumi.OutputState }

func (CommandOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Command)(nil)).Elem()
}

func (o CommandOutput) ToCommandOutput() CommandOutput {
	return o
}

func (o CommandOutput) ToCommandOutputWithContext(ctx context.Context) CommandOutput {
	return o
}

// command actions
func (o CommandOutput) Actions() CommandActionArrayOutput {
	return o.ApplyT(func(v *Command) CommandActionArrayOutput { return v.Actions }).(CommandActionArrayOutput)
}

// the description of the command to be created
func (o CommandOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Command) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// the name of the command to be created
func (o CommandOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Command) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

type CommandArrayOutput struct{ *pulumi.OutputState }

func (CommandArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*Command)(nil)).Elem()
}

func (o CommandArrayOutput) ToCommandArrayOutput() CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) ToCommandArrayOutputWithContext(ctx context.Context) CommandArrayOutput {
	return o
}

func (o CommandArrayOutput) Index(i pulumi.IntInput) CommandOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *Command {
		return vs[0].([]*Command)[vs[1].(int)]
	}).(CommandOutput)
}

type CommandMapOutput struct{ *pulumi.OutputState }

func (CommandMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*Command)(nil)).Elem()
}

func (o CommandMapOutput) ToCommandMapOutput() CommandMapOutput {
	return o
}

func (o CommandMapOutput) ToCommandMapOutputWithContext(ctx context.Context) CommandMapOutput {
	return o
}

func (o CommandMapOutput) MapIndex(k pulumi.StringInput) CommandOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *Command {
		return vs[0].(map[string]*Command)[vs[1].(string)]
	}).(CommandOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*CommandInput)(nil)).Elem(), &Command{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommandArrayInput)(nil)).Elem(), CommandArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*CommandMapInput)(nil)).Elem(), CommandMap{})
	pulumi.RegisterOutputType(CommandOutput{})
	pulumi.RegisterOutputType(CommandArrayOutput{})
	pulumi.RegisterOutputType(CommandMapOutput{})
}
