// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package splight

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
	"github.com/splightplatform/pulumi-splight/sdk/go/splight/internal"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "splight:index/alert:Alert":
		r = &Alert{}
	case "splight:index/asset:Asset":
		r = &Asset{}
	case "splight:index/assetAttribute:AssetAttribute":
		r = &AssetAttribute{}
	case "splight:index/assetMetadata:AssetMetadata":
		r = &AssetMetadata{}
	case "splight:index/component:Component":
		r = &Component{}
	case "splight:index/componentRoutine:ComponentRoutine":
		r = &ComponentRoutine{}
	case "splight:index/dashboard:Dashboard":
		r = &Dashboard{}
	case "splight:index/dashboardChart:DashboardChart":
		r = &DashboardChart{}
	case "splight:index/dashboardTab:DashboardTab":
		r = &DashboardTab{}
	case "splight:index/file:File":
		r = &File{}
	case "splight:index/fileFolder:FileFolder":
		r = &FileFolder{}
	case "splight:index/function:Function":
		r = &Function{}
	case "splight:index/node:Node":
		r = &Node{}
	case "splight:index/secret:Secret":
		r = &Secret{}
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	err = ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:splight" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	r := &Provider{}
	err := ctx.RegisterResource(typ, name, nil, r, pulumi.URN_(urn))
	return r, err
}

func init() {
	version, err := internal.PkgVersion()
	if err != nil {
		version = semver.Version{Major: 1}
	}
	pulumi.RegisterResourceModule(
		"splight",
		"index/alert",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/asset",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/assetAttribute",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/assetMetadata",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/component",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/componentRoutine",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/dashboard",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/dashboardChart",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/dashboardTab",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/file",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/fileFolder",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/function",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/node",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"splight",
		"index/secret",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"splight",
		&pkg{version},
	)
}
