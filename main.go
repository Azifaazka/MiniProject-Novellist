package main

import (
	"github.com/labstack/echo/v4"

	conf "MiniProject-Novellist/config"
	rest "MiniProject-Novellist/handler"
)

func main() {
	config := conf.InitConfiguration()
	e := echo.New()
	
	rest.RegisterUserGroupAPI(e, config)
	rest.RegisterStoryGroupAPI(e, config)
	rest.RegisterBabGroupAPI(e, config)
	rest.RegisterCommentGroupAPI(e, config)

	e.Logger.Fatal(e.Start(config.SERVER_ADDRESS))
}