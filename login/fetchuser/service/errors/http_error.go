package errors

import "errors"

var ErrEmptyAccessTokenInRequest = errors.New("the access token is empty")
var ErrInvalidAccessTokenInRequest = errors.New("the access token entered is invalid")
