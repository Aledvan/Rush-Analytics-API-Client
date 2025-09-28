package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func StatusGetData(params endpoints.StatusParams) ([]byte, error) {
	builder := endpoints.StatusURLBuilder{Params: params}
	return api.APIService(builder, "status")
}
