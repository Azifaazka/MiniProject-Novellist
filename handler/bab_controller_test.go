package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateBabControllerAll(t *testing.T) {
	svc := MockSvcBab{}

	svc.On("CreateBabService",  mock.Anything).
	Return(errors.New("new")).Once()

	svc.On("CreateBabService",  mock.Anything).
	Return(nil).Once()

	babController := EchoControllerBab{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		babController.CreateBabController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		babController.CreateBabController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}
