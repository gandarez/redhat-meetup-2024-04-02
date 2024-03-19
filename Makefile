# globals
BINARY_NAME?=video-game-api
BUILD_DIR?="./build"

# basic Go commands
GOCMD=go
GOBUILD=$(GOCMD) build

# linting
define get_latest_lint_release
	curl -s "https://api.github.com/repos/golangci/golangci-lint/releases/latest" | grep '"tag_name":' | sed -E 's/.*"([^"]+)".*/\1/'
endef
LATEST_LINT_VERSION=$(shell $(call get_latest_lint_release))
INSTALLED_LINT_VERSION=$(shell golangci-lint --version 2>/dev/null | awk '{print "v"$$4}')

# get GOPATH according to OS
ifeq ($(OS),Windows_NT) # is Windows_NT on XP, 2000, 7, Vista, 10...
    GOPATH=$(go env GOPATH)
else
    GOPATH=$(shell go env GOPATH)
endif

.PHONY: install-linter
install-linter:
ifneq "$(INSTALLED_LINT_VERSION)" "$(LATEST_LINT_VERSION)"
	@echo "new golangci-lint version found:" $(LATEST_LINT_VERSION)
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(GOPATH)/bin latest
endif

# targets
build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 $(MAKE) build

build-darwin-amd64:
	GOOS=darwin GOARCH=amd64 $(MAKE) build

build-linux-amd64:
	GOOS=linux GOARCH=amd64 $(MAKE) build

build-windows-amd64:
	GOOS=windows GOARCH=amd64 $(MAKE) build-windows

.PHONY: build
build:
	CGO_ENABLED="0" GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -v \
		-o ${BUILD_DIR}/$(BINARY_NAME)-$(GOOS)-$(GOARCH) ./cmd/$(BINARY_NAME)/main.go

.PHONY: build-windows
build-windows:
	CGO_ENABLED="0" GOOS=$(GOOS) GOARCH=$(GOARCH) $(GOBUILD) -v \
		-o ${BUILD_DIR}/$(BINARY_NAME)-$(GOOS)-$(GOARCH).exe ./cmd/$(BINARY_NAME)/main.go


.PHONY: lint
lint: install-linter
	golangci-lint run ./...

.PHONY: test
test:
	go test -count=1 -race -covermode=atomic -coverprofile=coverage.out ./...

.PHONY: gen.go
gen.go:
	go generate ./...
