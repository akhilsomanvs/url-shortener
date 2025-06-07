package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/akhilsomanvs/url-shortener/internal/models"
	"github.com/akhilsomanvs/url-shortener/internal/storage/db"
	"github.com/akhilsomanvs/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateShortUrl(db *db.Database) func(context *gin.Context) {
	return func(context *gin.Context) {
		var url models.Url
		err := context.ShouldBind(&url)
		if err != nil {
			context.JSON(http.StatusBadRequest, models.NewApiResponseModel("Bad Request", err.Error()))
			return
		}

		// uniqueKey := utils.GenerateUniqueKey()
		uniqueKey := utils.GenerateHashKey(url.Url)
		//Check if it is already in the DB
		shortCode, err := db.Storage.GetUniqueShortUrl(uniqueKey, url.Url)
		if err != nil {
			//Original URL already exists in DB
			if shortCode.ShortCode == "" {
				context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Could not shorten URL", err.Error()))
				return
			} else {
				context.JSON(http.StatusOK, models.NewApiResponseModel("Success", shortCode))
				return
			}
		}
		//If there are no error
		//Store the newly created shorturl into DB
		err = db.Storage.SaveShortUrl(&shortCode)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Could save short url to DB", err.Error()))
			return
		}
		context.JSON(http.StatusCreated, models.NewApiResponseModel("Success", shortCode))
	}
}

// Retrieve an original URL from a short URL
func FetchOriginalURL(db *db.Database) func(context *gin.Context) {
	return func(context *gin.Context) {
		shortCode := context.Param("shortURL")
		if shortCode == "" {
			context.JSON(http.StatusBadRequest, models.NewApiResponseModel("Bad Request", "Could not parse event ID"))
			return
		}

		originalUrl, err := db.Storage.GetOriginalUrl(shortCode)
		if err != nil {
			context.JSON(http.StatusOK, models.NewApiResponseModel("Short Code Not Found", "This short code is not associated with any data"))
			return
		}

		data, err := json.MarshalIndent(originalUrl, "", "")
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Failed", "Failed to fetch data"))
			return
		}

		var i interface{}
		if err := json.Unmarshal([]byte(data), &i); err != nil {
			context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Failed", "Failed to fetch data"))
			return
		}

		if m, ok := i.(map[string]interface{}); ok {
			delete(m, "access_count") // No Need to show the access count
		}
		context.JSON(http.StatusOK, models.NewApiResponseModel("Success", i))

	}
}

// Update an existing short URL
func UpdateShortURL(db *db.Database) func(context *gin.Context) {
	return func(context *gin.Context) {
		shortCode := context.Param("shortURL")
		if shortCode == "" {
			context.JSON(http.StatusBadRequest, models.NewApiResponseModel("Bad Request", "Could not parse event ID"))
			return
		}

		var url models.Url
		err := context.ShouldBind(&url)
		if err != nil {
			context.JSON(http.StatusBadRequest, models.NewApiResponseModel("Bad Request", err.Error()))
			return
		}

		//Fetch the original URL from the DB
		originalUrl, err := db.Storage.GetOriginalUrl(shortCode)
		if err != nil {
			context.JSON(http.StatusOK, models.NewApiResponseModel("Short Code Not Found", "This short code is not associated with any data"))
			return
		}

		originalUrl.Url = url.Url
		err = db.Storage.UpdateShortUrl(&originalUrl)

		if err != nil {
			fmt.Println(err)
			context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Failed", "Failed to update data"))
			return
		}

		context.JSON(http.StatusOK, models.NewApiResponseModel("Success", originalUrl))
	}
}

// Delete an existing short URL
// Get statistics on the short URL (e.g., number of times accessed)
