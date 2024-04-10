package utils

import "errors"

var (
	ErrTokenRequired = errors.New("token is required")
	ErrTokenExpired = errors.New("token expired")
)