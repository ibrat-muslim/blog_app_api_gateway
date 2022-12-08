package main

import (
	"log"

	_ "github.com/lib/pq"

	"github.com/ibrat-muslim/blog_app_api_gateway/api"
	"github.com/ibrat-muslim/blog_app_api_gateway/config"
	grpcPkg "github.com/ibrat-muslim/blog_app_api_gateway/pkg/grpc_client"
)

func main() {
	cfg := config.Load(".")

	grpcClient, err := grpcPkg.New(cfg)
	if err != nil {
		log.Fatalf("failed to get grpc connection: %v", err)
	}

	apiServer := api.New(&api.RouterOptions{
		Cfg:        &cfg,
		GrpcClient: grpcClient,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
