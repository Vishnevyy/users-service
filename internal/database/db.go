package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/Vishnevyy/users-service/internal/user"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to init DB: %v", err)
	}

	if err := DB.AutoMigrate(&user.User{}); err != nil {
		log.Fatalf("auto-migrate failed: %v", err)
	}
}
