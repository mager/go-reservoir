# go-reservoir

[![Go Reference](https://pkg.go.dev/badge/github.com/mager/go-reservoir.svg)](https://pkg.go.dev/github.com/mager/go-reservoir)

A Go client for the Reservoir API.

## Endpoints supported

### v1

_https://docs.reservoir.tools/reference/_

- [x] Collections


## Example usage

```go
package main

import (
	"github.com/mager/go-reservoir/reservoir"
)

func main() {
	// Create a new client
	client := reservoir.NewReservoirClient("")

	// Get Collections
	slug := "azuki"
	c, err := client.GetCollection(slug)
	if err != nil {
		// TODO: Handle error
	}

	// Print the collection
    client.Log.Info(c.Name)
}
```

