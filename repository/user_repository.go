package repository

import (
	"fmt"

	"github.com/azifaazka/MiniProject-Novellist/database"
	"github.com/azifaazka/MiniProject-Novellist/model"
)

func CreateUsers(user model.User) error {
	res := database.DB.Create(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func GetAll() []model.User {
	users := []model.User{}
	database.DB.Find(&users)

	return users
}

func GetOneByID(id int) (user model.User, err error) {
	res := database.DB.Where("id = ?", id).Find(&user)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return
}

func UpdateOneByID(id int, user model.User) error {
	res := database.DB.Where("id = ?", id).UpdateColumns(&user)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func DeleteByID(id int) error {
	res := database.DB.Delete(&model.User{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}