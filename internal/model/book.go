package model

type Book struct {
	Id              uint64 `json:"book_id"`
	Title           string `json:"title" binding:"required"`
	Description     string `json:"description"`
	PublicationDate string `json:"publication_date"`
}
