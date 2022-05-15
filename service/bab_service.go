package service

import (
	"fmt"

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
)

type svcBab struct {
	c config.Config
	repo domain.AdapterRepositoryBab
}

func (s *svcBab) CreateBabService(bab model.Bab) error {
	return s.repo.CreateBabs(bab)
}

func (s *svcBab) UpdateBabService(id, idToken int, bab model.Bab) error {
	if id != idToken {
		return fmt.Errorf("error")
	}
	return s.repo.UpdateOneBabByID(id, bab)
}

func (s *svcBab) GetAllBabsService() []model.Bab {
	return s.repo.GetAllBab()
}

func (s *svcBab) GetBabByID(id int) (model.Bab, error) {
	return s.repo.GetOneBabByID(id)
}

func (s *svcBab) DeleteBabsByID(id int) error {
	return s.repo.DeleteBabsByID(id)
}

func NewServiceBab(repo domain.AdapterRepositoryBab, c config.Config) domain.AdapterServiceBab {
	return &svcBab{
		repo: repo,
		c: c,
	}
}