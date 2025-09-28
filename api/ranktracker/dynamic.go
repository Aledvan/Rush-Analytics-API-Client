package ranktracker

import (
	"ra-api-client/endpoints"
)

func DynamicGetData(params endpoints.RankTrackerParams) ([]byte, error) {
	return apiService(params, "dynamic")
}
