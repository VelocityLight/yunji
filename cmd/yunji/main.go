package main

import (
	"yunji/api"
	"yunji/internal/app/data_fetcher"
)

func main() {
	go data_fetcher.FetchData()

	router := api.Routers()
	router.Run(":8080")
}
