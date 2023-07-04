package util

import (
	"authentication/domain"
	"errors"
)

func GetHttpStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	case errors.Is(err, domain.ErrOTPInvalid):
		return 400
	case errors.Is(err, domain.ErrOTPExpired):
		return 400
	case errors.Is(err, domain.ErrBadRequest):
		return 400
	case errors.Is(err, domain.ErrUsernameAlreadyExist):
		return 400
	case errors.Is(err, domain.ErrEmailAlreadyExist):
		return 400
	default:
		return 500
	}
}
