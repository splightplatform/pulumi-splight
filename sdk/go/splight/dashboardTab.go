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
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/splightplatform/pulumi-splight/sdk/go/splight"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			_, err := splight.NewDashboardTab(ctx, "dashboardTabTest", &splight.DashboardTabArgs{
//				Dashboard: pulumi.String("1234-1234-1234-1234"),
//				Order:     pulumi.Int(0),
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
// $ pulumi import splight:index/dashboardTab:DashboardTab [options] splight_dashboard_tab.<name> <dashboard_tab_id>
// ```
type DashboardTab struct {
	pulumi.CustomResourceState

	// dashboard id where to place it
	Dashboard pulumi.StringOutput `pulumi:"dashboard"`
	// name for the tab
	Name pulumi.StringOutput `pulumi:"name"`
	// order within the dashboard
	Order pulumi.IntPtrOutput `pulumi:"order"`
}

// NewDashboardTab registers a new resource with the given unique name, arguments, and options.
func NewDashboardTab(ctx *pulumi.Context,
	name string, args *DashboardTabArgs, opts ...pulumi.ResourceOption) (*DashboardTab, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Dashboard == nil {
		return nil, errors.New("invalid value for required argument 'Dashboard'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource DashboardTab
	err := ctx.RegisterResource("splight:index/dashboardTab:DashboardTab", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetDashboardTab gets an existing DashboardTab resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetDashboardTab(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DashboardTabState, opts ...pulumi.ResourceOption) (*DashboardTab, error) {
	var resource DashboardTab
	err := ctx.ReadResource("splight:index/dashboardTab:DashboardTab", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering DashboardTab resources.
type dashboardTabState struct {
	// dashboard id where to place it
	Dashboard *string `pulumi:"dashboard"`
	// name for the tab
	Name *string `pulumi:"name"`
	// order within the dashboard
	Order *int `pulumi:"order"`
}

type DashboardTabState struct {
	// dashboard id where to place it
	Dashboard pulumi.StringPtrInput
	// name for the tab
	Name pulumi.StringPtrInput
	// order within the dashboard
	Order pulumi.IntPtrInput
}

func (DashboardTabState) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardTabState)(nil)).Elem()
}

type dashboardTabArgs struct {
	// dashboard id where to place it
	Dashboard string `pulumi:"dashboard"`
	// name for the tab
	Name *string `pulumi:"name"`
	// order within the dashboard
	Order *int `pulumi:"order"`
}

// The set of arguments for constructing a DashboardTab resource.
type DashboardTabArgs struct {
	// dashboard id where to place it
	Dashboard pulumi.StringInput
	// name for the tab
	Name pulumi.StringPtrInput
	// order within the dashboard
	Order pulumi.IntPtrInput
}

func (DashboardTabArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dashboardTabArgs)(nil)).Elem()
}

type DashboardTabInput interface {
	pulumi.Input

	ToDashboardTabOutput() DashboardTabOutput
	ToDashboardTabOutputWithContext(ctx context.Context) DashboardTabOutput
}

func (*DashboardTab) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardTab)(nil)).Elem()
}

func (i *DashboardTab) ToDashboardTabOutput() DashboardTabOutput {
	return i.ToDashboardTabOutputWithContext(context.Background())
}

func (i *DashboardTab) ToDashboardTabOutputWithContext(ctx context.Context) DashboardTabOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardTabOutput)
}

// DashboardTabArrayInput is an input type that accepts DashboardTabArray and DashboardTabArrayOutput values.
// You can construct a concrete instance of `DashboardTabArrayInput` via:
//
//	DashboardTabArray{ DashboardTabArgs{...} }
type DashboardTabArrayInput interface {
	pulumi.Input

	ToDashboardTabArrayOutput() DashboardTabArrayOutput
	ToDashboardTabArrayOutputWithContext(context.Context) DashboardTabArrayOutput
}

type DashboardTabArray []DashboardTabInput

