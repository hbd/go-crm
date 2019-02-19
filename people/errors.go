package people

import (
	"fmt"
	"net/http"
)

// ErrNotFound - Resource not found.
type ErrNotFound struct {
	code       int
	statusText string
	message    string
}

// NewErrNotFound returns a 404 Not Found error.
func NewErrNotFound(resource string) *ErrNotFound {
	return &ErrNotFound{
		http.StatusNotFound,
		http.StatusText(http.StatusNotFound),
		fmt.Sprintf("%s not found", resource),
	}
}

func (e *ErrNotFound) Error() string {
	return e.message
}

// Code return the error's code.
func (e *ErrNotFound) Code() int {
	return e.code
}
