package ranktracker

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func SchemaGetData(params endpoints.GdsParams) ([]byte, error) {
	builder := endpoints.GdsURLBuilder{Params: params}
	return api.APIService(builder, "schema")
}
