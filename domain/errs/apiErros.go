package errs

import (
	"net/http"
)

type ApiError struct {
	StatusCode int
	Message    string
}

func NewBadRequestError(msg string) *ApiError {
	return &ApiError{http.StatusBadRequest, msg}
}

func NewNotFoundError(msg string) *ApiError {
	return &ApiError{http.StatusNotFound, msg}
}

func NewUnauthorizedError(msg string) *ApiError {
	return &ApiError{http.StatusUnauthorized, msg}
}

func NewForbiddenError(msg string) *ApiError {
	return &ApiError{http.StatusForbidden, msg}
}

func NewUnprocessableEntityError(msg string) *ApiError {
	return &ApiError{http.StatusUnprocessableEntity, msg}
}

func NewUnexpectedError(msg string) *ApiError {
	return &ApiError{http.StatusInternalServerError, msg}
}
