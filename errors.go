package app

import "strings"

const (
	// ErrNotFound is an implementation agnostic error that should be returned
	// by any service implementation when a record was not located.
	ErrNotFound appError = "app: resource not found"
)

type appError string

func (e appError) Error() string {
	return string(e)
}

func (e appError) Public() string {
	return strings.Replace(string(e), "app: ", "", 1)
 // TODO: Improve this using golang.org/x/text/cases Title
}
