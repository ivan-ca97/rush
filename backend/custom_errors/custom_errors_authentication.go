package custom_errors

import (
	"fmt"
)

type AuthenticationError struct {
	Code    int
	Message string
}

func (e *AuthenticationError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *AuthenticationError) StatusCode() int {
	return e.Code
}
