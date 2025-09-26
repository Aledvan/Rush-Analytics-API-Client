package ranktracker

import (
	"ra-api-client/endpoints"
	"ra-api-client/response/ranktracker"
)

func PositionsHistoryGetData(params endpoints.RankTrackerParams) (*ranktracker.PositionsHistoryResponse, error) {
	return apiService(params, "positions_history", ranktracker.ParsePositionsHistoryResponse)
}
