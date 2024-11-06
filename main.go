package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/oadultradeepfield/keepactive-backend/config"
	"github.com/oadultradeepfield/keepactive-backend/handlers"
	"github.com/oadultradeepfield/keepactive-backend/middleware"
	"github.com/oadultradeepfield/keepactive-backend/services"
)

func main() {
    // Load .env only in development
    if gin.Mode() != gin.ReleaseMode {
        if err := godotenv.Load(); err != nil {
            log.Println("Warning: .env file not found")
        }
    }

    // Load configuration
    cfg := config.LoadConfig()

    // Set Gin mode
    if cfg.Environment == "production" {
        gin.SetMode(gin.ReleaseMode)
    }

    db := config.InitDB()

    // Initialize handlers
    authHandler := handlers.NewAuthHandler(db)
    websiteHandler := handlers.NewWebsiteHandler(db)

    // Initialize and start website pinger
    pinger := services.NewWebsitePinger(db)
    go pinger.Start()

    // Initialize Gin router
    r := gin.New()
    
    // Use logger and recovery middleware
    r.Use(gin.Logger())
    r.Use(gin.Recovery())
    
    // Add CORS middleware
    r.Use(middleware.CorsMiddleware(cfg.AllowedOrigins))

    // Health check endpoint
    r.GET("/health", func(c *gin.Context) {
        c.JSON(200, gin.H{"status": "ok"})
    })

    // Auth routes
    r.POST("/api/register", authHandler.Register)
    r.POST("/api/login", authHandler.Login)

    // Protected routes
    authorized := r.Group("/api")
    authorized.Use(middleware.AuthMiddleware())
    {
        authorized.POST("/websites", websiteHandler.Create)
        authorized.GET("/websites", websiteHandler.List)
        authorized.DELETE("/websites/:id", websiteHandler.Delete)
    }

    r.Run(":" + cfg.Port)
}