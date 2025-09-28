package ranktracker

import (
	"ra-api-client/endpoints"
)

func compareRegionsGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "compare_regions")
}
