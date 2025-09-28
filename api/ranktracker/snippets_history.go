package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func SnippetsHistoryGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	builder := endpoints.RankTrackerURLBuilder{Params: params}
	return api.APIService(builder, "snippets_history")
}
