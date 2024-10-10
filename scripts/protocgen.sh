#!/usr/bin/env bash

#== Requirements ==
#
## make sure your `go env GOPATH` is in the `$PATH`
## Install:
## + latest buf (v1.0.0-rc11 or later)
## + protobuf v3

set -eo pipefail

echo "Generating gogo proto code"
cd proto
buf mod update
cd ..
buf generate

# move proto files to the right places
cp -r ./github.com/firmachain/firmachain/x/* x/
rm -rf ./github.com

go mod tidy -compat=1.17

# ./scripts/protocgen2.sh