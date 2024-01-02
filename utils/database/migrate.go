package database

import (
	"debtomate/feature/entities"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		entities.UserModels{})

	if err != nil {
		return
	}
	return
}