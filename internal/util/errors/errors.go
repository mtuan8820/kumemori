package errors

import "fmt"

type ErrorType string

const (
	// ErrorTypeValidation represents validation errors
	ErrorTypeValidation ErrorType = "VALIDATION"

	// ErrorTypeNotFound represents resource not found errors
	ErrorTypeNotFound ErrorType = "NOT_FOUND"

	// ErrorTypePersistence represents persistence layer errors
	ErrorTypePersistence ErrorType = "PERSISTENCE"

	// ErrorTypeSystem represents internal system errors
	ErrorTypeSystem ErrorType = "SYSTEM"

	// ErrorTypeBusiness represents business logic errors
	ErrorTypeBusiness ErrorType = "BUSINESS"

	// ErrorTypeUnauthorized represents authentication/authorization errors
	ErrorTypeUnauthorized ErrorType = "UNAUTHORIZED"

	// ErrorTypeForbidden represents permission errors
	ErrorTypeForbidden ErrorType = "FORBIDDEN"

	// ErrorTypeConflict represents resource conflict errors
	ErrorTypeConflict ErrorType = "CONFLICT"
)

// AppError defines the application error structure
type AppError struct {
	Type    ErrorType
	Message string
	Cause   error
	Details map[string]any
	Code    int
}

// Error implements the error interface
func (e *AppError) Error() string {
	if e.Cause != nil {
		return fmt.Sprintf("%s: %s (caused by: %v)", e.Type, e.Message, e.Cause)
	}
	return fmt.Sprintf("%s: %s", e.Type, e.Message)
}

// Unwrap supports Go 1.13+ error unwrapping
func (e *AppError) Unwrap() error {
	return e.Cause
}

// WithDetails adds error details
func (e *AppError) WithDetails(details map[string]any) *AppError {
	e.Details = details
	return e
}

// WithCode adds an error code
func (e *AppError) WithCode(code int) *AppError {
	e.Code = code
	return e
}

func NewValidationError(message string, cause error) *AppError {
	return &AppError{
		Type:    ErrorTypeValidation,
		Message: message,
		Cause:   cause,
	}
}

// Wrap wraps a standard error as an application error
func Wrap(err error, errType ErrorType, message string) *AppError {
	return &AppError{
		Type:    errType,
		Message: message,
		Cause:   err,
	}
}

// Wrapf wraps an error with a formatted message
func Wrapf(err error, errType ErrorType, format string, args ...any) *AppError {
	return &AppError{
		Type:    errType,
		Message: fmt.Sprintf(format, args...),
		Cause:   err,
	}
}

// New creates a new application error
func New(errType ErrorType, message string) *AppError {
	return &AppError{
		Type:    errType,
		Message: message,
	}
}

// Newf creates a new application error with a formatted message
func Newf(errType ErrorType, format string, args ...any) *AppError {
	return &AppError{
		Type:    errType,
		Message: fmt.Sprintf(format, args...),
	}
}
