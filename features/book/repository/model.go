package repository

import (
	_book "be13/ca/features/book"

	"gorm.io/gorm"
)

// struct gorm model
type Book struct {
	gorm.Model
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
	// Books    []Book
}

// DTO
// mapping

// mengubah struct core ke struct model gorm
func fromCore(dataCore _book.Core) Book {
	bookGorm := Book{
		Title:       dataCore.Title,
		Publisher:   dataCore.Publisher,
		Author:      dataCore.Author,
		PublishYear: dataCore.PublishYear,
		UserID:      dataCore.UserID,
	}
	return bookGorm
}

// mengubah struct model gorm ke struct core
func (dataModel *Book) toCore() _book.Core {
	return _book.Core{
		ID:          dataModel.ID,
		Title:       dataModel.Title,
		Publisher:   dataModel.Publisher,
		Author:      dataModel.Author,
		PublishYear: dataModel.PublishYear,
		UserID:      dataModel.UserID,
	}
}

// mengubah slice struct model gorm ke slice struct core
func toCoreList(dataModel []Book) []_book.Core {
	var dataCore []_book.Core
	for _, v := range dataModel {
		dataCore = append(dataCore, v.toCore())
	}
	return dataCore
}
