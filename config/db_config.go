package config

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"testApi/models/dtos"
)

var DB *gorm.DB

func ConnectDB() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		return err
	}
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dsn := dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: " + err.Error())
		return err
	}

	err = DB.AutoMigrate(&dtos.User{})
	if err != nil {
		log.Fatal("Error auto-migrating database: " + err.Error())
		return err
	}

	return nil
}
