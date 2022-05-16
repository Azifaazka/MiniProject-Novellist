package service

import (
	"errors"
	"testing"

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/model"
	"github.com/stretchr/testify/assert"
)

func TestUpdateBab(t *testing.T) {
	testTable := []struct{
		name string
		f func(id int, bab model.Bab) error
		noError bool
		token, id int
	}{
		{
			name: "success",
			f: func(id int, bab model.Bab) error {
				return nil
			},
			noError: true,
			token: 1,
			id: 1,
		},
		{
			name: "error token != id",
			f: func(id int, bab model.Bab) error {
				return nil
			},
			noError: false,
			token: 1,
			id: 2,
		},
		{
			name: "error internal",
			f: func(id int, bab model.Bab) error {
				if id == 1 {

				}
				return errors.New("error")
			},
			noError: false,
			token: 1,
			id: 1,
		},
	}
	repo := repoMockTraditionalBab{
	}
	
	for _, v := range testTable {
		t.Run(v.name, func(t *testing.T) {
			repo.f = v.f
	
			svc := NewServiceBab(&repo, config.Config{})
			err := svc.UpdateBabService(v.id, model.Bab{})
			if v.noError {
				assert.NoError(t, err)
			}
		})
	}
}