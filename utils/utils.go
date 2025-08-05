package utils

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"io"
	"mime/multipart"
	"os"
	"strconv"
	"strings"
	"testApi/config"
	"time"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateUserJwt(email string, userId uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"id":     userId,
		"expiry": time.Now().Add(time.Hour * 1).UTC().Unix(),
	})
	return token.SignedString(config.JwtSecret)
}

func VerifyUserJwt(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, echo.NewHTTPError(401, "Invalid token")
		}
		return config.JwtSecret, nil
	})
	if err != nil || !token.Valid {
		return nil, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, echo.NewHTTPError(401, "Invalid token")
	}
	return claims, nil
}

func SaveImage(file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	err = createDirectoryIfNotExists("images/")
	if err != nil {
		return "", err
	}
	ext := strings.Split(file.Filename, ".")[1]

	now := time.Now().UTC().Unix()
	fileName := strconv.FormatInt(now, 10) + "." + ext
	dst, err := os.Create("images/" + fileName)
	if err != nil {
		return "", err
	}
	defer dst.Close()
	_, err = io.Copy(dst, src)
	if err != nil {
		return "", err
	}
	return fileName, nil
}

func createDirectoryIfNotExists(path string) error {
	_, err := os.ReadDir(path)
	if err == nil {
		return nil
	}
	return os.MkdirAll(path, os.ModePerm)
}
