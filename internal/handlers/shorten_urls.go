package handlers

import (
	"net/http"
	"time"

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

		uniqueKey := utils.GenerateUniqueKey()
		//Store this into DB
		var createdAt = time.Now()
		shortUrl := models.ShortUrl{
			Id:          0,
			Url:         url.Url,
			ShortCode:   uniqueKey,
			CreatedAt:   createdAt,
			UpdatedAt:   createdAt,
			AccessCount: 1,
		}
		err = db.Storage.SaveShortUrl(&shortUrl)
		if err != nil {
			context.JSON(http.StatusInternalServerError, models.NewApiResponseModel("Could save short url to DB", err.Error()))
			return
		}
		context.JSON(http.StatusOK, models.NewApiResponseModel("Success", shortUrl))
	}
}

// Retrieve an original URL from a short URL
// Update an existing short URL
// Delete an existing short URL
// Get statistics on the short URL (e.g., number of times accessed)
