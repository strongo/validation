package validation

import "testing"

func TestIsValidationError(t *testing.T) {
	if !IsValidationError(errValidation) {
		t.Fatal("expected to return true for errValidation")
	}
}
