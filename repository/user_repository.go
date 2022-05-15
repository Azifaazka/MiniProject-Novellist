package repository

import (
	"fmt"

	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
	"gorm.io/gorm"
)

type repositoryMysqlLayer struct {
	DB *gorm.DB
}

func (r *repositoryMysqlLayer) CreateUsers(user model.User) error {
	res := r.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayer) GetAll() []model.User {
	users := []model.User{}
	r.DB.Find(&users)

	return users
}

func (r *repositoryMysqlLayer) GetOneByID(id int) (user model.User, err error) {
	res := r.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryMysqlLayer) GetOneByEmail(email string) (user model.User, err error) {
	res := r.DB.Where("email = ?", email).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func (r *repositoryMysqlLayer) UpdateOneByID(id int, user model.User) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repositoryMysqlLayer) DeleteByID(id int) error {
	res := r.DB.Delete(&model.User{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRepository(db *gorm.DB) domain.AdapterRepository {
	return &repositoryMysqlLayer{
		DB: db,
	}
}