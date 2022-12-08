PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell echo $(shell git describe --tags --always) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

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

all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/firmachaind

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
