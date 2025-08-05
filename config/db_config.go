package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testApi/models/dtos"
)

var JwtSecret []byte
var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	JwtSecret = []byte(os.Getenv("JWT_SECRET"))
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"
	fmt.Println("DSN: ", dsn)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
		return err
	}

	err = DB.AutoMigrate(&dtos.User{}, &dtos.Session{}, &dtos.Break{})
	if err != nil {
		log.Fatal("Error auto-migrating database: " + err.Error())
		return err
	}

	return nil
}
