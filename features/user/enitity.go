package user

import "time"

type Core struct {
	ID        uint
	Name      string `validate:"required"`
	Email     string `validate:"required"`
	Password  string `validate:"required"`
	Phone     string `validate:"required"`
	Address   string
	Role      string `validate:"required"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (err error)
	GetById(id int) (data Core, err error)
	DeleteUser(id int) (err error)
	UpdateUser(input Core, id int) (err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (row int, err error)
	GetById(id int) (data Core, err error)
	DeleteUser(id int) (row int, err error)
	UpdateUser(input Core, id int) (err error)
}
