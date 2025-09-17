package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"ra-api-client/client"
	"ra-api-client/config"
	"ra-api-client/endpoints"
)

type RankTrackerService struct {
	client *client.HTTPClient
	config *config.Config
}

func NewRankTrackerService(httpClient *client.HTTPClient, cfg *config.Config) *RankTrackerService {
	return &RankTrackerService{
		client: httpClient,
		config: cfg,
	}
}

func (s *RankTrackerService) GetDynamicData(params endpoints.RankTrackerDynamicParams) (*RankTrackerResponse, error) {
	url := endpoints.RankTrackerDynamicURL(params)

	log.Printf("Sending request to: %s", url)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get rank tracker: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Для отладки - выводим сырой ответ
	log.Printf("Raw response: %s", string(body))

	var rankResponse RankTrackerResponse
	if err := json.Unmarshal(body, &rankResponse); err != nil {
		return nil, fmt.Errorf("failed to decode JSON response: %w. Raw response: %s", err, string(body))
	}

	// Проверяем разные варианты ошибок
	if !rankResponse.Success {
		errorMsg := rankResponse.Message
		if errorMsg == "" {
			errorMsg = rankResponse.Error
		}
		if errorMsg == "" {
			errorMsg = fmt.Sprintf("API request failed with status code: %d", resp.StatusCode)
		}
		return nil, fmt.Errorf("API returned error: %s", errorMsg)
	}

	return &rankResponse, nil
}

func PrintRankTrackerResponse(response *RankTrackerResponse) {
	fmt.Printf("Success: %t\n", response.Success)
	if response.Message != "" {
		fmt.Printf("Message: %s\n", response.Message)
	}
	if response.Error != "" {
		fmt.Printf("Error: %s\n", response.Error)
	}

	if len(response.Data) == 0 {
		fmt.Println("No data received")
		return
	}

	fmt.Println("Data:")
	for _, data := range response.Data {
		fmt.Printf("Date: %s\n", data.Date)
		for _, rank := range data.Ranks {
			fmt.Printf("  Keyword: %s, Rank: %d, URL: %s\n", rank.Keyword, rank.Rank, rank.Url)
		}
	}
}
