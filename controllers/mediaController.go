package controllers

import (
	"github.com/labstack/echo/v4"
	"testApi/models/response"
)

func GetFile(c echo.Context) error {
	fileName := c.Param("fileName")
	if fileName == "" {
		return c.JSON(400, response.BaseResponse{
			Message: "Invalid request, File name is required",
		})
	}
	return c.File("images/" + fileName)
}
