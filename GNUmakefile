TEST?=./...
PKG_NAME=ciscoasa
WEBSITE_REPO=github.com/hashicorp/terraform-website
HOSTNAME=registry.terraform.io
NAMESPACE=CiscoDevNet
BINARY=terraform-provider-${PKG_NAME}
VERSION=0.2
OS_ARCH=$(shell go env GOOS)_$(shell go env GOARCH)

default: build

build:
	go build -o ${BINARY}_${VERSION}_${OS_ARCH}

install: build
	mkdir -p ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${PKG_NAME}/${VERSION}/${OS_ARCH}
	mv ${BINARY}_${VERSION}_${OS_ARCH} ~/.terraform.d/plugins/${HOSTNAME}/${NAMESPACE}/${PKG_NAME}/${VERSION}/${OS_ARCH}

fmtcheck:
	@sh -c "'$(CURDIR)/scripts/gofmtcheck.sh'"

test: fmtcheck
	go test $(TEST) -timeout=30s -parallel=4


test-compile:
	@if [ "$(TEST)" = "./..." ]; then \
		echo "ERROR: Set TEST to a specific package. For example,"; \
		echo "  make test-compile TEST=./$(PKG_NAME)"; \
		exit 1; \
	fi
	go test -c $(TEST) $(TESTARGS)

testacc:
	TF_SCHEMA_PANIC_ON_ERROR=1 \
	TF_ACC=1 go test $(TEST) -v $(TESTARGS) -timeout 10m

testinfra-testacc:
	CISCOASA_SSLNOVERIFY=true \
	CISCOASA_OBJECT_PREFIX=acc \
	CISCOASA_INTERFACE_NAME=inside \
	CISCOASA_INTERFACE_HW_ID_BASE=TenGigabitEthernet0 \
	CISCOASA_USERNAME="$$(cd testinfra; terraform output asav_username)" \
	CISCOASA_PASSWORD="$$(cd testinfra; terraform output asav_password)" \
	CISCOASA_API_URL="https://$$(cd testinfra; terraform output asav_public_ip)" \
	TF_SCHEMA_PANIC_ON_ERROR=1 \
	TF_ACC=1 go test ./... -count 1 -v -cover $(TESTARGS) -timeout 10m

website:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)

website-test:
ifeq (,$(wildcard $(GOPATH)/src/$(WEBSITE_REPO)))
	echo "$(WEBSITE_REPO) not found in your GOPATH (necessary for layouts and assets), get-ting..."
	git clone https://$(WEBSITE_REPO) $(GOPATH)/src/$(WEBSITE_REPO)
endif
	@$(MAKE) -C $(GOPATH)/src/$(WEBSITE_REPO) website-provider-test PROVIDER_PATH=$(shell pwd) PROVIDER_NAME=$(PKG_NAME)


.PHONY: build fmtcheck test test-compile testacc testinfra-testacc website website-test
