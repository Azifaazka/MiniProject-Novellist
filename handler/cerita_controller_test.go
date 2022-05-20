package handler

import (
	"errors"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCeritaControllerAll(t *testing.T) {
	svc := MockSvcStory{}

	svc.On("CreateStoryService",  mock.Anything).
	Return(errors.New("new")).Once()

	svc.On("CreateStoryService",  mock.Anything).
	Return(nil).Once()

	storyController := EchoControllerStory{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		storyController.CreateStoryController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("POST", "/", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		storyController.CreateStoryController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestDeleteStoryController(t *testing.T) {
	svc := MockSvcStory{}

	svc.On("DeleteStorysByID",  mock.Anything).
	Return(errors.New("new")).Once()

	svc.On("DeleteStorysByID",  mock.Anything).
	Return(nil).Once()

	storyController := EchoControllerStory{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("DELETE", "/storys/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		storyController.DeleteStoryController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("DELETE", "/storys/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		storyController.DeleteStoryController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

func TestUpdateStoryController(t *testing.T) {
	svc := MockSvcStory{}

	svc.On("UpdateOneStoryByID",  mock.Anything).
	Return(errors.New("new")).Once()

	svc.On("UpdateOneStoryByID",  mock.Anything).
	Return(nil).Once()

	storyController := EchoControllerStory{
		svc: &svc,
	}

	e := echo.New()

	t.Run("error", func(t *testing.T) {
		r := httptest.NewRequest("PUT", "/storys/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		storyController.UpdateStoryController(echoContext)

		assert.Equal(t, 500, w.Result().StatusCode)
	})

	t.Run("success", func(t *testing.T) {
		e := echo.New()
		r := httptest.NewRequest("PUT", "/storys/:id", nil)
		w := httptest.NewRecorder()
		echoContext := e.NewContext(r, w)

		storyController.UpdateStoryController(echoContext)

		assert.Equal(t, 201, w.Result().StatusCode)
	})
}

