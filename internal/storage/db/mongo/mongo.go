package dbmongo

import (
	"context"
	"errors"
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
}

func InitMongoDB(cfg *config.Config) *MongoDatabse {
	ctx, cancel := context.WithTimeout(context.TODO(), 2*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic("Could not connect to DB " + err.Error())
	}
	if err = client.Ping(ctx, nil); err != nil {
		panic("Could not PING to DB " + err.Error())
	}
	return &MongoDatabse{
		Client: client,
	}
}

func (db *MongoDatabse) SaveShortUrl(shortUrl *models.ShortUrl) error {
	urlCollection := db.Client.Database("AppDatabase").Collection("ShortURL")
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
	urlCollection := db.Client.Database("AppDatabase").Collection("ShortURL")
	codeSlice := strings.Split(uniqueHash, "")
	startIndex := 0

	shortUrl := models.ShortUrl{
		Url:         orignalUrl,
		ShortCode:   "",
		AccessCount: 1,
	}
	for key := utils.GetHashWithKeyLength(codeSlice, startIndex); key != ""; {
		filter := bson.D{{"short_code", key}}
		err := urlCollection.FindOne(context.TODO(), filter).Decode(&shortUrl)
		if err != nil {
			//The short code does not exists in DB
			var createdAt = time.Now()
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
	urlCollection := db.Client.Database("AppDatabase").Collection("ShortURL")
	filter := bson.D{{Key: "short_code", Value: shortCode}}
	var shortUrl models.ShortUrl
	err := urlCollection.FindOne(context.TODO(), filter).Decode(&shortUrl)
	if err != nil {
		return models.ShortUrl{}, errors.New("Could not find URL")
	}
	return shortUrl, nil
}

func (db *MongoDatabse) UpdateShortUrl(shortUrl *models.ShortUrl) error {
	urlCollection := db.Client.Database("AppDatabase").Collection("ShortURL")
	filter := bson.D{{Key: "short_code", Value: shortUrl.ShortCode}}
	update := bson.D{
		{Key: "$set",
			Value: bson.D{
				{Key: "url", Value: shortUrl.Url},
				{Key: "update_at", Value: time.Now()},
			},
		},
	}

	_, err := urlCollection.UpdateOne(context.Background(), filter, update)
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
	urlCollection := db.Client.Database("AppDatabase").Collection("ShortURL")
	filter := bson.D{{Key: "short_code", Value: shortUrl}}
	_, err := urlCollection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return errors.New("Could not delete data " + err.Error())
	}

	return nil
}
func (db *MongoDatabse) GetShortUrlStats(shortUrl string) (models.ShortUrl, error) {
	return models.ShortUrl{}, nil
}
