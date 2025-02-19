#!/usr/bin/env bash

SWAGGER_DIR=./swagger-proto

if [[ ! -d $SWAGGER_DIR ]]; then 
  echo "Error: $SWAGGER_DIR does not exist. Please run this script from the root of the repository."
  exit 1
fi

set -eo pipefail

# prepare swagger generation
mkdir -p "$SWAGGER_DIR/proto"
printf "version: v1\ndirectories:\n  - proto\n  - third_party" > "$SWAGGER_DIR/buf.work.yaml"
printf "version: v1\nname: buf.build/osmosis/osmosis\n" > "$SWAGGER_DIR/proto/buf.yaml"
cp ./proto/buf.gen.swagger.yaml "$SWAGGER_DIR/proto/buf.gen.swagger.yaml"

# copy existing proto files
cp -r ./proto/firmachain "$SWAGGER_DIR/proto"

# create temporary folder to store intermediate results from `buf generate`
rm -rf ./tmp-swagger-gen &&  mkdir -p ./tmp-swagger-gen

# step into swagger folder
cd "$SWAGGER_DIR"

# create swagger files on an individual basis  w/ `buf build` and `buf generate` (needed for `swagger-combine`)
proto_dirs=$(find ./proto ./third_party -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
  # generate swagger files (filter query files)
  query_file=$(find "${dir}" -maxdepth 1 \( -name 'query.proto' -o -name 'service.proto' \))
  if [[ -n "$query_file" ]]; then
    buf generate --template proto/buf.gen.swagger.yaml "$query_file"
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

# Fix circular definitions in cosmos by removing them
jq 'del(.definitions["cosmos.tx.v1beta1.ModeInfo.Multi"].properties.mode_infos.items["$ref"])' ./tmp-swagger-gen/cosmos/tx/v1beta1/service.swagger.json > ./tmp-swagger-gen/cosmos/tx/v1beta1/fixed_service.swagger.json
jq 'del(.definitions["cosmos.autocli.v1.ServiceCommandDescriptor"].properties.sub_commands)' ./tmp-swagger-gen/cosmos/autocli/v1/query.swagger.json > ./tmp-swagger-gen/cosmos/autocli/v1/fixed_query.swagger.json

# Delete unnecessary files
rm -rf ./tmp-swagger-gen/cosmos/tx/v1beta1/service.swagger.json
rm -rf ./tmp-swagger-gen/cosmos/autocli/v1/query.swagger.json

# Delete cosmos/nft path since firmachain uses its own module
rm -rf ./tmp-swagger-gen/cosmos/nft

# Convert all *.swagger.json files into a single folder _all
files=$(find ./tmp-swagger-gen -name '*.swagger.json' -print0 | xargs -0)
mkdir -p ./tmp-swagger-gen/_all
counter=0
for f in $files; do
  echo "[+] $f"
  cp "$f" ./tmp-swagger-gen/_all/swagger-json-$counter.json
  counter=$(expr $counter + 1)
done


# Save the base JSON to a temporary file.
temp_file=$(mktemp)
echo "$base_json" > "$temp_file"

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

# combine swagger files
OUTPUT_FILE=./client/docs/static/swagger.yaml
npx swagger-combine "./tmp-swagger-gen/_all/FINAL.json" -o "./tmp-swagger-gen/tmp_swagger.yaml" -f yaml --continueOnConflictingPaths true --includeDefinitions true && \
echo "swagger files combined" && \
npx swagger-merger --input "./tmp-swagger-gen/tmp_swagger.yaml" -o "$OUTPUT_FILE" && 
[[ -f "$OUTPUT_FILE" ]] && \
echo "swagger files merged" && \
{
  # Cleanup.
  rm -rf tmp-swagger-gen
  echo "Swagger generation complete. Output at $OUTPUT_FILE"

} || {
  
  echo "Error: Swagger generation failed."
  exit 1
}


