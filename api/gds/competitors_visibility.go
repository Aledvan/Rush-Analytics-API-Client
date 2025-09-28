package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func CompetitorsVisibilityGetData(params endpoints.GdsParams) ([]byte, error) {
	builder := endpoints.GdsURLBuilder{Params: params}
	return api.APIService(builder, "competitors_visibility")
}
