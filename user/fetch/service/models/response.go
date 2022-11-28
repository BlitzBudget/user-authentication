package models

type HttpResponse struct {
	UserAttributes []*UserAttribute `validate:"required" json:"user_attributes"`
	Username       *string          `validate:"required" json:"username"`
}

type Response struct {
	UserAttributes []*UserAttribute `validate:"required" json:"user_attributes"`
	Username       *string          `validate:"required" json:"username"`
}

// Specifies whether the attribute is standard or custom.
type UserAttribute struct {

	// The name of the attribute.
	//
	// Name is a required field
	Name *string `min:"1" type:"string" required:"true"`

	// The value of the attribute.
	//
	// Value is a sensitive parameter and its value will be
	// replaced with "sensitive" in string returned by AttributeType's
	// String and GoString methods.
	Value *string `type:"string" sensitive:"true"`
}
