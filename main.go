package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	conf "github.com/azifaazka/MiniProject-Novellist/config"
	c "github.com/azifaazka/MiniProject-Novellist/controller"
	db "github.com/azifaazka/MiniProject-Novellist/database"
	m "github.com/azifaazka/MiniProject-Novellist/middleware"
)

func init() {
	config := conf.InitConfiguration()
	db.InitDB(config)
}

func main() {
	e := echo.New()
	apiUser := e.Group("/users", 
		middleware.Logger(),
		middleware.CORS(),
		m.APIKEYMiddleware, 
	)

	c.RegisterUserGroupAPI(apiUser)

	e.Logger.Fatal(e.Start(":8080"))
}