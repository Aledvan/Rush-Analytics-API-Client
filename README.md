# Rush-Analytics API Client

A Go client library for interacting with the Rush-Analytics API based on their official Swagger documentation available at [https://app.rush-analytics.ru/apiv2/doc/](https://app.rush-analytics.ru/apiv2/doc/).

## Features

- Easy-to-use Go functions for all Rush-Analytics API endpoints
- Structured parameter handling with type safety
- Error handling and response management
- Modular design with separate packages for different API sections

## Installation

```bash
go mod init your-project-name
go get github.com/Aledvan/Rush-Analytics-API-Client
```

## Usage
### Basic Example

```bash
package main

import (
	"fmt"
	"ra-api-client/api/ranktracker"
	"ra-api-client/endpoints"
)

func main() {
	params := endpoints.RankTrackerParams{
		PeriodStart: "2025-06-01",
		PeriodEnd:   "2025-06-30",
		ProjectID:   1054510,
		Page:        1,
	}

	resp, err := ranktracker.CompetitorsGetData(params)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Printf("Response: %+v\n", string(resp))
}
```

## Available API Modules
### Rank Tracker
```bash
import "ra-api-client/api/ranktracker"

// Examples:

resp, err := ranktracker.DynamicGetData(params)
resp, err := ranktracker.CompetitorsGetData(params)
```
### Project IDs
```bash
import "ra-api-client/api/projectids"

// Examples:
resp, err := projectids.ProjectIdsGetData(params)
```