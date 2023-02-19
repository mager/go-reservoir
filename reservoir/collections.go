package reservoir

import (
	"encoding/json"
	"fmt"
	"net/url"
	"time"
)

type CollectionsResp struct {
	Collections []Collection `json:"collections"`
}

type Breakdown struct {
	Bps       int    `json:"bps"`
	Recipient string `json:"recipient"`
}
type Royalties struct {
	Recipient string      `json:"recipient"`
	Breakdown []Breakdown `json:"breakdown"`
	Bps       int         `json:"bps"`
}
type Custom struct {
	Bps       int    `json:"bps"`
	Recipient string `json:"recipient"`
}
type Opensea struct {
	Bps       int    `json:"bps"`
	Recipient string `json:"recipient"`
}
type Allroyalties struct {
	Custom  []Custom  `json:"custom"`
	Opensea []Opensea `json:"opensea"`
}
type Lastbuy struct {
	Value interface{} `json:"value"`
}
type Currency struct {
	Contract string `json:"contract"`
	Name     string `json:"name"`
	Symbol   string `json:"symbol"`
	Decimals int    `json:"decimals"`
}
type Amount struct {
	Raw     string  `json:"raw"`
	Decimal float64 `json:"decimal"`
	Usd     float64 `json:"usd"`
	Native  float64 `json:"native"`
}
type Price struct {
	Currency Currency `json:"currency"`
	Amount   Amount   `json:"amount"`
}
type Token struct {
	Contract string `json:"contract"`
	Tokenid  string `json:"tokenId"`
	Name     string `json:"name"`
	Image    string `json:"image"`
}
type Floorask struct {
	ID           string `json:"id"`
	Sourcedomain string `json:"sourceDomain"`
	Price        Price  `json:"price"`
	Maker        string `json:"maker"`
	Validfrom    int    `json:"validFrom"`
	Validuntil   int    `json:"validUntil"`
	Token        Token  `json:"token"`
}
type Rank struct {
	OneDay    interface{} `json:"1day"`
	SevenDay  int         `json:"7day"`
	Three0Day int         `json:"30day"`
	Alltime   int         `json:"allTime"`
}
type Volume struct {
	OneDay    int     `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
	Alltime   float64 `json:"allTime"`
}
type Volumechange struct {
	OneDay    int     `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
}
type Floorsale struct {
	OneDay    float64 `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
}
type Floorsalechange struct {
	OneDay    float64 `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
}
type Collection struct {
	ID                        string          `json:"id"`
	Slug                      string          `json:"slug"`
	Createdat                 time.Time       `json:"createdAt"`
	Name                      string          `json:"name"`
	Image                     string          `json:"image"`
	Banner                    string          `json:"banner"`
	Discordurl                string          `json:"discordUrl"`
	Externalurl               string          `json:"externalUrl"`
	Twitterusername           string          `json:"twitterUsername"`
	Openseaverificationstatus string          `json:"openseaVerificationStatus"`
	Description               string          `json:"description"`
	Sampleimages              []string        `json:"sampleImages"`
	Tokencount                string          `json:"tokenCount"`
	Onsalecount               string          `json:"onSaleCount"`
	Primarycontract           string          `json:"primaryContract"`
	Tokensetid                string          `json:"tokenSetId"`
	Royalties                 Royalties       `json:"royalties"`
	Allroyalties              Allroyalties    `json:"allRoyalties"`
	Lastbuy                   Lastbuy         `json:"lastBuy"`
	Floorask                  Floorask        `json:"floorAsk"`
	Rank                      Rank            `json:"rank"`
	Volume                    Volume          `json:"volume"`
	Volumechange              Volumechange    `json:"volumeChange"`
	Floorsale                 Floorsale       `json:"floorSale"`
	Floorsalechange           Floorsalechange `json:"floorSaleChange"`
	Collectionbidsupported    bool            `json:"collectionBidSupported"`
}

// GetCollections gets a list of collections by slug
// https://docs.reservoir.tools/reference/getcollectionsv5
func (c *ReservoirClient) GetCollections(slug string) (CollectionsResp, error) {
	var resp CollectionsResp
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/collections/v5", c.baseURL))
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

	// Log response body
	return resp, err
}
