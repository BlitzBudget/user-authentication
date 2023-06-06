package errors

import "errors"

var ErrEmptyEmailInRequest = errors.New("the email id is empty")
var ErrInvalidEmailInRequest = errors.New("the email entered is invalid")

var ErrEmptyAccessTokenInRequest = errors.New("the access token provided is empty")
var ErrMinimumLengthRequiredAccessToken = errors.New("the access token provided is invalid")
