package repository

import (
	"be13/ca/features/book"
	"errors"

	"gorm.io/gorm"
)

type bookRepository struct {
	db *gorm.DB
}

// dependency injection
func New(db *gorm.DB) book.RepositoryInterface {
	return &bookRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *bookRepository) Create(input book.Core) (row int, err error) {
	bookGorm := fromCore(input)
	tx := repo.db.Create(&bookGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *bookRepository) GetAll() (data []book.Core, err error) {
	var books []Book

	tx := repo.db.Find(&books)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(books)
	return dataCore, nil
}

func (repo *bookRepository) GetById(id int) (data book.Core, err error) {
	var IdBook Book
	var IdBookCore = book.Core{}
	IdBook.ID = uint(id)
	tx := repo.db.First(&IdBook, IdBook.ID)
	if tx.Error != nil {
		return IdBookCore, tx.Error
	}
	IdBookCore = IdBook.toCore()
	return IdBookCore, nil
}

func (repo *bookRepository) DeleteBook(id int) (row int, err error) {
	idBook := Book{}

	tx := repo.db.Delete(&idBook, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete book by id failed")
	}
	return int(tx.RowsAffected), nil
}

func (repo *bookRepository) UpdateBook(datacore book.Core, id int) (err error) {
	bookGorm := fromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(bookGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update user failed")
	}
	return nil
}
