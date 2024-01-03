package service

import (
	"debtomate/module/entities"
	"debtomate/module/feature/borrowers/domain"
	"debtomate/module/feature/borrowers/mocks"
	"errors"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func setupTest(t *testing.T) (*mocks.BorrowersRepositoryInterface, domain.BorrowersServiceInterface) {
	repo := mocks.NewBorrowersRepositoryInterface(t)
	service := NewBorrowersService(repo)
	return repo, service
}

func TestGetBorrowersPage(t *testing.T) {
	currentPage := 2
	pageSize := 5

	expectedCurrentPage := 2
	expectedTotalPages := int(math.Ceil(float64(10) / float64(pageSize)))
	expectedNextPage := 0
	expectedPrevPage := 1

	t.Run("Success Case - Get Total Items", func(t *testing.T) {
		repo, service := setupTest(t)
		repo.On("GetTotalItems").Return(int64(10), nil)

		currentPageResult, totalPagesResult, nextPageResult, prevPageResult, err := service.GetBorrowersPage(currentPage, pageSize)

		assert.Nil(t, err)
		assert.Equal(t, expectedCurrentPage, currentPageResult)
		assert.Equal(t, expectedTotalPages, totalPagesResult)
		assert.Equal(t, expectedNextPage, nextPageResult)
		assert.Equal(t, expectedPrevPage, prevPageResult)
	})

	t.Run("Failed Case - GetTotalItems Error", func(t *testing.T) {
		repo, service := setupTest(t)
		mockError := errors.New("failed to get total items")
		repo.On("GetTotalItems").Return(int64(0), mockError)

		_, _, _, _, err := service.GetBorrowersPage(currentPage, pageSize)

		assert.NotNil(t, err)
		assert.EqualError(t, err, "failed to get total items")
	})

	t.Run("Success Case - PrevPage Handling", func(t *testing.T) {
		repo, service := setupTest(t)
		repo.On("GetTotalItems").Return(int64(10), nil)

		currentPage := 1

		currentPageResult, _, _, prevPageResult, _ := service.GetBorrowersPage(currentPage, pageSize)

		assert.Equal(t, currentPage, currentPageResult)
		assert.Equal(t, 0, prevPageResult)
	})

}

func TestGetAllBorrowers(t *testing.T) {
	var expectedResults []*entities.BorrowersModels
	expectedTotalItems := int64(10)

	page := 1
	pageSize := 5

	t.Run("Success Case - Get All Pagination", func(t *testing.T) {
		repo, service := setupTest(t)
		repo.On("GetPaginatedBorrowers", page, pageSize).Return(expectedResults, nil)
		repo.On("GetTotalItems").Return(expectedTotalItems, nil)

		results, totalItems, err := service.GetAllBorrowers(page, pageSize)

		assert.Nil(t, err)
		assert.Equal(t, expectedResults, results)
		assert.Equal(t, expectedTotalItems, totalItems)

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Get Pagination", func(t *testing.T) {
		repo, service := setupTest(t)
		expectedError := errors.New("failed to get paginated borrowers")
		repo.On("GetPaginatedBorrowers", page, pageSize).Return(nil, expectedError)

		_, _, err := service.GetAllBorrowers(page, pageSize)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedError.Error())

		repo.AssertExpectations(t)
	})

	t.Run("Failed Case - Error Get Total Items", func(t *testing.T) {
		repo, service := setupTest(t)
		expectedError := errors.New("failed to get total items")
		repo.On("GetPaginatedBorrowers", page, pageSize).Return(expectedResults, nil)
		repo.On("GetTotalItems").Return(int64(0), expectedError)

		_, _, err := service.GetAllBorrowers(page, pageSize)

		assert.Error(t, err)
		assert.EqualError(t, err, expectedError.Error())

		repo.AssertExpectations(t)
	})

}
