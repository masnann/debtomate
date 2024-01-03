package service

import (
	"debtomate/module/entities"
	"debtomate/module/feature/borrowers/domain"
	"math"
)

type BorrowersService struct {
	repo domain.BorrowersRepositoryInterface
}

func NewBorrowersService(repo domain.BorrowersRepositoryInterface) domain.BorrowersServiceInterface {
	return &BorrowersService{
		repo: repo,
	}
}

func (s *BorrowersService) GetAllBorrowers(page, pageSize int) ([]*entities.BorrowersModels, int64, error) {
	result, err := s.repo.GetPaginatedBorrowers(page, pageSize)
	if err != nil {
		return nil, 0, err
	}

	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return nil, 0, err
	}

	return result, totalItems, nil
}

func (s *BorrowersService) GetBorrowersPage(currentPage, pageSize int) (int, int, int, int, error) {
	totalItems, err := s.repo.GetTotalItems()
	if err != nil {
		return 0, 0, 0, 0, err
	}

	totalPages := int(math.Ceil(float64(totalItems) / float64(pageSize)))
	nextPage := currentPage + 1
	prevPage := currentPage - 1

	if nextPage > totalPages {
		nextPage = 0
	}

	if prevPage < 1 {
		prevPage = 0
	}

	return currentPage, totalPages, nextPage, prevPage, nil
}
