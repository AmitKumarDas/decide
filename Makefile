# list only this project's namespaced directories
PACKAGES = $(shell go list ./... | grep -v '/vendor/')

.PHONY: all
all: format metalint test ginkgo

# `make setup` needs to be run in a completely new environment
# In case of go related issues, run below commands & verify:
# go version    # ensure go1.9.1 or above
# go env        # ensure if GOPATH is set
# echo $PATH    # ensure if $GOPATH/bin is set
.PHONY: setup
setup:
	@echo "------------------"
	@echo "--> Running setup"
	@echo "------------------"
	@go get -u -v github.com/golang/lint/golint
	@go get -u -v golang.org/x/tools/cmd/goimports
	@go get -u -v github.com/golang/dep/cmd/dep
	@go get -u -v github.com/alecthomas/gometalinter
	@go get -u -v github.com/onsi/ginkgo/ginkgo
	@go get -u -v github.com/onsi/gomega/...
	@gometalinter --install

.PHONY: format
format:
	@echo "------------------"
	@echo "--> Running go fmt"
	@echo "------------------"
	@go fmt $(PACKAGES)

.PHONY: lint
lint:
	@echo "------------------"
	@echo "--> Running golint"
	@echo "------------------"
	@golint $(PACKAGES)
	@echo "------------------"
	@echo "--> Running go vet"
	@echo "------------------"
	@go vet $(PACKAGES)

.PHONY: metalint
metalint:
	@echo "------------------"
	@echo "--> Running metalinter"
	@echo "------------------"
	@gometalinter $(PACKAGES)

.PHONY: test
test:
	@echo "------------------"
	@echo "--> Running test"
	@echo "------------------"
	@go test $(PACKAGES)

.PHONY: ginkgo
ginkgo:
	@echo "------------------"
	@echo "--> Running ginkgo"
	@echo "------------------"
	@cd integration && ginkgo
