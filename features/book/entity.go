package book

type Core struct {
	ID          uint
	Title       string
	Publisher   string
	Author      string
	PublishYear string
	UserID      uint
}

type ServiceInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (err error)
	GetById(id int) (data Core, err error)
	DeleteBook(id int) (err error)
	UpdateBook(input Core, id int) (err error)
}

type RepositoryInterface interface {
	GetAll() (data []Core, err error)
	Create(input Core) (row int, err error)
	GetById(id int) (data Core, err error)
	DeleteBook(id int) (row int, err error)
	UpdateBook(input Core, id int) (err error)
}
