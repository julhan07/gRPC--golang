package server

import (
	"fmt"
	"log"
	"time"

	"grpcservices/product_service/cmd/server/routes" // Memanggil routes.go
	"grpcservices/product_service/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

type GatewayServer struct {
	cfg *config.Config
}

func NewGatewayServer(cfg *config.Config) *GatewayServer {
	return &GatewayServer{
		cfg: cfg,
	}
}

func (s *GatewayServer) Run() error {

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Printf("Error: %v", err)
			return c.Status(code).JSON(fiber.Map{
				"error":   err.Error(),
				"status":  code,
				"message": "An error occurred",
			})
		},
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	})

	// Add middlewares
	app.Use(recover.New())
	app.Use(logger.New(logger.Config{
		Format: "[${time}] ${status} - ${latency} ${method} ${path}\n",
	}))
	app.Use(cors.New(cors.Config{
		AllowOrigins:     fmt.Sprintf("http://%s:%s", s.cfg.ServerHost, s.cfg.ServerPort),
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS,PATCH",
		AllowHeaders:     "Origin,Content-Type,Accept,Authorization",
		AllowCredentials: true,
		ExposeHeaders:    "Content-Length,Content-Type",
		MaxAge:           86400,
	}))

	// Routes
	r := routes.NewRoutes(app, s.cfg) // Menginisialisasi routes
	r.Setup()                         // Mengatur routes

	log.Printf("HTTP Gateway server starting on port %s", s.cfg.ServerPort)
	return app.Listen(fmt.Sprintf("%s:%s", s.cfg.ServerHost, s.cfg.ServerPort))
}
