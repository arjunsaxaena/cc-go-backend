package util

import (
	"errors"
	"net/http"
)

func ErrRequiredInputMissing(field string) error {
	return errors.New("required input missing: " + field)
}

func ErrUnchangeable(field string) error {
	return errors.New("field unchangeable after creation: " + field)
}

func ErrHeaderMissing(field string) error {
	return errors.New("header missing: " + field)
}

var (
	ErrExpiredToken          = errors.New("expired token")
	ErrInvalidToken          = errors.New("invalid token")
	ErrInternal              = errors.New("internal error")
	ErrTokenMissing          = errors.New("token missing")
)

var CustomErrorType = map[error]int{
	ErrExpiredToken:                          http.StatusUnauthorized,
	ErrInvalidToken:                          http.StatusUnauthorized,
	ErrInternal:                              http.StatusInternalServerError,
	ErrTokenMissing:                          http.StatusUnauthorized,
}
