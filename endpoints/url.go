package endpoints

import (
	"fmt"
	"net/url"
	"path"
	"ra-api-client/config"
)

type Validator interface {
	Validate() error
}

func buildURL(valid Validator, pathParts []string, queryParams map[string]string) (string, error) {
	if valid != nil {
		if err := valid.Validate(); err != nil {
			return "", fmt.Errorf("invalid parameters: %w", err)
		}
	}

	cfg, err := config.GetConfig()
	if err != nil {
		return "", fmt.Errorf("failed to load config: %w", err)
	}

	base, err := url.Parse(cfg.API.BaseURL)
	if err != nil {
		return "", fmt.Errorf("invalid base URL: %w", err)
	}

	base.Path = path.Join(append([]string{base.Path}, pathParts...)...)

	q := base.Query()
	q.Set("apikey", cfg.API.APIKey)
	for k, v := range queryParams {
		q.Set(k, v)
	}
	base.RawQuery = q.Encode()

	return base.String(), nil
}

func RankTrackerURL(params RankTrackerParams, urlpart string) (string, error) {
	pathParts := []string{"result", "ranktracker", urlpart, fmt.Sprintf("%d", params.ProjectID), fmt.Sprintf("%d", params.Page)}
	queryParams := map[string]string{
		"periodStart": params.PeriodStart,
		"periodEnd":   params.PeriodEnd,
	}
	return buildURL(params, pathParts, queryParams)
}

func ResultURL(params ResultParams, urlpart string) (string, error) {
	pathParts := []string{"result", urlpart, fmt.Sprintf("%d", params.ProjectID)}
	queryParams := make(map[string]string)
	return buildURL(params, pathParts, queryParams)
}

func UserURL(urlpart string) (string, error) {
	pathParts := []string{urlpart}
	queryParams := make(map[string]string)
	return buildURL(nil, pathParts, queryParams)
}

func ProjectIdsURL(params ProjectIdsParams, urlpart string) (string, error) {
	var pathParts []string
	pathParts = append(pathParts, urlpart)
	if params.ProjectType != 0 {
		pathParts = append(pathParts, fmt.Sprintf("%d", params.ProjectType))
	}
	queryParams := make(map[string]string)
	return buildURL(params, pathParts, queryParams)
}

func StatusURL(params StatusParams, urlpart string) (string, error) {
	pathParts := []string{urlpart, fmt.Sprintf("%d", params.ProjectType), fmt.Sprintf("%d", params.ProjectID)}
	queryParams := make(map[string]string)
	return buildURL(params, pathParts, queryParams)
}

func GdsURL(params GdsParams, urlpart string) (string, error) {
	pathParts := []string{"gds", urlpart, fmt.Sprintf("%d", params.ProjectID)}
	queryParams := make(map[string]string)
	return buildURL(params, pathParts, queryParams)
}
