package configs

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	// Open a new connection to the database
	db, err := gorm.Open(
		postgres.Open("postgres://postgres:yourpassword@localhost:port/yourdatabasename"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	// Log the successful connection
	log.Println("Database connected successfully")

	DB = db
}
