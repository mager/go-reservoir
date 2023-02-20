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
	BPS       int    `json:"bps"`
	Recipient string `json:"recipient"`
}
type Royalties struct {
	Recipient string      `json:"recipient"`
	Breakdown []Breakdown `json:"breakdown"`
	BPS       int         `json:"bps"`
}
type Custom struct {
	BPS       int    `json:"bps"`
	Recipient string `json:"recipient"`
}
type Opensea struct {
	BPS       int    `json:"bps"`
	Recipient string `json:"recipient"`
}
type AllRoyalties struct {
	Custom  []Custom  `json:"custom"`
	Opensea []Opensea `json:"opensea"`
}
type LastBuy struct {
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
	USD     float64 `json:"usd"`
	Native  float64 `json:"native"`
}
type Price struct {
	Currency Currency `json:"currency"`
	Amount   Amount   `json:"amount"`
}
type Token struct {
	Contract string `json:"contract"`
	TokenID  string `json:"tokenId"`
	Name     string `json:"name"`
	Image    string `json:"image"`
}
type FloorAsk struct {
	ID           string `json:"id"`
	SourceDomain string `json:"sourceDomain"`
	Price        Price  `json:"price"`
	Maker        string `json:"maker"`
	ValidFrom    int    `json:"validFrom"`
	ValidUntil   int    `json:"validUntil"`
	Token        Token  `json:"token"`
}
type Rank struct {
	OneDay    interface{} `json:"1day"`
	SevenDay  int         `json:"7day"`
	Three0Day int         `json:"30day"`
	AllTime   int         `json:"allTime"`
}
type Volume struct {
	OneDay    float64 `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
	AllTime   float64 `json:"allTime"`
}
type VolumeChange struct {
	OneDay    float64 `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
}
type FloorSale struct {
	OneDay    float64 `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
}
type FloorSaleChange struct {
	OneDay    float64 `json:"1day"`
	SevenDay  float64 `json:"7day"`
	Three0Day float64 `json:"30day"`
}
type Collection struct {
	ID                        string          `json:"id"`
	Slug                      string          `json:"slug"`
	CreatedAt                 time.Time       `json:"createdAt"`
	Name                      string          `json:"name"`
	Image                     string          `json:"image"`
	Banner                    string          `json:"banner"`
	DiscordURL                string          `json:"discordUrl"`
	externalURL               string          `json:"externalUrl"`
	TwitterUsername           string          `json:"twitterUsername"`
	OpenseaVerificationstatus string          `json:"openseaVerificationStatus"`
	Description               string          `json:"description"`
	SampleImages              []string        `json:"sampleImages"`
	TokenCount                string          `json:"tokenCount"`
	OnSaleCount               string          `json:"onSaleCount"`
	PrimaryContract           string          `json:"primaryContract"`
	TokensetUD                string          `json:"tokenSetId"`
	Royalties                 Royalties       `json:"royalties"`
	AllRoyalties              AllRoyalties    `json:"allRoyalties"`
	LastBuy                   LastBuy         `json:"lastBuy"`
	FloorAsk                  FloorAsk        `json:"floorAsk"`
	Rank                      Rank            `json:"rank"`
	Volume                    Volume          `json:"volume"`
	VolumeChange              VolumeChange    `json:"volumeChange"`
	FloorSale                 FloorSale       `json:"floorSale"`
	FloorSaleChange           FloorSaleChange `json:"floorSaleChange"`
	CollectionBidSupported    bool            `json:"collectionBidSupported"`
	OwnerCount                int             `json:"ownerCount"`
}

type GetCollectionsOptions struct {
	Slug              string
	IncludeOwnerCount bool
}

// GetCollections gets a list of collections by slug
// https://docs.reservoir.tools/reference/getcollectionsv5
func (c *ReservoirClient) GetCollections(opts GetCollectionsOptions) (CollectionsResp, error) {
	var resp CollectionsResp
	var err error

	u, _ := url.Parse(fmt.Sprintf("%s/collections/v5", c.baseURL))
	q := u.Query()
	q.Add("slug", opts.Slug)

	if opts.IncludeOwnerCount {
		q.Add("includeOwnerCount", "true")
	}

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
