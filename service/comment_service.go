package service

import (

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
)

type svcKomen struct {
	c config.Config
	repo domain.AdapterRepositoryComment
}

func (s *svcKomen) CreateCommentService(komen model.Comment) error {
	return s.repo.CreateComment(komen)
}

func (s *svcKomen) GetAllCommentsService() []model.Comment {
	return s.repo.GetAllComment()
}

func (s *svcKomen) DeleteCommentsByID(id int) error {
	return s.repo.DeleteCommentByID(id)
}

func NewServiceKomen(repo domain.AdapterRepositoryComment, c config.Config) domain.AdapterServiceComment {
	return &svcKomen{
		repo: repo,
		c: c,
	}
}