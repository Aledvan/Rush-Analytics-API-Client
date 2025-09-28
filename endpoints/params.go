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

type ResultParams struct {
	ProjectID int
}

func (p ResultParams) Validate() error {
	if p.ProjectID <= 0 {
		return fmt.Errorf("ProjectID must be positive, got %d", p.ProjectID)
	}

	return nil
}

type ProjectIdsParams struct {
	ProjectType int
}

func (p ProjectIdsParams) Validate() error {
	return nil
}

type StatusParams struct {
	ProjectID   int
	ProjectType int
}

func (p StatusParams) Validate() error {
	if p.ProjectID <= 0 {
		return fmt.Errorf("ProjectID must be positive, got %d", p.ProjectID)
	}

	if p.ProjectType <= 0 {
		return fmt.Errorf("ProjectType must be positive, got %d", p.ProjectType)
	}

	return nil
}

type GdsParams struct {
	ProjectID int
}

func (p GdsParams) Validate() error {
	if p.ProjectID <= 0 {
		return fmt.Errorf("ProjectID must be positive, got %d", p.ProjectID)
	}

	return nil
}
