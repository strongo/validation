package validation

import (
	"errors"
	"testing"
)

func TestNewValidationError(t *testing.T) {
	t.Run("should fail", func(t *testing.T) {
		t.Run("empty message", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Error("expected to panic")
				}
			}()
			_ = NewValidationError("")
		})
	})
	t.Run("should pass", func(t *testing.T) {
		err := NewValidationError("test_err_1")
		if !IsValidationError(err) {
			t.Error("expected to return true from IsValidationError()")
		}
	})
}

func TestIsValidationError(t *testing.T) {
	t.Run("errValidation", func(t *testing.T) {
		if !IsValidationError(errValidation) {
			t.Fatal("expected to return true for errValidation")
		}
	})
	t.Run("errBadRecord", func(t *testing.T) {
		if !IsValidationError(errBadRecord) {
			t.Fatal("expected to return true for errBadRecord")
		}
	})
	t.Run("errBadRequest", func(t *testing.T) {
		if !IsValidationError(errBadRequest) {
			t.Fatal("expected to return true for errBadRequest")
		}
	})
}

func TestIsBadRecordError(t *testing.T) {
	t.Run("should fail", func(t *testing.T) {
		if IsBadRecordError(errors.New("some random error")) {
			t.Error("should fail for random error")
		}
		t.Run("errBadRequest", func(t *testing.T) {
			if IsBadRecordError(errBadRequest) {
				t.Error("should be false")
			}
		})
	})
	t.Run("should pass for", func(t *testing.T) {
		t.Run("errBadRecord", func(t *testing.T) {
			if !IsBadRecordError(errBadRecord) {
				t.Error("should be true")
			}
		})
		t.Run("NewErrRecordIsMissingRequiredField", func(t *testing.T) {
			if !IsBadRecordError(NewErrRecordIsMissingRequiredField("field1")) {
				t.Error("should be true")
			}
		})
		t.Run("NewErrBadRecordFieldValue", func(t *testing.T) {
			if !IsBadRecordError(NewErrBadRecordFieldValue("field1", "msg1")) {
				t.Error("should be true")
			}
		})
	})
}

func TestIsBadRequestError(t *testing.T) {
	t.Run("should fail", func(t *testing.T) {
		if IsBadRequestError(errors.New("some random error")) {
			t.Error("should fail for random error")
		}
		t.Run("errBadRecord", func(t *testing.T) {
			if IsBadRequestError(errBadRecord) {
				t.Error("should be false")
			}
		})
	})
	t.Run("should pass for", func(t *testing.T) {
		t.Run("errBadRequest", func(t *testing.T) {
			if !IsBadRequestError(errBadRequest) {
				t.Error("should be true")
			}
		})
		t.Run("NewErrRequestIsMissingRequiredField", func(t *testing.T) {
			if !IsBadRequestError(NewErrRequestIsMissingRequiredField("field1")) {
				t.Error("should be true")
			}
		})
		t.Run("NewErrBadRequestFieldValue", func(t *testing.T) {
			if !IsBadRequestError(NewErrBadRequestFieldValue("field1", "msg1")) {
				t.Error("should be true")
			}
		})
	})
}

func TestNewErrBadRecordFieldValue(t *testing.T) {
	t.Run("should panic", func(t *testing.T) {
		t.Run("empty name", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Error("expected to panic")
				}
			}()
			_ = NewErrBadRecordFieldValue("", "bad_field_123")
		})
		t.Run("empty message", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Error("expected to panic")
				}
			}()
			_ = NewErrBadRecordFieldValue("field1", "")
		})
	})
	t.Run("should pass", func(t *testing.T) {
		err := NewErrBadRecordFieldValue("field1", "bad_field_123")
		mustBeValidationError(t, err)
		mustBeBadFieldError(t, err)
		//
		mustBeBadRecordError(t, err)
	})
}

func TestNewErrBadRequestFieldValue(t *testing.T) {
	t.Run("should panic", func(t *testing.T) {
		t.Run("empty name", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Error("expected to panic")
				}
			}()
			_ = NewErrBadRequestFieldValue("", "bad_field_123")
		})
		t.Run("empty message", func(t *testing.T) {
			defer func() {
				if err := recover(); err == nil {
					t.Error("expected to panic")
				}
			}()
			_ = NewErrBadRequestFieldValue("field1", "")
		})
	})
	t.Run("should pass", func(t *testing.T) {
		err := NewErrBadRequestFieldValue("field1", "bad_field_123")
		mustBeValidationError(t, err)
		mustBeBadFieldError(t, err)
		//
		mustBeBadRequestError(t, err)
	})
}

func mustBeValidationError(t *testing.T, err error) {
	t.Helper()
	if !IsValidationError(err) {
		t.Errorf("expected to be IsValidationError, got %T: %v", err, err)
	}
}

func mustBeBadRecordError(t *testing.T, err error) {
	t.Helper()
	if !IsBadRecordError(err) {
		t.Errorf("expected to be IsBadRecordError, got %T: %v", err, err)
	}
}

func mustBeBadRequestError(t *testing.T, err error) {
	t.Helper()
	if !IsBadRequestError(err) {
		t.Errorf("expected to be IsBadRequestError, got %T: %v", err, err)
	}
}

func mustBeBadFieldError(t *testing.T, err error) {
	t.Helper()
	if !IsBadFieldValueError(err) {
		t.Errorf("expected to be IsBadFieldValueError, got %T: %v", err, err)
	}
}

type TestMock struct {
	helper int
	errorf []struct {
		format string
		args   []interface{}
	}
}

func (v *TestMock) Helper() {
	v.helper += 1
}

func (v *TestMock) Errorf(format string, args ...interface{}) {
	if format == "" {
		panic("format is a required parameter")
	}
	v.errorf = append(v.errorf, struct {
		format string
		args   []interface{}
	}{format: format, args: args})
}

func testHelperShouldBeCalledOnce(t *testing.T, testMock *TestMock) {
	t.Helper()
	if testMock.helper != 1 {
		t.Errorf("t.Helper() should be called once, got called %v times", testMock.helper)
	}
}
func TestMustBeFieldError(t *testing.T) {
	t.Run("should write a single error", func(t *testing.T) {
		testMock := new(TestMock)
		MustBeFieldError(testMock, errors.New("some error"), "field1")
		testHelperShouldBeCalledOnce(t, testMock)
		if len(testMock.errorf) != 1 {
			t.Errorf("t.Errorf() should be called once, got called %v times", len(testMock.errorf))
		}
	})
	t.Run("should not write any errors", func(t *testing.T) {
		testMock := new(TestMock)
		MustBeFieldError(testMock, NewErrRecordIsMissingRequiredField("field1"), "field1")
		testHelperShouldBeCalledOnce(t, testMock)
		if len(testMock.errorf) != 0 {
			t.Errorf("t.Errorf() should NOT be called, got called %v times", len(testMock.errorf))
		}
	})
}
