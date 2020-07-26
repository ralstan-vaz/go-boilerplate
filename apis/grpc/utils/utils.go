package utils

import (
	"fmt"

	"github.com/ralstan-vaz/go-errors"
	"github.com/ralstan-vaz/go-errors/grpc"
	"google.golang.org/grpc/status"
)

// HandleError formats, logs and sets a GRPC response for the error
func HandleError(errObj error) error {
	err := errors.Get(errObj)

	if err.Message == "" {
		err.Message = "Something Went Wrong"
	}

	fmt.Printf("%+v\n", err.Source)
	statusCode := grpc.StatusCode(err)

	return status.New(statusCode, err.Description).Err()
}
