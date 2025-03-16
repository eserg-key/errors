package errors

// newHTTP создает новую ошибку с заданным сообщением и HTTP-статусом
// newHTTP creates a new error with the specified message and HTTP status
func newHTTPError(message string, code Code) error {
	return &Error{Message: message, code: code, typeProtocol: httpProtocol}
}

// BadRequestHTTP создает ошибку с HTTP-статусом 400 (Bad Request)
// BadRequestHTTP creates an error with HTTP status 400 (Bad Request)
func BadRequestHTTP(message string) error {
	return newHTTPError(message, hTTPBadRequest)
}

// UnauthorizedHTTP создает ошибку с HTTP-статусом 401 (Unauthorized)
// UnauthorizedHTTP creates an error with HTTP status 401 (Unauthorized)
func UnauthorizedHTTP(message string) error {
	return newHTTPError(message, hTTPUnauthorized)
}

// PaymentRequiredHTTP создает ошибку с HTTP-статусом 402 (Payment Required)
// PaymentRequiredHTTP creates an error with HTTP status 402 (Payment Required)
func PaymentRequiredHTTP(message string) error {
	return newHTTPError(message, hTTPPaymentRequired)
}

// ForbiddenHTTP создает ошибку с HTTP-статусом 403 (Forbidden)
// ForbiddenHTTP creates an error with HTTP status 403 (Forbidden)
func ForbiddenHTTP(message string) error {
	return newHTTPError(message, hTTPForbidden)
}

// NotFoundHTTP создает ошибку с HTTP-статусом 404 (Not Found)
// NotFoundHTTP creates an error with HTTP status 404 (Not Found)
func NotFoundHTTP(message string) error {
	return newHTTPError(message, hTTPNotFound)
}

// MethodNotAllowedHTTP создает ошибку с HTTP-статусом 405 (Method Not Allowed)
// MethodNotAllowedHTTP creates an error with HTTP status 405 (Method Not Allowed)
func MethodNotAllowedHTTP(message string) error {
	return newHTTPError(message, hTTPMethodNotAllowed)
}

// NotAcceptableHTTP создает ошибку с HTTP-статусом 406 (Not Acceptable)
// NotAcceptableHTTP creates an error with HTTP status 406 (Not Acceptable)
func NotAcceptableHTTP(message string) error {
	return newHTTPError(message, hTTPNotAcceptable)
}

// ProxyAuthRequiredHTTP создает ошибку с HTTP-статусом 407 (Proxy Authentication Required)
// ProxyAuthRequiredHTTP creates an error with HTTP status 407 (Proxy Authentication Required)
func ProxyAuthRequiredHTTP(message string) error {
	return newHTTPError(message, hTTPStatusProxyAuthRequired)
}

// RequestTimeoutHTTP создает ошибку с HTTP-статусом 408 (Request Timeout)
// RequestTimeoutHTTP creates an error with HTTP status 408 (Request Timeout)
func RequestTimeoutHTTP(message string) error {
	return newHTTPError(message, hTTPRequestTimeout)
}

// ConflictHTTP создает ошибку с HTTP-статусом 409 (Conflict)
// ConflictHTTP creates an error with HTTP status 409 (Conflict)
func ConflictHTTP(message string) error {
	return newHTTPError(message, hTTPConflict)
}

// GoneHTTP создает ошибку с HTTP-статусом 410 (Gone)
// GoneHTTP creates an error with HTTP status 410 (Gone)
func GoneHTTP(message string) error {
	return newHTTPError(message, hTTPGone)
}

// LengthRequiredHTTP создает ошибку с HTTP-статусом 411 (Length Required)
// LengthRequiredHTTP creates an error with HTTP status 411 (Length Required)
func LengthRequiredHTTP(message string) error {
	return newHTTPError(message, hTTPStatusLengthRequired)
}

// PreconditionFailedHTTP создает ошибку с HTTP-статусом 412 (Precondition Failed)
// PreconditionFailedHTTP creates an error with HTTP status 412 (Precondition Failed)
func PreconditionFailedHTTP(message string) error {
	return newHTTPError(message, hTTPStatusPreconditionFailed)
}

