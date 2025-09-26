package ranktracker

import (
	"ra-api-client/endpoints"
	"ra-api-client/response/ranktracker"
)

func CompetitorsGetData(params endpoints.RankTrackerParams) (*[]string, error) {
	return apiService(params, "competitors", ranktracker.ParseCompetitorsResponse)
}
