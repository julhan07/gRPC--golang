package client

import (
	"context"
	userpb "grpcservices/user_service/proto"

	"google.golang.org/grpc"
)

type UserClient interface {
    GetUser(ctx context.Context, userID string) (*userpb.GetUserResponse, error)
}

type userClient struct {
    client userpb.UserServiceClient
}

func NewUserClient(conn *grpc.ClientConn) UserClient {
    return &userClient{
        client: userpb.NewUserServiceClient(conn),
    }
}

func (c *userClient) GetUser(ctx context.Context, userID string) (*userpb.GetUserResponse, error) {
    return c.client.GetUser(ctx, &userpb.GetUserRequest{UserId: userID})
}
