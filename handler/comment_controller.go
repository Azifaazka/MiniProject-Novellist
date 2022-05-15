package handler

import (
	"net/http"
	"strconv"

	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
	"github.com/labstack/echo/v4"
)

type EchoControllerComment struct {
	svc domain.AdapterServiceComment
}

func (ce *EchoControllerComment) CreateCommentController(c echo.Context) error {
	komen := model.Comment{}
	c.Bind(&komen)
	
	err := ce.svc.CreateCommentService(komen)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"comment": komen,
	})
}


func (ce *EchoControllerComment) DeleteCommentController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteCommentsByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerComment) GetCommentsController(c echo.Context) error {
	komen := ce.svc.GetAllCommentsService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"komens": komen,
	}, "  ")
}
