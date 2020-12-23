package validation

// github.com/strongo/validation

import (
	"errors"
	"fmt"
	"strings"
)

var errValidation = errors.New("validation error")

var errBadRequest = fmt.Errorf("%w: invalid request", errValidation)
var errBadRequestFieldValue = fmt.Errorf("%w: bad field value", errBadRequest)

var errBadRecord = fmt.Errorf("%w: invalid record", errValidation)
var errBadRecordFieldValue = fmt.Errorf("%w: invalid field value", errBadRecord)

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
		panic("field parameters is required")
	}
	if strings.TrimSpace(message) == "" {
		panic("message parameters is required")
	}
	return ErrBadFieldValue{
		Field:   field,
		Message: message,
		err:     errBadRequestFieldValue,
	}
}

// NewErrBadRecordFieldValue creates an error for a bad record field value
func NewErrBadRecordFieldValue(field, message string) ErrBadFieldValue {
	if field == "" {
		panic("field is a required parameter")
	}
	if message == "" {
		panic("message is a required parameter")
	}
	return ErrBadFieldValue{
		Field:   field,
		Message: message,
		err:     errBadRecordFieldValue,
	}
}

// NewValidationError creates a common validation error
func NewValidationError(message string) error {
	if message == "" {
		panic("message is a required parameter")
	}
	return fmt.Errorf("%w: %v", errValidation, message)
}

// IsValidationError checks if provided errors is a validation error
func IsValidationError(err error) bool {
	return errors.Is(err, errValidation)
}

// IsBadRequestError checks if provided errors is a validation error
func IsBadRequestError(err error) bool {
	return errors.Is(err, errBadRequest)
}

func IsBadFieldValueError(err error) bool {
	if errors.Is(err, errBadRequestFieldValue) {
		return true
	}
	if errors.Is(err, errBadRecordFieldValue) {
		return true
	}
	return false
}

// IsBadRecordError checks if provided errors is a validation error
func IsBadRecordError(err error) bool {
	return errors.Is(err, errBadRecord)
}

// NewErrRecordIsMissingRequiredField creates an error for a missing required field in a record
func NewErrRecordIsMissingRequiredField(field string) ErrBadFieldValue {
	return NewErrBadRecordFieldValue(field, "missing required field")
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
