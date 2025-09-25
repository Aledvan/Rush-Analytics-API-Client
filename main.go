package main

import (
	"fmt"
	"ra-api-client/api"
	"ra-api-client/endpoints"
)

func main() {
	params := endpoints.RankTrackerDynamicParams{
		ProjectID:   1054510,
		RegionID:    1,
		PeriodStart: "2025-06-01",
		PeriodEnd:   "2025-06-30",
	}

	resp, err := api.GetData(params)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Response: %t\n", resp)
}
