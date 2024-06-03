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

export COVERAGE_OUTPUT_DIR = $(WORKING_DIR)/.coverage

COLOR_RESET     := \033[0m
COLOR_INFO      := \033[0;32m

.PHONY: tidy tfgen schema-bridge provider sdk clean build

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

sdk:: tfgen
	@echo -e "${COLOR_INFO}Building SDK for Python${COLOR_RESET}"; \
	$(WORKING_DIR)/bin/$(TFGEN) python --out sdk/python/; \

build:
	@cd sdk/python && \
	python3 setup.py sdist

clean::
	@rm -rf $(WORKING_DIR)/bin
	@rm -f $(WORKING_DIR)/provider/cmd/${PROVIDER}/schema.json
	@echo "{}" > $(WORKING_DIR)/provider/cmd/${PROVIDER}/bridge-metadata.json
	@rm -rf sdk/python
