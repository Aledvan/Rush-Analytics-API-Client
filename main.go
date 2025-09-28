package main

import (
	"fmt"
	"ra-api-client/api/ranktracker"
	"ra-api-client/endpoints"
)

func main() {
	params := endpoints.RankTrackerParams{
		PeriodStart: "025-06-01",
		PeriodEnd:   "2025-06-30",
		ProjectID:   1054510,
		Page:        1,
	}

	resp, err := ranktracker.VisibilityGetData(params)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", string(resp))
}
