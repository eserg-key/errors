package errors

import (
	"context"
	"testing"
)

func TestNewHTTPError(t *testing.T) {
	err := BadRequestHTTP("Bad request")
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if err.Error() != "Bad request" {
		t.Errorf("Expected error message 'Bad request', got '%s'", err.Error())
	}

	if StatusHTTP(err) != int(hTTPBadRequest) {
		t.Errorf("Expected HTTP status 400, got %d", StatusHTTP(err))
	}
}

func TestNewGRPCError(t *testing.T) {
	err := InvalidArgumentGRPC("Invalid argument")
	if err == nil {
		t.Error("Expected an error, got nil")
	}

	if err.Error() != "Invalid argument" {
		t.Errorf("Expected error message 'Invalid argument', got '%s'", err.Error())
	}

	if StatusGRPC(err) != gRPCInvalidArgument {
		t.Errorf("Expected gRPC status 3, got %d", StatusGRPC(err))
	}
}

func TestWrap(t *testing.T) {
	originalErr := BadRequestHTTP("Invalid input")
	wrappedErr := Wrap(originalErr, "Validation failed")

	if wrappedErr.Error() != "Validation failed:Invalid input" {
		t.Errorf("Expected wrapped error message 'Validation failed:Invalid input', got '%s'", wrappedErr.Error())
	}

	if StatusHTTP(wrappedErr) != int(hTTPBadRequest) {
		t.Errorf("Expected HTTP status 400 after wrapping, got %d", StatusHTTP(wrappedErr))
	}
}

func TestStatusHTTP(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected int
	}{
		{"HTTP Bad Request", BadRequestHTTP("Bad request"), int(hTTPBadRequest)},
		{"gRPC Invalid Argument", InvalidArgumentGRPC("Invalid argument"), int(hTTPBadRequest)},
		{"Context Deadline Exceeded", context.DeadlineExceeded, int(hTTPGatewayTimeout)},
		{"Context Canceled", context.Canceled, int(hTTPRequestTimeout)},
		{"Nil Error", nil, int(hTTPOk)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := StatusHTTP(tt.err)
			if status != tt.expected {
				t.Errorf("Expected status %d, got %d", tt.expected, status)
			}
		})
	}
}

func TestStatusGRPC(t *testing.T) {
	tests := []struct {
		name     string
		err      error
		expected Code
	}{
		{"gRPC Invalid Argument", InvalidArgumentGRPC("Invalid argument"), gRPCInvalidArgument},
		{"HTTP Bad Request", BadRequestHTTP("Bad request"), gRPCInvalidArgument},
		{"Context Deadline Exceeded", context.DeadlineExceeded, gRPCDeadlineExceeded},
		{"Context Canceled", context.Canceled, gRPCCanceled},
		{"Nil Error", nil, gRPCOk},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			status := StatusGRPC(tt.err)
			if status != tt.expected {
				t.Errorf("Expected status %d, got %d", tt.expected, status)
			}
		})
	}
}

func TestStatusHTTPToGRPC(t *testing.T) {
	tests := []struct {
		name     string
		httpCode Code
		expected Code
	}{
		{"HTTP Bad Request", hTTPBadRequest, gRPCInvalidArgument},
		{"HTTP Unauthorized", hTTPUnauthorized, gRPCUnauthenticated},
		{"HTTP Internal Server Error", hTTPInternalServerError, gRPCInternal},
		{"Unknown HTTP Code", 999, gRPCUnknown},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			grpcCode := statusHTTPToGRPC(tt.httpCode)
			if grpcCode != tt.expected {
				t.Errorf("Expected gRPC code %d for HTTP code %d, got %d", tt.expected, tt.httpCode, grpcCode)
			}
		})
	}
}

func TestStatusGRPCToHTTP(t *testing.T) {
	tests := []struct {
		name     string
		grpcCode Code
		expected Code
	}{
		{"gRPC Invalid Argument", gRPCInvalidArgument, hTTPBadRequest},
		{"gRPC Unauthenticated", gRPCUnauthenticated, hTTPUnauthorized},
		{"gRPC Internal", gRPCInternal, hTTPInternalServerError},
		{"Unknown gRPC Code", 999, hTTPInternalServerError},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpCode := statusGRPCToHTTP(tt.grpcCode)
			if httpCode != tt.expected {
				t.Errorf("Expected HTTP code %d for gRPC code %d, got %d", tt.expected, tt.grpcCode, httpCode)
			}
		})
	}
}

func TestContextErrors(t *testing.T) {
	tests := []struct {
		name         string
		err          error
		expectedHTTP int
		expectedGRPC Code
	}{
		{"Deadline Exceeded", context.DeadlineExceeded, int(hTTPGatewayTimeout), gRPCDeadlineExceeded},
		{"Context Canceled", context.Canceled, int(hTTPRequestTimeout), gRPCCanceled},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			httpStatus := StatusHTTP(tt.err)
			if httpStatus != tt.expectedHTTP {
				t.Errorf("Expected HTTP status %d, got %d", tt.expectedHTTP, httpStatus)
			}

			grpcStatus := StatusGRPC(tt.err)
			if grpcStatus != tt.expectedGRPC {
				t.Errorf("Expected gRPC status %d, got %d", tt.expectedGRPC, grpcStatus)
			}
		})
	}
}
