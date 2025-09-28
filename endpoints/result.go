package endpoints

type ResultURLBuilder struct {
	Params ResultParams
}

func (r ResultURLBuilder) Build(endpoint string) (string, error) {
	return ResultURL(r.Params, endpoint)
}
