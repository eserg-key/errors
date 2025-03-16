package errors

// ProtocolType - type of protocol
type ProtocolType string

const (
	httpProtocol ProtocolType = "http"
	grpcProtocol ProtocolType = "grpc"
)

// A Code is a status code defined
type Code uint32

// HTTP status constants
const (
	hTTPOk                                 Code = 200
	hTTPBadRequest                         Code = 400
	hTTPUnauthorized                       Code = 401
	hTTPPaymentRequired                    Code = 402
	hTTPForbidden                          Code = 403
	hTTPNotFound                           Code = 404
	hTTPMethodNotAllowed                   Code = 405
	hTTPNotAcceptable                      Code = 406
	hTTPStatusProxyAuthRequired            Code = 407
	hTTPRequestTimeout                     Code = 408
	hTTPConflict                           Code = 409
	hTTPGone                               Code = 410
	hTTPStatusLengthRequired               Code = 411
	hTTPStatusPreconditionFailed           Code = 412
	hTTPPayloadTooLarge                    Code = 413
	hTTPURITooLong                         Code = 414
	hTTPUnsupportedMediaType               Code = 415
	hTTPStatusRequestedRangeNotSatisfiable Code = 416
	hTTPStatusExpectationFailed            Code = 417
	hTTPStatusTeapot                       Code = 418
	hTTPUnprocessableEntity                Code = 422
	hTTPTooManyRequests                    Code = 429
	hTTPInternalServerError                Code = 500
	hTTPNotImplemented                     Code = 501
	hTTPBadGateway                         Code = 502
	hTTPServiceUnavailable                 Code = 503
	hTTPGatewayTimeout                     Code = 504
	hTTPStatusHTTPVersionNotSupported      Code = 505
)

// gRPC status constants
const (
	gRPCOk                 Code = 0
	gRPCCanceled           Code = 1
	gRPCUnknown            Code = 2
	gRPCInvalidArgument    Code = 3
	gRPCDeadlineExceeded   Code = 4
	gRPCNotFound           Code = 5
	gRPCAlreadyExists      Code = 6
	gRPCPermissionDenied   Code = 7
	gRPCResourceExhausted  Code = 8
	gRPCFailedPrecondition Code = 9
	gRPCAborted            Code = 10
	gRPCOutOfRange         Code = 11
	gRPCUnimplemented      Code = 12
	gRPCInternal           Code = 13
	gRPCUnavailable        Code = 14
	gRPCDataLoss           Code = 15
	gRPCUnauthenticated    Code = 16
)
