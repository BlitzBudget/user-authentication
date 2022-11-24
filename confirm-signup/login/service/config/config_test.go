package config

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert := assert.New(t)
	clientId := "clientId"
	// Set environment variable
	os.Setenv("COGNITO_CLIENT_ID", clientId)
	// returns value of COGNITO_CLIENT_ID
	fmt.Println("COGNITO_CLIENT_ID:", os.Getenv("COGNITO_CLIENT_ID"))
	ClientId = os.Getenv("COGNITO_CLIENT_ID")

	assert.Equal(clientId, ClientId)
}

func TestConfigClientSecret(t *testing.T) {
	assert := assert.New(t)
	clientSecret := "clientSecret"
	// Set environment variable
	os.Setenv("COGNITO_CLIENT_SECRET", clientSecret)
	// returns value of COGNITO_CLIENT_SECRET
	fmt.Println("COGNITO_CLIENT_SECRET:", os.Getenv("COGNITO_CLIENT_SECRET"))
	ClientSecret = os.Getenv("COGNITO_CLIENT_SECRET")

	assert.Equal(clientSecret, ClientSecret)
}

func TestConfigUserPoolId(t *testing.T) {
	assert := assert.New(t)
	userPoolId := "userPoolId"
	// Set environment variable
	os.Setenv("COGNITO_USER_POOL_ID", userPoolId)
	// returns value of COGNITO_USER_POOL_ID
	fmt.Println("COGNITO_USER_POOL_ID:", os.Getenv("COGNITO_USER_POOL_ID"))
	UserPoolId = os.Getenv("COGNITO_USER_POOL_ID")

	assert.Equal(userPoolId, UserPoolId)
}