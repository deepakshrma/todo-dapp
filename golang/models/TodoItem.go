package models

type TodoItem struct {
	Id          int    `form:"Id"`
	Status      int    `form:"Status"`
	Title       string `form:"Title" binding:"required"`
	Description string `form:"Description"`
	CreatedBy   string `form:"CreatedBy"`
}
