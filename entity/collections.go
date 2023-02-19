package entity

import (
	"time"
)

type CollectionsResp struct {
	Collections  []Collection `json:"collections"`
	Continuation string       `json:"continuation"`
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
	Usd     float64 `json:"usd"`
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
	OneDay    int `json:"1day"`
	SevenDay  int `json:"7day"`
	Three0Day int `json:"30day"`
	AllTime   int `json:"allTime"`
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

type Royalties struct {
	Bps       int    `json:"bps"`
	Recipient string `json:"recipient"`
}

type Collection struct {
	ID                        string          `json:"id"`
	Slug                      string          `json:"slug"`
	CreatedAt                 time.Time       `json:"createdAt"`
	Name                      string          `json:"name"`
	Image                     string          `json:"image"`
	Banner                    string          `json:"banner"`
	DiscordURL                string          `json:"discordUrl"`
	ExternalURL               string          `json:"externalUrl"`
	TwitterUsername           string          `json:"twitterUsername"`
	OpenseaVerificationStatus string          `json:"openseaVerificationStatus"`
	Description               string          `json:"description"`
	SampleImages              []string        `json:"sampleImages"`
	TokenCount                string          `json:"tokenCount"`
	OnSaleCount               string          `json:"onSaleCount"`
	PrimaryContract           string          `json:"primaryContract"`
	TokenSetID                string          `json:"tokenSetId"`
	LastBuy                   LastBuy         `json:"lastBuy"`
	FloorAsk                  FloorAsk        `json:"floorAsk"`
	Rank                      Rank            `json:"rank"`
	Volume                    Volume          `json:"volume"`
	VolumeChange              VolumeChange    `json:"volumeChange"`
	FloorSale                 FloorSale       `json:"floorSale"`
	FloorSaleChange           FloorSaleChange `json:"floorSaleChange"`
	CollectionBidSupported    bool            `json:"collectionBidSupported"`
	Royalties                 Royalties       `json:"royalties,omitempty"`
}
