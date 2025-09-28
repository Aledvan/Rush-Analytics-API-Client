package api

import (
	"io"
	"ra-api-client/client"
	"ra-api-client/errors"
)

type URLBuilder interface {
	Build(endpoint string) (string, error)
}

func APIService(builder URLBuilder, endpoint string) ([]byte, error) {
	url, err := builder.Build(endpoint)
	if err != nil {
		return nil, err
	}

	httpClient := client.NewHttpClient()
	resp, err := httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.NewNetworkError("Read response body", err)
	}

	return body, nil
}
