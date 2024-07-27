package configs

import (
	"log"

	"github.com/muhammadderic/hrms/migrate"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DBResult struct {
	Result *gorm.DB
	Error  error
}

func ConnectDB() {
	// Open a new connection to the database
	db, err := gorm.Open(
		postgres.Open("postgres://postgres:yourpassword@localhost:port/yourdatabasename"),
		&gorm.Config{},
	)
	if err != nil {
		panic(err)
	}

	// Run the migrations
	migrate.Migrate(db)

	// Log the successful connection
	log.Println("Database connected successfully")

	DB = db
}
