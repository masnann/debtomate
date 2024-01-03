package domain

import (
	"debtomate/module/entities"
	"time"
)

type BorrowersResponse struct {
	ID        uint64    `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
}

func ResponseDetailBorrowers(data *entities.BorrowersModels) *BorrowersResponse {
	res := &BorrowersResponse{
		ID:        data.ID,
		Name:      data.Name,
		Address:   data.Address,
		Phone:     data.Phone,
		CreatedAt: data.CreatedAt,
	}
	return res
}

func ResponseArrayBorrowers(data []*entities.BorrowersModels) []*BorrowersResponse {
	res := make([]*BorrowersResponse, 0)

	for _, borrower := range data {
		borrowerRes := &BorrowersResponse{
			ID:        borrower.ID,
			Name:      borrower.Name,
			Address:   borrower.Address,
			Phone:     borrower.Phone,
			CreatedAt: borrower.CreatedAt,
		}
		res = append(res, borrowerRes)
	}

	return res
}
