GOFMT ?= gofmt -s
GOFMT_FILES?=$$(find . -name '*.go' |grep -v vendor)
VERSION = 1.1.2
CURRENT_USER = $$(id -u ${USER})
LOG_PATH = /mnt/d/trace.txt

vet:
	@echo "go vet ."
	@go vet $$(go list ./... | grep -v vendor/) ; if [ $$? -eq 1 ]; then \
		echo ""; \
		echo "Vet found suspicious constructs. Please check the reported constructs"; \
		echo "and fix them if necessary before submitting the code for review."; \
		exit 1; \
	fi

.PHONY: fmt-check
fmt-check:
	@diff=$$($(GOFMT) -d $(GOFMT_FILES)); \
	if [ -n "$$diff" ]; then \
		echo "Please run 'make fmt' and commit the result:"; \
		echo "$${diff}"; \
		exit 1; \
	fi;

fmt:
	@gofmt -w -s .

gen-api:
	@rm -fr ./internal/proxy/api/*
	@docker run -u ${CURRENT_USER}:${CURRENT_USER} --rm -v "${PWD}/internal/proxy/api:/local" openapitools/openapi-generator-cli generate \
		--additional-properties packageName=api \
		--additional-properties structPrefix=true \
		-i https://gitea.com/swagger.v1.json \
		-g go \
		-o /local
	@make fmt

build: 
	@go build -ldflags "-s -w" -o terraform-provider-gitea_${VERSION} -tags release ./cmd

install: build 
	@echo installing to 
	@echo ~/.terraform.d/plugins/terraform.local/local/gitea/${VERSION}/linux_amd64/terraform-provider-gitea_${VERSION}
	@mkdir -p ~/.terraform.d/plugins/terraform.local/local/gitea/${VERSION}/linux_amd64
	@mv ./terraform-provider-gitea_${VERSION} ~/.terraform.d/plugins/terraform.local/local/gitea/${VERSION}/linux_amd64/terraform-provider-gitea_${VERSION}

prepare-test: FORCE
	@make build
	@make install
	@make docs
	@rm -f ./test/.terraform.lock.hcl
	@rm -fR ./test/.terraform
	@terraform -chdir=./test init

docs: FORCE
	go generate ./...

FORCE: