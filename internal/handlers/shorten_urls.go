package handlers

import (
	"net/http"

	"github.com/akhilsomanvs/url-shortener/internal/models"
	"github.com/akhilsomanvs/url-shortener/internal/storage/db"
	"github.com/akhilsomanvs/url-shortener/internal/utils"
	"github.com/gin-gonic/gin"
)

func CreateShortUrl(db db.Database) func(context *gin.Context) {
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
// Update an existing short URL
// Delete an existing short URL
// Get statistics on the short URL (e.g., number of times accessed)
