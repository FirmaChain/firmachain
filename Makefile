PACKAGES=$(shell go list ./... | grep -v '/simulation')
VERSION := $(shell git rev-parse HEAD)
COMMIT := $(shell git log -1 --format='%H')
ldflags = -X github.com/firmachain/FirmaChain/version.Name=FirmaChain \
	-X github.com/firmachain/FirmaChain/version.ServerName=firma \
	-X github.com/firmachain/FirmaChain/version.ClientName=firma-cli \
	-X github.com/firmachain/FirmaChain/version.Version=$(VERSION) \
	-X github.com/firmachain/FirmaChain/version.Commit=$(COMMIT)


include Makefile.ledger

BUILD_FLAGS := -ldflags '$(ldflags)' -tags$(build_tags)

all: install

install: go.sum
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/firma
		go install -mod=readonly $(BUILD_FLAGS) ./cmd/firma-cli

go.sum: go.mod
		@echo "--> Ensure dependencies have not been modified"
		GO111MODULE=on go mod verify

test:
	@go test -mod=readonly $(PACKAGES)
