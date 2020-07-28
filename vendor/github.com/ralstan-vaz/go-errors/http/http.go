package http

import (
	"net/http"

	"github.com/ralstan-vaz/go-errors"
)

type getter interface {
	Get() *errors.Error
}

// StatusCode Check if its a custom error type and then
// gets the corresponding status code from the error kind.
// Returns InternalServerError by default
func StatusCode(err error) int {
	switch t := err.(type) {
	case getter:
		kind := t.Get().Kind
		return getHTTPCode(kind)
	}

	return http.StatusInternalServerError
}

// getHTTPCode return the httpStatus code corresponding to the error kind
func getHTTPCode(kind errors.Kind) int {
	kindToHTTPCode := map[errors.Kind]int{
		errors.NotFound:         http.StatusNotFound,
		errors.Unauthorized:     http.StatusUnauthorized,
		errors.Forbidden:        http.StatusForbidden,
		errors.Expired:          http.StatusBadRequest,
		errors.BadRequest:       http.StatusBadRequest,
		errors.ParameterMissing: http.StatusBadRequest,
		errors.InternalError:    http.StatusInternalServerError,
		errors.Unknown:          http.StatusInternalServerError,
	}

	if httpCode, found := kindToHTTPCode[kind]; found {
		return httpCode
	}

	return http.StatusInternalServerError
}
