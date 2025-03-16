# Errors Package
The errors package is a Go library designed to simplify error handling in applications that use both HTTP and gRPC protocols. It provides a unified structure for errors, allowing you to create, wrap, and convert errors with associated HTTP and gRPC status codes. This package is particularly useful for microservices and APIs that need to handle errors consistently across different communication protocols.

## Features
- Unified Error Structure: Define errors with both HTTP and gRPC status codes.

- Error Wrapping: Wrap errors with additional context while preserving the original error's status code.

- Protocol Conversion: Automatically convert between HTTP and gRPC status codes.

- Context-Aware Errors: Handle context-related errors (e.g., context.DeadlineExceeded) with appropriate status codes.

- Predefined Error Constructors: Easily create errors with predefined HTTP and gRPC status codes.

## Installation
To use the errors package, install it using go get:
```
go get -u github.com/eserg-key/errors
```
## Usage
Then, import it in your Go code:
```
import "github.com/eserg-key/errors"
```

## Creating Errors
You can create errors with specific HTTP or gRPC status codes using the provided constructors.
### HTTP Errors
```
err := errors.BadRequestHTTP("Invalid input provided")
err := errors.NotFoundHTTP("Resource not found")
err := errors.InternalServerHTTP("Internal server error")
```
### gRPC Errors
```
err := errors.InvalidArgumentGRPC("Invalid argument provided")
err := errors.NotFoundGRPC("Resource not found")
err := errors.InternalGRPC("Internal server error")
```

## Wrapping Errors
You can wrap existing errors to add additional context while preserving the original error's status code.
```
originalErr := errors.NotFoundHTTP("User not found")
wrappedErr := errors.Wrap(originalErr, "Failed to fetch user details")

fmt.Println(wrappedErr.Error()) // Output: Failed to fetch user details: User not found
```

## Retrieving Status Codes
You can retrieve the HTTP or gRPC status code from an error.
### HTTP Status Code
```
err := errors.BadRequestHTTP("Invalid input")
status := errors.StatusHTTP(err)
fmt.Println(status) // Output: 400
```
### gRPC Status Code
```
err := errors.InvalidArgumentGRPC("Invalid argument")
status := errors.StatusGRPC(err)
fmt.Println(status) // Output: 3 (gRPC InvalidArgument code)
```

## Handling Context Errors
The package automatically handles context-related errors (e.g., context.DeadlineExceeded and context.Canceled) and maps them to appropriate status codes.
```
err := context.DeadlineExceeded
httpStatus := errors.StatusHTTP(err)
grpcStatus := errors.StatusGRPC(err)

fmt.Println(httpStatus) // Output: 504 (HTTP Gateway Timeout)
fmt.Println(grpcStatus) // Output: 4 (gRPC DeadlineExceeded)
```

## Converting Between HTTP and gRPC Status Codes
The package provides utility functions to convert between HTTP and gRPC status codes.
### HTTP to gRPC
```
httpCode := errors.HTTPBadRequest
grpcCode := errors.StatusHTTPToGRPC(httpCode)
fmt.Println(grpcCode) // Output: 3 (gRPC InvalidArgument)
```
### gRPC to HTTP
```
grpcCode := errors.GRPCInvalidArgument
httpCode := errors.StatusGRPCToHTTP(grpcCode)
fmt.Println(httpCode) // Output: 400 (HTTP Bad Request)
```

## Examples
### Example 1: Creating and Wrapping Errors
```
package main

import (
	"fmt"
	"github.com/eserg-key/errors"
)

func main() {
	// Create an HTTP error
	err := errors.BadRequestHTTP("Invalid input")

	// Wrap the error with additional context
	wrappedErr := errors.Wrap(err, "Validation failed")

	// Retrieve the HTTP status code
	status := errors.StatusHTTP(wrappedErr)
	fmt.Println(status) // Output: 400

	// Print the error message
	fmt.Println(wrappedErr.Error()) // Output: Validation failed: Invalid input
}
```

## License
This project is licensed under the MIT License. See the LICENSE file for details.

## Acknowledgments
- Inspired by the github.com/pkg/errors package.
- Built for Go developers who need consistent error handling across HTTP and gRPC protocols.