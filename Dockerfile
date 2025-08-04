# syntax=docker/dockerfile:1

ARG GO_VERSION="1.23"
ARG RUNNER_IMAGE="alpine:3.20"

# Use a minimal base image (Alpine Linux)
FROM golang:${GO_VERSION}-alpine3.20 AS go-builder

WORKDIR /app

RUN apk add --no-cache ca-certificates build-base git linux-headers binutils-gold

# Download dependencies and CosmWasm libwasmvm if found.
ADD go.mod go.sum ./
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    go mod download

# Cosmwasm - Download correct libwasmvm version
RUN ARCH=$(uname -m) && WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm/v2 | sed 's/.* //') && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/libwasmvm_muslc.$ARCH.a \
    -O /lib/libwasmvm_muslc.$ARCH.a && \
    # verify checksum
    wget https://github.com/CosmWasm/wasmvm/releases/download/$WASMVM_VERSION/checksums.txt -O /tmp/checksums.txt && \
    sha256sum /lib/libwasmvm_muslc.$ARCH.a | grep $(cat /tmp/checksums.txt | grep libwasmvm_muslc.$ARCH | cut -d ' ' -f 1) 
  

# Copy the repository
COPY . ./

ARG GIT_VERSION 
ARG GIT_COMMIT
ARG GIT_BRANCH

RUN git checkout ${GIT_BRANCH}

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/root/go/pkg/mod \
    GOWORK=off go build \
    -mod=readonly \
    -tags "netgo,ledger,muslc" \
    -ldflags \
    "-X github.com/cosmos/cosmos-sdk/version.Name="firmachain" \
    -X github.com/cosmos/cosmos-sdk/version.AppName="firmachaind" \
    -X github.com/cosmos/cosmos-sdk/version.Version=${GIT_VERSION}@${GIT_BRANCH} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${GIT_COMMIT} \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o ./bin/firmachaind \
    ./cmd/firmachaind

# --------------------------------------------------------
FROM ${RUNNER_IMAGE}

COPY --from=go-builder /app/bin/firmachaind /usr/bin/firmachaind

# rest server, tendermint p2p, tendermint rpc
EXPOSE 1317 26656 26657

ENV PATH="/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin:$PATH"
CMD ["/usr/bin/firmachaind", "version"]
