package errors

import (
	"errors"
	"net/http"
)

var (
	InternalError = Error{
		ErrorCode:    "KA-001-E",
		ErrorMessage: "Internal Error.",
		StatusCode:   http.StatusInternalServerError,
	}

	RequestSucess = Error{
		ErrorCode:    "KA-001-S",
		ErrorMessage: "Request Processed Successfully. ",
		StatusCode:   http.StatusOK,
	}

	UserNotAuth = Error{
		ErrorCode:    "KA-002-E",
		ErrorMessage: "User Not Authorized.",
		StatusCode:   http.StatusUnauthorized,
	}

	ValidationFailed = Error{
		ErrorCode:    "KA-003-E",
		ErrorMessage: "Request Validation Failed",
		StatusCode:   http.StatusBadRequest,
	}

	UserNotFound = Error{
		ErrorCode:    "KA-004-E",
		ErrorMessage: "User Not Found.",
		StatusCode:   http.StatusBadRequest,
	}

	ErrInvalidRequest = Error{
		ErrorCode:    "KA-005-E",
		ErrorMessage: "Invalid Request",
		StatusCode:   http.StatusBadRequest,
	}

	ErrNoServiceFound = Error{
		ErrorCode:    "KA-006-E",
		ErrorMessage: "No Service Found",
		StatusCode:   http.StatusBadRequest,
	}

	ErrUserNotFound  = errors.New("user not found")
	ErrWrongUserId   = errors.New("user id is wrong")
	ErrWrongPassword = errors.New("password is wrong")
)
