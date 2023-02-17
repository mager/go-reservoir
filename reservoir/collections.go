package reservoir

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/mager/go-reservoir/data"
)

// GetCollections gets a list of collections by slug
// https://docs.reservoir.tools/reference/getcollectionsv5
func (c *ReservoirClient) GetCollections(slug string) (data.CollectionsResp, error) {
	var resp data.CollectionsResp
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/collections/v5/", c.baseURL))
	q := u.Query()
	q.Add("slug", slug)

	u.RawQuery = q.Encode()

	// Make the request
	httpResp, err := c.Get(u)

	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error getting collections: %s", err)
		return resp, err
	}

	defer httpResp.Body.Close()
	err = json.NewDecoder(httpResp.Body).Decode(&resp)
	// TODO: Test
	if err != nil {
		c.Log.Errorf("Error decoding collections: %s", err)
		return resp, err
	}

	// Loop through collections
	for _, collection := range resp.Collections {
		c.Log.Infow("DEBUG", "collection", collection)
	}

	return resp, err
}
