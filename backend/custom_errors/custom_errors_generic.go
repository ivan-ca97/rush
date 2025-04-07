package custom_errors

import (
	"fmt"
	"net/http"
)

type RecordNotFoundError struct {
	Message string
}

func (e *RecordNotFoundError) Error() string {
	return fmt.Sprintf("Error %d: %s", http.StatusNotFound, e.Message)
}

func (e *RecordNotFoundError) StatusCode() int {
	return http.StatusNotFound
}

type InternalServerError struct {
	Message string
}

func (e *InternalServerError) Error() string {
	return fmt.Sprintf("Error %d: %s", http.StatusInternalServerError, e.Message)
}

func (e *InternalServerError) StatusCode() int {
	return http.StatusInternalServerError
}
