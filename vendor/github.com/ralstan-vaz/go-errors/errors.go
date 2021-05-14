package errors

import (
	"fmt"
	"runtime"
	"runtime/debug"
	"strings"
)

// Kind defines the kind of error.
// This helps with grouping the errors at root level
type Kind string

// String converts the kind to a string value
func (k *Kind) String() string {
	return string(*k)
}

// Error contains all details with respect to the error
// Kind: Defines what category the error belong to
// Code: App level error code. Not necessary to have it. only requird for external calls
// Description: Is the description that will be shown to the end user. May not always be there.
// will be set to a default if it does not exist
// Message: Is the detailed message of the error
// Source: Contains the details from where the error originated
type Error struct {
	Kind        Kind
	Code        string
	Description string
	Message     string
	Source      Source
}

// Source Contains information regarding the source of the error
// Caller: Contains the method name of the caller
// File: Contains the file path in which the error originated
// Line: Contains the line number in the file
// StackTrace: Contains the complete stack trace of the request from the error
// Error: Contains the original error
type Source struct {
	Caller     string `json:"caller,omitempty"`
	File       string `json:"file,omitempty"`
	Line       int    `json:"line,omitempty"`
	StackTrace string `json:"stackTrace,omitempty"`
	Error      error  `json:"error,omitempty"`
}

// Get gets the error struct in return
func (e *Error) Get() *Error {
	return e
}

// Error return the error in the form of Kind and description
func (e *Error) Error() string {
	return fmt.Sprintf("%s: %s", e.Kind, e.Description)
}

// Wrap wraps the given error in error struct
func (e *Error) Wrap(err error) *Error {
	e.Source.Error = err
	return e
}

// SetCode sets an error code for the current error
func (e *Error) SetCode(code string) *Error {
	e.Code = code
	return e
}

type getter interface {
	Get() *Error
}

// Stop !!
// Create new error Kinds only if it does not fit in the below cases

// List of all the Error Kinds
const (
	NotFound         Kind = "NotFound"
	Unauthorized          = "Unauthorized"
	Forbidden             = "Forbidden"
	Expired               = "Expired"
	BadRequest            = "BadRequest"
	ParameterMissing      = "ParameterMissing"
	InternalError         = "InternalError"
	Unknown               = "Unknown"
)

// Get return the corresponding error struct
// If the error is not of the custom struct it will return InternalError kind by default
func Get(err error) *Error {
	switch t := err.(type) {
	case getter:
		return t.Get()
	}

	return NewInternalError(err)
}

// New creates an instance of an error
func New(err Error) *Error {
	return newError(err.Kind, err.Description, err.Code, err.Message)
}

// NewNotFound  creates an instance of a NotFound error
func NewNotFound(description string) *Error {
	return newError(NotFound, description, "", "")
}

// NewUnauthorized creates an instance of an Unauthorized error
func NewUnauthorized(description string) *Error {
	return newError(Unauthorized, description, "", "")
}

// NewForbidden creates an instance of a Forbidden error
func NewForbidden(description string) *Error {
	return newError(Forbidden, description, "", "")
}

// NewExpired creates an instance of an Expired error
func NewExpired(description string) *Error {
	return newError(Expired, description, "", "")
}

// NewBadRequest creates an instance of a BadRequest error
func NewBadRequest(description string) *Error {
	return newError(BadRequest, description, "", "")
}

// NewParameterMissing creates an instance of a ParameterMissing error
func NewParameterMissing(description string) *Error {
	return newError(ParameterMissing, description, "", "")
}

// NewInternalError creates an instance of an InternalError error
func NewInternalError(err error) *Error {
	internalErr := newError(InternalError, err.Error(), "", "")
	internalErr.Wrap(err)
	return internalErr
}

// Error validations

// IsNotFound validated if the error is not found
func IsNotFound(err error) bool {
	return getKind(err) == NotFound
}

// IsUnauthorized validated if the error is Unauthorized
func IsUnauthorized(err error) bool {
	return getKind(err) == Unauthorized
}

// IsForbidden validated if the error is Forbidden
func IsForbidden(err error) bool {
	return getKind(err) == Forbidden
}

// IsExpired validated if the error is of type Expired
func IsExpired(err error) bool {
	return getKind(err) == Expired
}

// IsBadRequest validated if the error is of type BadRequest
func IsBadRequest(err error) bool {
	return getKind(err) == BadRequest
}

// IsParameterMissing validated if the error is of type ParameterMissing
func IsParameterMissing(err error) bool {
	return getKind(err) == ParameterMissing
}

// IsInternalError validated if the error is of type InternalError
func IsInternalError(err error) bool {
	return getKind(err) == InternalError
}

// getKind returns the error type
func getKind(err error) Kind {
	switch t := err.(type) {
	case getter:
		return t.Get().Kind
	}
	return Unknown
}

// newError creates a new instance of an error
func newError(kind Kind, description, code, message string) *Error {
	caller, file, line := getCallerDetails(3)

	errorSource := Source{}
	errorSource.StackTrace = string(debug.Stack())
	errorSource.File = file
	errorSource.Line = line
	errorSource.Caller = caller

	if code == "" {
		code = kind.String()
	}

	return &Error{
		Kind:        kind,
		Description: description,
		Code:        code,
		Message:     message,
		Source:      errorSource,
	}
}

// getCallerDetails gets the callerName, filePath and the lineNo
func getCallerDetails(skip int) (string, string, int) {
	pc, file, line, _ := runtime.Caller(skip)

	fn := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	caller := fn[len(fn)-1]

	return caller, file, line
}
