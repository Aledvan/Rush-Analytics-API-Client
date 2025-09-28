package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func CompareRegionsGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	builder := endpoints.RankTrackerURLBuilder{Params: params}
	return api.APIService(builder, "compare_regions")
}
