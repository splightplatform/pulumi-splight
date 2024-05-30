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

.PHONY: tfgen provider build_python cleanup

tfgen:: 
	# Build the brigdge
	# Set the value for the version at compile time
	cd provider && go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN}

	# Run the bridge and generate/update the Pulumi schema from the Terraform schema
	$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}

provider:: 
	(cd provider && go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER})

build_sdks:: install_plugins provider build_python

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

clean::
	rm -rf $(WORKING_DIR)/bin
	rm -f provider/cmd/${PROVIDER}/schema.go
	rm -rf sdk/python
