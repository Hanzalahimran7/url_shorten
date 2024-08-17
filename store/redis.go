package store

import (
	"context"
	"log"

	"github.com/hanzalahimran7/url_shorten/model"
	"github.com/redis/go-redis/v9"
)

type redisDB struct {
	db *redis.Client
}

func GetRedisInstance(options *redis.Options) (Database, error) {
	rdb := redis.NewClient(options)
	ctx := context.Background()
	// Ping the Redis server
	pong, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Println("Connection failed with Redis:", err)
		return nil, err
	}
	log.Println("Connection successful with Redis:", pong)
	return &redisDB{db: rdb}, nil
}

func (r *redisDB) CreateUrl(url string) (model.Url, error)     { return model.Url{}, nil }
func (r *redisDB) GetUrl(url string) (model.Url, error)        { return model.Url{}, nil }
func (r *redisDB) GetUrlStats(url string) (model.Stats, error) { return model.Stats{}, nil }
