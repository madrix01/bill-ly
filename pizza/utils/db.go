package utils

import (
	"bill-ly/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// host -> localhost
// user -> bill-ly
// password -> bill-ly-6969
// db -> bill-db
// port -> 5432

func InitDB() *gorm.DB {
	dbURL := "postgres://bill-ly:bill-ly-6969@db:5432/bill-db"
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		log.Fatal("Unable to connect to db")
	}

	db.AutoMigrate(&models.User{})

	return db
}
