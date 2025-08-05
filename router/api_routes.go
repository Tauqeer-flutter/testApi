package router

import (
	"github.com/labstack/echo/v4"
	"testApi/controllers"
	"testApi/middlewares"
)

func SetupRoutes(router *echo.Group) {
	router.GET("/sessions/images/:fileName", controllers.GetFile)
	api := router.Group("/api/v1")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
		authGroup := api.Group("/sessions", middlewares.AuthMiddleware)
		{
			authGroup.POST("/start", controllers.StartSession)
		}
	}
}
