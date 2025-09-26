package ranktracker

import (
	"ra-api-client/endpoints"
	"ra-api-client/response/ranktracker"
)

func DynamicGetData(params endpoints.RankTrackerParams) (*ranktracker.DynamicResponse, error) {
	return apiService(params, "dynamic", ranktracker.ParseDynamicResponse)
}
