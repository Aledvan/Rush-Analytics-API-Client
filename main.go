package main

import (
	"fmt"
	"ra-api-client/api/ranktracker"
	"ra-api-client/endpoints"
)

func main() {
	params := endpoints.RankTrackerParams{
		ProjectID: 1054510,
		Page:      1,
	}

	resp, err := ranktracker.RegionsGetData(params)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", resp)
}
