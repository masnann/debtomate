package entities

type UserModels struct {
	ID       uint64 `json:"id" gorm:"column:id;primaryKey"`
	Name     string `json:"name" gorm:"column:name"`
	Email    string `json:"email" gorm:"column:email"`
	Password string `json:"password" gorm:"column:password"`
}

func (UserModels) TableName() string {
	return "users"
}
