package models

import "time"

type ShortUrl struct {
	Id          int64     `bson:"_id" json:"_id"`
	Url         string    `bson:"url" json:"url"`
	ShortCode   string    `bson:"short_code" json:"short_code"`
	CreatedAt   time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt   time.Time `bson:"updated_at" json:"updated_at"`
	AccessCount int       `bson:"access_count" json:"access_count"`
}
