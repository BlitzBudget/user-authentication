package helper

import (
	fetchuserModel "login/fetchuser/service/models"
	fetchWalletModel "login/fetch-wallet/service/models"
	"login/service/models"
	"testing"

	"github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/stretchr/testify/assert"
)

func TestParseResponse(t *testing.T) {
	assert := assert.New(t)

	session := "session"
	accessToken := "accessToken"
	refreshToken := "refreshToken"
	idToken := "idToken"
	authResultType := cognitoidentityprovider.AuthenticationResultType{
		AccessToken:  &accessToken,
		RefreshToken: &refreshToken,
		IdToken:      &idToken,
	}
	loginResponseModel := models.LoginResponseModel{
		Session:              &session,
		AuthenticationResult: &authResultType,
	}

	httpResponse := ParseResponse(&loginResponseModel)
	assert.Equal(*httpResponse.Session, session)
	assert.Equal(*httpResponse.AuthenticationResult.AccessToken, accessToken)
	assert.Equal(*httpResponse.AuthenticationResult.RefreshToken, refreshToken)
	assert.Equal(*httpResponse.AuthenticationResult.IdToken, idToken)
}

func TestParseFetchUserResponse(t *testing.T) {
	assert := assert.New(t)
	name := "name"
	value := "value"

	fetchUserModel := fetchuserModel.FetchUserResponse{
		UserAttributes: []*models.UserAttribute{
			{
				Name:  &name,
				Value: &value,
			},
		},
	}

	httpResponse := models.HttpResponse{}

	httpParsedResponse := ParseFetchUserResponse(&fetchUserModel, &httpResponse)
	assert.Equal(httpParsedResponse.UserAttributes, fetchUserModel.UserAttributes)
}

func TestParseWalletResponse(t *testing.T) {
	assert := assert.New(t)

	pk := "pk"
	sk := "sk"
	walletItems := fetchWalletModel.WalletResponseItems{}
	walletItem := fetchWalletModel.WalletResponseItem{
		Pk: &pk,
		Sk: &sk,
	}

	walletItems = append(walletItems, &walletItem)
	httpResponse := models.HttpResponse{}

	httpResponseItem := ParseWalletResponse(walletItems, &httpResponse)

	assert.Equal(httpResponseItem.Wallet[0], walletItems[0])
}
