package client

import (
	"context"
	productpb "grpcservices/product_service/proto"

	"google.golang.org/grpc"
)

type ProductClient interface {
    GetUserProducts(ctx context.Context, userID string) ([]*productpb.GetProductResponse, error)
}

type productClient struct {
    client productpb.ProductServiceClient
}

func NewProductClient(conn *grpc.ClientConn) ProductClient {
    return &productClient{
        client: productpb.NewProductServiceClient(conn),
    }
}

func (c *productClient) GetUserProducts(ctx context.Context, userID string) ([]*productpb.GetProductResponse, error) {
    resp, err := c.client.GetUserProducts(ctx, &productpb.GetUserProductsRequest{
        UserId: userID,
    })
    if err != nil {
        return nil, err
    }
    return resp.Products, nil
}
