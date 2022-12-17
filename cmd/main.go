package main

import (
	_ "github.com/lib/pq"

	"github.com/ibrat-muslim/blog_app_api_gateway/api"
	"github.com/ibrat-muslim/blog_app_api_gateway/config"
	grpcPkg "github.com/ibrat-muslim/blog_app_api_gateway/pkg/grpc_client"
	"github.com/ibrat-muslim/blog_app_api_gateway/pkg/logger"
)

func main() {
	cfg := config.Load(".")

	log := logger.New()

	grpcClient, err := grpcPkg.New(cfg)
	if err != nil {
		log.Fatalf("failed to get grpc connection: %v", err)
	}

	apiServer := api.New(&api.RouterOptions{
		Cfg:        &cfg,
		GrpcClient: grpcClient,
		Logger:     log,
	})

	err = apiServer.Run(cfg.HttpPort)
	if err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
