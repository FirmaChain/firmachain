PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
BRANCH := $(shell git branch --show-current)

include Makefile.ledger

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

ldflags = -X github.com/cosmos/cosmos-sdk/version.Name=FirmaChain \
	-X github.com/cosmos/cosmos-sdk/version.AppName=firmachaind \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)" \

ifeq ($(LEDGER_ENABLED),true)
	BUILD_FLAGS := -ldflags '$(ldflags)' -tags $(build_tags)
else
	BUILD_FLAGS := -ldflags '$(ldflags)' $(build_tags)
endif

DOCKER := $(shell which docker)

all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/firmachaind

docker-img-from-current-branch:
	docker build \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
		--build-arg GIT_BRANCH=$(BRANCH) \
		-t firmachain .

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)

###############################################################################
###                                  Proto                                  ###
###############################################################################

# Variables for the image and version
protoVer=0.14.0
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protocgen.sh