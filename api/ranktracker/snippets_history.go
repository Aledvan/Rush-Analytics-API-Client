package ranktracker

import (
	"ra-api-client/endpoints"
)

func SnippetsHistoryGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "snippets_history")
}
