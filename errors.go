package errors

import (
	"context"
	"github.com/pkg/errors"
	"strings"
)

// Error универсальная структура для ошибок с HTTP и gRPC статусами
// Error is a universal structure for errors with HTTP and gRPC statuses
type Error struct {
	Message      string
	code         Code
	typeProtocol ProtocolType
}

// Error возвращает текстовое представление ошибки
// Error returns a text representation of the error
func (e *Error) Error() string {
	return e.Message
}

// StatusHTTP возвращает HTTP-статус ошибки
// StatusHTTP returns the HTTP status of the error
func StatusHTTP(err error) int {
	if err == nil {
		return int(hTTPOk)
	}

	var er *Error
	if errors.As(err, &er) {
		if err.(*Error).typeProtocol == httpProtocol {
			return int(err.(*Error).code)
		}
		return int(statusGRPCToHTTP(err.(*Error).code))
	}

	// Handling other types of errors
	if errors.Is(err, context.DeadlineExceeded) {
		return int(hTTPGatewayTimeout)
	}
	if errors.Is(err, context.Canceled) {
		return int(hTTPRequestTimeout)
	}

	// By default, we return 500 Internal Server Error
	return int(hTTPInternalServerError)
}

// StatusGRPC возвращает gRPC статус ошибки
// StatusGRPC returns the gRPC status of the error
func StatusGRPC(err error) Code {
	if err == nil {
		return gRPCOk
	}

	var er *Error
	if errors.As(err, &er) {
		if err.(*Error).typeProtocol == grpcProtocol {
			return err.(*Error).code
		}
		return statusHTTPToGRPC(err.(*Error).code)
	}

	// Handling other types of errors
	if errors.Is(err, context.DeadlineExceeded) {
		return gRPCDeadlineExceeded
	}
	if errors.Is(err, context.Canceled) {
		return gRPCCanceled
	}

	// By default, we return codes.Unknown
	return gRPCUnknown
}

// Wrap обертывает ошибку с дополнительным сообщением, сохраняя код исходной ошибки
// Wrap wraps an error with an additional message, preserving the original error's code
func Wrap(err error, message string) error {
	var er *Error
	var builder strings.Builder
	builder.WriteString(message + ":")
	builder.WriteString(err.Error())

	if errors.As(err, &er) {
		return &Error{
			Message:      builder.String(),
			code:         err.(*Error).code,
			typeProtocol: err.(*Error).typeProtocol,
		}
	}

	return &Error{Message: builder.String()}
}
