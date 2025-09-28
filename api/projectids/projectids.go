package projectids

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func ProjectIdsGetData(params endpoints.ProjectIdsParams) ([]byte, error) {
	builder := endpoints.ProjectIdsURLBuilder{Params: params}
	return api.APIService(builder, "projectids")
}
