package model

import "time"

type Url struct {
	Id           string     `json:"id"`
	OriginalUrl  string     `json:"original_url"`
	ShortenedUrl string     `json:"shortened_url"`
	CreateAt     *time.Time `json:"created_at"`
	ExpiredAt    *time.Time `json:"expired_at"`
}
