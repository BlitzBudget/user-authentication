package errors

import "errors"

var ErrEmptyEmailInRequest = errors.New("the email id is empty")
var ErrInvalidEmailInRequest = errors.New("the email entered is invalid")

var ErrEmptyPasswordInRequest = errors.New("the password provided is empty")
var ErrMinimumLengthRequiredPassword = errors.New("the password provided must have a minimum length of 8 characters")
