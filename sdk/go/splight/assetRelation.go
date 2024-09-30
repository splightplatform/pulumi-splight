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
// $ pulumi import splight:index/assetRelation:AssetRelation [options] splight_relation.<name> <relation_id>
// ```
type AssetRelation struct {
	pulumi.CustomResourceState

	// asset where the relation origins
	Asset AssetRelationAssetOutput `pulumi:"asset"`
	// relation description
	Description pulumi.StringPtrOutput `pulumi:"description"`
	// relation name
	Name pulumi.StringOutput `pulumi:"name"`
	// target asset of the relation
	RelatedAsset AssetRelationRelatedAssetOutput `pulumi:"relatedAsset"`
	// kind of the target relation asset
	RelatedAssetKind AssetRelationRelatedAssetKindOutput `pulumi:"relatedAssetKind"`
}

// NewAssetRelation registers a new resource with the given unique name, arguments, and options.
func NewAssetRelation(ctx *pulumi.Context,
	name string, args *AssetRelationArgs, opts ...pulumi.ResourceOption) (*AssetRelation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Asset == nil {
		return nil, errors.New("invalid value for required argument 'Asset'")
	}
	if args.RelatedAsset == nil {
		return nil, errors.New("invalid value for required argument 'RelatedAsset'")
	}
	if args.RelatedAssetKind == nil {
		return nil, errors.New("invalid value for required argument 'RelatedAssetKind'")
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource AssetRelation
	err := ctx.RegisterResource("splight:index/assetRelation:AssetRelation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetAssetRelation gets an existing AssetRelation resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetAssetRelation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AssetRelationState, opts ...pulumi.ResourceOption) (*AssetRelation, error) {
	var resource AssetRelation
	err := ctx.ReadResource("splight:index/assetRelation:AssetRelation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering AssetRelation resources.
type assetRelationState struct {
	// asset where the relation origins
	Asset *AssetRelationAsset `pulumi:"asset"`
	// relation description
	Description *string `pulumi:"description"`
	// relation name
	Name *string `pulumi:"name"`
	// target asset of the relation
	RelatedAsset *AssetRelationRelatedAsset `pulumi:"relatedAsset"`
	// kind of the target relation asset
	RelatedAssetKind *AssetRelationRelatedAssetKind `pulumi:"relatedAssetKind"`
}

type AssetRelationState struct {
	// asset where the relation origins
	Asset AssetRelationAssetPtrInput
	// relation description
	Description pulumi.StringPtrInput
	// relation name
	Name pulumi.StringPtrInput
	// target asset of the relation
	RelatedAsset AssetRelationRelatedAssetPtrInput
	// kind of the target relation asset
	RelatedAssetKind AssetRelationRelatedAssetKindPtrInput
}

func (AssetRelationState) ElementType() reflect.Type {
	return reflect.TypeOf((*assetRelationState)(nil)).Elem()
}

type assetRelationArgs struct {
	// asset where the relation origins
	Asset AssetRelationAsset `pulumi:"asset"`
	// relation description
	Description *string `pulumi:"description"`
	// relation name
	Name *string `pulumi:"name"`
	// target asset of the relation
	RelatedAsset AssetRelationRelatedAsset `pulumi:"relatedAsset"`
	// kind of the target relation asset
	RelatedAssetKind AssetRelationRelatedAssetKind `pulumi:"relatedAssetKind"`
}

// The set of arguments for constructing a AssetRelation resource.
type AssetRelationArgs struct {
	// asset where the relation origins
	Asset AssetRelationAssetInput
	// relation description
	Description pulumi.StringPtrInput
	// relation name
	Name pulumi.StringPtrInput
	// target asset of the relation
	RelatedAsset AssetRelationRelatedAssetInput
	// kind of the target relation asset
	RelatedAssetKind AssetRelationRelatedAssetKindInput
}

func (AssetRelationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*assetRelationArgs)(nil)).Elem()
}

type AssetRelationInput interface {
	pulumi.Input

	ToAssetRelationOutput() AssetRelationOutput
	ToAssetRelationOutputWithContext(ctx context.Context) AssetRelationOutput
}

func (*AssetRelation) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetRelation)(nil)).Elem()
}

func (i *AssetRelation) ToAssetRelationOutput() AssetRelationOutput {
	return i.ToAssetRelationOutputWithContext(context.Background())
}

func (i *AssetRelation) ToAssetRelationOutputWithContext(ctx context.Context) AssetRelationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetRelationOutput)
}

// AssetRelationArrayInput is an input type that accepts AssetRelationArray and AssetRelationArrayOutput values.
// You can construct a concrete instance of `AssetRelationArrayInput` via:
//
//	AssetRelationArray{ AssetRelationArgs{...} }
type AssetRelationArrayInput interface {
	pulumi.Input

	ToAssetRelationArrayOutput() AssetRelationArrayOutput
	ToAssetRelationArrayOutputWithContext(context.Context) AssetRelationArrayOutput
}

