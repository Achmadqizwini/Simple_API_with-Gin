package delivery

import "be13/ca/features/book"

type BookRequest struct {
	Title       string `json:"title" form:"title"`
	Publisher   string `json:"publisher" form:"publisher"`
	Author      string `json:"author" form:"author"`
	PublishYear string `json:"publish_year" form:"publish_year"`
	UserID      uint   `json:"user_id" form:"user_id"`
}

func BookRequestToCore(dataReq BookRequest) book.Core {
	return book.Core{
		Title:       dataReq.Title,
		Publisher:   dataReq.Publisher,
		Author:      dataReq.Author,
		PublishYear: dataReq.PublishYear,
		UserID:      dataReq.UserID,
	}
}
