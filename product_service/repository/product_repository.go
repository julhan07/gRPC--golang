package repository

import (
	"fmt"
	"grpcservices/product_service/entities"
)

type ProductRepository interface {
	GetByID(id string) (*entities.Product, error)
	GetByUserID(userID string) ([]*entities.Product, error)
}

type productRepository struct {
	// Bisa menggunakan database seperti MySQL, MongoDB, dll
	products map[string]*entities.Product
}

func NewProductRepository() ProductRepository {
	// Simulasi database
	products := map[string]*entities.Product{
		"1": {ID: "1", Name: "Product 1", Price: 100, UserID: "user1"},
		"2": {ID: "2", Name: "Product 2", Price: 200, UserID: "user1"},
	}
	return &productRepository{products: products}
}

func (r *productRepository) GetByID(id string) (*entities.Product, error) {
	if product, exists := r.products[id]; exists {
		return product, nil
	}
	return nil, fmt.Errorf("product not found")
}

func (r *productRepository) GetByUserID(userID string) ([]*entities.Product, error) {
	var userProducts []*entities.Product
	for _, product := range r.products {
		if product.UserID == userID {
			userProducts = append(userProducts, product)
		}
	}
	return userProducts, nil
}
