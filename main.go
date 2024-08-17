package main

import (
	urlshorten "github.com/hanzalahimran7/url_shorten/app"
)

func main() {
	app := urlshorten.Initialize()
	app.Run()
}
