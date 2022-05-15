package handler

import (
	"MiniProject-Novellist/model"
)

type MockSvcTradition struct {
	MFunc func (email, password string) (string, int)
}

func (m *MockSvcTradition) SetFuncMock(f func (email, password string) (string, int)) {m.MFunc = f}
func (m *MockSvcTradition) CreateUserService(user model.User) error {return nil}
func (m *MockSvcTradition) UpdateUserService(id, idToken int, user model.User) error {return nil}
func (m *MockSvcTradition) GetAllUsersService() []model.User {return []model.User{}}
func (m *MockSvcTradition) GetUserByID(id int) (model.User, error) {return model.User{}, nil}
func (m *MockSvcTradition) LoginUser(email, password string) (string, int) {return m.MFunc(email, password)}
func (m *MockSvcTradition) DeleteByID(id int) error {return nil}
