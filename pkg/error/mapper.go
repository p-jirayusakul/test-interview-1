package error

import (
	"errors"
	"net/http"
)

func HTTPStatus(err error) int {
	var e *Error
	if !errors.As(err, &e) {
		return http.StatusInternalServerError
	}

	switch e.Code {
	case CodeInvalidInput:
		return http.StatusBadRequest
	case CodeUnauthorized:
		return http.StatusUnauthorized
	case CodeForbidden:
		return http.StatusForbidden
	case CodeNotFound:
		return http.StatusNotFound
	case CodeConflict:
		return http.StatusConflict
	case CodeBusiness:
		return http.StatusBadRequest
	case CodeDependencyUnavailable:
		return http.StatusServiceUnavailable
	case CodeUnknown, CodeSystem:
		return http.StatusInternalServerError
	default:
		return http.StatusInternalServerError
	}
}

func GetErrorCode(err error) string {
	if e, ok := errors.AsType[*Error](err); ok {
		return string(e.Code)
	}
	return string(CodeUnknown)
}
