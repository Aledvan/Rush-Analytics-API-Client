package ranktracker

import (
	"encoding/json"
	"fmt"
	"ra-api-client/errors"
)

func ParseCompetitorsResponse(body []byte) (*[]string, error) {
	var data []string

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.NewAPIError(
			400,
			fmt.Sprintf("Invalid JSON: %v. Raw: %s", err, string(body)),
		)
	}

	return &data, nil
}
