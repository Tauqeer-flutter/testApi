package router

import (
	"github.com/labstack/echo/v4"
	"testApi/controllers"
)

func SetupRoutes(router *echo.Group) {
	api := router.Group("")
	{
		api.POST("/login", controllers.Login)
	}
}
