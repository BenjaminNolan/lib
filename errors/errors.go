package errors

import (
	goerrors "errors"
)

// ensure Error always implements the error interface
var _ error = (*Error)(nil)

// ErrorSlice is used to detect errors which have been through errors.Join()
type ErrorSlice interface {
	Unwrap() []error
}

// Error is a type which implements the error interface, and thus can be used to declare errors
// as constants.
//
// Example:
//
//	const ErrMyError errors.Error = "my error"
type Error string

// Error returns the error content as a string
func (e Error) Error() string {
	return string(e)
}

// New creates a new Error instance with the given message
func New(text string) error {
	return Error(text)
}

// Join combines multiple errors into one error slice. Nil entries are ignored.
//
// This func handles flattening nested Joined errors, which means that errors.Is and errors.As can
// actually detect them properly. In the default implementation, this doesn't work, because you end
// up with the joined error just included into the new slice.
func Join(errs ...error) error {
	var unwrappedErrs []error

	for _, err := range errs {
		if err == nil {
			continue
		}

		unwrappedErrs = append(unwrappedErrs, Split(err)...)
	}

	return goerrors.Join(unwrappedErrs...)
}

// Is passes through to errors.Is
func Is(err error, targetErr error) bool {
	return goerrors.Is(err, targetErr)
}

// As passes through to errors.As
func As(err error, targetErr any) bool {
	return goerrors.As(err, targetErr)
}

// Unwrap passes through to errors.Unwrap
func Unwrap(err error) error {
	return goerrors.Unwrap(err)
}

// Split converts any error (back) into a slice of errors
func Split(err error) []error {
	if e, ok := err.(ErrorSlice); ok {
		return e.Unwrap()
	}

	return []error{err}
}
