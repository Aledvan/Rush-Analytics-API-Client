package ranktracker

import (
	"ra-api-client/endpoints"
)

func CompetitorsGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "competitors")
}
