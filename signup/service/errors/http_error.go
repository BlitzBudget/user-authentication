package errors

import "errors"

var ErrEmptyEmailInRequest = errors.New("the email id is empty")
var ErrInvalidEmailInRequest = errors.New("the email entered is invalid")

var ErrEmptyPasswordInRequest = errors.New("the password provided is empty")
var ErrMinimumLengthRequiredPassword = errors.New("the password provided must have a minimum length of 8 characters")

var ErrEmptyPhoneNumberInRequest = errors.New("the phone number is empty in the request")
var ErrEmptyUserRoleInRequest = errors.New("the user role is empty in the request")
var ErrEmptyOrganizationInRequest = errors.New("the organization is empty in the request")
var ErrNameShouldHaveOnlyCharactersInRequest = errors.New("the name should have only characters in the request")
var ErrLastNameShouldHaveOnlyCharactersInRequest = errors.New("the last name should have only characters in the request")
var ErrMiddleNameShouldHaveOnlyCharactersInRequest = errors.New("the middle name should have only characters in the request")

var ErrInvalidPhoneNumberInRequest = errors.New("the phone number is invalid in the request")
var ErrInvalidUserRoleInRequest = errors.New("the user role is invalid in the request")
var ErrInvalidOrganizationInRequest = errors.New("the organization is invalid in the request")
