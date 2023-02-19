package reservoir

type Attribute struct {
	Key            string    `json:"key"`
	Value          string    `json:"value"`
	FloorAskPrices []float64 `json:"floorAskPrices"`
	SampleImages   []string  `json:"sampleImages"`
}
