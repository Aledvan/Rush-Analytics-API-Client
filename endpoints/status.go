package endpoints

type StatusURLBuilder struct {
	Params StatusParams
}

func (r StatusURLBuilder) Build(endpoint string) (string, error) {
	return StatusURL(r.Params, endpoint)
}
