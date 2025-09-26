package client

import (
	"bytes"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"ra-api-client/config"
	"ra-api-client/errors"
	"time"
)

type HTTPClient struct {
	client *http.Client
	config *config.Config
	logger *slog.Logger
}

func NewHttpClient() *HTTPClient {
	cfg, _ := config.GetConfig()
	logger := slog.Default()
	return &HTTPClient{
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
		config: cfg,
		logger: logger,
	}
}

func (c *HTTPClient) Get(url string) (*http.Response, error) {
	c.logger.Info("Making GET request", "url", url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		c.logger.Error("Failed to create GET request", "url", url, "error", err)
		return nil, &errors.NetworkError{
			Op:  "GET " + url,
			Err: err,
		}
	}

	resp, err := c.client.Do(req)
	if err != nil {
		c.logger.Error("GET request failed", "url", url, "error", err)
		return nil, &errors.NetworkError{
			Op:  "GET " + url,
			Err: err,
		}
	}

	if resp.StatusCode >= 400 {
		body, readErr := io.ReadAll(resp.Body)
		if readErr != nil {
			c.logger.Error("Failed to read error response body", "url", url, "error", readErr)
			return nil, errors.NewAPIError(resp.StatusCode, http.StatusText(resp.StatusCode))
		}

		var errorResponse struct {
			Error string `json:"error"`
		}

		if err := json.Unmarshal(body, &errorResponse); err == nil && errorResponse.Error != "" {
			c.logger.Warn("API returned error", "url", url, "status", resp.StatusCode, "error", errorResponse.Error)
			return nil, errors.NewAPIError(resp.StatusCode, errorResponse.Error)
		}

		c.logger.Warn("API returned error", "url", url, "status", resp.StatusCode, "body", string(body))
		return nil, errors.NewAPIError(resp.StatusCode, string(body))
	}

	c.logger.Info("GET request successful", "url", url, "status", resp.StatusCode)
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
