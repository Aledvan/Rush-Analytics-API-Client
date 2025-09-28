package ranktracker

import (
	"ra-api-client/endpoints"
)

func RegionsGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "regions")
}
