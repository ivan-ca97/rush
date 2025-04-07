package custom_errors

import (
	"fmt"
	"net/http"
)

type RequestError struct {
	Code    int
	Message string
}

func (e *RequestError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *RequestError) StatusCode() int {
	return http.StatusBadRequest
}
