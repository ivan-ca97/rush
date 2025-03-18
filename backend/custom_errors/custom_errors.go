package custom_errors

type ExpectedError interface {
	error
	StatusCode() int
}
