package ranktracker

import (
	"ra-api-client/endpoints"
	"ra-api-client/response/ranktracker"
)

func VisibilityGetData(params endpoints.RankTrackerParams) (*ranktracker.VisibilityResponse, error) {
	return apiService(params, "visibility", ranktracker.ParseVisibilityResponse)
}
