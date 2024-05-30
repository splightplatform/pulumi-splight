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

.PHONY: provider build_sdks build_nodejs build_dotnet build_go build_python clean

tfgen::
	@cd provider && \
	go mod tidy && \
	go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN}
	@$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}

provider:: tfgen
	@cd provider && \
	go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER}

build_nodejs::
	@$(WORKING_DIR)/bin/$(TFGEN) nodejs --overlays provider/overlays/nodejs --out sdk/nodejs/

build_python::
	@$(WORKING_DIR)/bin/$(TFGEN) python --overlays provider/overlays/python --out sdk/python/

build_dotnet::
	@$(WORKING_DIR)/bin/$(TFGEN) dotnet --overlays provider/overlays/dotnet --out sdk/dotnet/

build_go::
	@$(WORKING_DIR)/bin/$(TFGEN) go --overlays provider/overlays/go --out sdk/go/

build:: provider build_nodejs build_python build_go build_dotnet

clean::
	rm -r $(WORKING_DIR)/bin
	rm -rf sdk/{dotnet,nodejs,go,python}
