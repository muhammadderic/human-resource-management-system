package migrate

import (
	"github.com/muhammadderic/hrms/models"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	if db == nil {
		panic("Database connection is not established")
	}

	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
}
