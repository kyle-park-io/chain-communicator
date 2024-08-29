SCRIPT_DIR=$(dirname "$(realpath "$0")")

rm -rf $SCRIPT_DIR/proto-go $SCRIPT_DIR/proto-ts
mkdir -p $SCRIPT_DIR/proto-go $SCRIPT_DIR/proto-ts

# go
protoc trx.proto -I=./proto --go_out=proto-go
# test protobuf
protoc test_trx.proto -I=./proto --go_out=proto-go

# js, ts
protoc trx.proto -I=./proto \
  --js_out=import_style=commonjs,binary:./proto-ts --plugin=protoc-gen-js=./node_modules/.bin/protoc-gen-js
protoc trx.proto -I=./proto \
  --ts_out=./proto-ts --plugin=protoc-gen-ts=./node_modules/.bin/protoc-gen-ts

echo "protoc finished"
