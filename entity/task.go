package entity

import (
	"time"
)

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      bool      `json:"status"`
	UserId      int       `json:"userId"`
	CategoryId  int       `json:"categoryId"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

// func (i *Task) TaskToTaskRespon() dto.task {
// 	return dto.ItemResponse{
// 		Id:          i.ItemId,
// 		ItemCode:    i.ItemCode,
// 		Quantity:    i.Quantity,
// 		Description: i.Description,
// 		OrderId:     i.OrderId,
// 		CreatedAt:   i.CreatedAt,
// 		UpdatedAt:   i.UpdatedAt,
// 	}
// }
