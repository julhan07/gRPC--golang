package routes

import (
	"grpcservices/product_service/config"
	"grpcservices/product_service/handler"
	"grpcservices/product_service/repository"
	"grpcservices/product_service/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

type Routes struct {
	app  *fiber.App
	conf *config.Config
}

func NewRoutes(app *fiber.App, conf *config.Config) *Routes {
	return &Routes{app: app, conf: conf}
}

func (r *Routes) Setup() {
	// Middleware
	r.app.Use(logger.New())

	// API Routes v1
	api := r.app.Group("/api")
	v1 := api.Group("/v1")

	productRepo := repository.NewProductRepository()
	productService := service.NewProductServiceInternal(productRepo, r.conf.UserService) // Menggunakan service yang diimpor
	producHandler := handler.NewProductHandler(*productService)

	// Products
	products := v1.Group("/products")
	products.Get("/", producHandler.GetProductByUserID)
	products.Get("/id/:id", producHandler.GetProductByID)

	// Health Check
	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": "ok",
		})
	})
}
