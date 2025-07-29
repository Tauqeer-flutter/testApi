package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"testApi/config"
	"testApi/router"
	"testApi/utils"
)

func main() {
	err := config.ConnectDB()
	if err != nil {
		return
	}
	e := echo.New()
	e.Binder = &utils.CustomBinder{}
	e.Validator = &utils.CustomValidator{
		Validator: validator.New(validator.WithRequiredStructEnabled()),
	}

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router.RequestHandler(e)
}
