package database

import (
	"github.com/alissonit/go-api-rest-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectToDatabase() {
	stringConnection := "host=localhost user=root password=root dbname=root port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
	DB, err = gorm.Open(postgres.Open(stringConnection))
	if err != nil {
		panic("Failed to connect to database!")
	}
	DB.AutoMigrate(&models.Student{})
}
