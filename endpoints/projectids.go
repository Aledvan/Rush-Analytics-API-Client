package endpoints

type ProjectIdsURLBuilder struct {
	Params ProjectIdsParams
}

func (r ProjectIdsURLBuilder) Build(endpoint string) (string, error) {
	return ProjectIdsURL(r.Params, endpoint)
}
