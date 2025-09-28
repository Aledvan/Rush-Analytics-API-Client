package endpoints

import (
	"fmt"
	"net/url"
	"path"
	"ra-api-client/config"
)

func RankTrackerURL(params RankTrackerParams, urlpart string) (string, error) {
	if err := params.Validate(); err != nil {
		return "", fmt.Errorf("invalid parameters: %w", err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	base, err := url.Parse(cfg.API.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	base.Path = path.Join(base.Path, "result", "ranktracker", urlpart,
		fmt.Sprintf("%d", params.ProjectID),
		fmt.Sprintf("%d", params.Page),
	)

	q := base.Query()
	q.Set("periodStart", params.PeriodStart)
	q.Set("periodEnd", params.PeriodEnd)
	q.Set("apikey", cfg.API.APIKey)
	base.RawQuery = q.Encode()

	return base.String(), nil
}

func ResultURL(params ResultParams, urlpart string) (string, error) {
	if err := params.Validate(); err != nil {
		return "", fmt.Errorf("invalid parameters: %w", err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	base, err := url.Parse(cfg.API.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	base.Path = path.Join(base.Path, "result", urlpart,
		fmt.Sprintf("%d", params.ProjectID),
	)

	q := base.Query()
	q.Set("apikey", cfg.API.APIKey)
	base.RawQuery = q.Encode()

	return base.String(), nil
}

func UserURL(urlpart string) (string, error) {

	cfg, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	base, err := url.Parse(cfg.API.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	base.Path = path.Join(base.Path, urlpart)

	q := base.Query()
	q.Set("apikey", cfg.API.APIKey)
	base.RawQuery = q.Encode()

	return base.String(), nil
}

func ProjectIdsURL(params ProjectIdsParams, urlpart string) (string, error) {
	if err := params.Validate(); err != nil {
		return "", fmt.Errorf("invalid parameters: %w", err)
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	base, err := url.Parse(cfg.API.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	base.Path = path.Join(base.Path, urlpart)

	q := base.Query()
	q.Set("apikey", cfg.API.APIKey)
	base.RawQuery = q.Encode()

	return base.String(), nil
}
