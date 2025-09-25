package api

import (
	"io"
	"ra-api-client/client"
	"ra-api-client/endpoints"
	"ra-api-client/errors"
	"ra-api-client/response"
)

func GetData(params endpoints.RankTrackerDynamicParams) (*response.RankTrackerResponse, error) {
	url, err := endpoints.RankTrackerDynamicURL(params)

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

	return response.ParseRankTrackerResponse(body)
}
