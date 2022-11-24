package helper

import "confirm-signup/service/models"

func ParseResponse(loginResponseModel *models.LoginResponseModel) *models.HttpResponse {
	responseModel := models.HttpResponse{
		Session:              loginResponseModel.Session,
		AuthenticationResult: loginResponseModel.AuthenticationResult,
		ChallengeName:        loginResponseModel.ChallengeName,
	}

	return &responseModel
}
