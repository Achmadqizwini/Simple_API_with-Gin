package service

import (
	"be13/ca/features/book"
	"errors"

	"github.com/go-playground/validator/v10"
)

type bookService struct {
	bookRepository book.RepositoryInterface
	validate       *validator.Validate
}

func New(repo book.RepositoryInterface) book.ServiceInterface {
	return &bookService{
		bookRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements book.ServiceInterface
func (service *bookService) Create(input book.Core) (err error) {
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("name, email, password harus diisi")

	// }
	_, errCreate := service.bookRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

// GetAll implements bookbook.ServiceInterface
func (service *bookService) GetAll() (data []book.Core, err error) {
	data, err = service.bookRepository.GetAll()
	return

}

func (service *bookService) GetById(id int) (data book.Core, err error) {
	data, errGet := service.bookRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get bookbook by id data, error query")
	}
	return data, nil
}

func (service *bookService) DeleteBook(id int) (err error) {
	_, errDel := service.bookRepository.DeleteBook(id)
	if errDel != nil {
		return errors.New("failed delete user, error query")
	}
	return nil
}

func (service *bookService) UpdateBook(dataCore book.Core, id int) (err error) {
	errUpdate := service.bookRepository.UpdateBook(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}
