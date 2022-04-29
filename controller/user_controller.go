package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/azifaazka/MiniProject-Novellist/model"
	"github.com/azifaazka/MiniProject-Novellist/repository"
	"github.com/labstack/echo/v4"
)

func createUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	
	err := repository.CreateUsers(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"users": user,
	})
}

func updateUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	user := model.User{}
	c.Bind(&user)

	err := repository.UpdateOneByID(intID, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func deleteUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := repository.DeleteByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func getUserController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)
	
	res, err := repository.GetOneByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users": res,
	})
}

func getUsersController(c echo.Context) error {
	users := repository.GetAll()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"users": users,
	}, "  ")
}