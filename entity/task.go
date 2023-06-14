package entity

type Task struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	CategoryId  int    `json:"categoryId"`
	Status      bool   `json:"status"`
	UserId      int    `json:"userId"`
}
