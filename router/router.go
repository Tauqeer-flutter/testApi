package router

import "github.com/labstack/echo/v4"

func RequestHandler(e *echo.Echo) {
	router := e.Group("/api/v1")
	SetupRoutes(router)
	e.Logger.Fatal(e.Start(":8080"))
}
