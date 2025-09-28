package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func CompetitorsLeadersGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	builder := endpoints.RankTrackerURLBuilder{Params: params}
	return api.APIService(builder, "competitors_leaders")
}
