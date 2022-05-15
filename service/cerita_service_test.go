package service

import (
	"errors"
	"testing"

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUpdateStory(t *testing.T) {
	testTable := []struct{
		name string
		f func(id int, cerita model.Cerita) error
		noError bool
		token, id int
	}{
		{
			name: "success",
			f: func(id int, cerita model.Cerita) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			name: "error token != id",
			f: func(id int, cerita model.Cerita) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			name: "error internal",
			f: func(id int, cerita model.Cerita) error {
				if id == 1 {

				}
				return errors.New("error")
			},
			noError: false,
			token: 1,
			id: 1,
		},
	}
	repo := repoMockTraditionalStory{
	}
	
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.f = v.f
	
			svc := NewServiceStory(&repo, config.Config{})
			err := svc.UpdateStoryService(v.id,v.token, model.Cerita{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}

func BenchmarkUpdateTraditionalStory(t *testing.B) {
	repo := repoMockTraditionalStory{
	}

	repo.f = func(id int, cerita model.Cerita) error {
		return errors.New("error")
	}
	svc := NewServiceStory(&repo, config.Config{})

	for i := 0; i <= t.N; i++ {
		_ = svc.UpdateStoryService(1,2, model.Cerita{})
	}
}

func BenchmarkUpdateMockLibraryStory(t *testing.B) {
	repo := repoMockStory{
	}

	repo.On("UpdateOneStoryByID", mock.Anything, mock.Anything).
	Return(errors.New("error"))
	svc := NewServiceStory(&repo, config.Config{})
	for i := 0; i <= t.N; i++ {
		_ = svc.UpdateStoryService(1,2, model.Cerita{})
	}
}