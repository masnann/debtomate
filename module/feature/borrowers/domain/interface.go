package domain

import (
	"debtomate/module/entities"
	"github.com/gofiber/fiber/v2"
)

type BorrowersRepositoryInterface interface {
	GetPaginatedBorrowers(page, pageSize int) ([]*entities.BorrowersModels, error)
	GetTotalItems() (int64, error)
}

type BorrowersServiceInterface interface {
	GetAllBorrowers(page, pageSize int) ([]*entities.BorrowersModels, int64, error)
	GetBorrowersPage(currentPage, pageSize int) (int, int, int, int, error)
}

type BorrowersHandlerInterface interface {
	GetAllBorrowers(c *fiber.Ctx) error
}
