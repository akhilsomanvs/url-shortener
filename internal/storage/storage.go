package storage

import "github.com/akhilsomanvs/url-shortener/internal/models"

type ShortUrlHandler interface {
	// InitDB(cfg *config.Config)
	SaveShortUrl(shortUrl models.ShortUrl) error
	UpdateShortUrl(shortUrl models.ShortUrl) error
	DeleteShortUrl(shortUrl string) error
	GetOriginalUrl(shortUrl string) (string, error)
	GetShortUrlStats(shortUrl string) (models.ShortUrl, error)
}

type Storage interface {
	ShortUrlHandler
}
