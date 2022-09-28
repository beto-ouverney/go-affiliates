package customerror

import "fmt"

// Application error codes.
const (
	ECONFLICT = "conflict"  // action cannot be performed - Example: duplicate email
	EINTERNAL = "internal"  // internal error - Example: internal errors, referring to the language itself or to the server where the code runs. Example: save a file, marshal a json
	EINVALID  = "invalid"   // validation failed - Logic errors created by us. Example: saving a user to the database
	ENOTFOUND = "not_found" // entity does not exist - Logic errors created by us. Example: saving a user to the database
)

// Error defines a standard application error.
type CustomError struct {
	Code    string // Machine-readable error code
	Message string // Human-readable message (final user)
	Op      string // Logical operation (operation role)
	Err     error  // Embedded error  (operation role)
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("%v", e.Err)
}

// NewError creates a new custom error
func NewError(code, message, op string, err error) *CustomError {
	return &CustomError{
		Code:    code,
		Message: message,
		Op:      op,
		Err:     err,
	}
}
