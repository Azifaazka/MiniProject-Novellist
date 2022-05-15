package handler

import (
	"MiniProject-Novellist/model"
	"github.com/stretchr/testify/mock"
)

type MockSvc struct {mock.Mock}
type MockSvcStory struct {mock.Mock}
type MockSvcBab struct {mock.Mock}
type MockSvcComment struct {mock.Mock}

func (m *MockSvc) CreateUserService(user model.User) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvc) UpdateUserService(id, idToken int, user model.User) error {return nil}
func (m *MockSvc) GetAllUsersService() []model.User {return []model.User{}}
func (m *MockSvc) GetUserByID(id int) (model.User, error) {return model.User{}, nil}
func (m *MockSvc) LoginUser(email, password string) (string, int) {return "success", 200}
func (m *MockSvc) DeleteByID(id int) error {return nil}

func (m *MockSvcStory) CreateStoryService(cerita model.Cerita) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvcStory) UpdateStoryService(id, idToken int, cerita model.Cerita) error {return nil}
func (m *MockSvcStory) GetAllStorysService() []model.Cerita {return []model.Cerita{}}
func (m *MockSvcStory) GetStoryByID(id int) (model.Cerita, error) {return model.Cerita{}, nil}
func (m *MockSvcStory) DeleteStorysByID(id int) error {return nil}

func (m *MockSvcBab) CreateBabService(bab model.Bab) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvcBab) UpdateBabService(id, idToken int, bab model.Bab) error {return nil}
func (m *MockSvcBab) GetAllBabsService() []model.Bab {return []model.Bab{}}
func (m *MockSvcBab) GetBabByID(id int) (model.Bab, error) {return model.Bab{}, nil}
func (m *MockSvcBab) DeleteBabsByID(id int) error {return nil}

func (m *MockSvcComment) CreateCommentService(komen model.Comment) error {
	ret := m.Called()
	return ret.Error(0)
}
func (m *MockSvcComment) GetAllCommentsService() []model.Comment {return []model.Comment{}}
func (m *MockSvcComment) DeleteCommentsByID(id int) error {return nil}