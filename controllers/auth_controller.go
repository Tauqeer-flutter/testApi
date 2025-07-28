package controllers

import (
	"github.com/labstack/echo/v4"
	"testApi/services"
)

func Login(c echo.Context) error {
	return services.Login(c)
}
