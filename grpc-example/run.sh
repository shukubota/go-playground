protoc \
  --go_out=./protobuf/server --go_opt=paths=source_relative \
  --go-grpc_out=./protobuf/server --go-grpc_opt=paths=source_relative \
  protobuf/*.proto

protoc \
  --js_out=import_style=commonjs,binary:./protobuf/web \
  --grpc-web_out=import_style=typescript,mode=grpcwebtext:./protobuf/web \
  protobuf/*.proto