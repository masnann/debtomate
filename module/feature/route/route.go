package route

import (
	"debtomate/module/feature/auth"
	"debtomate/module/feature/borrowers"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func SetupRoutes(app *fiber.App, db *gorm.DB) {
	auth.InitializeAuth(db)
	auth.SetupRoutesAuth(app)
	borrowers.InitializeBorrowers(db)
	borrowers.SetupRoutesBorrowers(app)
}
