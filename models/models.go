package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
    gorm.Model
    Email    string `gorm:"unique"`
    Password string
    Websites []Website
}

type Website struct {
    gorm.Model
    Name       string
    URL        string
    Duration   int    // in days
    Status     string // "ok" or "failed"
    LastPinged time.Time
    UserID     uint
}
