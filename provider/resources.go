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
		LogoURL:           "",
		PluginDownloadURL: "",
		Description:       "A Pulumi package for creating and managing splight cloud resources.",
		Keywords:          []string{"splightplatform", "splight", "category/cloud"},
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
