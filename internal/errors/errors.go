package errors

import (
	"errors"
	"gorm.io/gorm"
	"net/http"
)

type AppError struct {
	Code       string // "bookmark_not_found"
	Message    string // "bookmark not found"
	Details    string // underlying error text
	HTTPStatus int    // 404
}

func (e *AppError) Error() string {
	if e.Details != "" {
		return e.Details
	}
	return e.Message
}

type ErrorBody struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Details string `json:"details,omitempty"`
}

type Response struct {
	Error ErrorBody `json:"error"`
}

func UnauthorizedError() *AppError {
	return &AppError{
		Code:       "unauthorized_error",
		Message:    "unauthorized",
		HTTPStatus: http.StatusUnauthorized,
	}
}

func ForbiddenError() *AppError {
	return &AppError{
		Code:       "forbidden_error",
		Message:    "forbidden",
		HTTPStatus: http.StatusForbidden,
	}
}

func NewError(err error) error {
	var appErr *AppError
	switch {
	case errors.Is(err, gorm.ErrRecordNotFound):
		appErr = NotFoundError()
	case errors.Is(err, gorm.ErrInvalidData):
		appErr = RecordInvalidError()
	default:
		appErr = InternalError()
	}
	appErr.Details = err.Error()
	return appErr
}

func NotFoundError() *AppError {
	return &AppError{
		Code:       "not_found_error",
		Message:    "not found",
		HTTPStatus: http.StatusNotFound,
	}
}

func RecordInvalidError() *AppError {
	return &AppError{
		Code:       "record_invalid_error",
		Message:    "record invalid",
		HTTPStatus: http.StatusUnprocessableEntity,
	}
}

func InvalidIDError() *AppError {
	return &AppError{
		Code:       "invalid_id",
		Message:    "invalid id",
		HTTPStatus: http.StatusUnprocessableEntity,
	}
}

func InternalError() *AppError {
	return &AppError{
		Code:       "internal_error",
		Message:    "internal server error",
		HTTPStatus: http.StatusInternalServerError,
	}
}
