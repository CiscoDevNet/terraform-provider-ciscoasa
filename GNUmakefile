TEST?=./...
PKG_NAME=ciscoasa
WEBSITE_REPO=github.com/hashicorp/terraform-website

default: build

build:
	go install

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
