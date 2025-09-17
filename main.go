package main

import (
	"fmt"
	"log"
	"ra-api-client/api"
	"ra-api-client/client"
	"ra-api-client/config"
	"ra-api-client/endpoints"
)

func main() {
	// Загружаем конфигурацию
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to load config:", err)
	}

	fmt.Printf("Loaded config - BaseURL: %s\n", cfg.API.BaseURL)
	fmt.Printf("API Key length: %d characters\n", len(cfg.API.APIKey))

	// Создаем HTTP клиент
	httpClient := client.NewHTTPClient(cfg)

	// Создаем сервис для работы с API
	rankService := api.NewRankTrackerService(httpClient, cfg)

	// Подготавливаем параметры запроса
	params := endpoints.RankTrackerDynamicParams{
		BaseURL:     cfg.API.BaseURL,
		APIKey:      cfg.API.APIKey,
		ProjectID:   1054510,
		RegionID:    1,
		PeriodStart: "2025-06-01",
		PeriodEnd:   "2025-06-30",
	}

	fmt.Printf("Request URL: %s\n", endpoints.RankTrackerDynamicURL(params))

	// Выполняем запрос
	response, err := rankService.GetDynamicData(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	// Выводим результат
	api.PrintRankTrackerResponse(response)
}
