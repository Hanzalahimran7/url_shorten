package model

import (
	"time"

	"github.com/google/uuid"
)

type Url struct {
	Id           uuid.UUID  `json:"id"`
	OriginalUrl  string     `json:"original_url"`
	ShortenedUrl string     `json:"shortened_url"`
	CreateAt     *time.Time `json:"created_at"`
	ExpiredAt    *time.Time `json:"expired_at"`
}

type Stats struct {
	Id           uuid.UUID `json:"id"`
	OriginalUrl  string    `json:"original_url"`
	ShortenedUrl string    `json:"shortened_url"`
	Visits       int       `json:"visits"`
}
