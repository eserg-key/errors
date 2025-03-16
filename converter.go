package errors

// statusHTTPToGRPC converts HTTP status into gRPC code
func statusHTTPToGRPC(httpCode Code) Code {
	switch httpCode {
	case hTTPBadRequest:
		return gRPCInvalidArgument
	case hTTPUnauthorized:
		return gRPCUnauthenticated
	case hTTPPaymentRequired:
		return gRPCResourceExhausted
	case hTTPForbidden:
		return gRPCPermissionDenied
	case hTTPNotFound:
		return gRPCNotFound
	case hTTPMethodNotAllowed:
		return gRPCUnimplemented
	case hTTPNotAcceptable:
		return gRPCInvalidArgument
	case hTTPStatusProxyAuthRequired:
		return gRPCUnauthenticated
	case hTTPRequestTimeout:
		return gRPCDeadlineExceeded
	case hTTPConflict:
		return gRPCAborted
	case hTTPGone:
		return gRPCNotFound
	case hTTPStatusLengthRequired:
		return gRPCInvalidArgument
	case hTTPStatusPreconditionFailed:
		return gRPCFailedPrecondition
	case hTTPPayloadTooLarge:
		return gRPCResourceExhausted
	case hTTPURITooLong:
		return gRPCInvalidArgument
	case hTTPUnsupportedMediaType:
		return gRPCInvalidArgument
	case hTTPStatusRequestedRangeNotSatisfiable:
		return gRPCInvalidArgument
	case hTTPStatusExpectationFailed:
		return gRPCInvalidArgument
	case hTTPStatusTeapot:
		return gRPCInternal
	case hTTPUnprocessableEntity:
		return gRPCInvalidArgument
	case hTTPTooManyRequests:
		return gRPCResourceExhausted
	case hTTPInternalServerError:
		return gRPCInternal
	case hTTPNotImplemented:
		return gRPCUnimplemented
	case hTTPBadGateway:
		return gRPCUnavailable
	case hTTPServiceUnavailable:
		return gRPCUnavailable
	case hTTPGatewayTimeout:
		return gRPCDeadlineExceeded
	case hTTPStatusHTTPVersionNotSupported:
		return gRPCUnimplemented
	default:
		return gRPCUnknown
	}
}

// statusGRPCToHTTP converts the gRPC code to an HTTP status
func statusGRPCToHTTP(grpcCode Code) Code {
	switch grpcCode {
	case gRPCCanceled:
		return hTTPRequestTimeout
	case gRPCUnknown:
		return hTTPInternalServerError
	case gRPCInvalidArgument:
		return hTTPBadRequest
	case gRPCDeadlineExceeded:
		return hTTPGatewayTimeout
	case gRPCNotFound:
		return hTTPNotFound
	case gRPCAlreadyExists:
		return hTTPConflict
	case gRPCPermissionDenied:
		return hTTPForbidden
	case gRPCResourceExhausted:
		return hTTPTooManyRequests
	case gRPCFailedPrecondition:
		return hTTPBadRequest
	case gRPCAborted:
		return hTTPConflict
	case gRPCOutOfRange:
		return hTTPBadRequest
	case gRPCUnimplemented:
		return hTTPNotImplemented
	case gRPCInternal:
		return hTTPInternalServerError
	case gRPCUnavailable:
		return hTTPServiceUnavailable
	case gRPCDataLoss:
		return hTTPInternalServerError
	case gRPCUnauthenticated:
		return hTTPUnauthorized
	default:
		return hTTPInternalServerError
	}
}
