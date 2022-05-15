package repository

import (
	"fmt"

	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
	"gorm.io/gorm"
)

type repositoryMysqlLayerBab struct {
	DB *gorm.DB
}

func (r *repositoryMysqlLayerBab) CreateBabs(bab model.Bab) error {
	res := r.DB.Create(&bab)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayerBab) GetAllBab() []model.Bab {
	bab := []model.Bab{}
	r.DB.Find(&bab)

	return bab
}

func (r *repositoryMysqlLayerBab) GetOneBabByID(id int) (bab model.Bab, err error) {
	res := r.DB.Where("id = ?", id).Find(&bab)
	if res.RowsAffected < 1 {
		err = fmt.Errorf("not found")
	}

	return 
}

func (r *repositoryMysqlLayerBab) UpdateOneBabByID(id int, bab model.Bab) error {
	res := r.DB.Where("id = ?", id).UpdateColumns(&bab)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error update")
	}

	return nil
}

func (r *repositoryMysqlLayerBab) DeleteBabsByID(id int) error {
	res := r.DB.Delete(&model.Bab{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRepositoryBab(db *gorm.DB) domain.AdapterRepositoryBab {
	return &repositoryMysqlLayerBab{
		DB: db,
	}
}
