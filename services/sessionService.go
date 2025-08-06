package services

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"testApi/config"
	"testApi/models/dtos"
	"testApi/models/requests"
	"testApi/models/response"
	"testApi/utils"
	"time"
)

func StartSession(c echo.Context) error {
	start := c.FormValue("start_time")
	fmt.Println("Start time: ", start)
	startTime, err := time.Parse(time.RFC3339, start)
	if err != nil {
		return c.JSON(400, response.BaseResponse{
			Status:  false,
			Message: err.Error(), //"Invalid request, Start time is required",
		})
	}
	file, err := c.FormFile("work_start_image")
	if err != nil {
		return c.JSON(400, response.BaseResponse{
			Status:  false,
			Message: "Invalid request, Work start image is required",
		})
	}
	fileName, err := utils.SaveImage(file)
	if err != nil {
		return c.JSON(500, response.BaseResponse{
			Status:  false,
			Message: "Something went wrong while uploading the image",
		})
	}
	session := dtos.Session{
		StartTime:         startTime.UTC(),
		WorkStartFilePath: fileName,
		UserId:            uint(c.Get("userId").(float64)),
		Mode:              dtos.Working,
	}
	value := config.DB.Create(&session)
	if value.Error != nil {
		return c.JSON(500, response.BaseResponse{
			Status:  false,
			Message: "Something went wrong while creating the session",
		})
	}
	return c.JSON(200, response.SessionCreatedResponse{
		Status:  true,
		Message: "Session created successfully",
		Session: session,
	})
}

func BreakSession(c echo.Context) error {
	var request requests.BreakSessionRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(400, response.BaseResponse{
			Message: err.Error(),
		})
	}
	err := config.DB.Exec("UPDATE test_db.sessions SET work_duration = ?, break_duration = ?, extra_duration = ?, mode = ? WHERE id = ?;", request.WorkDuration, request.BreakDuration, request.ExtraDuration, dtos.OnBreak, request.SessionId).Error
	if err != nil {
		return c.JSON(500, response.BaseResponse{
			Message: err.Error(), //"Something went wrong while updating the session",
		})
	}
	return c.JSON(200, response.BaseResponse{
		Status:  true,
		Message: "Session updated successfully",
	})
}
