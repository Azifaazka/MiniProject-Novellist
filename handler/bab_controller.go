package handler
import (
	"net/http"
	"strconv"

	"MiniProject-Novellist/domain"
	"MiniProject-Novellist/model"
	"github.com/labstack/echo/v4"
)

type EchoControllerBab struct {
	svc domain.AdapterServiceBab
}

func (ce *EchoControllerBab) CreateBabController(c echo.Context) error {
	bab := model.Bab{}
	c.Bind(&bab)
	
	err := ce.svc.CreateBabService(bab)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"messages": err.Error(),
		})
	}
	
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"messages": "success",
		"babs": bab,
	})
}

func (ce *EchoControllerBab) UpdateBabController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	bab := model.Bab{}
	c.Bind(&bab)

	err := ce.svc.UpdateBabService(intID, bab)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no change",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "edited",
	})
}

func (ce *EchoControllerBab) DeleteBabsController(c echo.Context) error {
	id := c.Param("id")
	intID, _ := strconv.Atoi(id)

	err := ce.svc.DeleteBabsByID(intID)
	if err != nil {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"messages": "no id or no delete",
		})
	}

	return c.JSON(http.StatusNoContent, map[string]interface{}{
		"messages": "deleted",
	})
}

func (ce *EchoControllerBab) GetBabController(c echo.Context) error {
	babs := ce.svc.GetAllBabsService()

	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"babs": babs,
	})
}

func (ce *EchoControllerBab) GetBabsController(c echo.Context) error {
	babs := ce.svc.GetAllBabsService()

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"babs": babs,
	}, "  ")
}
