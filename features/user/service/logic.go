package service

import (
	"be13/ca/features/user"
	"errors"

	"github.com/go-playground/validator/v10"
)

type userService struct {
	userRepository user.RepositoryInterface
	validate       *validator.Validate
}

func New(repo user.RepositoryInterface) user.ServiceInterface {
	return &userService{
		userRepository: repo,
		validate:       validator.New(),
	}
}

// Create implements user.ServiceInterface
func (service *userService) Create(input user.Core) (err error) {
	// if input.Name == "" || input.Email == "" || input.Password == "" {
	// 	return errors.New("name, email, password harus diisi")

	// }
	input.Role = "user"
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return errValidate
	}
	_, errCreate := service.userRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

// GetAll implements user.ServiceInterface
func (service *userService) GetAll() (data []user.Core, err error) {
	data, err = service.userRepository.GetAll()
	return

}

func (service *userService) GetById(id int) (data user.Core, err error) {
	data, errGet := service.userRepository.GetById(id)
	if errGet != nil {
		return data, errors.New("failed get user by id data, error query")
	}
	return data, nil
}

func (service *userService) DeleteUser(id int) (err error) {
	_, errDel := service.userRepository.DeleteUser(id)
	if errDel != nil {
		return errors.New("failed delete user, error query")
	}
	return nil
}

func (service *userService) UpdateUser(dataCore user.Core, id int) (err error) {
	errUpdate := service.userRepository.UpdateUser(dataCore, id)
	if errUpdate != nil {
		return errors.New("failed update data, error query")
	}
	return nil

}
