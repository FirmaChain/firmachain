# Go Version
ARG GO_VERSION="1.23.4"
ARG GO_IMAGE_DIGEST="sha256:6a84ccdb73e005d0ee7bfff6066f230612ca9dff3e88e31bfc752523c3a271f8"
# WASMVM Version
ARG WASMVM_VERSION="v2.2.9-rc.3"
ARG WASMVM_SHA256="c2e4018d532138fad3113a140ccb6c73cfa41b5c9889633b08299e59568616b9"

FROM golang:${GO_VERSION}-alpine3.20@${GO_IMAGE_DIGEST} AS builder

WORKDIR /src

RUN apk add --no-cache binutils-gold build-base ca-certificates file git linux-headers xz

COPY go.mod go.sum ./
COPY .modcache/ /modcache/
ENV GOPROXY="file:///modcache,https://proxy.golang.org,direct"
RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    test "$(go env GOVERSION)" = "go1.23.4" \
    && go mod download

ARG TARGETARCH
ARG WASMVM_VERSION
ARG WASMVM_SHA256

RUN --mount=type=cache,target=/go/pkg/mod \
    test "${TARGETARCH}" = "amd64" \
    && test "$(go list -mod=readonly -m -f '{{.Replace.Path}} {{.Replace.Version}}' github.com/CosmWasm/wasmd)" = "github.com/CosmWasm/priv_wasmd_sec v0.54.10-rc.3" \
    && test "$(go list -mod=readonly -m -f '{{.Replace.Path}} {{.Replace.Version}}' github.com/CosmWasm/wasmvm/v2)" = "github.com/CosmWasm/priv_wasmvm_sec/v2 ${WASMVM_VERSION}" \
    && unxz -c "$(go list -mod=readonly -m -f '{{.Dir}}' github.com/CosmWasm/wasmvm/v2)/internal/api/libwasmvm_muslc.x86_64.a.xz" \
      > /lib/libwasmvm_muslc.x86_64.a \
    && echo "${WASMVM_SHA256}  /lib/libwasmvm_muslc.x86_64.a" | sha256sum -c -

COPY . .

ARG VERSION
ARG COMMIT

RUN --mount=type=cache,target=/root/.cache/go-build \
    --mount=type=cache,target=/go/pkg/mod \
    test -n "${VERSION}" \
    && test -n "${COMMIT}" \
    && CGO_ENABLED=1 GOOS=linux GOARCH=amd64 GOWORK=off go build \
    -mod=readonly \
    -tags "netgo,muslc" \
    -ldflags "-X github.com/cosmos/cosmos-sdk/version.Name=FirmaChain \
    -X github.com/cosmos/cosmos-sdk/version.AppName=firmachaind \
    -X github.com/cosmos/cosmos-sdk/version.Version=${VERSION} \
    -X github.com/cosmos/cosmos-sdk/version.Commit=${COMMIT} \
    -X github.com/cosmos/cosmos-sdk/version.BuildTags=netgo,muslc \
    -w -s -linkmode=external -extldflags '-Wl,-z,muldefs -static'" \
    -trimpath \
    -o /out/firmachaind \
    ./cmd/firmachaind \
    && file /out/firmachaind | grep -Eq "ELF 64-bit.*x86-64.*statically linked"

FROM scratch AS binary

COPY --from=builder /out/firmachaind /firmachaind

FROM alpine:3.20 AS runner

RUN apk add --no-cache ca-certificates curl jq \
    && addgroup -S -g 10001 firmachain \
    && adduser -S -D -H -u 10001 -G firmachain firmachain \
    && install -d -o firmachain -g firmachain /var/lib/firmachain

COPY --from=builder /out/firmachaind /usr/local/bin/firmachaind

ENV HOME="/var/lib/firmachain"

WORKDIR /var/lib/firmachain
USER 10001:10001

# REST, gRPC, CometBFT P2P, CometBFT RPC ports
EXPOSE 1317 9090 26656 26657

ENTRYPOINT ["/usr/local/bin/firmachaind"]
CMD ["start", "--home", "/var/lib/firmachain"]
