package ranktracker

import (
	"encoding/json"
	"fmt"
	"ra-api-client/errors"
)

type PositionHistory struct {
	Date     string `json:"date"`
	Position int    `json:"position"`
}

type KeywordData struct {
	Keyword          string            `json:"keyword"`
	PositionsHistory []PositionHistory `json:"positions_history"`
}

type PositionsHistoryResponse struct {
	Data []KeywordData `json:"-"`
}

func ParsePositionsHistoryResponse(body []byte) (*PositionsHistoryResponse, error) {
	var data []KeywordData

	if err := json.Unmarshal(body, &data); err != nil {
		return nil, errors.NewAPIError(
			400,
			fmt.Sprintf("Invalid JSON: %v. Raw: %s", err, string(body)),
		)
	}

	resp := &PositionsHistoryResponse{
		Data: data,
	}

	return resp, nil
}
