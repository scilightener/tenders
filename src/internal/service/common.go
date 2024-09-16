package service

import "errors"

var (
	ErrUnknownError = errors.New("unknown error")
	ErrUnprivileged = errors.New("you have no access to this entity/operation")
)
