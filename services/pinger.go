package services

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/oadultradeepfield/keepactive-backend/models"
	"gorm.io/gorm"
)

type WebsitePinger struct {
    db     *gorm.DB
    client *http.Client
}

func NewWebsitePinger(db *gorm.DB) *WebsitePinger {
    return &WebsitePinger{
        db: db,
        client: &http.Client{
            Timeout: 10 * time.Second,
        },
    }
}

func (p *WebsitePinger) Start() {
    for {
        var websites []models.Website
        p.db.Find(&websites)

        for _, website := range websites {
            timeSinceLastPing := time.Since(website.LastPinged)

            var randomInterval time.Duration
            if website.Duration > 1 {
                randomInterval = time.Duration(rand.Intn(website.Duration-1)+1) * 24 * time.Hour
            } else {
                randomInterval = 24 * time.Hour
            }

            if timeSinceLastPing >= randomInterval {
                resp, err := p.client.Get(website.URL)
                
                status := "ok"
                if err != nil || resp.StatusCode >= 400 {
                    status = "failed"
                } else {
                    resp.Body.Close()
                }

                p.db.Model(&website).Updates(models.Website{
                    Status:     status,
                    LastPinged: time.Now(),
                })
            }
        }

        time.Sleep(1 * time.Hour)
    }
}
