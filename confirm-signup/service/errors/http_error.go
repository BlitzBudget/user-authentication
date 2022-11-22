package errors

import "errors"

var ErrEmptyEmailInRequest = errors.New("the email id is empty")
var ErrInvalidEmailInRequest = errors.New("the email entered is invalid")

var ErrEmptyConfirmationCodeInRequest = errors.New("the confirmation code is empty")
var ErrInvalidConfirmationCodeInRequest = errors.New("the confirmation code entered is invalid")
