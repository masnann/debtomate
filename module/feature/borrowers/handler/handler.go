package handler

import (
	"debtomate/module/feature/borrowers/domain"
	"debtomate/utils/response"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type BorrowersHandler struct {
	service domain.BorrowersServiceInterface
}

func NewBorrowersHandler(service domain.BorrowersServiceInterface) domain.BorrowersHandlerInterface {
	return &BorrowersHandler{
		service: service,
	}
}

func (h *BorrowersHandler) GetAllBorrowers(c *fiber.Ctx) error {
	currentPage, err := strconv.Atoi(c.Query("page"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page number")
	}

	pageSize, err := strconv.Atoi(c.Query("page_size"))
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusBadRequest, "Invalid page size")
	}

	result, totalItems, err := h.service.GetAllBorrowers(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get borrowers: "+err.Error())
	}

	currentPage, totalPages, nextPage, prevPage, err := h.service.GetBorrowersPage(currentPage, pageSize)
	if err != nil {
		return response.ErrorBuildResponse(c, fiber.StatusInternalServerError, "Failed to get page info: "+err.Error())
	}

	return response.PaginationBuildResponse(c, fiber.StatusOK, "Success get pagination", domain.ResponseArrayBorrowers(result), currentPage, int(totalItems), totalPages, nextPage, prevPage)
}
