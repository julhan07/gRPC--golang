package main

import (
	"grpcservices/user_service/client"
	"grpcservices/user_service/config"
	pb "grpcservices/user_service/proto"
	"grpcservices/user_service/repository"
	"grpcservices/user_service/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Setup gRPC server
	lis, err := net.Listen("tcp", cfg.ServerPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Connect to Product Service
	productConn, err := grpc.Dial(cfg.ProductServiceURL, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to product service: %v", err)
	}
	defer productConn.Close()

	// Initialize dependencies
	productClient := client.NewProductClient(productConn)
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo, productClient)

	// Create gRPC server
	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userService)

	log.Printf("Starting user service on %s", lis.Addr().String())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
