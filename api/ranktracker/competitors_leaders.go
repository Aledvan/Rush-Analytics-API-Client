package ranktracker

import (
	"ra-api-client/endpoints"
)

func CompetitorsLeadersGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "competitors_leaders")
}
