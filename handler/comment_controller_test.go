package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCommentControllerAll(t *testing.T) {
	svc := MockSvcComment{}

	svc.On("CreateCommentService",  mock.Anything).
	Return(errors.New("new")).Once()

	svc.On("CreateCommentService",  mock.Anything).
	Return(nil).Once()

	commentController := EchoControllerComment{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		commentController.CreateCommentController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		commentController.CreateCommentController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}
