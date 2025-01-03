#!/usr/bin/env bash

#== Requirements ==
#
## make sure your `go env GOPATH` is in the `$PATH`
## Install:
## + latest buf (v1.0.0-rc11 or later)
## + protobuf v3

set -eo pipefail
echo "Removing old proto go files..."
find . -type f -name "*.pb.go" -exec rm -f {} +
find . -type f -name "*.pb.gw.go" -exec rm -f {} +
echo "Generating gogo proto code..."
cd proto
buf mod update
cd ..
buf generate --template ./proto/buf.gen.gogo.yaml

# move proto files to the right places
if [ -d "./github.com" ]; then
    echo "Copying generated files..."
    cp -r ./github.com/firmachain/firmachain/v05/x/* x/
    rm -rf ./github.com
else
    echo "No files in ./github.com, skipping..."
fi

echo "Running go mod tidy..."
go mod tidy

echo "Done."

# ./scripts/protocgen2.sh