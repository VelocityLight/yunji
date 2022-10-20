package provider

type DataProvider interface {
    // Fetch data from cloud and store in the db.
	FetchData() error
}
