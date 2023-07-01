package util

import (
	"authentication/domain"
	"errors"
)

func GetHttpStatusCode(err error) int {
	switch {
	case errors.Is(err, domain.ErrAuthFailed):
		return 401
	default:
		return 500
	}
}
