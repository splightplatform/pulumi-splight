PROJECT_NAME := splight Package

SHELL            := /bin/bash
PACK             := splight
PROJECT          := github.com/splightplatform/pulumi-splight
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell cat version)

WORKING_DIR     := $(shell pwd)

OS := $(shell uname)

.PHONY: development provider build_sdks  build_python cleanup

# TODO: review for dev.md
development:: install_plugins provider lint_provider build_sdks install_sdks cleanup # Build the provider & SDKs for a development environment

# Required for the codegen action that runs in pulumi/pulumi and pulumi/pulumi-terraform-bridge
# build:: install_plugins provider build_sdks install_sdks
# only_build:: build

tfgen:: 
	# Build the brigdge
	# Set the value for the version at compile time
	cd provider && go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN}

	# Run the bridge and generate/update the Pulumi schema from the Terraform schema
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}

# TODO: improve
# TODO: reorder cmd and pkg folders, i dont like that structure
provider:: 
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_python

# TODO: why the repeated command names?
# Convert the provider version to a Pypi compatible version
build_python:: PYPI_VERSION := $(shell pulumictl convert-version --language python --version "$(VERSION)")

# TODO: debug and see whats necessary
build_python::
	$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/
	cd sdk/python/ && \
        cp ../../README.md . && \
        python3 setup.py clean --all 2>/dev/null && \
        rm -rf ./bin/ ../python.bin/ && cp -R . ../python.bin && mv ../python.bin ./bin && \
        sed -i.bak -e 's/^VERSION = .*/VERSION = "$(PYPI_VERSION)"/g' -e 's/^PLUGIN_VERSION = .*/PLUGIN_VERSION = "$(VERSION)"/g' ./bin/setup.py && \
        rm ./bin/setup.py.bak && \
        cd ./bin && python3 setup.py build sdist

cleanup:: # cleans up the temporary directory
	rm -r $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go

clean::
	rm -rf sdk/{python}