type AssetRelationArray []AssetRelationInput

func (AssetRelationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssetRelation)(nil)).Elem()
}

func (i AssetRelationArray) ToAssetRelationArrayOutput() AssetRelationArrayOutput {
	return i.ToAssetRelationArrayOutputWithContext(context.Background())
}

func (i AssetRelationArray) ToAssetRelationArrayOutputWithContext(ctx context.Context) AssetRelationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetRelationArrayOutput)
}

// AssetRelationMapInput is an input type that accepts AssetRelationMap and AssetRelationMapOutput values.
// You can construct a concrete instance of `AssetRelationMapInput` via:
//
//	AssetRelationMap{ "key": AssetRelationArgs{...} }
type AssetRelationMapInput interface {
	pulumi.Input

	ToAssetRelationMapOutput() AssetRelationMapOutput
	ToAssetRelationMapOutputWithContext(context.Context) AssetRelationMapOutput
}

type AssetRelationMap map[string]AssetRelationInput

func (AssetRelationMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssetRelation)(nil)).Elem()
}

func (i AssetRelationMap) ToAssetRelationMapOutput() AssetRelationMapOutput {
	return i.ToAssetRelationMapOutputWithContext(context.Background())
}

func (i AssetRelationMap) ToAssetRelationMapOutputWithContext(ctx context.Context) AssetRelationMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssetRelationMapOutput)
}

type AssetRelationOutput struct{ *pulumi.OutputState }

func (AssetRelationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssetRelation)(nil)).Elem()
}

func (o AssetRelationOutput) ToAssetRelationOutput() AssetRelationOutput {
	return o
}

func (o AssetRelationOutput) ToAssetRelationOutputWithContext(ctx context.Context) AssetRelationOutput {
	return o
}

// asset where the relation origins
func (o AssetRelationOutput) Asset() AssetRelationAssetOutput {
	return o.ApplyT(func(v *AssetRelation) AssetRelationAssetOutput { return v.Asset }).(AssetRelationAssetOutput)
}

// relation description
func (o AssetRelationOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssetRelation) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

// relation name
func (o AssetRelationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AssetRelation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

// target asset of the relation
func (o AssetRelationOutput) RelatedAsset() AssetRelationRelatedAssetOutput {
	return o.ApplyT(func(v *AssetRelation) AssetRelationRelatedAssetOutput { return v.RelatedAsset }).(AssetRelationRelatedAssetOutput)
}

// kind of the target relation asset
func (o AssetRelationOutput) RelatedAssetKind() AssetRelationRelatedAssetKindOutput {
	return o.ApplyT(func(v *AssetRelation) AssetRelationRelatedAssetKindOutput { return v.RelatedAssetKind }).(AssetRelationRelatedAssetKindOutput)
}

type AssetRelationArrayOutput struct{ *pulumi.OutputState }

func (AssetRelationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*AssetRelation)(nil)).Elem()
}

func (o AssetRelationArrayOutput) ToAssetRelationArrayOutput() AssetRelationArrayOutput {
	return o
}

func (o AssetRelationArrayOutput) ToAssetRelationArrayOutputWithContext(ctx context.Context) AssetRelationArrayOutput {
	return o
}

func (o AssetRelationArrayOutput) Index(i pulumi.IntInput) AssetRelationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *AssetRelation {
		return vs[0].([]*AssetRelation)[vs[1].(int)]
	}).(AssetRelationOutput)
}

type AssetRelationMapOutput struct{ *pulumi.OutputState }

func (AssetRelationMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*AssetRelation)(nil)).Elem()
}

func (o AssetRelationMapOutput) ToAssetRelationMapOutput() AssetRelationMapOutput {
	return o
}

func (o AssetRelationMapOutput) ToAssetRelationMapOutputWithContext(ctx context.Context) AssetRelationMapOutput {
	return o
}

func (o AssetRelationMapOutput) MapIndex(k pulumi.StringInput) AssetRelationOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *AssetRelation {
		return vs[0].(map[string]*AssetRelation)[vs[1].(string)]
	}).(AssetRelationOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AssetRelationInput)(nil)).Elem(), &AssetRelation{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetRelationArrayInput)(nil)).Elem(), AssetRelationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssetRelationMapInput)(nil)).Elem(), AssetRelationMap{})
	pulumi.RegisterOutputType(AssetRelationOutput{})
	pulumi.RegisterOutputType(AssetRelationArrayOutput{})
	pulumi.RegisterOutputType(AssetRelationMapOutput{})
}