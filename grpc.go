package errors

// newHTTP создает новую ошибку с заданным сообщением и HTTP-статусом
// newHTTP creates a new error with the specified message and HTTP status
func newGRPCError(message string, code Code) error {
	return &Error{Message: message, code: code, typeProtocol: grpcProtocol}
}

// CanceledGRPC создает ошибку с gRPC кодом 1
// CanceledGRPC creates an error with gRPC code 1
func CanceledGRPC(message string) error {
	return newGRPCError(message, gRPCCanceled)
}

// UnknownGRPC создает ошибку с gRPC кодом 2
// UnknownGRPC creates an error with gRPC code 2
func UnknownGRPC(message string) error {
	return newGRPCError(message, gRPCUnknown)
}

// InvalidArgumentGRPC создает ошибку с gRPC кодом 3
// InvalidArgumentGRPC creates an error with gRPC code 3
func InvalidArgumentGRPC(message string) error {
	return newGRPCError(message, gRPCInvalidArgument)
}

// DeadlineExceededGRPC создает ошибку с gRPC кодом 4
// DeadlineExceededGRPC creates an error with gRPC code 4
func DeadlineExceededGRPC(message string) error {
	return newGRPCError(message, gRPCDeadlineExceeded)
}

// NotFoundGRPC создает ошибку с gRPC кодом 5
// NotFoundGRPC creates an error with gRPC code 5
func NotFoundGRPC(message string) error {
	return newGRPCError(message, gRPCNotFound)
}

// AlreadyExistsGRPC создает ошибку с gRPC кодом 6
// AlreadyExistsGRPC creates an error with gRPC code 6
func AlreadyExistsGRPC(message string) error {
	return newGRPCError(message, gRPCAlreadyExists)
}

// PermissionDeniedGRPC создает ошибку с gRPC кодом 7
// PermissionDeniedGRPC creates an error with gRPC code 7
func PermissionDeniedGRPC(message string) error {
	return newGRPCError(message, gRPCPermissionDenied)
}

// ResourceExhaustedGRPC создает ошибку с gRPC кодом 8
// ResourceExhaustedGRPC creates an error with gRPC code 8
func ResourceExhaustedGRPC(message string) error {
	return newGRPCError(message, gRPCResourceExhausted)
}

// FailedPreconditionGRPC создает ошибку с gRPC кодом 9
// FailedPreconditionGRPC creates an error with gRPC code 9
func FailedPreconditionGRPC(message string) error {
	return newGRPCError(message, gRPCFailedPrecondition)
}

// AbortedGRPC создает ошибку с gRPC кодом 10
// AbortedGRPC creates an error with gRPC code 10
func AbortedGRPC(message string) error {
	return newGRPCError(message, gRPCAborted)
}

// OutOfRangeGRPC создает ошибку с gRPC кодом 11
// OutOfRangeGRPC creates an error with gRPC code 11
func OutOfRangeGRPC(message string) error {
	return newGRPCError(message, gRPCOutOfRange)
}

// UnimplementedGRPC создает ошибку с gRPC кодом 12
// UnimplementedGRPC creates an error with gRPC code 12
func UnimplementedGRPC(message string) error {
	return newGRPCError(message, gRPCUnimplemented)
}

// InternalGRPC создает ошибку с gRPC кодом 13
// InternalGRPC creates an error with gRPC code 13
func InternalGRPC(message string) error {
	return newGRPCError(message, gRPCInternal)
}

// UnavailableGRPC создает ошибку с gRPC кодом 14
// UnavailableGRPC creates an error with gRPC code 14
func UnavailableGRPC(message string) error {
	return newGRPCError(message, gRPCUnavailable)
}

// DataLossGRPC создает ошибку с gRPC кодом 15
// DataLossGRPC creates an error with gRPC code 15
func DataLossGRPC(message string) error {
	return newGRPCError(message, gRPCDataLoss)
}

// UnauthenticatedGRPC создает ошибку с gRPC кодом 16
// UnauthenticatedGRPC creates an error with gRPC code 16
func UnauthenticatedGRPC(message string) error {
	return newGRPCError(message, gRPCUnauthenticated)
}
