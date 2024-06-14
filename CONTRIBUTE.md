## Splight Terraform Provider

### Requirements

- Go (1.21.x)
- Python (3.9.x)
- Yarn (1.22.x)
- Goreleaser (1.26.2)
- Dotnet (8.0.104)

### Building locally

Update the provider/go.mod file to specify the desired version of the [Splight Terraform Provider](https://github.com/splightplatform/terraform-provider-splight).

Use the same version number in the `version` file, formatted as X.Y.Z (without the 'v').

By doing this, its easier to track which version of the Terraform provider is used by the Pulumi provider.

TIP: Go mod also allows you to reference a local repository of the terraform provider. This makes development interactive.

For each change to made to it, you will need to recompile the Pulumi provider and SDKs by following the steps below:

#### Clean the project

This will:
- Delete the `bin` and `sdk/*` directories, which contain the compiled binaries and SDKs.
- Delete the Pulumi schema.
- Replace the bridge-metadata file with an empty dictionary (required as its embedded in the resources.go file, shared by both, the tfgen and the provider)

```bash
make clean
```

#### Sync dependencies

```bash
make tidy
```

#### Compile the tfgen utility

This tool is used to generate the Pulumi `schema.json` and `bridge-metadata.json` files, which are used to compile the Pulumi provider.

It is also used to generate the SDKs for the different languages.

```bash
make tfgen
```

#### Build the provider

Generate the files mentioned above and compile the Pulumi provider:

```bash
make provider
```

We offer an optional `make snapshot` command, which leverages Goreleaser to create a local release of the Pulumi provider.
This command is particularly useful for testing the release process prior to officially releasing a new version.
You can also use the generated binaries for different architectures and operating systems to test the provider by referencing it in your Pulumi project.

#### Generate the SDKs

```bash
make sdks
```

The `make build` command is used by the CI/CD pipeline to build the SDKs before publishing them to their respective package managers.

### Using locally

Create a new Pulumi project:

```bash
mkdir myproject
cd myproject
```

Create and activate the virtual environment of the Pulumi project:

```bash
python -m venv venv
source venv/bin/activate
```

Go into the generated SDK directory and install the Python SDK:

```bash
cd <splight-pulumi-repo>/sdk/python
pip install .
```

Go back to the Pulumi project folder and create the following files:

`Pulumi.yaml`

```yaml
name: MyProject
runtime:
  name: python
  options:
    virtualenv: venv # Path to the virtual environment
description: My Project Description
plugins:
  providers:
  - name: splight # Override the Splight provider with the local version
    path: <path_to_provider> # Should not end with a slash
```

`__main__.py`

```python
import geojson
import pulumi_splight as splight


splight.Asset(
    resource_name="MyAsset",
    name="MyAsset",
    description="My Asset Description",
    geometry=geojson.dumps(
        geojson.GeometryCollection(
            geometries=[{"type": "Point", "coordinates": [0, 0]}]
        )
    ),
)
```

Try to create infrastructure using the Pulumi CLI:

```bash
pulumi up --parallel 1 # Do not use concurrent requests
```

### Release

Create a tag that matches the version in the `version` file.

```bash
git tag -a vX.Y.Z -m "Release vX.Y.Z"
git push origin vX.Y.Z
```

This will trigger the CI/CD pipeline to:
- Use GoReleaser to build the provider and upload the binaries to a GitHub release.
- Build the SDKs and publish them to their respective package managers.
