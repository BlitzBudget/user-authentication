package errors

import "errors"

var ErrEmptyAccessTokenInRequest = errors.New("the access token is empty")

var ErrEmptyPreviousPasswordInRequest = errors.New("the previous password provided is empty")
var ErrMinimumLengthRequiredPreviousPassword = errors.New("the previous password provided must have a minimum length of 8 characters")

var ErrEmptyProposedPasswordInRequest = errors.New("the proposed password provided is empty")
var ErrMinimumLengthRequiredProposedPassword = errors.New("the proposed password provided must have a minimum length of 8 characters")
