package reservoir

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

func GenerateReservoirClient() *ReservoirClient {
	logger, _ := zap.NewProduction()

	return &ReservoirClient{
		Log: logger.Sugar(),

		apiKey:  "test",
		baseURL: "https://api.reservoir.tools",
		client: &http.Client{
			Timeout: time.Second * 10,
		},
		requestDelay: time.Millisecond * 250,
	}

}
