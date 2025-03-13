package errors

import (
	"fmt"
	"net/http"

	"github.com/hashicorp/go-multierror"
	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
)

// Error универсальная структура для ошибок с HTTP и gRPC статусами
// Error is a universal structure for errors with HTTP and gRPC statuses
type Error struct {
	Message  string
	HTTPCode int
	GRPCCode codes.Code
}

// Error возвращает текстовое представление ошибки
// creates a new error without a code
func (e *Error) Error() string {
	return e.Message
}

// New создает новую ошибку без кода
// New creates a new error with the not code
func New(message string) error {
	return &Error{Message: message}
}

// NewHTTP создает новую ошибку с заданным сообщением и HTTP-статусом
// NewHTTP creates a new error with the specified message and HTTP status
func NewHTTP(message string, httpCode int) error {
	return &Error{Message: message, HTTPCode: httpCode}
}

// NewGRPC создает новую ошибку с заданным сообщением и gRPC кодом
// NewGRPC creates a new error with the specified message and gRPC code
func NewGRPC(message string, grpcCode codes.Code) error {
	return &Error{Message: message, GRPCCode: grpcCode}
}

// Wrap обертывает ошибку с дополнительным сообщением, сохраняя код исходной ошибки
// Wrap wraps an error with an additional message, preserving the original error's code
func Wrap(err error, message string) error {
	httpCode := http.StatusInternalServerError
	grpcCode := codes.Unknown

	if e, ok := err.(*Error); ok {
		httpCode = e.HTTPCode
		grpcCode = e.GRPCCode
	}

	return &Error{
		Message:  fmt.Sprintf("%s: %v", message, err),
		HTTPCode: httpCode,
		GRPCCode: grpcCode,
	}
}

// HTTPStatus возвращает HTTP-статус ошибки
// HTTPStatus returns the HTTP status of the error
func HTTPStatus(err error) int {
	if e, ok := err.(*Error); ok {
		return e.HTTPCode
	}
	return http.StatusInternalServerError
}

// GRPCStatus возвращает gRPC статус ошибки
// GRPCStatus returns the gRPC status of the error
func GRPCStatus(err error) codes.Code {
	if e, ok := err.(*Error); ok {
		return e.GRPCCode
	}
	return codes.Unknown
}

// As проверяет, можно ли привести ошибку err к типу, на который указывает target
// As checks whether the error err can be cast to the type pointed to by target
func As(err error, target interface{}) bool {
	return errors.As(err, target)
}

// Is проверяет, является ли ошибка err или любая ошибка в ее цепочке эквивалентной target
// Is checks whether the error err or any error in its chain is equivalent to target
func Is(err, target error) bool {
	return errors.Is(err, target)
}

// Unwrap извлекает следующую ошибку в цепочке
// Unwrap extracts the next error in the chain
func Unwrap(err error) error {
	return errors.Unwrap(err)
}

// Append добавляет одну или несколько ошибок к существующей ошибке, сохраняя код исходной ошибки
// Append appends one or more errors to an existing error, preserving the original error's code
func Append(err error, errs ...error) *multierror.Error {
	httpCode := http.StatusInternalServerError
	grpcCode := codes.Unknown

	if e, ok := err.(*Error); ok {
		httpCode = e.HTTPCode
		grpcCode = e.GRPCCode
	}

	// Добавляем ошибки
	result := multierror.Append(err, errs...)

	// Сохраняем код исходной ошибки
	if result != nil {
		for _, e := range result.Errors {
			if errWithCode, ok := e.(*Error); ok {
				errWithCode.HTTPCode = httpCode
				errWithCode.GRPCCode = grpcCode
			}
		}
	}

	return result
}

// Flatten упрощает ошибку, преобразуя вложенные multierror.Error в плоский список ошибок
// Flatten flattens the error, converting nested multierror.Error into a flat list of errors
func Flatten(err error) error {
	return multierror.Flatten(err)
}

// Prefix добавляет префикс к каждой ошибке в multierror.Error, сохраняя код исходной ошибки
// Prefix adds a prefix to each error in multierror.Error, preserving the original error's code
func Prefix(err error, prefix string) error {
	httpCode := http.StatusInternalServerError
	grpcCode := codes.Unknown

	if e, ok := err.(*Error); ok {
		httpCode = e.HTTPCode
		grpcCode = e.GRPCCode
	}

	prefixedErr := multierror.Prefix(err, prefix)

	// Сохраняем код исходной ошибки
	if errWithCode, ok := prefixedErr.(*Error); ok {
		errWithCode.HTTPCode = httpCode
		errWithCode.GRPCCode = grpcCode
	}

	return prefixedErr
}

// Конструкторы для часто используемых HTTP ошибок

// BadRequest создает ошибку с HTTP-статусом 400 (Bad Request)
// BadRequest creates an error with HTTP status 400 (Bad Request)
func BadRequest(message string) error {
	return NewHTTP(message, http.StatusBadRequest)
}

// Unauthorized создает ошибку с HTTP-статусом 401
// Unauthorized creates an error with HTTP status 401
func Unauthorized(message string) error {
	return NewHTTP(message, http.StatusUnauthorized)
}

// Forbidden создает ошибку с HTTP-статусом 403
// Forbidden creates an error with HTTP status 403
func Forbidden(message string) error {
	return NewHTTP(message, http.StatusForbidden)
}

// NotFound создает ошибку с HTTP-статусом 404 (Not Found)
// NotFound creates an error with HTTP status 404 (Not Found)
func NotFound(message string) error {
	return NewHTTP(message, http.StatusNotFound)
}

// InternalServer создает ошибку с HTTP-статусом 500 (Internal Server Error)
// InternalServer creates an error with HTTP status 500 (Internal Server Error)
func InternalServer(message string) error {
	return NewHTTP(message, http.StatusInternalServerError)
}

// BadGateway создает ошибку с HTTP-статусом 502
// BadGateway creates an error with HTTP status 502
func BadGateway(message string) error {
	return NewHTTP(message, http.StatusBadGateway)
}

// Конструкторы для часто используемых gRPC ошибок

// GRPCNotFound создает ошибку с gRPC кодом NotFound
// GRPCNotFound creates an error with gRPC code NotFound
func GRPCNotFound(message string) error {
	return NewGRPC(message, codes.NotFound)
}

// GRPCInternal создает ошибку с gRPC кодом Internal
// GRPCInternal creates an error with gRPC code Internal
func GRPCInternal(message string) error {
	return NewGRPC(message, codes.Internal)
}

// GRPCInvalidArgument создает ошибку с gRPC кодом InvalidArgument
// GRPCInvalidArgument creates an error with gRPC code InvalidArgument
func GRPCInvalidArgument(message string) error {
	return NewGRPC(message, codes.InvalidArgument)
}

// GRPCUnavailable создает ошибку с gRPC кодом Unavailable
// GRPCUnavailable creates an error with gRPC code Unavailable
func GRPCUnavailable(message string) error {
	return NewGRPC(message, codes.Unavailable)
}
