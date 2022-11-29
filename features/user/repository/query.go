package repository

import (
	"be13/ca/features/user"
	"errors"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// dependency injection
func New(db *gorm.DB) user.RepositoryInterface {
	return &userRepository{
		db: db,
	}
}

// Create implements user.Repository
func (repo *userRepository) Create(input user.Core) (row int, err error) {
	userGorm := fromCore(input)
	tx := repo.db.Create(&userGorm) // proses insert data
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failed")
	}
	return int(tx.RowsAffected), nil
}

// GetAll implements user.Repository
func (repo *userRepository) GetAll() (data []user.Core, err error) {
	var users []User

	tx := repo.db.Find(&users)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var dataCore = toCoreList(users)
	return dataCore, nil
}

func (repo *userRepository) GetById(id int) (data user.Core, err error) {
	var IdUser User
	var IdUserCore = user.Core{}
	IdUser.ID = uint(id)
	tx := repo.db.First(&IdUser, IdUser.ID)
	if tx.Error != nil {
		return IdUserCore, tx.Error
	}
	IdUserCore = IdUser.toCore()
	return IdUserCore, nil
}

func (repo *userRepository) DeleteUser(id int) (row int, err error) {
	idUser := User{}

	tx := repo.db.Delete(&idUser, id)
	if tx.Error != nil {
		return -1, tx.Error
	}
	if tx.RowsAffected == 0 {
		return -1, errors.New("delete user by id failed")
	}
	return int(tx.RowsAffected), nil
}

func (repo *userRepository) UpdateUser(datacore user.Core, id int) (err error) {
	userGorm := fromCore(datacore)
	tx := repo.db.Where("id= ?", id).Updates(userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return errors.New("update user failed")
	}
	return nil
}
