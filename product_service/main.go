package main

import (
	"grpcservices/product_service/cmd/server"
	"grpcservices/product_service/config"
	"log"
	"sync"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	// Create gRPC server
	grpcServer := server.NewGRPCServer(cfg)
	// Create Gateway server
	httpServer := server.NewGatewayServer(cfg)
	// Run both servers in goroutines
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := grpcServer.Run(); err != nil {
			log.Fatalf("Failed to run gRPC server: %v", err)
		}
	}()

	// Run HTTP gateway
	go func() {
		defer wg.Done()
		if err := httpServer.Run(); err != nil {
			log.Fatalf("Failed to run gateway server: %v", err)
		}
	}()

	wg.Wait()
}
