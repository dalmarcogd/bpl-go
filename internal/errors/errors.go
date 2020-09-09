package errors

import "errors"

var (
	UserNotFound = errors.New("user not found")
	UserIdRequired = errors.New("user id is required")
)
