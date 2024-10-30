Code to generate grpc files
```
protoc --proto_path=api/proto \
       --go_out=api/proto/generated \
       --go_opt=paths=source_relative \
       --go-grpc_out=api/proto/generated \
       --go-grpc_opt=paths=source_relative \
       api/proto/*.proto
```