// PayloadTooLargeHTTP создает ошибку с HTTP-статусом 413 (Payload Too Large)
// PayloadTooLargeHTTP creates an error with HTTP status 413 (Payload Too Large)
func PayloadTooLargeHTTP(message string) error {
	return newHTTPError(message, hTTPPayloadTooLarge)
}

// URITooLongHTTP создает ошибку с HTTP-статусом 414 (URI Too Long)
// URITooLongHTTP creates an error with HTTP status 414 (URI Too Long)
func URITooLongHTTP(message string) error {
	return newHTTPError(message, hTTPURITooLong)
}

// UnsupportedMediaTypeHTTP создает ошибку с HTTP-статусом 415 (Unsupported Media Type)
// UnsupportedMediaTypeHTTP creates an error with HTTP status 415 (Unsupported Media Type)
func UnsupportedMediaTypeHTTP(message string) error {
	return newHTTPError(message, hTTPUnsupportedMediaType)
}

// RangeNotSatisfiableHTTP создает ошибку с HTTP-статусом 416 (Range Not Satisfiable)
// RangeNotSatisfiableHTTP creates an error with HTTP status 416 (Range Not Satisfiable)
func RangeNotSatisfiableHTTP(message string) error {
	return newHTTPError(message, hTTPStatusRequestedRangeNotSatisfiable)
}

// ExpectationFailedHTTP создает ошибку с HTTP-статусом 417 (Expectation Failed)
// ExpectationFailedHTTP creates an error with HTTP status 417 (Expectation Failed)
func ExpectationFailedHTTP(message string) error {
	return newHTTPError(message, hTTPStatusExpectationFailed)
}

// TeapotHTTP создает ошибку с HTTP-статусом 418 (I'm a teapot)
// TeapotHTTP creates an error with HTTP status 418 (I'm a teapot)
func TeapotHTTP(message string) error {
	return newHTTPError(message, hTTPStatusTeapot)
}

// UnprocessableEntityHTTP создает ошибку с HTTP-статусом 422 (Unprocessable Entity)
// UnprocessableEntityHTTP creates an error with HTTP status 422 (Unprocessable Entity)
func UnprocessableEntityHTTP(message string) error {
	return newHTTPError(message, hTTPUnprocessableEntity)
}

// TooManyRequestsHTTP создает ошибку с HTTP-статусом 429 (Too Many Requests)
// TooManyRequestsHTTP creates an error with HTTP status 429 (Too Many Requests)
func TooManyRequestsHTTP(message string) error {
	return newHTTPError(message, hTTPTooManyRequests)
}

// InternalServerHTTP создает ошибку с HTTP-статусом 500 (Internal Server Error)
// InternalServerHTTP creates an error with HTTP status 500 (Internal Server Error)
func InternalServerHTTP(message string) error {
	return newHTTPError(message, hTTPInternalServerError)
}

// NotImplementedHTTP создает ошибку с HTTP-статусом 501 (Not Implemented)
// NotImplementedHTTP creates an error with HTTP status 501 (Not Implemented)
func NotImplementedHTTP(message string) error {
	return newHTTPError(message, hTTPNotImplemented)
}

// BadGatewayHTTP создает ошибку с HTTP-статусом 502 (Bad Gateway)
// BadGatewayHTTP creates an error with HTTP status 502 (Bad Gateway)
func BadGatewayHTTP(message string) error {
	return newHTTPError(message, hTTPBadGateway)
}

// ServiceUnavailableHTTP создает ошибку с HTTP-статусом 503 (Service Unavailable)
// ServiceUnavailableHTTP creates an error with HTTP status 503 (Service Unavailable)
func ServiceUnavailableHTTP(message string) error {
	return newHTTPError(message, hTTPServiceUnavailable)
}

// GatewayTimeoutHTTP создает ошибку с HTTP-статусом 504 (Gateway Timeout)
// GatewayTimeoutHTTP creates an error with HTTP status 504 (Gateway Timeout)
func GatewayTimeoutHTTP(message string) error {
	return newHTTPError(message, hTTPGatewayTimeout)
}

// VersionNotSupportedHTTP создает ошибку с HTTP-статусом 505 (HTTP Version Not Supported)
// VersionNotSupportedHTTP creates an error with HTTP status 505 (HTTP Version Not Supported)
func VersionNotSupportedHTTP(message string) error {
	return newHTTPError(message, hTTPStatusHTTPVersionNotSupported)
}
