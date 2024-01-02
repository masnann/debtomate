package handler

import (
	"debtomate/module/feature/auth/domain"
	"debtomate/utils/response"
	"debtomate/utils/validator"
	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service domain.AuthServiceInterface
}

func NewAuthHandler(service domain.AuthServiceInterface) domain.AuthHandlerInterface {
	return &AuthHandler{
		service: service,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	req := new(validator.LoginRequest)
	if err := c.BodyParser(req); err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "err")
	}

	if err := validator.ValidateLoginRequest(req); err != nil {
		return response.ErrorResponse(c, fiber.StatusBadRequest, "Validation failed: "+err.Error())
	}

	user, token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return response.ErrorResponse(c, fiber.StatusInternalServerError, "Status internal server error: "+err.Error())
	}

	return response.SuccessResponse(c, fiber.StatusOK, "Login successfully", domain.LoginFormatter(user, token))
}
