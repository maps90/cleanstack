package httpx

import (
	"database/sql"
	"fmt"
	"net/http"
	"sort"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

// HTTPError struct
type (
	stackTracer interface {
		StackTrace() errors.StackTrace
	}
	HTTPError struct {
		RequestID string `json:"request_id"`
		// StatusCode is http status code
		StatusCode int `json:"status_code"`
		// Code is the error code
		Code string `json:"code"`
		// Message is the error message that may be displayed to end users
		Message string `json:"error_message"`
		// DeveloperMessage is the error message that is mainly meant for developers
		DeveloperMessage string `json:"developer_message,omitempty"`
		DeveloperTrace   string `json:"developer_trace,omitempty"`
		// Details specifies the additional error information
		Details interface{} `json:"details"`
	}
	validationError struct {
		Field string `json:"field"`
		Code  string `json:"code"`
		Error string `json:"error"`
	}
	Params map[string]interface{}
)

// Error returns the error message.
func (e HTTPError) Error() string {
	return e.Message
}

// InternalServerError creates a new API error representing an internal server error (HTTP 500)
func ErrInternalServerError(err error) *HTTPError {
	errs := &HTTPError{
		Message:          "We have encountered an error, please try again later.",
		DeveloperMessage: err.Error(),
	}
	errs.StatusCode = http.StatusInternalServerError
	return errs
}

// NotFound creates a new API error representing a resource-not-found error (HTTP 404)
func ErrNotFound() *HTTPError {
	errs := &HTTPError{
		Message: "The requested resource was not found.",
	}
	errs.StatusCode = http.StatusNotFound
	return errs
}

// Unauthorized creates a new API error representing an authentication failure (HTTP 401)
func ErrUnauthorized(err error) *HTTPError {
	errs := &HTTPError{
		Message:          "Restricted access, please check your credentials.",
		DeveloperMessage: fmt.Sprintf("UNAUTHORIZED : %v", err),
	}
	errs.StatusCode = http.StatusUnauthorized
	return errs
}

// ErrValidation converts a data validation error into an API error (HTTP 400)
func ErrValidation(errs validation.Errors) *HTTPError {
	result := make([]validationError, 0)
	fields := make([]string, 0)
	for field := range errs {
		fields = append(fields, field)
	}
	sort.Strings(fields)
	for _, field := range fields {
		err := errs[field].(validation.Error)
		result = append(result, validationError{
			Field: field,
			Code:  err.Code(),
			Error: err.Error(),
		})
	}

	err := &HTTPError{
		Message: "Failed to validate the request.",
	}
	err.StatusCode = http.StatusBadRequest
	err.Details = result

	return err
}

// ErrorHandler override echo.HTTPErrorHandler
func ErrorHandler(err error, c echo.Context) {
	var e = &Context{c}
	var pkgError = errors.Cause(err)

	if err == sql.ErrNoRows {
		e.JSONErr(ErrNotFound())
		return
	}

	switch pkgError.(type) {
	case *HTTPError:
		e.JSONErr(err.(*HTTPError))
	case validation.Errors:
		e.JSONErr(ErrValidation(err.(validation.Errors)))
	case *echo.HTTPError:
		switch pkgError.(*echo.HTTPError).Code {
		case http.StatusUnauthorized:
			e.JSONErr(ErrUnauthorized(err))
		case http.StatusNotFound:
			e.JSONErr(ErrNotFound())
		case http.StatusBadRequest:
			var a = &HTTPError{
				StatusCode:       http.StatusBadRequest,
				Message:          "Failed to validate the request.",
				DeveloperMessage: err.Error(),
			}
			if st, ok := err.(stackTracer); ok {
				a.DeveloperTrace = fmt.Sprintf("called at: %+v", st.StackTrace()[0:2])
			}
			e.JSONErr(a)
		default:
			errEcho := err.(*echo.HTTPError)
			e.JSONErr(&HTTPError{
				StatusCode: errEcho.Code,
				Message:    errEcho.Error(),
			})
		}
	default:
		e.JSONErr(ErrInternalServerError(err))
	}
}
