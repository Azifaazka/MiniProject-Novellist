package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateUserControllerAll(t *testing.T) {
	svc := MockSvc{}

	svc.On("CreateUserService",  mock.Anything).
	Return(errors.New("new")).Once()

	svc.On("CreateUserService",  mock.Anything).
	Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.CreateUserController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteUserController(t *testing.T) {
	svc := MockSvc{}

	svc.On("DeleteByID",  mock.Anything).
	Return(errors.New("new")).Once()

	svc.On("DeleteByID",  mock.Anything).
	Return(nil).Once()

	usrController := EchoController{
		svc: &svc,
	}

	// e := echo.New()

	// t.Run("error", func(t *testing.T) {
	// 	r := httptest.NewRequest("DELETE", "/users/:id", nil)
	// 	w := httptest.NewRecorder()
	// 	echoContext := e.NewContext(r, w)

	// 	usrController.DeleteUserController(echoContext)

	// 	assert.Equal(t, 500, w.Result().StatusCode)
	// })

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/users/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		usrController.DeleteUserController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}