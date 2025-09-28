package result

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func Top10GetData(params endpoints.ResultParams) ([]byte, error) {
	builder := endpoints.ResultURLBuilder{Params: params}
	return api.APIService(builder, "top10")
}
