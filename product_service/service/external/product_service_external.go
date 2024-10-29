package service

import (
	"context"
	"grpcservices/product_service/client"
	pb "grpcservices/product_service/proto"
	"grpcservices/product_service/repository"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ProductService struct {
	pb.UnimplementedProductServiceServer
	repo       repository.ProductRepository
	userClient client.UserClient
}

func NewProductService(repo repository.ProductRepository, userClient client.UserClient) *ProductService {
	return &ProductService{
		repo:       repo,
		userClient: userClient,
	}
}

func (s *ProductService) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	product, err := s.repo.GetByID(req.ProductId)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "product not found")
	}

	// Get user info
	userResp, err := s.userClient.GetUser(ctx, product.UserID)
	if err != nil {
		// Log error but continue
		log.Printf("Failed to get user info: %v", err)
	} else {
		log.Printf("Product belongs to user: %s", userResp.Name)
	}

	return &pb.GetProductResponse{
		ProductId: product.ID,
		Name:      product.Name,
		Price:     product.Price,
		UserId:    product.UserID, // Sesuaikan dengan field name di proto
	}, nil
}

func (s *ProductService) GetUserProducts(ctx context.Context, req *pb.GetUserProductsRequest) (*pb.GetUserProductsResponse, error) {
	products, err := s.repo.GetByUserID(req.UserId)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get user products")
	}

	var protoProducts []*pb.GetProductResponse
	for _, product := range products {
		protoProducts = append(protoProducts, &pb.GetProductResponse{
			ProductId: product.ID,
			Name:      product.Name,
			Price:     product.Price,
			UserId:    product.UserID, // Sesuaikan dengan field name di proto
		})
	}

	return &pb.GetUserProductsResponse{
		Products: protoProducts,
	}, nil
}
