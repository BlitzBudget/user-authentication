package helper

import (
	fetchWalletService "login/fetch-wallet/service/models"
	fetchuserModel "login/fetchuser/service/models"
	"login/service/models"
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

func ParseWalletResponse(walletResponseItem []*fetchWalletService.WalletResponseItem, httpResponse *models.HttpResponse) *models.HttpResponse {
	httpResponse.Wallet = walletResponseItem
	return httpResponse
}
