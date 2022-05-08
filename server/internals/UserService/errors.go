package userservice

import "errors"

type BadUserInputError struct {
	msg string
}

func (e BadUserInputError) Error() string {
	return e.msg
}

func newBadUserInputError(msg string) BadUserInputError {
	return BadUserInputError{msg: msg}
}

type FailedUserCreationError struct {
	msg string
}

func (e FailedUserCreationError) Error() string {
	return e.msg
}

func newFailedUserCreationError(msg string) FailedUserCreationError {
	return FailedUserCreationError{msg: msg}
}

var ErrLoginFailed = errors.New("login failed")

var ErrInvalidToken = errors.New("invalid token")
