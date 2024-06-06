PROJECT_NAME := splight Package

SHELL            := /bin/bash
PACK             := splight
PROJECT          := github.com/splightplatform/pulumi-splight
PROVIDER_PATH    := provider
VERSION_PATH     := ${PROVIDER_PATH}/pkg/version.Version
LANGUAGES	 := nodejs python dotnet go

TFGEN           := pulumi-tfgen-${PACK}
PROVIDER        := pulumi-resource-${PACK}
VERSION         := $(shell cat version)

WORKING_DIR     := $(shell pwd)

export COVERAGE_OUTPUT_DIR = $(WORKING_DIR)/.coverage

COLOR_RESET     := \033[0m
COLOR_INFO      := \033[0;32m

.PHONY: tidy tfgen schema-bridge provider sdks clean build build-nodejs build-python build-dotnet provider

tidy::
	@cd provider && \
	go mod tidy

tfgen:: tidy
	@cd provider && \
	go build -o $(WORKING_DIR)/bin/${TFGEN} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${TFGEN}

schema-bridge:: tfgen
	@$(WORKING_DIR)/bin/${TFGEN} schema --out provider/cmd/${PROVIDER}

provider:: schema-bridge
	@cd provider && \
	go build -o $(WORKING_DIR)/bin/${PROVIDER} -ldflags "-X ${PROJECT}/${VERSION_PATH}=${VERSION}" ${PROJECT}/${PROVIDER_PATH}/cmd/${PROVIDER}

sdks:: tfgen
	@for sdk in $(LANGUAGES); do \
		echo -e "${COLOR_INFO}Generating SDK for $$sdk${COLOR_RESET}"; \
		$(WORKING_DIR)/bin/$(TFGEN) $$sdk --out sdk/$$sdk/; \
	done

build-python: 
	@cd sdk/python && \
	python3 setup.py sdist

build-nodejs:
	@cd sdk/nodejs && \
	yarn install && \
        yarn build && \
	cp package.json yarn.lock bin/

build-dotnet:
	@cd sdk/dotnet/ && \
        dotnet build /p:Version=${VERSION}

# TODO: missing readme for each package
build: build-python build-nodejs build-dotnet # Used by CI/CD

snapshot: provider
	@goreleaser --snapshot --clean

clean::
	@rm -rf $(WORKING_DIR)/bin
	@rm -f $(WORKING_DIR)/provider/cmd/${PROVIDER}/schema.json
	@echo "{}" > $(WORKING_DIR)/provider/cmd/${PROVIDER}/bridge-metadata.json
	@rm -rf sdk/{nodejs,python}
