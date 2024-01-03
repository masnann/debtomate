package repository

import (
	"debtomate/module/entities"
	"debtomate/module/feature/borrowers/domain"
	"gorm.io/gorm"
)

type BorrowersRepository struct {
	db *gorm.DB
}

func NewBorrowersRepository(db *gorm.DB) domain.BorrowersRepositoryInterface {
	return &BorrowersRepository{
		db: db,
	}
}

func (r *BorrowersRepository) GetTotalItems() (int64, error) {
	var totalItems int64

	if err := r.db.Model(&entities.BorrowersModels{}).Count(&totalItems).Error; err != nil {
		return 0, err
	}

	return totalItems, nil
}

func (r *BorrowersRepository) GetPaginatedBorrowers(page, pageSize int) ([]*entities.BorrowersModels, error) {
	var borrowers []*entities.BorrowersModels

	offset := (page - 1) * pageSize

	if err := r.db.Offset(offset).Limit(pageSize).Find(&borrowers).Error; err != nil {
		return nil, err
	}

	return borrowers, nil
}
