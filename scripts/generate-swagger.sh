#!/usr/bin/env bash

set -eu pipefail

echo "Reinitializing the working directory..."

SWAGGER_DIR=./swagger-proto

# prepare swagger generation
rm -rf $SWAGGER_DIR && mkdir -p "$SWAGGER_DIR/proto"
printf "version: v1\ndirectories:\n  - proto\n  - third_party" > "$SWAGGER_DIR/buf.work.yaml"
printf "version: v1\nname: buf.build/firmachain/firmachain\n" > "$SWAGGER_DIR/proto/buf.yaml"
cp ./proto/buf.gen.swagger.yaml "$SWAGGER_DIR/proto/buf.gen.swagger.yaml"

# copy existing proto files
cp -r ./proto/firmachain "$SWAGGER_DIR/proto"

# step into swagger folder
cd "$SWAGGER_DIR"/proto

echo "Copying protos and downloading dependencies..."

function get_go_dir() {
  go list -f '{{ .Dir }}' -m "$1"
}

cosmos_sdk_dir=$(get_go_dir "github.com/cosmos/cosmos-sdk")
wasmd=$(get_go_dir  "github.com/CosmWasm/wasmd")
pfm=$(get_go_dir "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8")
ibc_go_dir=$(get_go_dir "github.com/cosmos/ibc-go/v8")

# Change permissions termporarly in order to copy the needed files
chmod -R +w $cosmos_sdk_dir
chmod -R +w $wasmd
chmod -R +w $pfm
chmod -R +w $ibc_go_dir

third_party_dir="./third_party"
mkdir -p ${third_party_dir}

cp -r "$cosmos_sdk_dir"/proto/* ${third_party_dir}
cp -r "$wasmd"/proto/* ${third_party_dir}
cp -r "$pfm"/proto/* ${third_party_dir}
cp -r "$ibc_go_dir"/proto/* ${third_party_dir}

find ${third_party_dir} -maxdepth 1 -type f -delete

mv ${third_party_dir}/* ./

# download necessary proto files
mkdir -p "./gogoproto"
mkdir -p "./google/api"
mkdir -p "./cosmos/ics23/v1"
mkdir -p "./cosmos_proto"
curl -SSL -o "./gogoproto/gogo.proto" https://raw.githubusercontent.com/cosmos/gogoproto/main/gogoproto/gogo.proto
curl -sSL -o "./google/api/annotations.proto" https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto
curl -sSL -o "./google/api/http.proto" https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto
curl -sSL -o "./cosmos/ics23/v1/proofs.proto" https://raw.githubusercontent.com/cosmos/ics23/master/proto/cosmos/ics23/v1/proofs.proto
curl -sSL -o "./cosmos_proto/cosmos.proto" https://raw.githubusercontent.com/cosmos/cosmos-proto/refs/heads/main/proto/cosmos_proto/cosmos.proto

# create swagger files on an individual basis  w/ `buf build` and `buf generate` (needed for `swagger-combine`)
proto_dirs=$(find \
  "." \
  -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
echo "From proto dirs, finding query.proto and service.proto files..."
for dir in $proto_dirs; do
echo "checking proto files in $dir"
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ -n "$query_file" ]]; then
    echo "[+] $query_file"
    buf generate --template ./buf.gen.swagger.yaml "$query_file"
  fi
done

cd ..

# Build the base JSON structure.
base_json=$(jq -n '{
  swagger: "2.0",
  info: { title: "Firmachain - gRPC Gateway docs", description: "A REST interface for state queries, legacy transactions", version: "1.0.0" },
  consumes: ["application/json"],
  produces: ["application/json"],
  paths: {},
  definitions: {}
}')

echo "Deleting not necessary outputs..."
# Delete unnecessary files
rm -rf ./tmp-swagger-gen/cosmos/tx/v1beta1/service.swagger.json
rm -rf ./tmp-swagger-gen/cosmos/autocli/v1/query.swagger.json
rm -rf ./tmp-swagger-gen/cosmos/orm/query/v1alpha1/query.swagger.json
# Delete cosmos/nft path since firmachain uses its own module
rm -rf ./tmp-swagger-gen/cosmos/nft
# Delete not installed modules
rm -rf ./tmp-swagger-gen/cosmos/group

echo "Copying all *.swagger.json into a single folder..."
files=$(find ./tmp-swagger-gen -name '*.swagger.json' -print0 | xargs -0 | tr ' ' '\n' | sort)
mkdir -p ./tmp-swagger-gen/_all
counter=0
for f in $files; do
  echo "[+] $f"
  case "$f" in
    *firmachain*) cp "$f" ./tmp-swagger-gen/_all/01-firma-$counter.json ;;
    *cosmos*) cp "$f" ./tmp-swagger-gen/_all/02-cosmos-$counter.json ;;
    *cosmwasm*) cp "$f" ./tmp-swagger-gen/_all/03-cosmwasm-$counter.json ;;
    *ibc*) cp "$f" ./tmp-swagger-gen/_all/04-ibc-$counter.json ;;
    *) cp "$f" ./tmp-swagger-gen/_all/05-other-$counter.json ;;
  esac
  counter=$(expr $counter + 1)
done


# Save the base JSON to a temporary file.
temp_file=$(mktemp)
echo "$base_json" > "$temp_file"

echo "Generating a final json file..."
# Loop through all JSON files in the target directory and merge their "paths" and "definitions".
for file in ./tmp-swagger-gen/_all/*.json; do
  content=$(cat "$file")
  temp_file2=$(mktemp)
  jq --argjson new "$content" '
    .paths += ($new.paths // {}) |
    .definitions += ($new.definitions // {})' "$temp_file" > "$temp_file2"
  mv "$temp_file2" "$temp_file"
done

# Save the final merged JSON to FINAL.json.
jq . "$temp_file" > "./tmp-swagger-gen/_all/FINAL.json"
rm "$temp_file"

cd ..

# combine swagger files
echo "Generating swagger.yaml..."
OUTPUT_FILE=./client/docs/static/swagger.yaml
npx swagger-combine "${SWAGGER_DIR}/tmp-swagger-gen/_all/FINAL.json" -o "${SWAGGER_DIR}/tmp-swagger-gen/tmp_swagger.yaml" -f yaml --continueOnConflictingPaths true --includeDefinitions true && \
echo "swagger files combined" && \
npx swagger-merger --input "${SWAGGER_DIR}/tmp-swagger-gen/tmp_swagger.yaml" -o "$OUTPUT_FILE" && 
[[ -f "$OUTPUT_FILE" ]] && \
echo "swagger files merged" && \
{

  # Restore permissions
  chmod -R -w $cosmos_sdk_dir
  chmod -R -w $wasmd
  chmod -R -w $pfm
  chmod -R -w $ibc_go_dir

  # Cleanup.
  rm -rf ${SWAGGER_DIR}

  echo "Swagger generation complete. Output at $OUTPUT_FILE"

} || {
  
  # # Restore permissions
  chmod -R -w $cosmos_sdk_dir
  chmod -R -w $wasmd
  chmod -R -w $pfm
  chmod -R -w $ibc_go_dir

  echo "Error: Swagger generation failed."
  exit 1
}


