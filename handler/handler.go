package handler 

import (

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"MiniProject-Novellist/config"
	"MiniProject-Novellist/database"
	m "MiniProject-Novellist/middleware"
	"MiniProject-Novellist/repository"
	"MiniProject-Novellist/service"
)

func RegisterUserGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewServiceUser(repo, conf)

	cont := EchoController{
		svc: svc,
	}

	apiUser := e.Group("/users", 
		middleware.Logger(),
		middleware.CORS(),
		m.APIKEYMiddleware, 
	)

	apiUser.GET("", cont.GetUsersController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSONPretty(404, map[string]interface{}{
					"messages": "token invalid",
				}, "  ")
			},
			SuccessHandler: func(c echo.Context) {
			},
		},
	))

	apiUser.POST("/login", cont.LoginUserController)
	apiUser.GET("/:id", cont.GetUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.PUT("/:id", cont.UpdateUserController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
	apiUser.DELETE("/:id", cont.DeleteUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiUser.POST("", cont.CreateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func RegisterStoryGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepositoryStory(db)

	svc := service.NewServiceStory(repo, conf)

	cont := EchoControllerStory{
		svc: svc,
	}

	apiStory := e.Group("/storys", 
		middleware.Logger(),
		middleware.CORS(),
	)

	apiStory.GET("", cont.GetStorysController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSONPretty(404, map[string]interface{}{
					"messages": "token invalid",
				}, "  ")
			},
			SuccessHandler: func(c echo.Context) {
			},
		},
	))

	apiStory.GET("/:id", cont.GetStoryController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiStory.PUT("/:id", cont.UpdateStoryController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
	apiStory.DELETE("/:id", cont.DeleteStoryController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiStory.POST("", cont.CreateStoryController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func RegisterBabGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepositoryBab(db)

	svc := service.NewServiceBab(repo, conf)

	cont := EchoControllerBab{
		svc: svc,
	}

	apiBab := e.Group("/babs", 
		middleware.Logger(),
		middleware.CORS(),
	)

	apiBab.GET("", cont.GetBabController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSONPretty(404, map[string]interface{}{
					"messages": "token invalid",
				}, "  ")
			},
			SuccessHandler: func(c echo.Context) {
			},
		},
	))

	apiBab.GET("/:id", cont.GetBabController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiBab.PUT("/:id", cont.UpdateBabController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
		},
	))
	apiBab.DELETE("/:id", cont.DeleteBabsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiBab.POST("", cont.CreateBabController, middleware.JWT([]byte(conf.JWT_KEY)))
}

func RegisterCommentGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepositoryComment(db)

	svc := service.NewServiceKomen(repo, conf)

	cont := EchoControllerComment{
		svc: svc,
	}

	apiComment := e.Group("/comments", 
		middleware.Logger(),
		middleware.CORS(),
	)

	apiComment.GET("", cont.GetCommentsController, middleware.JWTWithConfig(
		middleware.JWTConfig{
			SigningKey: []byte(conf.JWT_KEY),
			ErrorHandlerWithContext: func(err error, c echo.Context) error {
				return c.JSONPretty(404, map[string]interface{}{
					"messages": "token invalid",
				}, "  ")
			},
			SuccessHandler: func(c echo.Context) {
			},
		},
	))

	apiComment.GET("/:id", cont.GetCommentsController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiComment.DELETE("/:id", cont.DeleteCommentController, middleware.JWT([]byte(conf.JWT_KEY)))
	apiComment.POST("", cont.CreateCommentController, middleware.JWT([]byte(conf.JWT_KEY)))
}
