---
title: Splight Installation & Configuration
meta_desc: Information on how to install the Splight provider.
layout: package
---

## Installation

The Splight provider is available as a package in following Pulumi languages:

* JavaScript/TypeScript: [`@splightplatform/pulumi-splight`](https://www.npmjs.com/package/@splightplatform/pulumi-splight)
* Python: [`pulumi-splight`](https://pypi.org/project/pulumi-splight/)
* Go: [`github.com/splightplatform/pulumi-splight/sdk/go/splight`](https://pkg.go.dev/github.com/splightplatform/pulumi-splight/sdk/go/splight)
* .NET: [`Splight.Splight`](https://www.nuget.org/packages/Splight.Splight)

## Setup

The following configuration points are available for the `Splight` provider:

- `splight:hostname` - the API endpoint for the Splight platform
- `splight:token` - the API token to authenticate with the Splight platform (Splight <access_id> <secret_key>)

If these are not provided, the provider will fall back to the environment variables:
- `SPLIGHT_ACCESS_ID` - the access ID to authenticate with the Splight platform
- `SPLIGHT_SECRET_KEY` - the secret key to authenticate with the Splight platform
- `SPLIGHT_PLATFORM_API_HOST` - the API endpoint for the Splight platform

and as a last resort, it will use the active workspace configuration from the Splight CLI.
