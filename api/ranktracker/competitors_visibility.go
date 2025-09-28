package ranktracker

import (
	"ra-api-client/endpoints"
)

func CompetitorsVisibilityGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "competitors_visibility")
}
