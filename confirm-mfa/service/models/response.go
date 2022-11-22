package models

type Response struct {
	AccessToken        *string            `validate:"required" json:"access_token"`
	ExpiresIn          *int64             `validate:"required" json:"expires_in"`
	IdToken            *string            `validate:"required" json:"id_token"`
	RefreshToken       *string            `validate:"required" json:"refresh_token"`
	TokenType          *string            `validate:"required" json:"token_type"`
	DeviceMetaData     *DeviceMetaData    `validate:"required" json:"device_meta_data"`
	ChallangeName      *string            `validate:"required" json:"challenge_name"`
	ChallangeParameter map[string]*string `validate:"required" json:"challenge_parameter"`
}

type DeviceMetaData struct {
	DeviceGroupKey *string `validate:"required" json:"device_group_key"`
	DeviceKey      *string `validate:"required" json:"device_key"`
}
