package main

import (
	"fmt"

	urlshorten "github.com/hanzalahimran7/url_shorten/app"
	"github.com/redis/go-redis/v9"
)

func main() {
	db_options := redis.Options{
		Addr:     fmt.Sprintf("%v:%d", "localhost", 6379),
		Password: "",
		DB:       0,
	}
	app := urlshorten.Initialize(&db_options)
	app.Run()
}
