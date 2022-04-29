package repository

import (
	"fmt"

	"github.com/azifaazka/MiniProject-Novellist/database"
	"github.com/azifaazka/MiniProject-Novellist/model"
)

func CreateStorys(cerita model.Cerita) error {
	res := database.DB.Create(&cerita)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}
	return nil
}

func GetAll() []model.Cerita {
	storys := []model.Cerita{}
	database.DB.Find(&storys)

	return storys
}

func GetOneStoryByID(id int) (cerita model.Cerita, err error) {
	res := database.DB.Where("id = ?", id).Find(&cerita)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}
	return
}

func UpdateOneStoryByID(id int, cerita model.Cerita) error {
	res := database.DB.Where("id = ?", id).UpdateColumns(&cerita)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func DeleteStoryByID(id int) error {
	res := database.DB.Delete(&model.Cerita{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}