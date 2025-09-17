package client

import (
	"encoding/json"
	_ "fmt"
	"net/http"
	"ra-api-client/config"
	"ra-api-client/errors"
)

type HTTPClient struct {
	client *http.Client
	config *config.Config
}

func NewHTTPClient(cfg *config.Config) *HTTPClient {
	return &HTTPClient{
		client: &http.Client{},
		config: cfg,
	}
}

func (c *HTTPClient) Get(url string) (*http.Response, error) {
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, &errors.NetworkError{
			Op:  "GET " + url,
			Err: err,
		}
	}

	if resp.StatusCode >= 400 {
		var errorResponse struct {
			Error string `json:"error"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&errorResponse); err == nil && errorResponse.Error != "" {
			return nil, errors.NewAPIError(resp.StatusCode, errorResponse.Error)
		}
		return nil, errors.NewAPIError(resp.StatusCode, http.StatusText(resp.StatusCode))
	}

	return resp, nil
}
