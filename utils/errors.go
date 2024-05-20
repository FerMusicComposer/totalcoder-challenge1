package utils

import (
	"fmt"
)

type ErrorType int

const (
	_ ErrorType = iota
	errInvalidDate
	errUnauthorized
	errInternal
	errNotFound
	errForbidden
	errBadRequest
	errNotAllowed
	errInvalidPayload
)

type AppError struct {
	ErrType ErrorType
	Code    int
	Err     error
}

var (
	InvalidDate    = AppError{ErrType: errInvalidDate, Code: -1}
	Unauthorized   = AppError{ErrType: errUnauthorized, Code: 401}
	Internal       = AppError{ErrType: errInternal, Code: 500}
	NotFound       = AppError{ErrType: errNotFound, Code: 404}
	Forbidden      = AppError{ErrType: errForbidden, Code: 403}
	BadRequest     = AppError{ErrType: errBadRequest, Code: 400}
	NotAllowed     = AppError{ErrType: errNotAllowed, Code: 405}
	InvalidPayload = AppError{ErrType: errInvalidPayload, Code: 400}
)

// =================
// API ERROR CONFIG
// =================

func (appErr AppError) Error() string {
	switch appErr.ErrType {
	case errInvalidDate:
		return "The provided date is invalid. Check the format"
	case errUnauthorized:
		return fmt.Sprintf("Access denied: %d Unauthorized", appErr.Code)
	case errInternal:
		return "Internal server error"
	case errNotFound:
		return fmt.Sprintf("%d No records match the query parameters", appErr.Code)
	case errForbidden:
		return fmt.Sprintf("%d Forbidden", appErr.Code)
	case errBadRequest:
		return fmt.Sprintf("%d Bad request", appErr.Code)
	case errNotAllowed:
		return fmt.Sprintf("%d Method not allowed", appErr.Code)
	case errInvalidPayload:
		return fmt.Sprintf("%d Invalid payload", appErr.Code)
	default:
		return "Unknown error"
	}
}

//===============
// ERROR HELPERS
//===============

// with returns an error with a particular code (e.g unauthorized)
func (appErr AppError) With(code int) AppError {
	err := appErr
	err.Code = code
	return err
}

func (appErr AppError) From(location string, err error) AppError {
	errFrom := appErr
	errFrom.Err = fmt.Errorf("from %s: %v", location, err)
	return errFrom
}

func (appErr AppError) FromNoErr(location string) AppError {
	errFrom := appErr
	errFrom.Err = fmt.Errorf("from %s", location)
	return errFrom
}

func (appErr *AppError) Unwrap() error {
	return appErr.Err
}

func (appErr *AppError) Is(target error) bool {
	err, ok := target.(*AppError) // reflection to check if target is an AppError

	if !ok {
		return false
	}

	return appErr.ErrType == err.ErrType
}
