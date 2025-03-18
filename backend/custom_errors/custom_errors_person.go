package custom_errors

import "fmt"

type PersonError struct {
	Code    int
	Message string
}

func (e *PersonError) Error() string {
	return fmt.Sprintf("Error %d: %s", e.Code, e.Message)
}

func (e *PersonError) StatusCode() int {
	return e.Code
}
