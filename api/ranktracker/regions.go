package ranktracker

import (
	"ra-api-client/endpoints"
	"ra-api-client/response/ranktracker"
)

func RegionsGetData(params endpoints.RankTrackerParams) (*ranktracker.RegionsResponse, error) {
	return apiService(params, "regions", ranktracker.ParseRegionsResponse)
}
