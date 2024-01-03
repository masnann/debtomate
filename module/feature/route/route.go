package route

import (
	"debtomate/module/feature/auth"
	"debtomate/module/feature/borrowers"
	"debtomate/utils/token"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB, jwt token.JWTInterface) {
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
	borrowers.InitializeBorrowers(db)
	borrowers.SetupRoutesBorrowers(app, jwt)
}
