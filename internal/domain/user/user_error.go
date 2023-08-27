package user

import "errors"

var ErrUserNotFound = errors.New("USER_NOT_FOUND")
var ErrUserIncorrectInformation = errors.New("USER_INCORRECT_INFORMATION")
