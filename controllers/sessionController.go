package controllers

import (
	"github.com/labstack/echo/v4"
	"testApi/services"
)

func StartSession(c echo.Context) error {
	return services.StartSession(c)
}
