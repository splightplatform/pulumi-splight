// Copyright 2016-2024, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package splight

import (
	"path"

	// Allow embedding bridge-metadata.json in the provider.
	_ "embed"

	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"

	splight "github.com/splightplatform/terraform-provider-splight/provider"

	"github.com/splightplatform/pulumi-splight/provider/pkg/version"
)

const (
	// Default name of the package in the package
	// registries for nodejs and python:
	mainPkg  = "splight"
	mainMod  = "index" // the splight module
	nugetOrg = "Splight"
)

//go:embed cmd/pulumi-resource-splight/bridge-metadata.json
var metadata []byte

// Provider returns additional overlaid schema and metadata associated with the provider.
func Provider() tfbridge.ProviderInfo {
	// Create a Pulumi provider mapping
	provider := tfbridge.ProviderInfo{
		// Instantiate the Terraform provider
		P: shimv2.NewProvider(splight.Provider()),

		Name:              "splight",
		Version:           version.Version,
		Publisher:         "splightplatform",
		LogoURL:           "https://splight-logos.s3.us-east-1.amazonaws.com/splight.png?response-content-disposition=inline&X-Amz-Security-Token=IQoJb3JpZ2luX2VjEAMaCXVzLWVhc3QtMSJHMEUCIH8TGLuih3jbzwNoLz%2FXm%2Bm%2F34tEf3flrlGy0XznSU4MAiEA6AaOBRggCgHFV6rBxcHOsIjJGqBt5ps4MheiiKOvEJ8qhAMInP%2F%2F%2F%2F%2F%2F%2F%2F%2F%2FARACGgw2MDkwNjc1OTg4NzciDHT%2FJ2KGaR7rh9pTzyrYAk6t8C3hmWJ2ECrlDckzcgBgnVa3KtIEvakQjPWcZc2iJfV4yQsIVlxaCNGiu0L9IjQWXfBRJf1nGb2kMcnqkYXmkyuqsp9iQRdCe5hkC%2FhscAvlI50olc98Df159G3bx87aQlUCAHJkuupfK89MQulriKpRJeP4CxrI7nGaZap1tvSqPr0XAGLUoi%2FxIJOPnmtvgBR97ZgslOqtPD57kGlVnWirZOnJofmrrmgM0CohIR3QWxflbgNTiGcsUDIvxaIZDSKn%2FV1N3kJv8uLb6Q%2F1i3eFNzUEwpIxPROC2uG2B%2BTvId%2FDiVUiDA9AeX3HP2tosexL2gtKlKUPnJwpSQwNYnQCGeJu3aIzSZ%2F4fIAjkXN4ff01tEN5HCqwVXzeG%2BibOTRPWsICU%2F5Qn%2FX03P5UT8uCG1MupB02iyPpilic6%2BH8yqjBwp9VPre8jZLRNAwx8N7cE4KoMIPirrMGOrMCOo5MOQhSLKtrvQUrqAOaZa1YupPy%2BtbC%2FM8fTPe%2BD4JkBM86vmquZnRIfa1EOKB7gTH%2F3U8adSLjA8ZwzTEiZXB5%2B9hSHEIqeALcayydJfMAmFOBBTO%2Fs7K2nyNG4TACw9ZqdIVR0frafRU61CkzEgr2PTvxfwAZ2SHvy69almj2uutmjoX%2BoxmVsajWi9tFHKXnJymuMBX20%2FnQtV5IDrDA%2B3OUHJ5s%2FXs5Ket55VHzUxhe81Lma%2FJbaBhDnVLWMYq8W3ooKc%2B5cfEVTwmT%2BnHpMFzv56%2Bi3PFu%2Bm8vMUDqbp0CrRjzO%2BCJwJfeJByDe9yJx1iRAiSNMtROGJP2fRQzmVlulRAv7ERRuyZSbLfX5X6X11OTTYiz9l7REVf9410OHyI4Qa%2FIMtm5EwNfvRmzlQ%3D%3D&X-Amz-Algorithm=AWS4-HMAC-SHA256&X-Amz-Date=20240614T140820Z&X-Amz-SignedHeaders=host&X-Amz-Expires=300&X-Amz-Credential=ASIAY3T2CAQOT7DTUQWW%2F20240614%2Fus-east-1%2Fs3%2Faws4_request&X-Amz-Signature=1c5e607f016f3f27525dfc6b353545244610411fa62c25a92620d58a37a936d5",
		Description:       "A Pulumi package for creating and managing Splight resources.",
		Keywords:          []string{"splightplatform", "splight", "category/Infrastructure"},
		License:           "Apache-2.0",
		Homepage:          "https://www.splight-ai.com",
		Repository:        "https://github.com/splightplatform/pulumi-splight",
		GitHubOrg:         "splightplatform",
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
		Config:            map[string]*tfbridge.SchemaInfo{},
		JavaScript: &tfbridge.JavaScriptInfo{
			PackageName:          "@splightplatform/pulumi-splight",
			RespectSchemaVersion: true,
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0", // So we can access strongly typed node definitions.
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			PackageName:          "pulumi_splight",
			RespectSchemaVersion: true,
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			RespectSchemaVersion: true,
			ImportBasePath: path.Join(
				"github.com/splightplatform/pulumi-splight/sdk/",
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			RespectSchemaVersion: true,
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			RootNamespace: nugetOrg,
		},
	}

	provider.MustComputeTokens(tokens.SingleModule("splight_", mainMod, tokens.MakeStandard(mainPkg)))
	provider.MustApplyAutoAliases()
	provider.SetAutonaming(255, "-")

	return provider
}
