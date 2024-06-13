# Splight Resource Provider

The Splight provider enables seamless interaction with resources supported by Splight.

To begin using the provider, you need to configure it with the appropriate credentials. Refer to the [Configuration](#configuration) section for more details.

For information on development, please refer to the [Contribute Guide](CONTRIBUTE.md).

## Installing

This package is available for several languages/platforms:

### Node.js (JavaScript/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @splightplatform/pulumi-splight
```

or `yarn`:

```bash
yarn add @splightplatform/pulumi-splight
```

### Python

To use from Python, install using `pip`:

```bash
pip install pulumi-splight
```

### Go

To use from Go, use `go get` to grab the latest version of the library:

```bash
go get github.com/splightplatform/pulumi-splight/sdk/go
```

### .NET

To use from .NET, install using `dotnet add package`:

```bash
dotnet add package Splight.Splight
```

## Configuration

The following configuration points are available for the `Splight` provider:

- `splight:hostname` - the API endpoint for the Splight platform
- `splight:token` - the API token to authenticate with the Splight platform (Splight <access_id> <secret_key>)

If these are not provided, the provider will fall back to the environment variables:
- `SPLIGHT_ACCESS_ID` - the access ID to authenticate with the Splight platform
- `SPLIGHT_SECRET_KEY` - the secret key to authenticate with the Splight platform
- `SPLIGHT_PLATFORM_API_HOST` - the API endpoint for the Splight platform

and as a last resort, it will use the active workspace configuration from the Splight CLI.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/splight/api-docs/).
