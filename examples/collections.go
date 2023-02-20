package main

import (
	"github.com/mager/go-reservoir/reservoir"
)

func main() {
	// Create a new client
	client := reservoir.NewReservoirClient("") // Your API key goes here

	// Get Collections
	slug := "hot-dougs"
	opts := reservoir.GetCollectionsOptions{
		Slug:              slug,
		IncludeOwnerCount: true,
	}
	c, err := client.GetCollections(opts)
	if err != nil {
		// TODO: Handle error
	}

	// Loop through collections and print the name
	for _, collection := range c.Collections {
		client.Log.Infow("Fetched collection", "name", collection.Name)
	}
}
