Code to generate grpc files
```
protoc --proto_path=proto \
       --go_out=proto/generated \
       --go_opt=paths=source_relative \
       --go-grpc_out=proto/generated \
       --go-grpc_opt=paths=source_relative \
       proto/*.proto
```