package helper

import (
	"login/service/models"
	fetchuserModel "login/fetchuser/service/models"
)

func ParseResponse(loginResponseModel *models.LoginResponseModel) *models.HttpResponse {
	responseModel := models.HttpResponse{
		Session:              loginResponseModel.Session,
		AuthenticationResult: loginResponseModel.AuthenticationResult,
		ChallengeName:        loginResponseModel.ChallengeName,
	}

	return &responseModel
}

func ParseFetchUserResponse(fetchUserResponse *fetchuserModel.FetchUserResponse, httpResponse *models.HttpResponse) *models.HttpResponse {
	httpResponse.UserAttributes = fetchUserResponse.UserAttributes
	return httpResponse
}
