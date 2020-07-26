package grpc

import (
	"google.golang.org/grpc/codes"

	"github.com/ralstan-vaz/go-errors"
)

type getter interface {
	Get() *errors.Error
}

// StatusCode Check if its a custom error type and then
// gets the corresponding status code from the error kind.
// Returns Internal Error by default
func StatusCode(err error) codes.Code {
	switch t := err.(type) {
	case getter:
		kind := t.Get().Kind
		return getGRPCCode(kind)
	}

	return codes.Internal
}

// getGRPCCode return the GRPC Status code corresponding to the error kind
func getGRPCCode(kind errors.Kind) codes.Code {
	kindToGRPCCode := map[errors.Kind]codes.Code{
		errors.NotFound:         codes.NotFound,
		errors.Unauthorized:     codes.Unauthenticated,
		errors.Forbidden:        codes.PermissionDenied,
		errors.Expired:          codes.InvalidArgument,
		errors.BadRequest:       codes.InvalidArgument,
		errors.ParameterMissing: codes.InvalidArgument,
		errors.InternalError:    codes.Internal,
		errors.Unknown:          codes.Internal,
	}

	if grpcCode, found := kindToGRPCCode[kind]; found {
		return grpcCode
	}

	return codes.Internal
}
