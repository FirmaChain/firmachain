#!/usr/bin/env bash

set -euo pipefail

ROOT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
OUTPUT_DIR="${1:-${ROOT_DIR}/build/v0.5.2}"
PRIVATE_CACHE_DIR="${ROOT_DIR}/.modcache"
COSMWASM_DOWNLOAD_CACHE="$(GOTOOLCHAIN=go1.23.4 go env GOMODCACHE)/cache/download/github.com/!cosm!wasm"
TEMP_OUTPUT_DIR="$(mktemp -d "${TMPDIR:-/tmp}/firmachain-v0.5.2.XXXXXX")"

cleanup() {
  rm -rf "${PRIVATE_CACHE_DIR}"
  rm -rf "${TEMP_OUTPUT_DIR}"
}

trap cleanup EXIT

cd "${ROOT_DIR}"

if [[ "$(GOTOOLCHAIN=go1.23.4 go env GOVERSION)" != "go1.23.4" ]]; then
  echo "Go 1.23.4 is required" >&2
  exit 1
fi

export GOTOOLCHAIN="go1.23.4"
export GOPRIVATE="github.com/CosmWasm/priv_wasmd_sec,github.com/CosmWasm/priv_wasmvm_sec"

go mod download

rm -rf "${PRIVATE_CACHE_DIR}"
mkdir -p "${PRIVATE_CACHE_DIR}/github.com/!cosm!wasm"
cp -R "${COSMWASM_DOWNLOAD_CACHE}/priv_wasmd_sec" \
  "${COSMWASM_DOWNLOAD_CACHE}/priv_wasmvm_sec" \
  "${PRIVATE_CACHE_DIR}/github.com/!cosm!wasm/"

mkdir -p "${OUTPUT_DIR}"

BUILD_COMMIT="$(git rev-parse HEAD)"
if [[ -n "$(git status --porcelain --untracked-files=normal)" ]]; then
  BUILD_COMMIT="${BUILD_COMMIT}-dirty"
fi

docker build \
  --platform linux/amd64 \
  --target binary \
  --build-arg VERSION="${VERSION:-v0.5.2}" \
  --build-arg COMMIT="${BUILD_COMMIT}" \
  --output "type=local,dest=${TEMP_OUTPUT_DIR}" \
  .

test -f "${TEMP_OUTPUT_DIR}/firmachaind"
install -m 0755 "${TEMP_OUTPUT_DIR}/firmachaind" "${OUTPUT_DIR}/.firmachaind.tmp"
mv "${OUTPUT_DIR}/.firmachaind.tmp" "${OUTPUT_DIR}/firmachaind"

file "${OUTPUT_DIR}/firmachaind"

if command -v sha256sum >/dev/null 2>&1; then
  sha256sum "${OUTPUT_DIR}/firmachaind"
else
  shasum -a 256 "${OUTPUT_DIR}/firmachaind"
fi
