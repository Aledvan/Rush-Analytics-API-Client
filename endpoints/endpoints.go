package endpoints

import (
	"fmt"
	"net/url"
	"path"
	"ra-api-client/config"
)

type RankTrackerDynamicParams struct {
	ProjectID   int
	RegionID    int
	PeriodStart string
	PeriodEnd   string
}

func RankTrackerDynamicURL(params RankTrackerDynamicParams) (string, error) {
	cfg, err := config.GetConfig()
	base, err := url.Parse(cfg.API.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	base.Path = path.Join(base.Path, "result", "ranktracker", "dynamic",
		fmt.Sprintf("%d", params.ProjectID),
		fmt.Sprintf("%d", params.RegionID),
	)

	q := base.Query()
	q.Set("periodStart", params.PeriodStart)
	q.Set("periodEnd", params.PeriodEnd)
	q.Set("apikey", cfg.API.APIKey)
	base.RawQuery = q.Encode()

	return base.String(), nil
}
