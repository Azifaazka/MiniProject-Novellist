package service

import (

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
)

type svcStory struct {
	c config.Config
	repo domain.AdapterRepositoryStory
}

func (s *svcStory) CreateStoryService(cerita model.Cerita) error {
	return s.repo.CreateStorys(cerita)
}

func (s *svcStory) UpdateStoryService(id int, cerita model.Cerita) error {
	return s.repo.UpdateOneStoryByID(id, cerita)
}

func (s *svcStory) GetAllStorysService() []model.Cerita {
	return s.repo.GetAllStory()
}

func (s *svcStory) GetStoryByID(id int) (model.Cerita, error) {
	return s.repo.GetOneStoryByID(id)
}

func (s *svcStory) DeleteStorysByID(id int) error {
	return s.repo.DeleteStoryByID(id)
}

func NewServiceStory(repo domain.AdapterRepositoryStory, c config.Config) domain.AdapterServiceStory {
	return &svcStory{
		repo: repo,
		c: c,
	}
}