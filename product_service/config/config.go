package config

import (
	"grpcservices/product_service/client"
	"grpcservices/product_service/infra"
	"log"
)

type Config struct {
	ServerHost     string
	ServerPort     string
	GrpcServerPort string
	UserServiceURL string
	DBConnection   string
	UserService    client.UserClient
}

func LoadConfig() (*Config, error) {

	cfg := &Config{
		ServerHost:     "localhost",
		ServerPort:     "3000",
		GrpcServerPort: "50051",
		UserServiceURL: "localhost:50052",
		DBConnection:   "localhost:27017",
	}

	userConn, err := infra.ConnectGRPCClient(cfg.UserServiceURL)
	if err != nil {
		log.Fatalf(err.Error())
	}

	userClient := client.NewUserClient(userConn)
	cfg.UserService = userClient

	return cfg, nil
}
