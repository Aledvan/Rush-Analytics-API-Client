package endpoints

import (
	"fmt"
)

type RankTrackerParams struct {
	ProjectID   int
	Page        int
	PeriodStart string
	PeriodEnd   string
}

func (p RankTrackerParams) Validate() error {
	if p.ProjectID <= 0 {
		return fmt.Errorf("ProjectID must be positive, got %d", p.ProjectID)
	}
	if p.Page <= 0 {
		return fmt.Errorf("Page must be positive, got %d", p.Page)
	}

	return nil
}
