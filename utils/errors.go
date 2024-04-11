package utils

import "errors"

var (
	ErrTokenRequired  = errors.New("token is required")
	ErrDomainRequired = errors.New("domain is required")
	ErrTokenExpired   = errors.New("token expired")
)
