package borrowers

import (
	"debtomate/module/feature/borrowers/domain"
	"debtomate/module/feature/borrowers/handler"
	"debtomate/module/feature/borrowers/repository"
	"debtomate/module/feature/borrowers/service"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var (
	borrowersRepo    domain.BorrowersRepositoryInterface
	borrowersService domain.BorrowersServiceInterface
	borrowersHandler domain.BorrowersHandlerInterface
)

func InitializeBorrowers(db *gorm.DB) {
	borrowersRepo = repository.NewBorrowersRepository(db)
	borrowersService = service.NewBorrowersService(borrowersRepo)
	borrowersHandler = handler.NewBorrowersHandler(borrowersService)
}

func SetupRoutesBorrowers(app *fiber.App) {
	api := app.Group("/api/v1/borrowers")
	api.Get("", borrowersHandler.GetAllBorrowers)
}
