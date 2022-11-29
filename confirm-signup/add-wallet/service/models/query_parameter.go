package models

// Create struct to hold info about new item
type QueryParameter struct {
	Pk           *string `validate:"required" json:"pk"`
	Sk           string  `json:"sk"`
	Currency     *string `validate:"required" json:"wallet_currency"`
	Name         *string `validate:"required" json:"wallet_name"`
	CreationDate *string `json:"creation_date"`
	UpdatedDate  *string `json:"updated_date"`
}
