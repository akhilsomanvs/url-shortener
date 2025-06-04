package db

import (
	"github.com/akhilsomanvs/url-shortener/internal/config"
	"github.com/akhilsomanvs/url-shortener/internal/storage"
	dbmongo "github.com/akhilsomanvs/url-shortener/internal/storage/db/mongo"
)

type Database struct {
	Storage storage.Storage
}

func InitDB(cfg *config.Config) Database {
	storage := dbmongo.InitMongoDB(cfg)
	return Database{
		Storage: storage,
	}
}
