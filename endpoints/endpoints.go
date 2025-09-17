package endpoints

import (
	"fmt"
)

type RankTrackerDynamicParams struct {
	BaseURL     string
	APIKey      string
	ProjectID   int
	RegionID    int
	PeriodStart string
	PeriodEnd   string
}

func RankTrackerDynamicURL(params RankTrackerDynamicParams) string {
	return fmt.Sprintf(
		"%s/result/ranktracker/dynamic/%d/%d?periodStart=%s&periodEnd=%s&apikey=%s",
		params.BaseURL,
		params.ProjectID,
		params.RegionID,
		params.PeriodStart,
		params.PeriodEnd,
		params.APIKey,
	)
}
