package handler
import (
	"net/http"
	"strconv"

	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type EchoControllerStory struct {
	svc domain.AdapterServiceStory
}

func (ce *EchoControllerStory) CreateStoryController(c echo.Context) error {
	cerita := model.Cerita{}
	c.Bind(&cerita)
	
	err := ce.svc.CreateStoryService(cerita)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"Storys": cerita,
	})
}

func (ce *EchoControllerStory) UpdateStoryController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	cerita := model.Cerita{}
	c.Bind(&cerita)

	bearer := c.Get("cerita").(*jwt.Token)
	claim := bearer.Claims.(jwt.MapClaims)

	err := ce.svc.UpdateStoryService(intID, int(claim["id"].(float64)), cerita)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change or unauthorization",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
		"id": intID,
		"expeted id": claim["id"],
	})
}

func (ce *EchoControllerStory) DeleteStoryController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteStorysByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerStory) GetStoryController(c echo.Context) error {
	id := c.Param("id")
	intID, err := strconv.Atoi(id)
	
	res, err := ce.svc.GetStoryByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"storys": res,
	})
}

func (ce *EchoControllerStory) GetStorysController(c echo.Context) error {
	cerita := ce.svc.GetAllStorysService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"storys": cerita,
	}, "  ")
}

