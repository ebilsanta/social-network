protoc --proto_path=services/proto \
       --go_out=services/proto/generated \
       --go_opt=paths=source_relative \
       --go-grpc_out=services/proto/generated \
       --go-grpc_opt=paths=source_relative \
       services/proto/*.proto