package models

type ErrorHttpResponse struct {
	Message *string `validate:"required" json:"message"`
}
