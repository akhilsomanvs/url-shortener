package routes

import (
	"github.com/akhilsomanvs/url-shortener/internal/handlers"
	"github.com/akhilsomanvs/url-shortener/internal/storage/db"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine, db *db.Database) {
	server.POST("/shorten", handlers.CreateShortUrl(db))
	server.GET("/shorten/:shortURL", handlers.FetchOriginalURL(db))
	server.PUT("/shorten/:shortURL", handlers.UpdateShortURL(db))
}
