package database

import (
	"fmt"
	"log"
	"os"
	"final-project-go/models"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	_"github.com/lib/pq"

)

var (
	db *gorm.DB
	err error
)

func GetEnv(key string) string {
	err = godotenv.Load(".env")
	if err != nil {
		log.Fatal("error while loading .env files", err)
	}

	return os.Getenv(key)
}

func StartDB() {
	var (
		username = GetEnv("DB_USERNAME")
		password = GetEnv("DB_PASSWORD")
		dbname = GetEnv("DB_NAME")
		dbport = GetEnv("PORT")
		host = GetEnv("HOST")
	)

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, username, password, dbname, dbport)

	db, err = gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		log.Fatal("error connecting to database", err)
	}

	db.Debug().AutoMigrate(models.Comment{}, models.Photo{}, models.SocialMedia{}, models.User{})
}

func GetDB() *gorm.DB {
	return db
}


