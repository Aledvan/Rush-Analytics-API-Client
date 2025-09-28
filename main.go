package main

import (
	"fmt"
	"ra-api-client/api/projectids"
	"ra-api-client/endpoints"
)

func main() {
	params := endpoints.ProjectIdsParams{
		ProjectType: 1,
	}

	resp, err := projectids.ProjectIdsGetData(params)
	if err != nil {
		fmt.Printf("Ошибка: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", string(resp))
}
