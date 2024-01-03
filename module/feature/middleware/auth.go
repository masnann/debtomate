package middleware

import (
	"debtomate/utils/response"
	"debtomate/utils/token"
	"github.com/gofiber/fiber/v2"
	"strings"
)

func Auth(jwtService token.JWTInterface) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" || len(strings.Split(authHeader, "Bearer ")) != 2 {
			return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid or missing authorization header.")
		}

		tokenString := strings.Split(authHeader, "Bearer ")[1]

		validToken, err := jwtService.ValidateToken(tokenString)
		if err != nil || !validToken.Valid {
			return response.ErrorBuildResponse(c, fiber.StatusUnauthorized, "Unauthorized: Invalid token")
		}

		return c.Next()
	}
}
