package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func PositionsGetData(params endpoints.GdsParams) ([]byte, error) {
	builder := endpoints.GdsURLBuilder{Params: params}
	return api.APIService(builder, "positions")
}
