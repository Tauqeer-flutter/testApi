package services

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"net/http"
	"testApi/config"
	"testApi/models/dtos"
	"testApi/models/requests"
	"testApi/models/response"
	"testApi/utils"
)

func Login(c echo.Context) error {
	var request requests.LoginRequest
	err := c.Bind(&request)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Status:  true,
		Message: "Login successful",
	})
}

func Register(c echo.Context) error {
	var user dtos.User
	err := c.Bind(&user)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: err.Error(),
		})
	}
	var foundUsers []dtos.User
	err = config.DB.Raw("SELECT * FROM users WHERE email = ?", user.Email).Scan(&foundUsers).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
	}
	if len(foundUsers) > 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "User already exists",
		})
	}
	passwordHash, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "Something went wrong!",
		})
	}
	user.Password = passwordHash
	result := config.DB.Create(&user)
	log.Debug("Result: ", result)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: result.Error.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.BaseResponse{
		Status:  true,
		Message: "User created successfully",
	})
}
