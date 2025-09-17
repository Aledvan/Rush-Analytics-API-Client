package errors

import (
	"log"
	"net/http"
)

func HandleError(err error) {
	switch e := err.(type) {
	case *APIError:
		log.Printf("API Error: %v", e)
	case *ConfigError:
		log.Printf("Config Error: %v", e)
	case *NetworkError:
		log.Printf("Network Error: %v", e)
	default:
		log.Printf("Unknown Error: %v", e)
	}
}

func HTTPStatusFromError(err error) int {
	switch err.(type) {
	case *APIError:
		return http.StatusBadRequest
	case *ConfigError:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}
