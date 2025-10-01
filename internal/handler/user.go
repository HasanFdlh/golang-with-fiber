package handler

import (
	"ms-golang-fiber/config"
	"ms-golang-fiber/internal/model"
	"ms-golang-fiber/internal/usecase"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

type UserUsecase interface {
	Register(user *model.User) error
	Login(email, password string) (model.User, error)
	FindAll() ([]model.User, error)
}

type UserHandler struct {
	usecase usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) *UserHandler {
	return &UserHandler{u}
}

func (h *UserHandler) Register(c *fiber.Ctx) error {
	var user model.RegisterUserRequest

	if err := c.BodyParser(&user); err != nil {
		return config.Error(c, fiber.StatusBadRequest, "Request failed", err.Error())
	}

	// Validasi input
	if err := config.Validate.Struct(user); err != nil {
		return config.Error(c, fiber.StatusBadRequest, "Validation failed", err.Error())
	}

	// Kalau lolos validasi
	if err := h.usecase.Register(&user); err != nil {
		return config.Error(c, fiber.StatusInternalServerError, "Registration failed", err.Error())
	}
	return config.Success(c, user)
}

func (h *UserHandler) Login(c *fiber.Ctx) error {
	var input model.User
	if err := c.BodyParser(&input); err != nil {
		return config.Error(c, fiber.StatusBadRequest, "Request failed", err.Error())
	}
	user, err := h.usecase.Login(input.Email, input.Password)
	if err != nil {
		return config.Error(c, fiber.StatusInternalServerError, "invalid credentials", err.Error())
	}

	// Generate JWT
	claims := jwt.MapClaims{
		"email": user.Email,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, _ := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	return config.Success(c, fiber.Map{"token": t})
}

func (h *UserHandler) GetUsers(c *fiber.Ctx) error {
	users, err := h.usecase.FindAll()
	if err != nil {
		return config.Error(c, fiber.StatusInternalServerError, "Failed to fetch users", err.Error())
	}
	return config.Success(c, users)
}
