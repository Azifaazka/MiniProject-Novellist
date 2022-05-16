package service

import (
	"errors"
	"testing"

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/model"
	"github.com/stretchr/testify/assert"

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
			err := svc.UpdateStoryService(v.id, model.Cerita{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}
