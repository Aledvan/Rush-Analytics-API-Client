package endpoints

type RankTrackerURLBuilder struct {
	Params RankTrackerParams
}

func (r RankTrackerURLBuilder) Build(endpoint string) (string, error) {
	return RankTrackerURL(r.Params, endpoint)
}
