package errors

import "fmt"

type APIError struct {
	statusCode int
	message    string
}

func (e *APIError) Error() string {
	return fmt.Sprintf("API error %d: %s", e.statusCode, e.message)
}

func (e *APIError) StatusCode() int {
	return e.statusCode
}

func NewAPIError(statusCode int, message string) error {
	return &APIError{statusCode: statusCode, message: message}
}

type ConfigError struct {
	message string
}

func (e *ConfigError) Error() string {
	return fmt.Sprintf("Config error: %s", e.message)
}

func NewConfigError(message string) error {
	return &ConfigError{message: message}
}

type NetworkError struct {
	Op  string
	Err error
}

func (e *NetworkError) Error() string {
	return fmt.Sprintf("Network error during %s: %v", e.Op, e.Err)
}

func (e *NetworkError) Unwrap() error {
	return e.Err
}

func NewNetworkError(op string, err error) error {
	return &NetworkError{Op: op, Err: err}
}
