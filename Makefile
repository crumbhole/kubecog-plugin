.PHONY:	all clean code-vet code-fmt test

DEPS := $(shell find . -type f -name "*.go" -printf "%p ")

unexport BW_SESSION
unexport VAULT_ADDR
unexport KUBECOG_PATH

all: code-vet code-fmt test build/kubecog-plugin

clean:
	$(RM) -rf build

get: $(DEPS)
	go get ./...
	go get github.com/otiai10/copy

test: get
	go test ./...

test_verbose: get
	go test -v ./...

build/kubecog-plugin: $(DEPS)
	mkdir -p build
	go build -o build ./...

code-vet: get $(DEPS)
## Run go vet for this project. More info: https://golang.org/cmd/vet/
	@echo go vet
	go vet $$(go list ./... )

code-fmt: get $(DEPS)
## Run go fmt for this project
	@echo go fmt
	go fmt $$(go list ./... )

lint: $(DEPS)
## Run golint for this project
	@echo golint
	golint $$(go list ./... )
