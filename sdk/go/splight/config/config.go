// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
	"github.com/splightplatform/pulumi-splight/sdk/go/splight/internal"
)

var _ = internal.GetEnvOrDefault

func GetHostname(ctx *pulumi.Context) string {
	return config.Get(ctx, "splight:hostname")
}
func GetToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "splight:token")
}
