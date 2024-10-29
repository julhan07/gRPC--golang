package service

import (
	"context"
	"grpcservices/product_service/client"
	"grpcservices/product_service/entities"
	"grpcservices/product_service/repository"

	"github.com/gofiber/fiber/v2/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductServiceInternal struct {
	repo       repository.ProductRepository
	userClient client.UserClient
}

func NewProductServiceInternal(repo repository.ProductRepository, userClient client.UserClient) *ProductServiceInternal {
	return &ProductServiceInternal{
		repo:       repo,
		userClient: userClient,
	}
}

func (s *ProductServiceInternal) GetProduct(ctx context.Context, productID string) (*entities.Product, error) {
	product, err := s.repo.GetByID(productID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found")
	}
	// Get user info
	userResp, err := s.userClient.GetUser(ctx, product.UserID)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	return &entities.Product{
		ID:       product.ID,
		Name:     product.Name,
		Price:    product.Price,
		UserID:   userResp.UserId,
		UserName: userResp.Name,
	}, nil
}

func (s *ProductServiceInternal) GetUserProducts(ctx context.Context, userID string) ([]*entities.Product, error) {
	userResp, err := s.userClient.GetUser(ctx, userID)
	if err != nil {
		log.Errorf("error", err.Error())
		return nil, status.Errorf(codes.NotFound, err.Error())
	}

	products, err := s.repo.GetByUserID(userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user products")
	}

	var list []*entities.Product
	for _, product := range products {
		list = append(list, &entities.Product{
			ID:       product.ID,
			Name:     product.Name,
			Price:    product.Price,
			UserID:   userResp.UserId, // Sesuaikan dengan field name di proto
			UserName: userResp.Name,
		})
	}

	return list, nil
}
