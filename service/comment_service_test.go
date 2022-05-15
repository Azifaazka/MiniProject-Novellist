package service

// import (
// 	"errors"
// 	"testing"

// 	// "MiniProject-Novellist/config"
// 	"MiniProject-Novellist/model"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// func TestUpdateComment(t *testing.T) {
// 	testTable := []struct{
// 		name string
// 		f func(id int, komen model.Comment) error
// 		noError bool
// 		token, id int
// 	}{
// 		{
// 			name: "success",
// 			f: func(id int, komen model.Comment) error {
// 				return nil
// 			},
// 			noError: true,
// 			token: 1,
// 			id: 1,
// 		},
// 		{
// 			name: "error token != id",
// 			f: func(id int, komen model.Comment) error {
// 				return nil
// 			},
// 			noError: false,
// 			token: 1,
// 			id: 2,
// 		},
// 		{
// 			name: "error internal",
// 			f: func(id int, komen model.Comment) error {
// 				if id == 1 {
// 					return
// 				}
// 				return errors.New("error")
// 			},
// 			noError: false,
// 			token: 1,
// 			id: 1,
// 		},
// 	}
// 	repo := repoMockTraditionalComment{
// 	}
// }