package client

import (
	"bytes"
	"encoding/json"
	"net/http"
	"ra-api-client/config"
	"ra-api-client/errors"
)

type HTTPClient struct {
	client *http.Client
	config *config.Config
}

func NewHttpClient() *HTTPClient {
	cfg, _ := config.GetConfig()
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

func (c *HTTPClient) Post(url string, body interface{}) (*http.Response, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, &errors.NetworkError{
			Op:  "POST " + url,
			Err: err,
		}
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, &errors.NetworkError{
			Op:  "POST " + url,
			Err: err,
		}
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, &errors.NetworkError{
			Op:  "POST " + url,
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
