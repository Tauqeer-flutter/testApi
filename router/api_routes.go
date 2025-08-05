package router

import (
	"github.com/labstack/echo/v4"
	"testApi/controllers"
	"testApi/middlewares"
)

func SetupRoutes(router *echo.Group) {
	api := router.Group("")
	{
		api.POST("/login", controllers.Login)
		api.POST("/register", controllers.Register)
		authGroup := api.Group("/sessions", middlewares.AuthMiddleware)
		{
			authGroup.POST("/start", controllers.StartSession)
		}
	}
}
