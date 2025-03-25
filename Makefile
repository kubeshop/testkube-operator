
# Image URL to use all building/pushing image targets
IMG ?= controller:latest
# Produce CRDs that work back to Kubernetes 1.11 (no version conversion)
CRD_OPTIONS ?= "crd:allowDangerousTypes=true"

# Get the currently used golang install path (in GOPATH/bin, unless GOBIN is set)
ifeq (,$(shell go env GOBIN))
GOBIN=$(shell go env GOPATH)/bin
else
GOBIN=$(shell go env GOBIN)
endif

# Setting SHELL to bash allows bash commands to be executed by recipes.
# This is a requirement for 'setup-envtest.sh' in the test target.
# Options are set to exit when a recipe line exits non-zero or a piped command fails.
SHELL = /usr/bin/env bash -o pipefail
.SHELLFLAGS = -ec

all: build

##@ General

# The help target prints out all targets with their descriptions organized
# beneath their categories. The categories are represented by '##@' and the
# target descriptions by '##'. The awk commands is responsible for reading the
# entire set of makefiles included in this invocation, looking for lines of the
# file as xyz: ## something, and then pretty-format the target and help. Then,
# if there's a line with ##@ something, that gets pretty-printed as a category.
# More info on the usage of ANSI control characters for terminal formatting:
# https://en.wikipedia.org/wiki/ANSI_escape_code#SGR_parameters
# More info on the awk command:
# http://linuxcommand.org/lc3_adv_awk.php

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

manifests: manifests-create manifests-clean ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects. Delete unnecessary entries.

manifests-create: controller-gen ## Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
	$(CONTROLLER_GEN) $(CRD_OPTIONS) rbac:roleName=manager-role webhook paths="./..." output:crd:artifacts:config=config/crd/bases

manifests-clean: yq
	@for file in testworkflows.testkube.io_testworkflows.yaml testworkflows.testkube.io_testworkflowtemplates.yaml testworkflows.testkube.io_testworkflowexecutions.yaml; do \
		for key in securityContext volumes dnsPolicy affinity tolerations hostAliases dnsConfig topologySpreadConstraints schedulingGates resourceClaims volumeMounts fieldRef resourceFieldRef configMapKeyRef secretKeyRef pvcs matchExpressions matchLabels env envFrom readinessProbe; do \
			yq --no-colors -i "del(.. | select(has(\"$$key\")).$$key | .. | select(has(\"description\")).description)" "config/crd/bases/$$file"; \
		done; \
		yq --no-colors -i \
		'with(..; . | select(has("additionalProperties")) | select(.additionalProperties | has("type")) | select(.additionalProperties.type == "dynamicList") | \
			.["x-kubernetes-preserve-unknown-fields"] = true | \
			del(.additionalProperties) \
		) | \
		with(..; . | select(has("properties")) | select(.properties | to_entries | filter(.value | has("type")) | filter(.value.type == "dynamicList") | length > 0) | \
			.["x-kubernetes-preserve-unknown-fields"] = true | \
			del(.properties) \
		)' \
		"config/crd/bases/$$file"; \
	done

generate: controller-gen ## Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
	$(CONTROLLER_GEN) object:headerFile="hack/boilerplate.go.txt" paths="./..."
	$(CONTROLLER_GEN) object:headerFile="pkg/tcl/header.txt" paths="./pkg/tcl/..."

fmt: ## Run go fmt against code.
	go fmt ./...

vet: ## Run go vet against code.
	go vet ./...

ENVTEST_ASSETS_DIR=$(shell pwd)/testbin
test: manifests generate fmt vet ## Run tests.
	mkdir -p ${ENVTEST_ASSETS_DIR}
	test -f ${ENVTEST_ASSETS_DIR}/setup-envtest.sh || curl -sSLo ${ENVTEST_ASSETS_DIR}/setup-envtest.sh https://raw.githubusercontent.com/kubernetes-sigs/controller-runtime/v0.8.3/hack/setup-envtest.sh
	source ${ENVTEST_ASSETS_DIR}/setup-envtest.sh; fetch_envtest_tools $(ENVTEST_ASSETS_DIR); setup_envtest_env $(ENVTEST_ASSETS_DIR); go test ./... -coverprofile cover.out

##@ Build

build: generate fmt vet ## Build manager binary.
	go build -o bin/manager cmd/main.go

.PHONY: build-local-linux
build-local-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o dist/manager cmd/main.go

.PHONY: docker-build-local
docker-build-local: build-local-linux
	docker build -t controller:latest -f local.Dockerfile .

.PHONY: kind-load-local
kind-load-local: docker-build-local
	kind load docker-image controller:latest --name testkube

docker-build: test ## Build docker image with the manager.
	docker build -t ${IMG} .

docker-push: ## Push docker image with the manager.
	docker push ${IMG}

##@ Build

run: manifests generate fmt vet ## Run a controller from your host.
	go run ./cmd/main.go

run-no-webhook: manifests generate fmt vet ## Run a controller from your host.
	ENABLE_WEBHOOKS=false go run ./cmd/main.go

##@ Deployment

install: manifests kustomize ## Install CRDs into the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl apply -f -

uninstall: manifests kustomize ## Uninstall CRDs from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/crd | kubectl delete -f -

deploy: manifests kustomize ## Deploy controller to the K8s cluster specified in ~/.kube/config.
	cd config/manager && $(KUSTOMIZE) edit set image controller=${IMG}
	$(KUSTOMIZE) build config/default | kubectl apply -f -

undeploy: ## Undeploy controller from the K8s cluster specified in ~/.kube/config.
	$(KUSTOMIZE) build config/default | kubectl delete -f -

##@ Setup

CONTROLLER_GEN = $(shell pwd)/bin/controller-gen
CONTROLLER_GEN_VERSION = v0.17.2
controller-gen: ## Download controller-gen locally if necessary.
	@if [ -f "$(CONTROLLER_GEN)" ] && [[ ! `"$(CONTROLLER_GEN)" --version` =~ "$(CONTROLLER_GEN_VERSION)" ]]; then\
		rm "$(CONTROLLER_GEN)";\
	fi
	$(call go-get-tool,$(CONTROLLER_GEN),sigs.k8s.io/controller-tools/cmd/controller-gen@$(CONTROLLER_GEN_VERSION))

KUSTOMIZE = $(shell pwd)/bin/kustomize
kustomize: ## Download kustomize locally if necessary.
	$(call go-get-tool,$(KUSTOMIZE),sigs.k8s.io/kustomize/kustomize/v5@v5.2.1)

YQ = $(shell pwd)/bin/yq
yq: ## Download YQ locally if necessary.
	$(call go-get-tool,$(YQ),github.com/mikefarah/yq/v4@v4.43.1)

# go-get-tool will 'go get' any package $2 and install it to $1.
PROJECT_DIR := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))
define go-get-tool
@[ -f $(1) ] || { \
set -e ;\
TMP_DIR=$$(mktemp -d) ;\
echo "Downloading $(2)" ;\
GOPATH="$$TMP_DIR" GOBIN="$$TMP_DIR/bin" go install $(2) ;\
cp -r "$$TMP_DIR/bin" "$(PROJECT_DIR)" ;\
chmod -R 777 "$$TMP_DIR" ; rm -rf "$$TMP_DIR" ;\
}
endef
