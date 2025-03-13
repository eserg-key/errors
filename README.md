# Errors Package
The errors package is a versatile Go library designed to simplify error handling in both HTTP and gRPC contexts. It provides a unified way to create, wrap, and manage errors while preserving HTTP status codes and gRPC error codes. This package is ideal for projects that need to handle errors consistently across different protocols.

## Features
- HTTP and gRPC Error Support: Create errors with HTTP status codes or gRPC error codes.

- Error Wrapping: Wrap errors with additional context while preserving the original error's status or code.

- Multi-Error Handling: Append multiple errors into a single error object.

- Error Flattening: Flatten nested multi-errors into a single list of errors.

- Error Prefixing: Add a prefix to error messages while preserving the original error's status or code.

- Compatibility: Works seamlessly with standard Go error handling and popular error libraries like github.com/pkg/errors.

## Installation
To install the package, use go get:
```
go get github.com/eserg-key/errors
```
## Usage
Importing the Package
```
import "github.com/eserg-key/errors"
```

## Creating HTTP Errors
The package provides constructors for common HTTP errors:
```
err := errors.BadRequest("Invalid input")
err := errors.NotFound("Resource not found")
err := errors.InternalServer("Internal server error")
```

## Creating gRPC Errors
For gRPC, you can create errors with specific gRPC codes:
```
err := errors.GRPCNotFound("Resource not found")
err := errors.GRPCInternal("Internal server error")
err := errors.GRPCInvalidArgument("Invalid argument")
```

## Creating Custom Errors
You can create custom errors with specific HTTP status codes or gRPC codes:
```azure
// Custom HTTP error
err := errors.New("Custom error message", http.StatusForbidden)

// Custom gRPC error
err := errors.NewGRPC("Custom error message", codes.Unavailable)
```

## Wrapping Errors
Wrap an existing error to add more context while preserving the original error's status or code:
```azure
originalErr := errors.NotFound("Resource not found")
wrappedErr := errors.Wrap(originalErr, "Failed to fetch resource")
```

## Handling Multiple Errors
Append multiple errors into a single error object:
```azure
err1 := errors.BadRequest("Invalid input")
err2 := errors.NotFound("Resource not found")
combinedErr := errors.Append(err1, err2)
```

## Flattening Errors
Flatten nested multi-errors into a single list of errors:
```azure
flattenedErr := errors.Flatten(combinedErr)
```

## Adding Prefixes to Errors
Add a prefix to an error message while preserving the original error's status or code:
```azure
prefixedErr := errors.Prefix(err, "API Error")
```

## Extracting HTTP and gRPC Status Codes
Retrieve the HTTP status code or gRPC code from an error:
```azure
httpStatus := errors.HTTPStatus(err)
grpcCode := errors.GRPCStatus(err)
```

## Checking and Unwrapping Errors
Use standard Go error handling functions:
```azure
// Check if an error is of a specific type
var targetErr *errors.Error
    if errors.As(err, &targetErr) {
    // Handle the error
}

// Check if an error matches another error
        if errors.Is(err, targetErr) {
    // Handle the error
}

// Unwrap an error
unwrappedErr := errors.Unwrap(err)

// Get the root cause of an error
rootCause := errors.Cause(err)
```

## Examples
### Example 1: Creating and Wrapping Errors
```azure
package main

    import (
	"fmt"
	"github.com/yourusername/errors"
	"net/http"
)

func main() {
	err := errors.NotFound("User not found")
	wrappedErr := errors.Wrap(err, "Failed to fetch user")

	fmt.Println(wrappedErr.Error()) // Output: Failed to fetch user: User not found
	fmt.Println(errors.HTTPStatus(wrappedErr)) // Output: 404
}
```

### Example 2: Handling Multiple Errors
```azure
package main

    import (
	"fmt"
	"github.com/yourusername/errors"
)

func main() {
	err1 := errors.BadRequest("Invalid email")
	err2 := errors.BadRequest("Invalid password")
	combinedErr := errors.Append(err1, err2)

	fmt.Println(combinedErr.Error()) // Output: 2 errors occurred:
    //         * Invalid email
	                                //         * Invalid password
}
```

### Example 3: Using gRPC Errors
```azure
package main

    import (
	"fmt"
	"github.com/yourusername/errors"
	"google.golang.org/grpc/codes"
)

func main() {
	err := errors.GRPCNotFound("User not found")
	fmt.Println(err.Error()) // Output: User not found
	fmt.Println(errors.GRPCStatus(err)) // Output: 5 (codes.NotFound)
}
```

## Contributing
Contributions are welcome! If you find a bug or have a feature request, please open an issue or submit a pull request.
1. Fork the repository.
2. Create a new branch for your changes.
3. Commit your changes and push to the branch.
4. Submit a pull request.

## License
This project is licensed under the MIT License.