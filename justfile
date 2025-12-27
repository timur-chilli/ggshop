set dotenv-load := true

customer_side_api_root := "./customer_side/api"
customer_side_pb_out := "./customer_side/internal/pb"

warehouse_api_root := "./warehouse/api"
warehouse_pb_out := "./warehouse/internal/pb"

build-customer-api:
    # set PATH="$PATH:$(go env GOPATH)/bin"
    # set PATH="$PATH:$(go env GOPATH)"

    protoc -I {{customer_side_api_root}} \
    -I {{customer_side_api_root}}/google/api \
    --go_out={{customer_side_pb_out}} --go_opt=paths=source_relative \
    --go-grpc_out={{customer_side_pb_out}} --go-grpc_opt=paths=source_relative \
    {{customer_side_api_root}}/customer_side_api/customer_side.proto {{customer_side_api_root}}/models/ggorder_model.proto

    # Генерация gRPC-Gateway
    protoc -I {{customer_side_api_root}} \
    -I {{customer_side_api_root}}/google/api \
    --grpc-gateway_out={{customer_side_pb_out}} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt logtostderr=true \
    {{customer_side_api_root}}/customer_side_api/customer_side.proto

    # Генерация OpenAPI
    protoc -I {{customer_side_api_root}} \
    -I {{customer_side_api_root}}/google/api \
    --openapiv2_out={{customer_side_pb_out}}/swagger \
    --openapiv2_opt logtostderr=true \
    {{customer_side_api_root}}/customer_side_api/customer_side.proto

build-warehouse-api:
    protoc -I {{warehouse_api_root}} \
    -I {{warehouse_api_root}}/google/api \
    --go_out={{warehouse_pb_out}} --go_opt=paths=source_relative \
    --go-grpc_out={{warehouse_pb_out}} --go-grpc_opt=paths=source_relative \
    {{warehouse_api_root}}/warehouse_api/warehouse.proto {{warehouse_api_root}}/models/ggorder_model.proto

    # Генерация gRPC-Gateway
    protoc -I {{warehouse_api_root}} \
    -I {{warehouse_api_root}}/google/api \
    --grpc-gateway_out={{warehouse_pb_out}} \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt logtostderr=true \
    {{warehouse_api_root}}/warehouse_api/warehouse.proto

    # Генерация OpenAPI
    protoc -I {{warehouse_api_root}} \
    -I {{warehouse_api_root}}/google/api \
    --openapiv2_out={{warehouse_pb_out}}/swagger \
    --openapiv2_opt logtostderr=true \
    {{warehouse_api_root}}/warehouse_api/warehouse.proto

run-customer:
    go run ./customer_side/cmd/app/main.go

run-warehouse:
    go run ./warehouse/cmd/app/main.go

up:
    docker-compose up -d

down: 
    docker-compose down

logs:
    docker-compose logs -f

mockery:
    mockery init 
