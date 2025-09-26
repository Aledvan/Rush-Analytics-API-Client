package ranktracker

import (
	"encoding/json"
	"fmt"
	"ra-api-client/errors"
)

type Visibility struct {
	Date       string `json:"date"`
	Visibility string `json:"visibility"`
}

type VisibilityGroup struct {
	Group        string       `json:"group"`
	Visibilities []Visibility `json:"visibilities"`
}

type VisibilityResponse struct {
	Keywords []Visibility `json:"visibility"`
}

func ParseVisibilityResponse(body []byte) (*VisibilityResponse, error) {
	var resp VisibilityResponse

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, errors.NewAPIError(
			400,
			fmt.Sprintf("Invalid JSON: %v. Raw: %s", err, string(body)),
		)
	}

	return &resp, nil
}
