package user

import (
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func BalanceGetData() ([]byte, error) {
	builder := endpoints.UserURLBuilder{}
	return api.APIService(builder, "balance")
}
