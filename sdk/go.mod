// This file ensures that 'go mod' can correctly install our Golang SDK.
// It is placed in this directory because 'go mod' cannot navigate deep into repository folders.
// Additionally, placing it here prevents it from being deleted by the 'make clean' command.
module github.com/splightplatform/pulumi-splight/sdk

go 1.21

require github.com/pulumi/pulumi/sdk/v3 v3.119.0
