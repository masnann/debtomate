package entities

import "time"

type BorrowersModels struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey"`
	Name      string     `json:"name" gorm:"column:name"`
	Address   string     `json:"address" gorm:"column:address"`
	Phone     string     `json:"phone" gorm:"column:phone"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;index"`
}

func (BorrowersModels) TableName() string {
	return "borrowers"
}
