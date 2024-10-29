package server

import (
	"fmt"
	"grpcservices/product_service/config"
	"grpcservices/product_service/infra"
	pb "grpcservices/product_service/proto"
	"grpcservices/product_service/repository"
	service "grpcservices/product_service/service/external"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	conf *config.Config
}

func NewGRPCServer(conf *config.Config) *GRPCServer {
	return &GRPCServer{
		conf: conf,
	}
}

func (s *GRPCServer) Run() error {

	userRepo := repository.NewProductRepository()
	err := infra.StartGRPCServer(s.conf.GrpcServerPort, func(gs *grpc.Server) {
		pb.RegisterProductServiceServer(gs, service.NewProductService(userRepo, s.conf.UserService))
	})
	if err != nil {
		return fmt.Errorf("failed to start gRPC server: %w", err)
	}

	return err
}
