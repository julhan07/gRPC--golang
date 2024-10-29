package service

import (
	"context"
	"grpcservices/user_service/client"
	pb "grpcservices/user_service/proto"
	"grpcservices/user_service/repository"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	repo          repository.UserRepository
	productClient client.ProductClient
}

func NewUserService(repo repository.UserRepository, productClient client.ProductClient) *UserService {
	return &UserService{
		repo:          repo,
		productClient: productClient,
	}
}

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.repo.GetByID(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}

	// Get user's products
	products, err := s.productClient.GetUserProducts(ctx, user.ID)
	if err != nil {
		// Log error but continue
		log.Printf("Failed to get user products: %v", err)
	} else {
		log.Printf("User %s has %d products", user.ID, len(products))
	}

	return &pb.GetUserResponse{
		UserId: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := s.repo.Create(req.Name, req.Email)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create user")
	}

	return &pb.CreateUserResponse{
		UserId: user.ID,
		Name:   user.Name,
		Email:  user.Email,
	}, nil
}

func (s *UserService) ListUsers(ctx context.Context, req *pb.ListUsersRequest) (*pb.ListUsersResponse, error) {
	users, total, err := s.repo.List(req.Page, req.Limit)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to list users")
	}

	var protoUsers []*pb.GetUserResponse
	for _, user := range users {
		protoUsers = append(protoUsers, &pb.GetUserResponse{
			UserId: user.ID,
			Name:   user.Name,
			Email:  user.Email,
		})
	}

	return &pb.ListUsersResponse{
		Users: protoUsers,
		Total: total,
	}, nil
}
