package dbmongo

import (
	"context"
	"fmt"
	"time"

	"github.com/akhilsomanvs/url-shortener/internal/config"
	"github.com/akhilsomanvs/url-shortener/internal/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabse struct {
	Client *mongo.Client
}

func InitMongoDB(cfg *config.Config) *MongoDatabse {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongo://%s", cfg.Addr)))
	if err != nil {
		panic("Could not connect to DB")
	}

	return &MongoDatabse{
		Client: client,
	}
}

func (db *MongoDatabse) SaveShortUrl(shortUrl models.ShortUrl) error {
	return nil
}
func (db *MongoDatabse) UpdateShortUrl(shortUrl models.ShortUrl) error  { return nil }
func (db *MongoDatabse) DeleteShortUrl(shortUrl string) error           { return nil }
func (db *MongoDatabse) GetOriginalUrl(shortUrl string) (string, error) { return "", nil }
func (db *MongoDatabse) GetShortUrlStats(shortUrl string) (models.ShortUrl, error) {
	return models.ShortUrl{}, nil
}
