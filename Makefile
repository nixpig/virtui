ifneq (,$(wildcard .env))
	include .env
	export
endif

BINARY_NAME := virtui
PACKAGE_PATH := cmd/virtui.go
BUILD_OUTPUT_DIR := tmp/bin
CGO_ENABLED := 1

.PHONY: tidy
tidy: 
	go fmt ./...
	go mod tidy -v

.PHONY: audit
audit:
	go mod verify
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck@latest -checks=all,-ST1000,-U1000 ./...
	go run golang.org/x/vuln/cmd/govulncheck@latest ./...

.PHONY: test
test: 
	go run gotest.tools/gotestsum@latest ./...

.PHONY: coverage
coverage:
	go test -v -race -buildvcs -covermode atomic -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out

.PHONY: coveralls
coveralls:
	go run github.com/mattn/goveralls@latest -coverprofile=coverage.out -service=github

.PHONY: build
build:
	go build -o ${BUILD_OUTPUT_DIR}/${BINARY_NAME} ${PACKAGE_PATH}

.PHONY: run
run:
	go run ${PACKAGE_PATH}

.PHONY: clean
clean:
	rm -rf tmp *.out
	go clean

.PHONY: vhs
vhs:
	vhs demo.tape

.PHONY: env
env: 
	# Environment variables
	BINARY_NAME=${BINARY_NAME}
	PACKAGE_PATH=${PACKAGE_PATH}
	BUILD_OUTPUT_DIR=${BUILD_OUTPUT_DIR}
	CGO_ENABLED=${CGO_ENABLED}

