default: build

build:
	go install

fmtcheck:

test:
	go test ./...

testacc: fmtcheck
	TF_SCHEMA_PANIC_ON_ERROR=1 \
	TF_ACC=1 go test ./... -v $(TESTARGS) -timeout 10m

testinfra-testacc:
	CISCOASA_SSLNOVERIFY=true \
	CISCOASA_OBJECT_PREFIX=acc \
	CISCOASA_INTERFACE_NAME=inside \
	CISCOASA_USERNAME="$$(cd testinfra; terraform output asav_username)" \
	CISCOASA_PASSWORD="$$(cd testinfra; terraform output asav_password)" \
	CISCOASA_API_URL="https://$$(cd testinfra; terraform output asav_public_ip)" \
	TF_SCHEMA_PANIC_ON_ERROR=1 \
	TF_ACC=1 go test ./... -count 1 -v -cover $(TESTARGS) -timeout 10m

.PHONY: build fmtcheck test testacc testinfra-testacc
