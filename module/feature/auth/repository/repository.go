package repository

import (
	"debtomate/module/entities"
	"debtomate/module/feature/auth/domain"
	"gorm.io/gorm"
)

type AuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) domain.AuthRepositoryInterface {
	return &AuthRepository{
		db: db,
	}
}

func (r *AuthRepository) GetUsersByEmail(email string) (*entities.AdminModels, error) {
	var user entities.AdminModels
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
