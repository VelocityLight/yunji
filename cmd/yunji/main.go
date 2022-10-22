package main

import (
	"yunji/api"
	"yunji/internal/app/data_fetcher"
)

func main() {
	go data_fetcher.FetchData()

	router := api.NewGinRouter()
	router = api.RouteWebsite(router, "website/build/")
	api.NewHTTPHandler(router)
	router.Run(":8080")
}
