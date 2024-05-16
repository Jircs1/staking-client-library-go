#
# Tool Prerequisites Check
# This ensures that you have the necessary executables installed to run this makefile.
#

BUILD_PREREQUISITES = git go buf protoc-gen-go protoc-gen-go-grpc protoc-gen-csf protoc-gen-grpc-gateway protoc-gen-go-aip mockery

#
# Informational Targets
# These rarely change, and only serve to print out helpful details about this makefile.
#

.PHONY: usage
usage:
	@ echo "Usage: make [`cat Makefile | grep "^[A-z\%\-]*:" | awk '{print $$1}' | sed "s/://g" | sed "s/%/[1-3]/g" | xargs`]"

#
# Functional Targets
# These serve specific purposes for the build/release of this client.
#

.PHONY: install_prerequisites
install_prerequisites: # Install build prerequisites
	@ scripts/install-go-tools.sh
	@ scripts/install-go-validation-tools.sh

.PHONY: clean
clean:
	@ rm -rf gen/*
	@ rm -rf docs/openapi/*

.PHONY: build_deps
build_deps:
	@ printf $(foreach exec,$(BUILD_PREREQUISITES), \
        $(if $(shell which $(exec)),"", \
        $(error "No $(exec) in PATH. Prerequisites are: $(BUILD_PREREQUISITES)")))

.PHONY: buf_gen
buf_gen: 
	@ buf generate --template protos/buf.gen.orchestration.yaml buf.build/cdp/orchestration --path coinbase/staking/orchestration/v1
	@ buf generate --template protos/buf.gen.rewards.yaml buf.build/cdp/rewards --path coinbase/staking/rewards/v1

	# We don't need internal snippets and test files generated by gapic plugin.
	# w.r.t the internal/snippets folder, gapic provides an option to supply an option to omit generating them here:
	# https://github.com/googleapis/gapic-generator-go?tab=readme-ov-file#invocation 
	# but unfortunately, the option doesn't work and is likely to be removed, per this issue:
	# https://github.com/googleapis/gapic-generator-go/issues/1332
	@ rm -rf gen/client/coinbase/staking/orchestration/v1/internal
	@ rm -rf gen/client/coinbase/staking/rewards/v1/internal
	@ rm -rf gen/client/coinbase/staking/orchestration/v1/*test.go
	@ rm -rf gen/client/coinbase/staking/rewards/v1/*test.go

.PHONY: gen
gen: install_prerequisites build_deps clean buf_gen format

.PHONY: build
build:
	@ go build -o bin/cosmos_list_rewards examples/cosmos/list-rewards/main.go
	@ go build -o bin/ethereum/create_and_process_workflow examples/ethereum/create-and-process-workflow/main.go
	@ go build -o bin/ethereum/create_workflow examples/ethereum/create-workflow/main.go
	@ go build -o bin/ethereum/list_rewards examples/ethereum/list-rewards/partial-eth/main.go
	@ go build -o bin/ethereum/list_rewards examples/ethereum/list-rewards/validator/main.go
	@ go build -o bin/ethereum/list_stakes examples/ethereum/list-stakes/main.go
	@ go build -o bin/example examples/hello-world/main.go
	@ go build -o bin/list_workflows examples/list-workflows/main.go
	@ go build -o bin/solana/solana_list_rewards examples/solana/list-rewards/main.go
	@ go build -o bin/solana/solana_create_workflow examples/solana/create-workflow/main.go

.PHONY: lint
lint:
	@ echo "Linting app..."
	@ docker run --rm -v "${PWD}:/app" -w /app golangci/golangci-lint:v1.52.2 golangci-lint run --timeout 3m ./...
	@ echo "Done linting app"

.PHONY: lintfix
lintfix:
	@ echo "Linting app and fixing issues..."
	@ docker run --rm -v "${PWD}:/app" -w /app golangci/golangci-lint:v1.52.2 golangci-lint run --timeout 3m --fix ./...
	@ echo "Done linting app and fixing issues"

.PHONY: format
format:
	@ echo "Formatting app..."
	@ docker run --rm -v "${PWD}:/app" -w /app golang:1.20 bash -c "go install mvdan.cc/gofumpt@v0.6.0 && go install github.com/daixiang0/gci@v0.13.4 && gofumpt -l -w . && gci write --skip-generated ."
	@ echo "Done formatting app"
