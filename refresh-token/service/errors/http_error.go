package errors

import "errors"

var ErrEmptyRefreshTokenInRequest = errors.New("the refresh token is empty")
var ErrInvalidRefreshTokenInRequest = errors.New("the refresh token entered is invalid")

var ErrEmptyUserIdInRequest = errors.New("the user id is empty")
