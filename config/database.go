package config

import (
	"log"
	"os"

	"github.com/oadultradeepfield/keepactive-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto migrate schemas
    db.AutoMigrate(&models.User{}, &models.Website{})

    return db
}
