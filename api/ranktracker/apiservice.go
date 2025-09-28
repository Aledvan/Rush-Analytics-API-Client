package ranktracker

import (
	"io"
	"ra-api-client/client"
	"ra-api-client/endpoints"
	"ra-api-client/errors"
)

func apiService(
	params endpoints.RankTrackerParams,
	endpoint string,
) ([]byte, error) {
	url, err := endpoints.RankTrackerURL(params, endpoint)
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
