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
	var foundUser dtos.User
	err = config.DB.Raw("SELECT * FROM users WHERE email = ?", request.Email).Scan(&foundUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
	}
	log.Debug("Found user: ", foundUser)
	if foundUser.Id == 0 {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Invalid credentials",
		})
	} else if !utils.CheckPasswordHash(request.Password, foundUser.Password) {
		return c.JSON(http.StatusBadRequest, response.BaseResponse{
			Message: "Invalid credentials",
		})
	}
	token, err := utils.GenerateUserJwt(foundUser.Email, foundUser.Id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: "Something went wrong!",
		})
	}
	return c.JSON(http.StatusOK, response.SuccessAuthResponse{
		Status:  true,
		Message: "Login successful",
		Token:   token,
		User: response.UserData{
			Id:         foundUser.Id,
			FirstName:  foundUser.FirstName,
			LastName:   foundUser.LastName,
			Email:      foundUser.Email,
			Age:        foundUser.Age,
			IsVerified: foundUser.IsVerified,
			CreatedAt:  foundUser.CreatedAt,
			UpdatedAt:  foundUser.UpdatedAt,
		},
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
	token, err := utils.GenerateUserJwt(user.Email, user.Id)
	if err != nil {
		log.Error(err)
		return c.JSON(http.StatusInternalServerError, response.BaseResponse{
			Message: err.Error(),
		})
	}
	return c.JSON(http.StatusOK, response.SuccessAuthResponse{
		Status:  true,
		Message: "Registration successful",
		Token:   token,
		User: response.UserData{
			Id:         user.Id,
			FirstName:  user.FirstName,
			LastName:   user.LastName,
			Email:      user.Email,
			Age:        user.Age,
			IsVerified: user.IsVerified,
			CreatedAt:  user.CreatedAt,
			UpdatedAt:  user.UpdatedAt,
		},
	})
}
