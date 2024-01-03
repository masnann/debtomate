package domain

import (
	"debtomate/module/entities"
	"github.com/gofiber/fiber/v2"
)

type AuthRepositoryInterface interface {
	GetUsersByEmail(email string) (*entities.AdminModels, error)
}

type AuthServiceInterface interface {
	Login(email, password string) (*entities.AdminModels, string, error)
}

type AuthHandlerInterface interface {
	Login(c *fiber.Ctx) error
}
