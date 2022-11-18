package reservoir

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
	"time"

	"go.uber.org/zap"
)

func generateReservoirClient() *ReservoirClient {
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

func TestNewReservoirClient(t *testing.T) {
	c := generateReservoirClient()
	type args struct {
		apiKey string
	}
	tests := []struct {
		name string
		args args
		want *ReservoirClient
	}{
		{
			name: "happy path",
			args: args{
				apiKey: c.apiKey,
			},
			want: c,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := NewReservoirClient(tt.args.apiKey)

			if !reflect.DeepEqual(client.apiKey, tt.want.apiKey) {
				t.Errorf("NewReservoirClient() = %v, want %v", client.apiKey, tt.want.apiKey)
			}

			if !reflect.DeepEqual(client.baseURL, tt.want.baseURL) {
				t.Errorf("NewReservoirClient() = %v, want %v", client.baseURL, tt.want.baseURL)
			}

			if !reflect.DeepEqual(client.client, tt.want.client) {
				t.Errorf("NewReservoirClient() = %v, want %v", client.client, tt.want.client)
			}

			if !reflect.DeepEqual(client.requestDelay, tt.want.requestDelay) {
				t.Errorf("NewReservoirClient() = %v, want %v", client.requestDelay, tt.want.requestDelay)
			}
		})
	}
}

func TestGetRequest(t *testing.T) {
	c := generateReservoirClient()

	type args struct {
		u *url.URL
	}
	tests := []struct {
		name    string
		args    args
		want    *http.Request
		wantErr bool
	}{
		{
			name: "happy path",
			args: args{
				u: &url.URL{
					Scheme: "https",
					Host:   "api.reservoir.tools",
					Path:   "/v1/collections",
					RawQuery: url.Values{
						"slug": []string{"azuki"},
					}.Encode(),
				},
			},
			want: &http.Request{
				Method: "GET",
				URL: &url.URL{
					Scheme: "https",
					Host:   "api.reservoir.tools",
					Path:   "/v1/collections",
					RawQuery: url.Values{
						"slug": []string{"azuki"},
					}.Encode(),
				},
				Header: http.Header{
					"X-Api-Key":    []string{"test"},
					"Content-Type": []string{"application/json"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetRequest(tt.args.u)

			if (err != nil) != tt.wantErr {
				t.Errorf("GetRequest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Method, tt.want.Method) {
				t.Errorf("GetRequest() = %v, want %v", got.Method, tt.want.Method)
			}

			if !reflect.DeepEqual(got.URL, tt.want.URL) {
				t.Errorf("GetRequest() = %v, want %v", got.URL, tt.want.URL)
			}

			if !reflect.DeepEqual(got.Header, tt.want.Header) {
				t.Errorf("GetRequest() = %v, want %v", got.Header, tt.want.Header)
			}
		})
	}
}

// TODO: TestGet()
