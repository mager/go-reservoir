package examples

import (
	"github.com/mager/go-reservoir/reservoir"
)

func main() {
	// Create a new client
	client := reservoir.NewReservoirClient("") // Your API key goes here

	// Get Collections
	slug := "azuki"
	c, err := client.GetCollections(slug)
	if err != nil {
		// TODO: Handle error
	}

	client.Log.Info(c.Name)
}
