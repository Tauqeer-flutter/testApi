package controllers

import (
	"github.com/labstack/echo/v4"
	"io"
	"os"
	"testApi/models/response"
)

func StartSession(c echo.Context) error {
	//startTime := c.FormValue("start_time")
	file, err := c.FormFile("work_start_image")
	if err != nil {
		return c.JSON(400, response.BaseResponse{
			Status:  false,
			Message: "Invalid request, Work start image is required",
		})
	}
	src, err := file.Open()
	if err != nil {
		return c.JSON(500, response.BaseResponse{
			Status:  false,
			Message: "Something went wrong!",
		})
	}
	defer src.Close()
	dst, err := os.Create(file.Filename)
	if err != nil {
		return c.JSON(500, response.BaseResponse{
			Message: "Something went wrong!",
		})
	}
	defer dst.Close()
	_, err = os.Create(file.Filename)
	if err != nil {
		return c.JSON(500, response.BaseResponse{
			Message: "Something went wrong!",
		})
	}
	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(500, response.BaseResponse{
			Message: "Something went wrong!",
		})
	}
	return c.JSON(200, response.BaseResponse{
		Status:  true,
		Message: "Success",
	})
}
