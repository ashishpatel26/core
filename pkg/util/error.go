package util

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/status"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
)

// UserError implements a new error type for user facing errors
type UserError struct {
	Code    codes.Code
	Message string
}

// Error returns error messages
func (e *UserError) Error() string {
	return e.Message
}

// GRPCStatus is used by gRPC to return the correct gRPC status codes
func (e *UserError) GRPCStatus() *status.Status {
	return status.New(e.Code, e.Message)
}

// NewUserError returns an instance of UserError with the appropriate code and message
func NewUserError(code codes.Code, message string) error {
	return &UserError{Code: code, Message: message}
}

func pqError(err *pq.Error) (code codes.Code) {
	switch err.Code {
	case "23505":
		code = codes.AlreadyExists
	default:
		code = codes.Unknown
	}
	return
}

// NewUserErrorWrap wraps pq errors and returns an instance of UserError
func NewUserErrorWrap(err error, entity string) error {
	var (
		code    codes.Code
		message string
		pqErr   *pq.Error
		userErr *UserError
	)
	if errors.As(err, &pqErr) {
		code = pqError(pqErr)
		message = fmt.Sprintf("%v already exists.", entity)
	} else if errors.As(err, &userErr) {
		return err
	} else {
		code = codes.Unknown
		message = "Unknown error."
	}

	return NewUserError(code, message)
}
