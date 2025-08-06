package controllers

import (
	"github.com/labstack/echo/v4"
	"testApi/services"
)

func StartSession(c echo.Context) error {
	return services.StartSession(c)
}

func BreakSession(c echo.Context) error {
	return services.BreakSession(c)
}

func EndBreak(c echo.Context) error {
	return services.EndBreak(c)
}
