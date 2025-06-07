package dbmongo

import (
	"context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/akhilsomanvs/url-shortener/internal/config"
	"github.com/akhilsomanvs/url-shortener/internal/models"
	"github.com/akhilsomanvs/url-shortener/internal/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDatabse struct {
	Client *mongo.Client
	Config *config.Config
}

func InitMongoDB(cfg *config.Config) *MongoDatabse {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(fmt.Sprintf("mongodb://%s", cfg.GetStorageAddress())))
	if err != nil {
		panic("Could not connect to DB " + err.Error())
	}
	if err = client.Ping(ctx, nil); err != nil {
		panic("Could not PING to DB " + err.Error())
	}
	return &MongoDatabse{
		Client: client,
		Config: cfg,
	}
}

func (db *MongoDatabse) SaveShortUrl(shortUrl *models.ShortUrl) error {
	urlCollection := db.Client.Database(db.Config.Database.Name).Collection(db.Config.Database.CollectionName)
	_, err := urlCollection.InsertOne(context.TODO(), shortUrl)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			//The URL is already in the database. So there is no error
			return nil
		}
		return err
	}
	return nil
}

func (db *MongoDatabse) GetUniqueShortUrl(uniqueHash string, orignalUrl string) (models.ShortUrl, error) {
	urlCollection := db.Client.Database(db.Config.Database.Name).Collection(db.Config.Database.CollectionName)
	codeSlice := strings.Split(uniqueHash, "")
	startIndex := 0

	shortUrl := models.ShortUrl{
		Url:         orignalUrl,
		ShortCode:   "",
		AccessCount: 1,
	}
	for key := utils.GetHashWithKeyLength(codeSlice, startIndex); key != ""; {
		filter := bson.D{{Key: "short_code", Value: key}}
		err := urlCollection.FindOne(context.TODO(), filter).Decode(&shortUrl)
		if err != nil {
			//The short code does not exists in DB
			var createdAt = time.Now()
			// shortUrl.Id, _ = primitive.ObjectIDFromHex(key)
			shortUrl.ShortCode = key
			shortUrl.CreatedAt = createdAt
			shortUrl.UpdatedAt = createdAt
			return shortUrl, nil
		} else {
			if shortUrl.Url == orignalUrl {
				shortUrl.ShortCode = key
				return shortUrl, errors.New("Original URL exists in DB")
			} else {
				startIndex++
			}
		}
	}
	return models.ShortUrl{}, errors.New("Original URL exists in DB")
}

func (db *MongoDatabse) GetOriginalUrl(shortCode string) (models.ShortUrl, error) {
	urlCollection := db.Client.Database(db.Config.Database.Name).Collection(db.Config.Database.CollectionName)
	filter := bson.D{{Key: "short_code", Value: shortCode}}
	var shortUrl models.ShortUrl
	err := urlCollection.FindOne(context.TODO(), filter).Decode(&shortUrl)
	if err != nil {
		return models.ShortUrl{}, errors.New("Could not find URL")
	}
	return shortUrl, nil
}

func (db *MongoDatabse) UpdateShortUrl(shortUrl *models.ShortUrl) error {
	urlCollection := db.Client.Database(db.Config.Database.Name).Collection(db.Config.Database.CollectionName)
	filter := bson.D{{Key: "short_code", Value: shortUrl.ShortCode}}
	update := bson.D{
		{Key: "$set",
			Value: shortUrl,
		},
	}

	_, err := urlCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return errors.New("Could Not update " + err.Error())
	}
	err = urlCollection.FindOne(context.TODO(), filter).Decode(&shortUrl)
	if err != nil {
		return errors.New("Could Not FIND " + err.Error())
	}

	return nil
}
func (db *MongoDatabse) DeleteShortUrl(shortUrl string) error {
	urlCollection := db.Client.Database(db.Config.Database.Name).Collection(db.Config.Database.CollectionName)
	filter := bson.D{{Key: "short_code", Value: shortUrl}}
	_, err := urlCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Could not delete data " + err.Error())
	}

	return nil
}
