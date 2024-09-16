package employee

import "errors"

var (
	ErrUsernameTooLong = errors.New("username is too long")
	ErrUsernameEmpty   = errors.New("username can't be empty")
	ErrNotFound        = errors.New("employee not found")
)
