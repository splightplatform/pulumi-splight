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
//			_, err := splight.GetLines(ctx, map[string]interface{}{}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetLines(ctx *pulumi.Context, opts ...pulumi.InvokeOption) (*GetLinesResult, error) {
	opts = internal.PkgInvokeDefaultOpts(opts)
	var rv GetLinesResult
	err := ctx.Invoke("splight:index/getLines:getLines", nil, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of values returned by getLines.
type GetLinesResult struct {
	// The provider-assigned unique ID for this managed resource.
	Id   string        `pulumi:"id"`
	Tags []GetLinesTag `pulumi:"tags"`
}

func GetLinesOutput(ctx *pulumi.Context, opts ...pulumi.InvokeOption) GetLinesResultOutput {
	return pulumi.ToOutput(0).ApplyT(func(int) (GetLinesResultOutput, error) {
		opts = internal.PkgInvokeDefaultOpts(opts)
		var rv GetLinesResult
		secret, err := ctx.InvokePackageRaw("splight:index/getLines:getLines", nil, &rv, "", opts...)
		if err != nil {
			return GetLinesResultOutput{}, err
		}

		output := pulumi.ToOutput(rv).(GetLinesResultOutput)
		if secret {
			return pulumi.ToSecret(output).(GetLinesResultOutput), nil
		}
		return output, nil
	}).(GetLinesResultOutput)
}

// A collection of values returned by getLines.
type GetLinesResultOutput struct{ *pulumi.OutputState }

func (GetLinesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetLinesResult)(nil)).Elem()
}

func (o GetLinesResultOutput) ToGetLinesResultOutput() GetLinesResultOutput {
	return o
}

func (o GetLinesResultOutput) ToGetLinesResultOutputWithContext(ctx context.Context) GetLinesResultOutput {
	return o
}

// The provider-assigned unique ID for this managed resource.
func (o GetLinesResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetLinesResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetLinesResultOutput) Tags() GetLinesTagArrayOutput {
	return o.ApplyT(func(v GetLinesResult) []GetLinesTag { return v.Tags }).(GetLinesTagArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetLinesResultOutput{})
}
