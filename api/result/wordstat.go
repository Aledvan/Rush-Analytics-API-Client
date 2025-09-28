package result

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func WordstatGetData(params endpoints.ResultParams) ([]byte, error) {
	builder := endpoints.ResultURLBuilder{Params: params}
	return api.APIService(builder, "wordstat")
}
