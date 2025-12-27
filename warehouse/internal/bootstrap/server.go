package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	server "github.com/timur-chilli/ggshop/warehouse/internal/api/warehouse_service_api"
	ggorderInfoConsumer "github.com/timur-chilli/ggshop/warehouse/internal/consumer/provide_ggorder_info_consumer"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	warehouseAPI "github.com/timur-chilli/ggshop/warehouse/internal/pb/warehouse_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppRun(api server.WarehouseServiceAPI, provideGGOrderInfoConsumer *ggorderInfoConsumer.ProvideGGOrderInfoConsumer) {
	go provideGGOrderInfoConsumer.Consume(context.Background())
	go func() {
		if err := runGRPCServer(api); err != nil {
			panic(fmt.Errorf("failed to run gRPC server: %v", err))
		}
	}()

	if err := runGatewayServer(); err != nil {
		panic(fmt.Errorf("failed to run gateway server: %v", err))
	}
}

func runGRPCServer(api server.WarehouseServiceAPI) error {
	lis, err := net.Listen("tcp", os.Getenv("warehouseGRPCPort"))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	warehouseAPI.RegisterWarehouseServiceServer(s, &api)

	slog.Info(fmt.Sprintf("gRPC-Gateway server listening on %v", os.Getenv("warehouseGRPCPort")))
	return s.Serve(lis)
}

func runGatewayServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	warehouseSwaggerPath := os.Getenv("warehouseSwaggerPath")
	if _, err := os.Stat(warehouseSwaggerPath); os.IsNotExist(err) {
		panic(fmt.Errorf("swagger file not found: %s", warehouseSwaggerPath))
	}

	r := chi.NewRouter()
	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, warehouseSwaggerPath)
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := warehouseAPI.RegisterWarehouseServiceHandlerFromEndpoint(ctx, mux, os.Getenv("warehouseGRPCPort"), opts)
	if err != nil {
		panic(err)
	}

	r.Mount("/", mux)

	slog.Info(fmt.Sprintf("gRPC-Gateway server listening on %v", os.Getenv("warehouseHTTPGatewayPort")))
	return http.ListenAndServe(os.Getenv("warehouseHTTPGatewayPort"), r)
}
