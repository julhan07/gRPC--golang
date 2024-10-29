package infra

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ConnectGRPCClient(serviceURL string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(
		serviceURL,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to service: %w", err)
	}
	log.Printf("Successfully connected to service: %s", serviceURL)
	return conn, nil
}

func StartGRPCServer(port string, registerFuncs ...func(*grpc.Server)) error {
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return fmt.Errorf("failed to listen on port %s: %w", port, err)
	}
	defer listener.Close()

	grpcServer := grpc.NewServer()
	for _, registerFunc := range registerFuncs {
		registerFunc(grpcServer)
	}

	log.Printf("gRPC server starting on port %s", port)
	return grpcServer.Serve(listener)
}
