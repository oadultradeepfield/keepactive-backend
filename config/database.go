package config

import (
	"fmt"
	"log"
	"os"

	"github.com/oadultradeepfield/keepactive-backend/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func InitDB() *gorm.DB {
    dsn := os.Getenv("DATABASE_URL")
    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
        Logger: logger.Default.LogMode(logger.Info),
        PrepareStmt: true,
    })
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Auto migrate schemas
    db.AutoMigrate(&models.User{}, &models.Website{})
    session := db.Session(&gorm.Session{PrepareStmt: true})
    if session != nil {
        fmt.Println("Migration successful")
    }

    return db
}
