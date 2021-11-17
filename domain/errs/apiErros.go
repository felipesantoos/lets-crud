package errs

import (
	"net/http"
)

type ApiError struct {
	StatusCode int
	Message    string
}

func NewBadRequestError(message string) *ApiError {
	return &ApiError{http.StatusBadRequest, message}
}

func NewNotFoundError(message string) *ApiError {
	return &ApiError{http.StatusNotFound, message}
}

func NewUnauthorizedError(message string) *ApiError {
	return &ApiError{http.StatusUnauthorized, message}
}

func NewForbiddenError(message string) *ApiError {
	return &ApiError{http.StatusForbidden, message}
}

func NewUnprocessableEntityError(message string) *ApiError {
	return &ApiError{http.StatusUnprocessableEntity, message}
}

func NewUnexpectedError(message string) *ApiError {
	return &ApiError{http.StatusInternalServerError, message}
}
