package service

import (
	"MiniProject-Novellist/model"
	"github.com/stretchr/testify/mock"
)

type repoMock struct {mock.Mock}
type repoMockStory struct {mock.Mock}
type repoMockBab struct {mock.Mock}
type repoMockComment struct {mock.Mock}

func (r *repoMock) CreateUsers(user model.User) error {panic("implemente")}
func (r *repoMock) GetAll() []model.User {panic("implemente")}
func (r *repoMock) GetOneByID(id int) (user model.User, err error) {panic("implemente")}
func (r *repoMock) GetOneByEmail(email string) (user model.User, err error) {panic("implemente")}
func (r *repoMock) UpdateOneByID(id int, user model.User) error {
	ret := r.Called()
	return ret.Error(0)
}
func (r *repoMock) DeleteByID(id int) error {panic("implemente")}

func (r *repoMockStory) CreateStorys(cerita model.Cerita) error {panic("implemente")}
func (r *repoMockStory) GetAllStory() []model.Cerita {panic("implemente")}
func (r *repoMockStory) GetOneStoryByID(id int) (cerita model.Cerita, err error) {panic("implemente")}
func (r *repoMockStory) UpdateOneStoryByID(id int, cerita model.Cerita) error {
	ret := r.Called()
	return ret.Error(0)
}
func (r *repoMockStory) DeleteStoryByID(id int) error {panic("implemente")}

func (r *repoMockBab) CreateBabs(bab model.Bab) error {panic("implemente")}
func (r *repoMockBab) GetAllBab() []model.Bab {panic("implemente")}
func (r *repoMockBab) GetOneBabByID(id int) (bab model.Bab, err error) {panic("implemente")}
func (r *repoMockBab) UpdateOneBabByID(id int, bab model.Bab) error {
	ret := r.Called()
	return ret.Error(0)
}
func (r *repoMockBab) DeleteBabsByID(id int) error {panic("implemente")}

