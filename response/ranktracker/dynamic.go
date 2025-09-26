package ranktracker

import (
	"encoding/json"
	"fmt"
	"ra-api-client/errors"
)

type Position struct {
	Date     string `json:"date"`
	Position string `json:"position"`
}

type Keyword struct {
	Keyword      string     `json:"keyword"`
	Wordstat     int        `json:"wordstat"`
	Wordstat1    int        `json:"wordstat ''"`
	WordstatBang int        `json:"wordstat '!'"`
	Group        string     `json:"group"`
	TargetURL    string     `json:"targeturl"`
	RelevantURL  string     `json:"relevanturl"`
	Positions    []Position `json:"positions"`
}

type DynamicResponse struct {
	Keywords []Keyword `json:"keywords"`
}

func ParseDynamicResponse(body []byte) (*DynamicResponse, error) {
	var resp DynamicResponse

	if err := json.Unmarshal(body, &resp); err != nil {
		return nil, errors.NewAPIError(
			400,
			fmt.Sprintf("Invalid JSON: %v. Raw: %s", err, string(body)),
		)
	}

	return &resp, nil
}
