package routes

import (
	"ms-golang-fiber/internal/handler"
	"ms-golang-fiber/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoutes(r fiber.Router, userHandler *handler.UserHandler) {

	auth := r.Group("/auth")
	auth.Post("/register", userHandler.Register)
	auth.Post("/login", userHandler.Login)

	users := r.Group("/users")
	users.Use(middleware.JWTProtected())
	users.Get("/", userHandler.GetUsers)
}
