package errors

import (
	"errors"
	"fmt"
	"google.golang.org/grpc/codes"
	"net/http"
	"testing"
)

// Тестирование конструкторов HTTP ошибок
func TestHTTPErrorConstructors(t *testing.T) {
	tests := []struct {
		name     string
		errorFn  func(string) error
		expected int
	}{
		{"BadRequest", BadRequest, http.StatusBadRequest},
		{"Unauthorized", Unauthorized, http.StatusUnauthorized},
		{"Forbidden", Forbidden, http.StatusForbidden},
		{"NotFound", NotFound, http.StatusNotFound},
		{"InternalServer", InternalServer, http.StatusInternalServerError},
		{"BadGateway", BadGateway, http.StatusBadGateway},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errorFn("test message")
			if HTTPStatus(err) != tt.expected {
				t.Errorf("%s: expected HTTP status %d, got %d", tt.name, tt.expected, HTTPStatus(err))
			}
		})
	}
}

// Тестирование конструкторов gRPC ошибок
func TestGRPCErrorConstructors(t *testing.T) {
	tests := []struct {
		name     string
		errorFn  func(string) error
		expected codes.Code
	}{
		{"GRPCNotFound", GRPCNotFound, codes.NotFound},
		{"GRPCInternal", GRPCInternal, codes.Internal},
		{"GRPCInvalidArgument", GRPCInvalidArgument, codes.InvalidArgument},
		{"GRPCUnavailable", GRPCUnavailable, codes.Unavailable},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.errorFn("test message")
			if GRPCStatus(err) != tt.expected {
				t.Errorf("%s: expected gRPC code %v, got %v", tt.name, tt.expected, GRPCStatus(err))
			}
		})
	}
}

// Тестирование функции New
func TestNew(t *testing.T) {
	err := New("test message")
	if err.Error() != "test message" {
		t.Errorf("New: expected message 'test message', got '%s'", err.Error())
	}
}

// Тестирование функции NewHTTP
func TestNewHTTP(t *testing.T) {
	err := NewHTTP("test message", http.StatusBadRequest)
	if err.Error() != "test message" {
		t.Errorf("New: expected message 'test message', got '%s'", err.Error())
	}
	if HTTPStatus(err) != http.StatusBadRequest {
		t.Errorf("New: expected HTTP status %d, got %d", http.StatusBadRequest, HTTPStatus(err))
	}
}

// Тестирование функции NewGRPC
func TestNewGRPC(t *testing.T) {
	err := NewGRPC("test message", codes.Internal)
	if err.Error() != "test message" {
		t.Errorf("NewGRPC: expected message 'test message', got '%s'", err.Error())
	}
	if GRPCStatus(err) != codes.Internal {
		t.Errorf("NewGRPC: expected gRPC code %v, got %v", codes.Internal, GRPCStatus(err))
	}
}

// Тестирование функции Wrap
func TestWrap(t *testing.T) {
	originalErr := NewHTTP("original error", http.StatusBadRequest)
	wrappedErr := Wrap(originalErr, "wrapped message")

	if wrappedErr.Error() != "wrapped message: original error" {
		t.Errorf("Wrap: expected message 'wrapped message: original error', got '%s'", wrappedErr.Error())
	}
	if HTTPStatus(wrappedErr) != http.StatusBadRequest {
		t.Errorf("Wrap: expected HTTP status %d, got %d", http.StatusBadRequest, HTTPStatus(wrappedErr))
	}
}

// Тестирование функции Append
func TestAppend(t *testing.T) {
	err1 := NewHTTP("error 1", http.StatusBadRequest)
	err2 := NewHTTP("error 2", http.StatusInternalServerError)
	combinedErr := Append(err1, err2)

	if len(combinedErr.Errors) != 2 {
		t.Errorf("Append: expected 2 errors, got %d", len(combinedErr.Errors))
	}
	for _, err := range combinedErr.Errors {
		if HTTPStatus(err) != http.StatusBadRequest {
			t.Errorf("Append: expected HTTP status %d for all errors, got %d", http.StatusBadRequest, HTTPStatus(err))
		}
	}
}

// Тестирование функции Flatten
func TestFlatten(t *testing.T) {
	err1 := NewHTTP("error 1", http.StatusBadRequest)
	err2 := NewHTTP("error 2", http.StatusInternalServerError)
	combinedErr := Append(err1, err2)
	flattenedErr := Flatten(combinedErr)

	if flattenedErr.Error() != "2 errors occurred:\n\t* error 1\n\t* error 2\n\n" {
		t.Errorf("Flatten: unexpected flattened error message: %s", flattenedErr.Error())
	}
}

// Тестирование функции Prefix
func TestPrefix(t *testing.T) {
	err := NewHTTP("error", http.StatusInternalServerError)
	prefixedErr := Prefix(err, "prefix")

	if prefixedErr.Error() != "prefix error" {
		t.Errorf("Prefix: expected message 'prefix: error', got '%s'", prefixedErr.Error())
	}
	if HTTPStatus(prefixedErr) != http.StatusInternalServerError {
		t.Errorf("Prefix: expected HTTP status %d, got %d", http.StatusInternalServerError, HTTPStatus(prefixedErr))
	}
}

// Тестирование функции HTTPStatus
func TestHTTPStatus(t *testing.T) {
	err := NewHTTP("test message", http.StatusNotFound)
	if HTTPStatus(err) != http.StatusNotFound {
		t.Errorf("HTTPStatus: expected status %d, got %d", http.StatusNotFound, HTTPStatus(err))
	}

	// Проверка для не типизированной ошибки
	nonTypedErr := errors.New("non-typed error")
	if HTTPStatus(nonTypedErr) != http.StatusInternalServerError {
		t.Errorf("HTTPStatus: expected default status %d for non-typed error, got %d", http.StatusInternalServerError, HTTPStatus(nonTypedErr))
	}
}

// Тестирование функции GRPCStatus
func TestGRPCStatus(t *testing.T) {
	err := NewGRPC("test message", codes.NotFound)
	if GRPCStatus(err) != codes.NotFound {
		t.Errorf("GRPCStatus: expected code %v, got %v", codes.NotFound, GRPCStatus(err))
	}

	// Проверка для не типизированной ошибки
	nonTypedErr := errors.New("non-typed error")
	if GRPCStatus(nonTypedErr) != codes.Unknown {
		t.Errorf("GRPCStatus: expected default code %v for non-typed error, got %v", codes.Unknown, GRPCStatus(nonTypedErr))
	}
}

// Тестирование функции As
func TestAs(t *testing.T) {
	err := NewHTTP("test message", http.StatusBadRequest)
	var target *Error
	if !As(err, &target) {
		t.Error("As: expected to cast error to *Error, but failed")
	}
}

// Тестирование функции Is
func TestIs(t *testing.T) {
	err := NewHTTP("test message", http.StatusBadRequest)
	if !Is(err, err) {
		t.Error("Is: expected error to be equal to itself, but it wasn't")
	}
}

// Тестирование функции Unwrap
func TestUnwrap(t *testing.T) {
	originalErr := errors.New("original error")
	wrappedErr := fmt.Errorf("wrapped: %w", originalErr)
	if Unwrap(wrappedErr) != originalErr {
		t.Error("Unwrap: expected to unwrap to original error, but failed")
	}
}
