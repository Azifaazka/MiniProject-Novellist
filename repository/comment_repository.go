package repository

import (
	"fmt"

	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
	"gorm.io/gorm"
)

type repositoryMysqlLayerComment struct {
	DB *gorm.DB
}

func (r *repositoryMysqlLayerComment) CreateComment(komen model.Comment) error {
	res := r.DB.Create(&komen)
	if res.RowsAffected < 1 {
		return fmt.Errorf("error insert")
	}

	return nil
}

func (r *repositoryMysqlLayerComment) GetAllComment() []model.Comment {
	komen := []model.Comment{}
	r.DB.Find(&komen)

	return komen
}

func (r *repositoryMysqlLayerComment) DeleteCommentByID(id int) error {
	res := r.DB.Delete(&model.Comment{
		ID: id,
	})

	if res.RowsAffected < 1 {
		return fmt.Errorf("error delete")
	}

	return nil
}

func NewMysqlRepositoryComment(db *gorm.DB) domain.AdapterRepositoryComment {
	return &repositoryMysqlLayerComment{
		DB: db,
	}
}
