package ranktracker

import (
	"ra-api-client/endpoints"
)

func CompetitorsPositionsGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "competitors_positions")
}
