package models

import "time"

type ShortUrl struct {
	Id          int64     `json:"id"`
	Url         string    `json:"url"`
	ShortCode   string    `json:"shortCode"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	AccessCount int       `json:"accessCount"`
}
