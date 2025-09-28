package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func TopsGetData(params endpoints.GdsParams) ([]byte, error) {
	builder := endpoints.GdsURLBuilder{Params: params}
	return api.APIService(builder, "tops")
}
