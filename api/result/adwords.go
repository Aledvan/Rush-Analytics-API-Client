package result

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func AdwordsGetData(params endpoints.ResultParams) ([]byte, error) {
	builder := endpoints.ResultURLBuilder{Params: params}
	return api.APIService(builder, "adwords")
}
