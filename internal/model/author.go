package model

type Author struct {
	Id      uint64 `json:"author_id"`
	Name    string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
}
