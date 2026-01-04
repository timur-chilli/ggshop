package bootstrap

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os"

	server "github.com/timur-chilli/ggshop/customer_side/internal/api/customer_side_service_api"
	customerCreateOrderConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/customer_create_order_consumer"
	customerGetOrderConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/customer_get_order_consumer"
	ggorderInfoConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/ggorder_info_consumer"
	ggorderInfoEditConsumer "github.com/timur-chilli/ggshop/customer_side/internal/consumer/ggorder_info_edit_consumer"

	"github.com/go-chi/chi/v5"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	httpSwagger "github.com/swaggo/http-swagger"
	customerSideAPI "github.com/timur-chilli/ggshop/customer_side/internal/pb/customer_side_api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func AppRun(api server.CustomerSideServiceAPI, ggorderInfoEditConsumer *ggorderInfoEditConsumer.AskGGOrderInfoEditConsumer,
	ggorderInfoConsumer *ggorderInfoConsumer.AskGGOrderInfoConsumer, createOrderConsumer *customerCreateOrderConsumer.CustomerCreateOrderConsumer, 
	getOrderConsumer *customerGetOrderConsumer.CustomerGetOrderConsumer) {
	go ggorderInfoEditConsumer.Consume(context.Background())
	go ggorderInfoConsumer.Consume(context.Background())
	go createOrderConsumer.Consume(context.Background())
	go getOrderConsumer.Consume(context.Background())

	go func() {
		if err := runGRPCServer(api); err != nil {
			panic(fmt.Errorf("failed to run gRPC server: %v", err))
		}
	}()

	if err := runGatewayServer(); err != nil {
		panic(fmt.Errorf("failed to run gateway server: %v", err))
	}
}

func runGRPCServer(api server.CustomerSideServiceAPI) error {
	lis, err := net.Listen("tcp", os.Getenv("customerGRPCPort"))
	if err != nil {
		return err
	}

	s := grpc.NewServer()
	customerSideAPI.RegisterCustomerSideServiceServer(s, &api)

	slog.Info(fmt.Sprintf("gRPC-server server listening on %v", os.Getenv("customerGRPCPort")))
	return s.Serve(lis)
}

func runGatewayServer() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	customerSwaggerPath := os.Getenv("customerSwaggerPath")
	if _, err := os.Stat(customerSwaggerPath); os.IsNotExist(err) {
		panic(fmt.Errorf("swagger file not found: %s", customerSwaggerPath))
	}

	r := chi.NewRouter()
	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, customerSwaggerPath)
	})

	r.Get("/docs/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger.json"),
	))

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	err := customerSideAPI.RegisterCustomerSideServiceHandlerFromEndpoint(ctx, mux, os.Getenv("customerGRPCPort"), opts)
	if err != nil {
		panic(err)
	}

	r.Mount("/", mux)

	slog.Info(fmt.Sprintf("gRPC-Gateway server listening on %v", os.Getenv("customerHTTPGatewayPort")))
	return http.ListenAndServe(os.Getenv("customerHTTPGatewayPort"), r)
}
