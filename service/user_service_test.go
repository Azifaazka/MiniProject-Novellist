package service

import (
	"errors"
	"testing"

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdate(t *testing.T) {
	testTable := []struct{
		name string
		f func(id int, user model.User) error
		noError bool
		token, id int
	}{
		{
			name: "success",
			f: func(id int, user model.User) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			name: "error token != id",
			f: func(id int, user model.User) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			name: "error internal",
			f: func(id int, user model.User) error {
				if id == 1 {

				}
				return errors.New("error")
			},
			noError: false,
			token: 1,
			id: 1,
		},
	}
	repo := repoMockTraditional{
	}
	
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.f = v.f
	
			svc := NewServiceUser(&repo, config.Config{})
			err := svc.UpdateUserService(v.id,v.token, model.User{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkUpdateTraditional(t *testing.B) {
	repo := repoMockTraditional{
	}

	repo.f = func(id int, user model.User) error {
		return errors.New("error")
	}
	svc := NewServiceUser(&repo, config.Config{})

	for i := 0; i <= t.N; i++ {
		_ = svc.UpdateUserService(1,2, model.User{})
	}
}

func BenchmarkUpdateMockLibrary(t *testing.B) {
	repo := repoMock{
	}

	repo.On("UpdateOneByID", mock.Anything, mock.Anything).
	Return(errors.New("error"))
	svc := NewServiceUser(&repo, config.Config{})
	for i := 0; i <= t.N; i++ {
		_ = svc.UpdateUserService(1,2, model.User{})
	}
}