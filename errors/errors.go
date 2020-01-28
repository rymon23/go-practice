// Package errors provides error handling for gRPC services.
package errors

import (
	"errors"
	"fmt"
	"strings"
	"bitbucket.org/EMN/mulch/mulch"
	"google.golang.org/grpc/codes"
)

// New forwards to New in the standard errors package.
var New = errors.New

// Error identifies a composed error.
type Error struct {
	account *mulch.Account
	code    codes.Code
	message string
	err     error
}

// E creates a new error for the supplied arguments.
func E(args ...interface{}) error {
	err := &Error{}
	for _, arg := range args {
		switch arg := arg.(type) {
		case *mulch.Account:
			err.account = arg
		case codes.Code:
			err.code = arg
		case string:
			err.message = arg
		case error:
			err.err = arg
		case nil:
		default:
			return fmt.Errorf("unknown type %T (%v) in error call", arg, arg)
		}
	}
	return err
}

// Wrap provides compatability with pkg/errors.
func Wrap(err error, args ...interface{}) error {
	if err == nil {
		return nil
	}
	args = append(args, err)
	return E(args...)
}

// Code returns the gRPC status code relavant to the error.
func (e *Error) Code() codes.Code {
	if err, ok := e.err.(*Error); ok {
		return err.Code()
	}
	if e.code == 0 {
		return codes.Internal
	}
	return e.code
}

// Account returns the account associated with the error.
func (e *Error) Account() *mulch.Account {
	if e.account != nil {
		return e.account
	}
	if err, ok := e.err.(*Error); ok {
		return err.Account()
	}
	return nil
}

// Error returns the error message given by the composed error.
func (e *Error) Error() string {
	msg := e.error()
	if a := e.Account(); a != nil {
		return fmt.Sprintf("%s (%s): %s", a.Name, a.Email, msg)
	}
	return msg
}

// UserError returns the error message suitable for the remote client.
func (e *Error) UserError() string {
	return e.error()
}

// IsPublic returns true if the error can be considered consumable by the
// end user. Some errors may contain sensitive information should not be
// revealed to the client.
//
// An error is considered public for certain status codes or if the global
// Public variable was passed to the error creation function.
func (e *Error) IsPublic() bool {
	switch e.code {
	case codes.PermissionDenied:
		return true
	case codes.NotFound:
		return true
	case codes.InvalidArgument:
		return true
	case codes.FailedPrecondition:
		return true
	default:
		return false
	}
}

// error returns the standard error message shared by Error and UserError.
func (e *Error) error() string {
	parts := []string{}
	if e.message != "" {
		parts = append(parts, e.message)
	}
	if e.err != nil {
		if err, ok := e.err.(*Error); ok {
			parts = append(parts, err.error())
		} else {
			parts = append(parts, e.err.Error())
		}
	}
	return strings.Join(parts, ": ")
}

// Equal returns true of the supplied errors are equal. Package Error types
// will unwrap their wrapped values for comparison.
func Equal(a, b error) bool {
	if a == b {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	if err, ok := a.(*Error); ok {
		return Equal(err.err, b)
	}
	if err, ok := b.(*Error); ok {
		return Equal(a, err.err)
	}

	return false
}
