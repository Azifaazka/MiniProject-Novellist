package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	m "github.com/azifaazka/MiniProject-Novellist/middleware"
)

func RegisterUserGroupAPI(e *echo.Group) {
	e.GET("", getUsersController, m.CustomMiddleware)
	e.GET("/:id", getUserController, m.CustomMiddleware)
	e.PUT("/:id", updateUserController)
	e.DELETE("/:id", deleteUserController)
	e.POST("", createUserController, middleware.BasicAuth(m.AuthUser))
}