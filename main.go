package main

import (
	"log"
	"ms-golang-fiber/config"
	"ms-golang-fiber/internal/handler"
	"ms-golang-fiber/internal/migration"
	"ms-golang-fiber/internal/repository"
	"ms-golang-fiber/internal/usecase"
	"ms-golang-fiber/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/spf13/viper"
)

func main() {
	// Load env
	config.LoadConfig()

	// Init logger
	config.InitFiberLogger()

	// Init DB
	dbPostgres := config.InitPostgres()
	dbMysql := config.InitMysql()

	// Run migrations hanya untuk Postgres (atau MySQL jika mau juga)
	migration.Migrate(dbPostgres)
	migration.Migrate(dbMysql)

	// Init Redis & Minio
	config.InitRedis()
	config.InitMinio()

	// Init Validator
	config.InitValidator()

	// Dependencies
	userRepo := repository.NewUserRepository(dbPostgres, dbMysql)
	userUsecase := usecase.NewUserUsecase(userRepo)
	userHandler := handler.NewUserHandler(userUsecase)

	// Init Fiber
	app := fiber.New(fiber.Config{
		ErrorHandler: config.ErrorHandler, // global error JSON
	})

	// Middleware global
	app.Use(cors.New())
	app.Use(logger.New(config.InitFiberLogger()))

	// Health check
	app.Get("/ping", func(c *fiber.Ctx) error {
		return config.Success(c, "pong")
	})

	// Routes
	api := app.Group("/api")
	routes.UserRoutes(api, userHandler)

	// Not found handler
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"status":  "error",
			"message": "Endpoint not found",
		})
	})

	// Start server
	port := viper.GetString("APP_PORT")
	if port == "" {
		port = "3000"
	}
	log.Println("🚀 Server running on port", port)
	app.Listen(":" + port)
}
