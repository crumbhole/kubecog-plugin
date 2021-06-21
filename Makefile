.PHONY:	all clean code-vet code-fmt test

DEPS := $(shell find . -type f -name "*.go" -printf "%p ")

unexport BW_SESSION
unexport VAULT_ADDR
unexport CRUMBLECOG_PATH

all: code-vet code-fmt test build/crumblecog-plugin

clean:
	$(RM) -rf build

test:
	go test ./...

build/crumblecog-plugin: $(DEPS)
	mkdir -p build
	go build -o build ./...

code-vet: $(DEPS)
## Run go vet for this project. More info: https://golang.org/cmd/vet/
	@echo go vet
	go vet $$(go list ./... )

code-fmt: $(DEPS)
## Run go fmt for this project
	@echo go fmt
	go fmt $$(go list ./... )
