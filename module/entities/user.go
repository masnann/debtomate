package entities

import "time"

type UserModels struct {
	ID        uint64     `json:"id" gorm:"column:id;primaryKey"`
	Name      string     `json:"name" gorm:"column:name"`
	Email     string     `json:"email" gorm:"column:email"`
	Password  string     `json:"password" gorm:"column:password"`
	Role      string     `json:"role" gorm:"column:role"`
	CreatedAt time.Time  `json:"created_at" gorm:"column:created_at"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"column:deleted_at;index"`
}

func (UserModels) TableName() string {
	return "users"
}
