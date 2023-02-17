package reservoir

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockTransport struct {
	statusCode int
	body       string
}

func (t *mockTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return &http.Response{
		StatusCode: t.statusCode,
		Body:       ioutil.NopCloser(strings.NewReader(t.body)),
	}, nil
}

func TestGetCollections(t *testing.T) {
	// Create a mock HTTP client and a mock response
	mockClient := &http.Client{
		Transport: &mockTransport{
			statusCode: 200,
			body:       `{"collections":[{"name":"Collection 1"},{"name":"Collection 2"}]}`,
		},
	}
	// Create a new ReservoirClient instance with the mock HTTP client
	client := NewReservoirClient("api-key")
	client.client = mockClient

	// Call the GetCollections method
	resp, err := client.GetCollections("slug")

	// Assert that the response is as expected
	assert.Nil(t, err)
	assert.Equal(t, 2, len(resp.Collections))
	assert.Equal(t, "Collection 1", resp.Collections[0].Name)
	assert.Equal(t, "Collection 2", resp.Collections[1].Name)
}