func (DashboardTabArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardTab)(nil)).Elem()
}

func (i DashboardTabArray) ToDashboardTabArrayOutput() DashboardTabArrayOutput {
	return i.ToDashboardTabArrayOutputWithContext(context.Background())
}

func (i DashboardTabArray) ToDashboardTabArrayOutputWithContext(ctx context.Context) DashboardTabArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardTabArrayOutput)
}

// DashboardTabMapInput is an input type that accepts DashboardTabMap and DashboardTabMapOutput values.
// You can construct a concrete instance of `DashboardTabMapInput` via:
//
//	DashboardTabMap{ "key": DashboardTabArgs{...} }
type DashboardTabMapInput interface {
	pulumi.Input

	ToDashboardTabMapOutput() DashboardTabMapOutput
	ToDashboardTabMapOutputWithContext(context.Context) DashboardTabMapOutput
}

type DashboardTabMap map[string]DashboardTabInput

func (DashboardTabMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardTab)(nil)).Elem()
}

func (i DashboardTabMap) ToDashboardTabMapOutput() DashboardTabMapOutput {
	return i.ToDashboardTabMapOutputWithContext(context.Background())
}

func (i DashboardTabMap) ToDashboardTabMapOutputWithContext(ctx context.Context) DashboardTabMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DashboardTabMapOutput)
}

type DashboardTabOutput struct{ *pulumi.OutputState }

func (DashboardTabOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DashboardTab)(nil)).Elem()
}

func (o DashboardTabOutput) ToDashboardTabOutput() DashboardTabOutput {
	return o
}

func (o DashboardTabOutput) ToDashboardTabOutputWithContext(ctx context.Context) DashboardTabOutput {
	return o
}

// dashboard id where to place it
func (o DashboardTabOutput) Dashboard() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardTab) pulumi.StringOutput { return v.Dashboard }).(pulumi.StringOutput)
}

// name for the tab
func (o DashboardTabOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *DashboardTab) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// order within the dashboard
func (o DashboardTabOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DashboardTab) pulumi.IntPtrOutput { return v.Order }).(pulumi.IntPtrOutput)
}

type DashboardTabArrayOutput struct{ *pulumi.OutputState }

func (DashboardTabArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*DashboardTab)(nil)).Elem()
}

func (o DashboardTabArrayOutput) ToDashboardTabArrayOutput() DashboardTabArrayOutput {
	return o
}

func (o DashboardTabArrayOutput) ToDashboardTabArrayOutputWithContext(ctx context.Context) DashboardTabArrayOutput {
	return o
}

func (o DashboardTabArrayOutput) Index(i pulumi.IntInput) DashboardTabOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *DashboardTab {
		return vs[0].([]*DashboardTab)[vs[1].(int)]
	}).(DashboardTabOutput)
}

type DashboardTabMapOutput struct{ *pulumi.OutputState }

func (DashboardTabMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*DashboardTab)(nil)).Elem()
}

func (o DashboardTabMapOutput) ToDashboardTabMapOutput() DashboardTabMapOutput {
	return o
}

func (o DashboardTabMapOutput) ToDashboardTabMapOutputWithContext(ctx context.Context) DashboardTabMapOutput {
	return o
}

func (o DashboardTabMapOutput) MapIndex(k pulumi.StringInput) DashboardTabOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *DashboardTab {
		return vs[0].(map[string]*DashboardTab)[vs[1].(string)]
	}).(DashboardTabOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardTabInput)(nil)).Elem(), &DashboardTab{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardTabArrayInput)(nil)).Elem(), DashboardTabArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DashboardTabMapInput)(nil)).Elem(), DashboardTabMap{})
	pulumi.RegisterOutputType(DashboardTabOutput{})
	pulumi.RegisterOutputType(DashboardTabArrayOutput{})
	pulumi.RegisterOutputType(DashboardTabMapOutput{})
}
