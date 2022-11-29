package delivery

import (
	"be13/ca/features/book"
)

type BookResponse struct {
	ID          uint   `json:"ID"`
	Title       string `json:"title"`
	Publisher   string `json:"publisher"`
	Author      string `json:"author"`
	PublishYear string `json:"publish_year"`
	UserID      uint   `json:"user_id"`
}

func DataIDBookRespon(book book.Core) BookResponse {
	dataResponse := BookResponse{
		ID:          book.ID,
		Title:       book.Title,
		Publisher:   book.Publisher,
		Author:      book.Author,
		PublishYear: book.PublishYear,
		UserID:      book.UserID,
	}
	return dataResponse
}

func ListModelBook(dataModel []book.Core) []BookResponse {
	var result []BookResponse
	for _, v := range dataModel {
		result = append(result, DataIDBookRespon(v))
	}
	return result
}
