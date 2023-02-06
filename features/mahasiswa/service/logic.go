package service

import (
	"be13/ca/features/mahasiswa"
	"errors"
)

type mahasiswaService struct {
	mahasiswaRepository mahasiswa.RepositoryInterface
}

func New(repo mahasiswa.RepositoryInterface) mahasiswa.ServiceInterface {
	return &mahasiswaService{
		mahasiswaRepository: repo,
	}
}

// Create implements mahasiswa.ServiceInterface
func (srv *mahasiswaService) Create(input mahasiswa.Core) (err error) {
	errCreate := srv.mahasiswaRepository.Create(input)
	if errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

// DeleteUser implements mahasiswa.ServiceInterface
func (srv *mahasiswaService) Delete(id int) (err error) {
	errCreate := srv.mahasiswaRepository.Delete(id)
	if errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

// UpdateUser implements mahasiswa.ServiceInterface
func (srv *mahasiswaService) Update(input mahasiswa.Core, id int) (err error) {
	errCreate := srv.mahasiswaRepository.Update(input, id)
	if errCreate != nil {
		return errors.New("failed insert data, error query")
	}
	return nil
}

// Read implements mahasiswa.ServiceInterface
func (srv *mahasiswaService) Read(id int) (data []mahasiswa.NilaiMhs, err error) {
	res, errCreate := srv.mahasiswaRepository.Read(id)
	if errCreate != nil {
		return res, errors.New("failed insert data, error query")
	}
	return res, nil
}
