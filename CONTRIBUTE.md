## Splight Terraform Provider

### Requirements

- Go (1.21.x)
- Python (3.9.x)
- Yarn (1.22.x)
- Goreleaser (1.26.2)
- Dotnet (8.0.104)

### Generation

Set the terraform-provider-splight version to the desired one in the `go.mod` file.

Install it along with the dependencies:

```bash
make tidy
```

Compile the tfgen utility:

```bash
make tfgen
```

Compile the provider:

```bash
make provider
```

Generate the Python SDK:

```
make sdk
```

### Installing the SDK

Activate the virtual environment of the Pulumi project:

```bash
source venv/bin/activate
```

Use your favourite way of installing Python packages (i.e. `pip`):

```bash
pip install sdk/python
```

Try to create infrastructure using the Pulumi CLI:

```bash
pulumi up
```

To implement changes to the Terraform provider, you need to regenerate the provider, update the SDK, and reinstall it each time.
