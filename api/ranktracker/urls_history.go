package ranktracker

import (
	"ra-api-client/endpoints"
)

func UrlsHistoryGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "urls_history")
}
