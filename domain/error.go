package domain

import "errors"

var (
	ErrAuthFailed          = errors.New("authentication failed")
	ErrInternalServerError = errors.New("internal server error")
)
