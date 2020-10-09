package strongo_validation

import (
	"errors"
	"fmt"
	"strings"
)

var errValidation = errors.New("validation error")
var errBadRequest = fmt.Errorf("%w: bad request", errValidation)
var errBadRecord = fmt.Errorf("%w: bad record", errValidation)

// NewBadRequestError create error
func NewBadRequestError(err error) error {
	return fmt.Errorf("%w: %s", errBadRequest, err)
}

// ErrBadFieldValue reports error for a bad field value
type ErrBadFieldValue struct {
	err     error
	Field   string `json:"field"`
	Message string `json:"message"`
}

func (v ErrBadFieldValue) Error() string {
	return fmt.Sprintf("bad value for field [%v]: %v", v.Field, v.Message)
}

func (v ErrBadFieldValue) Unwrap() error {
	return v.err
}

// NewErrBadRequestFieldValue creates an error for a bad request field value
func NewErrBadRequestFieldValue(field, message string) ErrBadFieldValue {
	if strings.TrimSpace(field) == "" {
		panic(fmt.Sprintf("field parameters is empty string, message: %v", message))
	}
	return ErrBadFieldValue{
		Field:   field,
		Message: message,
		err:     errBadRequest,
	}
}

// NewErrBadRecordFieldValue creates an error for a bad record field value
func NewErrBadRecordFieldValue(field, message string) *ErrBadFieldValue {
	return &ErrBadFieldValue{
		Field:   field,
		Message: message,
		err:     errBadRecord,
	}
}

// NewValidationError creates a common validation error
func NewValidationError(err error, message string) error {
	if message == "" {
		message = errValidation.Error()
	}
	return fmt.Errorf("%v: %w", message, err)
}

// IsValidationError checks if provided errors is a validation error
func IsValidationError(err error) bool {
	return errors.Is(err, errValidation)
}

// IsBadRequestError checks if provided errors is a validation error
func IsBadRequestError(err error) bool {
	return errors.Is(err, errBadRequest)
}

// IsBadRecordError checks if provided errors is a validation error
func IsBadRecordError(err error) bool {
	return errors.Is(err, errBadRecord)
}

// NewErrRecordIsMissingRequiredField creates an error for a missing required field in a record
func NewErrRecordIsMissingRequiredField(field string) ErrBadFieldValue {
	return NewErrBadRequestFieldValue(field, "missing required field")
}

// NewErrRequestIsMissingRequiredField creates an error for a missing required field in a request
func NewErrRequestIsMissingRequiredField(field string) error {
	return NewBadRequestError(NewErrBadRequestFieldValue(field, "missing required field"))
}

// MustBeFieldError makes sure the provided error is a field error
func MustBeFieldError(t interface {
	Helper()
	Errorf(format string, args ...interface{})
}, err error, field string) {
	t.Helper()
	if err == nil {
		t.Errorf("Should return validation error for field [%v], got err == nil", field)
	} else if !IsValidationError(err) {
		t.Errorf("Must be validation error, got: %v", err)
	} else {
		errStr := err.Error()
		if !strings.Contains(errStr, fmt.Sprintf("[%v]", field)) {
			t.Errorf("Should mention [team] field, got: %v", errStr)
		}
	}
}
