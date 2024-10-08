package store

import (
	"context"

	"github.com/hanzalahimran7/url_shorten/model"
)

type Database interface {
	CreateUrl(context.Context, model.Url) error
	GetUrl(url string) (model.Url, error)
	GetUrlStats(url string) (model.Stats, error)
}
