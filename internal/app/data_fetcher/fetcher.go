package data_fetcher

import (
	provider "yunji/internal/app/data_fetcher/data_provider"
)

type DataFetcher struct {
	Providers []provider.DataProvider
}

func (fetcher *DataFetcher) Register(provider provider.DataProvider) {
	fetcher.Providers = append(fetcher.Providers, provider)
}

func (fetcher *DataFetcher) initProviders() {
	fetcher.Register(provider.AWSRealtimeMockProvider{})
}

func FetchData() {
	fetcher := DataFetcher{}
	fetcher.initProviders()

	for _, provider := range fetcher.Providers {
		provider.FetchData()
	}
}
