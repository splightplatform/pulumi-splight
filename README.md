# Splight Resource Provider

The Splight provider enables seamless interaction with resources supported by Splight.

To begin using the provider, you need to configure it with the appropriate credentials.

Refer to the [Configuration](#Configuration) section for more details.

## Installing

This package is available for Python.

Install using `pip`:

```bash
pip install pulumi_splight
```

## Configuration

The following configuration points are available for the `splight` provider:

- `splight:hostname` - the API endpoint for the Splight platform
- `splight:token` - the API token to authenticate with the Splight platform (Splight <access_id> <secret_key>)

If these are not provided, the provider will fall back to the environment variables:
- `SPLIGHT_ACCESS_ID` - the access ID to authenticate with the Splight platform
- `SPLIGHT_SECRET_KEY` - the secret key to authenticate with the Splight platform
- `SPLIGHT_PLATFORM_API_HOST` - the API endpoint for the Splight platform

and as a last resort, it will use the active workspace configuration from the Splight CLI.

## Reference

For detailed reference documentation, please visit [the Pulumi registry](https://www.pulumi.com/registry/packages/splight/api-docs/).
