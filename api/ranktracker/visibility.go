package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func VisibilityGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	builder := endpoints.RankTrackerURLBuilder{Params: params}
	return api.APIService(builder, "visibility")
}
