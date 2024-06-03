## Splight Terraform Provider


### Requirements

- Go (1.21.x)
- Python (3.9.x)

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

You must generate the SDK and install it each time you make changes to the terraform provider.
