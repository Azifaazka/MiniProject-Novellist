package repository

import (
	"fmt"

	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
	"gorm.io/gorm"
)

type repositoryMysqlLayerStory struct {
	DB *gorm.DB
}

func (r *repositoryMysqlLayerStory) CreateStorys(cerita model.Cerita) error {
	res := r.DB.Create(&cerita)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayerStory) GetAllStory() []model.Cerita {
	storys := []model.Cerita{}
	r.DB.Find(&storys)

	return storys
}

func (r *repositoryMysqlLayerStory) GetOneStoryByID(id int) (cerita model.Cerita, err error) {
	res := r.DB.Where("id = ?", id).Find(&cerita)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return 
}

func (r *repositoryMysqlLayerStory) UpdateOneStoryByID(id int, cerita model.Cerita) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&cerita)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repositoryMysqlLayerStory) DeleteStoryByID(id int) error {
	res := r.DB.Delete(&model.Cerita{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRepositoryStory(db *gorm.DB) domain.AdapterRepositoryStory {
	return &repositoryMysqlLayerStory{
		DB: db,
	}
}
