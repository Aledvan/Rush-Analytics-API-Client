package ranktracker

import (
	"ra-api-client/endpoints"
)

func PositionsHistoryGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "positions_history")
}
