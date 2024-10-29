package handler

import (
	"grpcservices/product_service/service"
	"log"

	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	productService service.ProductServiceInternal
}

func NewProductHandler(productService service.ProductServiceInternal) *ProductHandler {
	return &ProductHandler{
		productService: productService,
	}
}

func (h *ProductHandler) GetProductByID(c *fiber.Ctx) error {
	productId := c.Params("id")
	if productId == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Product ID is required",
		})
	}

	resp, err := h.productService.GetProduct(c.Context(), productId)
	if err != nil {
		log.Printf("Error getting products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get products",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   resp,
	})
}

func (h *ProductHandler) GetProductByUserID(c *fiber.Ctx) error {
	userID := c.Query("user_id")
	if userID == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "UserID is required",
		})
	}

	resp, err := h.productService.GetUserProducts(c.Context(), userID)
	if err != nil {
		log.Printf("Error getting products: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to get products",
		})
	}

	return c.JSON(fiber.Map{
		"status": "success",
		"data":   resp,
	})
}
