package endpoints

type GdsURLBuilder struct {
	Params GdsParams
}

func (r GdsURLBuilder) Build(endpoint string) (string, error) {
	return GdsURL(r.Params, endpoint)
}
