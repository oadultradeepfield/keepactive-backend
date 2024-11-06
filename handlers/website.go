package handlers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/oadultradeepfield/keepactive-backend/models"
	"gorm.io/gorm"
)

type WebsiteHandler struct {
    db *gorm.DB
}

func NewWebsiteHandler(db *gorm.DB) *WebsiteHandler {
    return &WebsiteHandler{db: db}
}

func (h *WebsiteHandler) Create(c *gin.Context) {
    var input struct {
        Name     string `json:"name" binding:"required"`
        URL      string `json:"url" binding:"required,url"`
        Duration int    `json:"duration" binding:"required,min=1"`
    }

    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("userID").(uint)

    website := models.Website{
        Name:       input.Name,
        URL:        input.URL,
        Duration:   input.Duration,
        Status:     "ok",
        LastPinged: time.Now(),
        UserID:     userID,
    }

    if err := h.db.Create(&website).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create website"})
        return
    }

    c.JSON(http.StatusCreated, website)
}

func (h *WebsiteHandler) List(c *gin.Context) {
    userID := c.MustGet("userID").(uint)

    var websites []models.Website
    if err := h.db.Where("user_id = ?", userID).Find(&websites).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch websites"})
        return
    }

    c.JSON(http.StatusOK, websites)
}

func (h *WebsiteHandler) Delete(c *gin.Context) {
    userID := c.MustGet("userID").(uint)
    websiteID := c.Param("id")

    result := h.db.Where("id = ? AND user_id = ?", websiteID, userID).Delete(&models.Website{})
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete website"})
        return
    }

    if result.RowsAffected == 0 {
        c.JSON(http.StatusNotFound, gin.H{"error": "Website not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Website deleted successfully"})
}
