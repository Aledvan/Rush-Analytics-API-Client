package ranktracker

import (
	"encoding/json"
	"fmt"
	"ra-api-client/errors"
)

type SearchEngine struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Device string `json:"device"`
}

type RegionsResponse struct {
	SearchEngine []SearchEngine `json:"searchEngine"`
}

func ParseRegionsResponse(body []byte) (*RegionsResponse, error) {
	var resp RegionsResponse

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, errors.NewAPIError(
			400,
			fmt.Sprintf("Invalid JSON: %v. Raw: %s", err, string(body)),
		)
	}

	return &resp, nil
}
