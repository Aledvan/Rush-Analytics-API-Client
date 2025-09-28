package ranktracker

import (
	"ra-api-client/endpoints"
)

func VisibilityGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "visibility")
}
