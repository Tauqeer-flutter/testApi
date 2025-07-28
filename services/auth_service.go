package services

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"testApi/models/requests"
	"testApi/models/response"
)

func Login(c echo.Context) error {
	var request requests.LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Invalid request",
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Status:  true,
		Message: "Login successful",
	})
}
