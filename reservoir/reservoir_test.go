package reservoir

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"
)

func TestNewReservoirClient(t *testing.T) {
	c := GenerateReservoirClient()
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
	c := GenerateReservoirClient()

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

func TestGet(t *testing.T) {
	c := GenerateReservoirClient()

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
			got, err := c.Get(tt.args.u)

			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got.Request.Method, tt.want.Method) {
				t.Errorf("Get() = %v, want %v", got.Request.Method, tt.want.Method)
			}

			if !reflect.DeepEqual(got.Request.URL, tt.want.URL) {
				t.Errorf("Get() = %v, want %v", got.Request.URL, tt.want.URL)
			}

			if !reflect.DeepEqual(got.Request.Header, tt.want.Header) {
				t.Errorf("Get() = %v, want %v", got.Request.Header, tt.want.Header)
			}

			// TODO: Fix test
			// if !reflect.DeepEqual(got.StatusCode, 200) {
			// 	t.Errorf("Get() = %v, want %v", got.StatusCode, 200)
			// }

			// // Test response body
			// body, err := ioutil.ReadAll(got.Body)
			// if err != nil {
			// 	t.Error(err)
			// }

			// var collections []data.Collection
			// err = json.Unmarshal(body, &collections)
			// if err != nil {
			// 	t.Error(err)
			// }

			// if !reflect.DeepEqual(len(collections), 1) {
			// 	t.Errorf("GetRequest() = %v, want %v", len(collections), 1)
			// }

			// if !reflect.DeepEqual(collections[0].Slug, "azuki") {
			// 	t.Errorf("GetRequest() = %v, want %v", collections[0].Slug, "azuki")
			// }

			// if !reflect.DeepEqual(collections[0].Name, "Azuki") {
			// 	t.Errorf("GetRequest() = %v, want %v", collections[0].Name, "Azuki")
			// }
		})
	}
}
