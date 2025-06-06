package storage

import (
	"github.com/akhilsomanvs/url-shortener/internal/models"
)

type ShortUrlHandler interface {
	// InitDB(cfg *config.Config)
	GetUniqueShortUrl(uniqueHash string, orignalUrl string) (models.ShortUrl, error)
	SaveShortUrl(shortUrl *models.ShortUrl) error
	UpdateShortUrl(shortUrl models.ShortUrl) error
	DeleteShortUrl(shortUrl string) error
	GetOriginalUrl(shortUrl string) (models.ShortUrl, error)
	GetShortUrlStats(shortUrl string) (models.ShortUrl, error)
}

type Storage interface {
	ShortUrlHandler
}
