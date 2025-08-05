package middlewares

import (
	"github.com/labstack/echo/v4"
	"testApi/models/response"
	"testApi/utils"
	"time"
)

func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return c.JSON(401, response.BaseResponse{
				Status:  false,
				Message: "Unauthorized",
			})
		}
		tokenString := authHeader[len("Bearer "):]
		if tokenString == "" {
			return c.JSON(401, response.BaseResponse{
				Status:  false,
				Message: "Unauthorized",
			})
		}
		claims, err := utils.VerifyUserJwt(tokenString)
		if err != nil {
			return c.JSON(401, response.BaseResponse{
				Status:  false,
				Message: "Unauthorized",
			})
		}
		expirySeconds := claims["expiry"]
		expiry := time.Unix(int64(expirySeconds.(float64)), 0).UTC()
		if err != nil {
			return c.JSON(401, response.BaseResponse{
				Message: "Unauthorized: Could not parse expiry date",
			})
		}
		if expiry.Before(time.Now()) {
			return c.JSON(401, response.BaseResponse{
				Status:  false,
				Message: "Unauthorized: Token expired",
			})
		}
		c.Set("userId", claims["id"])
		c.Set("email", claims["email"])
		return next(c)
	}
}
